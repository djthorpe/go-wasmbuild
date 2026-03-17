//go:build js && wasm

package js

import (
	"fmt"
	"net/url"
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// FetchRequest represents an HTTP request configuration.
type FetchRequest struct {
	url     string
	method  string
	headers map[string]string
	body    any
}

// FetchResponse wraps a JavaScript Response object.
type FetchResponse struct {
	jsResponse Value
}

// FetchOption is a function that configures a FetchRequest.
type FetchOption func(*FetchRequest)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Fetch creates a new HTTP request and returns a Promise that resolves to a *FetchResponse.
func Fetch(url string, opts ...FetchOption) *Promise {
	req := &FetchRequest{
		url:     url,
		method:  "GET",
		headers: make(map[string]string),
	}

	for _, opt := range opts {
		opt(req)
	}

	// Build JS fetch options
	jsOpts := js.Global().Get("Object").New()
	jsOpts.Set("method", req.method)

	if len(req.headers) > 0 {
		headers := js.Global().Get("Object").New()
		for k, v := range req.headers {
			headers.Set(k, v)
		}
		jsOpts.Set("headers", headers)
	}

	if req.body != nil {
		jsOpts.Set("body", req.body)
	}

	// Call fetch and wrap the JS promise
	jsPromise := js.Global().Call("fetch", req.url, jsOpts)

	return FromJSPromise(jsPromise).Then(func(value Value) (Value, error) {
		// Check response.ok - fetch only rejects on network errors
		if !value.Get("ok").Bool() {
			return Undefined(), fmt.Errorf("HTTP %d: %s",
				value.Get("status").Int(),
				value.Get("statusText").String())
		}
		return value, nil
	})
}

// Get is a convenience method for Fetch with GET method.
func Get(url string, opts ...FetchOption) *Promise {
	return Fetch(url, append([]FetchOption{WithMethod("GET")}, opts...)...)
}

// Post is a convenience method for Fetch with POST method.
func Post(url string, body any, opts ...FetchOption) *Promise {
	return Fetch(url, append([]FetchOption{WithMethod("POST"), WithBody(body)}, opts...)...)
}

// Put is a convenience method for Fetch with PUT method.
func Put(url string, body any, opts ...FetchOption) *Promise {
	return Fetch(url, append([]FetchOption{WithMethod("PUT"), WithBody(body)}, opts...)...)
}

// Delete is a convenience method for Fetch with DELETE method.
func Delete(url string, opts ...FetchOption) *Promise {
	return Fetch(url, append([]FetchOption{WithMethod("DELETE")}, opts...)...)
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithMethod sets the HTTP method for the request.
func WithMethod(method string) FetchOption {
	return func(r *FetchRequest) {
		r.method = method
	}
}

// WithHeader adds a header to the request.
func WithHeader(key, value string) FetchOption {
	return func(r *FetchRequest) {
		if r.headers == nil {
			r.headers = make(map[string]string)
		}
		r.headers[key] = value
	}
}

// WithHeaders sets multiple headers on the request.
func WithHeaders(headers map[string]string) FetchOption {
	return func(r *FetchRequest) {
		if r.headers == nil {
			r.headers = make(map[string]string)
		}
		for k, v := range headers {
			r.headers[k] = v
		}
	}
}

// WithBody sets the request body.
func WithBody(body any) FetchOption {
	return func(r *FetchRequest) {
		r.body = body
	}
}

// WithQuery appends the given query parameters to the request URL.
// Existing parameters in the URL are preserved; new keys are added.
func WithQuery(params url.Values) FetchOption {
	return func(r *FetchRequest) {
		if len(params) == 0 {
			return
		}
		u, err := url.Parse(r.url)
		if err != nil {
			return
		}
		q := u.Query()
		for k, vs := range params {
			for _, v := range vs {
				q.Add(k, v)
			}
		}
		u.RawQuery = q.Encode()
		r.url = u.String()
	}
}

// WithJSON sets the request body as JSON and adds Content-Type header.
func WithJSON(body string) FetchOption {
	return func(r *FetchRequest) {
		if r.headers == nil {
			r.headers = make(map[string]string)
		}
		r.headers["Content-Type"] = "application/json"
		r.body = body
	}
}

///////////////////////////////////////////////////////////////////////////////
// RESPONSE METHODS

// ResponseFrom creates a FetchResponse from a js.Value.
func ResponseFrom(jsResponse Value) *FetchResponse {
	return &FetchResponse{jsResponse: jsResponse}
}

// OK returns true if the response status is in the 200-299 range.
func (r *FetchResponse) OK() bool {
	return r.jsResponse.Get("ok").Bool()
}

// Status returns the HTTP status code.
func (r *FetchResponse) Status() int {
	return r.jsResponse.Get("status").Int()
}

// StatusText returns the HTTP status text.
func (r *FetchResponse) StatusText() string {
	return r.jsResponse.Get("statusText").String()
}

// Headers returns the response headers as a map.
func (r *FetchResponse) Headers() map[string]string {
	headers := make(map[string]string)
	jsHeaders := r.jsResponse.Get("headers")

	// Use forEach to iterate over the Headers object
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) >= 2 {
			value := args[0].String()
			key := args[1].String()
			headers[key] = value
		}
		return nil
	})
	defer cb.Release()

	jsHeaders.Call("forEach", cb)
	return headers
}

// JSON returns a Promise that resolves to the parsed JSON body.
func (r *FetchResponse) JSON() *Promise {
	return FromJSPromise(r.jsResponse.Call("json"))
}

// Text returns a Promise that resolves to the response body as text.
func (r *FetchResponse) Text() *Promise {
	return FromJSPromise(r.jsResponse.Call("text"))
}

// Blob returns a Promise that resolves to the response body as a Blob.
func (r *FetchResponse) Blob() *Promise {
	return FromJSPromise(r.jsResponse.Call("blob"))
}
