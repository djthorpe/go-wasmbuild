package mvc

import (
	"net/url"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// A provider fetches data from a remote source, either periodically, on demand, or
// in reaction to some event.
type Provider interface {
	Fetch(path string, opts ...js.FetchOption) *js.Promise
}

type provider struct {
}

var _ Provider = (*provider)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new provider with a "base" URL
func NewProvider(url *url.URL) *provider {
	this := new(provider)
	if url == nil {
		return nil
	}

	return this
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Fetch data from the remote source on demand
func (p *provider) Fetch(path string, opts ...js.FetchOption) *js.Promise {
	return nil
}

/*
// Fetch data from the remote source on a scheduled interval
func (p *provider) FetchWithInterval(path string, interval time.Duration, opts ...js.FetchOption) {

}

// Stop any periodic fetching
func (p *provider) Cancel() {

}
*/
