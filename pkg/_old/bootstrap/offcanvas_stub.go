//go:build !js || !wasm

package bootstrap

// offcanvas stub for non-WASM builds (testing only)
type offcanvas struct {
	component
}
