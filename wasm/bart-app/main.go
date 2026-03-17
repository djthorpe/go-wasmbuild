package main

import (
	_ "time/tzdata" // embed timezone database for WASM

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	bart "github.com/djthorpe/go-wasmbuild/wasm/bart-app/bart"
)

var (
	stationsview  mvc.View
	departureview mvc.View
	scheduleview  mvc.View
	controller    = bart.NewController()
)

// Application displays BART station data
func main() {
	// Wire model → view: full station list refresh
	controller.Stations.OnSet(func(stations []bart.Station) {
		views := make([]any, len(stations))
		for i, station := range stations {
			row := bart.StationRow(station, controller.SelectStation)
			controller.RegisterStationRow(station.Abbr, row)
			views[i] = row
		}
		stationsview.Content(views)
	})

	// Wire model → view: real-time departures for selected station
	controller.ETD.AddEventListener(func(stations []bart.ETDStation) {
		if len(stations) == 0 {
			return
		}
		departureview.Content(bart.DeparturesView(stations[0]))
	})

	// Wire model → view: schedule for selected station
	controller.Schedule.AddEventListener(func(scheds []bart.StationSchedule) {
		if len(scheds) == 0 {
			return
		}
		scheduleview.Content(bart.ScheduleView(scheds[0]))
	})

	// Create stations table
	stationsview = bs.Table(
		bs.WithBorder(),
		bs.WithStripedRows(),
		bs.WithRowHover(),
		mvc.WithClass("m-0"),
	).Header(
		bs.Container(
			bs.WithFlex(bs.Center),
			"Stations",
			bs.Button(bs.Icon("arrow-repeat"), bs.WithSize(bs.Small), mvc.WithClass("ms-auto")).AddEventListener("click", load),
		),
	)
	controller.SetStationsTable(stationsview.(mvc.ActiveGroup))

	// Create right-panel views (populated on station click)
	departureview = bs.Container(mvc.WithClass("p-3 text-muted"), "Select a station to see departures")
	scheduleview = bs.Container()

	// Load stations on startup
	controller.Start()

	// Run the application
	mvc.New(
		bs.FluidContainer(
			mvc.WithClass("p-0"),
			bs.Row(
				mvc.WithClass("g-0"),
				bs.Col4(stationsview),
				bs.Col8(departureview, scheduleview),
			),
		),
	).Run()
}

func load(evt dom.Event) {
	controller.Refresh()
}
