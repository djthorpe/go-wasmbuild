package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestPara_Basic(t *testing.T) {
	para := bs.Para()
	assert.NotNil(t, para)
	assert.NotNil(t, para.Element())
}

func TestPara_TagName(t *testing.T) {
	para := bs.Para()
	element := para.Element()
	assert.Equal(t, "P", element.TagName())
}

func TestPara_AppendText(t *testing.T) {
	para := bs.Para().Append("This is a paragraph.")
	element := para.Element()
	assert.Equal(t, "This is a paragraph.", element.TextContent())
}

func TestPara_AppendMultipleTexts(t *testing.T) {
	para := bs.Para().Append("First sentence. ", "Second sentence.")
	element := para.Element()
	assert.Equal(t, "First sentence. Second sentence.", element.TextContent())
}

func TestPara_WithClass(t *testing.T) {
	para := bs.Para(bs.WithClass("lead", "text-muted"))
	classList := para.Element().ClassList()
	assert.True(t, classList.Contains("lead"))
	assert.True(t, classList.Contains("text-muted"))
}

func TestPara_WithMargin(t *testing.T) {
	para := bs.Para(bs.WithMargin(bs.BOTTOM, 0))
	classList := para.Element().ClassList()
	assert.True(t, classList.Contains("mb-0"))
}

func TestPara_WithPadding(t *testing.T) {
	para := bs.Para(bs.WithPadding(bs.PaddingAll, 2))
	classList := para.Element().ClassList()
	assert.True(t, classList.Contains("p-2"))
}

func TestPara_OuterHTML(t *testing.T) {
	para := bs.Para().Append("Test paragraph")
	outerHTML := para.Element().OuterHTML()
	assert.Equal(t, "<p data-component=\"para\">test paragraph</p>", strings.ToLower(outerHTML))
}

func TestPara_ComponentInterface(t *testing.T) {
	para := bs.Para()
	var component dom.Component = para
	assert.NotNil(t, component)
	assert.NotNil(t, component.Element())
}

func TestPara_ChainedAppends(t *testing.T) {
	para := bs.Para().Append("First part. ").Append("Second part.")
	assert.Equal(t, "First part. Second part.", para.Element().TextContent())
}

func TestPara_WithNestedSpan(t *testing.T) {
	// Test nesting a span inside a paragraph
	para := bs.Para().Append("This is ").Append(
		bs.Span(bs.WithClass("fw-bold")).Append("important"),
	).Append(" text.")

	element := para.Element()
	assert.True(t, element.HasChildNodes())
	assert.Contains(t, element.TextContent(), "This is important text.")
}

func TestPara_MultipleOptions(t *testing.T) {
	para := bs.Para(
		bs.WithClass("lead"),
		bs.WithMargin(bs.BOTTOM, 3),
		bs.WithPadding(bs.TOP, 2),
	)

	classList := para.Element().ClassList()
	assert.True(t, classList.Contains("lead"))
	assert.True(t, classList.Contains("mb-3"))
	assert.True(t, classList.Contains("pt-2"))
}
