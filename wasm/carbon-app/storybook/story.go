package storybook

import (
	"fmt"
	"strings"
	"time"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// DefaultThemes keeps the example stories aligned on the same Carbon theme set.
var DefaultThemes = []carbon.Attr{carbon.ThemeWhite, carbon.ThemeG10, carbon.ThemeG90, carbon.ThemeG100}

var docsPanel struct {
	panel mvc.VisibleState
	body  mvc.View
}

var componentDocs = map[string]string{}

// SetDocsPanel wires the shared documentation panel used by page headers.
func SetDocsPanel(panel mvc.VisibleState, body mvc.View) {
	docsPanel.panel = panel
	docsPanel.body = body
}

// RegisterComponentDoc associates a story title with an embedded component doc.
func RegisterComponentDoc(title, filename string) {
	if strings.TrimSpace(title) == "" {
		return
	}
	componentDocs[title] = strings.TrimSpace(filename)
}

// Story builds a reusable component demo frame for the example app.
func Story(title, description string, preview mvc.View, observed mvc.View, controls ...mvc.View) dom.Element {
	// Add the title and description if provided, then the preview content
	children := make([]any, 0, 5)
	if title != "" {
		children = append(children, storyHeader(title))
	}
	if description != "" {
		children = append(children, carbon.Para(description))
	}
	children = append(children, preview)

	// Add any controls, then the event indicators if an observed view is provided.
	if len(controls) > 0 {
		children = append(children, controlsPanel(controls...))
	}
	if observed != nil {
		children = append(children, eventIndicators(observed))
	}

	// Return the story
	return mvc.HTML("DIV", append([]any{mvc.WithClass("component-story")}, children...)...)
}

func storyHeader(title string) dom.Element {
	return mvc.HTML("DIV",
		mvc.WithStyle("display:flex;align-items:center;gap:0.5rem"),
		carbon.Head(2, title),
	)
}

// PageHeader builds a top-level component page heading with a single docs link.
func PageHeader(title, filename string) dom.Element {
	RegisterComponentDoc(title, filename)
	row := []any{carbon.Head(1, title)}
	if link := docsLink(title); link != nil {
		row = append(row, link)
	}
	children := []any{mvc.HTML("DIV",
		append([]any{mvc.WithStyle("display:flex;align-items:center;gap:0.75rem")}, row...)...,
	)}
	if description, err := componentDocDescription(filename); err == nil && description != "" {
		children = append(children, carbon.Lead(description))
	}
	return mvc.HTML("DIV",
		append([]any{mvc.WithStyle("padding:1.5rem 2rem")}, children...)...,
	)
}

func docsLink(title string) mvc.View {
	if docsPanel.panel == nil || docsPanel.body == nil {
		return nil
	}
	if strings.TrimSpace(componentDocs[title]) == "" {
		return nil
	}
	link := carbon.Link(
		"javascript:void(0)",
		carbon.With(carbon.LinkInline, carbon.SizeSmall),
		carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize16)),
	).SetLabel("View docs for " + title)
	link.AddEventListener(carbon.EventClick, func(dom.Event) {
		openComponentDoc(title)
	})
	return link
}

func openComponentDoc(title string) {
	if docsPanel.panel == nil || docsPanel.body == nil {
		return
	}
	filename := componentDocs[title]
	if strings.TrimSpace(filename) == "" {
		return
	}
	docsPanel.body.Content(ComponentDoc(filename))
	docsPanel.panel.SetVisible(true)
}

// Dropdown builds a Carbon dropdown for a set of Attr options.
func Dropdown(label string, selected carbon.Attr, options []carbon.Attr, onChange func(carbon.Attr)) mvc.View {
	onChange(selected)
	items := make([]any, 0, len(options))
	for _, option := range options {
		items = append(items, carbon.DropdownItem(string(option)).SetValue(string(option)))
	}
	return carbon.Dropdown("", items).
		SetLabel(label).
		SetValue(string(selected)).
		AddEventListener(carbon.EventSelected, func(e dom.Event) {
			if v := mvc.ViewFromEventTarget(e, carbon.ViewDropdown); v != nil {
				onChange(carbon.Attr(v.Root().Value()))
			}
		})
}

// CheckboxGroup builds a Carbon checkbox group containing one checkbox.
func CheckboxGroup(legend, label string, selected bool, onChange func(bool)) mvc.View {
	chk := carbon.Checkbox(label)
	chk.SetActive(selected)
	chk.AddEventListener(carbon.EventChange, func(e dom.Event) {
		if v := mvc.ViewFromEventTarget(e, carbon.ViewCheckbox); v != nil {
			if a, ok := v.(interface{ Active() bool }); ok {
				onChange(a.Active())
			}
		}
	})
	return carbon.CheckboxGroup("", chk).SetLabel(legend)
}

// IconDropdown builds a Carbon dropdown for the bundled icon names.
func IconDropdown(label string, selected carbon.IconName, options []carbon.IconName, onChange func(carbon.IconName)) mvc.View {
	items := make([]any, 0, len(options))
	for _, option := range options {
		item := carbon.DropdownItem(string(option)).SetValue(string(option))
		if option == selected {
			item.SetActive(true)
		}
		items = append(items, item)
	}

	return carbon.Dropdown("",
		mvc.WithAttr("style", "width:100%"),
		items,
	).SetLabel(label).SetValue(string(selected)).AddEventListener(carbon.EventSelected, func(e dom.Event) {
		if v := mvc.ViewFromEventTarget(e, carbon.ViewDropdown); v != nil {
			onChange(carbon.IconName(v.Root().Value()))
		}
	})
}

// IconSizeDropdown builds a Carbon dropdown for icon sizes.
func IconSizeDropdown(label string, selected carbon.IconSize, options []carbon.IconSize, onChange func(carbon.IconSize)) mvc.View {
	items := make([]any, 0, len(options))
	for _, option := range options {
		value := string(option)
		item := carbon.DropdownItem(value).SetValue(value)
		if option == selected {
			item.SetActive(true)
		}
		items = append(items, item)
	}

	return carbon.Dropdown("",
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
		items,
	).SetLabel(label).SetValue(string(selected)).AddEventListener(carbon.EventSelected, func(e dom.Event) {
		if v := mvc.ViewFromEventTarget(e, carbon.ViewDropdown); v != nil {
			switch v.Root().Value() {
			case "20":
				onChange(carbon.IconSize20)
			case "24":
				onChange(carbon.IconSize24)
			case "32":
				onChange(carbon.IconSize32)
			default:
				onChange(carbon.IconSize16)
			}
		}
	})
}

func controlsPanel(controls ...mvc.View) dom.Element {
	cols := make([]any, 0, len(controls))
	for i, control := range controls {
		// Each cell needs its own z-index so that the first dropdown's open
		// menu paints above the cells that follow it in the DOM.
		zIndex := len(controls) - i
		cols = append(cols, carbon.Col4(
			mvc.WithClass("controls-panel__cell"),
			mvc.WithStyle(fmt.Sprintf("position:relative;z-index:%d;overflow:visible", zIndex)),
			control,
		))
	}
	return mvc.HTML("DIV", mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite), "controls-panel"),
		mvc.WithStyle("position:relative;z-index:100;overflow:visible"),
		carbon.Grid(append([]any{mvc.WithClass("controls-panel__grid"), mvc.WithStyle("overflow:visible")}, cols...)...),
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
		tile := carbon.TileDecorator(
			carbon.WithFill(),
			carbon.WithHeight("9rem"),
			carbon.WithBackground("var(--cds-layer-02,#e0e0e0)"),
			carbon.Head(4, carbon.GoName(evt)),
			val,
		)
		localTile := tile
		localVal := val
		v.AddEventListener(evt, func(e dom.Event) {
			text := "—"
			if target, ok := e.Target().(dom.Element); ok {
				if value := target.GetAttribute("value"); value != "" {
					text = value
				} else if title := target.GetAttribute("title"); title != "" {
					text = title
				} else if href := target.GetAttribute("href"); href != "" {
					text = href
				} else if content := target.TextContent(); content != "" {
					text = content
				}
			}
			localVal.Root().SetInnerHTML(text)
			baseStyle := localTile.Root().GetAttribute("style")
			flashStyle := strings.Trim(baseStyle, "; ")
			if flashStyle != "" {
				flashStyle += ";"
			}
			flashStyle += "filter:contrast(1.12) brightness(0.94) saturate(0.92);transform:translateY(-1px);box-shadow:0 0 0 1px var(--cds-border-strong-01,#8d8d8d)"
			localTile.Root().SetAttribute("style", flashStyle)
			js.SetTimeout(600*time.Millisecond, func() {
				localTile.Root().SetAttribute("style", baseStyle)
			})
		})
		tiles[i] = mvc.HTML("DIV", mvc.WithAttr("style", wrapperStyle), tile)
	}
	return mvc.HTML("DIV", append([]any{mvc.WithAttr("style", rowStyle)}, tiles...)...)
}
