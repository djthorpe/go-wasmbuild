package navigation

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var headerChoices = []carbon.Attr{"Home", "Products", "Docs", "Support"}

func HeaderView() []any {
	return []any{
		storybook.PageHeader("Header", "Header.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicHeaderStory(),
		),
	}
}

func basicHeaderStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentActive := headerChoices[1]

	home := carbon.HeaderNavItem("#home", "Home")
	products := carbon.HeaderNavItem("#products", "Products")
	docs := carbon.HeaderNavItem("#docs", "Docs")
	support := carbon.HeaderNavItem("#support", "Support")
	searchAction := carbon.Button(
		mvc.WithAriaLabel("Search"),
		mvc.WithAttr("tooltip-text", "Search"),
		mvc.WithAttr("title", "Search"),
		carbon.Icon(carbon.IconSearch, carbon.With(carbon.IconSize20)),
	).SetValue("search")
	profileAction := carbon.Button(
		mvc.WithAriaLabel("Profile"),
		mvc.WithAttr("tooltip-text", "Profile"),
		mvc.WithAttr("title", "Profile"),
		carbon.Icon(carbon.IconUserAvatar, carbon.With(carbon.IconSize20)),
	).SetValue("profile")
	globalActions := carbon.HeaderNavGlobal(searchAction, profileAction)

	header := carbon.Header(home, products, docs, support, globalActions).
		SetLabel("#", "Go Wasm Build", "Carbon")
	header.Apply(mvc.WithStyle("position:absolute;inset:0 0 auto 0;z-index:1"))

	preview := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:0;overflow:hidden;border:1px solid var(--cds-border-subtle,#c6c6c6);min-height:15rem"),
		mvc.HTML("DIV",
			mvc.WithStyle("position:relative;overflow:hidden;min-height:3rem;background:var(--cds-background,#ffffff)"),
			header,
		),
		mvc.HTML("DIV",
			mvc.WithStyle("padding:7rem 1.5rem 1.5rem;color:var(--cds-text-primary,#161616);background:var(--cds-layer-accent-01,#f4f4f4)"),
			carbon.Head(3, "Header preview"),
			carbon.Para("Use the active-item control to switch the selected header navigation entry while keeping the Carbon shell structure intact."),
		),
	)

	setActive := func(choice carbon.Attr) {
		currentActive = choice
		switch choice {
		case "Home":
			header.SetActive(home)
		case "Docs":
			header.SetActive(docs)
		case "Support":
			header.SetActive(support)
		default:
			header.SetActive(products)
		}
	}
	render := func() {
		preview.Apply(carbon.With(currentTheme)...)
		setActive(currentActive)
	}
	render()

	return storybook.Story(
		"Header Navigation",
		"The header story exercises the shell header, top-level nav items, and the right-aligned global actions wrapper. Active state stays at the component boundary through the shared nav-group helpers.",
		preview,
		header,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Active item", currentActive, headerChoices, func(choice carbon.Attr) {
			setActive(choice)
		}),
	)
}
