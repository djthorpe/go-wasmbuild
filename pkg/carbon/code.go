package carbon

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
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewCode, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(codeblock), element, func(self, child mvc.View) {
			self.(*codeblock).View = child
		})
	})
}

// CodeInline returns a <cds-code-snippet type="inline"> for embedding a short
// code fragment within a sentence or paragraph.
func CodeInline(args ...any) *codeblock {
	return mvc.NewView(new(codeblock), ViewCode, "cds-code-snippet", func(self, child mvc.View) {
		self.(*codeblock).View = child
	}, mvc.WithAttr("type", "inline"), args).(*codeblock)
}

// CodeSingle returns a <cds-code-snippet type="single"> for a single-line
// code snippet with a copy button.
func CodeSingle(args ...any) *codeblock {
	return mvc.NewView(new(codeblock), ViewCode, "cds-code-snippet", func(self, child mvc.View) {
		self.(*codeblock).View = child
	}, mvc.WithAttr("type", "single"), args).(*codeblock)
}

// CodeMulti returns a <cds-code-snippet type="multi"> for a multi-line
// code block with a copy button and optional expand/collapse.
func CodeMulti(args ...any) *codeblock {
	return mvc.NewView(new(codeblock), ViewCode, "cds-code-snippet", func(self, child mvc.View) {
		self.(*codeblock).View = child
	}, mvc.WithAttr("type", "multi"), args).(*codeblock)
}
