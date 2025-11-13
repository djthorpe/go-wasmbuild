//go:build js

package mapbox

import (
	"syscall/js"
)

// MarkerOptions represents options for creating a Marker.
type MarkerOptions struct {
	value js.Value
}

// NewMarkerOptions creates a new MarkerOptions object.
func NewMarkerOptions() *MarkerOptions {
	return &MarkerOptions{
		value: js.Global().Get("Object").New(),
	}
}

// JSValue returns the underlying js.Value.
func (mo *MarkerOptions) JSValue() js.Value {
	return mo.value
}

// SetElement sets the DOM element for the marker.
func (mo *MarkerOptions) SetElement(element js.Value) *MarkerOptions {
	mo.value.Set("element", element)
	return mo
}

// SetOffset sets the offset for the marker.
func (mo *MarkerOptions) SetOffset(point *Point) *MarkerOptions {
	mo.value.Set("offset", point.JSValue())
	return mo
}

// SetAnchor sets the anchor point for the marker.
func (mo *MarkerOptions) SetAnchor(anchor string) *MarkerOptions {
	mo.value.Set("anchor", anchor)
	return mo
}

// SetColor sets the color for the default marker.
func (mo *MarkerOptions) SetColor(color string) *MarkerOptions {
	mo.value.Set("color", color)
	return mo
}

// SetScale sets the scale for the marker.
func (mo *MarkerOptions) SetScale(scale float64) *MarkerOptions {
	mo.value.Set("scale", scale)
	return mo
}

// SetDraggable sets whether the marker is draggable.
func (mo *MarkerOptions) SetDraggable(draggable bool) *MarkerOptions {
	mo.value.Set("draggable", draggable)
	return mo
}

// SetClickTolerance sets the click tolerance for the marker.
func (mo *MarkerOptions) SetClickTolerance(tolerance int) *MarkerOptions {
	mo.value.Set("clickTolerance", tolerance)
	return mo
}

// SetRotation sets the rotation for the marker.
func (mo *MarkerOptions) SetRotation(rotation float64) *MarkerOptions {
	mo.value.Set("rotation", rotation)
	return mo
}

// SetRotationAlignment sets the rotation alignment for the marker.
func (mo *MarkerOptions) SetRotationAlignment(alignment string) *MarkerOptions {
	mo.value.Set("rotationAlignment", alignment)
	return mo
}

// SetPitchAlignment sets the pitch alignment for the marker.
func (mo *MarkerOptions) SetPitchAlignment(alignment string) *MarkerOptions {
	mo.value.Set("pitchAlignment", alignment)
	return mo
}

// Marker represents a marker on the map.
type Marker struct {
	value js.Value
}

// NewMarker creates a new Marker instance.
func NewMarker(options ...*MarkerOptions) *Marker {
	var markerInstance js.Value
	if len(options) > 0 {
		markerInstance = js.Global().Get("mapboxgl").Get("Marker").New(options[0].JSValue())
	} else {
		markerInstance = js.Global().Get("mapboxgl").Get("Marker").New()
	}
	return &Marker{value: markerInstance}
}

// MarkerFromValue wraps an existing js.Value as a Marker.
func MarkerFromValue(value js.Value) *Marker {
	return &Marker{value: value}
}

// JSValue returns the underlying js.Value.
func (m *Marker) JSValue() js.Value {
	return m.value
}

// AddTo adds the marker to a map.
func (m *Marker) AddTo(mapInstance *Map) *Marker {
	m.value.Call("addTo", mapInstance.JSValue())
	return m
}

// Remove removes the marker from the map.
func (m *Marker) Remove() *Marker {
	m.value.Call("remove")
	return m
}

// GetLngLat returns the geographical location of the marker.
func (m *Marker) GetLngLat() *LngLat {
	lngLat := m.value.Call("getLngLat")
	return &LngLat{value: lngLat}
}

// SetLngLat sets the geographical location of the marker.
func (m *Marker) SetLngLat(lngLat *LngLat) *Marker {
	m.value.Call("setLngLat", lngLat.JSValue())
	return m
}

// GetElement returns the marker's HTML element.
func (m *Marker) GetElement() js.Value {
	return m.value.Call("getElement")
}

// SetPopup binds a popup to the marker.
func (m *Marker) SetPopup(popup *Popup) *Marker {
	m.value.Call("setPopup", popup.JSValue())
	return m
}

// GetPopup returns the popup bound to the marker.
func (m *Marker) GetPopup() *Popup {
	popup := m.value.Call("getPopup")
	if popup.IsUndefined() || popup.IsNull() {
		return nil
	}
	return &Popup{value: popup}
}

// TogglePopup opens or closes the popup bound to the marker.
func (m *Marker) TogglePopup() *Marker {
	m.value.Call("togglePopup")
	return m
}

// GetOffset returns the marker's offset.
func (m *Marker) GetOffset() *Point {
	offset := m.value.Call("getOffset")
	return &Point{value: offset}
}

// SetOffset sets the marker's offset.
func (m *Marker) SetOffset(offset *Point) *Marker {
	m.value.Call("setOffset", offset.JSValue())
	return m
}

// SetDraggable sets whether the marker is draggable.
func (m *Marker) SetDraggable(draggable bool) *Marker {
	m.value.Call("setDraggable", draggable)
	return m
}

// IsDraggable returns whether the marker is draggable.
func (m *Marker) IsDraggable() bool {
	return m.value.Call("isDraggable").Bool()
}

// GetRotation returns the marker's rotation.
func (m *Marker) GetRotation() float64 {
	return m.value.Call("getRotation").Float()
}

// SetRotation sets the marker's rotation.
func (m *Marker) SetRotation(rotation float64) *Marker {
	m.value.Call("setRotation", rotation)
	return m
}

// GetRotationAlignment returns the marker's rotation alignment.
func (m *Marker) GetRotationAlignment() string {
	return m.value.Call("getRotationAlignment").String()
}

// SetRotationAlignment sets the marker's rotation alignment.
func (m *Marker) SetRotationAlignment(alignment string) *Marker {
	m.value.Call("setRotationAlignment", alignment)
	return m
}

// GetPitchAlignment returns the marker's pitch alignment.
func (m *Marker) GetPitchAlignment() string {
	return m.value.Call("getPitchAlignment").String()
}

// SetPitchAlignment sets the marker's pitch alignment.
func (m *Marker) SetPitchAlignment(alignment string) *Marker {
	m.value.Call("setPitchAlignment", alignment)
	return m
}

// PopupOptions represents options for creating a Popup.
type PopupOptions struct {
	value js.Value
}

// NewPopupOptions creates a new PopupOptions object.
func NewPopupOptions() *PopupOptions {
	return &PopupOptions{
		value: js.Global().Get("Object").New(),
	}
}

// JSValue returns the underlying js.Value.
func (po *PopupOptions) JSValue() js.Value {
	return po.value
}

// SetCloseButton sets whether to show a close button.
func (po *PopupOptions) SetCloseButton(show bool) *PopupOptions {
	po.value.Set("closeButton", show)
	return po
}

// SetCloseOnClick sets whether to close the popup on map click.
func (po *PopupOptions) SetCloseOnClick(close bool) *PopupOptions {
	po.value.Set("closeOnClick", close)
	return po
}

// SetCloseOnMove sets whether to close the popup on map move.
func (po *PopupOptions) SetCloseOnMove(close bool) *PopupOptions {
	po.value.Set("closeOnMove", close)
	return po
}

// SetFocusAfterOpen sets whether to focus the popup after opening.
func (po *PopupOptions) SetFocusAfterOpen(focus bool) *PopupOptions {
	po.value.Set("focusAfterOpen", focus)
	return po
}

// SetAnchor sets the anchor point for the popup.
func (po *PopupOptions) SetAnchor(anchor string) *PopupOptions {
	po.value.Set("anchor", anchor)
	return po
}

// SetOffset sets the offset for the popup.
func (po *PopupOptions) SetOffset(offset interface{}) *PopupOptions {
	switch v := offset.(type) {
	case *Point:
		po.value.Set("offset", v.JSValue())
	case js.Value:
		po.value.Set("offset", v)
	case float64:
		po.value.Set("offset", v)
	}
	return po
}

// SetClassName sets the CSS class name for the popup.
func (po *PopupOptions) SetClassName(className string) *PopupOptions {
	po.value.Set("className", className)
	return po
}

// SetMaxWidth sets the maximum width of the popup.
func (po *PopupOptions) SetMaxWidth(maxWidth string) *PopupOptions {
	po.value.Set("maxWidth", maxWidth)
	return po
}

// Popup represents a popup on the map.
type Popup struct {
	value js.Value
}

// NewPopup creates a new Popup instance.
func NewPopup(options ...*PopupOptions) *Popup {
	var popupInstance js.Value
	if len(options) > 0 {
		popupInstance = js.Global().Get("mapboxgl").Get("Popup").New(options[0].JSValue())
	} else {
		popupInstance = js.Global().Get("mapboxgl").Get("Popup").New()
	}
	return &Popup{value: popupInstance}
}

// PopupFromValue wraps an existing js.Value as a Popup.
func PopupFromValue(value js.Value) *Popup {
	return &Popup{value: value}
}

// JSValue returns the underlying js.Value.
func (p *Popup) JSValue() js.Value {
	return p.value
}

// AddTo adds the popup to a map.
func (p *Popup) AddTo(mapInstance *Map) *Popup {
	p.value.Call("addTo", mapInstance.JSValue())
	return p
}

// Remove removes the popup from the map.
func (p *Popup) Remove() *Popup {
	p.value.Call("remove")
	return p
}

// IsOpen returns whether the popup is open.
func (p *Popup) IsOpen() bool {
	return p.value.Call("isOpen").Bool()
}

// GetLngLat returns the geographical location of the popup.
func (p *Popup) GetLngLat() *LngLat {
	lngLat := p.value.Call("getLngLat")
	if lngLat.IsUndefined() || lngLat.IsNull() {
		return nil
	}
	return &LngLat{value: lngLat}
}

// SetLngLat sets the geographical location of the popup.
func (p *Popup) SetLngLat(lngLat *LngLat) *Popup {
	p.value.Call("setLngLat", lngLat.JSValue())
	return p
}

// SetHTML sets the HTML content of the popup.
func (p *Popup) SetHTML(html string) *Popup {
	p.value.Call("setHTML", html)
	return p
}

// SetText sets the text content of the popup.
func (p *Popup) SetText(text string) *Popup {
	p.value.Call("setText", text)
	return p
}

// SetDOMContent sets the DOM content of the popup.
func (p *Popup) SetDOMContent(element js.Value) *Popup {
	p.value.Call("setDOMContent", element)
	return p
}

// GetElement returns the popup's HTML element.
func (p *Popup) GetElement() js.Value {
	return p.value.Call("getElement")
}

// SetMaxWidth sets the maximum width of the popup.
func (p *Popup) SetMaxWidth(maxWidth string) *Popup {
	p.value.Call("setMaxWidth", maxWidth)
	return p
}

// GetMaxWidth returns the maximum width of the popup.
func (p *Popup) GetMaxWidth() string {
	return p.value.Call("getMaxWidth").String()
}

// SetOffset sets the offset of the popup.
func (p *Popup) SetOffset(offset interface{}) *Popup {
	switch v := offset.(type) {
	case *Point:
		p.value.Call("setOffset", v.JSValue())
	case js.Value:
		p.value.Call("setOffset", v)
	case float64:
		p.value.Call("setOffset", v)
	}
	return p
}

// GetOffset returns the offset of the popup.
func (p *Popup) GetOffset() js.Value {
	return p.value.Call("getOffset")
}

// Controls

// NavigationControl represents a navigation control.
type NavigationControl struct {
	value js.Value
}

// NewNavigationControl creates a new navigation control.
func NewNavigationControl(options ...js.Value) *NavigationControl {
	var controlInstance js.Value
	if len(options) > 0 {
		controlInstance = js.Global().Get("mapboxgl").Get("NavigationControl").New(options[0])
	} else {
		controlInstance = js.Global().Get("mapboxgl").Get("NavigationControl").New()
	}
	return &NavigationControl{value: controlInstance}
}

// JSValue returns the underlying js.Value.
func (nc *NavigationControl) JSValue() js.Value {
	return nc.value
}

// GeolocateControl represents a geolocate control.
type GeolocateControl struct {
	value js.Value
}

// NewGeolocateControl creates a new geolocate control.
func NewGeolocateControl(options ...js.Value) *GeolocateControl {
	var controlInstance js.Value
	if len(options) > 0 {
		controlInstance = js.Global().Get("mapboxgl").Get("GeolocateControl").New(options[0])
	} else {
		controlInstance = js.Global().Get("mapboxgl").Get("GeolocateControl").New()
	}
	return &GeolocateControl{value: controlInstance}
}

// JSValue returns the underlying js.Value.
func (gc *GeolocateControl) JSValue() js.Value {
	return gc.value
}

// Trigger starts the geolocation process.
func (gc *GeolocateControl) Trigger() *GeolocateControl {
	gc.value.Call("trigger")
	return gc
}

// AttributionControl represents an attribution control.
type AttributionControl struct {
	value js.Value
}

// NewAttributionControl creates a new attribution control.
func NewAttributionControl(options ...js.Value) *AttributionControl {
	var controlInstance js.Value
	if len(options) > 0 {
		controlInstance = js.Global().Get("mapboxgl").Get("AttributionControl").New(options[0])
	} else {
		controlInstance = js.Global().Get("mapboxgl").Get("AttributionControl").New()
	}
	return &AttributionControl{value: controlInstance}
}

// JSValue returns the underlying js.Value.
func (ac *AttributionControl) JSValue() js.Value {
	return ac.value
}

// ScaleControl represents a scale control.
type ScaleControl struct {
	value js.Value
}

// NewScaleControl creates a new scale control.
func NewScaleControl(options ...js.Value) *ScaleControl {
	var controlInstance js.Value
	if len(options) > 0 {
		controlInstance = js.Global().Get("mapboxgl").Get("ScaleControl").New(options[0])
	} else {
		controlInstance = js.Global().Get("mapboxgl").Get("ScaleControl").New()
	}
	return &ScaleControl{value: controlInstance}
}

// JSValue returns the underlying js.Value.
func (sc *ScaleControl) JSValue() js.Value {
	return sc.value
}

// SetUnit sets the unit for the scale control.
func (sc *ScaleControl) SetUnit(unit string) *ScaleControl {
	sc.value.Call("setUnit", unit)
	return sc
}

// FullscreenControl represents a fullscreen control.
type FullscreenControl struct {
	value js.Value
}

// NewFullscreenControl creates a new fullscreen control.
func NewFullscreenControl(options ...js.Value) *FullscreenControl {
	var controlInstance js.Value
	if len(options) > 0 {
		controlInstance = js.Global().Get("mapboxgl").Get("FullscreenControl").New(options[0])
	} else {
		controlInstance = js.Global().Get("mapboxgl").Get("FullscreenControl").New()
	}
	return &FullscreenControl{value: controlInstance}
}

// JSValue returns the underlying js.Value.
func (fc *FullscreenControl) JSValue() js.Value {
	return fc.value
}
