package main

import (
	"fmt"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// ServeBroadcaster manages broadcasting messages to SSE clients
type ServeBroadcaster struct {
	mu      sync.Mutex
	clients map[chan ServeMessage]bool
}

// ServeMessage represents a message sent to SSE clients
type ServeMessage struct {
	Type string // "reload" or "error"
	Data string
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewServeBroadcaster() *ServeBroadcaster {
	return &ServeBroadcaster{
		clients: make(map[chan ServeMessage]bool),
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (rb *ServeBroadcaster) register(client chan ServeMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.clients[client] = true
}

func (rb *ServeBroadcaster) unregister(client chan ServeMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	delete(rb.clients, client)
	close(client)
}

func (rb *ServeBroadcaster) error(err error) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	for client := range rb.clients {
		select {
		case client <- ServeMessage{
			Type: "build-error",
			Data: fmt.Sprintf("Compilation error: %v", err),
		}:
		default:
			// Client not ready, skip
		}
	}
}

func (rb *ServeBroadcaster) reload() {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	for client := range rb.clients {
		select {
		case client <- ServeMessage{
			Type: "reload",
			Data: "build succeeded",
		}:
		default:
			// Client not ready, skip
		}
	}
}
