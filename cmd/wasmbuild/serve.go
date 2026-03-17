package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type ServeCmd struct {
	BuildPath
	Listen string `default:"localhost:9090" help:"Address to listen on (e.g., localhost:9090 or 0.0.0.0:9090)"`
}

type ServeContext struct {
	DepContext
	Listen string `json:"listen"`
	Watch  bool   `json:"watch"`

	// Broadcast notifications to clients
	broadcaster *ServeBroadcaster `json:"-"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// ServeContext creates a ServeContext from a DepContext, returning all the
// information needed to serve a WASM application.
func (d DepContext) ServeContext(ctx *Context, listen string, watch bool) (*ServeContext, error) {
	// Return the ServeContext
	return &ServeContext{
		DepContext: d,
		Listen:     listen,
		Watch:      watch,
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *ServeCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveFile(ctx.Config, c.Path)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create a build context from the configuration
	buildContext, err := config.BuildContext(ctx, c.Path, "", true)
	if err != nil {
		return err
	}

	// Create a dependency context from the build context
	dep, err := buildContext.DepContext(ctx)
	if err != nil {
		return err
	}

	// Create the server context from the configuration
	serveContext, err := dep.ServeContext(ctx, c.Listen, true)
	if err != nil {
		return err
	}

	// Compile the .wasm file
	file, err := serveContext.CompileExec(ctx)
	if err != nil {
		return err
	} else {
		serveContext.wasm = file
	}

	// Log the serve context
	ctx.log.Info(serveContext)

	// Start the server
	return serveContext.Serve(ctx,
		serveContext.WasmExecJS,
		serveContext.WasmExecHTML,
		serveContext.FavIcon,
	)
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *ServeContext) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (c *ServeContext) Serve(ctx *Context, files ...*File) error {
	handler := http.NewServeMux()

	// Serve files
	for _, f := range files {
		if f != nil {
			handler.Handle(f.URL(), f.Handler())
		}
	}

	// Serve assets
	for _, asset := range c.Assets {
		if filepath.IsAbs(asset) == false {
			asset = filepath.Join(c.Path, asset)
		}
		if info, err := os.Stat(asset); err != nil {
			return err
		} else if info.Mode().IsDir() {
			prefix := "/" + filepath.Base(asset) + "/"
			handler.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(asset))))
		} else if info.Mode().IsRegular() {
			file, err := NewFileFromSource(asset, filepath.Base(asset))
			if err != nil {
				return err
			} else {
				handler.Handle(file.URL(), file.Handler())
			}
		}
	}

	// Server notify handler
	if c.Watch {
		c.broadcaster = NewServeBroadcaster()
		handler.HandleFunc("/_notify", c.NotifyHandler)
	}

	// WASM file handler
	handler.HandleFunc(c.wasm.URL(), func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/wasm")
		w.Write(c.wasm.Data)
	})

	// Output listening info
	url, err := url.Parse(fmt.Sprintf("http://%s%s", c.Listen, c.WasmExecHTML.URL()))
	if err != nil {
		return err
	}
	fmt.Println(url.String())

	// If watch flag is set, we build a dependency watcher
	var wg sync.WaitGroup
	if c.Watch {
		// Watch for dependency changes
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.DepContext.Run(ctx.ctx); err != nil {
				ctx.log.Error(err)
				ctx.cancel()
			}
		}()

		// Respond to modification events
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.ctx.Done():
					return
				case <-c.DepContext.modified:
					// Re-compile the .wasm file
					if file, err := c.CompileExec(ctx); err != nil {
						c.broadcaster.error(err)
						ctx.log.Error(err)
					} else {
						c.wasm = file
						c.broadcaster.reload()
					}
				}
			}
		}()

	}

	// Start HTTP server
	server := &http.Server{
		Addr:    c.Listen,
		Handler: logging(handler, ctx.log),
	}

	// Start server in goroutine
	wg.Add(1)
	serverErr := make(chan error, 1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.ctx.Done():
		ctx.log.Info("Shutting down server...")
		// Gracefully shutdown the server
		if err := server.Close(); err != nil {
			ctx.log.Error("Error shutting down server: ", err)
		}
	case err := <-serverErr:
		return err
	}

	// Wait for all go-routines to end
	wg.Wait()

	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// HANDLERS

// notify handler
func (c *ServeContext) NotifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// We need to be able to flush the data
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Create a client channel and register it
	notify := make(chan ServeMessage)
	c.broadcaster.register(notify)
	defer c.broadcaster.unregister(notify)

	// Send initial connection message
	fmt.Fprintf(w, "event: connected\ndata: connected\n\n")
	flusher.Flush()

	// Keep connection alive and wait for reload signals
	for {
		select {
		case msg := <-notify:
			fmt.Println("Notify client:", msg.Type)
			switch msg.Type {
			case "reload":
				fmt.Fprintf(w, "event: reload\ndata: reload\n\n")
			case "build-error":
				fmt.Fprintf(w, "event: build-error\n")
				// For multi-line error messages, prefix each line with "data: "
				lines := strings.Split(msg.Data, "\n")
				for _, line := range lines {
					fmt.Fprintf(w, "data: %s\n", line)
				}
				fmt.Fprintf(w, "\n")
			}
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}
