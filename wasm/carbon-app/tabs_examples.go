package main

import (
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TabExamples() mvc.View {
	return ExamplePage("Tabs",
		cds.LeadPara(
			`Tabs use the `, cds.Code("cds-tabs"), ` and `, cds.Code("cds-tab"),
			` web components. Build a complete tab group with `, cds.Code("cds.TabSet()"),
			` and `, cds.Code("cds.TabPane()"), `. `,
			`Pass `, cds.Code(".Selected()"), ` to pre-select a pane, `,
			cds.Code(".Disabled()"), ` to disable one, and `,
			cds.Code("cds.WithTabsType(cds.TabsContained)"), ` for the box style.`,
		),
		ExampleRow("Basic", Example_Tabs_001, "Four panes; the second is selected by default."),
		ExampleRow("Contained", Example_Tabs_002, "Box-style tabs using cds.WithTabsType(cds.TabsContained)."),
		ExampleRow("Disabled tab", Example_Tabs_003, "The third tab is disabled and cannot be selected."),
		ExampleRow("Interactive", Example_Tabs_004, "Listen for cds-tabs-selected; read the new value from the event target."),
	)
}

const tabsStyle = "padding-bottom:var(--cds-spacing-07,2rem);"

func tabPanel(title, body string, bullets ...string) mvc.View {
	items := make([]any, 0, len(bullets))
	for _, bullet := range bullets {
		items = append(items, mvc.HTML("li", bullet))
	}
	content := []any{
		cds.Heading(5, title),
		cds.Para(body),
	}
	if len(items) > 0 {
		content = append(content, mvc.HTML("ul", items...))
	}
	return cds.Section(content...)
}

func Example_Tabs_001() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", tabsStyle),
		cds.TabSet(
			cds.TabPane("Overview",
				tabPanel("Overview", "This overview pane summarizes the structure of the component and its most common usage.",
					"Use TabSet for simple examples",
					"Use TabPane to define each labeled section",
				),
			),
			cds.TabPane("Details",
				tabPanel("Details", "This details pane exposes the default selected state and shows how individual panels can carry richer content.",
					"The second tab starts selected",
					"Each pane can contain headings, text, and lists",
				),
			).Selected(),
			cds.TabPane("Usage",
				tabPanel("Usage", "This usage pane is intentionally different so the content swap is obvious when the selected tab changes.",
					"Switch tabs to compare panel bodies",
					"Panel visibility is handled in TabSet",
				),
			),
			cds.TabPane("References",
				tabPanel("References", "This references pane could hold linked docs, related examples, or API guidance.",
					"Related: buttons, notifications, accordion",
					"API: TabsContained, TabPane.Selected, TabPane.Disabled",
				),
			),
		),
	), sourcecode()
}

func Example_Tabs_002() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", tabsStyle),
		cds.TabSet(
			cds.WithTabsType(cds.TabsContained),
			cds.TabPane("Overview", tabPanel("Overview", "Contained tabs are visually heavier and work better when the content feels sectioned and boxed.")),
			cds.TabPane("Details", tabPanel("Details", "This pane starts active and demonstrates the contained treatment with a different body copy.")).Selected(),
			cds.TabPane("Usage", tabPanel("Usage", "Use the contained variant when tabs should read as discrete controls rather than a thin navigation strip.")),
			cds.TabPane("References", tabPanel("References", "This final pane rounds out the example with distinct content for the fourth tab.")),
		),
	), sourcecode()
}

func Example_Tabs_003() (mvc.View, string) {
	return cds.Section(
		mvc.WithAttr("style", tabsStyle),
		cds.TabSet(
			cds.TabPane("Active", tabPanel("Active", "This first tab is available and shows standard content.")),
			cds.TabPane("Also active", tabPanel("Also active", "This second tab is also available and should switch normally.")),
			cds.TabPane("Disabled", tabPanel("Disabled", "You cannot select this tab; it remains unavailable to pointer and keyboard interaction.")).Disabled(),
			cds.TabPane("Last", tabPanel("Last", "This last pane confirms that tabs after a disabled one still work correctly.")),
		),
	), sourcecode()
}

func Example_Tabs_004() (mvc.View, string) {
	status := cds.Para(
		mvc.WithAttr("style", "margin-top:var(--cds-spacing-05,1rem);color:var(--cds-text-secondary,#525252);"),
		"Select a tab…",
	)
	tabSet := cds.TabSet(
		cds.TabPane("Alpha", tabPanel("Alpha", "Alpha focuses on the first step in the flow.", "Initialize the state", "Render the first section")),
		cds.TabPane("Beta", tabPanel("Beta", "Beta shifts to a second phase with different supporting points.", "Load dependent data", "Apply intermediate rules")),
		cds.TabPane("Gamma", tabPanel("Gamma", "Gamma represents a later stage and intentionally uses different copy.", "Validate the current selection", "Prepare the final output")),
		cds.TabPane("Delta", tabPanel("Delta", "Delta is the wrap-up state for the interactive example.", "Review the result", "Persist the final change")),
	)
	// cds-tabs-selected fires on the <cds-tabs> element after selection.
	// el.Value() reads the JS .value property, which is set synchronously to
	// the selected tab's value string before the event is dispatched.
	tabSet.AddEventListener("cds-tabs-selected", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			selected := el.Value()
			status.Content("Selected tab: ", cds.Strong(selected))
		}
	})
	return cds.Section(
		mvc.WithAttr("style", tabsStyle),
		cds.Section(tabSet, status),
	), sourcecode()
}
