package storybook

import (
	"fmt"
	"time"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// DefaultThemes keeps the example stories aligned on the same Carbon theme set.
var DefaultThemes = []carbon.Attr{carbon.ThemeWhite, carbon.ThemeG10, carbon.ThemeG90, carbon.ThemeG100}

// Story builds a reusable component demo frame for the example app.
func Story(title, description string, preview dom.Element, observed mvc.View, controls ...dom.Element) dom.Element {
	children := make([]any, 0, 5)
	children = append(children,
		carbon.Head(2, title),
		carbon.Para(description),
		preview,
	)
	if len(controls) > 0 {
		children = append(children, controlsPanel(controls...))
	}
	if observed != nil {
		children = append(children, eventIndicators(observed))
	}
	return mvc.HTML("DIV", append([]any{mvc.WithClass("component-story")}, children...)...)
}

// AttrDropdown builds a Carbon dropdown for a set of Attr options.
func AttrDropdown(label string, selected carbon.Attr, options []carbon.Attr, onChange func(carbon.Attr)) dom.Element {
	items := make([]any, 0, len(options)+1)
	items = append(items, carbon.DropdownTitleText(label))
	for _, option := range options {
		item := carbon.DropdownItem(mvc.WithAttr("value", string(option)), string(option))
		if option == selected {
			item.SetSelected(true)
		}
		items = append(items, item)
	}

	dd := carbon.Dropdown(append([]any{
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
	}, items...)...)
	dd.SetValue(string(selected))
	dd.AddEventListener(carbon.EventSelected, func(dom.Event) {
		onChange(carbon.Attr(dd.Value()))
	})
	return dd.Root()
}

// CheckboxGroup builds a Carbon checkbox group containing one checkbox.
func CheckboxGroup(legend, label string, initiallyChecked bool, onChange func(bool)) dom.Element {
	chk := carbon.Checkbox(label)
	if initiallyChecked {
		chk.SetState(carbon.CheckboxStateTrue)
	} else {
		chk.SetState(carbon.CheckboxStateFalse)
	}
	chk.AddEventListener(carbon.EventCheckboxChanged, func(dom.Event) {
		onChange(chk.State() == carbon.CheckboxStateTrue)
	})

	group := carbon.CheckboxGroup(mvc.WithAttr("legend-text", legend))
	group.AddCheckbox(chk)
	return group.Root()
}

// IconDropdown builds a Carbon dropdown for the bundled icon names.
func IconDropdown(label string, selected carbon.IconName, options []carbon.IconName, onChange func(carbon.IconName)) dom.Element {
	items := make([]any, 0, len(options)+1)
	items = append(items, carbon.DropdownTitleText(label))
	for _, option := range options {
		item := carbon.DropdownItem(mvc.WithAttr("value", string(option)), string(option))
		if option == selected {
			item.SetSelected(true)
		}
		items = append(items, item)
	}

	dd := carbon.Dropdown(append([]any{
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
	}, items...)...)
	dd.SetValue(string(selected))
	dd.AddEventListener(carbon.EventSelected, func(dom.Event) {
		onChange(carbon.IconName(dd.Value()))
	})
	return dd.Root()
}

// IconSizeDropdown builds a Carbon dropdown for icon sizes.
func IconSizeDropdown(label string, selected carbon.IconSize, options []carbon.IconSize, onChange func(carbon.IconSize)) dom.Element {
	items := make([]any, 0, len(options)+1)
	items = append(items, carbon.DropdownTitleText(label))
	for _, option := range options {
		value := fmt.Sprintf("%d", option)
		item := carbon.DropdownItem(mvc.WithAttr("value", value), value)
		if option == selected {
			item.SetSelected(true)
		}
		items = append(items, item)
	}

	dd := carbon.Dropdown(append([]any{
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
	}, items...)...)
	dd.SetValue(fmt.Sprintf("%d", selected))
	dd.AddEventListener(carbon.EventSelected, func(dom.Event) {
		switch dd.Value() {
		case "20":
			onChange(carbon.IconSize20)
		case "24":
			onChange(carbon.IconSize24)
		case "32":
			onChange(carbon.IconSize32)
		default:
			onChange(carbon.IconSize16)
		}
	})
	return dd.Root()
}

func controlsPanel(controls ...dom.Element) dom.Element {
	cols := make([]any, 0, len(controls))
	for _, control := range controls {
		cols = append(cols, carbon.Col4(mvc.WithClass("controls-panel__cell"), control))
	}
	return mvc.HTML("DIV", mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite), "controls-panel"),
		carbon.Grid(append([]any{mvc.WithClass("controls-panel__grid")}, cols...)...),
	)
}

func eventIndicators(v mvc.View) dom.Element {
	evts := mvc.RegisteredEvents(v.Name())
	tiles := make([]any, len(evts))
	const (
		rowStyle     = "display:flex;align-items:stretch;gap:0.5rem;width:100%;margin:1rem 0"
		wrapperStyle = "flex:1;min-width:0"
	)
	for i, evt := range evts {
		val := carbon.Para()
		tile := carbon.TileDecorator(carbon.Head(4, carbon.GoName(evt)), val)
		tile.SetFill(true)
		tile.SetHeight("9rem")
		tile.SetBackground("var(--cds-layer-02,#e0e0e0)")
		localTile := tile
		localVal := val
		v.AddEventListener(evt, func(e dom.Event) {
			text := "—"
			if target, ok := e.Target().(dom.Element); ok {
				if value := target.GetAttribute("value"); value != "" {
					text = value
				} else if href := target.GetAttribute("href"); href != "" {
					text = href
				} else if content := target.TextContent(); content != "" {
					text = content
				}
			}
			localVal.Root().SetInnerHTML(text)
			localTile.SetActive(true)
			js.SetTimeout(600*time.Millisecond, func() {
				localTile.SetActive(false)
			})
		})
		tiles[i] = mvc.HTML("DIV", mvc.WithAttr("style", wrapperStyle), tile)
	}
	return mvc.HTML("DIV", append([]any{mvc.WithAttr("style", rowStyle)}, tiles...)...)
}
