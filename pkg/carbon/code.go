package carbon

import (
	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	js "github.com/djthorpe/go-wasmbuild/pkg/js"
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type codeSnippet struct{ base }

var _ mvc.View = (*codeSnippet)(nil)
var _ mvc.EnabledState = (*codeSnippet)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func init() {
	mvc.RegisterView(ViewCodeSnippet, func(element dom.Element) mvc.View {
		return mvc.NewViewWithElement(new(codeSnippet), element, setView)
	})
}

// Code returns a <cds-code-snippet type="inline"> suitable for embedding short
// code fragments within a sentence.
//
//	carbon.Code("go build ./...")
func Code(args ...any) *codeSnippet {
	snippet := mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "inline"), args).(*codeSnippet)
	snippet.syncPresentation()
	return snippet
}

// CodeSnippet returns a <cds-code-snippet type="single"> — a one-line code
// block with a copy button and horizontal scroll on overflow.
//
//	carbon.CodeSnippet("GOOS=js GOARCH=wasm go build .")
func CodeSnippet(args ...any) *codeSnippet {
	snippet := mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "single"), args).(*codeSnippet)
	snippet.syncPresentation()
	return snippet
}

// CodeBlock returns a <cds-code-snippet type="multi"> — a multi-line code
// block that collapses long content behind a "Show more" button.
//
//	carbon.CodeBlock("line1\nline2\nline3")
func CodeBlock(args ...any) *codeSnippet {
	snippet := mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "multi"), args).(*codeSnippet)
	snippet.syncPresentation()
	return snippet
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (c *codeSnippet) Apply(opts ...mvc.Opt) mvc.View {
	c.View.Apply(opts...)
	c.syncPresentation()
	return c
}

///////////////////////////////////////////////////////////////////////////////
// ENABLED STATE

// Enabled reports whether the snippet's copy button is active.
func (c *codeSnippet) Enabled() bool {
	return c.Root().GetAttribute("disabled") != "true"
}

// SetEnabled enables or disables the copy button.
func (c *codeSnippet) SetEnabled(enabled bool) mvc.View {
	if enabled {
		c.Root().RemoveAttribute("disabled")
	} else {
		c.Root().SetAttribute("disabled", "true")
	}
	return c
}

// WithCodeFeedback returns an option that overrides the temporary copied
// feedback message shown by the snippet. When empty, the component default is
// used.
func WithCodeFeedback(msg string) mvc.Opt {
	if msg == "" {
		return mvc.WithoutAttr("feedback")
	}
	return mvc.WithAttr("feedback", msg)
}

// WithCodeCopyText returns an option that overrides the text placed on the
// clipboard. When empty, the component copies its own visible content.
func WithCodeCopyText(text string) mvc.Opt {
	if text == "" {
		return mvc.WithoutAttr("copy-text")
	}
	return mvc.WithAttr("copy-text", text)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (c *codeSnippet) syncPresentation() {
	root := c.Root()
	if node, ok := root.JSValue().(js.Value); ok && !node.IsUndefined() && !node.IsNull() {
		node.Set("wrapText", root.HasAttribute(string(CodeWrapText)))
	}
	setTagBoolProperty(root, string(CodeHideCopyButton), root.HasAttribute(string(CodeHideCopyButton)))
}
