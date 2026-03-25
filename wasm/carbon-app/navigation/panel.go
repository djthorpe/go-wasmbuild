package navigation

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var panelSizes = []carbon.Attr{"16rem", "24rem", "32rem"}

func PanelView() []any {
	return []any{
		storybook.PageHeader("Panels", "HeaderPanel.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicPanelStory(),
		),
	}
}

func basicPanelStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentSize := panelSizes[0] // 16rem — Carbon default
	visible := true

	// Forward-declare refresh so the close button handler can call it.
	var refresh func()

	// CloseButton automatically calls SetVisible(false) on the nearest parent
	// VisibleState. We also add a second click handler to sync local state.
	closeBtn := carbon.CloseButton(mvc.WithStyle("position:absolute;top:0;right:0"))
	closeBtn.AddEventListener(carbon.EventClick, func(dom.Event) {
		visible = false
		refresh()
	})

	panel := carbon.HeaderPanel(
		mvc.HTML("DIV", mvc.WithStyle("position:relative;padding:1rem 1.5rem"),
			carbon.Head(4, "Panel"),
			closeBtn,
			carbon.Para("The header panel slides in from the right side of the UI shell. Use it for supplementary navigation, account settings, or any contextual content that does not belong in the main content area."),
		),
	)
	panel.SetVisible(visible)

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;position:relative;min-height:14rem;overflow:hidden"),
		mvc.HTML("DIV", mvc.WithStyle("padding:1rem 1.5rem"),
			carbon.Head(4, "Page content"),
			carbon.Para("This content sits in the main area. When the panel is open it overlays from the right — the content underneath is not pushed or reflowed."),
		),
		panel,
	)

	refresh = func() {
		canvas.Apply(carbon.With(currentTheme)...)
		// Only override inline-size when expanded — Carbon uses inline-size:0 to
		// collapse the panel, so a permanent inline override would prevent hiding.
		style := "position:absolute;top:0;right:0;height:100%"
		if visible {
			style += ";inline-size:" + string(currentSize)
		}
		panel.Root().SetAttribute("style", style)
		panel.SetVisible(visible)
	}
	refresh()

	return storybook.Story(
		"Basic Panel",
		"The header panel is a right-side overlay anchored to the UI shell header. It is shown or hidden by toggling the expanded property — typically driven by a button in the header global actions bar. Carbon's default width is 16rem; larger sizes can be applied via an inline style override.",
		canvas,
		nil,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			refresh()
		}),
		storybook.Dropdown("Size", currentSize, panelSizes, func(size carbon.Attr) {
			currentSize = size
			refresh()
		}),
		storybook.CheckboxGroup("Visibility", "Show panel", visible, func(v bool) {
			visible = v
			refresh()
		}),
	)
}
