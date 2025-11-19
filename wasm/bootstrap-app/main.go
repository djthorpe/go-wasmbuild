package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application displays examples of Bootstrap components
func main() {
	app := mvc.New()
	router := mvc.Router(mvc.WithClass("container-fluid", "my-2"))
	var navItems []mvc.View
	for _, group := range exampleGroups {
		var items []mvc.View
		for _, page := range group.pages {
			router = router.Page(page.id, page.build())
			items = append(items, bs.NavItem(page.id, page.label))
		}
		if len(items) == 0 {
			continue
		}
		dropdownChildren := make([]any, len(items))
		for i, item := range items {
			dropdownChildren[i] = item
		}
		navItems = append(navItems, bs.NavDropdown(dropdownChildren...).Label(group.label))
	}

	navArgs := []any{
		bs.WithPosition(bs.Sticky | bs.Top),
		bs.WithTheme(bs.Dark),
		bs.WithSize(bs.Medium),
	}
	for _, item := range navItems {
		navArgs = append(navArgs, item)
	}
	navArgs = append(navArgs, bs.NavItem("https://github.com/djthorpe/go-wasmbuild", bs.Icon("github", mvc.WithClass("me-1")), "GitHub"))

	app.Append(
		bs.NavBar("main", navArgs...).Label(bs.Icon("bootstrap-fill")),
		router,
	)

	// Wait
	select {}
}
