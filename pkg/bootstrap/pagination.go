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
	BootstrapView
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
	p := new(pagination)
	p.BootstrapView.View = mvc.NewViewExEx(p, ViewPagination, templatePagination, args)
	return p
}

func newPaginationFromElement(element Element) mvc.View {
	if element.TagName() != "NAV" {
		return nil
	}
	p := new(pagination)
	p.BootstrapView.View = mvc.NewViewWithElement(p, element)
	return p
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (pagination *pagination) Self() mvc.View {
	return pagination
}

func (pagination *pagination) Content(args ...any) *pagination {
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
	pagination.ReplaceSlot("body", wrapChildren(args...))
	return pagination
}
