package content

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

const (
	structuredListHang       = carbon.Attr("Hang")
	structuredListFlush      = carbon.Attr("Flush")
	structuredListDefault    = carbon.Attr("Default")
	structuredListCondensed  = carbon.Attr("Condensed")
	structuredListStarter    = carbon.Attr("Starter")
	structuredListTeam       = carbon.Attr("Team")
	structuredListEnterprise = carbon.Attr("Enterprise")
)

func StructuredListView() []any {
	return []any{
		storybook.PageHeader("StructuredList", "StructuredList.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			structuredListDefaultStory(),
			structuredListSelectableStory(),
		),
	}
}

func structuredListDefaultStory() dom.Element {
	list := carbon.StructuredList(
		carbon.StructuredListHeader("Plan", "Support", "Notes"),
		carbon.StructuredListRow("Starter", "Community", "Good for quick evaluation and early prototypes."),
		carbon.StructuredListRow("Team", "Business hours", "Balanced default for internal tools and line-of-business apps."),
		carbon.StructuredListRow("Enterprise", "24/7", "Best when rollout, governance, and uptime commitments matter."),
	)
	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		list,
	)

	applyAlignment := func(alignment carbon.Attr) {
		switch alignment {
		case structuredListFlush:
			list.Apply(carbon.With(carbon.StructuredListFlush)...)
		default:
			list.Apply(mvc.WithoutAttr("flush"))
		}
	}

	applySize := func(size carbon.Attr) {
		switch size {
		case structuredListCondensed:
			list.Apply(carbon.With(carbon.StructuredListCondensed)...)
		default:
			list.Apply(mvc.WithoutAttr("condensed"))
		}
	}

	return storybook.Story(
		"Default Structured List",
		"The default variant is a lighter-weight alternative to a data table when the content is mostly read-only and comparison is simple.",
		canvas,
		nil,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Alignment", structuredListHang, []carbon.Attr{structuredListHang, structuredListFlush}, applyAlignment),
		storybook.Dropdown("Size", structuredListDefault, []carbon.Attr{structuredListDefault, structuredListCondensed}, applySize),
	)
}

func structuredListSelectableStory() dom.Element {
	starter := carbon.StructuredListRow("Starter", "Community", "A lighter option for small, self-serve teams.")
	starter.SetValue("starter")
	team := carbon.StructuredListRow("Team", "Business hours", "Adds support coverage without changing the layout model.")
	team.SetValue("team")
	enterprise := carbon.StructuredListRow("Enterprise", "24/7", "Use selection when the list represents a single choice.")
	enterprise.SetValue("enterprise")

	list := carbon.StructuredList(
		mvc.WithAttr("selection-name", "structured-list-plan"),
		carbon.StructuredListHeader("Plan", "Support", "Notes"),
		starter,
		team,
		enterprise,
	)
	list.SetActive(team)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;width:100%;padding:1rem"),
		carbon.Compact("Selectable structured lists are appropriate when one row represents a single choice, such as a pricing plan or environment target."),
		list,
	)

	applySize := func(size carbon.Attr) {
		switch size {
		case structuredListCondensed:
			list.Apply(carbon.With(carbon.StructuredListCondensed)...)
		default:
			list.Apply(mvc.WithoutAttr("condensed"))
		}
	}

	applySelected := func(selected carbon.Attr) {
		switch selected {
		case structuredListStarter:
			list.SetActive(starter)
		case structuredListEnterprise:
			list.SetActive(enterprise)
		default:
			list.SetActive(team)
		}
	}

	selectedRowControl := storybook.Dropdown("Selected row", structuredListTeam, []carbon.Attr{structuredListStarter, structuredListTeam, structuredListEnterprise}, applySelected)
	syncSelectedRowControl := func(selected carbon.Attr) {
		if dropdown, ok := selectedRowControl.(mvc.ActiveGroup); ok {
			for _, child := range selectedRowControl.Children() {
				if item, ok := child.(mvc.ValueState); ok && item.Value() == string(selected) {
					dropdown.SetActive(child)
					break
				}
			}
		}
		if dropdown, ok := selectedRowControl.(mvc.ValueState); ok {
			dropdown.SetValue(string(selected))
		}
	}
	list.AddEventListener(carbon.EventChange, func(e dom.Event) {
		row := mvc.ViewFromEvent(e, carbon.ViewStructuredListRow)
		if row == nil {
			return
		}
		selected, ok := row.(mvc.ValueState)
		if !ok {
			return
		}
		syncSelectedRowControl(carbon.Attr(selected.Value()))
	})

	return storybook.Story(
		"Selectable Structured List",
		"The selectable variant turns rows into a single-choice set. Carbon treats it as a presentation and selection pattern, not as a full data table.",
		canvas,
		list,
		storybook.Dropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", structuredListDefault, []carbon.Attr{structuredListDefault, structuredListCondensed}, applySize),
		selectedRowControl,
	)
}
