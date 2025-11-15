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
		bs.Heading(3, "Offcanvas", mvc.WithClass("mt-5")), Example(Example_Offcanvas_001),
		bs.Heading(3, "Offcanvas Position", mvc.WithClass("mt-5")), Example(Example_Offcanvas_002),
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

/*
func DialogExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(3).Content("Modal Examples"),
		bs.Modal("modal1").Header(
			mvc.HTML("H4", mvc.WithInnerText("This is the title")),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "modal")),
		).Content(
			"This is the modal content!",
		),
		bs.ButtonGroup().Content(
			bs.Button(bs.WithModal("modal1"), "Open Modal"),
		),
	)
}
*/
