package mvc

import (
	"regexp"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	reTagName = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9:_-]*$`)
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// HTML returns an element with the given tag name and class/attribute options
func HTML(tagName string, args ...any) dom.Element {
	var root, e dom.Element
	if reTagName.MatchString(tagName) {
		root = elementFactory(tagName)
	} else {
		root = elementFactory("DIV")
		root.SetInnerHTML(tagName)
		if root = root.FirstElementChild(); root == nil {
			panic("HTML: invalid tag name" + tagName)
		}
	}

	// Separate options and content
	opts, content := gatherOpts(args...)

	// Apply options to root element
	if err := applyOpts(root, opts...); err != nil {
		panic(err)
	}

	// Set the content appending element to a leaf node
	e = root
	for {
		if c := e.FirstElementChild(); c == nil {
			break
		} else {
			e = c
		}
	}

	// Append content
	for _, c := range content {
		node := NodeFromAny(c)
		if node != nil {
			e.AppendChild(node)
		}
	}

	// Return element
	return root
}

// CData returns a text node with the given text
func Text(text string) dom.Text {
	return doc.CreateTextNode(text)
}

// Placeholder returns a placeholder element which is not rendered and has no effect
func Placeholder(args ...any) dom.Element {
	return HTML("SCRIPT", args...)
}
