package main

import (
	"encoding/json"
	"fmt"

	// Packages
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"
)

const (
	StationUrl = "https://api.bart.gov/api/stn.aspx?cmd=stns&key=MW9S-E7SL-26DU-VV8V&json=y"
)

func main() {
	// Fetch station data - now we can flatten the promise chain!
	jsutil.Fetch(StationUrl).Then(func(value jsutil.Value) error {
		// Wrap the response value
		resp := jsutil.NewResponse(value)
		if !resp.Ok() {
			return fmt.Errorf("request failed with status %d", resp.Status())
		}

		// Return the inner promise as a PromiseError to flatten the chain
		// This will cause the outer Then to wait for resp.Text() to complete
		return jsutil.NewPromiseError(resp.Text().Then(func(textValue jsutil.Value) error {
			// Convert the JavaScript string value to a Go string
			jsonText := jsutil.ToString(textValue)

			// Unmarshal into our struct
			var response StationsResponse
			if err := json.Unmarshal([]byte(jsonText), &response); err != nil {
				return err
			}

			// Print it out
			data, err := json.MarshalIndent(response.Root.Stations, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
			return nil
		}))
	}).Catch(func(err error) {
		// This now catches ALL errors, including from the inner promise!
		fmt.Println("Error:", err)
	})

	// Run the application
	select {}
}
