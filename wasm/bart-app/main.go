package main

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	bart "github.com/djthorpe/go-wasmbuild/wasm/bart-app/bart"
)

var stationsview mvc.View

// Application displays examples of MVC components
func main() {
	// Create stations table
	stationsview = bs.Table(
		bs.WithBorder(),
		bs.WithStripedRows(),
		mvc.WithClass("m-0"),
	).Header(
		bs.Container(
			bs.WithFlex(bs.Center),
			"Stations",
			bs.Button(bs.Icon("arrow-repeat"), bs.WithSize(bs.Small), mvc.WithClass("ms-auto")).AddEventListener("click", load),
		),
	)

	// Run the application
	mvc.New(
		bs.FluidContainer(
			mvc.WithClass("p-0"),
			bs.Row(
				mvc.WithClass("g-0"),
				bs.Col4(stationsview),
				bs.Col8(
					bs.WithPosition(bs.Center), bs.WithColor(bs.Light),
					"RIGHT",
				),
			),
		),
	).Run()
}

func load(evt dom.Event) {
	bart.FetchStations(func(stations []bart.Station, err error) {
		if err != nil {
			fmt.Println("Error fetching stations:", err)
			return
		}

		// Re-create the stations
		views := make([]any, len(stations))
		for i, station := range stations {
			views[i] = bart.StationRow(station)
		}
		stationsview.Content(views)
	})
}
