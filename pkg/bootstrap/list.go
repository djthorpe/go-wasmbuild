package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
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

func init() {
	mvc.RegisterView(ViewList, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(list), element, setView)
	})
	mvc.RegisterView(ViewListGroup, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(list), element, setView)
	})
	mvc.RegisterView(ViewDefinitionList, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(deflist), element, setView)
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func List(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "OL", setView, args)
}

func BulletList(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", setView, args)
}

func UnstyledList(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewList, "UL", setView, mvc.WithClass("list-unstyled"), args)
}

func ListGroup(args ...any) mvc.View {
	return mvc.NewView(new(list), ViewListGroup, "UL", setView, mvc.WithClass("list-group"), args)
}

func DefinitionList(args ...any) mvc.View {
	return mvc.NewView(new(deflist), ViewDefinitionList, "DL", setView, mvc.WithClass("row"), args)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

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
		case *option:
			nodes = append(nodes, mvc.HTML("DT", mvc.WithClass("col-3"), mvc.WithInnerText(child.Name)))
			nodes = append(nodes, mvc.HTML("DD", mvc.WithClass("col-9"), mvc.WithInnerText(child.Value)))
		default:
			panic("Content[deflist]: child must be of type Option")
		}
	}
	return deflist.View.Content(nodes...)
}
