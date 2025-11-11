package main

import (
	"fmt"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func main() {
	doc := dom.GetWindow().Document()
	body := doc.Body()
	h1 := mvc.HTML("h1", mvc.WithClass("text-center", "my-4"))
	h1.Prepend(doc.CreateTextNode("hello, world!"))
	body.Prepend(h1)

	// Print the body
	fmt.Println(doc.Body())
}
