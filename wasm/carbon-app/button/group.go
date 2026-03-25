package button

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func GroupView() []any {
	return []any{
		storybook.PageHeader("ButtonGroup", "ButtonGroup.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeG10),
			basicButtonGroupStory(),
			toolbarButtonGroupStory(),
		),
	}
}

func basicButtonGroupStory() dom.Element {
	btn1 := carbon.Button(carbon.With(carbon.KindPrimary), "Primary").SetValue("btn-primary")
	btn2 := carbon.Button(carbon.With(carbon.KindSecondary), "Secondary").SetValue("btn-secondary")
	btn3 := carbon.Button(carbon.With(carbon.KindTertiary), "Tertiary").SetValue("btn-tertiary")
	grp := carbon.ButtonGroup()
	grp.Content(btn1, btn2, btn3)
	canvas := carbon.Section(mvc.WithClass("canvas"), grp)
	enabled1 := carbon.Checkbox("Primary")
	enabled1.SetActive(true)
	enabled2 := carbon.Checkbox("Secondary")
	enabled2.SetActive(true)
	enabled3 := carbon.Checkbox("Tertiary")
	enabled3.SetActive(true)
	enabledControl := carbon.CheckboxGroup("").SetLabel("Enabled")
	syncEnabled := func() {
		views := make([]mvc.View, 0, 3)
		if enabled1.Active() {
			views = append(views, btn1)
		}
		if enabled2.Active() {
			views = append(views, btn2)
		}
		if enabled3.Active() {
			views = append(views, btn3)
		}
		grp.SetEnabled(views...)
	}
	enabled1.AddEventListener(carbon.EventChange, func(dom.Event) { syncEnabled() })
	enabled2.AddEventListener(carbon.EventChange, func(dom.Event) { syncEnabled() })
	enabled3.AddEventListener(carbon.EventChange, func(dom.Event) { syncEnabled() })
	enabledControl.Content(enabled1, enabled2, enabled3)

	return storybook.Story(
		"Basic Button Group",
		"A button group arranges multiple buttons horizontally with correct Carbon spacing.",
		canvas,
		grp,
		storybook.Dropdown("Theme", carbon.ThemeG10, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(a carbon.Attr) {
			btn1.Apply(carbon.With(a)...)
			btn2.Apply(carbon.With(a)...)
			btn3.Apply(carbon.With(a)...)
		}),
		enabledControl,
	)
}

func toolbarButtonGroupStory() dom.Element {
	newBtn := carbon.Button(carbon.Icon(carbon.IconAdd, carbon.With(carbon.IconSize20)))
	newBtn.SetValue("toolbar-new")
	newBtn.SetLabel("New")
	searchBtn := carbon.Button(carbon.Icon(carbon.IconSearch, carbon.With(carbon.IconSize20)))
	searchBtn.SetValue("toolbar-search")
	searchBtn.SetLabel("Search")
	settingsBtn := carbon.Button(carbon.Icon(carbon.IconSettings, carbon.With(carbon.IconSize20)))
	settingsBtn.SetValue("toolbar-settings")
	settingsBtn.SetLabel("Settings")
	profileBtn := carbon.Button(carbon.Icon(carbon.IconUserAvatar, carbon.With(carbon.IconSize20)))
	profileBtn.SetValue("toolbar-profile")
	profileBtn.SetLabel("Profile")

	leftGroup := carbon.ButtonGroup()
	leftGroup.Content(newBtn, searchBtn)
	rightGroup := carbon.ButtonGroup()
	rightGroup.Content(settingsBtn, profileBtn)

	separator := mvc.HTML("DIV",
		mvc.WithAttr("aria-hidden", "true"),
		mvc.WithStyle("width:1px;height:2rem;background:var(--cds-border-subtle,#c6c6c6);margin:0 0.5rem"),
	)
	toolbar := mvc.HTML("DIV",
		mvc.WithStyle("display:flex;align-items:center;gap:0.25rem;padding:0.5rem;background:var(--cds-layer-01,#f4f4f4);border-radius:0.25rem;width:max-content"),
		leftGroup,
		separator,
		rightGroup,
	)
	canvas := carbon.Section(mvc.WithClass("canvas"), toolbar)
	applySize := func(size carbon.Attr) {
		newBtn.Apply(carbon.With(size)...)
		searchBtn.Apply(carbon.With(size)...)
		settingsBtn.Apply(carbon.With(size)...)
		profileBtn.Apply(carbon.With(size)...)
	}

	return storybook.Story(
		"Toolbar Button Group",
		"Multiple button groups can be composed into a toolbar, with a visual separator between groups of icon-only actions.",
		canvas,
		leftGroup,
		storybook.Dropdown("Theme", carbon.ThemeG10, storybook.DefaultThemes, func(theme carbon.Attr) {
			canvas.Apply(carbon.With(theme)...)
		}),
		storybook.Dropdown("Size", carbon.SizeLarge, buttonSizes, func(size carbon.Attr) {
			applySize(size)
		}),
	)
}
