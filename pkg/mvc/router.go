package mvc

import (
	"strings"

	// Packages
	wasm "github.com/djthorpe/go-wasmbuild"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
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
	r := new(router)
	r.View = NewView(r, ViewRouter, "div", args...)
	window := dom.GetWindow()

	// Set event listener for hashchange
	window.AddEventListener("hashchange", func(wasm.Event) {
		r.refresh(window.Location().Hash())
	})

	return r
}

// Create a Table from an existing element
func newRouterFromElement(element wasm.Element) View {
	if element.TagName() != "DIV" {
		return nil
	}
	r := new(router)
	r.View = NewViewWithElement(r, element)
	return r
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (router *router) Self() View {
	return router
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
