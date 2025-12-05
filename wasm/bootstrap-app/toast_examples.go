package main

import (
	// Packages

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ToastExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("toast_examples.md"),
		ExampleCard("Basic Toast", Example_Toast_001),
		ExampleCard("With Header", Example_Toast_002),
		ExampleCard("With Color", Example_Toast_003),
		ExampleCard("Toast Group", Example_Toast_004),
	)
}

func Example_Toast_001() (mvc.View, string) {
	toast := bs.Toast(
		mvc.WithClass("my-3"),
		"Hello, world! This is a toast message.",
	)
	return bs.Container(
		bs.Button("Show Toast").AddEventListener("click", func(e dom.Event) {
			toast.Show()
		}),
		toast,
	), sourcecode()
}

func Example_Toast_002() (mvc.View, string) {
	toast := bs.Toast(
		mvc.WithClass("my-3"),
		"Hello, world! This is a toast message.",
	).Header(
		bs.Icon("alarm", mvc.WithClass("me-2")),
		"Notification",
		bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "toast"), mvc.WithClass("ms-auto")),
	)
	return bs.Container(
		bs.Button("Show Toast").AddEventListener("click", func(e dom.Event) {
			toast.Show()
		}),
		toast,
	), sourcecode()
}

func Example_Toast_003() (mvc.View, string) {
	toast := bs.Toast(
		mvc.WithClass("my-3"), bs.WithColor(bs.Dark),
		"Hello, world! This is a toast message.",
	).Header(
		bs.Icon("alarm", mvc.WithClass("me-2")),
		"Notification",
		bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "toast"), mvc.WithClass("ms-auto")),
	)
	return bs.Container(
		bs.Button("Show Toast").AddEventListener("click", func(e dom.Event) {
			toast.Show()
		}),
		toast,
	), sourcecode()
}

func Example_Toast_004() (mvc.View, string) {
	group := bs.ToastGroup(
		mvc.WithClass("position-absolute", "top-0", "start-0", "w-100", "h-100", "d-flex", "flex-column", "align-items-end", "justify-content-end", "p-3", "pe-none", "overflow-hidden"),
	)
	return bs.Container(
		bs.Button("Show Toast", mvc.WithClass("mb-3")).AddEventListener("click", func(e dom.Event) {
			toast := bs.Toast(
				mvc.WithClass("pe-auto"),
				mvc.Counter("toast"),
			)
			group.Slot("body").AppendChild(toast.Root())
			toast.Show()
		}),
		bs.Container(
			mvc.WithClass("position-relative", "overflow-hidden"),
			mvc.WithStyle("height: 300px;"),
			bs.WithBorder(),
			group,
		),
	), sourcecode()
}
