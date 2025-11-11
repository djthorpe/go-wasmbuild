package main

import (
	"fmt"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// Application is an example of using a table view and controller
func main() {
	// Set the content of the application
	app := mvc.New().Content(
		bs.Link("#badge", mvc.WithClass("m-2")).Content("Badge Examples"),
		bs.Link("#link", mvc.WithClass("m-2")).Content("Link Examples"),
		mvc.Router().Page(
			"#badge", BadgeExamples(),
		).Page(
			"#link", LinkExamples(),
		),
	)

	// Print the application
	fmt.Println(app)

	// Wait
	select {}
}

func BadgeExamples() mvc.View {
	return bs.Container(bs.WithSize(bs.SizeSmall), mvc.WithClass("m-2")).Content(
		bs.Heading(1).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(2).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(3).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(4).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(5).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(6).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
	)
}

func LinkExamples() mvc.View {
	return bs.Container(bs.WithSize(bs.SizeSmall), mvc.WithClass("m-2")).Content(
		bs.Heading(1).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(2).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(3).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(4).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(5).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
		bs.Heading(6).Content(
			"Example heading ",
			bs.Badge().Content("New"),
		),
	)
}
