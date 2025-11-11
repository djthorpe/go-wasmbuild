package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Rules() mvc.View {
	return bs.Container(bs.WithMargin(bs.All, 5)).Append(
		bs.Heading(3).Append("Rules"),
		bs.Rule(),
		bs.VerticalRule(mvc.WithStyle("height:50px")),
	)
}
