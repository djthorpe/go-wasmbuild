package main

import (
	// Packages
	carbon "github.com/djthorpe/go-wasmbuild/pkg/carbon"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
	buttonstories "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/button"
	iconstories "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/icon"
	navigationstories "github.com/djthorpe/go-wasmbuild/wasm/carbon-app/navigation"
)

func main() {
	children := []any{
		carbon.Head(1, "Carbon App"),
	}
	for _, story := range buttonstories.Stories() {
		children = append(children, story)
	}
	for _, story := range iconstories.Stories() {
		children = append(children, story)
	}
	for _, story := range navigationstories.Stories() {
		children = append(children, story)
	}

	sec := carbon.Section(children...)
	mvc.New(sec).Run()
}
