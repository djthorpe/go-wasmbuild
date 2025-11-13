package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
	// Create the bootstrap app
	bs.New().Insert(
		bs.Container(
			bs.WithBreakpoint(bs.BreakpointLarge),
			bs.WithMargin(bs.TOP|bs.BOTTOM, 5),
		).Insert(
			bs.Heading(1, bs.WithTextAlign(bs.CENTER)).Insert("Hello, World!"),
		),
	)

	// Run the application
	select {}
}
