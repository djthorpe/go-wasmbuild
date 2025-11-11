//go:build js

package mapbox

import (
	"syscall/js"
)

// EventListener represents a callback function for map events.
type EventListener func(event js.Value)

// EventOptions represents options for event handling.
type EventOptions struct {
	value js.Value
}

// NewEventOptions creates a new EventOptions object.
func NewEventOptions() *EventOptions {
	return &EventOptions{
		value: js.Global().Get("Object").New(),
	}
}

// JSValue returns the underlying js.Value.
func (eo *EventOptions) JSValue() js.Value {
	return eo.value
}

// SetOnce sets whether the listener should be called only once.
func (eo *EventOptions) SetOnce(once bool) *EventOptions {
	eo.value.Set("once", once)
	return eo
}

// SetPassive sets whether the listener should be passive.
func (eo *EventOptions) SetPassive(passive bool) *EventOptions {
	eo.value.Set("passive", passive)
	return eo
}

// Event type constants for Mapbox GL JS.
const (
	// Map load events
	EventLoad                 = "load"
	EventIdle                 = "idle"
	EventRemove               = "remove"
	EventRender               = "render"
	EventResize               = "resize"
	EventWebGLContextLost     = "webglcontextlost"
	EventWebGLContextRestored = "webglcontextrestored"

	// Map movement events
	EventMoveStart = "movestart"
	EventMove      = "move"
	EventMoveEnd   = "moveend"

	// Drag events
	EventDragStart = "dragstart"
	EventDrag      = "drag"
	EventDragEnd   = "dragend"

	// Zoom events
	EventZoomStart = "zoomstart"
	EventZoom      = "zoom"
	EventZoomEnd   = "zoomend"

	// Rotation events
	EventRotateStart = "rotatestart"
	EventRotate      = "rotate"
	EventRotateEnd   = "rotateend"

	// Pitch events
	EventPitchStart = "pitchstart"
	EventPitch      = "pitch"
	EventPitchEnd   = "pitchend"

	// Mouse events
	EventMouseDown   = "mousedown"
	EventMouseUp     = "mouseup"
	EventMouseOver   = "mouseover"
	EventMouseMove   = "mousemove"
	EventMouseOut    = "mouseout"
	EventContextMenu = "contextmenu"

	// Click events
	EventClick    = "click"
	EventDblClick = "dblclick"

	// Touch events
	EventTouchStart  = "touchstart"
	EventTouchEnd    = "touchend"
	EventTouchMove   = "touchmove"
	EventTouchCancel = "touchcancel"

	// Wheel events
	EventWheel = "wheel"

	// Data events
	EventData       = "data"
	EventSourceData = "sourcedata"
	EventStyleData  = "styledata"

	// Error events
	EventError = "error"

	// Style events
	EventStyleLoad         = "style.load"
	EventStyleImageMissing = "styleimagemissing"

	// Source events
	EventSourceDataLoading = "sourcedataloading"

	// Terrain events
	EventTerrain = "terrain"

	// Cooperation events
	EventCooperativeGestureStart = "cooperativegesturestart"
	EventCooperativeGestureEnd   = "cooperativegestureend"
)

// MapEventHandling provides event handling methods for the Map.
// This is embedded in the Map struct to provide event functionality.
type MapEventHandling struct {
	value js.Value
}

// On adds an event listener to the map.
func (m *Map) On(eventType string, listener EventListener, options ...EventOptions) *Map {
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		return nil
	})

	// Store the callback to prevent garbage collection
	// In a real implementation, you'd want to manage these references properly
	if len(options) > 0 {
		m.value.Call("on", eventType, callback, options[0].JSValue())
	} else {
		m.value.Call("on", eventType, callback)
	}
	return m
}

// Once adds an event listener that will be called only once.
func (m *Map) Once(eventType string, listener EventListener) *Map {
	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		// Release the callback after it's called once
		callback.Release()
		return nil
	})

	m.value.Call("once", eventType, callback)
	return m
}

// Off removes an event listener from the map.
// If no listener is provided, removes all listeners for the event type.
func (m *Map) Off(eventType string, listener ...js.Func) *Map {
	if len(listener) > 0 {
		m.value.Call("off", eventType, listener[0])
	} else {
		m.value.Call("off", eventType)
	}
	return m
}

// Fire triggers an event on the map.
func (m *Map) Fire(eventType string, properties ...js.Value) *Map {
	if len(properties) > 0 {
		m.value.Call("fire", eventType, properties[0])
	} else {
		m.value.Call("fire", eventType)
	}
	return m
}

// ListenersCount returns the number of listeners for a given event type.
func (m *Map) ListenersCount(eventType string) int {
	return m.value.Call("listens", eventType).Int()
}

// HasListeners returns true if there are listeners for the given event type.
func (m *Map) HasListeners(eventType string) bool {
	return m.ListenersCount(eventType) > 0
}

// Layer event handling methods

// OnLayer adds an event listener for a specific layer.
func (m *Map) OnLayer(eventType string, layerId string, listener EventListener) *Map {
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		return nil
	})

	m.value.Call("on", eventType, layerId, callback)
	return m
}

// OnceLayer adds an event listener for a specific layer that will be called only once.
func (m *Map) OnceLayer(eventType string, layerId string, listener EventListener) *Map {
	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		callback.Release()
		return nil
	})

	m.value.Call("once", eventType, layerId, callback)
	return m
}

// OffLayer removes an event listener from a specific layer.
func (m *Map) OffLayer(eventType string, layerId string, listener ...js.Func) *Map {
	if len(listener) > 0 {
		m.value.Call("off", eventType, layerId, listener[0])
	} else {
		m.value.Call("off", eventType, layerId)
	}
	return m
}

// Source event handling methods

// OnSource adds an event listener for a specific source.
func (m *Map) OnSource(eventType string, sourceId string, listener EventListener) *Map {
	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		return nil
	})

	m.value.Call("on", eventType, sourceId, callback)
	return m
}

// OnceSource adds an event listener for a specific source that will be called only once.
func (m *Map) OnceSource(eventType string, sourceId string, listener EventListener) *Map {
	var callback js.Func
	callback = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) > 0 {
			listener(args[0])
		}
		callback.Release()
		return nil
	})

	m.value.Call("once", eventType, sourceId, callback)
	return m
}

// OffSource removes an event listener from a specific source.
func (m *Map) OffSource(eventType string, sourceId string, listener ...js.Func) *Map {
	if len(listener) > 0 {
		m.value.Call("off", eventType, sourceId, listener[0])
	} else {
		m.value.Call("off", eventType, sourceId)
	}
	return m
}

// Convenience methods for common events

// OnLoad adds a listener for the map load event.
func (m *Map) OnLoad(listener EventListener) *Map {
	return m.On(EventLoad, listener)
}

// OnClick adds a listener for map click events.
func (m *Map) OnClick(listener EventListener) *Map {
	return m.On(EventClick, listener)
}

// OnMouseMove adds a listener for mouse move events.
func (m *Map) OnMouseMove(listener EventListener) *Map {
	return m.On(EventMouseMove, listener)
}

// OnZoomEnd adds a listener for zoom end events.
func (m *Map) OnZoomEnd(listener EventListener) *Map {
	return m.On(EventZoomEnd, listener)
}

// OnMoveEnd adds a listener for move end events.
func (m *Map) OnMoveEnd(listener EventListener) *Map {
	return m.On(EventMoveEnd, listener)
}

// OnError adds a listener for error events.
func (m *Map) OnError(listener EventListener) *Map {
	return m.On(EventError, listener)
}

// OnStyleLoad adds a listener for style load events.
func (m *Map) OnStyleLoad(listener EventListener) *Map {
	return m.On(EventStyleLoad, listener)
}

// OnData adds a listener for data events.
func (m *Map) OnData(listener EventListener) *Map {
	return m.On(EventData, listener)
}

// OnSourceData adds a listener for source data events.
func (m *Map) OnSourceData(listener EventListener) *Map {
	return m.On(EventSourceData, listener)
}

// OnStyleData adds a listener for style data events.
func (m *Map) OnStyleData(listener EventListener) *Map {
	return m.On(EventStyleData, listener)
}

// Event utility functions

// GetEventLngLat extracts the geographical coordinates from a map event.
func GetEventLngLat(event js.Value) *LngLat {
	lngLat := event.Get("lngLat")
	if lngLat.IsUndefined() || lngLat.IsNull() {
		return nil
	}
	return &LngLat{value: lngLat}
}

// GetEventPoint extracts the pixel coordinates from a map event.
func GetEventPoint(event js.Value) *Point {
	point := event.Get("point")
	if point.IsUndefined() || point.IsNull() {
		return nil
	}
	return &Point{value: point}
}

// GetEventTarget extracts the event target from a map event.
func GetEventTarget(event js.Value) js.Value {
	return event.Get("target")
}

// GetEventType extracts the event type from a map event.
func GetEventType(event js.Value) string {
	eventType := event.Get("type")
	if eventType.IsUndefined() {
		return ""
	}
	return eventType.String()
}

// GetEventOriginalEvent extracts the original DOM event from a map event.
func GetEventOriginalEvent(event js.Value) js.Value {
	return event.Get("originalEvent")
}

// IsEventDefaultPrevented checks if preventDefault was called on the event.
func IsEventDefaultPrevented(event js.Value) bool {
	defaultPrevented := event.Get("defaultPrevented")
	if defaultPrevented.IsUndefined() {
		return false
	}
	return defaultPrevented.Bool()
}

// PreventEventDefault calls preventDefault on the event.
func PreventEventDefault(event js.Value) {
	event.Call("preventDefault")
}
