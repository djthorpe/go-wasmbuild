package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestSpan_Basic(t *testing.T) {
	span := bs.Span()
	assert.NotNil(t, span)
	assert.NotNil(t, span.Element())
}

func TestSpan_TagName(t *testing.T) {
	span := bs.Span()
	element := span.Element()
	assert.Equal(t, "SPAN", element.TagName())
}

func TestSpan_AppendText(t *testing.T) {
	span := bs.Span().Append("Hello, World!")
	element := span.Element()
	assert.Equal(t, "Hello, World!", element.TextContent())
}

func TestSpan_AppendMultipleTexts(t *testing.T) {
	span := bs.Span().Append("Hello", ", ", "World", "!")
	element := span.Element()
	assert.Equal(t, "Hello, World!", element.TextContent())
}

func TestSpan_WithClass(t *testing.T) {
	span := bs.Span(bs.WithClass("text-primary", "fw-bold"))
	classList := span.Element().ClassList()
	assert.True(t, classList.Contains("text-primary"))
	assert.True(t, classList.Contains("fw-bold"))
}

func TestSpan_WithMargin(t *testing.T) {
	span := bs.Span(bs.WithMargin(bs.START, 2))
	classList := span.Element().ClassList()
	assert.True(t, classList.Contains("ms-2"))
}

func TestSpan_WithPadding(t *testing.T) {
	span := bs.Span(bs.WithPadding(bs.PaddingAll, 3))
	classList := span.Element().ClassList()
	assert.True(t, classList.Contains("p-3"))
}

func TestSpan_OuterHTML(t *testing.T) {
	span := bs.Span().Append("Test")
	outerHTML := span.Element().OuterHTML()
	assert.Equal(t, `<span data-component="span">test</span>`, strings.ToLower(outerHTML))
}

func TestSpan_ComponentInterface(t *testing.T) {
	span := bs.Span()
	var component dom.Component = span
	assert.NotNil(t, component)
	assert.NotNil(t, component.Element())
}

func TestSpan_ChainedAppends(t *testing.T) {
	span := bs.Span().Append("First").Append(" ").Append("Second")
	assert.Equal(t, "First Second", span.Element().TextContent())
}

func TestSpan_WithBadge(t *testing.T) {
	// Test nesting a badge inside a span
	span := bs.Span().Append("Notifications ").Append(
		bs.Badge(bs.WithColor(bs.PRIMARY)).Append("5"),
	)

	element := span.Element()
	assert.True(t, element.HasChildNodes())
	assert.Contains(t, element.TextContent(), "Notifications")
	assert.Contains(t, element.TextContent(), "5")
}
