package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tableToolbar struct{ base }

type tableToolbarContent struct{ base }

type tableToolbarSearch struct{ base }

var _ mvc.View = (*tableToolbar)(nil)
var _ mvc.View = (*tableToolbarContent)(nil)
var _ mvc.View = (*tableToolbarSearch)(nil)
var _ mvc.EnabledState = (*tableToolbarSearch)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const toolbarSearchInputEvent = "cds-search-input"

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTableToolbar, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableToolbar), element, setView)
	})
	mvc.RegisterView(ViewTableToolbarContent, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableToolbarContent), element, setView)
	})
	mvc.RegisterView(ViewTableToolbarSearch, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableToolbarSearch), element, setView)
	}, EventInput, EventChange, EventFocus, EventNoFocus)
}

// TableToolbar returns a <cds-table-toolbar> web component.
//
// Search children are inserted directly into the toolbar. Any other content
// is automatically grouped into a <cds-table-toolbar-content> wrapper so call
// sites can pass buttons and menus directly.
func TableToolbar(args ...any) *tableToolbar {
	return mvc.NewView(new(tableToolbar), ViewTableToolbar, "cds-table-toolbar", setView, args...).(*tableToolbar)
}

// TableToolbarContent returns a <cds-table-toolbar-content> wrapper.
func TableToolbarContent(args ...any) *tableToolbarContent {
	content := mvc.NewView(new(tableToolbarContent), ViewTableToolbarContent, "cds-table-toolbar-content", setView, args...).(*tableToolbarContent)
	return content
}

// TableToolbarSearch returns a <cds-table-toolbar-search> web component.
func TableToolbarSearch(args ...any) *tableToolbarSearch {
	search := mvc.NewView(new(tableToolbarSearch), ViewTableToolbarSearch, "cds-table-toolbar-search", setView, args...).(*tableToolbarSearch)
	if search.Label() == "" {
		search.SetLabel("Search")
	}
	return search
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TABLE TOOLBAR

func (t *tableToolbar) Content(args ...any) mvc.View {
	if t == nil {
		return nil
	}

	direct := make([]any, 0, len(args))
	actions := make([]any, 0, len(args))
	for _, arg := range args {
		switch value := arg.(type) {
		case *tableToolbarSearch:
			direct = append(direct, value)
		case *tableToolbarContent:
			direct = append(direct, value)
		default:
			actions = append(actions, value)
		}
	}
	if len(actions) > 0 {
		direct = append(direct, TableToolbarContent(actions...))
	}
	return t.View.Content(direct...)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - TABLE TOOLBAR SEARCH

func (s *tableToolbarSearch) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	if event == EventInput {
		s.View.AddEventListener(toolbarSearchInputEvent, handler)
		return s
	}
	return s.View.AddEventListener(event, handler)
}

func (s *tableToolbarSearch) RemoveEventListener(event string) mvc.View {
	if event == EventInput {
		s.View.RemoveEventListener(toolbarSearchInputEvent)
		return s
	}
	return s.View.RemoveEventListener(event)
}

func (s *tableToolbarSearch) Enabled() bool {
	return s != nil && !s.Root().HasAttribute("disabled")
}

func (s *tableToolbarSearch) SetEnabled(enabled bool) mvc.View {
	if s == nil {
		return nil
	}
	if enabled {
		s.Root().RemoveAttribute("disabled")
	} else {
		s.Root().SetAttribute("disabled", "")
	}
	return s
}

func (s *tableToolbarSearch) Value() string {
	if s == nil {
		return ""
	}
	if value := s.Root().Value(); value != "" {
		return value
	}
	return s.Root().GetAttribute("value")
}

func (s *tableToolbarSearch) SetValue(value string) *tableToolbarSearch {
	if s == nil {
		return nil
	}
	s.Root().SetValue(value)
	s.Root().SetAttribute("value", value)
	return s
}

func (s *tableToolbarSearch) Label() string {
	if s == nil {
		return ""
	}
	return s.Root().GetAttribute("label-text")
}

func (s *tableToolbarSearch) SetLabel(label string) *tableToolbarSearch {
	if s == nil {
		return nil
	}
	if label == "" {
		s.Root().RemoveAttribute("label-text")
	} else {
		s.Root().SetAttribute("label-text", label)
	}
	return s
}

func (s *tableToolbarSearch) Placeholder() string {
	if s == nil {
		return ""
	}
	return s.Root().GetAttribute("placeholder")
}

func (s *tableToolbarSearch) SetPlaceholder(placeholder string) *tableToolbarSearch {
	if s == nil {
		return nil
	}
	if placeholder == "" {
		s.Root().RemoveAttribute("placeholder")
	} else {
		s.Root().SetAttribute("placeholder", placeholder)
	}
	return s
}

func (s *tableToolbarSearch) Expanded() bool {
	return s != nil && s.Root().HasAttribute("expanded")
}

func (s *tableToolbarSearch) SetExpanded(expanded bool) *tableToolbarSearch {
	if s == nil {
		return nil
	}
	if expanded {
		s.Root().SetAttribute("expanded", "")
	} else {
		s.Root().RemoveAttribute("expanded")
	}
	return s
}
