package mvc

import (
	"strings"

	// Packages
	wasm "github.com/djthorpe/go-wasmbuild"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type router struct {
	View
	pages []page
	sel   ActiveGroup
}

type page struct {
	path string
	view View
	sel  []View
}

var _ View = (*router)(nil)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewRouter = "mvc-router"
)

func init() {
	RegisterView(ViewRouter, func(element wasm.Element) View {
		return NewViewWithElement(new(router), element, func(self, child View) {
			self.(*router).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a Router
func Router(args ...any) *router {
	self := NewView(new(router), ViewRouter, "DIV", func(self, child View) {
		self.(*router).View = child
	}, args).(*router)
	window := dom.GetWindow()

	// Set event listener for hashchange
	window.AddEventListener("hashchange", func(wasm.Event) {
		self.refresh(window.Location().Hash())
	})

	return self
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// SetActive registers an ActiveGroup (e.g. a SideNav) that the router will
// update whenever the active page changes.
func (router *router) Active(sel ActiveGroup) *router {
	router.sel = sel
	return router
}

func (router *router) Page(path string, view View, sel ...View) *router {
	if path != "" && !strings.HasPrefix(path, "#") {
		panic("Router.Page: path must start with '#'")
	}

	// Append page to list of pages
	router.pages = append(router.pages, page{path: path, view: view, sel: sel})

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
	var active *page
	if p := router.match(hash); p != nil {
		active = p
	} else if len(router.pages) > 0 {
		active = &router.pages[0]
	}
	if active == nil {
		return
	}
	router.ReplaceSlotChildren(ContentSlot, active.view)
	if router.sel != nil {
		router.sel.SetActive(active.sel...)
	}
}
