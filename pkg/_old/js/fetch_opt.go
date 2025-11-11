package js

import "strings"

///////////////////////////////////////////////////////////////////////////////
// TYPES

// FetchOpt is a function that configures fetch options.
type FetchOpt func(*fetchopts) error

type fetchopts struct {
	Method      string            // HTTP method (GET, POST, PUT, DELETE, etc.)
	Headers     map[string]string // Request headers
	Body        any               // Request body (string or FormData)
	Mode        string            // Request mode: "cors", "no-cors", "same-origin", "navigate"
	Credentials string            // Credentials: "omit", "same-origin", "include"
	Cache       string            // Cache mode: "default", "no-store", "reload", "no-cache", "force-cache", "only-if-cached"
	Redirect    string            // Redirect mode: "follow", "error", "manual"
	Referrer    string            // Referrer URL or "no-referrer", "client"
	Integrity   string            // Subresource integrity value
	KeepAlive   bool              // Keep connection alive
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func applyOpts(opts ...FetchOpt) (*fetchopts, error) {
	fetchOptions := &fetchopts{
		Headers: make(map[string]string),
	}
	for _, opt := range opts {
		if err := opt(fetchOptions); err != nil {
			return nil, err
		}
	}
	return fetchOptions, nil
}

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

func WithMethod(method string) FetchOpt {
	return func(opts *fetchopts) error {
		opts.Method = strings.TrimSpace(strings.ToUpper(method))
		return nil
	}
}

func WithHeaders(pairs ...string) FetchOpt {
	return func(opts *fetchopts) error {
		for i := 0; i < len(pairs)-1; i += 2 {
			key := strings.TrimSpace(pairs[i])
			value := strings.TrimSpace(pairs[i+1])
			if key != "" {
				opts.Headers[key] = value
			}
		}
		return nil
	}
}

func WithBody(body any) FetchOpt {
	return func(opts *fetchopts) error {
		opts.Body = body
		return nil
	}
}
