package main

import (
	// Packages

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func ProgressExamples() mvc.View {
	var progress, striped mvc.ViewWithValue
	return bs.Container(
		bs.Heading(3, "Simple Progress Bar"),
		bs.Grid(
			func() mvc.View {
				progress = bs.Progress()
				return progress
			}(),
			bs.RangeInput("range-input").AddEventListener("input", func(e Event) {
				if view := mvc.ViewFromEvent(e).(mvc.ViewWithValue); view != nil {
					progress.SetValue(view.Value())
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
				if view := mvc.ViewFromEvent(e).(mvc.ViewWithValue); view != nil {
					striped.SetValue(view.Value())
				}
			}),
		),
		bs.HRule(),
		bs.Heading(3, "Colours"),
		bs.Grid(
			bs.StripedProgress(bs.WithColor(bs.Secondary)).SetValue("35"),
			bs.StripedProgress(bs.WithColor(bs.Success)).SetValue("75"),
			bs.StripedProgress(bs.WithColor(bs.Danger)).SetValue("15"),
			bs.StripedProgress(bs.WithColor(bs.Warning)).SetValue("25"),
		),
	)
}
