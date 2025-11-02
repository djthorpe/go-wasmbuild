package main

// Station represents a BART station with its details
type Station struct {
	Name          string `json:"name"`
	Abbr          string `json:"abbr"`
	GTFSLatitude  string `json:"gtfs_latitude"`
	GTFSLongitude string `json:"gtfs_longitude"`
	Address       string `json:"address"`
	City          string `json:"city"`
	County        string `json:"county"`
	State         string `json:"state"`
	Zipcode       string `json:"zipcode"`
}

// StationsResponse represents the complete API response structure
type StationsResponse struct {
	Root Root `json:"root"`
}

// Root contains the main response data
type Root struct {
	URI      URI      `json:"uri"`
	Stations Stations `json:"stations"`
	Message  string   `json:"message"`
}

// URI contains the API endpoint URL
type URI struct {
	CDATASection string `json:"#cdata-section"`
}

// Stations contains the list of stations
type Stations struct {
	Station []Station `json:"station"`
}
