package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TagExamples() mvc.View {
	return cds.Section(
		cds.LeadPara(
			`Tags use the `, cds.Code("cds-tag"), ` web component. `,
			`Set the colour with `, cds.Code("cds.WithTagType()"), `. `,
			`Add `, cds.Code("cds.WithTagFilter()"), ` to make a tag interactive, `,
			`or `, cds.Code("cds.WithTagSmall()"), ` for the compact size.`,
		),
		ExampleRow("Colors", Example_Tag_001, "All available tag color palettes."),
		ExampleRow("Small", Example_Tag_002, "Use cds.WithTagSmall() for a compact variant."),
		ExampleRow("Filter (dismissible)", Example_Tag_003, "Filter tags emit an event when dismissed; handle it to remove them from the DOM."),
		ExampleRow("Interactive", Example_Tag_004, "Click on a tag to toggle its selected state."),
	)
}

func Example_Tag_001() (mvc.View, string) {
	type typed struct {
		t TagType
		l string
	}
	tags := []typed{
		{cds.TagBlue, "Blue"},
		{cds.TagCyan, "Cyan"},
		{cds.TagTeal, "Teal"},
		{cds.TagGreen, "Green"},
		{cds.TagRed, "Red"},
		{cds.TagMagenta, "Magenta"},
		{cds.TagPurple, "Purple"},
		{cds.TagGray, "Gray"},
		{cds.TagCoolGray, "Cool gray"},
		{cds.TagWarmGray, "Warm gray"},
		{cds.TagHighContrast, "High contrast"},
		{cds.TagOutline, "Outline"},
	}
	items := make([]any, 0, len(tags))
	for _, t := range tags {
		items = append(items, cds.Tag(t.l, cds.WithTagType(t.t)))
	}
	return cds.Section(
		mvc.WithAttr("style", tagPreviewStyle),
		items,
	), sourcecode()
}

func Example_Tag_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", tagPreviewStyle),
		cds.Tag("Blue", cds.WithTagType(cds.TagBlue), cds.WithTagSmall()),
		cds.Tag("Green", cds.WithTagType(cds.TagGreen), cds.WithTagSmall()),
		cds.Tag("Red", cds.WithTagType(cds.TagRed), cds.WithTagSmall()),
		cds.Tag("Purple", cds.WithTagType(cds.TagPurple), cds.WithTagSmall()),
		cds.Tag("Outline", cds.WithTagType(cds.TagOutline), cds.WithTagSmall()),
	), sourcecode()
}

func Example_Tag_003() (mvc.View, string) {
	response := cds.Para(mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);"), "Dismiss a tag")
	container := cds.Section(
		mvc.WithAttr("style", tagPreviewStyle),
		cds.Tag("Design", cds.WithTagType(cds.TagBlue), cds.WithTagFilter()),
		cds.Tag("Development", cds.WithTagType(cds.TagGreen), cds.WithTagFilter()),
		cds.Tag("Research", cds.WithTagType(cds.TagPurple), cds.WithTagFilter()),
		cds.Tag("Testing", cds.WithTagType(cds.TagTeal), cds.WithTagFilter()),
	)
	container.AddEventListener("cds-tag-closed", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			response.Content("Dismissed: ", cds.Strong(el.TextContent()))
			el.Remove()
		}
	})
	return cds.Section(container, response), sourcecode()
}

func Example_Tag_004() (mvc.View, string) {
	response := cds.Para(mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);"), "Click a tag")
	container := cds.Section(
		mvc.WithAttr("style", tagPreviewStyle),
		cds.Tag("Draft", cds.WithTagType(cds.TagGray)),
		cds.Tag("In review", cds.WithTagType(cds.TagBlue)),
		cds.Tag("Approved", cds.WithTagType(cds.TagGreen)),
		cds.Tag("Rejected", cds.WithTagType(cds.TagRed)),
	)
	return cds.Section(
		container,
		response,
	).AddEventListener("click", func(e dom.Event) {
		t := mvc.ViewFromEvent(e)
		if t != nil && t.Name() == cds.ViewTag {
			response.Content("Clicked: ", cds.Strong(t.Root().TextContent()))
		}
	}), sourcecode()
}

const tagPreviewStyle = "display:flex;flex-wrap:wrap;align-items:center;gap:var(--cds-spacing-03,0.5rem);padding:var(--cds-spacing-04,0.75rem);border:1px solid var(--cds-border-subtle-01,#e0e0e0);"

// TagType alias so examples can reference cds.TagType directly without the import path.
type TagType = cds.TagType
