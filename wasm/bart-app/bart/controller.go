package bart

import (
	"time"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Controller owns the BART provider and all domain models. It wires provider
// responses into models so that views only need to subscribe to model events.
//
//	c := bart.NewController()
//	c.Stations.OnSet(func(stations []bart.Station) { ... })
//	c.Stations.OnAdded(func(e mvc.AddedEvent[bart.Station]) { ... })
//	c.Start()
type Controller struct {
	// Stations is the full list of BART stations, keyed by abbreviation.
	// OnSet fires after Refresh or initial load. OnAdded/OnDeleted fire for
	// incremental changes.
	Stations mvc.KeyedModel[string, Station]

	// ETD holds the real-time departure data for the currently selected
	// station. It contains at most one element. OnSet fires whenever new
	// departure data arrives (on the polling interval).
	ETD mvc.Model[ETDStation]

	// Schedule holds the scheduled departures for the currently selected
	// station. It contains at most one element. OnSet fires on each fetch.
	Schedule mvc.Model[StationSchedule]

	// Provider is the BART API provider. Exposed here for direct fetches if needed,
	// but most interactions should go through the public methods on Controller.
	provider *Provider
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewController creates a Controller with a wired-up BART provider.
// Register listeners on the public models before calling Start.
func NewController() *Controller {
	c := &Controller{provider: NewProvider()}

	// provider → Stations model
	c.provider.OnStations(func(stations []Station, err error) {
		if err != nil {
			return
		}
		c.Stations.Set(stations)
	})

	// provider → ETD model (0 or 1 elements)
	c.provider.OnDepartures(func(stations []ETDStation, err error) {
		if err != nil || len(stations) == 0 {
			c.ETD.Set(nil)
			return
		}
		c.ETD.Set(stations[:1])
	})

	// provider → Schedule model (0 or 1 elements)
	c.provider.OnSchedule(func(sched StationSchedule, err error) {
		if err != nil {
			c.Schedule.Set(nil)
			return
		}
		c.Schedule.Set([]StationSchedule{sched})
	})

	return c
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Start loads the station list. Call after registering all model listeners.
func (c *Controller) Start() {
	c.provider.Stations()
}

// SelectStation starts a 1-minute polling interval for real-time departures
// and fetches today's schedule for the given station abbreviation.
// Cancels any existing departure interval first.
func (c *Controller) SelectStation(abbr string) {
	c.provider.CancelDepartures()
	c.provider.DeparturesWithInterval(abbr, time.Minute)
	c.provider.Schedule(abbr)
}

// Refresh cancels any active departure polling and reloads the station list.
func (c *Controller) Refresh() {
	c.provider.CancelDepartures()
	c.provider.Stations()
}
