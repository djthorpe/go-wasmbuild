package navigation

import (
	"time"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

var sideNavChoices = []carbon.Attr{"Overview", "Reports", "Stations", "Vehicles", "Maintenance", "Settings"}

func SideNavView() []any {
	return []any{
		storybook.PageHeader("SideNav", "SideNav.md"),
		carbon.Section(
			mvc.WithStyle("padding:1.5rem 2rem"),
			carbon.With(carbon.ThemeWhite),
			basicSideNavStory(),
		),
	}
}

func basicSideNavStory() dom.Element {
	currentTheme := carbon.ThemeWhite
	currentActive := sideNavChoices[1]

	overview := carbon.SideNavLink("#overview", "Overview")
	reports := carbon.SideNavLink("#reports", "Reports")
	stations := carbon.SideNavGroupItem("#stations", "Stations")
	vehicles := carbon.SideNavGroupItem("#vehicles", "Vehicles")
	maintenance := carbon.SideNavGroupItem("#maintenance", "Maintenance")
	fleet := carbon.SideNavGroup("Fleet", stations, vehicles, maintenance)
	settings := carbon.SideNavLink("#settings", "Settings")

	side := carbon.SideNav(
		mvc.WithAttr("is-not-child-of-header", ""),
		mvc.WithAttr("expanded", ""),
		overview,
		reports,
		fleet,
		settings,
	)
	localizeSideNav(side.Root())

	preview := carbon.Section(
		mvc.WithClass("canvas"),
		mvc.WithStyle("display:grid;gap:0;border:1px solid var(--cds-border-subtle,#c6c6c6);overflow:hidden"),
		mvc.HTML("DIV",
			mvc.WithStyle("display:flex;min-height:24rem"),
			mvc.HTML("DIV", mvc.WithStyle("flex:0 0 16rem;min-width:16rem;background:var(--cds-background,#ffffff)"), side),
			mvc.HTML("DIV",
				mvc.WithStyle("flex:1;padding:2rem 2rem 2rem 1.5rem;color:var(--cds-text-primary,#161616);background:linear-gradient(180deg,var(--cds-layer-accent-01,#f4f4f4) 0%,var(--cds-background,#ffffff) 100%)"),
				carbon.Head(3, "Side nav preview"),
				carbon.Para("Top-level links and grouped items participate in the same ActiveGroup selection model, which keeps router-style activation straightforward."),
			),
		),
	)

	setActive := func(choice carbon.Attr) {
		currentActive = choice
		switch choice {
		case "Overview":
			side.SetActive(overview)
		case "Stations":
			fleet.Root().SetAttribute("expanded", "")
			side.SetActive(stations)
		case "Vehicles":
			fleet.Root().SetAttribute("expanded", "")
			side.SetActive(vehicles)
		case "Maintenance":
			fleet.Root().SetAttribute("expanded", "")
			side.SetActive(maintenance)
		case "Settings":
			side.SetActive(settings)
		default:
			side.SetActive(reports)
		}
	}
	render := func() {
		preview.Apply(carbon.With(currentTheme)...)
		setActive(currentActive)
	}
	render()

	return storybook.Story(
		"Side Navigation",
		"The side-nav wrapper coordinates top-level links and nested menu items through the shared ActiveGroup API. This preview keeps the component in its expanded shell mode so the active-state and section-toggle behavior are visible.",
		preview,
		side,
		storybook.Dropdown("Theme", currentTheme, storybook.DefaultThemes, func(theme carbon.Attr) {
			currentTheme = theme
			render()
		}),
		storybook.Dropdown("Active item", currentActive, sideNavChoices, func(choice carbon.Attr) {
			setActive(choice)
		}),
	)
}

func localizeSideNav(element dom.Element) {
	if element == nil {
		return
	}
	element.SetAttribute("style", "display:block;position:relative;inset:auto;inline-size:16rem;max-inline-size:16rem;block-size:100%;z-index:1")
	localizeSideNavShadow(element, 0)
}

func localizeSideNavShadow(element dom.Element, attempt int) {
	value, ok := element.JSValue().(js.Value)
	if !ok || value.IsUndefined() || value.IsNull() {
		return
	}
	shadowRoot := value.Get("shadowRoot")
	if shadowRoot.IsUndefined() || shadowRoot.IsNull() {
		if attempt < 6 {
			js.SetTimeout(20*time.Millisecond, func() {
				localizeSideNavShadow(element, attempt+1)
			})
		}
		return
	}
	nav := shadowRoot.Call("querySelector", ".cds--side-nav")
	if nav.IsUndefined() || nav.IsNull() {
		if attempt < 6 {
			js.SetTimeout(20*time.Millisecond, func() {
				localizeSideNavShadow(element, attempt+1)
			})
		}
		return
	}
	style := nav.Get("style")
	if style.IsUndefined() || style.IsNull() {
		return
	}
	style.Call("setProperty", "position", "relative")
	style.Call("setProperty", "inset", "auto")
	style.Call("setProperty", "inset-block-start", "auto")
	style.Call("setProperty", "inset-inline-start", "auto")
	style.Call("setProperty", "inline-size", "16rem")
	style.Call("setProperty", "max-inline-size", "16rem")
	style.Call("setProperty", "block-size", "100%")
	style.Call("setProperty", "min-block-size", "100%")
	style.Call("setProperty", "transform", "none")
	style.Call("setProperty", "transition", "none")
	style.Call("setProperty", "will-change", "auto")
	style.Call("setProperty", "z-index", "1")
	overlay := shadowRoot.Call("querySelector", ".cds--side-nav__overlay")
	if !overlay.IsUndefined() && !overlay.IsNull() {
		overlayStyle := overlay.Get("style")
		if !overlayStyle.IsUndefined() && !overlayStyle.IsNull() {
			overlayStyle.Call("setProperty", "display", "none")
			overlayStyle.Call("setProperty", "position", "absolute")
			overlayStyle.Call("setProperty", "inset", "auto")
			overlayStyle.Call("setProperty", "block-size", "0")
			overlayStyle.Call("setProperty", "inline-size", "0")
		}
	}
}
