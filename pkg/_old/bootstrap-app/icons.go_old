package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Icons() mvc.View {
	return bs.Container(bs.WithMargin(bs.All, 5)).Append(
		// Title
		bs.Heading(1).Append(
			bs.Icon("bootstrap-fill"),
			" Bootstrap Icons Examples",
		),
		bs.Rule(),

		// Basic Icons Section
		bs.Heading(3).Append("Basic Icons"),
		bs.Para(
			"Icons can be used inline with text: ",
			bs.Icon("heart"),
			" Heart ",
			bs.Icon("star"),
			" Star ",
			bs.Icon("bell"),
			" Bell",
		),

		// Filled vs Outline Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Filled vs Outline Icons"),
		bs.Para(
			bs.Icon("heart"),
			" Outline Heart ",
			bs.Icon("heart-fill"),
			" Filled Heart",
		),
		bs.Para(
			bs.Icon("star"),
			" Outline Star ",
			bs.Icon("star-fill"),
			" Filled Star",
		),

		// Common Action Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Common Action Icons"),
		bs.Para(
			bs.Icon("search"),
			" Search ",
			bs.Icon("download"),
			" Download ",
			bs.Icon("upload"),
			" Upload ",
			bs.Icon("trash"),
			" Trash ",
			bs.Icon("pencil"),
			" Edit ",
			bs.Icon("check"),
			" Check ",
			bs.Icon("x"),
			" Close",
		),

		// Navigation Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Navigation Icons"),
		bs.Para(
			bs.Icon("house"),
			" Home ",
			bs.Icon("gear"),
			" Settings ",
			bs.Icon("person"),
			" Profile ",
			bs.Icon("envelope"),
			" Messages ",
			bs.Icon("cart"),
			" Cart",
		),

		// Arrow Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Arrow Icons"),
		bs.Para(
			bs.Icon("arrow-up"),
			" Up ",
			bs.Icon("arrow-down"),
			" Down ",
			bs.Icon("arrow-left"),
			" Left ",
			bs.Icon("arrow-right"),
			" Right ",
			bs.Icon("arrow-up-circle"),
			" Circle",
		),

		// Colored Icons (using text color utilities)
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Colored Icons"),
		bs.Para().Append(
			createColoredIcon("heart-fill", "danger", "Danger"),
			createColoredIcon("star-fill", "warning", "Warning"),
			createColoredIcon("check-circle-fill", "success", "Success"),
			createColoredIcon("info-circle-fill", "info", "Info"),
			createColoredIcon("exclamation-triangle-fill", "warning", "Alert"),
		),

		// Sized Icons (using font-size utilities)
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Different Sizes"),
		bs.Para().Append(
			createSizedIcon("heart-fill", "1", "Small"),
			createSizedIcon("heart-fill", "3", "Medium"),
			createSizedIcon("heart-fill", "5", "Large"),
		),

		// Icons in Buttons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Icons in Buttons"),
		bs.Para().Append(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append(
				bs.Icon("download"),
				" Download",
			),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(
				bs.Icon("check"),
				" Confirm",
			),
			bs.Button(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append(
				bs.Icon("trash"),
				" Delete",
			),
			bs.Button(bs.WithColor(bs.SECONDARY)).Append(
				bs.Icon("gear"),
				" Settings",
			),
		),

		// Icon-only Buttons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Icon-only Buttons"),
		bs.Para().Append(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append(bs.Icon("heart")),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(bs.Icon("star")),
			bs.Button(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append(bs.Icon("bell")),
			bs.Button(bs.WithColor(bs.DANGER)).Append(bs.Icon("trash")),
		),

		// File and Media Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("File & Media Icons"),
		bs.Para(
			bs.Icon("file-earmark"),
			" File ",
			bs.Icon("folder"),
			" Folder ",
			bs.Icon("image"),
			" Image ",
			bs.Icon("camera"),
			" Camera ",
			bs.Icon("film"),
			" Video ",
			bs.Icon("music-note"),
			" Music",
		),

		// Social & Communication Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("Social & Communication"),
		bs.Para(
			bs.Icon("share"),
			" Share ",
			bs.Icon("chat"),
			" Chat ",
			bs.Icon("telephone"),
			" Phone ",
			bs.Icon("envelope"),
			" Email ",
			bs.Icon("at"),
			" Mention",
		),

		// List with Icons
		bs.Heading(3, mvc.WithClass("mt-4")).Append("List with Icons"),
		createIconList(),
	)
}

// Helper function to create a colored icon with label
func createColoredIcon(iconName, color, label string) mvc.View {
	icon := bs.Icon(iconName)
	icon.Opts(mvc.WithClass("text-" + color))
	span := mvc.Span().Content(icon, " ", label)
	span.Opts(mvc.WithClass("me-3"))
	return span
}

// Helper function to create a sized icon with label
func createSizedIcon(iconName, size, label string) mvc.View {
	icon := bs.Icon(iconName)
	icon.Opts(mvc.WithClass("fs-" + size))
	span := mvc.Span().Content(icon, " ", label)
	span.Opts(mvc.WithClass("me-3"))
	return span
}

// Helper function to create a list with icons
func createIconList() mvc.View {
	items := []struct {
		icon string
		text string
	}{
		{"check-circle", "Task completed successfully"},
		{"clock", "Scheduled for later"},
		{"exclamation-triangle", "Warning: Review required"},
		{"info-circle", "Additional information available"},
		{"x-circle", "Action cancelled"},
	}

	list := bs.Para()
	list.Opts(mvc.WithClass("list-unstyled"))
	for _, item := range items {
		icon := bs.Icon(item.icon)
		icon.Opts(mvc.WithClass("me-2"))
		span := mvc.Span().Content(icon, item.text)
		span.Opts(mvc.WithClass("d-block", "mb-2"))
		list.Append(span)
	}
	return list
}
