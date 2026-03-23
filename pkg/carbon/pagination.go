package carbon

import (
	"strconv"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination struct{ base }

var _ mvc.View = (*pagination)(nil)
var _ mvc.EnabledState = (*pagination)(nil)
var _ mvc.PaginationState = (*pagination)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewPagination, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(pagination), element, setView)
	}, EventPaginationChanged, EventPaginationPageSize)
}

// Pagination returns a <cds-pagination> web component.
func Pagination(args ...any) *pagination {
	return mvc.NewView(new(pagination), ViewPagination, "cds-pagination", setView, args...).(*pagination)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (p *pagination) Enabled() bool {
	return !p.Root().HasAttribute("disabled")
}

func (p *pagination) SetEnabled(enabled bool) mvc.View {
	if enabled {
		p.Root().RemoveAttribute("disabled")
	} else {
		p.Root().SetAttribute("disabled", "")
	}
	return p
}

func (p *pagination) SetPage(page uint) *pagination {
	limit := p.Limit()
	if limit == 0 {
		p.SetOffset(0)
	} else if page <= 1 {
		p.SetOffset(0)
	} else {
		p.SetOffset((page - 1) * limit)
	}
	setTagIntProperty(p.Root(), "page", int(mvc.Page(p)))
	return p
}

func (p *pagination) Offset() uint {
	return uint(tagIntProperty(p.Root(), "start"))
}

func (p *pagination) SetOffset(offset uint) mvc.View {
	setTagIntProperty(p.Root(), "start", int(offset))
	setTagIntProperty(p.Root(), "page", int(mvc.Page(p)))
	return p
}

func (p *pagination) Limit() uint {
	return uint(tagIntProperty(p.Root(), "page-size"))
}

func (p *pagination) SetLimit(limit uint) mvc.View {
	setTagIntProperty(p.Root(), "page-size", int(limit))
	setTagIntProperty(p.Root(), "page", int(mvc.Page(p)))
	return p
}

func (p *pagination) Count() uint {
	return uint(tagIntProperty(p.Root(), "total-items"))
}

func (p *pagination) SetCount(count uint) mvc.View {
	setTagIntProperty(p.Root(), "total-items", int(count))
	setTagIntProperty(p.Root(), "page", int(mvc.Page(p)))
	return p
}

func (p *pagination) PagesUnknown() bool {
	return tagBoolProperty(p.Root(), "pages-unknown")
}

func (p *pagination) SetPagesUnknown(enabled bool) *pagination {
	setTagBoolProperty(p.Root(), "pages-unknown", enabled)
	return p
}

func (p *pagination) SetBackwardText(text string) *pagination {
	return p.setTextAttr("backward-text", text)
}

func (p *pagination) SetForwardText(text string) *pagination {
	return p.setTextAttr("forward-text", text)
}

func (p *pagination) SetItemsPerPageText(text string) *pagination {
	return p.setTextAttr("items-per-page-text", text)
}

func (p *pagination) SetPageSizeLabelText(text string) *pagination {
	return p.setTextAttr("page-size-label-text", text)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (p *pagination) setTextAttr(key, value string) *pagination {
	if value == "" {
		p.Root().RemoveAttribute(key)
	} else {
		p.Root().SetAttribute(key, value)
	}
	return p
}

func tagIntProperty(element dom.Element, key string) int {
	if element == nil {
		return 0
	}
	if node, ok := element.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		if value := node.Get(key); !value.IsUndefined() && !value.IsNull() {
			if result, err := strconv.Atoi(value.String()); err == nil {
				return result
			}
		}
	}
	if value := element.GetAttribute(key); value != "" {
		if result, err := strconv.Atoi(value); err == nil {
			return result
		}
	}
	return 0
}

func setTagIntProperty(element dom.Element, key string, value int) {
	if element == nil {
		return
	}
	if node, ok := element.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Set(key, value)
	}
	if value <= 0 {
		element.RemoveAttribute(key)
		return
	}
	text := strconv.Itoa(value)
	element.SetAttribute(key, text)
}
