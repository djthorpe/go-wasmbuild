package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func BadgeExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("badge_examples.md"),
		bs.HRule(),
	)
}
