//go:build js

package mapbox

import (
	"syscall/js"
)

// MapboxGL provides access to the global mapboxgl namespace.
type MapboxGL struct {
	value js.Value
}

// New returns a new MapboxGL instance that wraps the global mapboxgl object.
// Returns nil if mapboxgl is not available in the global scope.
func New() *MapboxGL {
	mapboxgl := js.Global().Get("mapboxgl")
	if mapboxgl.IsUndefined() || mapboxgl.IsNull() {
		return nil
	}

	return &MapboxGL{
		value: mapboxgl,
	}
}

// JSValue returns the underlying js.Value.
func (m *MapboxGL) JSValue() js.Value {
	return m.value
}

// Version returns the version of Mapbox GL JS being used.
func (m *MapboxGL) Version() string {
	version := m.value.Get("version")
	if version.IsUndefined() {
		return ""
	}
	return version.String()
}

// Supported returns true if WebGL is supported by the browser.
func (m *MapboxGL) Supported(options ...map[string]interface{}) bool {
	if len(options) > 0 {
		return m.value.Call("supported", js.ValueOf(options[0])).Bool()
	}
	return m.value.Call("supported").Bool()
}

// SetRTLTextPlugin sets the URL for the right-to-left text plugin.
// This is used to support languages that use right-to-left scripts.
func (m *MapboxGL) SetRTLTextPlugin(pluginURL string, callback func(error js.Value)) {
	if callback != nil {
		callbackFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			var err js.Value
			if len(args) > 0 {
				err = args[0]
			}
			callback(err)
			return nil
		})
		defer callbackFunc.Release()
		m.value.Call("setRTLTextPlugin", pluginURL, callbackFunc)
	} else {
		m.value.Call("setRTLTextPlugin", pluginURL)
	}
}

// GetRTLTextPluginStatus returns the status of the right-to-left text plugin.
// Returns one of: "unavailable", "loading", "loaded", "error"
func (m *MapboxGL) GetRTLTextPluginStatus() string {
	status := m.value.Call("getRTLTextPluginStatus")
	if status.IsUndefined() {
		return "unavailable"
	}
	return status.String()
}

// ClearStorage clears Mapbox GL JS's web storage (localStorage and sessionStorage).
func (m *MapboxGL) ClearStorage(callback func()) {
	if callback != nil {
		callbackFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			callback()
			return nil
		})
		defer callbackFunc.Release()
		m.value.Call("clearStorage", callbackFunc)
	} else {
		m.value.Call("clearStorage")
	}
}

// WorkerCount gets or sets the number of web workers used for processing map data.
func (m *MapboxGL) WorkerCount() int {
	return m.value.Get("workerCount").Int()
}

// SetWorkerCount sets the number of web workers used for processing map data.
func (m *MapboxGL) SetWorkerCount(count int) {
	m.value.Set("workerCount", count)
}

// MaxParallelImageRequests gets or sets the maximum number of parallel image requests.
func (m *MapboxGL) MaxParallelImageRequests() int {
	return m.value.Get("maxParallelImageRequests").Int()
}

// SetMaxParallelImageRequests sets the maximum number of parallel image requests.
func (m *MapboxGL) SetMaxParallelImageRequests(count int) {
	m.value.Set("maxParallelImageRequests", count)
}

// AccessToken gets the Mapbox access token.
func (m *MapboxGL) AccessToken() string {
	token := m.value.Get("accessToken")
	if token.IsUndefined() {
		return ""
	}
	return token.String()
}

// SetAccessToken sets the Mapbox access token.
func (m *MapboxGL) SetAccessToken(token string) {
	m.value.Set("accessToken", token)
}

// BaseApiURL gets the base API URL for Mapbox services.
func (m *MapboxGL) BaseApiURL() string {
	url := m.value.Get("baseApiUrl")
	if url.IsUndefined() {
		return ""
	}
	return url.String()
}

// SetBaseApiURL sets the base API URL for Mapbox services.
func (m *MapboxGL) SetBaseApiURL(url string) {
	m.value.Set("baseApiUrl", url)
}
