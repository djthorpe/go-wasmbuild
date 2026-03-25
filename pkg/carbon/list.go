package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type list struct{ base }

var _ mvc.View = (*list)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewList, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(list), element, setView)
	})
}

// List returns an unordered list (<ul>) view.
func List(args ...any) *list {
	l := mvc.NewView(new(list), ViewList, "UL", setView, args).(*list)
	l.syncPresentation()
	return l
}

func (l *list) Apply(opts ...mvc.Opt) mvc.View {
	l.View.Apply(opts...)
	l.syncPresentation()
	return l
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (l *list) Content(args ...any) mvc.View {
	if tag := l.Root().TagName(); tag == "UL" || tag == "OL" {
		children := make([]any, 0, len(args))
		for _, arg := range args {
			if child, ok := arg.(*list); ok && child.Root().TagName() == "LI" {
				children = append(children, child)
			} else {
				children = append(children, ListItem(arg))
			}
		}
		return l.View.Content(children...)
	}
	return l.View.Content(args...)
}

func (l *list) syncPresentation() {
	root := l.Root()
	tag := root.TagName()
	if tag != "UL" && tag != "OL" {
		return
	}
	styleType := strings.TrimSpace(root.GetAttribute("data-carbon-list-style"))
	generatedStyle := strings.TrimSpace(root.GetAttribute("data-carbon-style-list"))
	baseStyle := strings.TrimSpace(strings.ReplaceAll(root.GetAttribute("style"), generatedStyle, ""))
	baseStyle = strings.Trim(baseStyle, "; ")
	if styleType == "" {
		root.RemoveAttribute("data-carbon-style-list")
		if baseStyle == "" {
			root.RemoveAttribute("style")
		} else {
			root.SetAttribute("style", baseStyle)
		}
		return
	}
	generatedStyle = "list-style-type:" + styleType
	root.SetAttribute("data-carbon-style-list", generatedStyle)
	if baseStyle == "" {
		root.SetAttribute("style", generatedStyle)
	} else {
		root.SetAttribute("style", baseStyle+";"+generatedStyle)
	}
}
