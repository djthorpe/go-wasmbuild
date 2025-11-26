package bootstrap

import (
	// Package imports
	dom "github.com/djthorpe/go-wasmbuild"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type codeblock struct {
	mvc.View
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

func init() {
	mvc.RegisterView(ViewCodeBlock, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(codeblock), element, func(self, child mvc.View) {
			self.(*codeblock).View = child
		})
	})
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func CodeBlock(args ...any) *codeblock {
	return mvc.NewView(new(codeblock), ViewCodeBlock, "PRE", func(self, child mvc.View) {
		self.(*codeblock).View = child
	}, mvc.WithClass("codeblock"), args).(*codeblock)
}
