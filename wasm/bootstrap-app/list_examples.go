package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ListExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("list_examples.md"),
		bs.HRule(),
	)
}
