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

func List(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "OL", args...)
}

func ListGroup(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewListGroup, "UL", mvc.WithClass("list-group"), args)
}

func BulletList(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", args...)
}

func UnstyledList(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", mvc.WithClass("list-unstyled"), args)
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
