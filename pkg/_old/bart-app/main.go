package main

import (
	"encoding/json"
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-dom"
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	StationUrl = "https://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y"
)

func main() {
	// Fetch station data
	jsutil.Fetch(StationUrl).Then(func(value jsutil.Value) error {
		// Wrap the response value
		resp := jsutil.NewResponse(value)
		if !resp.Ok() {
			return fmt.Errorf("request failed with status %d", resp.Status())
		}

		return jsutil.NewPromiseError(resp.Text().Then(func(textValue jsutil.Value) error {
			var response StationsResponse
			if err := json.Unmarshal([]byte(jsutil.ToString(textValue)), &response); err != nil {
				return err
			}

			body := StationTable(response.Root.Stations.Station)
			fmt.Println(body)
			return nil
		}))
	}).Catch(func(err error) {
		fmt.Println("Error:", err)
	})

	select {}
}

func StationTable([]Station) dom.Element {
	return mvc.Element("div", mvc.Class("station-table"))
}
