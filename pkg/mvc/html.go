package mvc

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// HTML returns an element with the given tag name and class/attribute options
func HTML(tagName string, args ...any) dom.Element {
	e := elementFactory(tagName)

	// Separate options and content
	opts, content := gatherOpts(args...)

	// Apply options
	if err := applyOpts(e, opts...); err != nil {
		panic(err)
	}

	// Append content
	for _, c := range content {
		node := NodeFromAny(c)
		if node != nil {
			e.AppendChild(node)
		}
	}

	// Return element
	return e
}

// CData returns a text node with the given text
func Text(text string) dom.Text {
	return doc.CreateTextNode(text)
}

// Placeholder returns a placeholder element which is not rendered and has no effect
func Placeholder(args ...any) dom.Element {
	return HTML("SCRIPT", args...)
}
