package httpserver_test

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/httpserver"
)

func TestRunLogger(t *testing.T) {
	server, err := httpserver.New()
	if err != nil {
		t.Fatal("Error:", err)
	}

	// Add a logger
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}))
	middleware := httpserver.NewLogger(logger)

	// Add ping handler
	server.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	}, middleware)

	// Perform a ping
	w := httptest.NewRecorder()
	server.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/", nil))
	if resp := w.Body.String(); resp != "ping" {
		t.Fatalf(`Expected "ping", got %q`, resp)
	}
}
