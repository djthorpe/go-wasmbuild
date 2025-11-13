package httpserver

import (
	"log/slog"
	"net/http"
)

///////////////////////////////////////////////////////////////////////////////
// LOGGER

type logger struct {
	*slog.Logger
}

var _ Middleware = (*logger)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a logger middleware wrapper
func NewLogger(parent *slog.Logger) *logger {
	return &logger{parent}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (logger *logger) Wrap(parent http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("REQUEST")
		parent(w, r)
		logger.Info("RESPONSE")
	}
}
