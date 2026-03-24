package button

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

func CloseButtonView() []any {
	return []any{
		storybook.PageHeader("CloseButton", "CloseButton.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeG10),
			closeButtonStory(),
		),
	}
}

func closeButtonStory() dom.Element {
	panel := carbon.HeaderPanel(
		carbon.With(carbon.ThemeG10),
		mvc.WithStyle("position:relative;display:block;inline-size:22rem;padding:1rem;border:1px solid var(--cds-border-subtle,#c6c6c6);background:var(--cds-layer,#fff)"),
		carbon.CloseButton(mvc.WithStyle("position:absolute;top:0;right:0")),
		carbon.Head(3, "Dismissible Panel"),
		carbon.Para("CloseButton finds the nearest mvc.VisibleState ancestor and hides it when clicked."),
	)
	panel.SetVisible(true)

	reset := carbon.Button(carbon.With(carbon.KindTertiary), "Reset panel")
	reset.AddEventListener(carbon.EventClick, func(dom.Event) {
		panel.SetVisible(true)
	})

	canvas := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:1rem;justify-items:start"),
		panel,
		reset,
	)

	return storybook.Story(
		"CloseButton",
		"CloseButton is an icon-only dismiss control that hides the nearest parent view implementing mvc.VisibleState.",
		canvas,
		reset,
	)
}
