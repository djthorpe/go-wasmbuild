package bart

// BART API response structures
type StationsResponse struct {
	Root struct {
		Stations struct {
			Station []Station `json:"station"`
		} `json:"stations"`
	} `json:"root"`
}

type Station struct {
	Name      string `json:"name"`
	Abbr      string `json:"abbr"`
	Latitude  string `json:"gtfs_latitude"`
	Longitude string `json:"gtfs_longitude"`
	Address   string `json:"address"`
	City      string `json:"city"`
	County    string `json:"county"`
	State     string `json:"state"`
	Zipcode   string `json:"zipcode"`
}
