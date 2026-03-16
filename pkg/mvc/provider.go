package mvc

import (
	"fmt"
	"net/url"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Provider[T] fetches data from a remote base URL, decodes the response body
// into T using a caller-supplied function, and notifies the listener with
// the result. T can be any type — a struct, a slice, a map, anything the
// decode function can produce.
//
// It has no knowledge of paging, sorting, or views: those concerns belong in
// the ModelController that owns the provider.
//
// Basic usage:
//
//	p := mvc.NewProvider[[]Station](baseURL, func(body []byte) ([]Station, error) {
//	    var r StationsResponse
//	    return r.Root.Stations.Station, json.Unmarshal(body, &r)
//	})
//	p.AddEventListener(func(stations []Station, err error) { … })
//	p.Fetch("stn.aspx", url.Values{"cmd": {"stns"}, "key": {apiKey}})
type Provider[T any] struct {
	base     *url.URL
	decode   func([]byte) (T, error)
	listener func(T, error)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewProvider creates a provider that fetches from base and decodes each
// response body with decode. Returns nil if base is nil.
func NewProvider[T any](base *url.URL, decode func([]byte) (T, error)) *Provider[T] {
	if base == nil || decode == nil {
		return nil
	}
	return &Provider[T]{base: base, decode: decode}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Fetch issues a GET request to base+path?params, decodes the body, and calls
// the listener with (T, error). The call is asynchronous — the listener is
// invoked from the Promise completion callback.
func (p *Provider[T]) Fetch(path string, params url.Values, opts ...js.FetchOption) {
	u := *p.base
	baseQuery := p.base.Query() // save before ResolveReference drops them
	if path != "" {
		ref, err := url.Parse(path)
		if err == nil {
			u = *u.ResolveReference(ref)
		}
	}
	q := baseQuery
	for k, vs := range params {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()

	js.Get(u.String(), opts...).Done(func(val js.Value, err error) {
		if err != nil {
			p.notify(*new(T), err)
			return
		}
		resp := js.ResponseFrom(val)
		if resp == nil {
			p.notify(*new(T), fmt.Errorf("provider: nil response from %s", u.String()))
			return
		}
		resp.Text().Done(func(text js.Value, err error) {
			if err != nil {
				p.notify(*new(T), err)
				return
			}
			result, err := p.decode([]byte(text.String()))
			p.notify(result, err)
		})
	})
}

// AddEventListener registers fn as the single completion listener, replacing
// any previously registered one. fn receives (T, error) after every Fetch.
func (p *Provider[T]) AddEventListener(fn func(T, error)) {
	p.listener = fn
}

// RemoveEventListener clears the completion listener.
func (p *Provider[T]) RemoveEventListener() {
	p.listener = nil
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (p *Provider[T]) notify(result T, err error) {
	if p.listener != nil {
		p.listener(result, err)
	}
}
