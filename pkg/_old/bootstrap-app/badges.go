package main

import (
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func Badges() mvc.View {
	return bs.Container().Append(
		// Heading with title
		bs.Heading(1, mvc.WithClass("mt-4", "mb-4")).Append("Bootstrap Badge Examples"),

		// Badges in Headings
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Badges in Headings"),
		bs.Heading(1).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),
		bs.Heading(2).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),
		bs.Heading(3).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),
		bs.Heading(4).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),
		bs.Heading(5).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),
		bs.Heading(6).Append("Example heading ", bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")),

		// Background Colors
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Background Colors"),
		bs.Para(
			bs.Badge(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("Primary"),
			bs.Badge(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Secondary"),
			bs.Badge(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("Success"),
			bs.Badge(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("Danger"),
			bs.Badge(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("Warning"),
			bs.Badge(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("Info"),
			bs.Badge(bs.WithColor(bs.LIGHT), mvc.WithClass("me-2")).Append("Light"),
			bs.Badge(bs.WithColor(bs.DARK), mvc.WithClass("me-2")).Append("Dark"),
		),

		// Pill Badges
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Pill Badges"),
		bs.Para(
			bs.PillBadge(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("Primary"),
			bs.PillBadge(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Secondary"),
			bs.PillBadge(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("Success"),
			bs.PillBadge(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("Danger"),
			bs.PillBadge(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("Warning"),
			bs.PillBadge(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("Info"),
			bs.PillBadge(bs.WithColor(bs.LIGHT), mvc.WithClass("me-2")).Append("Light"),
			bs.PillBadge(bs.WithColor(bs.DARK), mvc.WithClass("me-2")).Append("Dark"),
		),

		// Badges in Buttons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Badges in Buttons"),
		bs.Para(
			bs.Button(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append(
				"Notifications ",
				bs.Badge(bs.WithColor(bs.SECONDARY)).Append("4"),
			),
			bs.Button(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(
				"Messages ",
				bs.Badge(bs.WithColor(bs.LIGHT)).Append("9"),
			),
			bs.Button(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append(
				"Alerts ",
				bs.Badge(bs.WithColor(bs.WARNING)).Append("2"),
			),
			bs.Button(bs.WithColor(bs.INFO)).Append(
				"Updates ",
				bs.Badge(bs.WithColor(bs.DARK)).Append("12"),
			),
		),

		// Counter Badges
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Counter Badges"),
		bs.Para(
			bs.Badge(bs.WithColor(bs.PRIMARY), mvc.WithClass("me-2")).Append("1"),
			bs.Badge(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("5"),
			bs.Badge(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("10"),
			bs.Badge(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("99+"),
			bs.PillBadge(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("3"),
			bs.PillBadge(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("42"),
		),

		// Status Badges
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Status Badges"),
		bs.Para(
			bs.PillBadge(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append("Active"),
			bs.PillBadge(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append("Pending"),
			bs.PillBadge(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append("Inactive"),
			bs.PillBadge(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append("Draft"),
			bs.PillBadge(bs.WithColor(bs.SECONDARY), mvc.WithClass("me-2")).Append("Archived"),
		),

		// Badges with Icons
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Badges with Icons"),
		bs.Para(
			bs.Badge(bs.WithColor(bs.SUCCESS), mvc.WithClass("me-2")).Append(
				bs.Icon("check-circle"),
				" Verified",
			),
			bs.Badge(bs.WithColor(bs.DANGER), mvc.WithClass("me-2")).Append(
				bs.Icon("x-circle"),
				" Error",
			),
			bs.Badge(bs.WithColor(bs.WARNING), mvc.WithClass("me-2")).Append(
				bs.Icon("exclamation-triangle"),
				" Warning",
			),
			bs.Badge(bs.WithColor(bs.INFO), mvc.WithClass("me-2")).Append(
				bs.Icon("info-circle"),
				" Info",
			),
		),

		// Positioned Badge on Button
		bs.Heading(2, mvc.WithClass("mt-4")).Append("Positioned Badges"),
		createPositionedBadgePara(),

		// Badge with indicators (dot notification)
		bs.Para().Append(
			createButtonWithIndicator("Alerts"),
		),
	)
}

// Helper function to create a paragraph with positioned badges
func createPositionedBadgePara() mvc.View {
	para := bs.Para(
		createButtonWithBadge("Inbox", "99+", "unread messages"),
		createButtonWithBadge("Profile", "5", "new notifications"),
	)
	para.Opts(mvc.WithClass("mb-3"))
	return para
}

// Helper function to create a button with a positioned badge
func createButtonWithBadge(text, count, screenReaderText string) mvc.View {
	btn := bs.Button(bs.WithColor(bs.PRIMARY))
	btn.Opts(mvc.WithClass("position-relative", "me-2"))

	badge := bs.PillBadge(bs.WithColor(bs.DANGER))
	badge.Opts(mvc.WithClass(
		"position-absolute",
		"top-0",
		"start-100",
		"translate-middle",
	))

	srText := mvc.Tag("SPAN", mvc.WithClass("visually-hidden")).Append(screenReaderText)
	badge.Append(count, srText)

	btn.Append(text, badge)
	return btn
}

// Helper function to create a button with a dot indicator
func createButtonWithIndicator(text string) mvc.View {
	btn := bs.Button(bs.WithColor(bs.PRIMARY))
	btn.Opts(mvc.WithClass("position-relative"))

	indicator := mvc.Tag("SPAN", mvc.WithClass(
		"position-absolute",
		"top-0",
		"start-100",
		"translate-middle",
		"p-2",
		"bg-danger",
		"border",
		"border-light",
		"rounded-circle",
	))

	srText := mvc.Tag("SPAN", mvc.WithClass("visually-hidden")).Append("New Alerts")
	indicator.Append(srText)

	btn.Append(text, indicator)
	return btn
}
