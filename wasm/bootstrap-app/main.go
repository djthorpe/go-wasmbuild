package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application displays examples of Bootstrap components
func main() {
	mvc.New().Content(
		bs.Link("#text", mvc.WithClass("m-2")).Content("Text"),
		bs.VRule(),
		bs.Link("#border", mvc.WithClass("m-2")).Content("Borders"),
		bs.VRule(),
		bs.Link("#badge", mvc.WithClass("m-2")).Content("Badges"),
		bs.VRule(),
		bs.Link("#list", mvc.WithClass("m-2")).Content("Lists"),
		bs.VRule(),
		bs.Link("#icon", mvc.WithClass("m-2")).Content("Icons"),
		bs.VRule(),
		bs.Link("#button", mvc.WithClass("m-2")).Content("Buttons"),
		bs.VRule(),
		bs.Link("#card", mvc.WithClass("m-2")).Content("Cards"),
		bs.VRule(),
		bs.Link("#modal", mvc.WithClass("m-2")).Content("Modal"),
		bs.VRule(),
		bs.Link("#input", mvc.WithClass("m-2")).Content("Input"),
		bs.VRule(),
		bs.Link("#tooltips", mvc.WithClass("m-2")).Content("Tooltips & Popovers"),
		bs.VRule(),
		bs.Link("#progress", mvc.WithClass("m-2")).Content("Progress Bars"),
		bs.VRule(),
		bs.Link("#navbar", mvc.WithClass("m-2")).Content("Navbars"),
		bs.VRule(),
		bs.Link("#nav", mvc.WithClass("m-2")).Content("Navigation & Pagination"),
		bs.VRule(),
		bs.Link("#table", mvc.WithClass("m-2")).Content("Tables"),
		bs.VRule(),
		bs.Link("#alert", mvc.WithClass("m-2")).Content("Alerts & Toasts"),

		mvc.Router(mvc.WithClass("container-fluid", "my-2")).
			Page("#text", TextExamples()).
			Page("#border", BorderExamples()).
			Page("#badge", BadgeExamples()).
			Page("#list", ListExamples()).
			Page("#icon", IconExamples()).
			Page("#button", ButtonExamples()).
			Page("#card", CardExamples()).
			Page("#modal", ModalExamples()).
			Page("#input", InputExamples()).
			Page("#tooltips", TooltipExamples()).
			Page("#progress", ProgressExamples()).
			Page("#navbar", NavBarExamples()).
			Page("#nav", NavExamples()).
			Page("#table", TableExamples()).
			Page("#alert", AlertExamples()),
	)

	// Wait
	select {}
}
