package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func ButtonExamples() mvc.View {
	return bs.Container(
		ButtonSingleExamples(),
		bs.HRule(),
		ButtonGroupExamples(),
		bs.HRule(),
		ButtonToolbarExamples(),
	)
}

func ButtonSingleExamples() mvc.View {
	return bs.Container(
		bs.Heading(1, "Button Examples"),
		bs.HRule(),
		bs.Heading(3, "Standard Buttons"),
		bs.Button(mvc.WithClass("mx-1")).Content("Default Button"),
		bs.Button(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1")).Content("Secondary Button"),
		bs.Button(bs.WithColor(bs.Danger), mvc.WithClass("mx-1")).Content("Danger Button"),
		bs.HRule(),
		bs.Heading(3, "Outline Buttons"),
		bs.OutlineButton(mvc.WithClass("mx-1")).Content("Default Button"),
		bs.OutlineButton(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1")).Content("Secondary Button"),
		bs.OutlineButton(bs.WithColor(bs.Danger), mvc.WithClass("mx-1")).Content("Danger Button"),
		bs.HRule(),
		bs.Heading(3, "Close Buttons"),
		bs.CloseButton(mvc.WithClass("mx-1")),
		bs.HRule(),
		bs.Heading(3, "Large Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithSize(bs.Large)).Content("Default Button"),
		bs.HRule(),
		bs.Heading(3, "Small Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithSize(bs.Small)).Content("Default Button"),
		bs.HRule(),
		bs.Heading(3, "Disabled Buttons"),
		bs.Button(mvc.WithClass("mx-1"), bs.WithDisabled(true), "Default Button"),
		bs.Button(bs.WithColor(bs.Secondary), mvc.WithClass("mx-1"), bs.WithDisabled(true), "Secondary Button"),
		bs.Button(bs.WithColor(bs.Danger), mvc.WithClass("mx-1"), bs.WithDisabled(true), "Danger Button"),
		bs.HRule(),
	)
}

func ButtonToolbarExamples() mvc.View {
	return bs.Container(
		bs.Heading(1, "Button Toolbar Examples"),
		bs.HRule(),
		bs.ButtonToolbar(mvc.WithAttr("aria-label", "Formatting toolbar"),
			bs.ButtonGroup(mvc.WithClass("me-2"), mvc.WithAttr("aria-label", "Text formatting"),
				toolbarIconButton("Bold", "type-bold"),
				toolbarIconButton("Italic", "type-italic"),
				toolbarIconButton("Underline", "type-underline"),
			),
			bs.ButtonGroup(mvc.WithClass("me-2"), mvc.WithAttr("aria-label", "Lists"),
				toolbarIconButton("Bulleted list", "list-ul"),
				toolbarIconButton("Numbered list", "list-ol"),
				toolbarIconButton("Checklist", "check2-square"),
			),
			bs.ButtonGroup(mvc.WithClass("me-2"), mvc.WithAttr("aria-label", "Alignment"),
				toolbarIconButton("Align left", "text-left"),
				toolbarIconButton("Align center", "text-center"),
				toolbarIconButton("Align right", "text-right"),
				toolbarIconButton("Justify", "justify"),
			),
			bs.ButtonGroup(mvc.WithAttr("aria-label", "Insert"),
				toolbarIconButton("Insert link", "link-45deg"),
				toolbarIconButton("Insert image", "image"),
			),
		),
		bs.HRule(),
	)
}

func toolbarIconButton(label, iconName string) mvc.View {
	return bs.OutlineButton(
		bs.WithColor(bs.Secondary),
		mvc.WithAttr("title", label),
		mvc.WithAttr("aria-label", label),
		mvc.WithClass("my-1"),
		bs.Icon(iconName, mvc.WithClass("bi", "me-2"), mvc.WithAttr("aria-hidden", "true")),
		label,
	)
}

func ButtonGroupExamples() mvc.View {
	return bs.Container(
		bs.Heading(1, "Button Groups"),
		bs.HRule(),
		bs.Heading(3, "Horizontal Button Group"),
		bs.ButtonGroup(
			bs.Button().Content("Left"),
			bs.Button().Content("Middle"),
			bs.Button().Content("Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2"),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2"),
			bs.OutlineButton(bs.WithColor(bs.Danger), "Left"),
			bs.OutlineButton(bs.WithColor(bs.Primary), "Middle"),
			bs.OutlineButton(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.ButtonGroup(mvc.WithClass("mx-2"), bs.WithSize(bs.Small),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}), bs.ButtonGroup(mvc.WithClass("mx-2"), bs.WithSize(bs.Large),
			bs.Button(bs.WithColor(bs.Danger), "Left"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Right"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
		bs.HRule(),
		bs.Heading(3, "Vertical Button Group"),
		bs.VButtonGroup(mvc.WithClass("mx-2"),
			bs.Button(bs.WithColor(bs.Danger), "Top"),
			bs.Button(bs.WithColor(bs.Warning), "Middle"),
			bs.Button(bs.WithColor(bs.Success), "Bottom"),
		).AddEventListener("click", func(e Event) {
			if view := mvc.ViewFromEvent(e); view != nil {
				view.Content("Clicked!")
			}
		}),
	)
}
