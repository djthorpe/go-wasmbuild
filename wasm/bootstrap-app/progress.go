package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ProgressExamples() mvc.View {
	return bs.Container(
		bs.Heading(3, "Simple Progress Bar"),
		bs.Progress().SetValue("100"),
		bs.HRule(),
	)
}
