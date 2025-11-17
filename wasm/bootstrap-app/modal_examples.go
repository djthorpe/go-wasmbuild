package main

import (
	// Packages

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ModalExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Modal Examples"),
		bs.HRule(),
		bs.Heading(3, "Modal Dialog", mvc.WithClass("mt-5")), Example(Example_Dialog_001),
		bs.Heading(3, "Sticky Modal Dialog", mvc.WithClass("mt-5")), Example(Example_Dialog_002),
		bs.Heading(3, "Modal Color", mvc.WithClass("mt-5")), Example(Example_Dialog_003),
		bs.Heading(3, "Offcanvas", mvc.WithClass("mt-5")), Example(Example_Offcanvas_001),
		bs.Heading(3, "Offcanvas Position", mvc.WithClass("mt-5")), Example(Example_Offcanvas_002),
		bs.Heading(3, "Offcanvas Color", mvc.WithClass("mt-5")), Example(Example_Offcanvas_003),
	)
}

func Example_Offcanvas_001() (mvc.View, string) {
	return bs.Container(
		bs.Offcanvas("offcanvas_001", "This is the offcanvas content!"),
		bs.Button(bs.WithOffcanvas("offcanvas_001"), "Start"),
	), sourcecode()
}

func Example_Offcanvas_002() (mvc.View, string) {
	offcanvas := bs.Offcanvas(
		"offcanvas_002",
		"This is the offcanvas content!",
	).Header(
		bs.Heading(5, "Offcanvas"),
		bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
	).(mvc.ViewWithVisibility)
	return bs.Container(
		offcanvas,
		bs.ButtonGroup(
			bs.Button("Start", mvc.WithID("start")),
			bs.Button("End", mvc.WithID("end")),
			bs.Button("Top", mvc.WithID("top")),
			bs.Button("Bottom", mvc.WithID("bottom")),
		).AddEventListener("click", func(e dom.Event) {
			v := mvc.ViewFromEvent(e)
			switch v.ID() {
			case "start":
				offcanvas.Apply(bs.WithPosition(bs.Start))
			case "end":
				offcanvas.Apply(bs.WithPosition(bs.End))
			case "top":
				offcanvas.Apply(bs.WithPosition(bs.Top))
			case "bottom":
				offcanvas.Apply(bs.WithPosition(bs.Bottom))
			}

			// Trigger the offcanvas programmatically
			offcanvas.Show()
		}),
	), sourcecode()
}

func Example_Offcanvas_003() (mvc.View, string) {
	return bs.Container(
		bs.Offcanvas("offcanvas_003", "This is the offcanvas content!", bs.WithColor(bs.Secondary), bs.WithTheme("dark")).Header(
			bs.Heading(5, "Offcanvas"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		),
		bs.Button(bs.WithOffcanvas("offcanvas_003"), "Start"),
	), sourcecode()
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
			"Sticky modals will not close when clicking outside of it, or pressing the ESC key.",
			"Press the Close button below to close this modal programmatically.",
		).Header(
			bs.Heading(5, "Modal Dialog"),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "modal")),
		).Footer(
			bs.Button("Close"),
		),
		bs.Button(bs.WithModal("dialog_002"), "Open Sticky Modal"),
	), sourcecode()
}

func Example_Dialog_003() (mvc.View, string) {
	modal := bs.Modal("dialog_003", bs.WithTheme("dark"),
		bs.ButtonToolbar(
			bs.ButtonGroup(bs.WithSize(bs.Small), mvc.WithClass("m-1"),
				bs.Button("Light", mvc.WithID("light")),
				bs.Button("Dark", mvc.WithID("dark")),
			),
			bs.ButtonGroup(bs.WithSize(bs.Small), mvc.WithClass("m-1"),
				bs.Button("Primary", mvc.WithID("primary")),
				bs.Button("Secondary", mvc.WithID("secondary")),
				bs.Button("Success", mvc.WithID("success")),
				bs.Button("Danger", mvc.WithID("danger")),
				bs.Button("Warning", mvc.WithID("warning")),
				bs.Button("Info", mvc.WithID("info")),
			),
		),
	)
	return bs.Container(
		modal.Header(
			bs.Heading(3, "Set modal theme and colour"),
		).Footer(
			bs.Button("Close"),
		).AddEventListener("click", func(e dom.Event) {
			v := mvc.ViewFromEvent(e)
			if v.Name() != bs.ViewButton || v.ID() == "" {
				return
			}
			switch v.ID() {
			case "light", "dark":
				modal.Apply(bs.WithTheme(bs.Color(v.ID())))
			default:
				// TODO: the colour should be applied to the modal-content not modal
				modal.Apply(bs.WithColor(bs.Color(v.ID())))
			}
		}),
		bs.Button(bs.WithModal("dialog_003"), "Open Color"),
	), sourcecode()
}
