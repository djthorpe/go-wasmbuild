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
var _ mvc.ActiveState = (*tile)(nil)

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
	return mvc.NewView(new(tile), ViewTile, "cds-tile", setView, args).(*tile)
}

// TileDecorator returns a <cds-tile> with the decorator variant attribute set,
// producing the accented left-border style.
//
//	carbon.TileDecorator(carbon.Head(3, "Title"), carbon.Para("Body text"))
func TileDecorator(args ...any) *tile {
	return mvc.NewView(new(tile), ViewTile, "cds-tile", setView,
		mvc.WithAttr("decorator", ""), args).(*tile)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// SetBackground sets the background colour of the tile by overriding the
// --cds-layer CSS custom property, which the shadow DOM uses internally.
func (t *tile) SetBackground(color string) *tile {
	t.Root().SetAttribute("data-carbon-layer", color)
	t.applyPresentation()
	return t
}

// SetFill controls whether the tile fills the width of its container.
func (t *tile) SetFill(fill bool) *tile {
	if fill {
		t.Root().SetAttribute("data-carbon-fill", "true")
	} else {
		t.Root().RemoveAttribute("data-carbon-fill")
	}
	t.applyPresentation()
	return t
}

// SetHeight sets a fixed CSS height for the tile host.
// Pass an empty string to clear the explicit height.
func (t *tile) SetHeight(height string) *tile {
	if height == "" {
		t.Root().RemoveAttribute("data-carbon-height")
	} else {
		t.Root().SetAttribute("data-carbon-height", height)
	}
	t.applyPresentation()
	return t
}

// Active reports whether the tile is in its active state.
func (t *tile) Active() bool {
	return t.Root().GetAttribute("data-carbon-active") == "true"
}

// SetActive applies a transient active presentation to the tile.
// The active state increases contrast and slightly lifts the tile, with a
// smooth transition back to the inactive state.
func (t *tile) SetActive(active bool) *tile {
	if active {
		t.Root().SetAttribute("data-carbon-active", "true")
	} else {
		t.Root().RemoveAttribute("data-carbon-active")
	}
	t.applyPresentation()
	return t
}

func (t *tile) applyPresentation() {
	root := t.Root()
	baseStyle := root.GetAttribute("data-carbon-style-base")
	if baseStyle == "" {
		baseStyle = root.GetAttribute("style")
		root.SetAttribute("data-carbon-style-base", baseStyle)
	}

	style := make([]string, 0, 4)
	if baseStyle != "" {
		style = append(style, baseStyle)
	}
	if layer := root.GetAttribute("data-carbon-layer"); layer != "" {
		style = append(style, "--cds-layer:"+layer)
	}
	if root.GetAttribute("data-carbon-fill") == "true" {
		style = append(style, "width:100%", "display:block")
	}
	if height := root.GetAttribute("data-carbon-height"); height != "" {
		style = append(style, "height:"+height)
	}
	style = append(style, "transition:filter 180ms ease,transform 180ms ease,box-shadow 180ms ease")
	if root.GetAttribute("data-carbon-active") == "true" {
		style = append(style,
			"filter:contrast(1.12) brightness(0.94) saturate(0.92)",
			"transform:translateY(-1px)",
			"box-shadow:0 0 0 1px var(--cds-border-strong-01,#8d8d8d)",
		)
	} else {
		style = append(style,
			"filter:contrast(1) brightness(1) saturate(1)",
			"transform:translateY(0)",
			"box-shadow:none",
		)
	}
	root.SetAttribute("style", strings.Join(style, ";"))
}
