package mvc

import (
	"encoding/json"
	"net/url"
	"time"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// JSONProvider fetches JSON data from a remote source, encoding and decoding
// values of type T. Use AddEventListener to receive decoded results.
type JSONProvider[T any] interface {
	// Get fetches from the given path and decodes the response as T.
	Get(path string, opts ...js.FetchOption)

	// GetWithInterval calls Get immediately and then repeatedly at the given
	// interval. Call Cancel to stop.
	GetWithInterval(path string, interval time.Duration, opts ...js.FetchOption)

	// Cancel stops any active interval.
	Cancel()

	// Post marshals body as JSON, posts to path, and decodes the response as T.
	Post(path string, body T, opts ...js.FetchOption)

	// Put marshals body as JSON, puts to path, and decodes the response as T.
	Put(path string, body T, opts ...js.FetchOption)

	// Patch marshals body as JSON, patches path, and decodes the response as T.
	Patch(path string, body T, opts ...js.FetchOption)

	// Delete sends a DELETE to path and decodes any response body as T.
	// If the server returns no body (e.g. 204 No Content), the listener
	// is called with the zero value of T and a nil error.
	Delete(path string, opts ...js.FetchOption)

	// AddEventListener registers a listener called on every completed request.
	// value is the zero value of T on error; err is nil on success.
	AddEventListener(fn func(T, error))
}

// jsonProvider is the concrete JSONProvider[T] implementation.
type jsonProvider[T any] struct {
	base      *url.URL
	listeners []func(T, error)
	timer     *js.Timer
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewJSONProvider creates a new JSONProvider with the given base URL.
// Returns nil if base is nil.
func NewJSONProvider[T any](base *url.URL) JSONProvider[T] {
	if base == nil {
		return nil
	}
	return &jsonProvider[T]{base: base}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *jsonProvider[T]) AddEventListener(fn func(T, error)) {
	p.listeners = append(p.listeners, fn)
}

func (p *jsonProvider[T]) Get(path string, opts ...js.FetchOption) {
	p.do(path, append([]js.FetchOption{js.WithMethod("GET")}, opts...)...)
}

func (p *jsonProvider[T]) GetWithInterval(path string, interval time.Duration, opts ...js.FetchOption) {
	p.Cancel()
	p.timer = js.SetInterval(interval, func() {
		p.Get(path, opts...)
	})
}

func (p *jsonProvider[T]) Cancel() {
	if p.timer != nil {
		p.timer.Cancel()
		p.timer = nil
	}
}

func (p *jsonProvider[T]) Post(path string, body T, opts ...js.FetchOption) {
	opts, err := withJSONBody(body, opts)
	if err != nil {
		p.emitError(err)
		return
	}
	p.do(path, append([]js.FetchOption{js.WithMethod("POST")}, opts...)...)
}

func (p *jsonProvider[T]) Put(path string, body T, opts ...js.FetchOption) {
	opts, err := withJSONBody(body, opts)
	if err != nil {
		p.emitError(err)
		return
	}
	p.do(path, append([]js.FetchOption{js.WithMethod("PUT")}, opts...)...)
}

func (p *jsonProvider[T]) Patch(path string, body T, opts ...js.FetchOption) {
	opts, err := withJSONBody(body, opts)
	if err != nil {
		p.emitError(err)
		return
	}
	p.do(path, append([]js.FetchOption{js.WithMethod("PATCH")}, opts...)...)
}

func (p *jsonProvider[T]) Delete(path string, opts ...js.FetchOption) {
	p.do(path, append([]js.FetchOption{js.WithMethod("DELETE")}, opts...)...)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// do performs the fetch, decodes the JSON response body as T, and emits.
func (p *jsonProvider[T]) do(path string, opts ...js.FetchOption) {
	u := p.base.JoinPath(path)
	js.Fetch(u.String(), opts...).Done(func(v js.Value, err error) {
		if err != nil {
			p.emitError(err)
			return
		}
		resp := js.ResponseFrom(v)
		resp.Text().Done(func(tv js.Value, err error) {
			if err != nil {
				p.emitError(err)
				return
			}
			body := tv.String()
			if body == "" {
				// No body (e.g. 204 No Content) — emit zero value
				p.emit(*new(T), nil)
				return
			}
			var result T
			if err := json.Unmarshal([]byte(body), &result); err != nil {
				p.emitError(err)
				return
			}
			p.emit(result, nil)
		})
	})
}

func (p *jsonProvider[T]) emit(value T, err error) {
	for _, fn := range p.listeners {
		fn(value, err)
	}
}

func (p *jsonProvider[T]) emitError(err error) {
	p.emit(*new(T), err)
}

// withJSONBody marshals body as JSON and prepends a WithJSON option.
func withJSONBody[T any](body T, opts []js.FetchOption) ([]js.FetchOption, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return append([]js.FetchOption{js.WithJSON(string(b))}, opts...), nil
}
