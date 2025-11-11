package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Links() mvc.View {
	return bs.Container().Append(
		bs.Heading(1, mvc.WithClass("mt-4", "mb-4")).Append("Bootstrap Link Examples"),
		bs.Link("#test1").Content("Left Aligned Link"),
		bs.Link("#test2").Content("Center Aligned Link"),
		bs.Link("#test3", bs.WithColor("danger")).Content("Right Aligned Link"),
	)
}
