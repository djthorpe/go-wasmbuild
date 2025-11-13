package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Alerts() mvc.View {
	return bs.Container(bs.WithSize(bs.SizeFluid)).Append(
		"Alerts ",
		bs.Alert().Append("Hello, World!"),
		bs.Alert(bs.WithColor(bs.INFO)).Append(
			bs.Icon("info-circle-fill"),
			" An example alert with an icon",
		),
		bs.Alert(bs.WithColor(bs.LIGHT)).Append("Light"),
		bs.Alert(bs.WithColor(bs.SUCCESS)).Append("Success"),
		bs.DismissableAlert(bs.WithColor(bs.DANGER)).Append(
			"Dismissable Alert",
			bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "alert")),
		),
	)
}
