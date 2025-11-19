package main

import (
	// Packages
	"fmt"
	"strconv"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func ButtonExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Button Examples"),
		bs.HRule(),
		bs.Heading(3, "Buttons With Color", mvc.WithClass("mt-5")), Example(Example_Buttons_001),
		bs.Heading(3, "Outlined Buttons", mvc.WithClass("mt-5")), Example(Example_Buttons_002),
		bs.Heading(3, "Close Button", mvc.WithClass("mt-5")), Example(Example_Buttons_003),
		bs.Heading(3, "Button with Labels", mvc.WithClass("mt-5")), Example(Example_Buttons_004),
		bs.Heading(3, "Button Sizes", mvc.WithClass("mt-5")), Example(Example_Buttons_005),
		bs.Heading(3, "Button States", mvc.WithClass("mt-5")), Example(Example_Buttons_006),
		bs.Heading(3, "Button Groups", mvc.WithClass("mt-5")), Example(Example_Buttons_007),
		bs.Heading(3, "Button Groups with Color", mvc.WithClass("mt-5")), Example(Example_Buttons_008),
		bs.Heading(3, "Button Groups Sizes", mvc.WithClass("mt-5")), Example(Example_Buttons_009),
		bs.Heading(3, "Vertical Button Group", mvc.WithClass("mt-5")), Example(Example_Buttons_010),
		bs.Heading(3, "Button Toolbar", mvc.WithClass("mt-5")), Example(Example_Buttons_011),
	)
}

func Example_Buttons_001() (mvc.View, string) {
	return bs.Container(
		bs.Button("Default Button", mvc.WithClass("m-1")),
		bs.Button("Secondary Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1")),
		bs.Button("Danger Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1")),
	), sourcecode()
}

func Example_Buttons_002() (mvc.View, string) {
	return bs.Container(
		bs.OutlineButton("Default Button", mvc.WithClass("m-1")),
		bs.OutlineButton("Secondary Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1")),
		bs.OutlineButton("Danger Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1")),
	), sourcecode()
}

func Example_Buttons_003() (mvc.View, string) {
	response := bs.Para("Click the close button")
	return bs.Container(
		bs.CloseButton(mvc.WithClass("m-1")),
		response,
	).AddEventListener("click", func(e Event) {
		if v := mvc.ViewFromEvent(e); v.Name() == bs.ViewButton {
			bs.SetContent(response, "Close button clicked!")
		}
	}), sourcecode()
}

func Example_Buttons_004() (mvc.View, string) {
	response := mvc.Text("50")
	return bs.Container(
		bs.Button("Inbox", mvc.WithClass("m-3")).Label(response),
		bs.RangeInput("unread_emails", mvc.WithClass("m-3")),
	).AddEventListener("input", func(e Event) {
		if v := mvc.ViewFromEvent(e); v != nil {
			fmt.Println("Input event from view:", v)
			if v, err := strconv.ParseUint(v.Value(), 10, 64); err == nil {
				if v == 0 {
					response.SetData("")
				} else if v >= 99 {
					response.SetData("99+")
				} else {
					response.SetData(strconv.FormatUint(v, 10))
				}
			}
		}
	}), sourcecode()
}

func Example_Buttons_005() (mvc.View, string) {
	return bs.Container(
		bs.Button("Large Button", mvc.WithClass("m-1"), bs.WithSize(bs.Large)),
		bs.Button("Regular Button", mvc.WithClass("m-1")),
		bs.Button("Small Button", mvc.WithClass("m-1"), bs.WithSize(bs.Small)),
	), sourcecode()
}

func Example_Buttons_006() (mvc.View, string) {
	return bs.Container(
		bs.Button("Disabled Button", mvc.WithClass("m-1"), bs.WithDisabled(true)),
		bs.Button("Secondary Disabled Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1"), bs.WithDisabled(true)),
		bs.Button("Danger Disabled Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1"), bs.WithDisabled(true)),
	), sourcecode()
}

func Example_Buttons_007() (mvc.View, string) {
	return bs.Container(
		bs.Para("Click on a button in the group to see event handling"),
		bs.ButtonGroup(
			bs.Button("Left"),
			bs.Button("Middle"),
			bs.Button("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
	), sourcecode()
}

func Example_Buttons_008() (mvc.View, string) {
	return bs.Container(
		bs.Para("Click on a button in the group to see event handling"),
		bs.ButtonGroup(mvc.WithClass("m-2"),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("m-2"),
			bs.OutlineButton(bs.WithColor(bs.Danger), "Left"),
			bs.OutlineButton(bs.WithColor(bs.Primary), "Middle"),
			bs.OutlineButton(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
	), sourcecode()
}

func Example_Buttons_009() (mvc.View, string) {
	return bs.Container(
		bs.Para("Click on a button in the group to see event handling"),
		bs.ButtonGroup(mvc.WithClass("m-2"), bs.WithSize(bs.Small),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("m-2"), bs.WithSize(bs.Large),
			bs.OutlineButton(bs.WithColor(bs.Danger), "Left"),
			bs.OutlineButton(bs.WithColor(bs.Primary), "Middle"),
			bs.OutlineButton(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
	), sourcecode()
}

func Example_Buttons_010() (mvc.View, string) {
	return bs.Container(
		bs.Para("Click on a button in the group to see event handling"),
		bs.VButtonGroup(mvc.WithClass("m-2"),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(view, "Clicked!")
			}
		}),
	), sourcecode()
}

func Example_Buttons_011() (mvc.View, string) {
	response := bs.Para("Click a button in the toolbar")
	toolbarIconButton := func(label, iconName string, displayText bool) mvc.View {
		args := []any{
			bs.WithBorder(), bs.WithColor(bs.Secondary),
			mvc.WithAttr("title", label), mvc.WithAriaLabel(label),
			bs.Icon(iconName, mvc.WithClass("me-2"), mvc.WithAttr("aria-hidden", "true")),
		}
		if displayText {
			args = append(args, label)
		}
		return bs.Button(args)
	}
	return bs.Container(
		response,
		bs.ButtonToolbar(mvc.WithAttr("aria-label", "Formatting toolbar"),
			bs.ButtonGroup(mvc.WithClass("m-2"), mvc.WithAttr("aria-label", "Text formatting"),
				toolbarIconButton("Bold", "type-bold", true),
				toolbarIconButton("Italic", "type-italic", true),
				toolbarIconButton("Underline", "type-underline", true),
			),
			bs.ButtonGroup(mvc.WithClass("m-2"), mvc.WithAttr("aria-label", "Lists"),
				toolbarIconButton("Bulleted list", "list-ul", false),
				toolbarIconButton("Numbered list", "list-ol", false),
				toolbarIconButton("Checklist", "check2-square", false),
			),
			bs.ButtonGroup(mvc.WithClass("m-2"), mvc.WithAttr("aria-label", "Alignment"),
				toolbarIconButton("Align left", "text-left", false),
				toolbarIconButton("Align center", "text-center", false),
				toolbarIconButton("Align right", "text-right", false),
				toolbarIconButton("Justify", "justify", false),
			),
			bs.ButtonGroup(mvc.WithClass("m-2"), mvc.WithAttr("aria-label", "Insert"),
				toolbarIconButton("Link", "link-45deg", true),
				toolbarIconButton("Image", "image", true),
			),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				bs.SetContent(response, "Button '"+view.Name()+"' clicked!")
			}
		}),
	), sourcecode()
}
