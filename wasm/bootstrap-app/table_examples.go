package main

import (
	// Packages

	"slices"

	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TableExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "Table Examples"), bs.HRule(),
		bs.Heading(3, "Table", mvc.WithClass("mt-5")), Example(Example_Table_001),
		bs.Heading(3, "Table Styling", mvc.WithClass("mt-5")), Example(Example_Table_002),
	)
}

func Example_Table_001() (mvc.View, string) {
	return bs.Table(
		bs.Row("Alice", "Engineer", "New York", "$80,000"),
		bs.Row("Bob", "Manager", "Chicago", "$90,000"),
		bs.Row("Charlie", "Designer", "San Francisco", "$85,000"),
		bs.Row("Diana", "Developer", "Seattle", "$95,000"),
	).Header(
		"Name", "Position", "Location", "Salary",
	).Footer(
		bs.WithColor(bs.Light),
		"", "", "Total", "$350,000",
	), sourcecode()
}

type ViewWithEnabled interface {
	mvc.View
	Enabled() []string
}

func Example_Table_002() (mvc.View, string) {
	table := bs.Table(
		bs.Row("Alice", "Engineer", "New York", "$80,000"),
		bs.Row("Bob", "Manager", "Chicago", "$90,000"),
		bs.Row("Charlie", "Designer", "San Francisco", "$85,000"),
		bs.Row("Diana", "Developer", "Seattle", "$95,000"),
	).Header(
		"Name", "Position", "Location", "Salary",
	).Footer(
		mvc.WithClass("table-group-divider"),
		"", "", "Total", "$350,000",
	)
	opts := func(state bool, class string) mvc.Opt {
		if state {
			return mvc.WithClass(class)
		} else {
			return mvc.WithoutClass(class)
		}
	}
	return bs.Container(
		bs.Form("table_002",
			mvc.WithClass("mb-3", "p-3"), bs.WithBorder(),
			bs.InlineSwitchGroup("table_002_styles", "Small", "Dark", "Bordered", "Striped Rows", "Hover Rows"),
		),
		table,
	).AddEventListener("input", func(e dom.Event) {
		enabled := mvc.ViewFromEvent(e).(ViewWithEnabled).Enabled()
		table.Apply([]mvc.Opt{
			opts(slices.Contains(enabled, "Small"), "table-sm"),
			opts(slices.Contains(enabled, "Dark"), "table-dark"),
			opts(slices.Contains(enabled, "Bordered"), "table-bordered"),
			opts(slices.Contains(enabled, "Striped Rows"), "table-striped"),
			opts(slices.Contains(enabled, "Hover Rows"), "table-hover"),
		}...)
	}), sourcecode()
}
