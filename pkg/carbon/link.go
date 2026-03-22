package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type link struct{ base }

var _ mvc.View = (*link)(nil)
var _ mvc.EnabledState = (*link)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewLink, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(link), element, setView)
	}, EventClick)
}

// Link returns a <cds-link> web component.
// The href is always applied as the first argument; use With() for size and
// inline presentation, and pass Icon(...) to include a slotted Carbon icon.
func Link(href string, args ...any) *link {
	normalizeLinkArgs(args...)
	return mvc.NewView(new(link), ViewLink, "cds-link", setView, append([]any{mvc.WithAttr("href", href)}, args...)...).(*link)
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

func (l *link) Enabled() bool {
	return !boolProperty(l.Root(), "disabled")
}

func (l *link) SetEnabled(enabled bool) *link {
	setBoolProperty(l.Root(), "disabled", !enabled)
	return l
}

///////////////////////////////////////////////////////////////////////////////
// LABELLING

// Label returns the link's accessible name (aria-label).
func (l *link) Label() string {
	return l.Root().GetAttribute("aria-label")
}

// SetLabel sets or clears the link's accessible name (aria-label).
func (l *link) SetLabel(label string) *link {
	if label == "" {
		l.Root().RemoveAttribute("aria-label")
	} else {
		l.Root().SetAttribute("aria-label", label)
	}
	return l
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Rel returns the link rel attribute.
func (l *link) Rel() string {
	return l.Root().GetAttribute("rel")
}

// SetRel updates the link rel attribute.
func (l *link) SetRel(rel string) *link {
	if rel == "" {
		l.Root().RemoveAttribute("rel")
	} else {
		l.Root().SetAttribute("rel", rel)
	}
	return l
}

// Target returns the link target attribute.
func (l *link) Target() string {
	return l.Root().GetAttribute("target")
}

// SetTarget updates the link target attribute.
func (l *link) SetTarget(target string) *link {
	if target == "" {
		l.Root().RemoveAttribute("target")
	} else {
		l.Root().SetAttribute("target", target)
	}
	return l
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func normalizeLinkArgs(args ...any) {
	for _, arg := range args {
		switch value := arg.(type) {
		case *icon:
			applyLinkIconSlot(value)
		case []any:
			normalizeLinkArgs(value...)
		}
	}
}

func applyLinkIconSlot(icon *icon) {
	root := icon.Root()
	if icon == nil {
		root.RemoveAttribute("slot")
	} else {
		root.SetAttribute("slot", "icon")
		if root.GetAttribute("aria-hidden") == "" && root.GetAttribute("aria-label") == "" {
			root.SetAttribute("aria-hidden", "true")
		}
		style := strings.TrimSpace(root.GetAttribute("style"))
		if !strings.Contains(style, "color:") {
			if style == "" {
				root.SetAttribute("style", "color:currentColor")
			} else {
				root.SetAttribute("style", strings.TrimRight(style, "; ")+";color:currentColor")
			}
		}
	}
}
