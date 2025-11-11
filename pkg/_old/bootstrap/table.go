package bootstrap

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type table struct {
	component
	caption  Element
	thead    Element
	tbody    Element
	tfoot    Element
	animated bool // Whether to use animations for row operations
}

// Ensure that table implements Component interface
var _ Component = (*table)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Table creates a new bootstrap table (table element with class="table")
// The table automatically creates a TBODY element for content.
// THEAD is created on-demand when Header() is called.
// TFOOT is created on-demand when Footer() is called.
// Use Header() to set column headers, Footer() for footer cells,
// and Append()/Insert() to add rows to the table body.
//
// Example:
//
//	Table().Header("Name", "Age", "City").Append(
//	    Row("John", "30", "NYC"),
//	    Row("Jane", "25", "LA"),
//	)
func Table(opt ...Opt) *table {
	// Create a table element
	root := dom.GetWindow().Document().CreateElement("TABLE")

	c := newComponent(TableComponent, root)

	// Apply options with default "table" class
	if err := c.applyTo(root, append([]Opt{WithClass("table")}, opt...)...); err != nil {
		panic(err)
	}

	// Check for animation marker class and remove it from the DOM
	var animated bool
	if root.ClassList().Contains("table-animated") {
		animated = true
		root.ClassList().Remove("table-animated")
	}

	// Create TBODY element (this is where content is appended by default)
	tbody := dom.GetWindow().Document().CreateElement("TBODY")
	root.AppendChild(tbody)

	// Inject animation CSS if needed
	if animated {
		injectTableAnimationCSS()
	}

	c.body = tbody // TBODY is the default append target

	return &table{
		component: *c,
		caption:   nil, // Created on demand by Caption()
		thead:     nil, // Created on demand by Header()
		tbody:     tbody,
		tfoot:     nil, // Created on demand by Footer()
		animated:  animated,
	}
}

///////////////////////////////////////////////////////////////////////////////
// TABLE ROW TYPE

type tableRow struct {
	component
}

// Ensure that tableRow implements Component interface
var _ Component = (*tableRow)(nil)

///////////////////////////////////////////////////////////////////////////////
// TABLE ROW LIFECYCLE

// Row creates a table row (TR element) with TD cells for each child argument.
// Each child argument becomes a TD cell in the row.
// Accepts string, Component, or Element children.
// Returns a Component that can be appended to a table.
//
// Example:
//
//	table.Append(
//	    Row("John", "30", "NYC"),
//	    Row("Jane", "25", Badge("Active")),
//	)
func Row(children ...any) *tableRow {
	// Create a TR element
	tr := dom.GetWindow().Document().CreateElement("TR")

	// Add each child as a TD cell
	for _, child := range children {
		td := dom.GetWindow().Document().CreateElement("TD")

		if component, ok := child.(Component); ok {
			td.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			td.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			td.AppendChild(node)
		}

		tr.AppendChild(td)
	}

	// Use newComponent and applyTo to set data-component attribute
	c := newComponent(TableRowComponent, tr)

	// Apply to set the data-component attribute
	if err := c.applyTo(tr); err != nil {
		panic(err)
	}

	return &tableRow{component: *c}
}

///////////////////////////////////////////////////////////////////////////////
// TABLE ROW METHODS

// WithActive adds the "table-active" class to the row.
// Returns *tableRow to allow method chaining.
func (r *tableRow) WithActive() *tableRow {
	r.root.ClassList().Add("table-active")
	return r
}

// WithoutActive removes the "table-active" class from the row.
// Returns *tableRow to allow method chaining.
func (r *tableRow) WithoutActive() *tableRow {
	r.root.ClassList().Remove("table-active")
	return r
}

// ToggleActive toggles the "table-active" class on the row.
// Returns *tableRow to allow method chaining.
func (r *tableRow) ToggleActive() *tableRow {
	classList := r.root.ClassList()
	if classList.Contains("table-active") {
		classList.Remove("table-active")
	} else {
		classList.Add("table-active")
	}
	return r
}

///////////////////////////////////////////////////////////////////////////////
// TABLE METHODS

// Caption sets the table caption text.
// Creates a CAPTION element (if it doesn't exist) and sets its text content.
// The caption is inserted as the first child of the table element.
// Returns *table to allow method chaining.
// If called multiple times, replaces the previous caption text.
func (t *table) Caption(text string) *table {
	// Create CAPTION if it doesn't exist
	if t.caption == nil {
		t.caption = dom.GetWindow().Document().CreateElement("CAPTION")
		// Insert CAPTION as the first child of the table
		if t.root.FirstChild() != nil {
			t.root.InsertBefore(t.caption, t.root.FirstChild())
		} else {
			t.root.AppendChild(t.caption)
		}
	}

	// Clear existing caption content
	for _, child := range t.caption.ChildNodes() {
		t.caption.RemoveChild(child)
	}

	// Set caption text
	t.caption.AppendChild(dom.GetWindow().Document().CreateTextNode(text))

	return t
}

// Header sets the table header row (THEAD > TR > TH elements).
// Each child argument becomes a TH cell in the header row.
// Accepts string, Component, or Element children.
// Returns *table to allow method chaining.
// If called multiple times, replaces the previous header.
func (t *table) Header(children ...any) *table {
	// Create THEAD if it doesn't exist
	if t.thead == nil {
		t.thead = dom.GetWindow().Document().CreateElement("THEAD")
		// Insert THEAD before TBODY
		t.root.InsertBefore(t.thead, t.tbody)
	}

	// Clear existing THEAD content
	for _, child := range t.thead.ChildNodes() {
		t.thead.RemoveChild(child)
	}

	// Create a TR element for the header row
	tr := dom.GetWindow().Document().CreateElement("TR")

	// Add each child as a TH cell
	for _, child := range children {
		th := dom.GetWindow().Document().CreateElement("TH")

		if component, ok := child.(Component); ok {
			th.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			th.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			th.AppendChild(node)
		}

		tr.AppendChild(th)
	}

	// Append the header row to THEAD
	t.thead.AppendChild(tr)

	return t
}

// Footer creates a TFOOT element (if it doesn't exist) and sets footer cells.
// Each child argument becomes a TD cell in the footer row.
// Accepts string, Component, or Element children.
// Returns *table to allow method chaining.
// If called multiple times, replaces the previous footer.
func (t *table) Footer(children ...any) *table {
	// Create TFOOT if it doesn't exist
	if t.tfoot == nil {
		t.tfoot = dom.GetWindow().Document().CreateElement("TFOOT")
		t.root.AppendChild(t.tfoot)
	}

	// Clear existing TFOOT content
	for _, child := range t.tfoot.ChildNodes() {
		t.tfoot.RemoveChild(child)
	}

	// Create a TR element for the footer row
	tr := dom.GetWindow().Document().CreateElement("TR")

	// Add each child as a TD cell
	for _, child := range children {
		td := dom.GetWindow().Document().CreateElement("TD")

		if component, ok := child.(Component); ok {
			td.AppendChild(component.Element())
		} else if str, ok := child.(string); ok {
			td.AppendChild(dom.GetWindow().Document().CreateTextNode(str))
		} else if node, ok := child.(Node); ok {
			td.AppendChild(node)
		}

		tr.AppendChild(td)
	}

	// Append the footer row to TFOOT
	t.tfoot.AppendChild(tr)

	return t
}

// Append adds Row components to the TBODY at the end.
// Only accepts Component objects with TR elements (created by Row()).
// Returns Component to satisfy the Component interface.
// Panics if a non-Row component is provided.
func (t *table) Append(children ...any) Component {
	for _, child := range children {
		if component, ok := child.(Component); ok {
			// Validate that it's a TR element
			if component.Element().TagName() != "TR" {
				panic("Table.Append() only accepts Row components (TR elements)")
			}
			t.tbody.AppendChild(component.Element())
		} else {
			panic("Table.Append() only accepts Row components")
		}
	}
	return t
}

// Insert prepends Row components to the TBODY at the beginning.
// Only accepts Component objects with TR elements (created by Row()).
// Returns Component to satisfy the Component interface.
// Panics if a non-Row component is provided.
func (t *table) Insert(children ...any) Component {
	// Get the first child of tbody (if any)
	firstChild := t.tbody.FirstChild()

	for _, child := range children {
		if component, ok := child.(Component); ok {
			// Validate that it's a TR element
			if component.Element().TagName() != "TR" {
				panic("Table.Insert() only accepts Row components (TR elements)")
			}
			if firstChild != nil {
				t.tbody.InsertBefore(component.Element(), firstChild)
			} else {
				t.tbody.AppendChild(component.Element())
			}
		} else {
			panic("Table.Insert() only accepts Row components")
		}
	}
	return t
}

// Count returns the number of rows in the table body (TBODY).
// Does not include header (THEAD) or footer (TFOOT) rows.
func (t *table) Count() int {
	count := 0
	for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
			count++
		}
	}
	return count
}

// InsertBefore inserts a Row component before the row at the specified index.
// Index is 0-based. If index is -1, appends to the end.
// If index is out of bounds (< -1 or >= Count()), does nothing.
// Only accepts Component objects with TR elements (created by Row()).
// Returns *table to allow method chaining.
// Panics if a non-Row component is provided.
// If animation is enabled (WithAnimation), adds fade-in animation class.
func (t *table) InsertBefore(index int, row Component) *table {
	// Validate that it's a TR element
	if row.Element().TagName() != "TR" {
		panic("Table.InsertBefore() only accepts Row components (TR elements)")
	}

	// Add animation class if enabled
	if t.animated {
		row.Element().ClassList().Add("table-row-fade-in")
		// Set up listener to remove animation class after it completes
		row.Element().AddEventListener("animationend", func(n Node) {
			if elem, ok := n.(Element); ok {
				elem.ClassList().Remove("table-row-fade-in")
			}
		})
	}

	// If index is -1, append to end
	if index == -1 {
		t.tbody.AppendChild(row.Element())
		return t
	}

	// If index is out of bounds, do nothing
	if index < -1 || index >= t.Count() {
		return t
	}

	// Find the row at the specified index
	currentIndex := 0
	for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
			if currentIndex == index {
				// Insert before this row
				t.tbody.InsertBefore(row.Element(), child)
				return t
			}
			currentIndex++
		}
	}

	return t
}

// Delete removes the row at the specified index from the table body.
// Index is 0-based. Returns *table to allow method chaining.
// If index is out of bounds, does nothing.
// If animation is enabled (WithAnimation), adds fade-out animation before removal.
func (t *table) Delete(index int) *table {
	if index < 0 {
		return t
	}

	// Find the row at the specified index
	currentIndex := 0
	for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
			if currentIndex == index {
				if t.animated {
					// Add animation class
					elem.ClassList().Remove("table-row-fade-in") // Remove any fade-in first
					elem.ClassList().Add("table-row-fade-out")
					// Remove after animation completes
					elem.AddEventListener("animationend", func(n Node) {
						t.tbody.RemoveChild(n)
					})
				} else {
					// Remove immediately
					t.tbody.RemoveChild(child)
				}
				return t
			}
			currentIndex++
		}
	}

	return t
}

// Replace replaces the row at the specified index with a new Row component.
// Index is 0-based. Only accepts Component objects with TR elements (created by Row()).
// Returns *table to allow method chaining.
// If index is out of bounds, does nothing.
// Panics if a non-Row component is provided.
// If animation is enabled (WithAnimation), animates the replacement.
func (t *table) Replace(index int, row Component) *table {
	// Validate that it's a TR element
	if row.Element().TagName() != "TR" {
		panic("Table.Replace() only accepts Row components (TR elements)")
	}

	if index < 0 {
		return t
	}

	// Find the row at the specified index
	currentIndex := 0
	for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
			if currentIndex == index {
				if t.animated {
					// Add fade-out to old row
					elem.ClassList().Add("table-row-fade-out")
					// Insert new row with fade-in
					row.Element().ClassList().Add("table-row-fade-in")
					t.tbody.InsertBefore(row.Element(), child)
					// Remove old row after its animation completes
					elem.AddEventListener("animationend", func(n Node) {
						t.tbody.RemoveChild(n)
					})
					// Clean up animation class on new row after animation completes
					row.Element().AddEventListener("animationend", func(n Node) {
						if newElem, ok := n.(Element); ok {
							newElem.ClassList().Remove("table-row-fade-in")
						}
					})
				} else {
					// Insert new row before old row, then remove old row immediately
					t.tbody.InsertBefore(row.Element(), child)
					t.tbody.RemoveChild(child)
				}
				return t
			}
			currentIndex++
		}
	}

	return t
}

// Active manages the active state of table rows using the "table-active" class.
// When called with indices: marks those rows as active, clears others, and returns the indices.
// When called without indices: returns the indices of currently active rows.
// Accepts one or more 0-based row indices.
// Indices that are out of bounds are ignored.
// Returns []int containing the indices of active rows.
func (t *table) Active(indices ...int) []int {
	if len(indices) > 0 {
		// Set mode: mark specified rows as active
		activeIndices := make(map[int]bool)
		for _, idx := range indices {
			if idx >= 0 {
				activeIndices[idx] = true
			}
		}

		// Iterate through all rows and set/remove active class
		currentIndex := 0
		for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
				classList := elem.ClassList()
				if activeIndices[currentIndex] {
					classList.Add("table-active")
				} else {
					classList.Remove("table-active")
				}
				currentIndex++
			}
		}

		return indices
	}

	// Get mode: return currently active row indices
	var activeRows []int
	currentIndex := 0
	for child := t.tbody.FirstChild(); child != nil; child = child.NextSibling() {
		if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
			if elem.ClassList().Contains("table-active") {
				activeRows = append(activeRows, currentIndex)
			}
			currentIndex++
		}
	}

	return activeRows
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

var tableAnimationCSSInjected bool

// injectTableAnimationCSS injects the CSS keyframes for table row animations
// This is called once when the first animated table is created
func injectTableAnimationCSS() {
	if tableAnimationCSSInjected {
		return
	}
	tableAnimationCSSInjected = true

	doc := dom.GetWindow().Document()
	style := doc.CreateElement("STYLE")
	cssText := doc.CreateTextNode(`
		/* Table row fade-out animation with scaleY for smooth height collapse */
		tr.table-row-fade-out {
			transform-origin: top;
			animation: tableRowFadeOut 0.3s ease-out forwards;
		}
		
		@keyframes tableRowFadeOut {
			0% {
				opacity: 1;
				transform: scaleY(1);
			}
			100% {
				opacity: 0;
				transform: scaleY(0);
			}
		}
		
		/* Table row fade-in animation with scaleY for smooth height expansion */
		tr.table-row-fade-in {
			transform-origin: top;
			animation: tableRowFadeIn 0.3s ease-in forwards;
		}
		
		@keyframes tableRowFadeIn {
			0% {
				opacity: 0;
				transform: scaleY(0);
			}
			100% {
				opacity: 1;
				transform: scaleY(1);
			}
		}
	`)
	style.AppendChild(cssText)
	doc.Body().AppendChild(style)
}
