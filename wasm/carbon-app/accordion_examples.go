package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func AccordionExamples() mvc.View {
	return ExamplePage("Accordion",
		cds.LeadPara(
			`Accordions use the `, cds.Code("cds-accordion"), ` and `,
			cds.Code("cds-accordion-item"), ` web components. `,
			`Pass `, cds.Code("cds.WithAccordionOpen()"), ` to expand an item by default, `,
			cds.Code("cds.WithAccordionSize()"), ` to change row height, and `,
			cds.Code("cds.WithAccordionAlign(cds.AccordionAlignStart)"), ` to move the chevron to the left.`,
		),
		ExampleRow("Basic", Example_Accordion_001, "Four items; the second is expanded by default."),
		ExampleRow("Sizes", Example_Accordion_002, "Three accordions showing sm, md (default), and lg row heights."),
		ExampleRow("Align start", Example_Accordion_003, "Chevron placed on the left, flush with the content."),
		ExampleRow("Interactive", Example_Accordion_004, "Listen for cds-accordion-item-toggled to react to expand/collapse."),
	)
}

const accordionStyle = "margin-bottom:var(--cds-spacing-07,2rem);"

func Example_Accordion_001() (mvc.View, string) {
	return cds.Accordion(
		mvc.WithAttr("style", accordionStyle),
		cds.AccordionItem("Getting started", cds.Para(lorem)),
		cds.AccordionItem("Configuration", cds.Para(lorem), cds.WithAccordionOpen()),
		cds.AccordionItem("Advanced usage", cds.Para(lorem)),
		cds.AccordionItem("Troubleshooting", cds.Para(lorem)),
	), sourcecode()
}

func Example_Accordion_002() (mvc.View, string) {
	makeAcc := func(size cds.AccordionSize, label string) mvc.View {
		return cds.Section(
			mvc.WithAttr("style", "margin-bottom:var(--cds-spacing-06,1.5rem);"),
			cds.Heading(5, label),
			cds.Accordion(mvc.WithAttr("style", accordionStyle), cds.WithAccordionSize(size),
				cds.AccordionItem("First item", cds.Para(lorem)),
				cds.AccordionItem("Second item", cds.Para(lorem), cds.WithAccordionOpen()),
			),
		)
	}
	return cds.Section(
		makeAcc(cds.AccordionSM, "Small (sm)"),
		makeAcc(cds.AccordionMD, "Medium (md — default)"),
		makeAcc(cds.AccordionLG, "Large (lg)"),
	), sourcecode()
}

func Example_Accordion_003() (mvc.View, string) {
	return cds.Accordion(
		mvc.WithAttr("style", accordionStyle),
		cds.WithAccordionAlign(cds.AccordionAlignStart),
		cds.AccordionItem("Design principles", cds.Para(lorem)),
		cds.AccordionItem("Accessibility", cds.Para(lorem), cds.WithAccordionOpen()),
		cds.AccordionItem("Responsive layout", cds.Para(lorem)),
	), sourcecode()
}

func Example_Accordion_004() (mvc.View, string) {
	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);color:var(--cds-text-secondary,#525252);"),
		"Expand or collapse an item…",
	)
	acc := cds.Accordion(
		mvc.WithAttr("style", accordionStyle),
		cds.AccordionItem("First section", cds.Para(lorem)),
		cds.AccordionItem("Second section", cds.Para(lorem)),
		cds.AccordionItem("Third section", cds.Para(lorem)),
	)
	acc.AddEventListener("cds-accordion-item-toggled", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			title := el.GetAttribute("title")
			// LitElement reflects `open` asynchronously; the attribute still holds
			// the OLD state when the event fires, so we invert it to get new state.
			expanded := !el.HasAttribute("open")
			action := "collapsed"
			if expanded {
				action = "expanded"
			}
			status.Content(cds.Strong(title), " — "+action)
		}
	})
	return cds.Section(acc, status), sourcecode()
}
