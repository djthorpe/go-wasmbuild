package bart

import (
	"fmt"

	// Packages
	cds "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

// StationRow creates a Carbon table row for a station.
func StationRow(station Station) mvc.View {
	return cds.TableRow(
		cds.TableCell(cds.Strong(station.Name)),
		cds.TableCell(station.Abbr),
		cds.TableCell(station.City+", "+station.State),
		cds.TableCell(
			cds.InlineLink(
				fmt.Sprintf("https://maps.google.com/?q=%s,%s", station.Latitude, station.Longitude),
				"Map",
				mvc.WithAttr("target", "_blank"),
			),
		),
	)
}
