package main

import (
	// Packages
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ButtonExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("button_examples.md"),
		bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Buttons with Color"), Example(Example_Buttons_001),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Outline Buttons"), Example(Example_Buttons_002),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Close Button"), Example(Example_Buttons_003),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Button with Indicator"), Example(Example_Buttons_004),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Button Sizes"), Example(Example_Buttons_005),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Button Group"), Example(Example_Buttons_006),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Vertical Button Group"), Example(Example_Buttons_007),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(4, "Button Toolbar Group"), Example(Example_Buttons_008),
		),
	)
}

func Example_Buttons_001() (mvc.View, string) {
	response := bs.Strong("Click a button", mvc.WithClass("m-3"))
	return bs.Container(
		bs.Button("Default Button", mvc.WithClass("m-1")),
		bs.Button("Secondary Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1")),
		bs.Button("Danger Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1")),
		response,
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event)
		if button.Name() == bs.ViewButton {
			response.Content("Clicked: ", button.Root().TextContent())
		}
	}), sourcecode()
}

func Example_Buttons_002() (mvc.View, string) {
	response := bs.Para("Click a button")
	return bs.Container(
		response,
		bs.OutlineButton("Default Button", mvc.WithClass("m-1")),
		bs.OutlineButton("Secondary Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1")),
		bs.OutlineButton("Danger Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1")),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event)
		if button.Name() == bs.ViewButton {
			response.Content("Button clicked: ", button.Root().TextContent())
		}
	}), sourcecode()
}

func Example_Buttons_003() (mvc.View, string) {
	response := bs.Para("Click the close button")
	return bs.Container(
		response,
		bs.CloseButton(bs.WithBorder(bs.Dark), mvc.WithClass("m-1")),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event)
		if button.Name() == bs.ViewButton {
			response.Content("Close Button Clicked")
		}
	}), sourcecode()
}

func Example_Buttons_004() (mvc.View, string) {
	response := mvc.Text("50")
	return bs.Row(
		bs.Col(
			bs.Button("Inbox", mvc.WithClass("m-3")).Label(response),
		),
		bs.Col(
			bs.RangeInput("value", bs.WithMinMax(0, 10)).Label("Change Unread Count"),
		),
	).AddEventListener("input", func(e dom.Event) {
		input := mvc.ViewFromEvent(e, bs.ViewInput)
		if input != nil {
			fmt.Println("Input value:", input.(bs.DataView).Value())
		}

	}), sourcecode()
}

func Example_Buttons_005() (mvc.View, string) {
	return bs.Container(
		bs.Button("Normal Button", mvc.WithClass("m-3")),
		bs.Button("Large Button", bs.WithSize(bs.Large), mvc.WithClass("m-3")),
		bs.Button("Small Button", bs.WithSize(bs.Small), mvc.WithClass("m-3")),
	), sourcecode()
}

func Example_Buttons_006() (mvc.View, string) {
	response := bs.Para("Click a button")
	return bs.Container(
		response,
		bs.ButtonGroup(
			mvc.WithClass("m-1"),
			bs.Button(bs.Icon("align-start", mvc.WithClass("me-1")), "Start"),
			bs.Button(bs.Icon("align-center", mvc.WithClass("me-1")), "Center"),
			bs.Button("End", bs.Icon("align-end", mvc.WithClass("ms-1"))),
		),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event, bs.ViewButton)
		if button != nil {
			response.Content("Button clicked: ", button.Root().TextContent())
		}
	}), sourcecode()
}

func Example_Buttons_007() (mvc.View, string) {
	response := bs.Para("Click a button")
	return bs.Container(
		response,
		bs.VButtonGroup(
			bs.OutlineButton("Option A"),
			bs.OutlineButton("Option B"),
			bs.OutlineButton("Option C"),
		),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event, bs.ViewButton)
		if button != nil {
			response.Content("Button clicked: ", button.Root().TextContent())
		}
	}), sourcecode()
}

func Example_Buttons_008() (mvc.View, string) {
	response := bs.Para("Click a button")
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
		bs.ButtonToolbar(
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
		),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event, bs.ViewButton)
		if button != nil {
			response.Content("Button clicked: ", button.Root().GetAttribute("title"))
		}
	}), sourcecode()
}
