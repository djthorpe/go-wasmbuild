package bart

import (
	"fmt"
	"time"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// StationCard creates a Bootstrap card view for a BART station
func StationCard(station Station) mvc.View {
	card := bs.Card().
		Header(bs.Badge(station.Abbr)).
		Footer()
	card.Content(
		bs.Heading(5, station.Name),
		bs.Para(station.Address),
		bs.Smaller(fmt.Sprintf("%s, %s %s", station.City, station.State, station.Zipcode)),
	)
	return card
}

// StationList creates a card group with all stations
func StationList(stations []Station) mvc.View {
	cards := make([]any, len(stations))
	for i, station := range stations {
		cards[i] = StationCard(station)
	}
	return bs.CardGroup(cards...)
}

// StationRow creates a compact row view for a station (for tables).
// onClick is called with the station abbreviation when the row is clicked.
func StationRow(station Station, onClick func(abbr string)) mvc.View {
	row := bs.TableRow(
		bs.Container(
			bs.Strong(station.Name),
			bs.Badge(station.Abbr, mvc.WithClass("m-1")),
			mvc.HTML("BR"),
			bs.Smaller(
				mvc.HTML("SPAN", station.Address, mvc.WithClass("me-2")),
				bs.IconLink(
					fmt.Sprintf("https://maps.google.com/?q=%s,%s", station.Latitude, station.Longitude),
					bs.Icon("geo-alt-fill"), "View on Map",
					mvc.WithAttr("target", "_blank"),
				),
			),
		),
	)
	row.AddEventListener("click", func(_ dom.Event) {
		onClick(station.Abbr)
	})
	return row
}

// DeparturesView renders real-time departure info for a single ETDStation.
func DeparturesView(station ETDStation) mvc.View {
	table := bs.Table(
		bs.WithBorder(),
		bs.WithStripedRows(),
		mvc.WithClass("m-0"),
	).Header(
		bs.Container(
			bs.WithFlex(bs.Center),
			bs.Strong(station.Name),
			bs.Badge(station.Abbr, mvc.WithClass("ms-2")),
		),
	)

	rows := make([]any, 0, len(station.ETD))
	for _, etd := range station.ETD {
		for _, est := range etd.Estimate {
			minutes := est.Minutes
			var label string
			if minutes == "Leaving" {
				label = "Now"
			} else {
				label = fmt.Sprintf("%s min", minutes)
			}
			rows = append(rows, bs.TableRow(
				bs.Container(
					bs.Strong(etd.Destination),
					mvc.HTML("BR"),
					bs.Smaller(
						mvc.HTML("SPAN", est.Direction, mvc.WithClass("me-2")),
						mvc.HTML("SPAN", fmt.Sprintf("Platform %s", est.Platform)),
					),
				),
				bs.Badge(
					label,
					mvc.WithClass("ms-auto"),
					mvc.WithAttr("style", fmt.Sprintf("background-color:%s", est.HexColor)),
				),
			))
		}
	}
	table.Content(rows...)
	return table
}

// ScheduleView renders upcoming scheduled departures for a station.
// Times from the API are in SFO (America/Los_Angeles) time and are converted
// to the user's local time for display. Past trains are excluded.
// BART's service day extends past midnight: times before 4 AM are treated as
// belonging to the next calendar day.
func ScheduleView(sched StationSchedule) mvc.View {
	loc, _ := time.LoadLocation("America/Los_Angeles")
	now := time.Now()

	table := bs.Table(
		bs.WithBorder(),
		bs.WithStripedRows(),
		mvc.WithClass("m-0"),
	).Header(
		bs.Container(
			bs.WithFlex(bs.Center),
			bs.Strong("Next Departures for "+sched.Name),
		),
	)

	const maxUpcoming = 20
	rows := make([]any, 0, maxUpcoming)
	for _, item := range sched.Items {
		t, err := time.ParseInLocation("1/2/2006 3:04 PM", sched.Date+" "+item.OrigTime, loc)
		if err != nil {
			continue
		}
		// Times before 4 AM belong to the next calendar day in BART's service schedule.
		if t.Hour() < 4 {
			t = t.AddDate(0, 0, 1)
		}
		if !t.After(now) {
			continue
		}
		rows = append(rows, bs.TableRow(
			bs.Container(
				bs.Strong(item.TrainHeadStation),
				mvc.HTML("BR"),
				bs.Smaller(
					mvc.HTML("SPAN", item.Platform, mvc.WithClass("me-2")),
					mvc.HTML("SPAN", item.Line),
				),
			),
			bs.Badge(
				t.Local().Format("3:04 PM"),
				mvc.WithClass("ms-auto"),
			),
		))
		if len(rows) >= maxUpcoming {
			break
		}
	}
	table.Content(rows...)
	return table
}
