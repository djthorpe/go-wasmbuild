# Mapbox GL JS Go Bindings

This package provides Go bindings for the Mapbox GL JS library, allowing you to use Mapbox maps in Go applications compiled to WebAssembly.

## Overview

The bindings provide Go wrappers for the main Mapbox GL JS components:

- **MapboxGL**: Main namespace with global functions and utilities
- **Map**: The core map component with camera controls and interaction
- **LngLat/LngLatBounds/Point**: Geographic and pixel coordinate utilities  
- **Markers & Popups**: UI elements for displaying information on the map
- **Sources & Layers**: Data sources and rendering layers
- **Events**: Event handling system for map interactions
- **Controls**: Built-in UI controls (navigation, geolocation, etc.)

## Quick Start

```go
//go:build js

package main

import (
    "syscall/js"
    "github.com/djthorpe/go-dom/pkg/mapbox"
)

func main() {
    // Initialize Mapbox GL JS
    mapboxgl := mapbox.New()
    if mapboxgl == nil {
        panic("Mapbox GL JS not available")
    }

    // Set access token
    mapboxgl.SetAccessToken("your-mapbox-token")

    // Create map options
    options := mapbox.NewMapOptions().
        SetContainer("map").  // HTML element ID
        SetStyle("mapbox://styles/mapbox/streets-v11").
        SetCenter(mapbox.NewLngLat(-74.5, 40)).
        SetZoom(9)

    // Create the map
    mapInstance := mapboxgl.NewMap(*options)

    // Add event listeners
    mapInstance.OnLoad(func(event js.Value) {
        println("Map loaded!")
    })

    mapInstance.OnClick(func(event js.Value) {
        lngLat := mapbox.GetEventLngLat(event)
        if lngLat != nil {
            println("Clicked at:", lngLat.Lng(), lngLat.Lat())
        }
    })

    // Keep the program running
    select {}
}
```

## Features

### Map Creation and Controls

```go
// Create map with options
options := mapbox.NewMapOptions().
    SetContainer("map").
    SetStyle("mapbox://styles/mapbox/satellite-v9").
    SetCenter(mapbox.NewLngLat(-122.4, 37.8)).
    SetZoom(10).
    SetBearing(45).
    SetPitch(60)

mapInstance := mapboxgl.NewMap(*options)

// Add navigation controls
navControl := mapbox.NewNavigationControl()
mapInstance.AddControl(navControl.JSValue(), "top-right")

// Add geolocate control
geoControl := mapbox.NewGeolocateControl()
mapInstance.AddControl(geoControl.JSValue(), "top-left")
```

### Working with Coordinates

```go
// Create coordinates
lngLat := mapbox.NewLngLat(-74.006, 40.7128) // New York City
point := mapbox.NewPoint(100, 200)           // Pixel coordinates

// Create bounds
bounds := mapbox.NewLngLatBounds(
    mapbox.NewLngLat(-74.1, 40.7),  // Southwest
    mapbox.NewLngLat(-73.9, 40.8),  // Northeast
)

// Coordinate conversion
pixelPoint := mapInstance.Project(lngLat)
geoPoint := mapInstance.Unproject(point)

// Fit map to bounds
mapInstance.FitBounds(bounds)
```

### Adding Markers and Popups

```go
// Create a marker
marker := mapbox.NewMarker().
    SetLngLat(mapbox.NewLngLat(-74.006, 40.7128)).
    AddTo(mapInstance)

// Create a popup
popup := mapbox.NewPopup().
    SetLngLat(mapbox.NewLngLat(-74.006, 40.7128)).
    SetHTML("<h3>Hello World!</h3><p>This is a popup.</p>").
    AddTo(mapInstance)

// Bind popup to marker
marker.SetPopup(popup)
```

### Data Sources and Layers

```go
// Add a GeoJSON source
geoJsonSource := mapbox.NewGeoJSONSource(js.Global().Get("Object").New())
geoJsonSource.SetData(geoJsonData) // js.Value containing GeoJSON
mapInstance.AddSource("my-data", &geoJsonSource.Source)

// Add a layer
layer := mapbox.NewCircleLayer("my-layer", "my-data")
mapInstance.AddLayer(layer)

// Style the layer
mapInstance.SetPaintProperty("my-layer", "circle-radius", js.ValueOf(8))
mapInstance.SetPaintProperty("my-layer", "circle-color", js.ValueOf("#ff0000"))
```

### Event Handling

```go
// Map events
mapInstance.OnLoad(func(event js.Value) {
    println("Map loaded")
})

mapInstance.OnMoveEnd(func(event js.Value) {
    center := mapInstance.GetCenter()
    println("Map moved to:", center.Lng(), center.Lat())
})

mapInstance.OnZoomEnd(func(event js.Value) {
    zoom := mapInstance.GetZoom()
    println("Zoom level:", zoom)
})

// Layer-specific events
mapInstance.OnLayer("click", "my-layer", func(event js.Value) {
    point := mapbox.GetEventPoint(event)
    features := mapInstance.QueryRenderedFeatures(point.JSValue(), nil)
    println("Clicked features:", features.Get("length").Int())
})

// Mouse events
mapInstance.OnMouseMove(func(event js.Value) {
    lngLat := mapbox.GetEventLngLat(event)
    if lngLat != nil {
        // Update coordinates display
    }
})
```

### Camera Animations

```go
// Fly to a location
options := js.Global().Get("Object").New()
options.Set("center", mapbox.NewLngLat(-122.4, 37.8).JSValue())
options.Set("zoom", 12)
options.Set("duration", 2000)
mapInstance.FlyTo(options)

// Ease to new position
options = js.Global().Get("Object").New()
options.Set("center", mapbox.NewLngLat(-74.006, 40.7128).JSValue())
options.Set("zoom", 10)
options.Set("bearing", 90)
options.Set("pitch", 45)
mapInstance.EaseTo(options)
```

## API Reference

### Core Classes

- `MapboxGL`: Main namespace and entry point
- `Map`: The map instance with camera and interaction methods
- `MapOptions`: Configuration options for map creation

### Geographic Types

- `LngLat`: Longitude/latitude coordinate pair
- `LngLatBounds`: Geographic bounding box
- `Point`: Pixel coordinate point

### UI Elements

- `Marker`: Point markers on the map
- `MarkerOptions`: Marker configuration
- `Popup`: Information popups
- `PopupOptions`: Popup configuration

### Data & Styling

- `Source`: Base data source type
- `GeoJSONSource`: GeoJSON data source
- `VectorSource`: Vector tile source
- `RasterSource`: Raster tile source
- `ImageSource`: Image source
- `VideoSource`: Video source
- `Layer`: Map layer for rendering data

### Controls

- `NavigationControl`: Zoom and rotation controls
- `GeolocateControl`: User location control
- `AttributionControl`: Attribution display
- `ScaleControl`: Scale indicator
- `FullscreenControl`: Fullscreen toggle

### Events

All event handling is provided through the `Map` methods:

- `On(eventType, listener)` - Add event listener
- `Once(eventType, listener)` - Add one-time event listener  
- `Off(eventType, listener)` - Remove event listener
- `OnLayer(eventType, layerId, listener)` - Layer-specific events

## Requirements

- Go 1.18+ with WebAssembly support
- Mapbox GL JS library loaded in the browser
- Valid Mapbox access token

## Browser Setup

Include Mapbox GL JS in your HTML:

```html
<!DOCTYPE html>
<html>
<head>
    <script src='https://api.mapbox.com/mapbox-gl-js/v2.15.0/mapbox-gl.js'></script>
    <link href='https://api.mapbox.com/mapbox-gl-js/v2.15.0/mapbox-gl.css' rel='stylesheet' />
</head>
<body>
    <div id='map' style='width: 100%; height: 500px;'></div>
    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        });
    </script>
</body>
</html>
```

## Building for WebAssembly

```bash
GOOS=js GOARCH=wasm go build -o main.wasm main.go
```

## Notes

- All coordinate methods expect and return geographic coordinates in decimal degrees
- Zoom levels typically range from 0 (world view) to 22 (street level)
- Bearing values are in degrees, with 0 pointing north
- Pitch values are in degrees, with 0 being perpendicular to the surface
- Event callbacks receive raw `js.Value` objects - use the provided utility functions to extract typed data
- Memory management: Some callback functions may need manual cleanup depending on usage patterns

## License

This package follows the same license as the parent project.
