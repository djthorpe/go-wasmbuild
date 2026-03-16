package bart

import (
	"encoding/json"
	"net/url"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

const (
	// https://www.bart.gov/schedules/developers/api
	apiEndpoint = "https://api.bart.gov/api/"
	apiKey      = "MW9S-E7SL-26DU-VV8V"
)

// StationsProvider wraps Provider[[]Station] and bakes in the endpoint path
// and query parameters, so callers just call Fetch() with no arguments.
type StationsProvider struct {
	*mvc.Provider[[]Station]
}

// Fetch requests the full station list from the BART API.
func (p *StationsProvider) Fetch() {
	p.Provider.Fetch("stn.aspx", url.Values{"cmd": {"stns"}})
}

// NewStationsProvider returns a StationsProvider ready to use.
func NewStationsProvider() *StationsProvider {
	base, _ := url.Parse(apiEndpoint)
	q := base.Query()
	q.Set("key", apiKey)
	q.Set("json", "y")
	base.RawQuery = q.Encode()

	return &StationsProvider{mvc.NewProvider[[]Station](base, func(body []byte) ([]Station, error) {
		var r StationsResponse
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, err
		}
		return r.Root.Stations.Station, nil
	})}
}
