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
	return mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "inline"), args).(*codeSnippet)
}

// CodeSnippet returns a <cds-code-snippet type="single"> — a one-line code
// block with a copy button and horizontal scroll on overflow.
//
//	carbon.CodeSnippet("GOOS=js GOARCH=wasm go build .")
func CodeSnippet(args ...any) *codeSnippet {
	return mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "single"), args).(*codeSnippet)
}

// CodeBlock returns a <cds-code-snippet type="multi"> — a multi-line code
// block that collapses long content behind a "Show more" button.
//
//	carbon.CodeBlock("line1\nline2\nline3")
func CodeBlock(args ...any) *codeSnippet {
	return mvc.NewView(new(codeSnippet), ViewCodeSnippet, "cds-code-snippet", setView,
		mvc.WithAttr("type", "multi"), args).(*codeSnippet)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

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

// SetWrapText controls whether long lines wrap instead of scroll.
// Only meaningful for single-line snippets.
func (c *codeSnippet) SetWrapText(wrap bool) *codeSnippet {
	// Set the Lit component property directly — setAttribute alone is not
	// reliable for Lit boolean properties on already-rendered elements.
	if node, ok := c.Root().JSValue().(js.Value); ok {
		node.Set("wrapText", wrap)
	}
	return c
}

// SetFeedback sets the message shown briefly after the user copies the snippet.
// Defaults to "Copied!" when unset.
func (c *codeSnippet) SetFeedback(msg string) *codeSnippet {
	if msg == "" {
		c.Root().RemoveAttribute("feedback")
	} else {
		c.Root().SetAttribute("feedback", msg)
	}
	return c
}

// SetCopyText overrides the text placed on the clipboard. When empty the
// component copies its own visible content.
func (c *codeSnippet) SetCopyText(text string) *codeSnippet {
	if text == "" {
		c.Root().RemoveAttribute("copy-text")
	} else {
		c.Root().SetAttribute("copy-text", text)
	}
	return c
}

// SetHideCopyButton hides or shows the copy-to-clipboard button.
func (c *codeSnippet) SetHideCopyButton(hide bool) *codeSnippet {
	if hide {
		c.Root().SetAttribute("hide-copy-button", "true")
	} else {
		c.Root().RemoveAttribute("hide-copy-button")
	}
	return c
}
