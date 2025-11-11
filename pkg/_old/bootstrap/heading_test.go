package bootstrap_test

import (
	"fmt"
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestHeading_Basic(t *testing.T) {
	tests := []struct {
		level       int
		expectedTag string
	}{
		{level: 1, expectedTag: "H1"},
		{level: 2, expectedTag: "H2"},
		{level: 3, expectedTag: "H3"},
		{level: 4, expectedTag: "H4"},
		{level: 5, expectedTag: "H5"},
		{level: 6, expectedTag: "H6"},
	}

	for _, tt := range tests {
		t.Run(tt.expectedTag, func(t *testing.T) {
			heading := bs.Heading(tt.level)

			assert.NotNil(t, heading, "Heading should not be nil")
			assert.NotNil(t, heading.Element(), "Heading element should not be nil")
			assert.Equal(t, dom.ELEMENT_NODE, heading.Element().NodeType(), "Heading should be an element node")
			assert.Equal(t, tt.expectedTag, heading.Element().TagName(), "Heading should have correct tag")
			assert.Equal(t, tt.level, heading.Level(), "Heading should have correct level")
			assert.Equal(t, "heading", heading.Name(), "Heading should have correct component name")
		})
	}
}

func TestHeading_InvalidLevel(t *testing.T) {
	tests := []struct {
		name  string
		level int
	}{
		{name: "level 0", level: 0},
		{name: "level -1", level: -1},
		{name: "level 7", level: 7},
		{name: "level 10", level: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() {
				bs.Heading(tt.level)
			}, "Should panic for invalid level %d", tt.level)
		})
	}
}

func TestHeading_WithClasses(t *testing.T) {
	heading := bs.Heading(1, bs.WithClass("display-1", "text-center"))
	element := heading.Element()

	classList := element.ClassList()
	assert.Equal(t, 2, classList.Length(), "Heading should have two classes")
	assert.True(t, classList.Contains("display-1"), "Heading should have display-1 class")
	assert.True(t, classList.Contains("text-center"), "Heading should have text-center class")
}

func TestHeading_OuterHTML(t *testing.T) {
	tests := []struct {
		name         string
		level        int
		options      []bs.Opt
		expectedHTML string
	}{
		{
			name:         "h1 without classes",
			level:        1,
			options:      nil,
			expectedHTML: "<h1 data-component=\"heading\"></h1>",
		},
		{
			name:         "h2 with single class",
			level:        2,
			options:      []bs.Opt{bs.WithClass("text-primary")},
			expectedHTML: `<h2 class="text-primary" data-component="heading"></h2>`,
		},
		{
			name:         "h3 with multiple classes",
			level:        3,
			options:      []bs.Opt{bs.WithClass("display-3", "text-center")},
			expectedHTML: `<h3 class="display-3 text-center" data-component="heading"></h3>`,
		},
		{
			name:         "h6 with bootstrap utility classes",
			level:        6,
			options:      []bs.Opt{bs.WithClass("fw-bold", "text-muted")},
			expectedHTML: `<h6 class="fw-bold text-muted" data-component="heading"></h6>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heading := bs.Heading(tt.level, tt.options...)
			assert.Equal(t, tt.expectedHTML, strings.ToLower(heading.Element().OuterHTML()),
				"OuterHTML should match expected format")
		})
	}
}

func TestHeading_ComponentInterface(t *testing.T) {
	heading := bs.Heading(1)

	// Verify it implements the Component interface methods
	var _ dom.Component = heading

	assert.NotNil(t, heading.Element(), "Element() should return non-nil")
	assert.Equal(t, "heading", heading.Name(), "Name() should return 'heading'")
}

func TestHeading_AppendText(t *testing.T) {
	heading := bs.Heading(2)
	heading.Append("Hello, World!")

	element := heading.Element()
	assert.Equal(t, 1, len(element.ChildNodes()), "Heading should have one child")
	assert.Equal(t, "Hello, World!", element.TextContent(), "Heading text should match")

	expectedHTML := "<h2 data-component=\"heading\">hello, world!</h2>"
	assert.Equal(t, expectedHTML, strings.ToLower(element.OuterHTML()),
		"OuterHTML should include text content")
}

func TestHeading_AppendMultipleTextNodes(t *testing.T) {
	heading := bs.Heading(1)
	heading.Append("Part 1", " ", "Part 2")

	element := heading.Element()
	assert.Equal(t, 3, len(element.ChildNodes()), "Heading should have three children")
	assert.Equal(t, "Part 1 Part 2", element.TextContent(), "Combined text should match")
}

func TestHeading_AppendElements(t *testing.T) {
	heading := bs.Heading(3)

	// Create a span element
	doc := domPkg.GetWindow().Document()
	span := doc.CreateElement("SPAN")
	span.SetAttribute("class", "badge")
	span.AppendChild(doc.CreateTextNode("New"))

	heading.Append("Title ", span)

	element := heading.Element()
	assert.Equal(t, 2, len(element.ChildNodes()), "Heading should have two children")

	expectedHTML := `<h3 data-component="heading">title <span class="badge">new</span></h3>`
	assert.Equal(t, expectedHTML, strings.ToLower(element.OuterHTML()),
		"OuterHTML should include nested element")
}

func TestHeading_AppendComponents(t *testing.T) {
	heading := bs.Heading(1)
	container := bs.Container(bs.WithClass("d-inline"))
	container.Append("Nested")

	heading.Append("Main ", container)

	element := heading.Element()
	assert.Equal(t, 2, len(element.ChildNodes()), "Heading should have two children")

	// Check that container was appended
	expectedHTML := `<h1 data-component="heading">main <div class="container d-inline" data-component="container">nested</div></h1>`
	assert.Equal(t, expectedHTML, strings.ToLower(element.OuterHTML()),
		"OuterHTML should include nested component")
}

func TestHeading_ChainedAppends(t *testing.T) {
	heading := bs.Heading(2)
	result := heading.Append("First").Append(" Second").Append(" Third")

	// The Append method returns Component interface, but it should be the same underlying object
	assert.Equal(t, heading.Element(), result.Element(), "Append should return same heading for chaining")
	assert.Equal(t, "First Second Third", heading.Element().TextContent(),
		"All appended text should be present")
}

func TestHeading_ElementProperties(t *testing.T) {
	t.Run("tag_names", func(t *testing.T) {
		for level := 1; level <= 6; level++ {
			heading := bs.Heading(level)
			expectedTag := strings.ToUpper(fmt.Sprintf("H%d", level))
			assert.Equal(t, expectedTag, heading.Element().TagName(),
				"Heading level %d should have tag %s", level, expectedTag)
		}
	})

	t.Run("node_type", func(t *testing.T) {
		heading := bs.Heading(1)
		assert.Equal(t, dom.ELEMENT_NODE, heading.Element().NodeType(),
			"Heading should be an element node")
	})

	t.Run("node_name", func(t *testing.T) {
		heading := bs.Heading(3)
		assert.Equal(t, "H3", strings.ToUpper(heading.Element().NodeName()),
			"Node name should match tag name")
	})
}

func TestHeading_Attributes(t *testing.T) {
	heading := bs.Heading(1)
	element := heading.Element()

	// Set custom attributes
	element.SetAttribute("id", "main-heading")
	element.SetAttribute("data-level", "1")

	assert.True(t, element.HasAttribute("id"), "Should have id attribute")
	assert.Equal(t, "main-heading", element.GetAttribute("id"), "ID should match")
	assert.Equal(t, "1", element.GetAttribute("data-level"), "Data attribute should match")
}

func TestHeading_ClassListOperations(t *testing.T) {
	heading := bs.Heading(2, bs.WithClass("initial"))
	element := heading.Element()
	classList := element.ClassList()

	t.Run("initial_class", func(t *testing.T) {
		assert.True(t, classList.Contains("initial"), "Should have initial class")
	})

	// Note: ClassList modifications don't sync back to the class attribute in non-WASM environment
	// This is a known limitation of the current DOM implementation
}

func TestHeading_BootstrapDisplayClasses(t *testing.T) {
	tests := []struct {
		name  string
		class string
		level int
	}{
		{name: "display-1", class: "display-1", level: 1},
		{name: "display-2", class: "display-2", level: 2},
		{name: "display-3", class: "display-3", level: 3},
		{name: "display-4", class: "display-4", level: 4},
		{name: "display-5", class: "display-5", level: 5},
		{name: "display-6", class: "display-6", level: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heading := bs.Heading(tt.level, bs.WithClass(tt.class))
			assert.True(t, heading.Element().ClassList().Contains(tt.class),
				"Heading should have %s class", tt.class)
		})
	}
}

func TestHeading_BootstrapUtilityClasses(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
	}{
		{
			name:    "text alignment",
			classes: []string{"text-start", "text-center", "text-end"},
		},
		{
			name:    "text colors",
			classes: []string{"text-primary", "text-secondary", "text-success", "text-danger"},
		},
		{
			name:    "font weight",
			classes: []string{"fw-light", "fw-normal", "fw-bold", "fw-bolder"},
		},
		{
			name:    "spacing",
			classes: []string{"mt-3", "mb-4", "ms-2", "me-1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heading := bs.Heading(1, bs.WithClass(tt.classes...))
			classList := heading.Element().ClassList()

			for _, class := range tt.classes {
				assert.True(t, classList.Contains(class),
					"Heading should have %s class", class)
			}
		})
	}
}

func TestHeading_MixedContent(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	heading := bs.Heading(1)

	// Create mixed content: text + element + text
	strong := doc.CreateElement("STRONG")
	strong.AppendChild(doc.CreateTextNode("Important"))

	heading.Append("This is ", strong, " text")

	expectedHTML := "<h1 data-component=\"heading\">this is <strong>important</strong> text</h1>"
	assert.Equal(t, expectedHTML, strings.ToLower(heading.Element().OuterHTML()),
		"Heading should support mixed content")
}

func TestHeading_EmptyHeading(t *testing.T) {
	heading := bs.Heading(1)
	element := heading.Element()

	assert.Equal(t, 0, len(element.ChildNodes()), "Empty heading should have no children")
	assert.Equal(t, "", element.TextContent(), "Empty heading should have empty text content")
	assert.Equal(t, "<h1 data-component=\"heading\"></h1>", strings.ToLower(element.OuterHTML()),
		"Empty heading OuterHTML should be correct")
}

func TestHeading_MultipleHeadings(t *testing.T) {
	headings := make([]dom.Component, 6)
	for i := 0; i < 6; i++ {
		headings[i] = bs.Heading(i + 1)
	}

	// Verify each heading is independent
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 6; j++ {
			assert.NotEqual(t, headings[i].Element(), headings[j].Element(),
				"Different headings should have different elements")
		}
	}
}

func Test_Heading_001(t *testing.T) {
	assert := assert.New(t)

	// Create headings of different levels
	h1 := bs.Heading(1)
	h1.Append("Heading 1")
	assert.Equal(`<h1 data-component="heading">heading 1</h1>`, strings.ToLower(h1.Element().OuterHTML()))

	h2 := bs.Heading(2, bs.WithClass("text-muted"))
	h2.Append("Heading 2")
	assert.Equal(`<h2 class="text-muted" data-component="heading">heading 2</h2>`, strings.ToLower(h2.Element().OuterHTML()))

	h3 := bs.Heading(3, bs.WithClass("display-3", "text-center"))
	h3.Append("Display Heading")
	assert.Equal(`<h3 class="display-3 text-center" data-component="heading">display heading</h3>`, strings.ToLower(h3.Element().OuterHTML()))
}
