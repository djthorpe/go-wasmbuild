package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
	model := NewModel("testdata.json", "Name", "Position", "Salary", "Location")
	offcanvas := NewOffcanvas()
	toast := NewToast()
	table := NewTable(offcanvas, toast, model)
	nav := NewNav(table)

	// Create the bootstrap app
	bs.New().Insert(
		nav.View(),
		bs.Container(
			bs.WithBreakpoint(bs.BreakpointLarge),
			bs.WithMargin(bs.TOP|bs.BOTTOM, 5),
		).Insert(
			table.View(),
		),
		offcanvas.View(),
		toast.View(),
	)

	// Load the model data
	model.Load()

	// Run the application
	select {}
}

/*

	/*
		// Helper function to update color ticks
		updateColorTicks := func(selectedID string) {
			for _, item := range colorItems {
				if comp := item.(Component); comp != nil {
					itemID := comp.ID()
					if itemID == "" {
						continue
					}

					elem := comp.Element()

					// Find if this is the selected item
					isSelected := itemID == selectedID

					// Check if tick already exists (look for span with tick)
					var tickSpan Element
					for child := elem.FirstChild(); child != nil; child = child.NextSibling() {
						if span, ok := child.(Element); ok && span.TagName() == "SPAN" {
							if span.TextContent() == "✓" {
								tickSpan = span
								break
							}
						}
					}

					if isSelected && tickSpan == nil {
						// Add tick
						newTick := dom.GetWindow().Document().CreateElement("SPAN")
						newTick.SetAttribute("style", "margin-left: auto; padding-left: 1rem;")
						newTick.AppendChild(dom.GetWindow().Document().CreateTextNode("✓"))
						elem.AppendChild(newTick)
						elem.SetAttribute("style", "display: flex; align-items: center;")
					} else if !isSelected && tickSpan != nil {
						// Remove tick
						elem.RemoveChild(tickSpan)
						elem.SetAttribute("style", "")
					} else if isSelected && tickSpan != nil {
						// Already has tick, ensure flex styling
						elem.SetAttribute("style", "display: flex; align-items: center;")
					}
				}
			}
		}

		// Initialize default color tick
		updateColorTicks("default")

		// Create navbar with brand header and color dropdown
		navbar := bs.NavBar(
			bs.WithClass("navbar-expand-lg"),
			bs.WithTheme(bs.DARK),
			bs.WithClass("bg-primary"),
		).Header(
			bs.Icon("people-fill"),
			" Employee Data Example",
		).Append(bs.NavDropdown("Color", colorItems...).AddEventListener("click", func(target Node) {
			if component := target.Component(); component != nil {
				colorID := component.ID()
				if colorID == "" {
					return
				}
				switch colorID {
				case "default":
					table.Apply(bs.WithColor(bs.TRANSPARENT))
					updateColorTicks(colorID)
				default:
					table.Apply(bs.WithColor(bs.Color(colorID)))
					updateColorTicks(colorID)
				}
			}
		}))

		// Add click handler to each color item
		/*
			for _, item := range colorItems {
				item.AddEventListener("click", func(target Node) {
					if component := target.Component(); component != nil {
						colorID := component.ID()
						if colorID == "" {
							return
						}
						switch colorID {
						case "default":
							table.Apply(bs.WithColor(bs.TRANSPARENT))
							updateColorTicks(colorID)
						default:
							table.Apply(bs.WithColor(bs.Color(colorID)))
							updateColorTicks(colorID)
						}
					}
				})
			}

		// Create styling dropdown with organized sections
		stylingItems := []Component{

		}

		// Helper function to update tick marks
		updateTicks := func(selectedID string, groupIDs []string) {
			for _, item := range stylingItems {
				if comp := item.(Component); comp != nil {
					itemID := comp.ID()
					if itemID == "" {
						continue
					}

					// Check if this item is in the current group
					isInGroup := false
					for _, gid := range groupIDs {
						if itemID == gid {
							isInGroup = true
							break
						}
					}

					if !isInGroup {
						continue
					}

					// Get the element and its text content
					elem := comp.Element()

					// Clear existing content and rebuild with or without tick
					for child := elem.FirstChild(); child != nil; {
						next := child.NextSibling()
						elem.RemoveChild(child)
						child = next
					}

					// Get the label text based on ID
					labelText := ""
					switch itemID {
					case "border-default", "size-default":
						labelText = "Default"
					case "border-bordered":
						labelText = "Bordered"
					case "border-borderless":
						labelText = "Borderless"
					case "size-small":
						labelText = "Small"
					case "divider-off", "hover-off":
						labelText = "Off"
					case "divider-on", "hover-on":
						labelText = "On"
					case "stripe-none":
						labelText = "None"
					case "stripe-row":
						labelText = "Row"
					case "stripe-column":
						labelText = "Column"
					}

					if itemID == selectedID {
						// Add text with tick on the right using flexbox
						elem.AppendChild(dom.GetWindow().Document().CreateTextNode(labelText))
						tickSpan := dom.GetWindow().Document().CreateElement("SPAN")
						tickSpan.SetAttribute("style", "margin-left: auto; padding-left: 1rem;")
						tickSpan.AppendChild(dom.GetWindow().Document().CreateTextNode("✓"))
						elem.AppendChild(tickSpan)
						elem.SetAttribute("style", "display: flex; align-items: center;")
					} else {
						elem.AppendChild(dom.GetWindow().Document().CreateTextNode(labelText))
						elem.SetAttribute("style", "")
					}
				}
			}
		}

		// Initialize default selections with ticks
		updateTicks("border-default", []string{"border-default", "border-bordered", "border-borderless"})
		updateTicks("size-default", []string{"size-default", "size-small"})
		updateTicks("divider-off", []string{"divider-off", "divider-on"})
		updateTicks("stripe-none", []string{"stripe-none", "stripe-row", "stripe-column"})
		updateTicks("hover-on", []string{"hover-off", "hover-on"})

		// Add click handler to each styling item
		for _, item := range stylingItems {
			item.AddEventListener("click", func(target Node) {
				if component := target.Component(); component != nil {
					itemID := component.ID()
					if itemID == "" {
						return
					}

					switch itemID {
					// Border options
					case "border-default":
						table.Element().ClassList().Remove("table-bordered", "table-borderless")
						updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})
					case "border-bordered":
						table.Element().ClassList().Remove("table-borderless")
						table.Element().ClassList().Add("table-bordered")
						updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})
					case "border-borderless":
						table.Element().ClassList().Remove("table-bordered")
						table.Element().ClassList().Add("table-borderless")
						updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})

					// Size options
					case "size-default":
						table.Element().ClassList().Remove("table-sm")
						updateTicks(itemID, []string{"size-default", "size-small"})
					case "size-small":
						table.Element().ClassList().Add("table-sm")
						updateTicks(itemID, []string{"size-default", "size-small"})

					// Group Divider options
					case "divider-off":
						table.Element().ClassList().Remove("table-group-divider")
						updateTicks(itemID, []string{"divider-off", "divider-on"})
					case "divider-on":
						table.Element().ClassList().Add("table-group-divider")
						updateTicks(itemID, []string{"divider-off", "divider-on"})

					// Stripe options
					case "stripe-none":
						table.Element().ClassList().Remove("table-striped", "table-striped-columns")
						updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})
					case "stripe-row":
						table.Element().ClassList().Remove("table-striped-columns")
						table.Element().ClassList().Add("table-striped")
						updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})
					case "stripe-column":
						table.Element().ClassList().Remove("table-striped")
						table.Element().ClassList().Add("table-striped-columns")
						updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})

					// Hover options
					case "hover-off":
						table.Element().ClassList().Remove("table-hover")
						updateTicks(itemID, []string{"hover-off", "hover-on"})
					case "hover-on":
						table.Element().ClassList().Add("table-hover")
						updateTicks(itemID, []string{"hover-off", "hover-on"})
					}
				}
			})
		}

		navbar.Append(bs.NavDropdown("Styling", stylingItems...))
*/
