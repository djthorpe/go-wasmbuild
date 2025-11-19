package main

import mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

// examplePage describes a routable example page and the view builder that renders it.
type examplePage struct {
	id    string
	label string
	build func() mvc.View
}

// exampleGroup organizes related pages for navigation dropdowns.
type exampleGroup struct {
	label string
	pages []examplePage
}

var exampleGroups = []exampleGroup{
	{
		label: "Typography",
		pages: []examplePage{
			{id: "#text", label: "Text", build: TextExamples},
			{id: "#list", label: "Lists", build: ListExamples},
			{id: "#badge", label: "Badges", build: BadgeExamples},
			{id: "#icon", label: "Icons", build: IconExamples},
		},
	},
	{
		label: "Interactivity",
		pages: []examplePage{
			{id: "#button", label: "Buttons", build: ButtonExamples},
			{id: "#modal", label: "Modals", build: ModalExamples},
			{id: "#alert", label: "Alerts & Toasts", build: AlertExamples},
			{id: "#tooltips", label: "Tooltips", build: TooltipExamples},
		},
	},
	{
		label: "Forms & Controls",
		pages: []examplePage{
			{id: "#input", label: "Input", build: InputExamples},
		},
	},
	{
		label: "Navigation",
		pages: []examplePage{
			{id: "#navbar", label: "Navbar", build: NavBarExamples},
			{id: "#nav", label: "Navigation", build: NavExamples},
		},
	},
	{
		label: "Decoration",
		pages: []examplePage{
			{id: "#border", label: "Borders", build: BorderExamples},
			{id: "#card", label: "Cards", build: CardExamples},
		},
	},
	{
		label: "Data",
		pages: []examplePage{
			{id: "#table", label: "Tables", build: TableExamples},
		},
	},
	{
		label: "Feedback",
		pages: []examplePage{
			{id: "#progress", label: "Progress", build: ProgressExamples},
		},
	},
}
