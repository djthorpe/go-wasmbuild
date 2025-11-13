package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func NavBarExamples() mvc.View {
	return bs.Container(
		bs.Heading(3, "Navigation Bar"),
		bs.NavBar("hello"),
	)
}
