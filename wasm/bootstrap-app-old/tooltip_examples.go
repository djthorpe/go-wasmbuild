package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TooltipExamples() mvc.View {
	return bs.Container(
		bs.Button("Hover me to see tooltip", bs.WithTooltip("Tooltip text example")),
		bs.HRule(),
	)
}
