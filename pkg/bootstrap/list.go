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
	BootstrapView
}

type deflist struct {
	BootstrapView
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

func List(args ...any) *list {
	l := new(list)
	l.BootstrapView.View = mvc.NewView(l, ViewList, "OL", args...)
	return l
}

func ListGroup(args ...any) *list {
	l := new(list)
	l.BootstrapView.View = mvc.NewView(l, ViewListGroup, "UL", mvc.WithClass("list-group"), args)
	return l
}

func DefinitionList(args ...any) *deflist {
	d := new(deflist)
	d.BootstrapView.View = mvc.NewView(d, ViewDefinitionList, "DL", mvc.WithClass("row"), args)
	return d
}

func BulletList(args ...any) *list {
	l := new(list)
	l.BootstrapView.View = mvc.NewView(l, ViewList, "UL", args...)
	return l
}

func UnstyledList(args ...any) *list {
	l := new(list)
	l.BootstrapView.View = mvc.NewView(l, ViewList, "UL", mvc.WithClass("list-unstyled"), args)
	return l
}

func newListFromElement(element Element) mvc.View {
	if element.TagName() != "UL" {
		return nil
	}
	l := new(list)
	l.BootstrapView.View = mvc.NewViewWithElement(l, element)
	return l
}

func newListGroupFromElement(element Element) mvc.View {
	if element.TagName() != "UL" {
		return nil
	}
	l := new(list)
	l.BootstrapView.View = mvc.NewViewWithElement(l, element)
	return l
}

func newDefinitionListFromElement(element Element) mvc.View {
	if element.TagName() != "DL" {
		return nil
	}
	d := new(deflist)
	d.BootstrapView.View = mvc.NewViewWithElement(d, element)
	return d
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (list *list) Self() mvc.View {
	return list
}

func (deflist *deflist) Self() mvc.View {
	return deflist
}

func (list *list) Content(args ...any) *list {
	nodes := make([]any, 0, len(args))
	for _, child := range args {
		col := mvc.HTML("LI")
		if list.Name() == ViewListGroup {
			col.ClassList().Add("list-group-item")
		}
		col.AppendChild(mvc.NodeFromAny(child))
		nodes = append(nodes, col)
	}
	if len(nodes) == 1 {
		list.ReplaceSlot("body", nodes[0])
		return list
	}
	div := mvc.HTML("div")
	for _, node := range nodes {
		div.AppendChild(mvc.NodeFromAny(node))
	}
	list.ReplaceSlot("body", div)
	return list
}

func (deflist *deflist) Content(args ...any) *deflist {
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
	if len(nodes) == 1 {
		deflist.ReplaceSlot("body", nodes[0])
		return deflist
	}
	div := mvc.HTML("div")
	for _, node := range nodes {
		div.AppendChild(mvc.NodeFromAny(node))
	}
	deflist.ReplaceSlot("body", div)
	return deflist
}
