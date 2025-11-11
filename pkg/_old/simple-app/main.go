package main

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	views "github.com/djthorpe/go-wasmbuild/pkg/mvc/views"
)

const (
	StationUrl = "https://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y"
)

func main() {
	// Create an application and append a table
	app := mvc.New().Append(
		views.Table().Append(
			views.TableRow("cell1", "cell2"),
			views.TableRow("cell3", "cell4"),
		),
	)
	if app == nil {
		fmt.Println("Failed to create app")
		return
	}

	// Fetch JSON from a data source
	js.Fetch(StationUrl).Then(func(value js.Value) error {
		// The response will be a js.Response object
		fmt.Println("Fetched data", value)
		return nil
	}).Catch(func(err error) {
		fmt.Println("Error:", err)
	})

	// Print the document
	fmt.Println(app.Root().OwnerDocument())
}
