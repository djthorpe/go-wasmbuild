package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ListExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-3"),
		Markdown("list_examples.md"),
		ExampleCard("Unstyled", Example_List_001),
		ExampleCard("Ordered", Example_List_002),
		ExampleCard("Bulleted", Example_List_003),
		ExampleCard("List Group", Example_List_004),
		ExampleCard("List Group with Badges", Example_List_005),
		ExampleCard("Definition", Example_List_006),
	)
}

func Example_List_001() (mvc.View, string) {
	return bs.UnstyledList(
		"Item 1",
		"Item 2",
		"Item 3",
	), sourcecode()
}

func Example_List_002() (mvc.View, string) {
	return bs.List(
		"Item 1",
		"Item 2",
		"Item 3",
	), sourcecode()
}

func Example_List_003() (mvc.View, string) {
	return bs.BulletList(
		"Item 1",
		"Item 2",
		"Item 3",
	), sourcecode()
}

func Example_List_004() (mvc.View, string) {
	return bs.ListGroup(
		"Item 1",
		"Item 2",
		"Item 3",
	), sourcecode()
}

func Example_List_005() (mvc.View, string) {
	return bs.ListGroup(
		bs.Row(
			bs.Col6(bs.Strong("Inbox")),
			bs.Col6(bs.PillBadge("99+"), bs.WithPosition(bs.End)),
		),
		bs.Row(
			bs.Col6("Drafts"),
			bs.Col6(bs.PillBadge("1"), bs.WithPosition(bs.End)),
		),
		bs.Row(
			bs.Col6("Sent"),
			bs.Col6(bs.PillBadge("0", bs.WithColor(bs.Danger)), bs.WithPosition(bs.End)),
		),
	), sourcecode()
}

func Example_List_006() (mvc.View, string) {
	return bs.DefinitionList(
		bs.Option("Term 1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam"),
		bs.Option("Term 2", "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."),
		bs.Option("Term 3", "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
	), sourcecode()
}
