package main

import (
	"fmt"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"
	"github.com/djthorpe/go-wasmbuild/pkg/mvc"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func main() {
	// Make a new application
	app := mvc.New("Bart App")

	// Create a heading
	app.Append(
		bs.Heading(1).Append("Hello, World!"),
		Links(),
		Badges(),
		Alerts(),
	)

	// Listen for hash changes
	app.AddEventListener("hashchange", func(node Node) {
		fmt.Println(node)
	})

	// Run the application
	select {}
}
