package bootstrap_test

import (
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	assert "github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////
// BASIC BUTTON TESTS

func TestButton_Basic(t *testing.T) {
	btn := bs.Button(bs.PRIMARY)

	assert.NotNil(t, btn, "Button should not be nil")
	assert.NotNil(t, btn.Element(), "Button element should not be nil")
	assert.Equal(t, dom.ELEMENT_NODE, btn.Element().NodeType(), "Button should be an element node")
	assert.Equal(t, "BUTTON", btn.Element().TagName(), "Button should be a button element")
}

func TestButton_DefaultClasses(t *testing.T) {
	btn := bs.Button(bs.PRIMARY)
	element := btn.Element()

	assert.True(t, element.HasAttribute("class"), "Button should have class attribute")

	classList := element.ClassList()
	assert.NotNil(t, classList, "Button should have class list")
	assert.True(t, classList.Contains("btn"), "Button should contain 'btn' class")
	assert.True(t, classList.Contains("btn-primary"), "Button should contain 'btn-primary' class")
}

func TestButton_DefaultType(t *testing.T) {
	btn := bs.Button(bs.PRIMARY)
	element := btn.Element()

	assert.True(t, element.HasAttribute("type"), "Button should have type attribute")
	assert.Equal(t, "button", element.GetAttribute("type"), "Button should have type='button'")
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON VARIANTS (Bootstrap 5.3 Documentation)

func TestButton_AllColorVariants(t *testing.T) {
	tests := []struct {
		name          string
		color         bs.Color
		expectedClass string
	}{
		{"primary button", bs.PRIMARY, "btn-primary"},
		{"secondary button", bs.SECONDARY, "btn-secondary"},
		{"success button", bs.SUCCESS, "btn-success"},
		{"danger button", bs.DANGER, "btn-danger"},
		{"warning button", bs.WARNING, "btn-warning"},
		{"info button", bs.INFO, "btn-info"},
		{"light button", bs.LIGHT, "btn-light"},
		{"dark button", bs.DARK, "btn-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			btn := bs.Button(tt.color)
			element := btn.Element()

			classList := element.ClassList()
			assert.True(t, classList.Contains("btn"), "Button should contain 'btn' class")
			assert.True(t, classList.Contains(tt.expectedClass), "Button should contain '%s' class", tt.expectedClass)
		})
	}
}

func TestButton_OuterHTML(t *testing.T) {
	tests := []struct {
		name          string
		constructor   func() dom.Component
		expectedClass string
	}{
		{
			name:          "default button",
			constructor:   func() dom.Component { return bs.Button(bs.PRIMARY) },
			expectedClass: "btn btn-primary",
		},
		{
			name:          "secondary button",
			constructor:   func() dom.Component { return bs.Button(bs.SECONDARY) },
			expectedClass: "btn btn-secondary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := tt.constructor()
			element := button.Element()
			// Check tag name
			assert.Equal(t, "BUTTON", element.TagName())
			// Check class attribute
			assert.Equal(t, tt.expectedClass, element.GetAttribute("class"))
			// Check type attribute
			assert.Equal(t, "button", element.GetAttribute("type"))
		})
	}
}

func TestButton_OutlineButton(t *testing.T) {
	tests := []struct {
		name          string
		color         bs.Color
		expectedClass string
	}{
		{"outline primary", bs.PRIMARY, "btn-outline-primary"},
		{"outline secondary", bs.SECONDARY, "btn-outline-secondary"},
		{"outline success", bs.SUCCESS, "btn-outline-success"},
		{"outline danger", bs.DANGER, "btn-outline-danger"},
		{"outline warning", bs.WARNING, "btn-outline-warning"},
		{"outline info", bs.INFO, "btn-outline-info"},
		{"outline light", bs.LIGHT, "btn-outline-light"},
		{"outline dark", bs.DARK, "btn-outline-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := bs.OutlineButton(tt.color)
			classList := button.Element().ClassList()
			assert.True(t, classList.Contains("btn"))
			assert.True(t, classList.Contains(tt.expectedClass))
		})
	}
}

func TestButton_WithSize(t *testing.T) {
	tests := []struct {
		name          string
		size          bs.Size
		expectedClass string
	}{
		{"small button", bs.SizeSmall, "btn-sm"},
		{"large button", bs.SizeLarge, "btn-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			button := bs.Button(bs.PRIMARY, bs.WithSize(tt.size))
			classList := button.Element().ClassList()
			assert.True(t, classList.Contains("btn"))
			assert.True(t, classList.Contains("btn-primary"))
			assert.True(t, classList.Contains(tt.expectedClass))
		})
	}
}

func TestButton_WithAdditionalClasses(t *testing.T) {
	button := bs.Button(bs.PRIMARY, bs.WithClass("custom-class", "another-class"))
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn"))
	assert.True(t, classList.Contains("custom-class"))
	assert.True(t, classList.Contains("another-class"))
}

func TestButton_ComponentInterface(t *testing.T) {
	button := bs.Button(bs.PRIMARY)
	var component dom.Component = button
	assert.NotNil(t, component)
	assert.NotNil(t, component.Element())
}

func TestButton_ChainedAppends(t *testing.T) {
	button := bs.Button(bs.PRIMARY).Append("Click ").Append("me")
	assert.Equal(t, "Click me", button.Element().TextContent())
}

func TestButton_WithBadge(t *testing.T) {
	// Test nesting a badge inside a button
	button := bs.Button(bs.PRIMARY).Append(
		"Notifications ",
		bs.Badge(bs.WithColor(bs.LIGHT)).Append("4"),
	)

	element := button.Element()
	assert.True(t, element.HasChildNodes())
	assert.Contains(t, element.TextContent(), "Notifications")
	assert.Contains(t, element.TextContent(), "4")
}

func TestButton_WithMargin(t *testing.T) {
	button := bs.Button(bs.PRIMARY, bs.WithMargin(bs.START, 2))
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn"))
	assert.True(t, classList.Contains("ms-2"))
}

func TestButton_WithPadding(t *testing.T) {
	button := bs.Button(bs.PRIMARY, bs.WithPadding(bs.PaddingAll, 3))
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn"))
	assert.True(t, classList.Contains("p-3"))
}

func TestButton_ComplexCombination(t *testing.T) {
	button := bs.Button(
		bs.SUCCESS,
		bs.WithSize(bs.SizeLarge),
		bs.WithMargin(bs.END, 2),
		bs.WithClass("shadow"),
	).Append("Large Success Button")

	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn"))
	assert.True(t, classList.Contains("btn-success"))
	assert.True(t, classList.Contains("btn-lg"))
	assert.True(t, classList.Contains("me-2"))
	assert.True(t, classList.Contains("shadow"))
	assert.Equal(t, "Large Success Button", button.Element().TextContent())
}

func TestButton_OutlineWithSize(t *testing.T) {
	button := bs.OutlineButton(bs.DANGER, bs.WithSize(bs.SizeSmall))
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn"))
	assert.True(t, classList.Contains("btn-outline-danger"))
	assert.True(t, classList.Contains("btn-sm"))
}

func TestCloseButton_Basic(t *testing.T) {
	button := bs.CloseButton()
	assert.NotNil(t, button)
	assert.NotNil(t, button.Element())
}

func TestCloseButton_TagName(t *testing.T) {
	button := bs.CloseButton()
	element := button.Element()
	assert.Equal(t, "BUTTON", element.TagName())
}

func TestCloseButton_DefaultClass(t *testing.T) {
	button := bs.CloseButton()
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn-close"))
}

func TestCloseButton_DefaultType(t *testing.T) {
	button := bs.CloseButton()
	element := button.Element()
	assert.Equal(t, "button", element.GetAttribute("type"))
}

func TestCloseButton_DefaultAriaLabel(t *testing.T) {
	button := bs.CloseButton()
	element := button.Element()
	assert.Equal(t, "Close", element.GetAttribute("aria-label"))
}

func TestCloseButton_WithCustomAriaLabel(t *testing.T) {
	button := bs.CloseButton(bs.WithAttribute("aria-label", "Dismiss"))
	element := button.Element()
	assert.Equal(t, "Dismiss", element.GetAttribute("aria-label"))
}

func TestCloseButton_WithDismissAttribute(t *testing.T) {
	button := bs.CloseButton(bs.WithAttribute("data-bs-dismiss", "modal"))
	element := button.Element()
	assert.Equal(t, "modal", element.GetAttribute("data-bs-dismiss"))
	assert.True(t, element.ClassList().Contains("btn-close"))
}

func TestCloseButton_WithAdditionalClass(t *testing.T) {
	button := bs.CloseButton(bs.WithClass("btn-close-white"))
	classList := button.Element().ClassList()
	assert.True(t, classList.Contains("btn-close"))
	assert.True(t, classList.Contains("btn-close-white"))
}

func TestCloseButton_OuterHTML(t *testing.T) {
	button := bs.CloseButton()
	html := button.Element().OuterHTML()
	assert.Contains(t, html, `class="btn-close"`)
	assert.Contains(t, html, `type="button"`)
	assert.Contains(t, html, `aria-label="Close"`)
	assert.Contains(t, html, `data-component="button"`)
}

///////////////////////////////////////////////////////////////////////////////
// ButtonGroup Tests

func TestButtonGroup_Basic(t *testing.T) {
	bg := bs.ButtonGroup()
	assert.NotNil(t, bg)
}

func TestButtonGroup_TagName(t *testing.T) {
	bg := bs.ButtonGroup()
	assert.Equal(t, "DIV", bg.Element().TagName())
}

func TestButtonGroup_DefaultClass(t *testing.T) {
	bg := bs.ButtonGroup()
	classList := bg.Element().ClassList()
	assert.True(t, classList.Contains("btn-group"))
}

func TestButtonGroup_DefaultRole(t *testing.T) {
	bg := bs.ButtonGroup()
	role := bg.Element().GetAttribute("role")
	assert.Equal(t, "group", role)
}

func TestButtonGroup_WithAdditionalClasses(t *testing.T) {
	bg := bs.ButtonGroup(bs.WithClass("custom-class"))
	classList := bg.Element().ClassList()
	assert.True(t, classList.Contains("btn-group"))
	assert.True(t, classList.Contains("custom-class"))
}

func TestButtonGroup_OuterHTML(t *testing.T) {
	bg := bs.ButtonGroup()
	html := bg.Element().OuterHTML()
	assert.Contains(t, html, `class="btn-group"`)
	assert.Contains(t, html, `role="group"`)
	assert.Contains(t, html, `data-component="button-group"`)
}

func TestButtonGroup_ComponentInterface(t *testing.T) {
	bg := bs.ButtonGroup()
	assert.Implements(t, (*dom.Component)(nil), bg)
}

///////////////////////////////////////////////////////////////////////////////
// VerticalButtonGroup Tests

func TestVerticalButtonGroup_Basic(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	assert.NotNil(t, bg)
}

func TestVerticalButtonGroup_TagName(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	assert.Equal(t, "DIV", bg.Element().TagName())
}

func TestVerticalButtonGroup_DefaultClass(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	classList := bg.Element().ClassList()
	assert.True(t, classList.Contains("btn-group-vertical"))
}

func TestVerticalButtonGroup_DefaultRole(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	role := bg.Element().GetAttribute("role")
	assert.Equal(t, "group", role)
}

func TestVerticalButtonGroup_WithAdditionalClasses(t *testing.T) {
	bg := bs.VerticalButtonGroup(bs.WithClass("custom-vertical"))
	classList := bg.Element().ClassList()
	assert.True(t, classList.Contains("btn-group-vertical"))
	assert.True(t, classList.Contains("custom-vertical"))
}

func TestVerticalButtonGroup_OuterHTML(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	html := bg.Element().OuterHTML()
	assert.Contains(t, html, `class="btn-group-vertical"`)
	assert.Contains(t, html, `role="group"`)
	assert.Contains(t, html, `data-component="button-group"`)
}

func TestVerticalButtonGroup_ComponentInterface(t *testing.T) {
	bg := bs.VerticalButtonGroup()
	assert.Implements(t, (*dom.Component)(nil), bg)
}

///////////////////////////////////////////////////////////////////////////////
// ButtonToolbar Tests

func TestButtonToolbar_Basic(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	assert.NotNil(t, toolbar)
}

func TestButtonToolbar_TagName(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	assert.Equal(t, "DIV", toolbar.Element().TagName())
}

func TestButtonToolbar_DefaultClass(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	classList := toolbar.Element().ClassList()
	assert.True(t, classList.Contains("btn-toolbar"))
}

func TestButtonToolbar_DefaultRole(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	role := toolbar.Element().GetAttribute("role")
	assert.Equal(t, "toolbar", role)
}

func TestButtonToolbar_WithAdditionalClasses(t *testing.T) {
	toolbar := bs.ButtonToolbar(bs.WithClass("custom-toolbar"))
	classList := toolbar.Element().ClassList()
	assert.True(t, classList.Contains("btn-toolbar"))
	assert.True(t, classList.Contains("custom-toolbar"))
}

func TestButtonToolbar_OuterHTML(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	html := toolbar.Element().OuterHTML()
	assert.Contains(t, html, `class="btn-toolbar"`)
	assert.Contains(t, html, `role="toolbar"`)
	assert.Contains(t, html, `data-component="button-group"`)
}

func TestButtonToolbar_ComponentInterface(t *testing.T) {
	toolbar := bs.ButtonToolbar()
	assert.Implements(t, (*dom.Component)(nil), toolbar)
}
