package mvc

import (
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func Element(tagName string, opts ...Opt) dom.Element {
	e := elementFactory(tagName)
	if len(opts) > 0 {
		if err := applyOpts(e, opts...); err != nil {
			panic(err)
		}
	}
	return e
}
