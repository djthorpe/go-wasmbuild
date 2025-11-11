//go:build js

package mapbox

import (
	"syscall/js"
)

// Source represents a data source for map layers.
type Source struct {
	value js.Value
}

// SourceFromValue wraps an existing js.Value as a Source.
func SourceFromValue(value js.Value) *Source {
	return &Source{value: value}
}

// JSValue returns the underlying js.Value.
func (s *Source) JSValue() js.Value {
	return s.value
}

// GetType returns the type of the source.
func (s *Source) GetType() string {
	sourceType := s.value.Get("type")
	if sourceType.IsUndefined() {
		return ""
	}
	return sourceType.String()
}

// GeoJSONSource represents a GeoJSON data source.
type GeoJSONSource struct {
	Source
}

// NewGeoJSONSource creates a new GeoJSON source.
func NewGeoJSONSource(options js.Value) *GeoJSONSource {
	sourceObj := js.Global().Get("Object").New()
	sourceObj.Set("type", "geojson")

	// Copy options to the source object
	if !options.IsUndefined() && !options.IsNull() {
		// Copy properties from options
		keys := js.Global().Get("Object").Call("keys", options)
		length := keys.Get("length").Int()
		for i := 0; i < length; i++ {
			key := keys.Index(i).String()
			sourceObj.Set(key, options.Get(key))
		}
	}

	return &GeoJSONSource{Source{value: sourceObj}}
}

// SetData sets the GeoJSON data for the source.
func (gs *GeoJSONSource) SetData(data js.Value) *GeoJSONSource {
	gs.value.Set("data", data)
	return gs
}

// GetData returns the GeoJSON data of the source.
func (gs *GeoJSONSource) GetData() js.Value {
	return gs.value.Get("data")
}

// VectorSource represents a vector tile source.
type VectorSource struct {
	Source
}

// NewVectorSource creates a new vector tile source.
func NewVectorSource(options js.Value) *VectorSource {
	sourceObj := js.Global().Get("Object").New()
	sourceObj.Set("type", "vector")

	if !options.IsUndefined() && !options.IsNull() {
		keys := js.Global().Get("Object").Call("keys", options)
		length := keys.Get("length").Int()
		for i := 0; i < length; i++ {
			key := keys.Index(i).String()
			sourceObj.Set(key, options.Get(key))
		}
	}

	return &VectorSource{Source{value: sourceObj}}
}

// RasterSource represents a raster tile source.
type RasterSource struct {
	Source
}

// NewRasterSource creates a new raster tile source.
func NewRasterSource(options js.Value) *RasterSource {
	sourceObj := js.Global().Get("Object").New()
	sourceObj.Set("type", "raster")

	if !options.IsUndefined() && !options.IsNull() {
		keys := js.Global().Get("Object").Call("keys", options)
		length := keys.Get("length").Int()
		for i := 0; i < length; i++ {
			key := keys.Index(i).String()
			sourceObj.Set(key, options.Get(key))
		}
	}

	return &RasterSource{Source{value: sourceObj}}
}

// ImageSource represents an image source.
type ImageSource struct {
	Source
}

// NewImageSource creates a new image source.
func NewImageSource(options js.Value) *ImageSource {
	sourceObj := js.Global().Get("Object").New()
	sourceObj.Set("type", "image")

	if !options.IsUndefined() && !options.IsNull() {
		keys := js.Global().Get("Object").Call("keys", options)
		length := keys.Get("length").Int()
		for i := 0; i < length; i++ {
			key := keys.Index(i).String()
			sourceObj.Set(key, options.Get(key))
		}
	}

	return &ImageSource{Source{value: sourceObj}}
}

// VideoSource represents a video source.
type VideoSource struct {
	Source
}

// NewVideoSource creates a new video source.
func NewVideoSource(options js.Value) *VideoSource {
	sourceObj := js.Global().Get("Object").New()
	sourceObj.Set("type", "video")

	if !options.IsUndefined() && !options.IsNull() {
		keys := js.Global().Get("Object").Call("keys", options)
		length := keys.Get("length").Int()
		for i := 0; i < length; i++ {
			key := keys.Index(i).String()
			sourceObj.Set(key, options.Get(key))
		}
	}

	return &VideoSource{Source{value: sourceObj}}
}

// Layer represents a map layer.
type Layer struct {
	value js.Value
}

// LayerFromValue wraps an existing js.Value as a Layer.
func LayerFromValue(value js.Value) *Layer {
	return &Layer{value: value}
}

// JSValue returns the underlying js.Value.
func (l *Layer) JSValue() js.Value {
	return l.value
}

// GetId returns the layer's ID.
func (l *Layer) GetId() string {
	id := l.value.Get("id")
	if id.IsUndefined() {
		return ""
	}
	return id.String()
}

// GetType returns the layer's type.
func (l *Layer) GetType() string {
	layerType := l.value.Get("type")
	if layerType.IsUndefined() {
		return ""
	}
	return layerType.String()
}

// GetSource returns the layer's source ID.
func (l *Layer) GetSource() string {
	source := l.value.Get("source")
	if source.IsUndefined() {
		return ""
	}
	return source.String()
}

// GetSourceLayer returns the layer's source layer (for vector sources).
func (l *Layer) GetSourceLayer() string {
	sourceLayer := l.value.Get("source-layer")
	if sourceLayer.IsUndefined() {
		return ""
	}
	return sourceLayer.String()
}

// GetLayout returns the layer's layout properties.
func (l *Layer) GetLayout() js.Value {
	return l.value.Get("layout")
}

// GetPaint returns the layer's paint properties.
func (l *Layer) GetPaint() js.Value {
	return l.value.Get("paint")
}

// GetFilter returns the layer's filter.
func (l *Layer) GetFilter() js.Value {
	return l.value.Get("filter")
}

// GetMinZoom returns the layer's minimum zoom level.
func (l *Layer) GetMinZoom() float64 {
	minZoom := l.value.Get("minzoom")
	if minZoom.IsUndefined() {
		return 0
	}
	return minZoom.Float()
}

// GetMaxZoom returns the layer's maximum zoom level.
func (l *Layer) GetMaxZoom() float64 {
	maxZoom := l.value.Get("maxzoom")
	if maxZoom.IsUndefined() {
		return 24
	}
	return maxZoom.Float()
}

// GetVisibility returns the layer's visibility.
func (l *Layer) GetVisibility() string {
	layout := l.GetLayout()
	if layout.IsUndefined() {
		return "visible"
	}
	visibility := layout.Get("visibility")
	if visibility.IsUndefined() {
		return "visible"
	}
	return visibility.String()
}

// Layer creation functions

// NewFillLayer creates a new fill layer.
func NewFillLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "fill")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewLineLayer creates a new line layer.
func NewLineLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "line")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewSymbolLayer creates a new symbol layer.
func NewSymbolLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "symbol")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewCircleLayer creates a new circle layer.
func NewCircleLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "circle")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewHeatmapLayer creates a new heatmap layer.
func NewHeatmapLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "heatmap")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewFillExtrusionLayer creates a new fill-extrusion layer.
func NewFillExtrusionLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "fill-extrusion")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewRasterLayer creates a new raster layer.
func NewRasterLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "raster")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewHillshadeLayer creates a new hillshade layer.
func NewHillshadeLayer(id, source string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "hillshade")
	layer.Set("source", source)
	return &Layer{value: layer}
}

// NewBackgroundLayer creates a new background layer.
func NewBackgroundLayer(id string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "background")
	return &Layer{value: layer}
}

// NewSkyLayer creates a new sky layer.
func NewSkyLayer(id string) *Layer {
	layer := js.Global().Get("Object").New()
	layer.Set("id", id)
	layer.Set("type", "sky")
	return &Layer{value: layer}
}

// Map methods for working with sources and layers

// AddSource adds a source to the map.
func (m *Map) AddSource(id string, source *Source) *Map {
	m.value.Call("addSource", id, source.JSValue())
	return m
}

// GetSource returns a source by ID.
func (m *Map) GetSource(id string) *Source {
	source := m.value.Call("getSource", id)
	if source.IsUndefined() || source.IsNull() {
		return nil
	}
	return &Source{value: source}
}

// RemoveSource removes a source from the map.
func (m *Map) RemoveSource(id string) *Map {
	m.value.Call("removeSource", id)
	return m
}

// AddLayer adds a layer to the map.
func (m *Map) AddLayer(layer *Layer, beforeId ...string) *Map {
	if len(beforeId) > 0 {
		m.value.Call("addLayer", layer.JSValue(), beforeId[0])
	} else {
		m.value.Call("addLayer", layer.JSValue())
	}
	return m
}

// GetLayer returns a layer by ID.
func (m *Map) GetLayer(id string) *Layer {
	layer := m.value.Call("getLayer", id)
	if layer.IsUndefined() || layer.IsNull() {
		return nil
	}
	return &Layer{value: layer}
}

// RemoveLayer removes a layer from the map.
func (m *Map) RemoveLayer(id string) *Map {
	m.value.Call("removeLayer", id)
	return m
}

// MoveLayer moves a layer to a new position.
func (m *Map) MoveLayer(id string, beforeId ...string) *Map {
	if len(beforeId) > 0 {
		m.value.Call("moveLayer", id, beforeId[0])
	} else {
		m.value.Call("moveLayer", id)
	}
	return m
}

// GetLayers returns all layers on the map.
func (m *Map) GetLayers() []*Layer {
	layersArray := m.value.Call("getStyle").Get("layers")
	if layersArray.IsUndefined() {
		return nil
	}

	length := layersArray.Get("length").Int()
	layers := make([]*Layer, length)
	for i := 0; i < length; i++ {
		layers[i] = &Layer{value: layersArray.Index(i)}
	}
	return layers
}

// SetFilter sets the filter for a layer.
func (m *Map) SetFilter(layerId string, filter js.Value, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("setFilter", layerId, filter, options[0])
	} else {
		m.value.Call("setFilter", layerId, filter)
	}
	return m
}

// GetFilter returns the filter for a layer.
func (m *Map) GetFilter(layerId string) js.Value {
	return m.value.Call("getFilter", layerId)
}

// SetLayerZoomRange sets the zoom range for a layer.
func (m *Map) SetLayerZoomRange(layerId string, minZoom, maxZoom float64) *Map {
	m.value.Call("setLayerZoomRange", layerId, minZoom, maxZoom)
	return m
}

// GetLayoutProperty returns a layout property value for a layer.
func (m *Map) GetLayoutProperty(layerId, name string) js.Value {
	return m.value.Call("getLayoutProperty", layerId, name)
}

// SetLayoutProperty sets a layout property for a layer.
func (m *Map) SetLayoutProperty(layerId, name string, value js.Value, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("setLayoutProperty", layerId, name, value, options[0])
	} else {
		m.value.Call("setLayoutProperty", layerId, name, value)
	}
	return m
}

// GetPaintProperty returns a paint property value for a layer.
func (m *Map) GetPaintProperty(layerId, name string) js.Value {
	return m.value.Call("getPaintProperty", layerId, name)
}

// SetPaintProperty sets a paint property for a layer.
func (m *Map) SetPaintProperty(layerId, name string, value js.Value, options ...js.Value) *Map {
	if len(options) > 0 {
		m.value.Call("setPaintProperty", layerId, name, value, options[0])
	} else {
		m.value.Call("setPaintProperty", layerId, name, value)
	}
	return m
}

// QueryRenderedFeatures returns an array of features at a given point or within a bounding box.
func (m *Map) QueryRenderedFeatures(pointOrBounds js.Value, options ...js.Value) js.Value {
	if len(options) > 0 {
		return m.value.Call("queryRenderedFeatures", pointOrBounds, options[0])
	}
	return m.value.Call("queryRenderedFeatures", pointOrBounds)
}

// QuerySourceFeatures returns an array of features from a source.
func (m *Map) QuerySourceFeatures(sourceId string, options ...js.Value) js.Value {
	if len(options) > 0 {
		return m.value.Call("querySourceFeatures", sourceId, options[0])
	}
	return m.value.Call("querySourceFeatures", sourceId)
}

// SetFeatureState sets the state of a feature.
func (m *Map) SetFeatureState(feature js.Value, state js.Value) *Map {
	m.value.Call("setFeatureState", feature, state)
	return m
}

// GetFeatureState returns the state of a feature.
func (m *Map) GetFeatureState(feature js.Value) js.Value {
	return m.value.Call("getFeatureState", feature)
}

// RemoveFeatureState removes the state of a feature.
func (m *Map) RemoveFeatureState(feature js.Value, key ...string) *Map {
	if len(key) > 0 {
		m.value.Call("removeFeatureState", feature, key[0])
	} else {
		m.value.Call("removeFeatureState", feature)
	}
	return m
}
