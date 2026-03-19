package main

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	button "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/button"
	content "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/content"
	form "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/form"
	navigation "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/navigation"
	storybook "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/storybook"
)

type codePanelController struct {
	panel   mvc.View
	visible bool
}

func (c *codePanelController) Visible() bool {
	return c.visible
}

func (c *codePanelController) SetVisible(visible bool) mvc.View {
	c.visible = visible
	style := "position:fixed;top:3rem;right:0;height:calc(100vh - 3rem);overflow:auto;z-index:9000"
	if visible {
		style += ";inline-size:36rem;background:var(--cds-layer,#fff);border-inline-start:1px solid var(--cds-border-subtle,#e0e0e0)"
	}
	c.panel.Root().SetAttribute("style", style)
	if panel, ok := c.panel.(mvc.VisibleState); ok {
		panel.SetVisible(visible)
	}
	return c.panel
}

func main() {
	// Side bar navigation
	SideNav := carbon.SideNav(
		carbon.SideNavGroup("Content",
			carbon.SideNavGroupItem("#heading", "Headings"),
			carbon.SideNavGroupItem("#text", "Text"),
			carbon.SideNavGroupItem("#markdown", "Markdown"),
			carbon.SideNavGroupItem("#code", "Source Code"),
			carbon.SideNavGroupItem("#icon", "Icons"),
			carbon.SideNavGroupItem("#grid", "Grids"),
		),
		carbon.SideNavGroup("Components",
			carbon.SideNavGroupItem("#button", "Buttons"),
			carbon.SideNavGroupItem("#button-group", "Button Groups"),
			carbon.SideNavGroupItem("#tile", "Tiles"),
			carbon.SideNavGroupItem("#tag", "Tags"),
		),
		carbon.SideNavGroup("Forms",
			carbon.SideNavGroupItem("#form-introduction", "Introduction"),
			carbon.SideNavGroupItem("#input", "Input"),
			carbon.SideNavGroupItem("#checkbox", "Checkboxes"),
			carbon.SideNavGroupItem("#dropdown", "Dropdowns"),
		),
		carbon.SideNavGroup("Data",
			carbon.SideNavGroupItem("#table", "Tables"),
			carbon.SideNavGroupItem("#pagination", "Pagination"),
		),
		carbon.SideNavGroup("Navigation",
			carbon.SideNavGroupItem("#header", "Headers"),
			carbon.SideNavGroupItem("#sidenav", "Panels"),
			carbon.SideNavGroupItem("#breadcrumb", "Breadcrumbs"),
			carbon.SideNavGroupItem("#tabs", "Tabs"),
		),
		carbon.SideNavGroup("Feedback",
			carbon.SideNavGroupItem("#modal", "Modals"),
			carbon.SideNavGroupItem("#notification", "Notifications"),
			carbon.SideNavGroupItem("#toast", "Toasts"),
			carbon.SideNavGroupItem("#progress-bar", "Progress Bars"),
		),
	)

	codeTitle := carbon.Head(4, "Select a story")
	codeBlock := carbon.CodeBlock("// Select a story's View code link to inspect its example.", carbon.With(carbon.ThemeG10))
	codeBlock.SetWrapText(true)
	closeBtn := carbon.CloseButton(mvc.WithStyle("position:absolute;top:0;right:0"))
	codePanel := carbon.HeaderPanel(
		mvc.HTML("DIV", mvc.WithStyle("position:relative;padding:1rem 1.5rem;display:grid;gap:1rem"),
			carbon.Head(3, "Code Example"),
			closeBtn,
			codeTitle,
			codeBlock,
		),
	)
	codePanelState := &codePanelController{panel: codePanel}
	closeBtn.AddEventListener(carbon.EventClick, func(dom.Event) {
		codePanelState.SetVisible(false)
	})
	codePanelState.SetVisible(false)
	storybook.SetCodePanel(codePanelState, codeTitle, codeBlock)

	// Main content area — min-height fills the viewport below the fixed header.
	Content := carbon.Section(
		mvc.WithStyle("min-height:100vh"),
		router(SideNav),
	)

	mvc.New(carbon.Section(
		carbon.Header(
			carbon.HeaderNavGlobal(
				carbon.Button(
					carbon.Icon(carbon.IconLogoGithub, carbon.With(carbon.IconSize24)),
					mvc.WithAttr("href", "https://github.com/djthorpe/go-wasmbuild"),
					mvc.WithAttr("target", "_blank"),
					mvc.WithAriaLabel("GitHub repository"),
					mvc.WithStyle("color:white;padding:0 0.75rem;display:flex;align-items:center"),
				),
			),
		).SetLabel("#", "Carbon Design System", "Storybook"),
		SideNav,
		Content,
		codePanel,
	), carbon.With(carbon.ThemeG90)).Run()
}

func router(nav mvc.View) mvc.View {
	type ItemSelector interface {
		Item(string) mvc.View
	}
	return mvc.Router().
		Active(nav.(mvc.ActiveGroup)).
		Page("#heading", carbon.Page(content.HeadingView()...), nav.(ItemSelector).Item("#heading")).
		Page("#text", carbon.Page(content.TextView()...), nav.(ItemSelector).Item("#text")).
		Page("#markdown", carbon.Page(content.MarkdownView()...), nav.(ItemSelector).Item("#markdown")).
		Page("#code", carbon.Page(content.CodeView()...), nav.(ItemSelector).Item("#code")).
		Page("#icon", carbon.Page(content.IconView()...), nav.(ItemSelector).Item("#icon")).
		Page("#grid", carbon.Page(content.GridView()...), nav.(ItemSelector).Item("#grid")).
		Page("#button", carbon.Page(button.View()...), nav.(ItemSelector).Item("#button")).
		Page("#button-group", carbon.Page(button.GroupView()...), nav.(ItemSelector).Item("#button-group")).
		Page("#tile", carbon.Page(content.TileView()...), nav.(ItemSelector).Item("#tile")).
		Page("#tag", carbon.Page(content.TagView()...), nav.(ItemSelector).Item("#tag")).
		Page("#form-introduction", carbon.Page(form.IntroductionView()...), nav.(ItemSelector).Item("#form-introduction")).
		Page("#input", carbon.Page(form.InputView()...), nav.(ItemSelector).Item("#input")).
		Page("#checkbox", carbon.Page(form.CheckboxView()...), nav.(ItemSelector).Item("#checkbox")).
		Page("#dropdown", carbon.Page(form.DropdownView()...), nav.(ItemSelector).Item("#dropdown")).
		Page("#sidenav", carbon.Page(navigation.PanelView()...), nav.(ItemSelector).Item("#sidenav"))
}
