package main

import (
	// Packages
	"fmt"

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ToastExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Toast Examples"), bs.HRule(),
		bs.Heading(3, "Toast", mvc.WithClass("mt-5")), Example(Example_Toast_001),
		bs.Heading(3, "Color", mvc.WithClass("mt-5")), Example(Example_Toast_002),
		bs.Heading(3, "Toast Group", mvc.WithClass("mt-5")), Example(Example_Toast_003),
	)
}

func Example_Toast_001() (mvc.View, string) {
	toast := bs.Toast(
		"toast-001", "Hello, world! This is a toast message.",
		mvc.WithClass("my-3"),
	).Header(
		bs.Strong("Toast Header", mvc.WithClass("me-auto")),
		bs.CloseButton(mvc.WithAttr("data-bs-dismiss", "toast")),
	).(mvc.ViewWithVisibility)

	return bs.Container(
		toast,
		bs.Button("Show Toast").AddEventListener("click", func(e dom.Event) {
			toast.Show()
		}),
	), sourcecode()
}

func Example_Toast_002() (mvc.View, string) {
	toast := bs.Toast(
		"toast-002",
		bs.Icon("bell-fill", mvc.WithClass("me-2")),
		"Hello, world! This is a toast message.",
		bs.WithColor(bs.Primary),
		mvc.WithClass("my-3"),
	)

	return bs.Container(
		toast,
		bs.Button("Show Toast").AddEventListener("click", func(e dom.Event) {
			toast.Show()
		}),
	), sourcecode()
}

func Example_Toast_003() (mvc.View, string) {
	toast3 := bs.Toast(
		"toast-003-1",
		bs.Icon("bell-fill", mvc.WithClass("me-2")),
		"Hello, world! This is a toast message.",
		bs.WithColor(bs.Warning),
		mvc.WithClass("my-3"),
	)
	toast4 := bs.Toast(
		"toast-003-2",
		bs.Icon("exclamation-octagon-fill", mvc.WithClass("me-2")),
		"Hello, world! This is a toast message.",
		bs.WithColor(bs.Danger),
		mvc.WithClass("my-3"),
	)
	group := bs.ToastGroup(toast3, toast4)
	position := bs.Select("select-003",
		bs.Option("None", ""),
		bs.Option("Top Left", "top-0 start-0"),
		bs.Option("Top Right", "top-0 end-0"),
		bs.Option("Bottom Left", "bottom-0 start-0"),
		bs.Option("Bottom Right", "bottom-0 end-0"),
	).AddEventListener("input", func(e dom.Event) {
		v := mvc.ViewFromEvent(e)
		switch v.Value() {
		case "":
			group.Apply(mvc.WithoutClass("position-absolute", "top-0", "bottom-0", "start-0", "end-0"))
			group.Apply(mvc.WithClass("position-static"))
		default:
			group.Apply(mvc.WithoutClass("position-static", "top-0", "bottom-0", "start-0", "end-0"))
			group.Apply(mvc.WithClass("position-absolute", v.Value()))
		}
		fmt.Printf("Selected position: %q\n", group.Root().ClassList().Values())
	})

	return bs.Container(
		group,
		bs.Container(
			bs.Button("Show Notification", mvc.WithClass("me-2")).AddEventListener("click", func(e dom.Event) {
				toast3.Show()
			}),
			bs.Button("Show Error", mvc.WithClass("me-2")).AddEventListener("click", func(e dom.Event) {
				toast4.Show()
			}),
			position,
			bs.WithFlex(bs.End),
		),
	), sourcecode()
}
