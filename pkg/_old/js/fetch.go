//go:build !(js && wasm)

package js

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Response wraps a JavaScript Response object (mock for non-WASM builds)
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
	// In mock mode, always return true
	return true
}

// Status returns the HTTP status code of the response
func (r *Response) Status() int {
	// In mock mode, return 200
	return 200
}

// StatusText returns the HTTP status message of the response
func (r *Response) StatusText() string {
	// In mock mode, return OK
	return "OK"
}

// Text returns a Promise that resolves to the response body as text
func (r *Response) Text() *Promise {
	return NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		// In mock mode, resolve with empty string
		resolve("")
	})
}

// JSON returns a Promise that resolves to the response body parsed as JSON
func (r *Response) JSON() *Promise {
	return NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		// In mock mode, resolve with empty object
		resolve(NewObject())
	})
}

///////////////////////////////////////////////////////////////////////////////
// FETCH

// Fetch performs a fetch request to the specified URL with optional FetchOpts options.
// Returns a Promise that resolves to a Response Value (mock for non-WASM builds).
// If there are errors with options, the promise is rejected.
func Fetch(url string, opts ...FetchOpt) *Promise {
	return NewPromiseFromTask(func(resolve func(any), reject func(error)) {
		opt, err := applyOpts(opts...)
		if err != nil {
			reject(err)
			return
		}

		// In mock mode, just resolve immediately with a mock Response Value
		response := NewObject()
		_ = opt // Use opt to avoid unused variable warning
		resolve(response)
	})
}
