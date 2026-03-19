package main

import (
	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	button "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/button"
	content "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/content"
	form "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/form"
	navigation "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/navigation"
)

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
		Page("#checkbox", carbon.Page(form.CheckboxView()...), nav.(ItemSelector).Item("#checkbox")).
		Page("#dropdown", carbon.Page(form.DropdownView()...), nav.(ItemSelector).Item("#dropdown")).
		Page("#sidenav", carbon.Page(navigation.PanelView()...), nav.(ItemSelector).Item("#sidenav"))
}
