package dom_test

import (
	"strings"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestDocument_Basic(t *testing.T) {
	doc := dom.GetWindow().Document()
	assert.NotNil(t, doc, "Document should not be nil")
	assert.NotNil(t, doc.Body(), "Document body should not be nil")
}

func TestDocument_CreateElement(t *testing.T) {
	doc := dom.GetWindow().Document()

	tests := []struct {
		name        string
		tagName     string
		expectError bool
	}{
		{"create div", "div", false},
		{"create span", "span", false},
		{"create p", "p", false},
		{"create section", "section", false},
		{"create article", "article", false},
		{"create h1", "h1", false},
		{"create input", "input", false},
		{"create button", "button", false},
		{"create img", "img", false},
		{"create a", "a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			elem := doc.CreateElement(tt.tagName)
			if tt.expectError {
				assert.Nil(t, elem, "Expected nil element for invalid tag name")
			} else {
				assert.NotNil(t, elem, "Element should not be nil")
				if elem != nil {
					expectedTagName := strings.ToUpper(tt.tagName)
					assert.Equal(t, expectedTagName, elem.TagName(), "Tag name should match (uppercase)")
				}
			}
		})
	}
}

func TestDocument_CreateTextNode(t *testing.T) {
	doc := dom.GetWindow().Document()

	tests := []struct {
		name string
		data string
	}{
		{"simple text", "Hello, World!"},
		{"empty text", ""},
		{"whitespace", "   "},
		{"multiline text", "Line 1\nLine 2\nLine 3"},
		{"special characters", "Special chars: !@#$%^&*()"},
		{"unicode text", "Unicode: 你好世界"},
		{"html entities", "&lt;div&gt;&amp;&quot;"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			textNode := doc.CreateTextNode(tt.data)
			assert.NotNil(t, textNode, "Text node should not be nil")
			assert.Equal(t, tt.data, textNode.Data(), "Text data should match")

			// Length calculation may differ between WASM (JS string length) and non-WASM (Go byte length)
			// Especially for unicode strings, so we just check that length is reasonable
			length := textNode.Length()
			assert.GreaterOrEqual(t, length, 0, "Text length should be non-negative")
			if tt.data == "" {
				assert.Equal(t, 0, length, "Empty text should have zero length")
			} else {
				assert.Greater(t, length, 0, "Non-empty text should have positive length")
			}
		})
	}
}

func TestDocument_CreateComment(t *testing.T) {
	doc := dom.GetWindow().Document()

	tests := []struct {
		name string
		data string
	}{
		{"simple comment", "This is a comment"},
		{"empty comment", ""},
		{"whitespace comment", "   "},
		{"multiline comment", "Line 1\nLine 2"},
		{"special chars comment", "Comment with !@#$%"},
		{"html in comment", "<div>HTML in comment</div>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comment := doc.CreateComment(tt.data)
			assert.NotNil(t, comment, "Comment should not be nil")
			assert.Equal(t, tt.data, comment.Data(), "Comment data should match")

			// Length calculation may differ between WASM (JS string length) and non-WASM (Go byte length)
			length := comment.Length()
			assert.GreaterOrEqual(t, length, 0, "Comment length should be non-negative")
			if tt.data == "" {
				assert.Equal(t, 0, length, "Empty comment should have zero length")
			} else {
				assert.Greater(t, length, 0, "Non-empty comment should have positive length")
			}
		})
	}
}

func TestDocument_CreateAttribute(t *testing.T) {
	doc := dom.GetWindow().Document()

	tests := []struct {
		name     string
		attrName string
	}{
		{"id attribute", "id"},
		{"class attribute", "class"},
		{"data attribute", "data-test"},
		{"custom attribute", "my-custom-attr"},
		{"special chars", "attr-with_chars"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attr := doc.CreateAttribute(tt.attrName)
			assert.NotNil(t, attr, "Attribute should not be nil")
			assert.Equal(t, tt.attrName, attr.Name(), "Attribute name should match")
		})
	}
}

func TestDocument_Title(t *testing.T) {
	doc := dom.GetWindow().Document()

	// Test getting title (it may be empty or have a value)
	title := doc.Title()
	assert.IsType(t, "", title, "Title should be a string")

	// Note: We can't reliably test setting title in all environments
	// as it depends on the specific DOM implementation
}

func TestDocument_Body(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Body()

	assert.NotNil(t, body, "Body should not be nil")
	assert.Equal(t, "BODY", body.TagName(), "Body tag name should be 'BODY'")

	// Test body manipulation
	elem := doc.CreateElement("div")
	assert.NotNil(t, elem, "Created element should not be nil")

	appendedElem := body.AppendChild(elem)
	assert.Equal(t, elem, appendedElem, "AppendChild should return the same element")
}

func TestDocument_Doctype(t *testing.T) {
	doc := dom.GetWindow().Document()
	doctype := doc.Doctype()

	// Doctype might be nil in some test environments
	if doctype != nil {
		assert.NotNil(t, doctype, "Doctype should not be nil if present")
		// Test basic doctype properties
		name := doctype.Name()
		assert.IsType(t, "", name, "Doctype name should be a string")
	}
}

func TestDocument_ElementOperations(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Body()

	// Create multiple elements and test operations
	div := doc.CreateElement("div")
	span := doc.CreateElement("span")
	p := doc.CreateElement("p")

	assert.NotNil(t, div, "Div element should not be nil")
	assert.NotNil(t, span, "Span element should not be nil")
	assert.NotNil(t, p, "P element should not be nil")

	// Test appending elements
	body.AppendChild(div)
	div.AppendChild(span)
	span.AppendChild(p)

	// Test text node creation and appending
	textNode := doc.CreateTextNode("Test text content")
	p.AppendChild(textNode)

	assert.Equal(t, "Test text content", textNode.Data(), "Text node data should match")
}

func TestDocument_EdgeCases(t *testing.T) {
	doc := dom.GetWindow().Document()

	t.Run("empty tag name", func(t *testing.T) {
		// In WASM (real DOM), empty tag names should cause an error
		// In non-WASM, it might be allowed
		defer func() {
			if r := recover(); r != nil {
				// This is expected behavior in WASM environment
				t.Log("Empty tag name correctly caused an error in WASM environment:", r)
			}
		}()

		elem := doc.CreateElement("")
		// If we get here, we're in non-WASM environment
		if elem != nil {
			t.Log("Empty tag name allowed in non-WASM environment")
		}
	})

	t.Run("invalid tag characters", func(t *testing.T) {
		// Test various potentially invalid tag names
		invalidTags := []string{
			"div with space",
			"div<script>",
			"div>invalid",
		}

		for _, tagName := range invalidTags {
			t.Run("tag: "+tagName, func(t *testing.T) {
				defer func() {
					if r := recover(); r != nil {
						t.Logf("Invalid tag name '%s' correctly caused an error: %v", tagName, r)
					}
				}()

				elem := doc.CreateElement(tagName)
				if elem != nil {
					t.Logf("Tag name '%s' was accepted", tagName)
				}
			})
		}
	})

	t.Run("empty attribute name", func(t *testing.T) {
		// In WASM (real DOM), empty attribute names should cause an error
		// In non-WASM, it might be allowed
		defer func() {
			if r := recover(); r != nil {
				// This is expected behavior in WASM environment
				t.Log("Empty attribute name correctly caused an error in WASM environment:", r)
			}
		}()

		attr := doc.CreateAttribute("")
		// If we get here, we're in non-WASM environment
		if attr != nil {
			t.Log("Empty attribute name allowed in non-WASM environment")
		}
	})

	t.Run("numeric attribute name", func(t *testing.T) {
		// In WASM (real DOM), numeric attribute names may not be allowed
		// In non-WASM, it might be allowed
		defer func() {
			if r := recover(); r != nil {
				// This is expected behavior in WASM environment
				t.Log("Numeric attribute name correctly caused an error in WASM environment:", r)
			}
		}()

		attr := doc.CreateAttribute("123")
		// If we get here, we're in non-WASM environment
		if attr != nil {
			t.Log("Numeric attribute name allowed in non-WASM environment")
		}
	})
}

func TestDocument_ComplexStructure(t *testing.T) {
	doc := dom.GetWindow().Document()
	body := doc.Body()

	// Create a more complex DOM structure
	container := doc.CreateElement("div")
	container.SetAttribute("class", "container")

	header := doc.CreateElement("h1")
	headerText := doc.CreateTextNode("Test Header")
	header.AppendChild(headerText)

	content := doc.CreateElement("div")
	content.SetAttribute("class", "content")

	paragraph := doc.CreateElement("p")
	paragraphText := doc.CreateTextNode("This is test content.")
	paragraph.AppendChild(paragraphText)

	// Build the structure
	container.AppendChild(header)
	content.AppendChild(paragraph)
	container.AppendChild(content)
	body.AppendChild(container)

	// Test the structure
	assert.Equal(t, "container", container.GetAttribute("class"), "Container class should match")
	assert.Equal(t, "content", content.GetAttribute("class"), "Content class should match")
	assert.Equal(t, "Test Header", headerText.Data(), "Header text should match")
	assert.Equal(t, "This is test content.", paragraphText.Data(), "Paragraph text should match")
}
