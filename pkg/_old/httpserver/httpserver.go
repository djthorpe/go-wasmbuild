package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type httpserver struct {
	server http.Server
	mux    *http.ServeMux
}

type Middleware interface {
	// Wrap a http handler in a middleware component
	Wrap(http.HandlerFunc) http.HandlerFunc
}

////////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 60 * time.Second
	defaultShutdownTimeout = 60 * time.Second
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func New(opts ...Opt) (*httpserver, error) {
	// Apply options
	defaults := []Opt{
		WithServeMux(http.NewServeMux()),
	}
	opt, err := apply(append(defaults, opts...)...)
	if err != nil {
		return nil, err
	}

	// If address not set, use unused port
	if opt.addr == "" {
		if addr_, err := freePort(); err != nil {
			return nil, err
		} else {
			opt.addr = addr_
		}
	}

	// Return the server
	return &httpserver{
		server: http.Server{
			Addr:         opt.addr,
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			Handler:      opt.mux,
		},
		mux: opt.mux,
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (httpserver *httpserver) String() string {
	type j struct {
		Listen string `json:"listen"`
	}
	data, err := json.MarshalIndent(j{Listen: httpserver.server.Addr}, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// http.Handler

func (httpserver *httpserver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpserver.mux.ServeHTTP(w, r)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Run the http server until cancelled, and return any errors
func (s *httpserver) Run(ctx context.Context) error {
	var wg sync.WaitGroup

	// Start a goroutine to listen for context cancellation
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Wait for context cancellation
		<-ctx.Done()

		// Graceful shutdown with timeout
		shutdown, cancel := context.WithTimeout(context.Background(), defaultShutdownTimeout)
		defer cancel()
		_ = s.server.Shutdown(shutdown)
	}()

	// Don't exit until the goroutine is done
	defer wg.Wait()

	// Start the HTTP server until shutdown
	if err := s.server.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		return nil
	} else {
		return err
	}
}

// Add a handler functiomn and middleware for each request
func (s *httpserver) Handle(path string, handler http.HandlerFunc, middleware ...Middleware) {
	if handler == nil {
		panic("HandleFunc: missing handler")
	}

	// Wrap middleware
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i].Wrap(handler)
	}

	// Register the wrapped handler
	s.mux.HandleFunc(path, handler)
}
