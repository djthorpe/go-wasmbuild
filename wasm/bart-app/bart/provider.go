package bart

import (
	"net/url"
	"time"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	// https://www.bart.gov/schedules/developers/api
	apiEndpoint = "https://api.bart.gov/api"
	apiKey      = "MW9S-E7SL-26DU-VV8V"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Provider is the entry point for all BART API data access.
// Create one with NewProvider; register listeners, then trigger fetches.
type Provider struct {
	base *url.URL

	// sub-providers (one per resource type)
	stationsProvider mvc.JSONProvider[StationsResponse]
	etdProvider      mvc.JSONProvider[ETDResponse]
	scheduleProvider mvc.JSONProvider[StationScheduleResponse]

	// fan-out listener lists
	stationListeners  []func([]Station, error)
	etdListeners      []func([]ETDStation, error)
	scheduleListeners []func(StationSchedule, error)
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewProvider creates a new BART API provider and wires up all sub-providers.
func NewProvider() *Provider {
	base, _ := url.Parse(apiEndpoint)
	p := &Provider{
		base:             base,
		stationsProvider: mvc.NewJSONProvider[StationsResponse](base),
		etdProvider:      mvc.NewJSONProvider[ETDResponse](base),
		scheduleProvider: mvc.NewJSONProvider[StationScheduleResponse](base),
	}

	// Fan-in: stationsProvider → Provider's station listeners
	p.stationsProvider.AddEventListener(func(resp StationsResponse, err error) {
		if err != nil {
			p.emitStations(nil, err)
			return
		}
		p.emitStations(resp.Root.Stations.Station, nil)
	})

	// Fan-in: etdProvider → Provider's ETD listeners
	p.etdProvider.AddEventListener(func(resp ETDResponse, err error) {
		if err != nil {
			p.emitDepartures(nil, err)
			return
		}
		p.emitDepartures(resp.Root.Station, nil)
	})

	// Fan-in: scheduleProvider → Provider's schedule listeners
	p.scheduleProvider.AddEventListener(func(resp StationScheduleResponse, err error) {
		if err != nil {
			p.emitSchedule(StationSchedule{}, err)
			return
		}
		p.emitSchedule(StationSchedule{
			Name:  resp.Root.Station.Name,
			Abbr:  resp.Root.Station.Abbr,
			Date:  resp.Root.Date,
			Items: resp.Root.Station.Items,
		}, nil)
	})

	return p
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// OnStations registers a listener that is called whenever a Stations fetch completes.
func (p *Provider) OnStations(fn func([]Station, error)) {
	p.stationListeners = append(p.stationListeners, fn)
}

// Stations triggers a fetch of all BART stations, notifying all OnStations listeners.
func (p *Provider) Stations() {
	p.stationsProvider.Get("stn.aspx", p.params(url.Values{"cmd": {"stns"}}))
}

// OnDepartures registers a listener called whenever a Departures fetch completes.
func (p *Provider) OnDepartures(fn func([]ETDStation, error)) {
	p.etdListeners = append(p.etdListeners, fn)
}

// Departures triggers a single fetch of real-time departures for the given
// station abbreviation (e.g. "RICH"), notifying all OnDepartures listeners.
// Pass "ALL" to fetch departures for every station.
func (p *Provider) Departures(station string) {
	p.etdProvider.Get("etd.aspx", p.params(url.Values{"cmd": {"etd"}, "orig": {station}}))
}

// DeparturesWithInterval cancels any existing interval, then starts polling
// departures for station at the given interval, firing immediately.
func (p *Provider) DeparturesWithInterval(station string, interval time.Duration) {
	p.etdProvider.GetWithInterval(
		"etd.aspx",
		interval,
		p.params(url.Values{"cmd": {"etd"}, "orig": {station}}),
	)
}

// CancelDepartures stops any active departure interval.
func (p *Provider) CancelDepartures() {
	p.etdProvider.Cancel()
}

// OnSchedule registers a listener called whenever a Schedule fetch completes.
func (p *Provider) OnSchedule(fn func(StationSchedule, error)) {
	p.scheduleListeners = append(p.scheduleListeners, fn)
}

// Schedule triggers a single fetch of today's scheduled departures for the
// given station abbreviation, notifying all OnSchedule listeners.
func (p *Provider) Schedule(station string) {
	p.scheduleProvider.Get("sched.aspx", p.params(url.Values{
		"cmd":  {"stnsched"},
		"orig": {station},
		"date": {"today"},
	}))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

// params builds a WithQuery option containing the standard BART auth params
// (key, json=y) merged with any caller-supplied values.
func (p *Provider) params(extra url.Values) js.FetchOption {
	q := url.Values{
		"key":  {apiKey},
		"json": {"y"},
	}
	for k, vs := range extra {
		q[k] = vs
	}
	return js.WithQuery(q)
}

// emitStations fans out to all registered station listeners.
func (p *Provider) emitStations(stations []Station, err error) {
	for _, fn := range p.stationListeners {
		fn(stations, err)
	}
}

// emitDepartures fans out to all registered ETD listeners.
func (p *Provider) emitDepartures(stations []ETDStation, err error) {
	for _, fn := range p.etdListeners {
		fn(stations, err)
	}
}

// emitSchedule fans out to all registered schedule listeners.
func (p *Provider) emitSchedule(sched StationSchedule, err error) {
	for _, fn := range p.scheduleListeners {
		fn(sched, err)
	}
}
