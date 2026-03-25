package carbon

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type structuredList struct {
	base
	changeBaseline dom.Element
}

var _ mvc.View = (*structuredList)(nil)
var _ mvc.ActiveGroup = (*structuredList)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const templateStructuredList = `
	<cds-structured-list>
		<cds-structured-list-head data-slot="header"></cds-structured-list-head>
		<cds-structured-list-body data-slot="body"></cds-structured-list-body>
	</cds-structured-list>
`

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewStructuredList, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(structuredList), element, setView)
	}, EventChange)
}

// StructuredList returns a Carbon structured list with a header slot and body slot.
func StructuredList(args ...any) *structuredList {
	s := mvc.NewView(new(structuredList), ViewStructuredList, templateStructuredList, setView, args).(*structuredList)
	s.syncPresentation()
	return s
}

// AddEventListener registers an event handler on the structured list.
// EventChange is bridged from row selection changes because Carbon does not
// emit a dedicated host change event for structured list selection.
func (s *structuredList) AddEventListener(event string, handler func(dom.Event)) mvc.View {
	if event == EventChange {
		ensureStructuredListChangeBridge(s)
	}
	s.View.AddEventListener(event, handler)
	return s
}

func (s *structuredList) Apply(opts ...mvc.Opt) mvc.View {
	s.View.Apply(opts...)
	s.syncPresentation()
	return s
}

func (s *structuredList) Content(args ...any) mvc.View {
	headers := make([]any, 0, 1)
	rows := make([]any, 0, len(args))
	for _, arg := range args {
		switch value := arg.(type) {
		case *structuredListHeader:
			headers = append(headers, value)
		case *structuredListRow:
			rows = append(rows, value)
		default:
			rows = append(rows, StructuredListRow(arg))
		}
	}
	if len(headers) > 0 {
		s.View.ReplaceSlotChildren("header", headers...)
	}
	if len(rows) > 0 {
		s.View.ReplaceSlotChildren("body", rows...)
	}
	s.syncPresentation()
	return s
}

// Active returns the currently selected rows.
// In Carbon's selectable structured list variant this will contain at most one row.
func (s *structuredList) Active() []mvc.View {
	active := make([]mvc.View, 0, 1)
	for _, child := range structuredListRowElements(s) {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if row, ok := v.(*structuredListRow); ok && row.Active() {
				active = append(active, row)
			}
		}
	}
	return active
}

// SetActive selects the first supplied row and deselects all others.
// Calling SetActive with no arguments clears the current selection.
func (s *structuredList) SetActive(views ...mvc.View) mvc.View {
	var target dom.Element
	for _, view := range views {
		if view != nil {
			target = view.Root()
			break
		}
	}
	for _, child := range structuredListRowElements(s) {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if row, ok := v.(*structuredListRow); ok {
				row.SetActive(target != nil && child.Equals(target))
			}
		}
	}
	return s
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func isStructuredListCell(arg any, tag string) bool {
	switch value := arg.(type) {
	case dom.Element:
		return strings.EqualFold(value.TagName(), tag)
	case mvc.View:
		return strings.EqualFold(value.Root().TagName(), tag)
	default:
		return false
	}
}

func structuredListRowElements(s *structuredList) []dom.Element {
	for _, child := range s.Root().Children() {
		if strings.EqualFold(child.TagName(), "CDS-STRUCTURED-LIST-BODY") {
			return child.Children()
		}
	}
	return nil
}

func structuredListHeaderRowElements(s *structuredList) []dom.Element {
	for _, child := range s.Root().Children() {
		if strings.EqualFold(child.TagName(), "CDS-STRUCTURED-LIST-HEAD") {
			return child.Children()
		}
	}
	return nil
}

func (s *structuredList) syncPresentation() {
	selectionName := s.Root().GetAttribute("selection-name")
	condensed := s.Root().HasAttribute("condensed")
	flush := s.Root().HasAttribute("flush")
	for _, row := range append(structuredListHeaderRowElements(s), structuredListRowElements(s)...) {
		syncStructuredListRowAttr(row, "selection-name", selectionName != "", selectionName)
		syncStructuredListRowAttr(row, "condensed", condensed, "")
		syncStructuredListRowAttr(row, "flush", flush, "")
	}
}

func syncStructuredListRowAttr(element dom.Element, name string, enabled bool, value string) {
	if enabled {
		element.SetAttribute(name, value)
	} else {
		element.RemoveAttribute(name)
	}
}

func structuredListActiveRowElement(s *structuredList) dom.Element {
	for _, child := range structuredListRowElements(s) {
		if v, err := mvc.ViewFromElement(child); err == nil {
			if row, ok := v.(*structuredListRow); ok && row.Active() {
				return child
			}
		}
	}
	return nil
}
