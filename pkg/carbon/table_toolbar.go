package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type tableToolbar struct{ base }

var _ mvc.View = (*tableToolbar)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewTableToolbar, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(tableToolbar), element, setView)
	})
}

// TableToolbar returns a <cds-table-toolbar> web component.
//
// Search children are inserted directly into the toolbar. Any other content
// is automatically grouped into a <cds-table-toolbar-content> wrapper so call
// sites can pass buttons and menus directly.
func TableToolbar(args ...any) *tableToolbar {
	return mvc.NewView(new(tableToolbar), ViewTableToolbar, "cds-table-toolbar", setView, args...).(*tableToolbar)
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
