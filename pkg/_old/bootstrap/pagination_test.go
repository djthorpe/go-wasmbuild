package bootstrap_test

import (
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS - Pagination

func TestPagination_Basic(t *testing.T) {
	pagination := bs.Pagination()
	assert.NotNil(t, pagination)
	assert.NotNil(t, pagination.Element())
}

func TestPagination_TagName(t *testing.T) {
	pagination := bs.Pagination()
	element := pagination.Element()
	assert.Equal(t, "NAV", element.TagName())
}

func TestPagination_DefaultAriaLabel(t *testing.T) {
	pagination := bs.Pagination()

	element := pagination.Element()
	ariaLabel := element.GetAttribute("aria-label")

	// Should not have a default aria-label
	assert.Equal(t, "", ariaLabel)
}

func TestPagination_CustomAriaLabel(t *testing.T) {
	pagination := bs.Pagination(bs.WithAriaLabel("Search results pages"))
	element := pagination.Element()
	assert.Equal(t, "Search results pages", element.GetAttribute("aria-label"))
}

func TestPagination_HasULChild(t *testing.T) {
	pagination := bs.Pagination()
	element := pagination.Element()
	children := element.Children()
	assert.Equal(t, 1, len(children))
	assert.Equal(t, "UL", children[0].TagName())
}

func TestPagination_ULHasPaginationClass(t *testing.T) {
	pagination := bs.Pagination()
	element := pagination.Element()
	children := element.Children()
	assert.True(t, children[0].ClassList().Contains("pagination"))
}

func TestPagination_AppendText(t *testing.T) {
	pagination := bs.Pagination().Append(bs.PaginationItem("1"))
	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 1, len(liElements))
	assert.Equal(t, "LI", liElements[0].TagName())
	assert.True(t, liElements[0].ClassList().Contains("page-item"))
	assert.Equal(t, "1", liElements[0].TextContent())
}

func TestPagination_AppendMultiple(t *testing.T) {
	pagination := bs.Pagination().
		Append(bs.PaginationItem("Previous")).
		Append(bs.PaginationItem("1")).
		Append(bs.PaginationItem("2")).
		Append(bs.PaginationItem("Next"))

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 4, len(liElements))
	assert.Equal(t, "Previous", liElements[0].TextContent())
	assert.Equal(t, "1", liElements[1].TextContent())
	assert.Equal(t, "2", liElements[2].TextContent())
	assert.Equal(t, "Next", liElements[3].TextContent())
}

func TestPagination_AppendComponent(t *testing.T) {
	item := bs.PaginationItem("Page 1")
	pagination := bs.Pagination().Append(item)

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	assert.Equal(t, 1, len(liElements))
	assert.Equal(t, "LI", liElements[0].TagName())
	assert.True(t, liElements[0].ClassList().Contains("page-item"))

	// Check that the link is inside the li
	liChildren := liElements[0].Children()
	assert.Equal(t, 1, len(liChildren))
	assert.Equal(t, "A", liChildren[0].TagName())
	assert.Equal(t, "Page 1", liChildren[0].TextContent())
}

func TestPagination_Insert(t *testing.T) {
	pagination := bs.Pagination().
		Append(bs.PaginationItem("2")).
		Append(bs.PaginationItem("3"))

	// Insert at the beginning
	pagination.Insert(bs.PaginationItem("1"))

	element := pagination.Element()
	ul := element.Children()[0] // Get the <ul>
	items := ul.Children()      // Get <li> elements

	// Should have 3 items
	assert.Equal(t, 3, len(items))
	// First item should be "1"
	assert.Contains(t, items[0].TextContent(), "1")
	// Second item should be "2"
	assert.Contains(t, items[1].TextContent(), "2")
	// Third item should be "3"
	assert.Contains(t, items[2].TextContent(), "3")
}

func TestPagination_InsertMultiple(t *testing.T) {
	pagination := bs.Pagination().
		Append(bs.PaginationItem("3"))

	// Insert multiple at the beginning
	pagination.Insert(
		bs.PaginationItem("1"),
		bs.PaginationItem("2"),
	)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.Equal(t, 3, len(items))
	assert.Contains(t, items[0].TextContent(), "1")
	assert.Contains(t, items[1].TextContent(), "2")
	assert.Contains(t, items[2].TextContent(), "3")
}

func TestPagination_AppendInvalidType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when appending non-PaginationItem")
		}
	}()

	pagination := bs.Pagination()
	pagination.Append("string") // Should panic
}

func TestPagination_InsertInvalidType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when inserting non-PaginationItem")
		}
	}()

	pagination := bs.Pagination()
	pagination.Insert(bs.Link("#")) // Should panic
}

func TestPagination_IsComponent(t *testing.T) {
	pagination := bs.Pagination()
	var _ dom.Component = pagination
	assert.NotNil(t, pagination)
}

func TestPagination_ComponentMethods(t *testing.T) {
	pagination := bs.Pagination()

	// Test Element() method
	element := pagination.Element()
	assert.NotNil(t, element)
	assert.Equal(t, "NAV", element.TagName())

	// Test Append() returns Component
	result := pagination.Append(bs.PaginationItem("Test"))
	assert.NotNil(t, result)
	var _ dom.Component = result

	// Test Insert() returns Component
	result = pagination.Insert(bs.PaginationItem("First"))
	assert.NotNil(t, result)
	var _ dom.Component = result
}

func TestPagination_AllPageItemsHaveClass(t *testing.T) {
	pagination := bs.Pagination().
		Append(bs.PaginationItem("1")).
		Append(bs.PaginationItem("2")).
		Append(bs.PaginationItem("3"))

	element := pagination.Element()
	ul := element.Children()[0]
	liElements := ul.Children()

	for i, li := range liElements {
		assert.True(t, li.ClassList().Contains("page-item"),
			"Element %d should have page-item class", i)
	}
}

func TestPagination_OuterHTML(t *testing.T) {
	pagination := bs.Pagination(bs.WithAriaLabel("Test navigation")).
		Append(bs.PaginationItem("1")).
		Append(bs.PaginationItem("2"))

	element := pagination.Element()
	html := element.OuterHTML()

	// HTML output uses lowercase tags (matching browser behavior)
	assert.Contains(t, html, "<nav")
	assert.Contains(t, html, "aria-label=\"Test navigation\"")
	assert.Contains(t, html, "<ul")
	assert.Contains(t, html, "class=\"pagination\"")
	assert.Contains(t, html, "<li")
	assert.Contains(t, html, "class=\"page-item\"")
}

func TestPagination_WithOptions(t *testing.T) {
	pagination := bs.Pagination(
		bs.WithClass("custom-class"),
		bs.WithAttribute("data-test", "value"))

	element := pagination.Element()
	assert.True(t, element.ClassList().Contains("custom-class"))
	assert.Equal(t, "value", element.GetAttribute("data-test"))
}

func TestPagination_EmptyAriaLabel(t *testing.T) {
	pagination := bs.Pagination()
	element := pagination.Element()
	ariaLabel := element.GetAttribute("aria-label")

	// Without WithAriaLabel, should have no aria-label
	assert.Equal(t, "", ariaLabel)
}

func TestPagination_ChainedAppend(t *testing.T) {
	pagination := bs.Pagination()

	result := pagination.
		Append(bs.PaginationItem("1")).
		Append(bs.PaginationItem("2")).
		Append(bs.PaginationItem("3"))

	// Verify chaining returns a Component
	var _ dom.Component = result

	// Verify all items were added
	element := pagination.Element()
	ul := element.Children()[0]
	assert.Equal(t, 3, len(ul.Children()))
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - PaginationItem

func TestPaginationItem_Basic(t *testing.T) {
	item := bs.PaginationItem("1")
	assert.NotNil(t, item)
	assert.NotNil(t, item.Element())
}

func TestPaginationItem_TagName(t *testing.T) {
	item := bs.PaginationItem("1")
	element := item.Element()
	assert.Equal(t, "LI", element.TagName())
}

func TestPaginationItem_HasPageItemClass(t *testing.T) {
	item := bs.PaginationItem("1")
	element := item.Element()
	assert.True(t, element.ClassList().Contains("page-item"))
}

func TestPaginationItem_HasLinkChild(t *testing.T) {
	item := bs.PaginationItem("1")
	element := item.Element()
	children := element.Children()

	assert.Equal(t, 1, len(children))
	assert.Equal(t, "A", children[0].TagName())
}

func TestPaginationItem_LinkHasPageLinkClass(t *testing.T) {
	item := bs.PaginationItem("1")
	link := item.Link()

	assert.NotNil(t, link)
	assert.True(t, link.ClassList().Contains("page-link"))
	assert.Equal(t, "#", link.GetAttribute("href"))
}

func TestPaginationItem_TextContent(t *testing.T) {
	item := bs.PaginationItem("Page 1")
	link := item.Link()

	assert.Equal(t, "Page 1", link.TextContent())
}

func TestPaginationItem_EmptyText(t *testing.T) {
	item := bs.PaginationItem("")
	link := item.Link()

	assert.Equal(t, "", link.TextContent())
}

func TestPaginationItem_WithOptions(t *testing.T) {
	item := bs.PaginationItem("1", bs.WithClass("active"))
	element := item.Element()

	assert.True(t, element.ClassList().Contains("page-item"))
	assert.True(t, element.ClassList().Contains("active"))
}

func TestPaginationItem_IsComponent(t *testing.T) {
	item := bs.PaginationItem("1")
	var _ dom.Component = item
}

func TestPaginationItem_InPagination(t *testing.T) {
	pagination := bs.Pagination()
	item1 := bs.PaginationItem("1")
	item2 := bs.PaginationItem("2")

	pagination.Append(item1)
	pagination.Append(item2)

	element := pagination.Element()
	ul := element.Children()[0]

	assert.Equal(t, 2, len(ul.Children()))
}

func TestPaginationItem_OuterHTML(t *testing.T) {
	item := bs.PaginationItem("Next")
	element := item.Element()
	html := element.OuterHTML()

	assert.Contains(t, html, "<li")
	assert.Contains(t, html, "class=\"page-item\"")
	assert.Contains(t, html, "<a")
	assert.Contains(t, html, "class=\"page-link\"")
	assert.Contains(t, html, "href=\"#\"")
	assert.Contains(t, html, "Next")
}

///////////////////////////////////////////////////////////////////////////////
// TESTS - Pagination Active/Disabled

func TestPagination_Active(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2"))
	pagination.Append(bs.PaginationItem("3"))

	// Set item at index 1 as active
	pagination.Active(1)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.False(t, items[0].ClassList().Contains("active"))
	assert.True(t, items[1].ClassList().Contains("active"))
	assert.False(t, items[2].ClassList().Contains("active"))
}

func TestPagination_Active_Multiple(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2"))
	pagination.Append(bs.PaginationItem("3"))
	pagination.Append(bs.PaginationItem("4"))

	// Set multiple items as active
	pagination.Active(0, 2)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.True(t, items[0].ClassList().Contains("active"))
	assert.False(t, items[1].ClassList().Contains("active"))
	assert.True(t, items[2].ClassList().Contains("active"))
	assert.False(t, items[3].ClassList().Contains("active"))
}

func TestPagination_Active_GetActive(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2", bs.WithClass("active")))
	pagination.Append(bs.PaginationItem("3"))

	// Get currently active items
	activeIndices := pagination.Active()

	assert.Equal(t, 1, len(activeIndices))
	assert.Equal(t, 1, activeIndices[0])
}

func TestPagination_Active_ClearsPrevious(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2"))
	pagination.Append(bs.PaginationItem("3"))

	// Set item 1 as active
	pagination.Active(1)

	// Now set item 2 as active (should clear item 1)
	pagination.Active(2)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.False(t, items[0].ClassList().Contains("active"))
	assert.False(t, items[1].ClassList().Contains("active"))
	assert.True(t, items[2].ClassList().Contains("active"))
}

func TestPagination_Disabled(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("Previous"))
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("Next"))

	// Disable first item (Previous)
	pagination.Disabled(0)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.True(t, items[0].ClassList().Contains("disabled"))
	assert.False(t, items[1].ClassList().Contains("disabled"))
	assert.False(t, items[2].ClassList().Contains("disabled"))
}

func TestPagination_Disabled_Multiple(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("Previous"))
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2"))
	pagination.Append(bs.PaginationItem("Next"))

	// Disable first and last items
	pagination.Disabled(0, 3)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.True(t, items[0].ClassList().Contains("disabled"))
	assert.False(t, items[1].ClassList().Contains("disabled"))
	assert.False(t, items[2].ClassList().Contains("disabled"))
	assert.True(t, items[3].ClassList().Contains("disabled"))
}

func TestPagination_Disabled_GetDisabled(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("Previous", bs.WithClass("disabled")))
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("Next", bs.WithClass("disabled")))

	// Get currently disabled items
	disabledIndices := pagination.Disabled()

	assert.Equal(t, 2, len(disabledIndices))
	assert.Equal(t, 0, disabledIndices[0])
	assert.Equal(t, 2, disabledIndices[1])
}

func TestPagination_ActiveAndDisabled(t *testing.T) {
	pagination := bs.Pagination()
	pagination.Append(bs.PaginationItem("Previous"))
	pagination.Append(bs.PaginationItem("1"))
	pagination.Append(bs.PaginationItem("2"))
	pagination.Append(bs.PaginationItem("3"))
	pagination.Append(bs.PaginationItem("Next"))

	// Set item 2 as active and first/last as disabled
	pagination.Active(2)
	pagination.Disabled(0, 4)

	element := pagination.Element()
	ul := element.Children()[0]
	items := ul.Children()

	assert.True(t, items[0].ClassList().Contains("disabled"))
	assert.False(t, items[1].ClassList().Contains("active"))
	assert.True(t, items[2].ClassList().Contains("active"))
	assert.False(t, items[3].ClassList().Contains("disabled"))
	assert.True(t, items[4].ClassList().Contains("disabled"))
}
