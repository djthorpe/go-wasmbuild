package mvc

import (
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
func Router(args ...any) RouterView {
	self := NewView(new(router), ViewRouter, "div", args).(RouterView)
	router := self.(*router)
	window := dom.GetWindow()

	// Set event listener for hashchange
	window.AddEventListener("hashchange", func(Event) {
		router.refresh(window.Location().Hash())
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

	// Append page to list of pages
	router.pages = append(router.pages, page{path: path, view: view})

	// Refresh the view for the current hash or default page
	router.refresh(dom.GetWindow().Location().Hash())

	// Append the view to the router body, with a DIV
	return router
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// match returns the page which matches the specified hash, or nil
func (router *router) match(hash string) *page {
	for _, page := range router.pages {
		if page.path == hash {
			return &page
		}
	}
	return nil
}

// refresh updates the router content based on the current hash. If no matching
// page exists, the first registered page (if any) becomes the default view.
func (router *router) refresh(hash string) {
	if page := router.match(hash); page != nil {
		router.Content(page.view)
		return
	}
	if len(router.pages) > 0 {
		router.Content(router.pages[0].view)
		return
	}
	router.Content()
}
