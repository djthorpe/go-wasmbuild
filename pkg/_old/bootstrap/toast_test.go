//go:build js && wasm

package bootstrap_test

import (
	"testing"
	"time"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestToast_Basic(t *testing.T) {
	toast := bs.Toast()
	assert.NotNil(t, toast)
	assert.NotNil(t, toast.Element())
}

func TestToast_TagName(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	assert.Equal(t, "DIV", element.TagName())
}

func TestToast_HasToastClass(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	assert.True(t, element.ClassList().Contains("toast"))
}

func TestToast_HasRoleAlert(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	assert.Equal(t, "alert", element.GetAttribute("role"))
}

func TestToast_HasAriaAttributes(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	assert.Equal(t, "assertive", element.GetAttribute("aria-live"))
	assert.Equal(t, "true", element.GetAttribute("aria-atomic"))
}

func TestToast_HasHeaderAndBody(t *testing.T) {
	toast := bs.Toast().Header("Title")
	element := toast.Element()
	children := element.Children()

	assert.Equal(t, 2, len(children))
	assert.Equal(t, "DIV", children[0].TagName())
	assert.Equal(t, "DIV", children[1].TagName())
}

func TestToast_HasOnlyBody(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	children := element.Children()

	// Without Header() called, should only have body
	assert.Equal(t, 1, len(children))
	assert.Equal(t, "DIV", children[0].TagName())
}

func TestToast_HeaderHasClass(t *testing.T) {
	toast := bs.Toast().Header("Title")
	element := toast.Element()
	children := element.Children()
	header := children[0].(dom.Element)
	assert.True(t, header.ClassList().Contains("toast-header"))
}

func TestToast_BodyHasClass(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	children := element.Children()
	body := children[0].(dom.Element)
	assert.True(t, body.ClassList().Contains("toast-body"))
}

func TestToast_HeaderAppendText(t *testing.T) {
	toast := bs.Toast().Header("Test Title")
	element := toast.Element()
	children := element.Children()
	header := children[0].(dom.Element)

	assert.Equal(t, "Test Title", header.TextContent())
}

func TestToast_WithOptions(t *testing.T) {
	toast := bs.Toast(bs.WithClass("custom-toast"), bs.WithID("myToast"))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("toast"))
	assert.True(t, element.ClassList().Contains("custom-toast"))
	assert.Equal(t, "myToast", element.GetAttribute("id"))
}

func TestToast_IsComponent(t *testing.T) {
	toast := bs.Toast()
	var _ dom.Component = toast
}

func TestToast_OuterHTML(t *testing.T) {
	toast := bs.Toast()
	element := toast.Element()
	html := element.OuterHTML()

	assert.Contains(t, html, "<div")
	assert.Contains(t, html, "class=\"toast\"")
	assert.Contains(t, html, "role=\"alert\"")
	assert.Contains(t, html, "aria-live=\"assertive\"")
	assert.Contains(t, html, "aria-atomic=\"true\"")
	assert.Contains(t, html, "toast-body")
}

func TestToast_OuterHTMLWithHeader(t *testing.T) {
	toast := bs.Toast().Header("Title")
	element := toast.Element()
	html := element.OuterHTML()

	assert.Contains(t, html, "<div")
	assert.Contains(t, html, "class=\"toast\"")
	assert.Contains(t, html, "role=\"alert\"")
	assert.Contains(t, html, "aria-live=\"assertive\"")
	assert.Contains(t, html, "aria-atomic=\"true\"")
	assert.Contains(t, html, "toast-header")
	assert.Contains(t, html, "toast-body")
}

func TestToast_WithColorScheme(t *testing.T) {
	toast := bs.Toast(bs.WithClass("text-bg-primary"), bs.WithClass("border-0"))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("toast"))
	assert.True(t, element.ClassList().Contains("text-bg-primary"))
	assert.True(t, element.ClassList().Contains("border-0"))
}

func TestToast_HeaderMethodChaining(t *testing.T) {
	toast := bs.Toast().Header("Title 1").Header("Title 2")
	element := toast.Element()
	children := element.Children()
	header := children[0].(dom.Element)

	content := header.TextContent()
	assert.Contains(t, content, "Title 1")
	assert.Contains(t, content, "Title 2")
}

func TestToast_HeaderWithMultipleArgs(t *testing.T) {
	toast := bs.Toast().Header("Part 1", " ", "Part 2")
	element := toast.Element()
	children := element.Children()
	header := children[0].(dom.Element)

	content := header.TextContent()
	assert.Contains(t, content, "Part 1")
	assert.Contains(t, content, "Part 2")
}

func TestToast_HeaderWithComponent(t *testing.T) {
	span := bs.Span().Append("Bold Text")
	toast := bs.Toast().Header(span)

	element := toast.Element()
	children := element.Children()
	header := children[0].(dom.Element)
	headerChildren := header.Children()

	assert.Equal(t, 1, len(headerChildren))
	assert.Equal(t, "SPAN", headerChildren[0].TagName())
}

func TestToast_WithoutAnimation(t *testing.T) {
	toast := bs.Toast(bs.WithoutAnimation())
	element := toast.Element()

	assert.Equal(t, "false", element.GetAttribute("data-bs-animation"))
}

func TestToast_WithTimeout(t *testing.T) {
	toast := bs.Toast(bs.WithTimeout(3 * time.Second))
	element := toast.Element()

	assert.Equal(t, "true", element.GetAttribute("data-bs-autohide"))
	assert.Equal(t, "3000", element.GetAttribute("data-bs-delay"))
}

func TestToast_WithTimeoutZero(t *testing.T) {
	toast := bs.Toast(bs.WithTimeout(0))
	element := toast.Element()

	assert.Equal(t, "false", element.GetAttribute("data-bs-autohide"))
}

func TestToast_WithColor(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.SUCCESS))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("toast"))
	assert.True(t, element.ClassList().Contains("text-bg-success"))
}

func TestToast_WithColorPrimary(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.PRIMARY))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("text-bg-primary"))
}

func TestToast_WithColorDanger(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.DANGER))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("text-bg-danger"))
}

func TestToast_WithColorWarning(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.WARNING))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("text-bg-warning"))
}

func TestToast_WithColorInfo(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.INFO))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("text-bg-info"))
}

func TestToast_WithColorTransparent(t *testing.T) {
	toast := bs.Toast(bs.WithColor(bs.TRANSPARENT))
	element := toast.Element()

	assert.True(t, element.ClassList().Contains("toast"))
	assert.False(t, element.ClassList().Contains("text-bg-"))
}

func TestToast_WithColorChange(t *testing.T) {
	// Create toast with primary color
	toast := bs.Toast(bs.WithColor(bs.PRIMARY))
	element := toast.Element()
	assert.True(t, element.ClassList().Contains("text-bg-primary"))

	// Change to success color - need to create a new toast
	toast2 := bs.Toast(bs.WithColor(bs.SUCCESS))
	element2 := toast2.Element()
	assert.True(t, element2.ClassList().Contains("text-bg-success"))
	assert.False(t, element2.ClassList().Contains("text-bg-primary"))
}
