package main

import (
	"fmt"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func TableExamples() mvc.View {
	return bs.Container(
		bs.Table(
			bs.TableRow("Jan", "$100.00"),
			bs.TableRow("Feb", "$200.00"),
			bs.TableRow("Mar", "$300.00"),
			bs.TableRow("Apr", "$400.00", bs.WithColor(bs.Primary)),
		).Header(
			"Month", "Revenue",
		).Footer(
			"Total", "$1000",
		),
		bs.Heading(3, "Striped Rows"),
		bs.Table(
			bs.TableRow("Jan", "$100.00"),
			bs.TableRow("Feb", "$200.00"),
			bs.TableRow("Mar", "$300.00"),
			bs.TableRow("Apr", "$400.00"),
			bs.WithStripedRows(),
		).Header(
			"Month", "Revenue",
		).Footer(
			"Total", "$1000",
		),
		bs.Heading(3, "Striped Columns"),
		bs.Table(
			bs.TableRow("Jan", "$100.00"),
			bs.TableRow("Feb", "$200.00"),
			bs.TableRow("Mar", "$300.00"),
			bs.TableRow("Apr", "$400.00"),
			bs.WithStripedColumns(),
		).Header(
			"Month", "Revenue",
		).Footer(
			"Total", "$1000",
		),
		bs.Heading(3, "Small Table with Hover"),
		bs.Table(
			bs.TableRow("Feb", "$200.00"),
			bs.TableRow("Mar", "$300.00"),
			bs.TableRow("Apr", "$400.00"),
			bs.WithRowHover(),
			bs.WithSize(bs.Small),
		).Header(
			"Month", "Revenue",
		).Footer(
			"Total", "$1000",
		).AddEventListener("click", func(e Event) {
			view := mvc.ViewFromEvent(e)
			if view == nil {
				return
			}
			switch view.Name() {
			case bs.ViewTableRow:
				// TODO: Determine the index of the row clicked in the table
				fmt.Println("Clicked on row:", view)
			}
		}),
	)
}
