//go:build js

package mapbox

import (
	"syscall/js"
)

// MapOptions represents the options for creating a new Map.
type MapOptions struct {
	value js.Value
}

// NewMapOptions creates a new MapOptions object.
func NewMapOptions() *MapOptions {
	return &MapOptions{
		value: js.Global().Get("Object").New(),
	}
}

// MapOptionsFromValue wraps an existing js.Value as MapOptions.
func MapOptionsFromValue(value js.Value) *MapOptions {
	return &MapOptions{value: value}
}

// JSValue returns the underlying js.Value.
func (mo *MapOptions) JSValue() js.Value {
	return mo.value
}

// SetContainer sets the HTML element or the string ID of an element in which to render the map.
func (mo *MapOptions) SetContainer(container interface{}) *MapOptions {
	switch v := container.(type) {
	case string:
		mo.value.Set("container", v)
	case js.Value:
		mo.value.Set("container", v)
	}
	return mo
}

// SetStyle sets the map's Mapbox style.
func (mo *MapOptions) SetStyle(style string) *MapOptions {
	mo.value.Set("style", style)
	return mo
}

// SetCenter sets the initial geographical centerpoint of the map.
func (mo *MapOptions) SetCenter(center *LngLat) *MapOptions {
	mo.value.Set("center", center.JSValue())
	return mo
}

// SetZoom sets the initial zoom level of the map.
func (mo *MapOptions) SetZoom(zoom float64) *MapOptions {
	mo.value.Set("zoom", zoom)
	return mo
}

// SetBearing sets the initial bearing (rotation) of the map, measured in degrees counter-clockwise from north.
func (mo *MapOptions) SetBearing(bearing float64) *MapOptions {
	mo.value.Set("bearing", bearing)
	return mo
}

// SetPitch sets the initial pitch (tilt) of the map, measured in degrees away from the plane of the screen.
func (mo *MapOptions) SetPitch(pitch float64) *MapOptions {
	mo.value.Set("pitch", pitch)
	return mo
}

// SetMinZoom sets the minimum zoom level of the map.
func (mo *MapOptions) SetMinZoom(minZoom float64) *MapOptions {
	mo.value.Set("minZoom", minZoom)
	return mo
}

// SetMaxZoom sets the maximum zoom level of the map.
func (mo *MapOptions) SetMaxZoom(maxZoom float64) *MapOptions {
	mo.value.Set("maxZoom", maxZoom)
	return mo
}

// SetMaxBounds sets the maximum extent of the map.
func (mo *MapOptions) SetMaxBounds(bounds *LngLatBounds) *MapOptions {
	mo.value.Set("maxBounds", bounds.JSValue())
	return mo
}

// SetInteractive sets whether the map is interactive.
func (mo *MapOptions) SetInteractive(interactive bool) *MapOptions {
	mo.value.Set("interactive", interactive)
	return mo
}

// SetScrollZoom enables or disables scroll to zoom interaction.
func (mo *MapOptions) SetScrollZoom(enable bool) *MapOptions {
	mo.value.Set("scrollZoom", enable)
	return mo
}

// SetBoxZoom enables or disables box zoom interaction.
func (mo *MapOptions) SetBoxZoom(enable bool) *MapOptions {
	mo.value.Set("boxZoom", enable)
	return mo
}

// SetDragRotate enables or disables drag to rotate interaction.
func (mo *MapOptions) SetDragRotate(enable bool) *MapOptions {
	mo.value.Set("dragRotate", enable)
	return mo
}

// SetDragPan enables or disables drag to pan interaction.
func (mo *MapOptions) SetDragPan(enable bool) *MapOptions {
	mo.value.Set("dragPan", enable)
	return mo
}

// SetKeyboard enables or disables keyboard navigation.
func (mo *MapOptions) SetKeyboard(enable bool) *MapOptions {
	mo.value.Set("keyboard", enable)
	return mo
}

// SetDoubleClickZoom enables or disables double click to zoom interaction.
func (mo *MapOptions) SetDoubleClickZoom(enable bool) *MapOptions {
	mo.value.Set("doubleClickZoom", enable)
	return mo
}

// SetTouchZoomRotate enables or disables touch zoom and rotate interaction.
func (mo *MapOptions) SetTouchZoomRotate(enable bool) *MapOptions {
	mo.value.Set("touchZoomRotate", enable)
	return mo
}

// SetTrackResize sets whether the map should automatically resize when the browser window resizes.
func (mo *MapOptions) SetTrackResize(enable bool) *MapOptions {
	mo.value.Set("trackResize", enable)
	return mo
}

// SetHash sets whether the map's position should be synced with the page's hash.
func (mo *MapOptions) SetHash(enable bool) *MapOptions {
	mo.value.Set("hash", enable)
	return mo
}

// SetRenderWorldCopies sets whether to render multiple copies of the world at low zoom levels.
func (mo *MapOptions) SetRenderWorldCopies(enable bool) *MapOptions {
	mo.value.Set("renderWorldCopies", enable)
	return mo
}

// SetAttributionControl sets whether an attribution control is added to the map.
func (mo *MapOptions) SetAttributionControl(enable bool) *MapOptions {
	mo.value.Set("attributionControl", enable)
	return mo
}

// SetLogoPosition sets the position of the Mapbox logo.
func (mo *MapOptions) SetLogoPosition(position string) *MapOptions {
	mo.value.Set("logoPosition", position)
	return mo
}

// SetFailIfMajorPerformanceCaveat sets whether to fail if the performance would be dramatically worse.
func (mo *MapOptions) SetFailIfMajorPerformanceCaveat(enable bool) *MapOptions {
	mo.value.Set("failIfMajorPerformanceCaveat", enable)
	return mo
}

// SetPreserveDrawingBuffer sets whether to preserve the drawing buffer.
func (mo *MapOptions) SetPreserveDrawingBuffer(enable bool) *MapOptions {
	mo.value.Set("preserveDrawingBuffer", enable)
	return mo
}

// SetAntialias sets whether to enable antialiasing.
func (mo *MapOptions) SetAntialias(enable bool) *MapOptions {
	mo.value.Set("antialias", enable)
	return mo
}

// SetRefreshExpiredTiles sets whether to refresh expired tiles automatically.
func (mo *MapOptions) SetRefreshExpiredTiles(enable bool) *MapOptions {
	mo.value.Set("refreshExpiredTiles", enable)
	return mo
}

// SetTransformRequest sets a function to transform requests made by the map.
func (mo *MapOptions) SetTransformRequest(fn js.Func) *MapOptions {
	mo.value.Set("transformRequest", fn)
	return mo
}

// SetAccessToken sets the access token for the map.
func (mo *MapOptions) SetAccessToken(token string) *MapOptions {
	mo.value.Set("accessToken", token)
	return mo
}

// Map represents a Mapbox GL JS map.
type Map struct {
	value js.Value
}

// NewMap creates a new Map instance.
func NewMap(options *MapOptions) *Map {
	mapInstance := js.Global().Get("mapboxgl").Get("Map").New(options.JSValue())
	return &Map{value: mapInstance}
}

// MapFromValue wraps an existing js.Value as a Map.
func MapFromValue(value js.Value) *Map {
	return &Map{value: value}
}

// JSValue returns the underlying js.Value.
func (m *Map) JSValue() js.Value {
	return m.value
}

// AddControl adds a control to the map.
func (m *Map) AddControl(control js.Value, position ...string) *Map {
	if len(position) > 0 {
		m.value.Call("addControl", control, position[0])
	} else {
		m.value.Call("addControl", control)
	}
	return m
}

// RemoveControl removes a control from the map.
func (m *Map) RemoveControl(control js.Value) *Map {
	m.value.Call("removeControl", control)
	return m
}

// HasControl checks if a control exists on the map.
func (m *Map) HasControl(control js.Value) bool {
	return m.value.Call("hasControl", control).Bool()
}

// Resize resizes the map according to the dimensions of its container element.
func (m *Map) Resize() *Map {
	m.value.Call("resize")
	return m
}

// GetBounds returns the geographical bounds visible in the current map view.
func (m *Map) GetBounds() *LngLatBounds {
	bounds := m.value.Call("getBounds")
	return &LngLatBounds{value: bounds}
}

// GetMaxBounds returns the maximum geographical bounds to which the user can pan or zoom the map.
func (m *Map) GetMaxBounds() *LngLatBounds {
	bounds := m.value.Call("getMaxBounds")
	if bounds.IsNull() {
		return nil
	}
	return &LngLatBounds{value: bounds}
}

// SetMaxBounds sets the maximum geographical bounds to which the user can pan or zoom the map.
func (m *Map) SetMaxBounds(bounds *LngLatBounds) *Map {
	if bounds != nil {
		m.value.Call("setMaxBounds", bounds.JSValue())
	} else {
		m.value.Call("setMaxBounds", js.Null())
	}
	return m
}

// SetMinZoom sets the minimum zoom level of the map.
func (m *Map) SetMinZoom(minZoom float64) *Map {
	m.value.Call("setMinZoom", minZoom)
	return m
}

// SetMaxZoom sets the maximum zoom level of the map.
func (m *Map) SetMaxZoom(maxZoom float64) *Map {
	m.value.Call("setMaxZoom", maxZoom)
	return m
}

// GetMinZoom returns the minimum zoom level of the map.
func (m *Map) GetMinZoom() float64 {
	return m.value.Call("getMinZoom").Float()
}

// GetMaxZoom returns the maximum zoom level of the map.
func (m *Map) GetMaxZoom() float64 {
	return m.value.Call("getMaxZoom").Float()
}

// Project converts a LngLat to a Point in pixel coordinates.
func (m *Map) Project(lngLat *LngLat) *Point {
	point := m.value.Call("project", lngLat.JSValue())
	return &Point{value: point}
}

// Unproject converts a Point in pixel coordinates to a LngLat.
func (m *Map) Unproject(point *Point) *LngLat {
	lngLat := m.value.Call("unproject", point.JSValue())
	return &LngLat{value: lngLat}
}

// IsMoving returns true if the map is currently moving.
func (m *Map) IsMoving() bool {
	return m.value.Call("isMoving").Bool()
}

// IsZooming returns true if the map is currently zooming.
func (m *Map) IsZooming() bool {
	return m.value.Call("isZooming").Bool()
}

// IsRotating returns true if the map is currently rotating.
func (m *Map) IsRotating() bool {
	return m.value.Call("isRotating").Bool()
}

// GetContainer returns the map's containing HTML element.
func (m *Map) GetContainer() js.Value {
	return m.value.Call("getContainer")
}

// GetCanvasContainer returns the HTML element containing the map's canvas.
func (m *Map) GetCanvasContainer() js.Value {
	return m.value.Call("getCanvasContainer")
}

// GetCanvas returns the map's canvas element.
func (m *Map) GetCanvas() js.Value {
	return m.value.Call("getCanvas")
}

// Loaded returns true if the style is loaded and all tiles are loaded.
func (m *Map) Loaded() bool {
	return m.value.Call("loaded").Bool()
}

// AreTilesLoaded returns true if all tiles in the current viewport are loaded.
func (m *Map) AreTilesLoaded() bool {
	return m.value.Call("areTilesLoaded").Bool()
}

// Remove removes the map from the page and releases its resources.
func (m *Map) Remove() {
	m.value.Call("remove")
}

// TriggerRepaint triggers a repaint of the map.
func (m *Map) TriggerRepaint() {
	m.value.Call("triggerRepaint")
}

// GetStyle returns a representation of the map's style.
func (m *Map) GetStyle() js.Value {
	return m.value.Call("getStyle")
}

// SetStyle sets the map's Mapbox style.
func (m *Map) SetStyle(style string, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("setStyle", style, options[0])
	} else {
		m.value.Call("setStyle", style)
	}
	return m
}

// IsStyleLoaded returns true if the style is loaded.
func (m *Map) IsStyleLoaded() bool {
	return m.value.Call("isStyleLoaded").Bool()
}

// GetCenter returns the map's geographical centerpoint.
func (m *Map) GetCenter() *LngLat {
	center := m.value.Call("getCenter")
	return &LngLat{value: center}
}

// SetCenter sets the map's geographical centerpoint.
func (m *Map) SetCenter(center *LngLat, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("setCenter", center.JSValue(), eventData[0])
	} else {
		m.value.Call("setCenter", center.JSValue())
	}
	return m
}

// PanBy pans the map by the specified offset.
func (m *Map) PanBy(offset *Point, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("panBy", offset.JSValue(), options[0])
	} else {
		m.value.Call("panBy", offset.JSValue())
	}
	return m
}

// PanTo pans the map to the specified location.
func (m *Map) PanTo(lngLat *LngLat, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("panTo", lngLat.JSValue(), options[0])
	} else {
		m.value.Call("panTo", lngLat.JSValue())
	}
	return m
}

// GetZoom returns the map's current zoom level.
func (m *Map) GetZoom() float64 {
	return m.value.Call("getZoom").Float()
}

// SetZoom sets the map's zoom level.
func (m *Map) SetZoom(zoom float64, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("setZoom", zoom, eventData[0])
	} else {
		m.value.Call("setZoom", zoom)
	}
	return m
}

// ZoomTo zooms the map to the specified zoom level.
func (m *Map) ZoomTo(zoom float64, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("zoomTo", zoom, options[0])
	} else {
		m.value.Call("zoomTo", zoom)
	}
	return m
}

// ZoomIn increases the map's zoom level by 1.
func (m *Map) ZoomIn(options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("zoomIn", options[0])
	} else {
		m.value.Call("zoomIn")
	}
	return m
}

// ZoomOut decreases the map's zoom level by 1.
func (m *Map) ZoomOut(options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("zoomOut", options[0])
	} else {
		m.value.Call("zoomOut")
	}
	return m
}

// GetBearing returns the map's current bearing.
func (m *Map) GetBearing() float64 {
	return m.value.Call("getBearing").Float()
}

// SetBearing sets the map's bearing (rotation).
func (m *Map) SetBearing(bearing float64, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("setBearing", bearing, eventData[0])
	} else {
		m.value.Call("setBearing", bearing)
	}
	return m
}

// RotateTo rotates the map to the specified bearing.
func (m *Map) RotateTo(bearing float64, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("rotateTo", bearing, options[0])
	} else {
		m.value.Call("rotateTo", bearing)
	}
	return m
}

// ResetNorth resets the map bearing to 0 (north).
func (m *Map) ResetNorth(options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("resetNorth", options[0])
	} else {
		m.value.Call("resetNorth")
	}
	return m
}

// ResetNorthPitch resets the map bearing to 0 and pitch to 0.
func (m *Map) ResetNorthPitch(options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("resetNorthPitch", options[0])
	} else {
		m.value.Call("resetNorthPitch")
	}
	return m
}

// SnapToNorth snaps the map bearing to 0 if it's within a threshold.
func (m *Map) SnapToNorth(options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("snapToNorth", options[0])
	} else {
		m.value.Call("snapToNorth")
	}
	return m
}

// GetPitch returns the map's current pitch.
func (m *Map) GetPitch() float64 {
	return m.value.Call("getPitch").Float()
}

// SetPitch sets the map's pitch (tilt).
func (m *Map) SetPitch(pitch float64, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("setPitch", pitch, eventData[0])
	} else {
		m.value.Call("setPitch", pitch)
	}
	return m
}

// FitBounds pans and zooms the map to contain its visible area within the specified geographical bounds.
func (m *Map) FitBounds(bounds *LngLatBounds, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("fitBounds", bounds.JSValue(), options[0])
	} else {
		m.value.Call("fitBounds", bounds.JSValue())
	}
	return m
}

// FitScreenCoordinates fits the map to a bounding box defined by screen coordinates.
func (m *Map) FitScreenCoordinates(p0, p1 *Point, bearing float64, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("fitScreenCoordinates", p0.JSValue(), p1.JSValue(), bearing, options[0])
	} else {
		m.value.Call("fitScreenCoordinates", p0.JSValue(), p1.JSValue(), bearing)
	}
	return m
}

// JumpTo jumps to a series of map camera options immediately.
func (m *Map) JumpTo(options js.Value, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("jumpTo", options, eventData[0])
	} else {
		m.value.Call("jumpTo", options)
	}
	return m
}

// EaseTo animates to a series of map camera options.
func (m *Map) EaseTo(options js.Value, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("easeTo", options, eventData[0])
	} else {
		m.value.Call("easeTo", options)
	}
	return m
}

// FlyTo animates to a series of map camera options with an animated flight path.
func (m *Map) FlyTo(options js.Value, eventData ...js.Value) *Map {
	if len(eventData) > 0 {
		m.value.Call("flyTo", options, eventData[0])
	} else {
		m.value.Call("flyTo", options)
	}
	return m
}

// Stop stops any animated transition or rotation that's currently in progress.
func (m *Map) Stop() *Map {
	m.value.Call("stop")
	return m
}
