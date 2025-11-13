//go:build js && wasm

package js

import "syscall/js"

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Response wraps a JavaScript Response object
type Response struct {
	Value
}

// NewResponse creates a new Response wrapper from a JavaScript Response value
func NewResponse(value Value) *Response {
	return &Response{Value: value}
}

///////////////////////////////////////////////////////////////////////////////
// RESPONSE METHODS

// Ok returns true if the response was successful (status in range 200-299)
func (r *Response) Ok() bool {
	return r.Get("ok").Bool()
}

// Status returns the HTTP status code of the response
func (r *Response) Status() int {
	return r.Get("status").Int()
}

// StatusText returns the HTTP status message of the response
func (r *Response) StatusText() string {
	return r.Get("statusText").String()
}

// Text returns a Promise that resolves to the response body as text
func (r *Response) Text() *Promise {
	textPromise := r.Call("text")
	return NewPromise(textPromise)
}

// JSON returns a Promise that resolves to the response body parsed as JSON
func (r *Response) JSON() *Promise {
	jsonPromise := r.Call("json")
	return NewPromise(jsonPromise)
}

///////////////////////////////////////////////////////////////////////////////
// FETCH

// Fetch performs a fetch request to the specified URL with optional FetchOpts options.
// Returns a Promise that resolves to a Response Value.
// If there are errors with options, the promise is rejected.
func Fetch(url string, opts ...FetchOpt) *Promise {
	opt, err := applyOpts(opts...)
	if err != nil {
		// Return a rejected promise
		return NewPromiseFromTask(func(resolve func(any), reject func(error)) {
			reject(err)
		})
	}

	// Call the JavaScript fetch function - it returns a promise that resolves to a Response
	fetchPromise := js.Global().Call("fetch", url, opt.Object())

	// Just wrap and return the JavaScript promise directly
	return NewPromise(fetchPromise)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (opts *fetchopts) Object() Value {
	fetchopts := NewObject()

	if opts.Method != "" {
		fetchopts.Set("method", opts.Method)
	}

	if len(opts.Headers) > 0 {
		headers := NewObject()
		for key, value := range opts.Headers {
			headers.Set(key, value)
		}
		fetchopts.Set("headers", headers)
	}

	if opts.Body != nil {
		fetchopts.Set("body", opts.Body)
	}

	if opts.Mode != "" {
		fetchopts.Set("mode", opts.Mode)
	}

	if opts.Credentials != "" {
		fetchopts.Set("credentials", opts.Credentials)
	}

	if opts.Cache != "" {
		fetchopts.Set("cache", opts.Cache)
	}

	if opts.Redirect != "" {
		fetchopts.Set("redirect", opts.Redirect)
	}

	if opts.Referrer != "" {
		fetchopts.Set("referrer", opts.Referrer)
	}

	if opts.Integrity != "" {
		fetchopts.Set("integrity", opts.Integrity)
	}

	if opts.KeepAlive {
		fetchopts.Set("keepalive", true)
	}

	return fetchopts
}
