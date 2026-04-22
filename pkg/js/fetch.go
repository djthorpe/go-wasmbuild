//go:build !(js && wasm)

package js

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

// FetchResponse wraps an HTTP response.
type FetchResponse struct {
	status     int
	statusText string
	ok         bool
	headers    map[string]string
	bodyData   []byte
}

// FetchOption is a function that configures a FetchRequest.
type FetchOption func(*FetchRequest)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Fetch creates a new HTTP request and returns a Promise that resolves to a *FetchResponse.
func Fetch(url string, opts ...FetchOption) *Promise {
	return NewPromise(func() (Value, error) {
		req := &FetchRequest{
			url:     url,
			method:  "GET",
			headers: make(map[string]string),
		}

		for _, opt := range opts {
			opt(req)
		}

		// Build the HTTP request
		var bodyReader io.Reader
		if req.body != nil {
			switch b := req.body.(type) {
			case string:
				bodyReader = strings.NewReader(b)
			case []byte:
				bodyReader = strings.NewReader(string(b))
			case io.Reader:
				bodyReader = b
			default:
				bodyReader = strings.NewReader(fmt.Sprintf("%v", b))
			}
		}

		httpReq, err := http.NewRequest(req.method, req.url, bodyReader)
		if err != nil {
			return Undefined(), err
		}

		for k, v := range req.headers {
			httpReq.Header.Set(k, v)
		}

		// Execute the request
		client := &http.Client{}
		httpResp, err := client.Do(httpReq)
		if err != nil {
			return Undefined(), err
		}
		defer httpResp.Body.Close()

		// Read the body
		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return Undefined(), err
		}

		// Build response headers
		respHeaders := make(map[string]string)
		for k, v := range httpResp.Header {
			if len(v) > 0 {
				respHeaders[k] = v[0]
			}
		}

		response := &FetchResponse{
			status:     httpResp.StatusCode,
			statusText: httpResp.Status,
			ok:         httpResp.StatusCode >= 200 && httpResp.StatusCode < 300,
			headers:    respHeaders,
			bodyData:   bodyBytes,
		}

		return ValueOf(response), nil
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
// RESPONSE METHODS (stubs for native build)

// OK returns true if the response status is in the 200-299 range.
func (r *FetchResponse) OK() bool {
	return r.ok
}

// Status returns the HTTP status code.
func (r *FetchResponse) Status() int {
	return r.status
}

// StatusText returns the HTTP status text.
func (r *FetchResponse) StatusText() string {
	return r.statusText
}

// Headers returns the response headers.
func (r *FetchResponse) Headers() map[string]string {
	return r.headers
}

// JSON returns a Promise that resolves to the parsed JSON body.
// Note: In native Go builds, this method is not implemented and will return an error.
// Use Text() followed by encoding/json.Unmarshal instead.
func (r *FetchResponse) JSON() *Promise {
	return NewPromise(func() (Value, error) {
		return Undefined(), fmt.Errorf("FetchResponse.JSON is only available in WASM builds")
	})
}

// Text returns a Promise that resolves to the response body as text.
func (r *FetchResponse) Text() *Promise {
	return NewPromise(func() (Value, error) {
		return ValueOf(string(r.bodyData)), nil
	})
}

// Blob returns a Promise that resolves to the response body as a Blob.
// Note: In native Go, this returns the raw body as a []byte wrapped in a Value,
// since the Blob type is browser-specific and not available here.
func (r *FetchResponse) Blob() *Promise {
	return NewPromise(func() (Value, error) {
		return ValueOf(r.bodyData), nil
	})
}

// ResponseFrom extracts a *FetchResponse from a Value.
// In native Go, the Value should wrap a *FetchResponse.
// Returns nil if the Value does not contain a *FetchResponse.
func ResponseFrom(v Value) *FetchResponse {
	if resp, ok := v.v.(*FetchResponse); ok {
		return resp
	}
	return nil
}
