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

type navChoice struct {
	label string
	view  mvc.View
}

type choiceControl interface {
	Root() dom.Element
	SetValue(string)
}

// Stories returns the navigation-related stories for the Carbon example app.
func Stories() []dom.Element {
	return []dom.Element{
		headerStory(),
		sideNavStory(),
	}
}

func headerStory() dom.Element {
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
		Label("#", "Go Wasm Build", "Carbon")
	header.Apply(mvc.WithAttr("style", "position:absolute;inset:0 0 auto 0;z-index:1"))
	header.SetActive(products)
	headerPreview := mvc.HTML("DIV",
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
		mvc.WithAttr("style", "position:relative;overflow:hidden;min-height:3rem;background:var(--cds-background,#ffffff)"),
		header,
	)

	canvas := mvc.HTML("DIV",
		mvc.WithClass("canvas"),
		mvc.HTML("DIV",
			mvc.WithAttr("style", "position:relative;overflow:hidden;border:1px solid var(--cds-border-subtle,#c6c6c6);min-height:15rem;background:var(--cds-background,#ffffff)"),
			headerPreview,
			mvc.HTML("DIV",
				mvc.WithAttr("style", "padding:7rem 1.5rem 1.5rem;color:var(--cds-text-primary,#161616)"),
				carbon.Head(3, "Header Preview").Root(),
				carbon.Para("Use the control to switch the active header nav item while keeping the Carbon shell structure intact.").Root(),
			),
		),
	)

	choices := []navChoice{
		{label: "Home", view: home},
		{label: "Products", view: products},
		{label: "Docs", view: docs},
		{label: "Support", view: support},
	}
	activeChoice := choiceDropdown("Active Item", "Products", choices, func(choice navChoice) {
		header.SetActive(choice.view)
	})
	bindChoiceControl(choices, activeChoice)

	return storybook.Story(
		"Header Navigation",
		"A Carbon header with product branding and top-level navigation items. The active state uses the new wrapper methods instead of raw attributes.",
		canvas,
		header,
		storybook.AttrDropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(headerPreview, a)
		}),
		activeChoice.Root(),
	)
}

func sideNavStory() dom.Element {
	overview := carbon.SideNavLink("#overview", "Overview")
	reports := carbon.SideNavLink("#reports", "Reports")
	stations := carbon.SideNavItem("#stations", "Stations")
	vehicles := carbon.SideNavItem("#vehicles", "Vehicles")
	maintenance := carbon.SideNavItem("#maintenance", "Maintenance")
	fleet := carbon.SideNavSection("Fleet", stations, vehicles, maintenance)
	settings := carbon.SideNavLink("#settings", "Settings")

	side := carbon.SideNav(
		mvc.WithAttr("is-not-child-of-header", ""),
		mvc.WithAttr("expanded", ""),
		overview,
		reports,
		fleet,
		settings,
	)
	side.Apply(mvc.WithAttr("style", "display:block;block-size:100%;inline-size:16rem;max-inline-size:16rem;z-index:1"))
	localizeSideNav(side.Root())
	side.SetActive(reports)

	canvas := mvc.HTML("DIV",
		mvc.WithClass("canvas", carbon.ClassForTheme(carbon.ThemeWhite)),
		mvc.HTML("DIV",
			mvc.WithAttr("style", "position:relative;overflow:hidden;border:1px solid var(--cds-border-subtle,#c6c6c6);min-height:24rem;background:var(--cds-background,#ffffff)"),
			mvc.HTML("DIV",
				mvc.WithAttr("style", "display:flex;min-height:24rem"),
				mvc.HTML("DIV", mvc.WithAttr("style", "flex:0 0 16rem;min-width:16rem;background:var(--cds-background,#ffffff)"), side),
				mvc.HTML("DIV",
					mvc.WithAttr("style", "flex:1;padding:2rem 2rem 2rem 1.5rem;color:var(--cds-text-primary,#161616);background:linear-gradient(180deg,var(--cds-layer-accent-01,#f4f4f4) 0%,var(--cds-background,#ffffff) 100%)"),
					carbon.Head(3, "Side Navigation Preview").Root(),
					carbon.Para("Use the controls to move the active state across top-level links and nested menu items.").Root(),
				),
			),
		),
	)

	choices := []navChoice{
		{label: "Overview", view: overview},
		{label: "Reports", view: reports},
		{label: "Stations", view: stations},
		{label: "Vehicles", view: vehicles},
		{label: "Maintenance", view: maintenance},
		{label: "Settings", view: settings},
	}
	activeChoice := choiceDropdown("Active Item", "Reports", choices, func(choice navChoice) {
		side.SetActive(choice.view)
	})
	bindChoiceControl(choices, activeChoice)

	return storybook.Story(
		"Side Navigation",
		"A responsive Carbon side nav with top-level links and a nested section. The wrapper's ActiveGroup support makes router-style selection straightforward.",
		canvas,
		side,
		storybook.AttrDropdown("Theme", carbon.ThemeWhite, storybook.DefaultThemes, func(a carbon.Attr) {
			applyTheme(canvas, a)
		}),
		activeChoice.Root(),
	)
}

func choiceDropdown(label, selected string, options []navChoice, onChange func(navChoice)) choiceControl {
	items := make([]any, 0, len(options)+1)
	items = append(items, carbon.DropdownTitleText(label))
	for _, option := range options {
		item := carbon.DropdownItem(mvc.WithAttr("value", option.label), option.label)
		if option.label == selected {
			item.SetSelected(true)
		}
		items = append(items, item)
	}

	dd := carbon.Dropdown(append([]any{
		mvc.WithAttr("style", "width:100%"),
		mvc.WithClass(carbon.ClassForTheme(carbon.ThemeWhite)),
	}, items...)...)
	dd.SetValue(selected)
	dd.AddEventListener(carbon.EventSelected, func(dom.Event) {
		for _, option := range options {
			if option.label == dd.Value() {
				onChange(option)
				return
			}
		}
	})
	return dd
}

func choiceLabelForView(options []navChoice, view mvc.View) string {
	if view == nil {
		return ""
	}
	root := view.Root()
	for _, option := range options {
		if option.view != nil && option.view.Root() == root {
			return option.label
		}
	}
	return ""
}

func bindChoiceControl(options []navChoice, control choiceControl) {
	for _, option := range options {
		if option.view == nil {
			continue
		}
		label := option.label
		option.view.AddEventListener(carbon.EventClick, func(dom.Event) {
			control.SetValue(label)
		})
	}
}

func applyTheme(canvas dom.Element, theme carbon.Attr) {
	cl := canvas.ClassList()
	for _, t := range storybook.DefaultThemes {
		cl.Remove(carbon.ClassForTheme(t))
	}
	cl.Add(carbon.ClassForTheme(theme))
	canvas.SetClassName(cl.Value())
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
