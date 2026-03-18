package main

import (
	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	buttons "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/buttons"
	headings "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/content"
)

func main() {
	// Side bar navigation
	SideNav := carbon.SideNav(
		carbon.SideNavGroup("Content",
			carbon.SideNavGroupItem("#headings", "Headings"),
			carbon.SideNavGroupItem("#text", "Text"),
			carbon.SideNavGroupItem("#icons", "Icons"),
		),
		carbon.SideNavGroup("Components",
			carbon.SideNavGroupItem("#buttons", "Buttons"),
			carbon.SideNavGroupItem("#button-groups", "Button Groups"),
		),
		carbon.SideNavGroup("Forms",
			carbon.SideNavGroupItem("#dropdowns", "Dropdowns"),
		),
		carbon.SideNavGroup("Data",
			carbon.SideNavGroupItem("#tables", "Tables"),
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
		).Label("#", "Carbon Design System", "Storybook"),
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
		Page("#headings", carbon.Page(headings.View()...), nav.(ItemSelector).Item("#headings")).
		Page("#text", carbon.Page(headings.TextView()...), nav.(ItemSelector).Item("#text")).
		Page("#buttons", carbon.Page(buttons.View()...), nav.(ItemSelector).Item("#buttons")).
		Page("#button-groups", carbon.Page(buttons.GroupView()...), nav.(ItemSelector).Item("#button-groups"))
}
