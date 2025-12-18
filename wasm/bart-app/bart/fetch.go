package bart

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/js"
)

const (
	// https://www.bart.gov/schedules/developers/api
	apiEndpoint = "https://api.bart.gov/api"
	apiKey      = "MW9S-E7SL-26DU-VV8V"
)

func Fetch(path string, params url.Values, callback func(*js.FetchResponse, error)) {
	u, err := url.Parse(apiEndpoint)
	if err != nil {
		panic(err)
	}
	u.Path = filepath.Join(u.Path, path)
	q := u.Query()
	q.Set("key", apiKey)
	q.Set("json", "y")
	for key, values := range params {
		for _, value := range values {
			q.Add(key, value)
		}
	}
	u.RawQuery = q.Encode()
	js.Fetch(u.String()).Done(func(value js.Value, err error) {
		if err != nil {
			callback(nil, err)
			return
		}
		callback(js.ResponseFrom(value), nil)
	})
}

func FetchStations(fn func([]Station, error)) {
	Fetch("stn.aspx", url.Values{"cmd": []string{"stns"}}, func(response *js.FetchResponse, err error) {
		if err != nil {
			fmt.Println("Fetch error:", err)
			return
		}

		response.Text().Done(func(textValue js.Value, err error) {
			if err != nil {
				fn(nil, err)
				return
			}

			// Parse JSON using Go's encoding/json
			var result StationsResponse
			if err := json.Unmarshal([]byte(textValue.String()), &result); err != nil {
				fn(nil, err)
				return
			} else {
				fn(result.Root.Stations.Station, nil)
			}
		})
	})
}
