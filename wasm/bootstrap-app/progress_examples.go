package main

import (
	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func ProgressExamples() mvc.View {
	var progress, striped mvc.View
	return bs.Container(
		bs.Heading(3, "Simple Progress Bar"),
		bs.Grid(
			func() mvc.View {
				progress = bs.Progress()
				return progress
			}(),
			bs.RangeInput("range-input").AddEventListener("input", func(e Event) {
				if view := mvc.ViewFromEvent(e); view != nil {
					progress.Set(view.Value())
				}
			}),
		),
		bs.HRule(),
		bs.Heading(3, "Striped Progress Bar"),
		bs.Grid(
			func() mvc.View {
				striped = bs.StripedProgress()
				return striped
			}(),
			bs.RangeInput("range-input").AddEventListener("input", func(e Event) {
				if view := mvc.ViewFromEvent(e); view != nil {
					striped.Set(view.Value())
				}
			}),
		),
		bs.HRule(),
		bs.Heading(3, "Colours"),
		bs.Grid(
			bs.StripedProgress(bs.WithColor(bs.Secondary)).Set("35"),
			bs.StripedProgress(bs.WithColor(bs.Success)).Set("75"),
			bs.StripedProgress(bs.WithColor(bs.Danger)).Set("15"),
			bs.StripedProgress(bs.WithColor(bs.Warning)).Set("25"),
		),
	)
}
