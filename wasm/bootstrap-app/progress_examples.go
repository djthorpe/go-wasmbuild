package main

import (
	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func ProgressExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Progress Examples"),
		bs.HRule(),
		bs.Heading(3, "Simple Progress Bar", mvc.WithClass("mt-5")), Example(Example_Progress_001),
		bs.Heading(3, "Striped Progress Bar", mvc.WithClass("mt-5")), Example(Example_Progress_002),
		bs.Heading(3, "Color Variants", mvc.WithClass("mt-5")), Example(Example_Progress_003),
	)
}

func Example_Progress_001() (mvc.View, string) {
	progress := bs.Progress()
	return bs.Container(
		bs.Grid(
			progress,
			bs.RangeInput("range-progress-basic").AddEventListener("input", func(e Event) {
				if view := mvc.ViewFromEvent(e); view != nil {
					progress.Set(view.Value())
				}
			}),
		),
	), sourcecode()
}

func Example_Progress_002() (mvc.View, string) {
	striped := bs.StripedProgress()
	return bs.Container(
		bs.Grid(
			striped,
			bs.RangeInput("range-progress-striped").AddEventListener("input", func(e Event) {
				if view := mvc.ViewFromEvent(e); view != nil {
					striped.Set(view.Value())
				}
			}),
		),
	), sourcecode()
}

func Example_Progress_003() (mvc.View, string) {
	return bs.Grid(
		bs.StripedProgress(bs.WithColor(bs.Secondary)).Set("35"),
		bs.StripedProgress(bs.WithColor(bs.Success)).Set("75"),
		bs.StripedProgress(bs.WithColor(bs.Danger)).Set("15"),
		bs.StripedProgress(bs.WithColor(bs.Warning)).Set("25"),
	), sourcecode()
}
