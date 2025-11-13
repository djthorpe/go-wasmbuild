package main

import (
	"fmt"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Table struct {
	view  Component
	table Component
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	TableLimit = 15
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTable(offcanvas *Offcanvas, toast *Toast, model *Model) *Table {
	controller := new(Table)
	pagination := NewPagination(model, toast, TableLimit)

	// Create the view
	table := bs.Table(
		bs.WithHover(), bs.WithAnimation(), bs.WithCursor(bs.CursorPointer),
	).Header(columns(model.Columns())...)

	// Set controller options
	controller.view = bs.Container().Append(
		table, pagination.View(),
	)
	controller.table = table

	// Listen for events from the model
	model.AddEventListener(EventLoad, func(e *jsutil.Event) {
		// Empty the table
		table.Empty()

		// Populate table with employee data
		rows := model.Get(pagination.Offset(), pagination.Limit())
		if len(rows) == 0 {
			return
		}
		for _, employee := range rows {
			row := bs.Row(employee.Name, employee.Position, formatSalary(employee.Salary), employee.Location)
			table.Append(row)
		}

		// Update pagination
		pagination.Update(model)
	})

	// Listen to events on the table rows
	table.AddEventListener("click", func(node Node) {
		component := node.Component()
		if component == nil {
			return
		}
		if component.Name() != string(bs.TableRowComponent) {
			return
		}

		// Show the offcanvas with employee details
		toast.view.Empty().Insert(fmt.Sprint(component))
		toast.Show()
	})

	// Return the controller
	return controller
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Table) View() Component {
	return this.view
}

func (this *Table) Table() Component {
	return this.table
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func columns(cols []string) []any {
	result := make([]any, len(cols))
	for i, col := range cols {
		result[i] = col
	}
	return result
}

// formatSalary converts a numeric salary to a formatted string with $ and commas
// e.g., 95000 -> "$95,000"
func formatSalary(salary float64) string {
	// Convert to integer
	salaryInt := int(salary)

	// Convert to string
	s := fmt.Sprintf("%d", salaryInt)

	// Add commas
	n := len(s)
	if n <= 3 {
		return "$" + s
	}

	// Build result with commas from right to left
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(s[i])
	}

	return "$" + result
}

/*
	// Helper function to show employee details in offcanvas
	showEmployeeDetails := func(employee *Employee) {
		if employee == nil {
			return
		}

		// Find the offcanvas body element and clear it
		offcanvasElem := offcanvas.Element()

		// Find the body div (it's the second child after the header)
		var bodyElem Element
		for child := offcanvasElem.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok {
				if elem.ClassList().Contains("offcanvas-body") {
					bodyElem = elem
					break
				}
			}
		}

		if bodyElem != nil {
			// Clear existing content
			for child := bodyElem.FirstChild(); child != nil; {
				next := child.NextSibling()
				bodyElem.RemoveChild(child)
				child = next
			}

			// Create a card with employee details
			card := bs.Card()

			// Card header with employee name
			card.Header(employee.Name) // Create a form with employee details for the card body
			form := bs.Form()

			// Helper to create a form field with label and input
			createField := func(label, name, value string) Component {
				field := bs.Container(bs.WithClass("mb-3")).Append(
					bs.Label(label),
					bs.Input(name, bs.WithValue(value)),
				)
				return field
			}

			// Create form fields
			posDiv := createField("Position", "position", employee.Position)
			// Salary field - use NumberInput for numeric input
			// Use the actual numeric salary value from the model
			salaryValue := fmt.Sprintf("%.0f", employee.Salary)
			fmt.Printf("Creating salary input with value: %s (from %.0f)\n", salaryValue, employee.Salary)
			salaryInput := bs.NumberInput("salary")
			salaryInput.Value(salaryValue)
			salDiv := bs.Container(bs.WithClass("mb-3")).Append(
				bs.Label("Salary"),
				salaryInput,
			)
			locDiv := createField("Location", "location", employee.Location) // Append all fields to form
			form.Append(posDiv, salDiv, locDiv)                              // Add submit event listener to form
			form.AddEventListener("submit", func(node Node) {
				// Get the form element as js.Value and create jsutil.Form wrapper
				formElem := form.Element()
				formJS := formElem.(interface{ JSValue() js.Value }).JSValue()
				jsForm := jsutil.NewForm(formJS)

				// Get form data as a map
				formData := jsForm.FormData().Entries()

				fmt.Printf("Form submitted with data:\n")
				fmt.Printf("  Name: %s\n", employee.Name)
				fmt.Printf("  Position: %s\n", formData["position"])
				fmt.Printf("  Salary: %s\n", formData["salary"])
				fmt.Printf("  Location: %s\n", formData["location"])

				// Hide the offcanvas
				offcanvasJSValue := js.Global().Get("document").Call("getElementById", "employeeDetails")
				bsOffcanvas := js.Global().Get("bootstrap").Get("Offcanvas")
				instance := bsOffcanvas.Call("getOrCreateInstance", offcanvasJSValue)
				instance.Call("hide")
			}) // Add form to card body
			card.Append(form)

			// Add Save button to card footer that triggers form submit
			saveBtn := bs.Button(bs.PRIMARY).Append("Save")
			saveBtn.AddEventListener("click", func(node Node) {
				// Trigger form submit event
				formElem := form.Element()
				// Create and dispatch a submit event
				submitEvent := js.Global().Get("Event").New("submit")
				formElem.(interface {
					Call(string, ...interface{}) js.Value
				}).Call("dispatchEvent", submitEvent)
			})
			card.Footer(saveBtn)

			bodyElem.AppendChild(card.Element())
		}

		// Use Bootstrap's JavaScript API to properly show the offcanvas with backdrop
		// Get the raw JavaScript value of the element
		offcanvasJSValue := js.Global().Get("document").Call("getElementById", "employeeDetails")
		bsOffcanvas := js.Global().Get("bootstrap").Get("Offcanvas")
		instance := bsOffcanvas.Call("getOrCreateInstance", offcanvasJSValue)
		instance.Call("show")
	}

	// Action toolbar
	toolbar := bs.ButtonToolbar(bs.WithMargin(bs.END, 3), bs.WithMargin(bs.BOTTOM, 3))
	addBtn := bs.Button(bs.SUCCESS, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("plus-circle"),
		" Add",
	)
	removeBtn := bs.Button(bs.DANGER, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("dash-circle"),
		" Remove",
	)
	updateBtn := bs.Button(bs.WARNING, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("arrow-repeat"),
		" Update",
	)
	clearBtn := bs.Button(bs.SECONDARY, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("x-circle"),
		" Clear",
	)

	// Helper function to update button states based on active rows
	updateButtonStates := func() {
		activeRows := table.Active()
		hasActiveRows := len(activeRows) > 0

		if hasActiveRows {
			// Enable all buttons - we have 4 buttons (0=Add, 1=Remove, 2=Update, 3=Clear)
			// When there are active rows, keep only Add enabled would be: toolbar.Disabled(1,2,3)
			// But we want ALL enabled, so disable NONE - pass empty/out of range indices
			// Actually, let's be explicit: we'll manually iterate or use a workaround
			// For now, let's disable button at index -1 (doesn't exist), forcing all to be enabled
			toolbar.Disabled(-1)
		} else {
			toolbar.Disabled(1, 2, 3) // Disable Remove, Update, Clear (indices 1, 2, 3)
		}
	}

	// Add mutation observer to table to automatically update button states
	// whenever row active states change
	observer := dom.GetWindow().NewMutationObserver(func() {
		updateButtonStates()
	})
	observer.Observe(tableElem, map[string]interface{}{
		"attributes":      true,
		"attributeFilter": []interface{}{"class"},
		"subtree":         true,
	})

	// Helper function to add click handler to a row
	addRowClickHandlerWithUpdate := func(row Component) {
		row.AddEventListener("click", func(node Node) {
			// node.Component() will walk up the DOM tree to find the row component
			// even if the click target is a TD element
			if rowComp := node.Component(); rowComp != nil {
				rowElem := rowComp.Element()

				// Find the row index by iterating through tbody children
				parent := rowElem.ParentNode()
				if parent == nil {
					return
				}

				rowIndex := -1
				currentIndex := 0
				for child := parent.FirstChild(); child != nil; child = child.NextSibling() {
					if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
						if elem == rowElem {
							rowIndex = currentIndex
							break
						}
						currentIndex++
					}
				}

				if rowIndex >= 0 {
					// Get the employee from the model using the row index
					employee := model.GetByIndex(rowIndex)
					if employee != nil {
						// Debug: print employee data
						fmt.Printf("Row clicked: index=%d, name=%s, salary=%.0f\n", rowIndex, employee.Name, employee.Salary)
						// Show employee details in offcanvas
						showEmployeeDetails(employee)
					} else {
						fmt.Printf("Row clicked: index=%d, employee is nil\n", rowIndex)
					}
				}

				// Toggle the active class directly on the element
				classList := rowElem.ClassList()
				if classList.Contains("table-active") {
					classList.Remove("table-active")
				} else {
					classList.Add("table-active")
				}
				// updateButtonStates() will be called automatically by mutation observer
			}
		})
	}

	addBtn.AddEventListener("click", func(node Node) {
		if model.Count() > 0 {
			// Get a random employee from the data
			randomEmployee := model.GetRandom()

			newRow := bs.Row(
				randomEmployee.Name,
				randomEmployee.Position,
				formatSalary(randomEmployee.Salary),
				randomEmployee.Location,
			)
			addRowClickHandlerWithUpdate(newRow)

			// Get active rows and insert before the first active row, or at the beginning
			activeRows := table.Active()
			if len(activeRows) > 0 {
				// Insert before the first active row
				table.InsertBefore(activeRows[0], newRow)
			} else {
				// Insert at the beginning (before index 0)
				table.InsertBefore(0, newRow)
			}
		}
	})
	removeBtn.AddEventListener("click", func(node Node) {
		activeRows := table.Active()
		if len(activeRows) > 0 {
			// Delete all active rows, starting from the highest index to avoid index shifting issues
			for i := len(activeRows) - 1; i >= 0; i-- {
				table.Delete(activeRows[i])
			}
			// updateButtonStates() will be called automatically by mutation observer
		}
	})
	updateBtn.AddEventListener("click", func(node Node) {
		activeRows := table.Active()
		if len(activeRows) > 0 && model.Count() > 0 {
			// Update all active rows with random employee data
			for _, rowIndex := range activeRows {
				// Get a random employee from the data
				randomEmployee := model.GetRandom()

				newRow := bs.Row(
					randomEmployee.Name,
					randomEmployee.Position,
					formatSalary(randomEmployee.Salary),
					randomEmployee.Location,
				).WithActive()
				addRowClickHandlerWithUpdate(newRow)
				table.Replace(rowIndex, newRow)
			}
		}
	})
	clearBtn.AddEventListener("click", func(node Node) {
		table.Active(-1)
		// updateButtonStates() will be called automatically by mutation observer
	})

	toolbar.Append(addBtn, removeBtn, updateBtn, clearBtn)

	// Initialize button states after buttons are appended
	updateButtonStates()

	container.Append(toolbar)

	// Create a loading message
	loadingDiv := bs.Container().Insert(
		"Loading employee data...",
	)
	container.Append(loadingDiv)

	// Function to calculate and update the total salary
	updateTotalSalary := func() {
		tbody := tableElem.FirstChild().(Element)
		totalSalary := 0.0

		// Iterate through all rows in tbody
		for child := tbody.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
				// Get the third cell (salary column, index 2)
				cells := elem.Children()
				if len(cells) >= 3 {
					salaryCell := cells[2]
					salaryText := salaryCell.TextContent()

					// Parse salary (remove $ and commas)
					cleanSalary := ""
					for _, ch := range salaryText {
						if ch != '$' && ch != ',' {
							cleanSalary += string(ch)
						}
					}

					// Convert to float
					salaryVal := js.Global().Get("parseFloat").Invoke(cleanSalary).Float()
					if !js.Global().Get("isNaN").Invoke(salaryVal).Bool() {
						totalSalary += salaryVal
					}
				}
			}
		}

		// Update footer with new total
		totalFormatted := js.Global().Get("Intl").Get("NumberFormat").New("en-US", map[string]interface{}{
			"style":    "currency",
			"currency": "USD",
		}).Call("format", totalSalary).String()

		// Find and update the footer
		tfoot := tableElem.LastChild()
		if tfootElem, ok := tfoot.(Element); ok && tfootElem.TagName() == "TFOOT" {
			tfootRow := tfootElem.FirstChild()
			if tfootRowElem, ok := tfootRow.(Element); ok {
				cells := tfootRowElem.Children()
				if len(cells) >= 4 {
					lastCell := cells[3]
					for child := lastCell.FirstChild(); child != nil; {
						next := child.NextSibling()
						lastCell.RemoveChild(child)
						child = next
					}
					lastCell.AppendChild(dom.GetWindow().Document().CreateTextNode(totalFormatted))
				}
			}
		}
	}

	// Add mutation observer to watch for table changes
	tableElem.AddEventListener("DOMSubtreeModified", func(node Node) {
		updateTotalSalary()
	})

	// Create pagination component
	pagination := bs.Pagination("Employee table navigation")

	// Declare updatePagination first so it can be called recursively
	var updatePagination func()

	// Function to refresh the table based on current model offset/limit
	refreshTable := func() {
		// Clear existing table rows (keep header)
		tbody := tableElem.FirstChild()
		if tbodyElem, ok := tbody.(Element); ok && tbodyElem.TagName() == "TBODY" {
			for child := tbodyElem.FirstChild(); child != nil; {
				next := child.NextSibling()
				tbodyElem.RemoveChild(child)
				child = next
			}
		}

		// Populate table with current page of employee data
		for _, employee := range model.GetAll() {
			row := bs.Row(employee.Name, employee.Position, formatSalary(employee.Salary), employee.Location)
			addRowClickHandlerWithUpdate(row)
			table.Append(row)
		}

		updateTotalSalary()
	}

	// Function to update pagination based on model state
	updatePagination = func() {
		// Clear existing pagination items by removing and recreating
		paginationElem := pagination.Element()

		// Find and remove the UL element
		for child := paginationElem.FirstChild(); child != nil; {
			next := child.NextSibling()
			paginationElem.RemoveChild(child)
			child = next
		}

		// Recreate UL element with centered alignment
		doc := dom.GetWindow().Document()
		list := doc.CreateElement("UL")
		list.SetAttribute("class", "pagination justify-content-center")
		paginationElem.AppendChild(list)

		totalPages := (model.Count() + model.Limit() - 1) / model.Limit()
		if totalPages == 0 {
			totalPages = 1
		}
		currentPage := (model.Offset() / model.Limit()) + 1

		// Helper to create a page link element
		createPageLink := func(text string, pageNum int, disabled bool, active bool) Element {
			doc := dom.GetWindow().Document()
			link := doc.CreateElement("A")
			link.SetAttribute("class", "page-link")
			link.SetAttribute("href", "javascript:void(0)")
			if disabled {
				link.SetAttribute("tabindex", "-1")
				link.SetAttribute("aria-disabled", "true")
			}
			if active {
				link.SetAttribute("aria-current", "page")
			}
			link.AppendChild(doc.CreateTextNode(text))

			if !disabled {
				link.AddEventListener("click", func(node Node) {
					model.SetOffsetLimit((pageNum-1)*model.Limit(), model.Limit())
					refreshTable()
					updatePagination()
				})
			}

			// Create LI wrapper
			li := doc.CreateElement("LI")
			li.SetAttribute("class", "page-item")
			if disabled {
				li.ClassList().Add("disabled")
			}
			if active {
				li.ClassList().Add("active")
			}
			li.AppendChild(link)

			return li
		}

		// Add Previous button
		list.AppendChild(createPageLink("Previous", currentPage-1, currentPage == 1, false))

		// Add page numbers (show max 5 pages)
		startPage := currentPage - 2
		if startPage < 1 {
			startPage = 1
		}
		endPage := startPage + 4
		if endPage > totalPages {
			endPage = totalPages
			startPage = endPage - 4
			if startPage < 1 {
				startPage = 1
			}
		}

		for page := startPage; page <= endPage; page++ {
			pageNum := page
			pageText := ""
			// Convert int to string
			pageText = js.Global().Get("String").Invoke(pageNum).String()
			list.AppendChild(createPageLink(pageText, pageNum, false, pageNum == currentPage))
		}

		// Add Next button
		list.AppendChild(createPageLink("Next", currentPage+1, currentPage == totalPages, false))
	}

	// Create an interval to check for data
	model.WaitForData(func() {
		container.Element().RemoveChild(loadingDiv.Element())

		// Populate table with employee data
		for _, employee := range model.GetAll() {
			row := bs.Row(employee.Name, employee.Position, formatSalary(employee.Salary), employee.Location)
			addRowClickHandlerWithUpdate(row)
			table.Append(row)
		}

		table.Footer("", "", "Total:", "$0.00")
		updateTotalSalary()
		container.Append(table)

		// Add pagination below the table
		updatePagination()
		container.Append(pagination)
	})
*/
