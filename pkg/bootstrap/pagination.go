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

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ViewPagination = "mvc-bs-pagination"
)

const (
	templatePagination = `
		<nav><ul class="pagination" data-slot></ul></nav>	
	`
	templatePaginationItem = `
		<li class="page-item"><a class="page-link" href="#"></a></li>
	`
)

func init() {
	mvc.RegisterView(ViewPagination, newPaginationFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Pagination(args ...any) *pagination {
	return mvc.NewViewExEx(new(pagination), ViewPagination, templatePagination, args).(*pagination)
}

func newPaginationFromElement(element Element) mvc.View {
	if element.TagName() != "NAV" {
		return nil
	}
	return mvc.NewViewWithElement(new(pagination), element)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (pagination *pagination) SetView(view mvc.View) {
	pagination.View = view
}

func (pagination *pagination) Content(args ...any) mvc.View {
	for i, arg := range args {
		switch arg := arg.(type) {
		case string:
			args[i] = mvc.HTML(templatePaginationItem, arg)
		case dom.Element:
			args[i] = mvc.HTML(templatePaginationItem, arg)
		case mvc.View:
			args[i] = mvc.HTML(templatePaginationItem, arg)
		default:
			panic(ErrInternalAppError.Withf("Content[pagination] unexpected argument '%T'", arg))
		}
	}
	return pagination.View.Content(args...)
}
