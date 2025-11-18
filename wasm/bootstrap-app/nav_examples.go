package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Navigation Examples"), bs.HRule(),
		bs.Heading(3, "Accordion", mvc.WithClass("mt-5")), Example(Example_Accordion_001),
		bs.Heading(3, "Flush Accordion", mvc.WithClass("mt-5")), Example(Example_Accordion_002),
		bs.Heading(3, "Dark Accordion", mvc.WithClass("mt-5")), Example(Example_Accordion_003),
	)
}

func Example_Accordion_001() (mvc.View, string) {
	return bs.Accordion("accordion-001",
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the first item's accordion body. "),
				"It is shown by default, until the collapse plugin adds the appropriate classes ",
				"that we use to style each element. These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #1"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the second item's accordion body. "),
				"These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #2"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the third item's accordion body. "),
				"These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #3"),
	), sourcecode()
}

func Example_Accordion_002() (mvc.View, string) {
	return bs.FlushAccordion("accordion-002",
		bs.AccordionItem(
			bs.Para(
				bs.Code("FlushAccordion"),
				" removes some borders and rounded corners to render accordions edge-to-edge ",
				"with their parent container.",
			),
		).Header("Accordion Item #1"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the second item's accordion body. "),
				"These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #2"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the third item's accordion body. "),
				"These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #3"),
	), sourcecode()
}

func Example_Accordion_003() (mvc.View, string) {
	return bs.Accordion("accordion-003", bs.WithTheme("dark"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the dark-themed accordion."),
				`Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
				sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
				quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis 
				aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.`,
			),
		).Header("Accordion Item #1"),
		bs.AccordionItem(
			bs.Para(
				bs.Strong("This is the second item's accordion body. "),
				"These classes control the overall appearance, ",
				"as well as the showing and hiding via CSS transitions. You can modify any of this ",
				"with custom CSS or overriding our default variables. It’s also worth noting that ",
				"just about any HTML can go within the ", bs.Code(".accordion-body"), ", ",
				"though the transition does limit overflow.",
			),
		).Header("Accordion Item #2"),
	), sourcecode()
}
