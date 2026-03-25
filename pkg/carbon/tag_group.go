package carbon

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

type tagGroup struct{ base }

var _ mvc.View = (*tagGroup)(nil)
var _ mvc.EnabledGroup = (*tagGroup)(nil)
var _ mvc.ActiveGroup = (*tagGroup)(nil)
var _ mvc.VisibleGroup = (*tagGroup)(nil)

func init() {
	mvc.RegisterView(ViewTagGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tagGroup), element, setView)
	}, EventTagDismissibleClosed, EventTagOperationalSelected)
}

// TagGroup returns a container for one or more tags.
// Child tag events bubble to the group, allowing group-level observation.
func TagGroup(args ...any) *tagGroup {
	args = append([]any{mvc.WithStyle("display:flex;flex-wrap:wrap;align-items:center;gap:0.75rem")}, args...)
	return mvc.NewView(new(tagGroup), ViewTagGroup, "DIV", setView, args).(*tagGroup)
}

// Content appends tags to the group. Panics if any arg is not a *tag.
func (g *tagGroup) Content(args ...any) mvc.View {
	children := make([]any, 0, len(args))
	for _, arg := range args {
		if t, ok := arg.(*tag); ok {
			children = append(children, t)
		} else {
			panic(fmt.Sprintf("TagGroup.Content: expected *tag, got %T", arg))
		}
	}
	return g.View.Content(children...)
}

func (g *tagGroup) Active() []mvc.View {
	active := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Active() {
			active = append(active, t)
		}
	}
	return active
}

func (g *tagGroup) SetActive(views ...mvc.View) mvc.View {
	active := make(map[mvc.View]bool, len(views))
	for _, view := range views {
		active[view] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetActive(active[child])
		}
	}
	return g
}

func (g *tagGroup) Enabled() []mvc.View {
	enabled := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Enabled() {
			enabled = append(enabled, t)
		}
	}
	return enabled
}

func (g *tagGroup) SetEnabled(views ...mvc.View) mvc.View {
	enabled := make(map[mvc.View]bool, len(views))
	for _, view := range views {
		enabled[view] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetEnabled(enabled[child])
		}
	}
	return g
}

func (g *tagGroup) Visible() []mvc.View {
	visible := make([]mvc.View, 0)
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok && t.Visible() {
			visible = append(visible, t)
		}
	}
	return visible
}

func (g *tagGroup) SetVisible(views ...mvc.View) mvc.View {
	visible := make(map[mvc.View]bool, len(views))
	for _, view := range views {
		visible[view] = true
	}
	for _, child := range g.Children() {
		if t, ok := child.(*tag); ok {
			t.SetVisible(visible[child])
		}
	}
	return g
}
