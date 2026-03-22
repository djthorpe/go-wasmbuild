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

var codePanel struct {
	panel mvc.VisibleState
	title mvc.View
	code  mvc.View
}

var codeExamples = map[string]string{}

// SetCodePanel wires the shared example panel used by story headers.
func SetCodePanel(panel mvc.VisibleState, title, code mvc.View) {
	codePanel.panel = panel
	codePanel.title = title
	codePanel.code = code
}

// RegisterCodeExample stores a code sample for a story title.
func RegisterCodeExample(title, code string) {
	if strings.TrimSpace(title) == "" {
		return
	}
	codeExamples[title] = strings.TrimSpace(code)
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
	children := []any{carbon.Head(2, title)}
	if link := codeLink(title); link != nil {
		children = append(children, link)
	}
	return mvc.HTML("DIV",
		append([]any{mvc.WithStyle("display:flex;align-items:center;gap:0.5rem")}, children...)...,
	)
}

func codeLink(title string) mvc.View {
	if codePanel.panel == nil || codePanel.title == nil || codePanel.code == nil {
		return nil
	}
	link := carbon.Link(
		"javascript:void(0)",
		carbon.With(carbon.LinkInline, carbon.SizeSmall),
		carbon.Icon(carbon.IconLaunch, carbon.With(carbon.IconSize16)),
	).SetLabel("View code for " + title)
	link.AddEventListener(carbon.EventClick, func(dom.Event) {
		openCodeExample(title)
	})
	return link
}

func openCodeExample(title string) {
	if codePanel.panel == nil || codePanel.title == nil || codePanel.code == nil {
		return
	}
	codePanel.title.Content(title)
	codePanel.code.Content(codeExample(title))
	codePanel.panel.SetVisible(true)
}

func codeExample(title string) string {
	if code, ok := codeExamples[title]; ok && strings.TrimSpace(code) != "" {
		return code
	}
	return fmt.Sprintf("// Code example for %q has not been added yet.\n// Register one with storybook.RegisterCodeExample(...).", title)
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
