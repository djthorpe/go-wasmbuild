package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type pagination struct {
	component
}

type paginationItem struct {
	component
	link Element
}

var _ Component = (*pagination)(nil)
var _ Component = (*paginationItem)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Pagination creates a new Bootstrap pagination component
func Pagination(opt ...Opt) *pagination {
	doc := dom.GetWindow().Document()

	// Create nav element
	nav := doc.CreateElement("NAV")

	// Create ul element with pagination class
	list := doc.CreateElement("UL")
	list.SetAttribute("class", "pagination")
	nav.AppendChild(list)

	// Apply options to nav element
	c := newComponent(PaginationComponent, nav)
	if err := c.applyTo(nav, opt...); err != nil {
		panic(err)
	}

	// Set the body to the list element so Empty() works correctly
	c.body = list

	return &pagination{
		component: *c,
	}
}

// PaginationItem creates a new Bootstrap pagination item component
// which is a <li class="page-item"><a class="page-link"> structure
func PaginationItem(text string, opt ...Opt) *paginationItem {
	doc := dom.GetWindow().Document()

	// Create li element
	li := doc.CreateElement("LI")

	// Create link element with page-link class
	link := doc.CreateElement("A")
	link.SetAttribute("class", "page-link")
	link.SetAttribute("href", "#")

	// Add text content
	if text != "" {
		link.AppendChild(doc.CreateTextNode(text))
	}

	// Append link to li
	li.AppendChild(link)

	// Create component and apply options with page-item class
	c := newComponent(PaginationItemComponent, li)
	if err := c.applyTo(li, append(opt, WithClass("page-item"))...); err != nil {
		panic(err)
	}

	return &paginationItem{
		component: *c,
		link:      link,
	}
}

///////////////////////////////////////////////////////////////////////////////
// METHODS - paginationItem

// Link returns the inner <a> element of the pagination item
func (p *paginationItem) Link() Element {
	return p.link
}

///////////////////////////////////////////////////////////////////////////////
// METHODS - pagination

// Append adds PaginationItem components to the pagination list
// Only accepts *paginationItem - use PaginationItem() to create items
// Panics if a non-PaginationItem type is passed
func (p *pagination) Append(children ...any) Component {
	for _, child := range children {
		// Only accept PaginationItem components
		item, ok := child.(*paginationItem)
		if !ok {
			panic("Pagination.Append only accepts *paginationItem - use PaginationItem() to create items")
		}
		p.body.AppendChild(item.Element())
	}

	return p
}

// Insert adds PaginationItem components at the beginning of the pagination list
// Only accepts *paginationItem - use PaginationItem() to create items
// Panics if a non-PaginationItem type is passed
func (p *pagination) Insert(children ...any) Component {
	firstChild := p.body.FirstChild()

	for _, child := range children {
		// Only accept PaginationItem components
		item, ok := child.(*paginationItem)
		if !ok {
			panic("Pagination.Insert only accepts *paginationItem - use PaginationItem() to create items")
		}

		if firstChild != nil {
			p.body.InsertBefore(item.Element(), firstChild)
		} else {
			p.body.AppendChild(item.Element())
		}
	}

	return p
}

// Active marks specified pagination items as active by adding the "active" class
// If no indices are provided, returns the currently active item indices
// Indices are zero-based and refer to the position in the list
func (p *pagination) Active(indices ...int) []int {
	if len(indices) > 0 {
		// Set mode: mark specified items as active
		activeIndices := make(map[int]bool)
		for _, idx := range indices {
			if idx >= 0 {
				activeIndices[idx] = true
			}
		}

		// Iterate through all items and set/remove active class
		currentIndex := 0
		for child := p.body.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "LI" {
				classList := elem.ClassList()
				if activeIndices[currentIndex] {
					classList.Add("active")
				} else {
					classList.Remove("active")
				}
				currentIndex++
			}
		}

		return indices
	}

	// Get mode: return currently active item indices
	var activeItems []int
	currentIndex := 0
	for child := p.body.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "LI" {
			if elem.ClassList().Contains("active") {
				activeItems = append(activeItems, currentIndex)
			}
			currentIndex++
		}
	}

	return activeItems
}

// Disabled marks specified pagination items as disabled by adding the "disabled" class
// If no indices are provided, returns the currently disabled item indices
// Indices are zero-based and refer to the position in the list
func (p *pagination) Disabled(indices ...int) []int {
	if len(indices) > 0 {
		// Set mode: mark specified items as disabled
		disabledIndices := make(map[int]bool)
		for _, idx := range indices {
			if idx >= 0 {
				disabledIndices[idx] = true
			}
		}

		// Iterate through all items and set/remove disabled class
		currentIndex := 0
		for child := p.body.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "LI" {
				classList := elem.ClassList()
				if disabledIndices[currentIndex] {
					classList.Add("disabled")
				} else {
					classList.Remove("disabled")
				}
				currentIndex++
			}
		}

		return indices
	}

	// Get mode: return currently disabled item indices
	var disabledItems []int
	currentIndex := 0
	for child := p.body.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "LI" {
			if elem.ClassList().Contains("disabled") {
				disabledItems = append(disabledItems, currentIndex)
			}
			currentIndex++
		}
	}

	return disabledItems
}
