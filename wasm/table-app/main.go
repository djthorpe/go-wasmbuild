package main

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	view "github.com/djthorpe/go-wasmbuild/pkg/mvc/view"
)

// Station represents a BART station with its details
type Station struct {
	Name          string `json:"name"`
	Abbr          string `json:"abbr"`
	GTFSLatitude  string `json:"gtfs_latitude"`
	GTFSLongitude string `json:"gtfs_longitude"`
	Address       string `json:"address"`
	City          string `json:"city"`
	County        string `json:"county"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
}

// Create a set of stations
var (
	stations = []Station{
		{
			Name:          "12th St. Oakland City Center",
			Abbr:          "12TH",
			GTFSLatitude:  "37.803768",
			GTFSLongitude: "-122.271450",
			Address:       "1245 Broadway",
			City:          "Oakland",
			County:        "Alameda",
			State:         "CA",
			Zipcode:       "94612",
		},
		{
			Name:          "16th St. Mission",
			Abbr:          "16TH",
			GTFSLatitude:  "37.765062",
			GTFSLongitude: "-122.419694",
			Address:       "2000 Mission St",
			City:          "San Francisco",
			County:        "San Francisco",
			State:         "CA",
			Zipcode:       "94110",
		},
	}
)

// Application is an example of using a table view and controller
func main() {

	// Create the table with the Station prototype
	table := view.Table(new(Station), mvc.WithClass("table", "table-hover")).Header(
		"name", "abbr", "gtfs_latitude", "gtfs_longitude", "address", "city", "county", "state", "zipcode",
	).Content(stations[0], stations[1])

	// Add rows to the table

	// Set the content of the application
	app := mvc.New().Content(
		table,
	)

	// Print the application
	fmt.Println(app)
}
