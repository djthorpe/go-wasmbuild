package httpserver

import (
	"fmt"
	"net"
	"net/http"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

type opts struct {
	addr string
	mux  *http.ServeMux
}

type Opt func(*opts) error

////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func apply(opt ...Opt) (*opts, error) {
	o := new(opts)
	for _, fn := range opt {
		if err := fn(o); err != nil {
			return nil, err
		}
	}
	return o, nil
}

////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func WithListenAddr(addr string) Opt {
	return func(o *opts) error {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return err
		} else {
			o.addr = net.JoinHostPort(host, port)
		}
		return nil
	}
}

func WithServeMux(mux *http.ServeMux) Opt {
	return func(o *opts) error {
		if mux == nil {
			return fmt.Errorf("WithServeMux: missing mux")
		}
		o.mux = mux
		return nil
	}
}
