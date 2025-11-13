package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func IconExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Icon Examples"), bs.HRule(),
		bs.Heading(3, "Common Glyphs", mvc.WithClass("mt-5")), Example(Example_Icon_001),
		bs.Heading(3, "Glyphs with Color", mvc.WithClass("mt-5")), Example(Example_Icon_002),
		bs.Heading(3, "Buttons with Icons", mvc.WithClass("mt-5")), Example(Example_Icon_003),
	)
}

func Example_Icon_001() (mvc.View, string) {
	return bs.Grid().Append(
		bs.Icon("alarm", mvc.WithClass("fs-1")),
		bs.Icon("award", mvc.WithClass("fs-1")),
		bs.Icon("bell", mvc.WithClass("fs-1")),
		bs.Icon("chat-dots", mvc.WithClass("fs-1")),
		bs.Icon("cloud", mvc.WithClass("fs-1")),
	), sourcecode()
}

func Example_Icon_002() (mvc.View, string) {
	return bs.Grid().Append(
		bs.Icon("heart-fill", bs.WithColor(bs.Danger), mvc.WithClass("fs-1")),
		bs.Icon("sun-fill", bs.WithColor(bs.Warning), mvc.WithClass("fs-1")),
		bs.Icon("moon-stars", bs.WithColor(bs.Primary), mvc.WithClass("fs-1")),
		bs.Icon("toggle2-on", bs.WithColor(bs.Success), mvc.WithClass("fs-1")),
		bs.Icon("wifi", bs.WithColor(bs.Info), mvc.WithClass("fs-1")),
	), sourcecode()
}

func Example_Icon_003() (mvc.View, string) {
	withButtonClasses := func(opts ...any) []any {
		base := []any{
			bs.WithSize(bs.Small),
			mvc.WithClass("d-flex", "align-items-center", "gap-2", "justify-content-center", "my-3"),
		}
		return append(base, opts...)
	}
	return bs.Grid().Append(
		bs.Button("Share ", bs.Icon("share"), withButtonClasses(bs.WithColor(bs.Primary))),
		bs.OutlineButton("Download ", bs.Icon("cloud-arrow-down"), withButtonClasses(bs.WithColor(bs.Success))),
		bs.Button("Next ", bs.Icon("arrow-right"), withButtonClasses(bs.WithColor(bs.Success))),
		bs.Button(bs.Icon("trash"), " Delete", withButtonClasses(bs.WithColor(bs.Danger))),
		bs.OutlineButton(bs.Icon("arrow-repeat"), " Refresh", withButtonClasses(bs.WithColor(bs.Dark))),
	), sourcecode()
}
