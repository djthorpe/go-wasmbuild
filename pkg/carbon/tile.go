package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tile struct{ base }

var _ mvc.View = (*tile)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTile, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tile), element, setView)
	})
}

// Tile returns a <cds-tile> web component — a static content container
// analogous to Bootstrap's Card.
//
//	carbon.Tile(carbon.Head(3, "Title"), carbon.Para("Body text"))
func Tile(args ...any) *tile {
	tile := mvc.NewView(new(tile), ViewTile, "cds-tile", setView, args).(*tile)
	tile.applyPresentation()
	return tile
}

// TileDecorator returns a <cds-tile> for backward compatibility.
// Carbon's tile decorator treatment is supplied via slotted content assigned
// to the `decorator` slot rather than a host attribute.
//
//	carbon.Tile(
//		mvc.HTML("SPAN", mvc.WithAttr("slot", "decorator"), "AI"),
//		carbon.Head(3, "Title"),
//		carbon.Para("Body text"),
//	)
func TileDecorator(args ...any) *tile {
	return Tile(args...)
}

func (t *tile) Apply(opts ...mvc.Opt) mvc.View {
	t.View.Apply(opts...)
	t.applyPresentation()
	return t
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (t *tile) applyPresentation() {
	root := t.Root()
	generatedStyle := strings.TrimSpace(root.GetAttribute("data-carbon-style-tile"))
	baseStyle := strings.TrimSpace(strings.ReplaceAll(root.GetAttribute("style"), generatedStyle, ""))
	baseStyle = strings.Trim(baseStyle, "; ")

	style := make([]string, 0, 3)
	if layer := root.GetAttribute("data-carbon-layer"); layer != "" {
		style = append(style, "--cds-layer:"+layer)
	}
	if root.GetAttribute("data-carbon-fill") == "true" {
		style = append(style, "width:100%", "display:block")
	}
	if height := root.GetAttribute("data-carbon-height"); height != "" {
		style = append(style, "height:"+height)
	}
	generatedStyle = strings.Join(style, ";")
	root.SetAttribute("data-carbon-style-tile", generatedStyle)
	if style := appendTileStyle(baseStyle, generatedStyle); style == "" {
		root.RemoveAttribute("style")
	} else {
		root.SetAttribute("style", style)
	}
}

func appendTileStyle(base, extra string) string {
	base = strings.Trim(base, "; ")
	extra = strings.Trim(extra, "; ")
	switch {
	case base == "":
		return extra
	case extra == "":
		return base
	default:
		return base + ";" + extra
	}
}
