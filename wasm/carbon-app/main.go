package main

import (
	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	buttons "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/buttons"
	content "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/content"
)

func main() {
	// Side bar navigation
	SideNav := carbon.SideNav(
		carbon.SideNavGroup("Content",
			carbon.SideNavGroupItem("#heading", "Headings"),
			carbon.SideNavGroupItem("#text", "Text"),
			carbon.SideNavGroupItem("#markdown", "Markdown"),
			carbon.SideNavGroupItem("#icon", "Icons"),
		),
		carbon.SideNavGroup("Components",
			carbon.SideNavGroupItem("#button", "Buttons"),
			carbon.SideNavGroupItem("#button-group", "Button Groups"),
		),
		carbon.SideNavGroup("Forms",
			carbon.SideNavGroupItem("#dropdown", "Dropdowns"),
		),
		carbon.SideNavGroup("Data",
			carbon.SideNavGroupItem("#table", "Tables"),
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
				carbon.Button(carbon.Icon(carbon.IconUserAvatar, carbon.With(carbon.IconSize24))),
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
		Page("#button", carbon.Page(buttons.View()...), nav.(ItemSelector).Item("#button")).
		Page("#button-group", carbon.Page(buttons.GroupView()...), nav.(ItemSelector).Item("#button-group"))
}
