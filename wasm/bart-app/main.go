package main

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	bart "github.com/djthorpe/go-wasmbuild/wasm/bart-app/bart"
)

var (
	stationsview mvc.View
	toastStack   mvc.View
	ctrl         *bart.StationsController
)

// Application displays BART station data
func main() {
	// Fixed-position overlay for toast notifications
	toastStack = mvc.New(
		mvc.WithAttr("style", "position:fixed;top:4.5rem;right:1rem;z-index:9999;"+
			"display:flex;flex-direction:column;gap:var(--cds-spacing-03,0.5rem);"),
	)
	toastStack.AddEventListener("cds-notification-closed", func(e dom.Event) {
		if el, ok := e.Target().(dom.Element); ok {
			el.Remove()
		}
	})

	// Controller: fetch errors fire an error toast
	ctrl = bart.NewStationsController(func(err error) {
		toastStack.Root().AppendChild(cds.ToastNotification(
			cds.WithNotificationKind(cds.NotificationError),
			cds.WithNotificationTitle("Failed to load stations"),
			cds.WithNotificationSubtitle(err.Error()),
			cds.WithNotificationTimeout(8000),
		).Root())
	})

	// Create stations table
	stationsview = cds.Table(
		cds.WithTableZebra(),
	).Header(
		cds.TableHeaderCell("Name"),
		cds.TableHeaderCell("Code"),
		cds.TableHeaderCell("Location"),
		cds.TableHeaderCell("Map"),
	)

	// Model listener: re-render the table whenever the data changes.
	// Registered after stationsview is created because AddEventListener
	// calls fn immediately with the current (empty) slice.
	ctrl.Model.AddEventListener("table", func(stations []bart.Station) {
		rows := make([]any, len(stations))
		for i, s := range stations {
			rows[i] = bart.StationRow(s)
		}
		stationsview.Content(rows)
		if len(stations) > 0 {
			toastStack.Root().AppendChild(cds.ToastNotification(
				cds.WithNotificationKind(cds.NotificationSuccess),
				cds.WithNotificationTitle("Stations loaded"),
				cds.WithNotificationSubtitle(fmt.Sprintf("%d stations retrieved.", len(stations))),
				cds.WithNotificationTimeout(4000),
			).Root())
		}
	})

	// Shell (fixed header)
	mvc.New(cds.Header(cds.WithTheme(cds.ThemeG100)).Label("/", "BART App"))

	// Content
	mvc.New(
		mvc.WithClass("cds--content"),
		cds.Section(
			cds.Heading(2, "BART Stations"),
			cds.Button("Load stations", cds.WithButtonKind(cds.ButtonPrimary)).
				AddEventListener("click", func(e dom.Event) {
					ctrl.Load()
				}),
			stationsview,
		),
	).Run()
}
