package mvc

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// HTML returns an element with the given tag name and class/attribute options
func HTML(tagName string, opts ...Opt) dom.Element {
	e := elementFactory(tagName)
	if len(opts) > 0 {
		if err := applyOpts(e, opts...); err != nil {
			panic(err)
		}
	}
	return e
}

// CData returns a text node with the given text
func CData(text string) dom.Text {
	return doc.CreateTextNode(text)
}

// Placeholder returns a placeholder element which is not rendered and has no effect
func Placeholder(opts ...Opt) dom.Element {
	return HTML("SCRIPT", opts...)
}
