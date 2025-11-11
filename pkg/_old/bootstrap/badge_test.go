package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	assert "github.com/stretchr/testify/assert"
)

func TestBadge_Basic(t *testing.T) {
	badge := bs.Badge()

	assert.NotNil(t, badge, "Badge should not be nil")
	assert.NotNil(t, badge.Element(), "Badge element should not be nil")
	assert.Equal(t, dom.ELEMENT_NODE, badge.Element().NodeType(), "Badge should be an element node")
	assert.Equal(t, "SPAN", badge.Element().TagName(), "Badge should be a span element")
}

func TestBadge_DefaultClass(t *testing.T) {
	badge := bs.Badge()
	element := badge.Element()

	assert.True(t, element.HasAttribute("class"), "Badge should have class attribute")
	assert.Equal(t, "badge", element.GetAttribute("class"), "Badge should have 'badge' class")

	classList := element.ClassList()
	assert.NotNil(t, classList, "Badge should have class list")
	assert.Equal(t, 1, classList.Length(), "Badge should have exactly one class")
	assert.True(t, classList.Contains("badge"), "Badge should contain 'badge' class")
}

func TestBadge_WithBackground(t *testing.T) {
	tests := []struct {
		name          string
		color         bs.Color
		expectedClass string
	}{
		{"primary badge", bs.PRIMARY, "text-bg-primary"},
		{"secondary badge", bs.SECONDARY, "text-bg-secondary"},
		{"success badge", bs.SUCCESS, "text-bg-success"},
		{"danger badge", bs.DANGER, "text-bg-danger"},
		{"warning badge", bs.WARNING, "text-bg-warning"},
		{"info badge", bs.INFO, "text-bg-info"},
		{"light badge", bs.LIGHT, "text-bg-light"},
		{"dark badge", bs.DARK, "text-bg-dark"},
		{"primary subtle badge", bs.PRIMARY_SUBTLE, "text-bg-primary-subtle"},
		{"danger subtle badge", bs.DANGER_SUBTLE, "text-bg-danger-subtle"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := bs.Badge(bs.WithColor(tt.color))
			element := badge.Element()

			classList := element.ClassList()
			assert.True(t, classList.Contains("badge"), "Badge should contain 'badge' class")
			assert.True(t, classList.Contains(tt.expectedClass), "Badge should contain '%s' class", tt.expectedClass)
		})
	}
}

func TestBadge_OuterHTML(t *testing.T) {
	tests := []struct {
		name         string
		constructor  func() dom.Component
		expectedHTML string
	}{
		{
			name:         "default badge",
			constructor:  func() dom.Component { return bs.Badge() },
			expectedHTML: `<span class="badge" data-component="badge"></span>`,
		},
		{
			name:         "primary badge",
			constructor:  func() dom.Component { return bs.Badge(bs.WithColor(bs.PRIMARY)) },
			expectedHTML: `<span class="badge text-bg-primary" data-component="badge"></span>`,
		},
		{
			name:         "danger badge",
			constructor:  func() dom.Component { return bs.Badge(bs.WithColor(bs.DANGER)) },
			expectedHTML: `<span class="badge text-bg-danger" data-component="badge"></span>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := tt.constructor()
			outerHTML := badge.Element().OuterHTML()
			// Normalize to lowercase for comparison
			assert.Equal(t, tt.expectedHTML, strings.ToLower(outerHTML), "Outer HTML should match expected format")
		})
	}
}

func TestBadge_WithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
	}{
		{
			name:            "badge with single additional class",
			options:         []bs.Opt{bs.WithClass("my-custom-class")},
			expectedClasses: []string{"badge", "my-custom-class"},
		},
		{
			name:            "badge with multiple additional classes",
			options:         []bs.Opt{bs.WithClass("class1", "class2", "class3")},
			expectedClasses: []string{"badge", "class1", "class2", "class3"},
		},
		{
			name:            "primary badge with additional classes",
			options:         []bs.Opt{bs.WithColor(bs.PRIMARY), bs.WithClass("custom-badge")},
			expectedClasses: []string{"badge", "text-bg-primary", "custom-badge"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := bs.Badge(tt.options...)
			element := badge.Element()
			classList := element.ClassList()

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Badge should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Badge should have exactly %d classes, got %d", len(tt.expectedClasses), classList.Length())
		})
	}
}

func TestBadge_ComponentInterface(t *testing.T) {
	badge := bs.Badge()

	// Test that badge implements Component interface
	var component dom.Component = badge
	assert.NotNil(t, component, "Badge should implement Component interface")
	assert.NotNil(t, component.Element(), "Component Element() should return an element")
}

func TestBadge_AppendText(t *testing.T) {
	badge := bs.Badge()
	badge.Append("Hello")

	element := badge.Element()
	assert.Equal(t, "Hello", element.TextContent(), "Badge should contain text 'Hello'")
}

func TestBadge_AppendMultipleTextNodes(t *testing.T) {
	badge := bs.Badge()
	badge.Append("Hello", " ", "World")

	element := badge.Element()
	assert.Equal(t, "Hello World", element.TextContent(), "Badge should contain text 'Hello World'")
}

func TestBadge_ChainedAppends(t *testing.T) {
	badge := bs.Badge().Append("Badge").Append(" ").Append("Text")

	element := badge.Element()
	assert.Equal(t, "Badge Text", element.TextContent(), "Badge should contain text 'Badge Text'")
}

func TestBadge_ElementProperties(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "tag name is SPAN",
			test: func(t *testing.T) {
				badge := bs.Badge()
				assert.Equal(t, "SPAN", badge.Element().TagName())
			},
		},
		{
			name: "node type is element",
			test: func(t *testing.T) {
				badge := bs.Badge()
				assert.Equal(t, dom.ELEMENT_NODE, badge.Element().NodeType())
			},
		},
		{
			name: "node name is SPAN",
			test: func(t *testing.T) {
				badge := bs.Badge()
				assert.Equal(t, "SPAN", badge.Element().NodeName())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestBadge_ClassListOperations(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "initial state",
			test: func(t *testing.T) {
				badge := bs.Badge()
				classList := badge.Element().ClassList()
				assert.Equal(t, 1, classList.Length())
				assert.True(t, classList.Contains("badge"))
			},
		},
		{
			name: "add classes",
			test: func(t *testing.T) {
				badge := bs.Badge()
				classList := badge.Element().ClassList()
				classList.Add("new-class", "another-class")
				assert.True(t, classList.Contains("badge"))
				assert.True(t, classList.Contains("new-class"))
				assert.True(t, classList.Contains("another-class"))
				assert.Equal(t, 3, classList.Length())
			},
		},
		{
			name: "remove classes",
			test: func(t *testing.T) {
				badge := bs.Badge(bs.WithClass("removable"))
				classList := badge.Element().ClassList()
				assert.True(t, classList.Contains("removable"))
				classList.Remove("removable")
				assert.False(t, classList.Contains("removable"))
				assert.True(t, classList.Contains("badge"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestBadge_MultipleBadges(t *testing.T) {
	badge1 := bs.Badge(bs.WithColor(bs.PRIMARY)).Append("Primary")
	badge2 := bs.Badge(bs.WithColor(bs.DANGER)).Append("Danger")
	badge3 := bs.Badge(bs.WithColor(bs.SUCCESS)).Append("Success")

	assert.Equal(t, "Primary", badge1.Element().TextContent())
	assert.Equal(t, "Danger", badge2.Element().TextContent())
	assert.Equal(t, "Success", badge3.Element().TextContent())

	assert.True(t, badge1.Element().ClassList().Contains("text-bg-primary"))
	assert.True(t, badge2.Element().ClassList().Contains("text-bg-danger"))
	assert.True(t, badge3.Element().ClassList().Contains("text-bg-success"))
}

func TestBadge_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "empty options",
			test: func(t *testing.T) {
				badge := bs.Badge()
				assert.NotNil(t, badge)
				assert.Equal(t, "badge", badge.Element().GetAttribute("class"))
			},
		},
		{
			name: "duplicate classes",
			test: func(t *testing.T) {
				badge := bs.Badge(bs.WithClass("duplicate", "duplicate"))
				classList := badge.Element().ClassList()
				// TokenList should handle duplicates (depends on implementation)
				assert.True(t, classList.Contains("duplicate"))
			},
		},
		{
			name: "empty text content",
			test: func(t *testing.T) {
				badge := bs.Badge()
				assert.Equal(t, "", badge.Element().TextContent())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestBadge_WithBorder(t *testing.T) {
	tests := []struct {
		name            string
		position        bs.Position
		color           bs.Color
		useColor        bool
		expectedClasses []string
	}{
		{
			name:            "badge with border",
			position:        bs.BorderAll,
			useColor:        false,
			expectedClasses: []string{"badge", "border"},
		},
		{
			name:            "badge with primary border",
			position:        bs.BorderAll,
			color:           bs.PRIMARY,
			useColor:        true,
			expectedClasses: []string{"badge", "border", "border-primary"},
		},
		{
			name:            "badge with danger border and background",
			position:        bs.BorderAll,
			color:           bs.DANGER,
			useColor:        true,
			expectedClasses: []string{"badge", "border", "border-danger"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var badge dom.Component
			if tt.useColor {
				if tt.name == "badge with danger border and background" {
					badge = bs.Badge(bs.WithColor(bs.DANGER), bs.WithBorder(tt.position, tt.color))
				} else {
					badge = bs.Badge(bs.WithBorder(tt.position, tt.color))
				}
			} else {
				badge = bs.Badge(bs.WithBorder(tt.position))
			}
			classList := badge.Element().ClassList()

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Badge should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}
		})
	}
}

func TestBadge_WithPadding(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
	}{
		{
			name:            "badge with padding all",
			options:         []bs.Opt{bs.WithPadding(bs.PaddingAll, 3)},
			expectedClasses: []string{"badge", "p-3"},
		},
		{
			name:            "badge with vertical padding",
			options:         []bs.Opt{bs.WithPadding(bs.TOP|bs.BOTTOM, 2)},
			expectedClasses: []string{"badge", "py-2"},
		},
		{
			name:            "badge with horizontal padding",
			options:         []bs.Opt{bs.WithPadding(bs.START|bs.END, 4)},
			expectedClasses: []string{"badge", "px-4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := bs.Badge(tt.options...)
			classList := badge.Element().ClassList()

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Badge should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}
		})
	}
}

func TestBadge_ComplexCombination(t *testing.T) {
	badge := bs.Badge(
		bs.WithColor(bs.PRIMARY),
		bs.WithBorder(bs.BorderAll, bs.PRIMARY),
		bs.WithPadding(bs.PaddingAll, 2),
		bs.WithClass("custom-badge", "rounded"),
	).Append("Complex Badge")

	classList := badge.Element().ClassList()
	expectedClasses := []string{"badge", "text-bg-primary", "border", "border-primary", "p-2", "custom-badge", "rounded"}

	for _, expectedClass := range expectedClasses {
		assert.True(t, classList.Contains(expectedClass),
			"Badge should contain class '%s', actual classes: %v", expectedClass, classList.Values())
	}

	assert.Equal(t, "Complex Badge", badge.Element().TextContent())
}

func TestBadge_RoundedPill(t *testing.T) {
	// Bootstrap badges can be rounded pills with "rounded-pill" class
	badge := bs.PillBadge(bs.WithColor(bs.SUCCESS)).Append("Pill Badge")

	classList := badge.Element().ClassList()
	assert.True(t, classList.Contains("badge"))
	assert.True(t, classList.Contains("text-bg-success"))
	assert.True(t, classList.Contains("rounded-pill"))
	assert.Equal(t, "Pill Badge", badge.Element().TextContent())
}

func TestBadge_PillBadgeConstructor(t *testing.T) {
	tests := []struct {
		name            string
		color           bs.Color
		text            string
		expectedClasses []string
	}{
		{
			name:            "primary pill badge",
			color:           bs.PRIMARY,
			text:            "New",
			expectedClasses: []string{"badge", "text-bg-primary", "rounded-pill"},
		},
		{
			name:            "danger pill badge",
			color:           bs.DANGER,
			text:            "99+",
			expectedClasses: []string{"badge", "text-bg-danger", "rounded-pill"},
		},
		{
			name:            "success pill badge",
			color:           bs.SUCCESS,
			text:            "Available",
			expectedClasses: []string{"badge", "text-bg-success", "rounded-pill"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := bs.PillBadge(bs.WithColor(tt.color)).Append(tt.text)

			classList := badge.Element().ClassList()
			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"PillBadge should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}

			assert.Equal(t, tt.text, badge.Element().TextContent())
		})
	}
}

func TestBadge_WithColorExamples(t *testing.T) {
	// Test examples from Bootstrap documentation
	tests := []struct {
		name          string
		color         bs.Color
		text          string
		expectedClass string
	}{
		{"Primary", bs.PRIMARY, "Primary", "text-bg-primary"},
		{"Secondary", bs.SECONDARY, "Secondary", "text-bg-secondary"},
		{"Success", bs.SUCCESS, "Success", "text-bg-success"},
		{"Danger", bs.DANGER, "Danger", "text-bg-danger"},
		{"Warning", bs.WARNING, "Warning", "text-bg-warning"},
		{"Info", bs.INFO, "Info", "text-bg-info"},
		{"Light", bs.LIGHT, "Light", "text-bg-light"},
		{"Dark", bs.DARK, "Dark", "text-bg-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badge := bs.Badge(bs.WithColor(tt.color)).Append(tt.text)

			classList := badge.Element().ClassList()
			assert.True(t, classList.Contains("badge"))
			assert.True(t, classList.Contains(tt.expectedClass))
			assert.Equal(t, tt.text, badge.Element().TextContent())
		})
	}
}
