package main

import (
	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	buttons "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/buttons"
	headings "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/headings"
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
		),
		carbon.SideNavGroup("Forms",
			carbon.SideNavGroupItem("#dropdowns", "Dropdowns"),
		),
		carbon.SideNavGroup("Data",
			carbon.SideNavGroupItem("#tables", "Tables"),
		),
	)

	// Main content area
	Content := carbon.Section(router(SideNav), carbon.With(carbon.ThemeG10))

	mvc.New(carbon.Section(
		carbon.Header(
			carbon.HeaderNavGlobal(
				carbon.Button(carbon.Icon(carbon.IconUserAvatar, carbon.With(carbon.IconSize24))),
			),
		).Label("#", "Go Wasm Build", "Carbon"),
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
		Page("#headings", headings.View(), nav.(ItemSelector).Item("#headings")).
		Page("#buttons", buttons.View(), nav.(ItemSelector).Item("#buttons"))
}
