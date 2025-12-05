package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func AlertExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("alert_examples.md"),
		ExampleCard("Alerts with Color", Example_Alerts_001),
		ExampleCard("Dismissable Alert", Example_Alerts_002),
	)
}

func Example_Alerts_001() (mvc.View, string) {
	return bs.Container(
		bs.Alert("This is a primary alert", bs.WithColor(bs.Primary)),
		bs.Alert("This is a secondary alert", bs.WithColor(bs.Secondary)),
		bs.Alert("This is a success alert", bs.WithColor(bs.Success)),
		bs.Alert("This is a danger alert", bs.WithColor(bs.Danger)),
		bs.Alert("This is a warning alert", bs.WithColor(bs.Warning)),
		bs.Alert("This is an info alert", bs.WithColor(bs.Info)),
		bs.Alert("This is a light alert", bs.WithColor(bs.Light)),
		bs.Alert("This is a dark alert", bs.WithColor(bs.Dark)),
	), sourcecode()
}

func Example_Alerts_002() (mvc.View, string) {
	return bs.Container(
		bs.Alert("Dismissing an Alert removes it from the DOM", bs.WithColor(bs.Primary), bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "alert"))),
	), sourcecode()
}
