package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Containers() mvc.View {
	return mvc.Div().Append(
		bs.Container(
			bs.WithBackground(bs.LIGHT),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All),
		).Append(
			"Light Container",
		),
		bs.Container(
			bs.WithBackground(bs.PRIMARY),
			bs.WithColor(bs.LIGHT),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All),
		).Append(
			"Primary Color Container",
		),
		bs.Container(
			bs.WithColor(bs.LIGHT),
			bs.WithBackground(bs.DARK),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
		).Append(
			"Dark Container",
		),
		bs.Container(
			bs.WithColor(bs.DARK),
			bs.WithBackground(bs.WHITE),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All, bs.DANGER_SUBTLE),
		).Append(
			"White Container with danger-subtle border",
		),
		bs.Container(
			bs.WithSize(bs.SizeLarge),
			bs.WithBackground(bs.LIGHT),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All),
		).Append(
			"Large Container",
		),
		bs.Container(
			bs.WithSize(bs.SizeSmall),
			bs.WithBackground(bs.LIGHT),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All),
		).Append(
			"Small Container",
		),
		bs.Container(
			bs.WithSize(bs.SizeFluid),
			bs.WithBackground(bs.LIGHT),
			bs.WithMargin(bs.All, 2),
			bs.WithPadding(bs.All, 2),
			bs.WithAlign(bs.Center),
			bs.WithBorder(bs.All),
		).Append(
			"Fluid Container",
		),
	)
}
