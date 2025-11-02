package mvc

import (
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// tag wraps a generic HTML tag element
type tag struct {
	e Element
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Tag(tagName string, opts ...Opt) *tag {
	e := elementFactory(tagName)
	if len(opts) > 0 {
		if err := applyOpts(e, opts...); err != nil {
			panic(err)
		}
	}
	return &tag{e}
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (t *tag) Root() Element {
	return t.e
}

func (t *tag) Append(children ...any) *tag {
	for _, child := range children {
		t.e.AppendChild(NodeFromAny(child))
	}
	return t
}
