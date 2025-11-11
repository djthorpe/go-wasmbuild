//go:build js

package mapbox

import (
	"syscall/js"
)

// LngLat represents a longitude/latitude coordinate pair.
type LngLat struct {
	value js.Value
}

// NewLngLat creates a new LngLat object from longitude and latitude values.
func NewLngLat(lng, lat float64) *LngLat {
	lngLatInstance := js.Global().Get("mapboxgl").Get("LngLat").New(lng, lat)
	return &LngLat{value: lngLatInstance}
}

// LngLatFromValue wraps an existing js.Value as a LngLat.
func LngLatFromValue(value js.Value) *LngLat {
	return &LngLat{value: value}
}

// JSValue returns the underlying js.Value.
func (ll *LngLat) JSValue() js.Value {
	return ll.value
}

// Lng returns the longitude.
func (ll *LngLat) Lng() float64 {
	return ll.value.Get("lng").Float()
}

// Lat returns the latitude.
func (ll *LngLat) Lat() float64 {
	return ll.value.Get("lat").Float()
}

// SetLng sets the longitude.
func (ll *LngLat) SetLng(lng float64) {
	ll.value.Set("lng", lng)
}

// SetLat sets the latitude.
func (ll *LngLat) SetLat(lat float64) {
	ll.value.Set("lat", lat)
}

// ToArray returns the LngLat as a [lng, lat] array.
func (ll *LngLat) ToArray() []float64 {
	result := ll.value.Call("toArray")
	return []float64{
		result.Index(0).Float(),
		result.Index(1).Float(),
	}
}

// ToString returns a string representation of the LngLat.
func (ll *LngLat) ToString() string {
	return ll.value.Call("toString").String()
}

// DistanceTo returns the approximate distance between this LngLat and another.
// The distance is calculated in meters using the Haversine formula.
func (ll *LngLat) DistanceTo(other *LngLat) float64 {
	return ll.value.Call("distanceTo", other.JSValue()).Float()
}

// Wrap returns a new LngLat with longitude wrapped to the range [-180, 180].
func (ll *LngLat) Wrap() *LngLat {
	wrapped := ll.value.Call("wrap")
	return &LngLat{value: wrapped}
}

// ToBounds creates a LngLatBounds from this LngLat with the specified radius.
// The radius is in degrees.
func (ll *LngLat) ToBounds(radius float64) *LngLatBounds {
	bounds := ll.value.Call("toBounds", radius)
	return &LngLatBounds{value: bounds}
}

// LngLatBounds represents a geographical bounding box.
type LngLatBounds struct {
	value js.Value
}

// NewLngLatBounds creates a new LngLatBounds object.
// If sw and ne are provided, creates bounds from southwest and northeast corners.
// Otherwise, creates an empty bounds that can be extended.
func NewLngLatBounds(sw, ne *LngLat) *LngLatBounds {
	var boundsInstance js.Value
	mapboxgl := js.Global().Get("mapboxgl")

	if sw != nil && ne != nil {
		boundsInstance = mapboxgl.Get("LngLatBounds").New(sw.JSValue(), ne.JSValue())
	} else {
		boundsInstance = mapboxgl.Get("LngLatBounds").New()
	}
	return &LngLatBounds{value: boundsInstance}
}

// LngLatBoundsFromArray creates LngLatBounds from a nested array.
// Accepts formats like [[west, south], [east, north]] or [west, south, east, north].
func LngLatBoundsFromArray(bounds [][]float64) *LngLatBounds {
	var jsArray js.Value
	if len(bounds) == 2 && len(bounds[0]) == 2 && len(bounds[1]) == 2 {
		// [[west, south], [east, north]] format
		jsArray = js.Global().Get("Array").New()
		sw := js.Global().Get("Array").New()
		sw.SetIndex(0, bounds[0][0])
		sw.SetIndex(1, bounds[0][1])
		ne := js.Global().Get("Array").New()
		ne.SetIndex(0, bounds[1][0])
		ne.SetIndex(1, bounds[1][1])
		jsArray.SetIndex(0, sw)
		jsArray.SetIndex(1, ne)
	}

	boundsInstance := js.Global().Get("mapboxgl").Get("LngLatBounds").New(jsArray)
	return &LngLatBounds{value: boundsInstance}
}

// LngLatBoundsFromValue wraps an existing js.Value as a LngLatBounds.
func LngLatBoundsFromValue(value js.Value) *LngLatBounds {
	return &LngLatBounds{value: value}
}

// JSValue returns the underlying js.Value.
func (llb *LngLatBounds) JSValue() js.Value {
	return llb.value
}

// GetSouthWest returns the southwest corner of the bounds.
func (llb *LngLatBounds) GetSouthWest() *LngLat {
	sw := llb.value.Call("getSouthWest")
	return &LngLat{value: sw}
}

// GetNorthEast returns the northeast corner of the bounds.
func (llb *LngLatBounds) GetNorthEast() *LngLat {
	ne := llb.value.Call("getNorthEast")
	return &LngLat{value: ne}
}

// GetNorthWest returns the northwest corner of the bounds.
func (llb *LngLatBounds) GetNorthWest() *LngLat {
	nw := llb.value.Call("getNorthWest")
	return &LngLat{value: nw}
}

// GetSouthEast returns the southeast corner of the bounds.
func (llb *LngLatBounds) GetSouthEast() *LngLat {
	se := llb.value.Call("getSouthEast")
	return &LngLat{value: se}
}

// GetWest returns the westernmost longitude.
func (llb *LngLatBounds) GetWest() float64 {
	return llb.value.Call("getWest").Float()
}

// GetSouth returns the southernmost latitude.
func (llb *LngLatBounds) GetSouth() float64 {
	return llb.value.Call("getSouth").Float()
}

// GetEast returns the easternmost longitude.
func (llb *LngLatBounds) GetEast() float64 {
	return llb.value.Call("getEast").Float()
}

// GetNorth returns the northernmost latitude.
func (llb *LngLatBounds) GetNorth() float64 {
	return llb.value.Call("getNorth").Float()
}

// GetCenter returns the center of the bounds as a LngLat.
func (llb *LngLatBounds) GetCenter() *LngLat {
	center := llb.value.Call("getCenter")
	return &LngLat{value: center}
}

// Extend extends the bounds to include a LngLat or another LngLatBounds.
func (llb *LngLatBounds) Extend(obj interface{}) *LngLatBounds {
	var jsObj js.Value
	switch v := obj.(type) {
	case *LngLat:
		jsObj = v.JSValue()
	case *LngLatBounds:
		jsObj = v.JSValue()
	default:
		return llb // Return unchanged if invalid type
	}

	extended := llb.value.Call("extend", jsObj)
	return &LngLatBounds{value: extended}
}

// Contains checks if the bounds contain a given LngLat.
func (llb *LngLatBounds) Contains(lngLat *LngLat) bool {
	return llb.value.Call("contains", lngLat.JSValue()).Bool()
}

// Intersects checks if these bounds intersect with another LngLatBounds.
func (llb *LngLatBounds) Intersects(other *LngLatBounds) bool {
	return llb.value.Call("intersects", other.JSValue()).Bool()
}

// ToArray returns the bounds as a nested array: [[west, south], [east, north]].
func (llb *LngLatBounds) ToArray() [][]float64 {
	result := llb.value.Call("toArray")
	sw := result.Index(0)
	ne := result.Index(1)

	return [][]float64{
		{sw.Index(0).Float(), sw.Index(1).Float()},
		{ne.Index(0).Float(), ne.Index(1).Float()},
	}
}

// ToString returns a string representation of the bounds.
func (llb *LngLatBounds) ToString() string {
	return llb.value.Call("toString").String()
}

// IsEmpty checks if the bounds are empty (have no area).
func (llb *LngLatBounds) IsEmpty() bool {
	return llb.value.Call("isEmpty").Bool()
}

// Point represents a pixel coordinate point.
type Point struct {
	value js.Value
}

// NewPoint creates a new Point with x and y coordinates.
func NewPoint(x, y float64) *Point {
	pointInstance := js.Global().Get("mapboxgl").Get("Point").New(x, y)
	return &Point{value: pointInstance}
}

// PointFromValue wraps an existing js.Value as a Point.
func PointFromValue(value js.Value) *Point {
	return &Point{value: value}
}

// JSValue returns the underlying js.Value.
func (p *Point) JSValue() js.Value {
	return p.value
}

// X returns the x coordinate.
func (p *Point) X() float64 {
	return p.value.Get("x").Float()
}

// Y returns the y coordinate.
func (p *Point) Y() float64 {
	return p.value.Get("y").Float()
}

// SetX sets the x coordinate.
func (p *Point) SetX(x float64) {
	p.value.Set("x", x)
}

// SetY sets the y coordinate.
func (p *Point) SetY(y float64) {
	p.value.Set("y", y)
}

// Clone returns a copy of the Point.
func (p *Point) Clone() *Point {
	cloned := p.value.Call("clone")
	return &Point{value: cloned}
}

// Add returns a new Point that is the sum of this Point and another.
func (p *Point) Add(other *Point) *Point {
	result := p.value.Call("add", other.JSValue())
	return &Point{value: result}
}

// Sub returns a new Point that is the difference of this Point and another.
func (p *Point) Sub(other *Point) *Point {
	result := p.value.Call("sub", other.JSValue())
	return &Point{value: result}
}

// Mult returns a new Point with coordinates multiplied by a scalar.
func (p *Point) Mult(scalar float64) *Point {
	result := p.value.Call("mult", scalar)
	return &Point{value: result}
}

// Div returns a new Point with coordinates divided by a scalar.
func (p *Point) Div(scalar float64) *Point {
	result := p.value.Call("div", scalar)
	return &Point{value: result}
}

// Rotate returns a new Point rotated around the origin by an angle in radians.
func (p *Point) Rotate(angle float64) *Point {
	result := p.value.Call("rotate", angle)
	return &Point{value: result}
}

// RotateAround returns a new Point rotated around another Point by an angle in radians.
func (p *Point) RotateAround(pivot *Point, angle float64) *Point {
	result := p.value.Call("rotateAround", pivot.JSValue(), angle)
	return &Point{value: result}
}

// MatMult applies a 2D matrix transformation to the Point.
func (p *Point) MatMult(matrix [4]float64) *Point {
	jsMatrix := js.Global().Get("Array").New()
	for i, val := range matrix {
		jsMatrix.SetIndex(i, val)
	}
	result := p.value.Call("matMult", jsMatrix)
	return &Point{value: result}
}

// Unit returns a unit vector (length 1) in the same direction as this Point.
func (p *Point) Unit() *Point {
	result := p.value.Call("unit")
	return &Point{value: result}
}

// Perp returns a Point perpendicular to this Point.
func (p *Point) Perp() *Point {
	result := p.value.Call("perp")
	return &Point{value: result}
}

// Round returns a new Point with coordinates rounded to the nearest integer.
func (p *Point) Round() *Point {
	result := p.value.Call("round")
	return &Point{value: result}
}

// Mag returns the magnitude (length) of the Point.
func (p *Point) Mag() float64 {
	return p.value.Call("mag").Float()
}

// Equals checks if this Point is equal to another Point.
func (p *Point) Equals(other *Point) bool {
	return p.value.Call("equals", other.JSValue()).Bool()
}

// Dist returns the distance between this Point and another Point.
func (p *Point) Dist(other *Point) float64 {
	return p.value.Call("dist", other.JSValue()).Float()
}

// DistSqr returns the squared distance between this Point and another Point.
func (p *Point) DistSqr(other *Point) float64 {
	return p.value.Call("distSqr", other.JSValue()).Float()
}

// Angle returns the angle of the Point in radians.
func (p *Point) Angle() float64 {
	return p.value.Call("angle").Float()
}

// AngleTo returns the angle from this Point to another Point in radians.
func (p *Point) AngleTo(other *Point) float64 {
	return p.value.Call("angleTo", other.JSValue()).Float()
}

// AngleWith returns the angle between this Point and another Point in radians.
func (p *Point) AngleWith(other *Point) float64 {
	return p.value.Call("angleWith", other.JSValue()).Float()
}

// AngleWithSep returns the signed angle between this Point and another Point in radians.
func (p *Point) AngleWithSep(other *Point) float64 {
	return p.value.Call("angleWithSep", other.JSValue()).Float()
}

// ToArray returns the Point as a [x, y] array.
func (p *Point) ToArray() []float64 {
	return []float64{p.X(), p.Y()}
}

// ToString returns a string representation of the Point.
func (p *Point) ToString() string {
	return p.value.Call("toString").String()
}
