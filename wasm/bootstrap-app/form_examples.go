package main

import (
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func FormExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("form_examples.md"),
		ExampleCard("Basic Form", Example_Form_001),
	)
}

func Example_Form_001() (mvc.View, string) {
	return bs.Form("form_001", bs.Card(
		bs.Input("username", bs.WithPlaceholder("Username or Email Address"), mvc.WithClass("mb-3")).Label("Username"),
		bs.SecureInput("password", bs.WithPlaceholder("Password"), mvc.WithClass("mb-3")).Label("Password"),
	).Footer(
		bs.WithPosition(bs.End),
		bs.Button(bs.WithSubmit(), "Submit"),
	)).AddEventListener("submit", func(evt dom.Event) {
		form := mvc.ViewFromEvent(evt, bs.ViewForm).(bs.DataView)
		fmt.Println("Form submitted", form.Value())
	}), sourcecode()
}
