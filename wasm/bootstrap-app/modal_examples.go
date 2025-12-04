package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ModalExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("modal_examples.md"),
		bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("p-3"),
			bs.Heading(3, "Modal Dialog"), Example(Example_Dialog_001),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(3, "Sticky Modal Dialog"), Example(Example_Dialog_002),
		), bs.Container(
			bs.WithBorder(), bs.WithColor(bs.Light), mvc.WithClass("my-3"), mvc.WithClass("p-3"),
			bs.Heading(3, "Dialog Size"), Example(Example_Dialog_003),
		),
	)
}

func Example_Dialog_001() (mvc.View, string) {
	return bs.Container(
		bs.Modal("dialog_001",
			"Modals can be dismissed by clicking the close button above, clicking outside the modal, or pressing the ESC key.",
		).Header(
			bs.Heading(5, "Modal Dialog"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "modal")),
		).Footer(
			"This is the footer",
		),
		bs.Button(bs.WithModal("dialog_001"), "Open Modal"),
	), sourcecode()
}

func Example_Dialog_002() (mvc.View, string) {
	return bs.Container(
		bs.StickyModal("dialog_002",
			"Sticky modals will not close when clicking outside of it, or pressing the ESC key. ",
			"Press the Close button below to close this modal programmatically.",
		).Header(
			bs.Heading(5, "Modal Dialog"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "modal")),
		).Footer(
			bs.Button("Close", mvc.WithAttr("data-bs-dismiss", "modal")),
		),
		bs.Button(bs.WithModal("dialog_002"), "Open Sticky Modal"),
	), sourcecode()
}

func Example_Dialog_003() (mvc.View, string) {
	modal := bs.Modal("dialog_003",
		bs.WithSize(bs.Small),
		"Showing Modal of different sizes.",
	).Header(
		bs.Heading(5, "Modal Dialog"),
		bs.CloseButton(mvc.WithID("close_modal")),
	)
	return bs.Container(
		modal,
		bs.Button(mvc.WithClass("m-1"), mvc.WithID("small_modal"), "Open Small Modal"),
		bs.Button(mvc.WithClass("m-1"), mvc.WithID("large_modal"), "Open Large Modal"),
		bs.Button(mvc.WithClass("m-1"), mvc.WithID("xlarge_modal"), "Open XLarge Modal"),
	).AddEventListener("click", func(e dom.Event) {
		button := mvc.ViewFromEvent(e, bs.ViewButton)
		switch {
		case button == nil:
			return
		case button.ID() == "close_modal":
			modal.Hide()
		case button.ID() == "small_modal":
			modal.Apply(bs.WithSize(bs.Small))
			modal.Show()
		case button.ID() == "large_modal":
			modal.Apply(bs.WithSize(bs.Large))
			modal.Show()
		case button.ID() == "xlarge_modal":
			modal.Apply(bs.WithSize(bs.XLarge))
			modal.Show()
		}
	}), sourcecode()
}
