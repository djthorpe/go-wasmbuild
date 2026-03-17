package bart

// BART API response structures

///////////////////////////////////////////////////////////////////////////////
// STATIONS

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

func (s Station) PrimaryKey() string { return s.Abbr }

///////////////////////////////////////////////////////////////////////////////
// ESTIMATED DEPARTURES (ETD)

type ETDResponse struct {
	Root struct {
		Station []ETDStation `json:"station"`
	} `json:"root"`
}

type ETDStation struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	ETD  []ETD  `json:"etd"`
}

func (s ETDStation) PrimaryKey() string { return s.Abbr }

type ETD struct {
	Destination  string     `json:"destination"`
	Abbreviation string     `json:"abbreviation"`
	Limited      string     `json:"limited"`
	Estimate     []Estimate `json:"estimate"`
}

type Estimate struct {
	Minutes     string `json:"minutes"`
	Platform    string `json:"platform"`
	Direction   string `json:"direction"`
	Length      string `json:"length"`
	Color       string `json:"color"`
	HexColor    string `json:"hexcolor"`
	BikeFlag    string `json:"bikeflag"`
	Delay       string `json:"delay"`
	CancelFlag  string `json:"cancelflag"`
	DynamicFlag string `json:"dynamicflag"`
}

///////////////////////////////////////////////////////////////////////////////
// STATION SCHEDULE

// StationScheduleResponse is the raw API response for the stnsched command.
type StationScheduleResponse struct {
	Root struct {
		Date    string          `json:"date"`
		Station ScheduleStation `json:"station"`
	} `json:"root"`
}

// ScheduleStation holds a station's full day schedule.
type ScheduleStation struct {
	Name  string         `json:"name"`
	Abbr  string         `json:"abbr"`
	Items []ScheduleItem `json:"item"`
}

// ScheduleItem is a single scheduled departure. JSON keys use @ because the
// BART API maps XML attributes directly into JSON.
type ScheduleItem struct {
	Line             string `json:"@line"`
	TrainHeadStation string `json:"@trainHeadStation"`
	OrigTime         string `json:"@origTime"`
	BikeFlag         string `json:"@bikeflag"`
	Load             string `json:"@load"`
	Platform         string `json:"@platform"`
}

// StationSchedule bundles a parsed station schedule with its SFO service date.
type StationSchedule struct {
	Name  string
	Abbr  string
	Date  string // "M/D/YYYY" as returned by the API
	Items []ScheduleItem
}
