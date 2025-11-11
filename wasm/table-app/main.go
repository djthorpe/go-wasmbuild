package main

import (
	"fmt"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	view "github.com/djthorpe/go-wasmbuild/pkg/mvc/view"
)

// Application is an example of using a table view and controller
func main() {

	// Create the table
	table := view.Table(mvc.WithClass("table")).Header(
		"Month", "Cost",
	).Footer(
		"Total", "$100",
	).Content(
		view.TableRow("Jan", "$100"),
	)

	// Set the content of the application
	app := mvc.New().Content(
		table,
	)

	// Print the application
	fmt.Println(app)
}
