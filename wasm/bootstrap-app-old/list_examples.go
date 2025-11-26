package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func ListExamples() mvc.View {
	return bs.Container(
		mvc.WithClass("my-4"),
		bs.Heading(2, "List Examples"), bs.HRule(),
		bs.Heading(3, "Unstyled List", mvc.WithClass("mt-5")), Example(Example_List_001),
		bs.Heading(3, "Ordered List", mvc.WithClass("mt-5")), Example(Example_List_002),
		bs.Heading(3, "Bulleted List", mvc.WithClass("mt-5")), Example(Example_List_003),
		bs.Heading(3, "List Group", mvc.WithClass("mt-5")), Example(Example_List_004),
		bs.Heading(3, "Definition List", mvc.WithClass("mt-5")), Example(Example_List_005),
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
	return bs.DefinitionList(
		bs.Option("Term 1", "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam"),
		bs.Option("Term 2", "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur."),
		bs.Option("Term 3", "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."),
	), sourcecode()
}
