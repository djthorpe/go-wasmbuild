package mvc

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func HTML(tagName string, opts ...Opt) dom.Element {
	e := elementFactory(tagName)
	if len(opts) > 0 {
		if err := applyOpts(e, opts...); err != nil {
			panic(err)
		}
	}
	return e
}
