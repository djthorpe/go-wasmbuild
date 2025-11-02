package bs

import (
	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// containers are elements to wrap any content
type container struct {
	View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewContainer = "mvc-bs-container"
)

func init() {
	RegisterView(ViewContainer, newContainerFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Container(opt ...Opt) View {
	opts := append([]Opt{WithClass("container")}, opt...)
	return NewView(new(container), ViewContainer, "DIV", opts...)
}

func newContainerFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return NewViewWithElement(new(container), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (container *container) SetView(view View) {
	container.View = view
}
