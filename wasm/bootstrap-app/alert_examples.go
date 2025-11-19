package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func AlertExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Alerts & Toasts"),
		bs.HRule(),
		bs.Heading(3, "Contextual Alerts", mvc.WithClass("mt-5")), Example(Example_Alert_001),
		bs.Heading(3, "Dismissible Alerts", mvc.WithClass("mt-5")), Example(Example_Alert_002),
		bs.Heading(3, "Toast Notifications", mvc.WithClass("mt-5")), Example(Example_Alert_003),
	)
}

func Example_Alert_001() (mvc.View, string) {
	return bs.Container(
		mvc.WithClass("d-flex", "flex-column", "gap-3"),
		alertBox([]string{"alert-primary"}, "A simple primary alert—check it out!"),
		alertBox([]string{"alert-success"}, "A success alert with a ", bs.Strong("bold"), " message."),
		alertBox([]string{"alert-warning"}, bs.Strong("Warning!"), " Better check yourself."),
		alertBox([]string{"alert-danger"}, "Danger alert—something went wrong."),
	), sourcecode()
}

func Example_Alert_002() (mvc.View, string) {
	return bs.Container(
		mvc.WithClass("d-flex", "flex-column", "gap-3"),
		alertBox([]string{"alert-warning", "alert-dismissible", "fade", "show"},
			bs.Strong("Heads up!"), " You should check in on some of those fields below.",
			dismissButton("Close warning", "alert"),
		),
		alertBox([]string{"alert-info", "alert-dismissible", "fade", "show"},
			"This is an info alert with a dismiss button.",
			dismissButton("Close info", "alert"),
		),
	), sourcecode()
}

func Example_Alert_003() (mvc.View, string) {
	return bs.Container(
		mvc.WithClass("d-flex", "flex-wrap", "gap-3"),
		toastBox("System", "11 mins ago", "Hello, world! This is a toast message."),
		toastBox("Build", "2 mins ago", "A new deployment finished successfully.", mvc.WithClass("text-bg-success")),
		toastBox("Warning", "Just now", "Pay attention to this toast.", mvc.WithClass("text-bg-warning")),
	), sourcecode()
}

func alertBox(classes []string, children ...any) any {
	args := []any{
		mvc.WithClass(append([]string{"alert"}, classes...)...),
		mvc.WithAttr("role", "alert"),
	}
	args = append(args, children...)
	return mvc.HTML("div", args...)
}

func dismissButton(label, target string) any {
	return mvc.HTML("button",
		mvc.WithClass("btn-close"),
		mvc.WithAttr("type", "button"),
		mvc.WithAttr("data-bs-dismiss", target),
		mvc.WithAttr("aria-label", label),
	)
}

func toastBox(title, timestamp, body string, opts ...any) any {
	args := []any{
		mvc.WithClass("toast", "show"),
		mvc.WithAttr("role", "alert"),
		mvc.WithAttr("aria-live", "assertive"),
		mvc.WithAttr("aria-atomic", "true"),
	}
	args = append(args, opts...)
	header := mvc.HTML("div",
		mvc.WithClass("toast-header"),
		bs.Strong(title, mvc.WithClass("me-auto")),
		mvc.HTML("small", timestamp),
		dismissButton("Close toast", "toast"),
	)
	bodyEl := mvc.HTML("div", mvc.WithClass("toast-body"), body)
	return mvc.HTML("div", append(args, header, bodyEl)...)
}
