package bart

import (
	"fmt"

	// Packages
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

// StationRow creates a compact row view for a station (for tables)
func StationRow(station Station) mvc.View {
	return bs.TableRow(
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
}
