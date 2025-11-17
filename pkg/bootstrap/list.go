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

type deflist struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewList           = "mvc-bs-list"
	ViewListGroup      = "mvc-bs-listgroup"
	ViewDefinitionList = "mvc-bs-deflist"
)

func init() {
	mvc.RegisterView(ViewList, newListFromElement)
	mvc.RegisterView(ViewListGroup, newListGroupFromElement)
	mvc.RegisterView(ViewDefinitionList, newDefinitionListFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func List(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "OL", args...)
}

func ListGroup(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewListGroup, "UL", mvc.WithClass("list-group"), args)
}

func DefinitionList(args ...any) mvc.View {
	return mvc.NewView(new(deflist), ViewDefinitionList, "DL", mvc.WithClass("row"), args)
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

func newDefinitionListFromElement(element Element) mvc.View {
	if element.TagName() != "DL" {
		return nil
	}
	return mvc.NewViewWithElement(new(deflist), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (list *list) SetView(view mvc.View) {
	list.View = view
}

func (deflist *deflist) SetView(view mvc.View) {
	deflist.View = view
}

func (list *list) Content(args ...any) mvc.View {
	nodes := make([]any, 0, len(args))
	for _, child := range args {
		col := mvc.HTML("LI")
		if list.Name() == ViewListGroup {
			col.ClassList().Add("list-group-item")
		}
		col.AppendChild(mvc.NodeFromAny(child))
		nodes = append(nodes, col)
	}
	return list.View.Content(nodes...)
}

func (deflist *deflist) Content(args ...any) mvc.View {
	nodes := make([]any, 0, len(args))
	for _, child := range args {
		switch child := child.(type) {
		case *inputoption:
			nodes = append(nodes, mvc.HTML("DT", mvc.WithClass("col-3"), mvc.WithInnerText(child.Name)))
			nodes = append(nodes, mvc.HTML("DD", mvc.WithClass("col-9"), mvc.WithInnerText(child.Value)))
		default:
			panic("Content[deflist]: child must be of type Option")
		}
	}
	return deflist.View.Content(nodes...)
}
