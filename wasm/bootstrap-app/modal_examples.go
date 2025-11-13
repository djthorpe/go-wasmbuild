package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ModalExamples() mvc.View {
	return bs.Container(
		DialogExamples(),
		OffcanvasExamples(),
	)
}

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

func OffcanvasExamples() mvc.View {
	return bs.Container().Content(
		bs.Heading(3).Content("Offcanvas Examples"),
		bs.Offcanvas("start", bs.WithPosition(bs.Start)).Header(
			mvc.HTML("H4", mvc.WithInnerText("This is the offcanvas title")),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		).Content(
			"This is the offcanvas content!",
		),
		bs.Offcanvas("end", bs.WithPosition(bs.End), bs.WithTheme(bs.Dark)).Header(
			mvc.HTML("H4", mvc.WithInnerText("This is the offcanvas title")),
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "offcanvas")),
		).Content(
			"This is the offcanvas content!",
		),
		bs.Offcanvas("top", bs.WithPosition(bs.Top)).Content("This is the offcanvas content!"),
		bs.Offcanvas("bottom", bs.WithPosition(bs.Bottom)).Content("This is the offcanvas content!"),
		bs.ButtonGroup().Content(
			bs.Button(bs.WithOffcanvas("start")).Content("Start"),
			bs.Button(bs.WithOffcanvas("end"), bs.WithColor(bs.Dark)).Content("End"),
			bs.Button(bs.WithOffcanvas("top")).Content("Top"),
			bs.Button(bs.WithOffcanvas("bottom")).Content("Bottom"),
		),
	)
}
