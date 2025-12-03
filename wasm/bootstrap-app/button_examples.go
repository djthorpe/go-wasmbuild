package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ButtonExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("button_examples.md"),
		bs.HRule(),
		bs.Heading(4, "Buttons with Color"), Example(Example_Buttons_001),
		bs.Heading(4, "Outline Buttons"), Example(Example_Buttons_002),
		bs.Heading(4, "Close Button"), Example(Example_Buttons_003),
		bs.Heading(4, "Button with Indicator"), Example(Example_Buttons_004),
		bs.Heading(4, "Button Sizes"), Example(Example_Buttons_005),
		bs.Heading(4, "Button Group"), Example(Example_Buttons_006),
	)
}

func Example_Buttons_001() (mvc.View, string) {
	response := bs.Para("Click a button")
	return bs.Container(
		response,
		bs.Button("Default Button", mvc.WithClass("m-1")),
		bs.Button("Secondary Button", bs.WithColor(bs.Secondary), mvc.WithClass("m-1")),
		bs.Button("Danger Button", bs.WithColor(bs.Danger), mvc.WithClass("m-1")),
	).AddEventListener("click", func(event dom.Event) {
		button := mvc.ViewFromEvent(event)
		if button.Name() == bs.ViewButton {
			response.Content("Button clicked: ", button.Root().TextContent())
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
	return bs.Container(
		bs.Button("Inbox", mvc.WithClass("m-3")).Label(response),
	), sourcecode()
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
