package main

import (
	"encoding/json"
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	bart "github.com/djthorpe/go-wasmbuild/wasm/bart-app/bart"
)

// Application displays examples of MVC components
func main() {
	// Run the application
	mvc.New(
		bs.Button("Load").AddEventListener("click", func(evt dom.Event) {
			bart.FetchStations(func(stations []bart.Station, err error) {
				if err != nil {
					fmt.Println("Error fetching stations:", err)
					return
				}

				// Print stations
				for _, station := range stations {
					data, err := json.MarshalIndent(station, "", "  ")
					if err != nil {
						fmt.Println("JSON marshal error:", err)
						return
					}
					fmt.Printf("%s: %s\n", station.Abbr, string(data))
				}
			})
		}),
	).Run()
}
