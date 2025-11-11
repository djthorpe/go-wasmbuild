package httpserver

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestNewDefaults(t *testing.T) {
	srv, err := New(WithListenAddr(":0"))
	if err != nil {
		t.Fatal("Error:", err)
	}

	if srv.server.ReadTimeout != defaultReadTimeout {
		t.Fatalf("expected read timeout %v, got %v", defaultReadTimeout, srv.server.ReadTimeout)
	}
	if srv.server.WriteTimeout != defaultWriteTimeout {
		t.Fatalf("expected write timeout %v, got %v", defaultWriteTimeout, srv.server.WriteTimeout)
	}
	if srv.server.Handler == nil {
		t.Fatal("expected default handler to be initialised")
	}
}

func TestNewCustomRouter(t *testing.T) {
	mux := http.NewServeMux()
	srv, err := New(WithListenAddr(":0"), WithServeMux(mux))
	if err != nil {
		t.Fatal("Error:", err)
	}
	if srv.server.Handler != mux {
		t.Fatalf("expected custom mux to be used, got %T", srv.server.Handler)
	}
}

func TestRunGracefulShutdown(t *testing.T) {
	addr, err := freePort()
	if err != nil {
		t.Fatal("Error:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv, err := New(WithListenAddr(addr), WithServeMux(mux))
	if err != nil {
		t.Fatal("Error:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Run(ctx)
	}()

	client := &http.Client{Timeout: 200 * time.Millisecond}
	deadline := time.NewTimer(time.Second)
	defer deadline.Stop()

	// Poll until the server is accepting requests, otherwise fail early.
	for {
		resp, err := client.Get("http://" + addr + "/ping")
		if err == nil {
			_ = resp.Body.Close()
			break
		}
		select {
		case <-deadline.C:
			t.Fatalf("server did not become ready: %v", err)
		default:
			time.Sleep(25 * time.Millisecond)
		}
	}

	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("expected graceful shutdown, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for server shutdown")
	}
}

func TestRunListenError(t *testing.T) {
	srv, err := New(WithListenAddr("bad-address"))
	if err != nil {
		t.Fatal("Error:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Run(ctx)
	}()

	// Allow Run to attempt ListenAndServe, then cancel to release the waiter.
	select {
	case err := <-errCh:
		// If it returns immediately we expect a listen error.
		if err == nil {
			t.Fatal("expected error from invalid listen address")
		}
		if errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("expected listen error, got %v", err)
		}
		return
	case <-time.After(50 * time.Millisecond):
		cancel()
	}

	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("expected error from invalid listen address")
		}
		if errors.Is(err, http.ErrServerClosed) {
			t.Fatalf("expected listen error, got %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for listen error")
	}
}
