package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TooltipExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Tooltip Examples"), bs.HRule(),
		bs.Heading(3, "Button Tooltips", mvc.WithClass("mt-5")), Example(Example_Tooltip_001),
		bs.Heading(3, "Icon Tooltips", mvc.WithClass("mt-5")), Example(Example_Tooltip_002),
	)
}

func Example_Tooltip_001() (mvc.View, string) {
	return bs.Container(
		mvc.WithClass("d-flex", "flex-wrap", "gap-3"),
		bs.Button("Tooltip on top", bs.WithTooltip("Tooltip on top"), mvc.WithAttr("data-bs-placement", "top")),
		bs.Button("Tooltip on right", bs.WithTooltip("Tooltip on right"), mvc.WithAttr("data-bs-placement", "right")),
		bs.Button("Tooltip on bottom", bs.WithTooltip("Tooltip on bottom"), mvc.WithAttr("data-bs-placement", "bottom")),
		bs.Button("Tooltip on left", bs.WithTooltip("Tooltip on left"), mvc.WithAttr("data-bs-placement", "left")),
	), sourcecode()
}

func Example_Tooltip_002() (mvc.View, string) {
	iconButton := func(icon, label, placement string) mvc.View {
		return bs.Button(
			bs.Icon(icon, mvc.WithClass("me-2")),
			label,
			bs.WithTooltip("Shortcut: "+label),
			mvc.WithAttr("data-bs-placement", placement),
		)
	}
	return bs.Container(
		mvc.WithClass("d-flex", "flex-wrap", "gap-3"),
		iconButton("type-bold", "Bold", "top"),
		iconButton("type-italic", "Italic", "top"),
		iconButton("type-underline", "Underline", "top"),
		iconButton("link-45deg", "Link", "bottom"),
		iconButton("image", "Image", "bottom"),
	), sourcecode()
}
