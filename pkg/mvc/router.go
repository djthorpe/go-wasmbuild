package mvc

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// INTERFACE

type RouterView interface {
	View

	// Append a page view to the router with a specific path
	Page(path string, view View) RouterView
}

///////////////////////////////////////////////////////////////////////////////
// TYPES

type router struct {
	View
	pages []page
}

type page struct {
	path string
	view View
}

var _ RouterView = (*router)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewRouter = "mvc-router"
)

func init() {
	RegisterView(ViewRouter, newRouterFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a Router
func Router(opts ...Opt) RouterView {
	self := NewView(new(router), ViewRouter, "div", opts...).(RouterView)

	// Set event listener for hashchange
	dom.GetWindow().AddEventListener("hashchange", func(event Event) {
		if window, ok := event.Target().(Window); ok {
			fmt.Println("hashchange to:", window.Location().Hash())
		}
	})

	return self
}

// Create a Table from an existing element
func newRouterFromElement(element Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	return NewViewWithElement(new(router), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (router *router) SetView(view View) {
	router.View = view
}

func (router *router) Page(path string, view View) RouterView {
	if path != "" && !strings.HasPrefix(path, "#") {
		panic("Router.Page: path must start with '#'")
	}

	// Append page to list
	// TODO
	router.pages = append(router.pages, page{path: path, view: view})

	// Append the view to the router body, with a DIV
	return router
}
