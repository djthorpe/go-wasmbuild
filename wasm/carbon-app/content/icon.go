package headings

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

// IconEntry is a single icon in the registry.
type IconEntry struct {
	ID   carbon.IconName
	Name string
}

func IconView() []any {
	// selectIcon is set by the playground story; the browser calls it on tile click.
	var selectIcon func(IconEntry)

	playground := iconPlaygroundStory(&selectIcon)
	browser := iconBrowserStory(func(e IconEntry) {
		if selectIcon != nil {
			selectIcon(e)
		}
	})

	return []any{
		mvc.HTML("DIV", mvc.WithStyle("padding:1.5rem 2rem"), carbon.Head(1, "Icons")),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			playground,
		),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeG10),
			browser,
		),
	}
}

// iconPlaygroundStory shows a single icon with live theme / size / colour controls.
// It stores a selection callback into *onSelect so the browser can drive it.
func iconPlaygroundStory(onSelect *func(IconEntry)) dom.Element {
	allIcons := Icons()
	currentEntry := allIcons[0]
	currentSize := carbon.IconSize32
	currentTheme := carbon.ThemeG10
	currentColor := "inherit"

	iconEl := carbon.Icon(currentEntry.ID, carbon.With(currentSize))
	nameLabel := carbon.Strong(currentEntry.Name)
	idLabel := carbon.Code(string(currentEntry.ID), mvc.WithStyle("font-size:0.75rem"))

	iconWrapper := mvc.HTML("DIV",
		mvc.WithStyle("display:flex;flex-direction:column;align-items:center;justify-content:center;gap:1rem;padding:4rem"),
		nameLabel,
		iconEl,
		idLabel,
	)

	canvas := carbon.Section(
		carbon.With(currentTheme),
		mvc.WithStyle("display:flex;align-items:center;justify-content:center;min-height:16rem"),
		iconWrapper,
	)

	refresh := func() {
		// Set size attribute before calling SetIcon so i.Size() picks it up.
		iconEl.Root().SetAttribute("size", string(currentSize))
		iconEl.SetIcon(currentEntry.ID)
		nameLabel.Root().SetInnerHTML(currentEntry.Name)
		idLabel.Root().SetInnerHTML(string(currentEntry.ID))

		// Update colour on the wrapper
		style := "display:flex;flex-direction:column;align-items:center;justify-content:center;gap:1rem;padding:4rem"
		if currentColor != "" {
			style += ";color:" + currentColor
		}
		iconWrapper.SetAttribute("style", style)

		// Swap theme class on canvas
		for _, t := range storybook.DefaultThemes {
			canvas.Root().ClassList().Remove(carbon.ClassForTheme(t))
		}
		canvas.Root().ClassList().Add(carbon.ClassForTheme(currentTheme))
	}

	// Expose the selection callback so the browser can drive this story.
	*onSelect = func(e IconEntry) {
		currentEntry = e
		refresh()
		// Scroll the playground canvas into view so the update is visible.
		if node, ok := canvas.Root().JSValue().(js.Value); ok {
			node.Call("scrollIntoView", map[string]any{"behavior": "smooth", "block": "nearest"})
		}
	}

	return storybook.Story(
		"Canvas",
		"Click any icon in the browser below to select it. Use the controls to adjust theme, size and colour.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(a carbon.Attr) {
			currentTheme = a
			refresh()
		}),
		storybook.IconSizeDropdown("Size", currentSize, []carbon.IconSize{
			carbon.IconSize16, carbon.IconSize20, carbon.IconSize24, carbon.IconSize32,
		}, func(s carbon.IconSize) {
			currentSize = s
			refresh()
		}),
		colorDropdown("Colour", currentColor, [][2]string{
			{"Default", "inherit"},
			{"Blue", "#0f62fe"},
			{"Teal", "#009d9a"},
			{"Green", "#24a148"},
			{"Red", "#da1e28"},
			{"Purple", "#8a3ffc"},
			{"Magenta", "#ee5396"},
			{"Yellow", "#f1c21b"},
			{"White", "#ffffff"},
			{"Black", "#161616"},
		}, func(c string) {
			currentColor = c
			refresh()
		}),
	)
}

// colorDropdown builds a dropdown of named colour options, returning the CSS colour value.
func colorDropdown(label, selected string, opts [][2]string, onChange func(string)) mvc.View {
	items := make([]any, 0, len(opts))
	for _, opt := range opts {
		item := carbon.DropdownItem(opt[0]).SetValue(opt[1])
		if opt[1] == selected {
			item.SetActive(true)
		}
		items = append(items, item)
	}
	onChange(selected)
	return carbon.Dropdown("", mvc.WithStyle("width:100%"), items).
		SetLabel(label).
		SetValue(selected).
		AddEventListener(carbon.EventSelected, func(e dom.Event) {
			if v := mvc.ViewFromEventTarget(e, carbon.ViewDropdown); v != nil {
				onChange(v.Root().Value())
			}
		})
}

func iconBrowserStory(onSelect func(IconEntry)) dom.Element {
	allNames := Icons()

	currentQuery := ""

	countLabel := carbon.Compact(fmt.Sprintf("Showing %d icons", len(allNames)))

	grid := mvc.HTML("DIV",
		mvc.WithStyle("display:flex;flex-wrap:wrap;justify-content:center;gap:0.5rem;margin-top:1rem"),
	)

	refresh := func() {
		terms := strings.Fields(strings.ToLower(currentQuery))
		names := Icons(terms...)
		countLabel.Root().SetInnerHTML(fmt.Sprintf("Showing %d of %d icons", len(names), len(allNames)))
		grid.SetInnerHTML("")
		for _, entry := range names {
			grid.AppendChild(iconTile(entry, carbon.IconSize32, onSelect))
		}
	}

	// Initial render
	for _, entry := range allNames {
		grid.AppendChild(iconTile(entry, carbon.IconSize32, onSelect))
	}

	searchInput := mvc.HTML("INPUT",
		mvc.WithAttr("type", "search"),
		mvc.WithAttr("placeholder", "Filter icons…"),
		mvc.WithAttr("style", "width:100%;padding:0.5rem 1rem;font-size:0.875rem;"+
			"border:1px solid var(--cds-border-strong-01,#8d8d8d);"+
			"background:var(--cds-field-01,#f4f4f4);"+
			"color:var(--cds-text-primary,#161616);outline:none"),
	)
	searchInput.AddEventListener("input", func(e dom.Event) {
		if target, ok := e.Target().(dom.Element); ok {
			currentQuery = strings.TrimSpace(target.Value())
			refresh()
		}
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:flex;flex-direction:column;gap:0.75rem"),
		searchInput,
		countLabel,
		grid,
	)

	return storybook.Story(
		"Browser",
		fmt.Sprintf("All %d Carbon icons are bundled into this example app for browsing. In production, applications would typically import only the icons they use.", len(allNames)),
		canvas,
		nil,
	)
}

func iconTile(entry IconEntry, size carbon.IconSize, onClick func(IconEntry)) dom.Element {
	tile := carbon.Tile(
		mvc.WithStyle("width:8rem;cursor:pointer"),
		mvc.HTML("DIV",
			mvc.WithStyle("display:flex;flex-direction:column;align-items:center;gap:0.5rem;text-align:center;height:100%"),
			carbon.Strong(entry.Name),
			carbon.Icon(entry.ID, carbon.With(size)),
			carbon.Code(string(entry.ID), mvc.WithStyle("font-size:0.625rem")),
		),
	)
	if onClick != nil {
		tile.Root().AddEventListener("click", func(dom.Event) {
			onClick(entry)
		})
	}
	return tile.Root()
}
