package bart

import (
	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// StationsController owns the station data model and its provider.
// Subscribe to Model to re-render views; call Load to trigger a fetch.
type StationsController struct {
	// Model holds the fetched station list. Add a named listener to be
	// notified on every Set (and immediately with the current empty slice):
	//
	//   ctrl.Model.AddEventListener("table", func(stations []Station) {
	//       // re-render the table from stations
	//   })
	Model    mvc.Model[Station]
	provider *StationsProvider
}

// NewStationsController returns a ready controller. onErr is called on fetch
// failure; pass nil to ignore errors.
func NewStationsController(onErr func(error)) *StationsController {
	c := new(StationsController)
	c.provider = NewStationsProvider()
	c.provider.AddEventListener(func(stations []Station, err error) {
		if err != nil {
			if onErr != nil {
				onErr(err)
			}
			return
		}
		c.Model.Set(stations)
	})
	return c
}

// Load fetches the full station list from the BART API.
func (c *StationsController) Load() {
	c.provider.Fetch()
}
