package bootstrap

import (
	// Packages

	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination struct {
	mvc.View
}

type paginationitem struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewPagination     = "mvc-bs-pagination"
	ViewPaginationItem = "mvc-bs-paginationitem"
)

const (
	templatePagination = `
		<nav><ul class="pagination" data-slot></ul></nav>	
	`
	templatePaginationItem = `
		<li class="page-item"><a class="page-link" role="button" data-slot></a></li>
	`
)

func init() {
	mvc.RegisterView(ViewPagination, newPaginationFromElement)
	mvc.RegisterView(ViewPaginationItem, newPaginationItemFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Pagination(args ...any) *pagination {
	return mvc.NewViewExEx(new(pagination), ViewPagination, templatePagination, args).(*pagination)
}

func PaginationItem(args ...any) *paginationitem {
	return mvc.NewViewExEx(new(paginationitem), ViewPaginationItem, templatePaginationItem, args).(*paginationitem)
}

func newPaginationFromElement(element Element) mvc.View {
	if element.TagName() != "NAV" {
		return nil
	}
	return mvc.NewViewWithElement(new(pagination), element)
}

func newPaginationItemFromElement(element Element) mvc.View {
	if element.TagName() != "LI" {
		return nil
	}
	return mvc.NewViewWithElement(new(paginationitem), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (pagination *pagination) SetView(view mvc.View) {
	pagination.View = view
}

func (paginationitem *paginationitem) SetView(view mvc.View) {
	paginationitem.View = view
}

func (pagination *pagination) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg := arg.(type) {
		case string:
			args[i] = PaginationItem(arg)
		case dom.Element:
			args[i] = PaginationItem(arg)
		case *paginationitem:
			// No-op
		case mvc.View:
			args[i] = PaginationItem(arg)
		default:
			panic(ErrInternalAppError.Withf("Content[pagination] unexpected argument '%T'", arg))
		}
	}
	return pagination.View.Content(args...)
}
