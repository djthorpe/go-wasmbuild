package mvc

import (
	"net/url"
	"time"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Provider fetches data from a remote source and notifies registered listeners.
// Use AddEventListener to react to each fetch result (data or error).
type Provider interface {
	// Fetch performs a single fetch and calls all listeners with the response or error.
	Fetch(path string, opts ...js.FetchOption)

	// FetchWithInterval starts periodic fetching at the given interval, calling
	// all listeners on each attempt. Fetches immediately, then repeats.
	// Call Cancel to stop.
	FetchWithInterval(path string, interval time.Duration, opts ...js.FetchOption)

	// Cancel stops any active interval fetch.
	Cancel()

	// AddEventListener registers a listener called on every fetch completion.
	// The response is nil on error; err is nil on success.
	AddEventListener(fn func(*js.FetchResponse, error))
}

// provider is the concrete Provider implementation.
type provider struct {
	base      *url.URL
	listeners []func(*js.FetchResponse, error)
	timer     *js.Timer
}

var _ Provider = (*provider)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewProvider creates a new provider with a base URL.
func NewProvider(base *url.URL) Provider {
	if base == nil {
		return nil
	}
	return &provider{
		base: base,
	}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *provider) AddEventListener(fn func(*js.FetchResponse, error)) {
	p.listeners = append(p.listeners, fn)
}

func (p *provider) Fetch(path string, opts ...js.FetchOption) {
	u := p.base.JoinPath(path)
	js.Fetch(u.String(), opts...).Done(func(v js.Value, err error) {
		if err != nil {
			p.emit(nil, err)
			return
		}
		p.emit(js.ResponseFrom(v), nil)
	})
}

func (p *provider) FetchWithInterval(path string, interval time.Duration, opts ...js.FetchOption) {
	p.Cancel()
	p.timer = js.SetInterval(interval, func() {
		p.Fetch(path, opts...)
	})
}

func (p *provider) Cancel() {
	if p.timer != nil {
		p.timer.Cancel()
		p.timer = nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (p *provider) emit(response *js.FetchResponse, err error) {
	for _, fn := range p.listeners {
		fn(response, err)
	}
}
