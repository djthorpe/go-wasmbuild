package bootstrap

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type list struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewList      = "mvc-bs-list"
	ViewListGroup = "mvc-bs-listgroup"
)

func init() {
	mvc.RegisterView(ViewList, newListFromElement)
	mvc.RegisterView(ViewListGroup, newListGroupFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func List(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(list), ViewList, "OL", opt...)
}

func ListGroup(opt ...mvc.Opt) mvc.View {
	opt = append([]mvc.Opt{mvc.WithClass("list-group")}, opt...)
	return mvc.NewView(new(list), ViewListGroup, "UL", opt...)
}

func BulletList(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", opt...)
}

func UnstyledList(opt ...mvc.Opt) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", append([]mvc.Opt{mvc.WithClass("list-unstyled")}, opt...)...)
}

func newListFromElement(element Element) mvc.View {
	if element.TagName() != "UL" {
		return nil
	}
	return mvc.NewViewWithElement(new(list), element)
}

func newListGroupFromElement(element Element) mvc.View {
	if element.TagName() != "UL" {
		return nil
	}
	return mvc.NewViewWithElement(new(list), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (list *list) SetView(view mvc.View) {
	list.View = view
}

func (list *list) Append(children ...any) mvc.View {
	// Wrap all children in divs with class "col"
	for _, child := range children {
		col := mvc.HTML("LI")
		if list.Name() == ViewListGroup {
			col.ClassList().Add("list-group-item")
		}
		col.AppendChild(mvc.NodeFromAny(child))
		list.View.Append(col)
	}
	return list
}
