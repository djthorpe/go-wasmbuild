package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestContainer_Basic(t *testing.T) {
	container := bs.Container()

	assert.NotNil(t, container, "Container should not be nil")
	assert.NotNil(t, container.Element(), "Container element should not be nil")
	assert.Equal(t, dom.ELEMENT_NODE, container.Element().NodeType(), "Container should be an element node")
	assert.Equal(t, "DIV", container.Element().TagName(), "Container should be a div element")
	assert.Equal(t, "container", container.Name(), "Container should have correct component name")
}

func TestContainer_DefaultClass(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	assert.True(t, element.HasAttribute("class"), "Container should have class attribute")
	assert.Equal(t, "container", element.GetAttribute("class"), "Container should have 'container' class")

	classList := element.ClassList()
	assert.NotNil(t, classList, "Container should have class list")
	assert.Equal(t, 1, classList.Length(), "Container should have exactly one class")
	assert.True(t, classList.Contains("container"), "Container should contain 'container' class")
}

func TestContainer_FluidOption(t *testing.T) {
	container := bs.Container(bs.WithBreakpoint(bs.BreakpointFluid))
	element := container.Element()

	assert.True(t, element.HasAttribute("class"), "Fluid container should have class attribute")
	assert.Equal(t, "container-fluid", element.GetAttribute("class"), "Fluid container should have 'container-fluid' class")

	classList := element.ClassList()
	assert.Equal(t, 1, classList.Length(), "Fluid container should have exactly one class")
	assert.True(t, classList.Contains("container-fluid"), "Fluid container should contain 'container-fluid' class")
	assert.False(t, classList.Contains("container"), "Fluid container should not contain 'container' class")
}

func TestContainer_OuterHTML(t *testing.T) {
	tests := []struct {
		name         string
		constructor  func() dom.Component
		expectedHTML string
	}{
		{
			name:         "default container",
			constructor:  func() dom.Component { return bs.Container() },
			expectedHTML: `<div class="container" data-component="container"></div>`,
		},
		{
			name:         "fluid container",
			constructor:  func() dom.Component { return bs.Container(bs.WithBreakpoint(bs.BreakpointFluid)) },
			expectedHTML: `<div class="container-fluid" data-component="container"></div>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := tt.constructor()
			outerHTML := container.Element().OuterHTML()
			// Normalize to lowercase for comparison since different implementations may vary
			assert.Equal(t, tt.expectedHTML, strings.ToLower(outerHTML), "Outer HTML should match expected format")
		})
	}
}

func TestContainer_WithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
	}{
		{
			name:            "container with single additional class",
			options:         []bs.Opt{bs.WithClass("my-custom-class")},
			expectedClasses: []string{"container", "my-custom-class"},
		},
		{
			name:            "container with multiple additional classes",
			options:         []bs.Opt{bs.WithClass("class1", "class2", "class3")},
			expectedClasses: []string{"container", "class1", "class2", "class3"},
		},
		{
			name:            "fluid container with additional classes",
			options:         []bs.Opt{bs.WithBreakpoint(bs.BreakpointFluid), bs.WithClass("custom-fluid")},
			expectedClasses: []string{"container-fluid", "custom-fluid"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(tt.options...)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

func TestContainer_ComponentInterface(t *testing.T) {
	container := bs.Container()

	// Verify it implements the Component interface methods
	var _ dom.Component = container

	assert.NotNil(t, container.Element(), "Element() should return non-nil")
	assert.Equal(t, "container", container.Name(), "Name() should return 'container'")
}

func TestContainer_AppendText(t *testing.T) {
	container := bs.Container()
	container.Append("Hello, World!")

	element := container.Element()
	assert.Equal(t, 1, len(element.ChildNodes()), "Should have one child node")
	assert.Equal(t, "Hello, World!", element.TextContent(), "Should have correct text content")
}

func TestContainer_AppendMultipleTextNodes(t *testing.T) {
	container := bs.Container()
	container.Append("First", " Second", " Third")

	element := container.Element()
	assert.Equal(t, 3, len(element.ChildNodes()), "Should have three child nodes")
	assert.Equal(t, "First Second Third", element.TextContent(), "Should concatenate all text")
}

func TestContainer_AppendElements(t *testing.T) {
	container := bs.Container()
	doc := domPkg.GetWindow().Document()

	// Create some child elements
	div := doc.CreateElement("DIV")
	div.SetAttribute("class", "child")
	div.AppendChild(doc.CreateTextNode("Child div"))

	span := doc.CreateElement("SPAN")
	span.AppendChild(doc.CreateTextNode("Child span"))

	container.Append(div, span)

	element := container.Element()
	assert.Equal(t, 2, len(element.ChildNodes()), "Should have two child elements")
	assert.Equal(t, "Child divChild span", element.TextContent(), "Should have correct text content")

	// Verify the children
	firstChild := element.ChildNodes()[0].(dom.Element)
	assert.Equal(t, "DIV", firstChild.TagName(), "First child should be a div")
	assert.Equal(t, "child", firstChild.GetAttribute("class"), "First child should have class")

	secondChild := element.ChildNodes()[1].(dom.Element)
	assert.Equal(t, "SPAN", secondChild.TagName(), "Second child should be a span")
}

func TestContainer_AppendComponents(t *testing.T) {
	container := bs.Container()

	// Create a heading component
	heading := bs.Heading(2, bs.WithClass("inner-heading"))
	heading.Append("Nested heading")

	// Append the heading component to the container
	container.Append(heading)

	element := container.Element()
	assert.Equal(t, 1, len(element.ChildNodes()), "Should have one child component")

	// Verify the heading was appended correctly
	childElement := element.ChildNodes()[0].(dom.Element)
	assert.Equal(t, "H2", childElement.TagName(), "Child should be an H2 element")
	assert.True(t, childElement.ClassList().Contains("inner-heading"), "Child should have inner-heading class")
	assert.Equal(t, "Nested heading", childElement.TextContent(), "Child should have correct text")
}

func TestContainer_AppendMixed(t *testing.T) {
	container := bs.Container()
	doc := domPkg.GetWindow().Document()

	// Create a heading component
	heading := bs.Heading(3)
	heading.Append("Section Title")

	// Create a plain element
	p := doc.CreateElement("P")
	p.AppendChild(doc.CreateTextNode("Paragraph text"))

	// Append mixed types: component, element, and text
	container.Append(heading, p, "Plain text")

	element := container.Element()
	assert.Equal(t, 3, len(element.ChildNodes()), "Should have three children")
	assert.Equal(t, "Section TitleParagraph textPlain text", element.TextContent(), "Should have all text content")
}

func TestContainer_ChainedAppends(t *testing.T) {
	container := bs.Container()
	result := container.Append("First").Append("Second").Append("Third")

	assert.Equal(t, container, result, "Append should return container for chaining")
	assert.Equal(t, "FirstSecondThird", container.Element().TextContent(),
		"All appended text should be present")
}

func TestContainer_ElementProperties(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	t.Run("basic element properties", func(t *testing.T) {
		assert.Equal(t, "DIV", element.TagName(), "Element should be a div")
		assert.Equal(t, dom.ELEMENT_NODE, element.NodeType(), "Should be element node type")
		assert.Equal(t, "", element.InnerHTML(), "Should have empty inner HTML initially")
	})

	t.Run("attributes", func(t *testing.T) {
		assert.True(t, element.HasAttributes(), "Should have attributes")
		assert.True(t, element.HasAttribute("class"), "Should have class attribute")
		assert.False(t, element.HasAttribute("id"), "Should not have id attribute initially")
	})

	t.Run("child nodes", func(t *testing.T) {
		assert.False(t, element.HasChildNodes(), "Should have no child nodes initially")
		assert.Nil(t, element.FirstChild(), "First child should be nil")
		assert.Nil(t, element.LastChild(), "Last child should be nil")
	})
}

func TestContainer_ContentManipulation(t *testing.T) {
	container := bs.Container()
	element := container.Element()

	// We need to get a document to create child elements
	doc := element.OwnerDocument()
	if doc == nil {
		// Fallback if OwnerDocument is not implemented
		doc = domPkg.GetWindow().Document()
	}

	t.Run("add child elements", func(t *testing.T) {
		// Create and add child elements
		child1 := doc.CreateElement("p")
		child1.AppendChild(doc.CreateTextNode("Paragraph 1"))

		child2 := doc.CreateElement("span")
		child2.AppendChild(doc.CreateTextNode("Span content"))

		element.AppendChild(child1)
		element.AppendChild(child2)

		assert.True(t, element.HasChildNodes(), "Should have child nodes after adding")
		assert.NotNil(t, element.FirstChild(), "First child should not be nil")
		assert.NotNil(t, element.LastChild(), "Last child should not be nil")
		assert.NotEqual(t, "", element.InnerHTML(), "Inner HTML should not be empty after adding children")
	})

	t.Run("set attributes", func(t *testing.T) {
		element.SetAttribute("id", "my-container")
		element.SetAttribute("data-test", "container-test")

		assert.True(t, element.HasAttribute("id"), "Should have id attribute")
		assert.True(t, element.HasAttribute("data-test"), "Should have data-test attribute")
		assert.Equal(t, "my-container", element.GetAttribute("id"), "ID should match")
		assert.Equal(t, "container-test", element.GetAttribute("data-test"), "Data attribute should match")
	})
}

func TestContainer_ClassListOperations(t *testing.T) {
	container := bs.Container()
	classList := container.Element().ClassList()

	t.Run("initial state", func(t *testing.T) {
		assert.Equal(t, 1, classList.Length(), "Should start with one class")
		assert.True(t, classList.Contains("container"), "Should contain container class")
	})

	t.Run("add classes", func(t *testing.T) {
		classList.Add("custom-class", "another-class")
		assert.Equal(t, 3, classList.Length(), "Should have 3 classes")
		assert.True(t, classList.Contains("custom-class"), "Should contain custom-class")
		assert.True(t, classList.Contains("another-class"), "Should contain another-class")
		assert.True(t, classList.Contains("container"), "Should still contain container class")
	})

	t.Run("remove classes", func(t *testing.T) {
		classList.Remove("custom-class")
		assert.Equal(t, 2, classList.Length(), "Should have 2 classes")
		assert.False(t, classList.Contains("custom-class"), "Should not contain custom-class")
		assert.True(t, classList.Contains("another-class"), "Should still contain another-class")
		assert.True(t, classList.Contains("container"), "Should still contain container class")
	})

	t.Run("toggle classes", func(t *testing.T) {
		result := classList.Toggle("toggle-class")
		assert.True(t, result, "Toggle should return true for added class")
		assert.True(t, classList.Contains("toggle-class"), "Should contain toggle-class")

		result = classList.Toggle("toggle-class")
		assert.False(t, result, "Toggle should return false for removed class")
		assert.False(t, classList.Contains("toggle-class"), "Should not contain toggle-class")
	})
}

func TestContainer_MultipleContainers(t *testing.T) {
	// Test creating multiple containers
	container1 := bs.Container()
	container2 := bs.Container(bs.WithBreakpoint(bs.BreakpointFluid))
	container3 := bs.Container(bs.WithClass("custom"))

	assert.NotEqual(t, container1.Element(), container2.Element(), "Different containers should have different elements")
	assert.NotEqual(t, container1.Element(), container3.Element(), "Different containers should have different elements")
	assert.NotEqual(t, container2.Element(), container3.Element(), "Different containers should have different elements")

	// Verify each has correct classes
	assert.True(t, container1.Element().ClassList().Contains("container"), "Container1 should have container class")
	assert.True(t, container2.Element().ClassList().Contains("container-fluid"), "Container2 should have container-fluid class")
	assert.True(t, container3.Element().ClassList().Contains("container"), "Container3 should have container class")
	assert.True(t, container3.Element().ClassList().Contains("custom"), "Container3 should have custom class")
}

func TestContainer_OptionsOrder(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
		description     string
	}{
		{
			name:            "fluid then custom class",
			options:         []bs.Opt{bs.WithBreakpoint(bs.BreakpointFluid), bs.WithClass("custom")},
			expectedClasses: []string{"container-fluid", "custom"},
			description:     "WithBreakpoint(BreakpointFluid) should replace container with container-fluid, then add custom",
		},
		{
			name:            "custom class then fluid",
			options:         []bs.Opt{bs.WithClass("custom"), bs.WithBreakpoint(bs.BreakpointFluid)},
			expectedClasses: []string{"container-fluid", "custom"},
			description:     "Order should not matter for final result",
		},
		{
			name:            "multiple custom classes with fluid",
			options:         []bs.Opt{bs.WithClass("class1"), bs.WithBreakpoint(bs.BreakpointFluid), bs.WithClass("class2")},
			expectedClasses: []string{"container-fluid", "class1", "class2"},
			description:     "All custom classes should be preserved with fluid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(tt.options...)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes: %s", tt.description)

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class %s: %s", expectedClass, tt.description)
			}
		})
	}
}

func TestContainer_EdgeCases(t *testing.T) {
	t.Run("empty options", func(t *testing.T) {
		container := bs.Container()
		assert.NotNil(t, container, "Container with no options should not be nil")
		assert.Equal(t, "container", container.Element().GetAttribute("class"),
			"Default container should have container class")
	})

	t.Run("nil options handling", func(t *testing.T) {
		// This should not panic
		assert.NotPanics(t, func() {
			container := bs.Container(nil)
			assert.NotNil(t, container, "Container should handle nil option gracefully")
		}, "Container creation should handle nil options")
	})

	t.Run("duplicate classes", func(t *testing.T) {
		container := bs.Container(bs.WithClass("container", "duplicate", "duplicate"))
		classList := container.Element().ClassList()

		// Count how many times "duplicate" appears
		duplicateCount := 0
		for _, class := range classList.Values() {
			if class == "duplicate" {
				duplicateCount++
			}
		}
		assert.Equal(t, 1, duplicateCount, "Duplicate classes should be handled properly")
	})

	t.Run("empty class names", func(t *testing.T) {
		// This tests edge case handling for empty strings
		assert.NotPanics(t, func() {
			container := bs.Container(bs.WithClass("", "valid-class", ""))
			classList := container.Element().ClassList()
			assert.True(t, classList.Contains("container"), "Should still have container class")
			assert.True(t, classList.Contains("valid-class"), "Should have valid class")
		}, "Should handle empty class names gracefully")
	})
}

// Legacy test for backward compatibility
func Test_Container_001(t *testing.T) {
	assert := assert.New(t)

	// Create a container
	container := bs.Container()
	assert.Equal(`<div class="container" data-component="container"></div>`, strings.ToLower(container.Element().OuterHTML()))

	// Create a fluid container
	fluid := bs.Container(bs.WithBreakpoint(bs.BreakpointFluid))
	assert.Equal(`<div class="container-fluid" data-component="container"></div>`, strings.ToLower(fluid.Element().OuterHTML()))
}

// TestContainer_ResponsiveBreakpoints tests all responsive breakpoint container options
func TestContainer_ResponsiveBreakpoints(t *testing.T) {
	tests := []struct {
		name       string
		breakpoint bs.Breakpoint
		expected   string
	}{
		{
			name:       "default container",
			breakpoint: bs.BreakpointDefault,
			expected:   "container",
		},
		{
			name:       "small breakpoint",
			breakpoint: bs.BreakpointSmall,
			expected:   "container-sm",
		},
		{
			name:       "medium breakpoint",
			breakpoint: bs.BreakpointMedium,
			expected:   "container-md",
		},
		{
			name:       "large breakpoint",
			breakpoint: bs.BreakpointLarge,
			expected:   "container-lg",
		},
		{
			name:       "extra large breakpoint",
			breakpoint: bs.BreakpointXLarge,
			expected:   "container-xl",
		},
		{
			name:       "extra extra large breakpoint",
			breakpoint: bs.BreakpointXXLarge,
			expected:   "container-xxl",
		},
		{
			name:       "fluid container",
			breakpoint: bs.BreakpointFluid,
			expected:   "container-fluid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(bs.WithBreakpoint(tt.breakpoint))
			element := container.Element()

			// Check class attribute
			assert.Equal(t, tt.expected, element.GetAttribute("class"),
				"Container should have '%s' class", tt.expected)

			// Check classList
			classList := element.ClassList()
			assert.Equal(t, 1, classList.Length(), "Container should have exactly one class")
			assert.True(t, classList.Contains(tt.expected),
				"Container should contain '%s' class", tt.expected)

			// Check OuterHTML
			expectedHTML := `<div class="` + tt.expected + `" data-component="container"></div>`
			assert.Equal(t, expectedHTML, strings.ToLower(element.OuterHTML()),
				"OuterHTML should match expected format")
		})
	}
}

// TestContainer_BreakpointWithAdditionalClasses tests responsive containers with additional classes
func TestContainer_BreakpointWithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name           string
		breakpoint     bs.Breakpoint
		additionalOpts []bs.Opt
		expectedClass  string
	}{
		{
			name:           "small with margin",
			breakpoint:     bs.BreakpointSmall,
			additionalOpts: []bs.Opt{bs.WithClass("mt-4")},
			expectedClass:  "container-sm mt-4",
		},
		{
			name:           "medium with padding",
			breakpoint:     bs.BreakpointMedium,
			additionalOpts: []bs.Opt{bs.WithClass("p-3")},
			expectedClass:  "container-md p-3",
		},
		{
			name:           "large with multiple utilities",
			breakpoint:     bs.BreakpointLarge,
			additionalOpts: []bs.Opt{bs.WithClass("mt-5", "mb-3")},
			expectedClass:  "container-lg mt-5 mb-3",
		},
		{
			name:           "extra large with custom class",
			breakpoint:     bs.BreakpointXLarge,
			additionalOpts: []bs.Opt{bs.WithClass("custom-container")},
			expectedClass:  "container-xl custom-container",
		},
		{
			name:           "extra extra large with text-center",
			breakpoint:     bs.BreakpointXXLarge,
			additionalOpts: []bs.Opt{bs.WithClass("text-center")},
			expectedClass:  "container-xxl text-center",
		},
		{
			name:           "fluid with margin",
			breakpoint:     bs.BreakpointFluid,
			additionalOpts: []bs.Opt{bs.WithClass("mt-4")},
			expectedClass:  "container-fluid mt-4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := append([]bs.Opt{bs.WithBreakpoint(tt.breakpoint)}, tt.additionalOpts...)
			container := bs.Container(opts...)
			element := container.Element()

			classAttr := element.GetAttribute("class")

			// Check that all expected classes are present
			expectedClasses := strings.Split(tt.expectedClass, " ")
			for _, expectedClass := range expectedClasses {
				assert.True(t, element.ClassList().Contains(expectedClass),
					"Container should contain '%s' class", expectedClass)
			}

			// Verify the full class attribute matches
			assert.Equal(t, tt.expectedClass, classAttr,
				"Container class attribute should match expected")
		})
	}
}

// TestContainer_BreakpointSwitching tests switching between different container types
func TestContainer_BreakpointSwitching(t *testing.T) {
	tests := []struct {
		name     string
		options  []bs.Opt
		expected string
	}{
		{
			name:     "fluid overrides small",
			options:  []bs.Opt{bs.WithBreakpoint(bs.BreakpointSmall), bs.WithBreakpoint(bs.BreakpointFluid)},
			expected: "container-fluid",
		},
		{
			name:     "large overrides medium",
			options:  []bs.Opt{bs.WithBreakpoint(bs.BreakpointMedium), bs.WithBreakpoint(bs.BreakpointLarge)},
			expected: "container-lg",
		},
		{
			name:     "xxl overrides all",
			options:  []bs.Opt{bs.WithBreakpoint(bs.BreakpointSmall), bs.WithBreakpoint(bs.BreakpointMedium), bs.WithBreakpoint(bs.BreakpointXXLarge)},
			expected: "container-xxl",
		},
		{
			name:     "default overrides fluid",
			options:  []bs.Opt{bs.WithBreakpoint(bs.BreakpointFluid), bs.WithBreakpoint(bs.BreakpointDefault)},
			expected: "container",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(tt.options...)
			element := container.Element()

			classAttr := element.GetAttribute("class")
			assert.Equal(t, tt.expected, classAttr,
				"Last option should win when multiple container types are specified")

			// Verify only one container class exists
			classList := element.ClassList()
			assert.Equal(t, 1, classList.Length(),
				"Should have exactly one class when switching between container types")
		})
	}
}

// TestContainer_WithBorder tests the WithBorder option
func TestContainer_WithBorder(t *testing.T) {
	tests := []struct {
		name            string
		position        bs.Position
		color           []bs.Color
		expectedClasses []string
	}{
		{
			name:            "border on all sides",
			position:        bs.BorderAll,
			color:           nil,
			expectedClasses: []string{"container", "border"},
		},
		{
			name:            "border top only",
			position:        bs.TOP,
			color:           nil,
			expectedClasses: []string{"container", "border-top"},
		},
		{
			name:            "border bottom only",
			position:        bs.BOTTOM,
			color:           nil,
			expectedClasses: []string{"container", "border-bottom"},
		},
		{
			name:            "border start only",
			position:        bs.START,
			color:           nil,
			expectedClasses: []string{"container", "border-start"},
		},
		{
			name:            "border end only",
			position:        bs.END,
			color:           nil,
			expectedClasses: []string{"container", "border-end"},
		},
		{
			name:            "border top and bottom",
			position:        bs.TOP | bs.BOTTOM,
			color:           nil,
			expectedClasses: []string{"container", "border-top", "border-bottom"},
		},
		{
			name:            "border start and end",
			position:        bs.START | bs.END,
			color:           nil,
			expectedClasses: []string{"container", "border-start", "border-end"},
		},
		{
			name:            "border all with primary color",
			position:        bs.BorderAll,
			color:           []bs.Color{bs.PRIMARY},
			expectedClasses: []string{"container", "border", "border-primary"},
		},
		{
			name:            "border top with danger color",
			position:        bs.TOP,
			color:           []bs.Color{bs.DANGER},
			expectedClasses: []string{"container", "border-top", "border-danger"},
		},
		{
			name:            "border all with success color",
			position:        bs.BorderAll,
			color:           []bs.Color{bs.SUCCESS},
			expectedClasses: []string{"container", "border", "border-success"},
		},
		{
			name:            "border bottom with warning color",
			position:        bs.BOTTOM,
			color:           []bs.Color{bs.WARNING},
			expectedClasses: []string{"container", "border-bottom", "border-warning"},
		},
		{
			name:            "border with subtle color",
			position:        bs.BorderAll,
			color:           []bs.Color{bs.PRIMARY_SUBTLE},
			expectedClasses: []string{"container", "border", "border-primary-subtle"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var container dom.Component
			if len(tt.color) > 0 {
				container = bs.Container(bs.WithBorder(tt.position, tt.color[0]))
			} else {
				container = bs.Container(bs.WithBorder(tt.position))
			}

			classList := container.Element().ClassList()
			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

// TestContainer_WithBorderReplacement tests that WithBorder replaces existing borders
func TestContainer_WithBorderReplacement(t *testing.T) {
	// First apply border-top
	container := bs.Container(bs.WithBorder(bs.TOP))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("border-top"), "Should have border-top")
	assert.Equal(t, 2, classList.Length(), "Should have container and border-top")

	// Now apply border-bottom - this should be done by recreating the container
	// since we can't directly modify the opts after creation
	t.Run("border replacement by recreating", func(t *testing.T) {
		newContainer := bs.Container(bs.WithBorder(bs.BOTTOM))
		newClassList := newContainer.Element().ClassList()

		assert.True(t, newClassList.Contains("border-bottom"), "Should have border-bottom")
		assert.False(t, newClassList.Contains("border-top"), "Should not have border-top")
		assert.Equal(t, 2, newClassList.Length(), "Should have container and border-bottom")
	})
}

// TestContainer_WithBackground tests the WithBackground option
func TestContainer_WithBackground(t *testing.T) {
	tests := []struct {
		name            string
		color           bs.Color
		expectedClasses []string
	}{
		{
			name:            "primary background",
			color:           bs.PRIMARY,
			expectedClasses: []string{"container", "bg-primary"},
		},
		{
			name:            "secondary background",
			color:           bs.SECONDARY,
			expectedClasses: []string{"container", "bg-secondary"},
		},
		{
			name:            "success background",
			color:           bs.SUCCESS,
			expectedClasses: []string{"container", "bg-success"},
		},
		{
			name:            "danger background",
			color:           bs.DANGER,
			expectedClasses: []string{"container", "bg-danger"},
		},
		{
			name:            "warning background",
			color:           bs.WARNING,
			expectedClasses: []string{"container", "bg-warning"},
		},
		{
			name:            "info background",
			color:           bs.INFO,
			expectedClasses: []string{"container", "bg-info"},
		},
		{
			name:            "light background",
			color:           bs.LIGHT,
			expectedClasses: []string{"container", "bg-light"},
		},
		{
			name:            "dark background",
			color:           bs.DARK,
			expectedClasses: []string{"container", "bg-dark"},
		},
		{
			name:            "white background",
			color:           bs.WHITE,
			expectedClasses: []string{"container", "bg-white"},
		},
		{
			name:            "black background",
			color:           bs.BLACK,
			expectedClasses: []string{"container", "bg-black"},
		},
		{
			name:            "primary subtle background",
			color:           bs.PRIMARY_SUBTLE,
			expectedClasses: []string{"container", "bg-primary-subtle"},
		},
		{
			name:            "danger subtle background",
			color:           bs.DANGER_SUBTLE,
			expectedClasses: []string{"container", "bg-danger-subtle"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(bs.WithBackground(tt.color))
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

// TestContainer_WithBorderAndBackground tests combining border and background
func TestContainer_WithBorderAndBackground(t *testing.T) {
	tests := []struct {
		name            string
		position        bs.Position
		borderColor     bs.Color
		bgColor         bs.Color
		expectedClasses []string
	}{
		{
			name:            "border with background",
			position:        bs.BorderAll,
			borderColor:     bs.PRIMARY,
			bgColor:         bs.LIGHT,
			expectedClasses: []string{"container", "border", "border-primary", "bg-light"},
		},
		{
			name:            "top border with danger background",
			position:        bs.TOP,
			borderColor:     bs.DANGER,
			bgColor:         bs.DANGER_SUBTLE,
			expectedClasses: []string{"container", "border-top", "border-danger", "bg-danger-subtle"},
		},
		{
			name:            "all borders with success colors",
			position:        bs.BorderAll,
			borderColor:     bs.SUCCESS,
			bgColor:         bs.SUCCESS_SUBTLE,
			expectedClasses: []string{"container", "border", "border-success", "bg-success-subtle"},
		},
		{
			name:            "multiple borders with background",
			position:        bs.TOP | bs.BOTTOM,
			borderColor:     bs.INFO,
			bgColor:         bs.INFO_SUBTLE,
			expectedClasses: []string{"container", "border-top", "border-bottom", "border-info", "bg-info-subtle"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(
				bs.WithBorder(tt.position, tt.borderColor),
				bs.WithBackground(tt.bgColor),
			)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

// TestContainer_WithBorderAndBackgroundWithBreakpoint tests all options together
func TestContainer_WithBorderAndBackgroundWithBreakpoint(t *testing.T) {
	tests := []struct {
		name            string
		breakpoint      bs.Breakpoint
		position        bs.Position
		borderColor     bs.Color
		bgColor         bs.Color
		expectedClasses []string
	}{
		{
			name:            "fluid container with border and background",
			breakpoint:      bs.BreakpointFluid,
			position:        bs.BorderAll,
			borderColor:     bs.PRIMARY,
			bgColor:         bs.LIGHT,
			expectedClasses: []string{"container-fluid", "border", "border-primary", "bg-light"},
		},
		{
			name:            "small container with top border and background",
			breakpoint:      bs.BreakpointSmall,
			position:        bs.TOP,
			borderColor:     bs.DANGER,
			bgColor:         bs.WARNING_SUBTLE,
			expectedClasses: []string{"container-sm", "border-top", "border-danger", "bg-warning-subtle"},
		},
		{
			name:            "large container with all sides border and dark background",
			breakpoint:      bs.BreakpointLarge,
			position:        bs.BorderAll,
			borderColor:     bs.DARK,
			bgColor:         bs.DARK_SUBTLE,
			expectedClasses: []string{"container-lg", "border", "border-dark", "bg-dark-subtle"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := bs.Container(
				bs.WithBreakpoint(tt.breakpoint),
				bs.WithBorder(tt.position, tt.borderColor),
				bs.WithBackground(tt.bgColor),
			)
			classList := container.Element().ClassList()

			assert.Equal(t, len(tt.expectedClasses), classList.Length(),
				"Should have expected number of classes")

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Should contain class: %s", expectedClass)
			}
		})
	}
}

// TestContainer_WithBorderNone tests that NONE position removes borders
func TestContainer_WithBorderNone(t *testing.T) {
	container := bs.Container(bs.WithBorder(bs.NONE))
	classList := container.Element().ClassList()

	assert.Equal(t, 1, classList.Length(), "Should only have container class")
	assert.True(t, classList.Contains("container"), "Should have container class")
	assert.False(t, classList.Contains("border"), "Should not have any border class")
}

// TestContainer_ComplexCombination tests complex combinations of all options
func TestContainer_ComplexCombination(t *testing.T) {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointMedium),
		bs.WithClass("mt-4", "p-3"),
		bs.WithBorder(bs.TOP|bs.BOTTOM, bs.PRIMARY),
		bs.WithBackground(bs.LIGHT),
	)

	expectedClasses := []string{
		"container-md",
		"mt-4",
		"p-3",
		"border-top",
		"border-bottom",
		"border-primary",
		"bg-light",
	}

	classList := container.Element().ClassList()
	assert.Equal(t, len(expectedClasses), classList.Length(),
		"Should have all expected classes")

	for _, expectedClass := range expectedClasses {
		assert.True(t, classList.Contains(expectedClass),
			"Should contain class: %s", expectedClass)
	}
}

// TestContainer_WithFlex_Center tests center alignment with WithFlex
func TestContainer_WithFlex_Center(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.CENTER))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-center"))
}

// TestContainer_WithFlex_Middle tests middle alignment with WithFlex
func TestContainer_WithFlex_Middle(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.MIDDLE))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-center"))
}

// TestContainer_WithFlex_Start tests start alignment with WithFlex
func TestContainer_WithFlex_Start(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.START))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-start"))
}

// TestContainer_WithFlex_Top tests top alignment with WithFlex
func TestContainer_WithFlex_Top(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.TOP))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-start"))
}

// TestContainer_WithFlex_End tests end alignment with WithFlex
func TestContainer_WithFlex_End(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.END))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-end"))
}

// TestContainer_WithFlex_Bottom tests bottom alignment with WithFlex
func TestContainer_WithFlex_Bottom(t *testing.T) {
	container := bs.Container(bs.WithFlex(bs.BOTTOM))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-end"))
}

// TestContainer_WithFlex_HorizontalFlow tests horizontal flow with WithFlex
func TestContainer_WithFlex_HorizontalFlow(t *testing.T) {
	// START | END means horizontal flow
	container := bs.Container(bs.WithFlex(bs.START | bs.END))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("flex-row"))
}

// TestContainer_WithFlex_VerticalFlow tests vertical flow with WithFlex
func TestContainer_WithFlex_VerticalFlow(t *testing.T) {
	// TOP | BOTTOM means vertical flow
	container := bs.Container(bs.WithFlex(bs.TOP | bs.BOTTOM))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("flex-column"))
}

// TestContainer_WithFlex_CenterWithHorizontalFlow tests center alignment with horizontal flow
func TestContainer_WithFlex_CenterWithHorizontalFlow(t *testing.T) {
	// CENTER with horizontal flow
	container := bs.Container(bs.WithFlex(bs.CENTER | bs.START | bs.END))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-center"))
	assert.True(t, classList.Contains("flex-row"))
}

// TestContainer_WithFlex_CenterWithVerticalFlow tests center alignment with vertical flow
func TestContainer_WithFlex_CenterWithVerticalFlow(t *testing.T) {
	// CENTER with vertical flow
	container := bs.Container(bs.WithFlex(bs.CENTER | bs.TOP | bs.BOTTOM))
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-center"))
	assert.True(t, classList.Contains("flex-column"))
}

// TestContainer_WithFlex_WithOtherOptions tests WithFlex combined with other options
func TestContainer_WithFlex_WithOtherOptions(t *testing.T) {
	// Test that WithFlex works with other options
	container := bs.Container(
		bs.WithFlex(bs.CENTER),
		bs.WithClass("gap-3"),
		bs.WithMargin(bs.TOP, 4),
	)
	classList := container.Element().ClassList()
	assert.True(t, classList.Contains("d-flex"))
	assert.True(t, classList.Contains("align-items-center"))
	assert.True(t, classList.Contains("gap-3"))
	assert.True(t, classList.Contains("mt-4"))
}
