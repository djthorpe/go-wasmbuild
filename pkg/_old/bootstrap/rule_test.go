package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestRule_Basic(t *testing.T) {
	rule := bs.Rule()
	assert.NotNil(t, rule)
	assert.NotNil(t, rule.Element())
}

func TestRule_TagName(t *testing.T) {
	rule := bs.Rule()
	element := rule.Element()
	assert.Equal(t, "HR", element.TagName())
}

func TestRule_WithClass(t *testing.T) {
	rule := bs.Rule(bs.WithClass("my-4"))
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("my-4"))
}

func TestRule_WithMargin(t *testing.T) {
	rule := bs.Rule(bs.WithMargin(bs.TOP|bs.BOTTOM, 5))
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("my-5"))
}

func TestRule_WithColor(t *testing.T) {
	// Bootstrap 5 supports colored horizontal rules
	rule := bs.Rule(bs.WithClass("border-primary"))
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("border-primary"))
}

func TestRule_OuterHTML(t *testing.T) {
	rule := bs.Rule()
	element := rule.Element()

	// Verify tag name and structure
	assert.Equal(t, "HR", element.TagName())

	// Verify OuterHTML contains the expected tag (HR is a void element, no closing tag)
	outerHTML := strings.ToLower(element.OuterHTML())
	assert.Contains(t, outerHTML, "<hr")
}

func TestRule_WithClassOuterHTML(t *testing.T) {
	rule := bs.Rule(bs.WithClass("my-4"))
	element := rule.Element()

	// Verify tag name
	assert.Equal(t, "HR", element.TagName())

	// Verify class is present
	classList := element.ClassList()
	assert.True(t, classList.Contains("my-4"), "Rule should have my-4 class")

	// Verify OuterHTML contains expected parts (HR is a void element, no closing tag)
	outerHTML := strings.ToLower(element.OuterHTML())
	assert.Contains(t, outerHTML, "<hr")
	assert.Contains(t, outerHTML, `class="my-4"`)
}

func TestRule_ComponentInterface(t *testing.T) {
	rule := bs.Rule()
	var component dom.Component = rule
	assert.NotNil(t, component)
	assert.NotNil(t, component.Element())
}

func TestRule_NoChildren(t *testing.T) {
	// HR elements typically don't have children, but the interface allows it
	rule := bs.Rule()
	element := rule.Element()
	assert.False(t, element.HasChildNodes())
}

func TestRule_MultipleOptions(t *testing.T) {
	rule := bs.Rule(
		bs.WithClass("border-primary"),
		bs.WithMargin(bs.TOP, 4),
		bs.WithMargin(bs.BOTTOM, 4),
	)

	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("border-primary"))
	assert.True(t, classList.Contains("mt-4"))
	assert.True(t, classList.Contains("mb-4"))
}

func TestRule_InContext(t *testing.T) {
	// Test rule in a typical usage context within a container
	container := bs.Container().Append(
		bs.Para().Append("First paragraph"),
		bs.Rule(bs.WithMargin(bs.TOP|bs.BOTTOM, 3)),
		bs.Para().Append("Second paragraph"),
	)

	element := container.Element()
	assert.True(t, element.HasChildNodes())
	// Should have 3 children: para, hr, para
	children := element.ChildNodes()
	assert.Equal(t, 3, len(children))
}

func TestVerticalRule_Basic(t *testing.T) {
	rule := bs.VerticalRule()
	assert.NotNil(t, rule)
	assert.NotNil(t, rule.Element())
}

func TestVerticalRule_TagName(t *testing.T) {
	rule := bs.VerticalRule()
	element := rule.Element()
	assert.Equal(t, "DIV", element.TagName())
}

func TestVerticalRule_HasVRClass(t *testing.T) {
	rule := bs.VerticalRule()
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("vr"))
}

func TestVerticalRule_OuterHTML(t *testing.T) {
	rule := bs.VerticalRule()
	element := rule.Element()

	// Verify tag name
	assert.Equal(t, "DIV", element.TagName())

	// Verify class is present
	classList := element.ClassList()
	assert.True(t, classList.Contains("vr"), "VerticalRule should have vr class")

	// Verify OuterHTML contains expected parts
	outerHTML := strings.ToLower(element.OuterHTML())
	assert.Contains(t, outerHTML, "<div")
	assert.Contains(t, outerHTML, `class="vr"`)
	assert.Contains(t, outerHTML, "</div>")
}

func TestVerticalRule_WithAdditionalClasses(t *testing.T) {
	rule := bs.VerticalRule(bs.WithClass("mx-2"))
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("vr"))
	assert.True(t, classList.Contains("mx-2"))
}

func TestVerticalRule_WithMargin(t *testing.T) {
	rule := bs.VerticalRule(bs.WithMargin(bs.START|bs.END, 2))
	classList := rule.Element().ClassList()
	assert.True(t, classList.Contains("vr"))
	assert.True(t, classList.Contains("mx-2"))
}

func TestVerticalRule_ComponentInterface(t *testing.T) {
	rule := bs.VerticalRule()
	var component dom.Component = rule
	assert.NotNil(t, component)
	assert.NotNil(t, component.Element())
}

func TestVerticalRule_InFlexContext(t *testing.T) {
	// Test vertical rule in a typical flex layout
	container := bs.Container(bs.WithClass("d-flex")).Append(
		bs.Span().Append("Item 1"),
		bs.VerticalRule(bs.WithMargin(bs.START|bs.END, 2)),
		bs.Span().Append("Item 2"),
		bs.VerticalRule(bs.WithMargin(bs.START|bs.END, 2)),
		bs.Span().Append("Item 3"),
	)

	element := container.Element()
	assert.True(t, element.HasChildNodes())
	// Should have 5 children: span, vr, span, vr, span
	children := element.ChildNodes()
	assert.Equal(t, 5, len(children))
}
