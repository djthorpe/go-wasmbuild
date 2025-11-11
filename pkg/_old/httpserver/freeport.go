package httpserver

import (
	"net"
)

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// FreePort returns an available local address, in case it's
// not specified
func freePort() (string, error) {
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	addr := listen.Addr().String()
	if err := listen.Close(); err != nil {
		return "", err
	}
	return addr, nil
}
