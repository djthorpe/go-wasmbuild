package carbon

// TileDecorator returns a <cds-tile> for backward compatibility.
// Carbon's tile decorator treatment is supplied via slotted content assigned
// to the `decorator` slot rather than a host attribute.
func TileDecorator(args ...any) *tile {
	return Tile(args...)
}
