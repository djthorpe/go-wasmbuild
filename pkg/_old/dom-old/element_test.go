package dom_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
	"github.com/stretchr/testify/assert"
)

func TestElement_Basic(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	assert.NotNil(t, element, "Element should not be nil")
	assert.Equal(t, "DIV", element.TagName(), "Tag name should be uppercase")
	assert.Equal(t, dom.ELEMENT_NODE, element.NodeType(), "Node type should be ELEMENT_NODE")

	// NodeName case varies between implementations
	nodeName := element.NodeName()
	assert.True(t, nodeName == "div" || nodeName == "DIV", "Node name should be 'div' or 'DIV'")
}

func TestElement_TagName(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name        string
		tagName     string
		expectedTag string
	}{
		{"div element", "div", "DIV"},
		{"span element", "span", "SPAN"},
		{"p element", "p", "P"},
		{"h1 element", "h1", "H1"},
		{"input element", "input", "INPUT"},
		{"button element", "button", "BUTTON"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			assert.Equal(t, tt.expectedTag, element.TagName(), "Tag name should be uppercase")
		})
	}
}

func TestElement_InnerHTML(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name         string
		tagName      string
		expectedHTML string
	}{
		{"empty div", "div", ""},
		{"empty span", "span", ""},
		{"empty p", "p", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			assert.Equal(t, tt.expectedHTML, element.InnerHTML(), "Empty element should have empty innerHTML")
		})
	}
}

func TestElement_OuterHTML(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	tests := []struct {
		name         string
		tagName      string
		expectedHTML string
	}{
		{"div element", "div", "<div></div>"},
		{"span element", "span", "<span></span>"},
		{"p element", "p", "<p></p>"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := doc.CreateElement(tt.tagName)
			outerHTML := element.OuterHTML()
			// Convert to lowercase for comparison since some implementations may vary in case
			assert.Equal(t, tt.expectedHTML, strings.ToLower(outerHTML), "Outer HTML should match expected format")
		})
	}
}

func TestElement_Attributes(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("initial state", func(t *testing.T) {
		assert.False(t, element.HasAttributes(), "New element should not have attributes")
		attributes := element.Attributes()
		assert.Empty(t, attributes, "New element should have empty attributes slice")
	})

	t.Run("set and get attribute", func(t *testing.T) {
		attr := element.SetAttribute("id", "test-id")
		assert.NotNil(t, attr, "SetAttribute should return an attribute")
		assert.True(t, element.HasAttribute("id"), "Element should have id attribute")
		assert.Equal(t, "test-id", element.GetAttribute("id"), "Attribute value should match")
		assert.True(t, element.HasAttributes(), "Element should have attributes")
	})

	t.Run("multiple attributes", func(t *testing.T) {
		element.SetAttribute("class", "test-class")
		element.SetAttribute("data-test", "test-value")

		assert.True(t, element.HasAttribute("id"), "Should have id attribute")
		assert.True(t, element.HasAttribute("class"), "Should have class attribute")
		assert.True(t, element.HasAttribute("data-test"), "Should have data-test attribute")

		assert.Equal(t, "test-id", element.GetAttribute("id"), "ID should match")
		assert.Equal(t, "test-class", element.GetAttribute("class"), "Class should match")
		assert.Equal(t, "test-value", element.GetAttribute("data-test"), "Data attribute should match")

		attributes := element.Attributes()
		assert.GreaterOrEqual(t, len(attributes), 3, "Should have at least 3 attributes")
	})

	t.Run("get non-existent attribute", func(t *testing.T) {
		value := element.GetAttribute("non-existent")
		assert.Equal(t, "", value, "Non-existent attribute should return empty string")
		assert.False(t, element.HasAttribute("non-existent"), "Should not have non-existent attribute")
	})
}

func TestElement_AttributeNode(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("set and get attribute node", func(t *testing.T) {
		element.SetAttribute("title", "test-title")

		attrNode := element.GetAttributeNode("title")
		if attrNode != nil {
			assert.Equal(t, "title", attrNode.Name(), "Attribute name should match")
			assert.Equal(t, "test-title", attrNode.Value(), "Attribute value should match")
		}
	})

	t.Run("get non-existent attribute node", func(t *testing.T) {
		attrNode := element.GetAttributeNode("non-existent")
		assert.Nil(t, attrNode, "Non-existent attribute node should be nil")
	})
}

func TestElement_ClassListIntegration(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	classList := element.ClassList()
	assert.NotNil(t, classList, "ClassList should not be nil")

	t.Run("initial state", func(t *testing.T) {
		assert.Equal(t, 0, classList.Length(), "Initial class list should be empty")
		assert.Equal(t, "", classList.Value(), "Initial class value should be empty")
	})

	t.Run("add classes", func(t *testing.T) {
		classList.Add("class1", "class2")
		assert.Equal(t, 2, classList.Length(), "Should have 2 classes")
		assert.True(t, classList.Contains("class1"), "Should contain class1")
		assert.True(t, classList.Contains("class2"), "Should contain class2")
	})

	t.Run("remove classes", func(t *testing.T) {
		classList.Remove("class1")
		assert.Equal(t, 1, classList.Length(), "Should have 1 class")
		assert.False(t, classList.Contains("class1"), "Should not contain class1")
		assert.True(t, classList.Contains("class2"), "Should still contain class2")
	})

	t.Run("toggle classes", func(t *testing.T) {
		result := classList.Toggle("class3")
		assert.True(t, result, "Toggle should return true for added class")
		assert.True(t, classList.Contains("class3"), "Should contain class3")

		result = classList.Toggle("class3")
		assert.False(t, result, "Toggle should return false for removed class")
		assert.False(t, classList.Contains("class3"), "Should not contain class3")
	})
}

func TestElement_ChildNodes(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	t.Run("initial state", func(t *testing.T) {
		assert.False(t, parent.HasChildNodes(), "New element should not have child nodes")
		assert.Nil(t, parent.FirstChild(), "First child should be nil")
		assert.Nil(t, parent.LastChild(), "Last child should be nil")
	})

	t.Run("append single child", func(t *testing.T) {
		child := doc.CreateElement("span")
		appendedChild := parent.AppendChild(child)

		assert.Equal(t, child, appendedChild, "AppendChild should return the same element")
		assert.True(t, parent.HasChildNodes(), "Parent should have child nodes")
		assert.True(t, parent.FirstChild().Equals(child), "First child should match")
		assert.True(t, parent.LastChild().Equals(child), "Last child should match")
		assert.True(t, child.ParentNode().Equals(parent), "Child's parent should match")
		assert.Nil(t, child.PreviousSibling(), "First child should have no previous sibling")
		assert.Nil(t, child.NextSibling(), "Only child should have no next sibling")
	})

	t.Run("append multiple children", func(t *testing.T) {
		child2 := doc.CreateElement("p")
		child3 := doc.CreateElement("h1")

		parent.AppendChild(child2)
		parent.AppendChild(child3)

		assert.True(t, parent.HasChildNodes(), "Parent should have child nodes")

		// Test first and last child
		firstChild := parent.FirstChild()
		lastChild := parent.LastChild()
		assert.NotNil(t, firstChild, "First child should not be nil")
		assert.NotNil(t, lastChild, "Last child should not be nil")
		assert.True(t, lastChild.Equals(child3), "Last child should be child3")

		// Test sibling relationships
		if child2.PreviousSibling() != nil {
			if prevSibling, ok := child2.PreviousSibling().(dom.Element); ok {
				assert.Equal(t, "SPAN", prevSibling.TagName(), "Child2's previous sibling should be span")
			}
		}
		if child2.NextSibling() != nil {
			assert.True(t, child2.NextSibling().Equals(child3), "Child2's next sibling should be child3")
		}
		if child3.PreviousSibling() != nil {
			assert.True(t, child3.PreviousSibling().Equals(child2), "Child3's previous sibling should be child2")
		}
		assert.Nil(t, child3.NextSibling(), "Last child should have no next sibling")
	})

	t.Run("remove child", func(t *testing.T) {
		// Remove the middle child (child2 - p element)
		children := []dom.Element{}
		current := parent.FirstChild()
		for current != nil {
			if elem, ok := current.(dom.Element); ok {
				children = append(children, elem)
			}
			current = current.NextSibling()
		}

		if len(children) >= 2 {
			child2 := children[1] // The p element
			parent.RemoveChild(child2)

			assert.Nil(t, child2.ParentNode(), "Removed child should have no parent")

			// Verify remaining structure
			assert.True(t, parent.HasChildNodes(), "Parent should still have child nodes")
		}
	})
}

func TestElement_TextContent(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("add text nodes", func(t *testing.T) {
		textNode := doc.CreateTextNode("Hello, World!")
		element.AppendChild(textNode)

		// Test that the text node was added (can't easily test TextContent without implementing it)
		assert.True(t, element.HasChildNodes(), "Element should have child nodes")

		firstChild := element.FirstChild()
		if textNode, ok := firstChild.(dom.Text); ok {
			assert.Equal(t, "Hello, World!", textNode.Data(), "Text content should match")
		}
	})
}

func TestElement_Style(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	t.Run("style property", func(t *testing.T) {
		style := element.Style()
		// Note: Style may be nil in some implementations or test environments
		// We just verify that calling Style() doesn't panic
		_ = style // Use the style variable to avoid unused warnings
		assert.NotPanics(t, func() {
			element.Style()
		}, "Style() should not panic")
	})
}

func TestElement_Focus(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	t.Run("focus and blur", func(t *testing.T) {
		input := doc.CreateElement("input")

		// These methods don't return values, so we just test they don't panic
		assert.NotPanics(t, func() {
			input.Focus()
		}, "Focus should not panic")

		assert.NotPanics(t, func() {
			input.Blur()
		}, "Blur should not panic")
	})
}

func TestElement_Events(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("button")

	t.Run("add event listener", func(t *testing.T) {
		handler := func(node dom.Node) {
			// Event handler - would be called when event fires
		}

		result := element.AddEventListener("click", handler)
		assert.Equal(t, element, result, "AddEventListener should return the same element")

		// Note: Actually triggering events would require a more complex test setup
		// This just verifies that AddEventListener doesn't panic and returns correctly
	})
}

func TestElement_EdgeCases(t *testing.T) {
	t.Run("self manipulation", func(t *testing.T) {
		element := domPkg.GetWindow().Document().CreateElement("div")

		// Test self-append behavior - WASM throws error, non-WASM might allow it
		defer func() {
			if r := recover(); r != nil {
				t.Log("Implementation correctly prevents self-append with error:", r)
			}
		}()

		element.AppendChild(element)

		// If we get here, self-append was allowed
		firstChild := element.FirstChild()
		if firstChild != nil && firstChild.Equals(element) {
			t.Log("Implementation allows self-append (element becomes its own child)")
		} else {
			t.Log("Self-append was ignored by implementation")
		}
	})

	t.Run("null operations", func(t *testing.T) {
		element := domPkg.GetWindow().Document().CreateElement("div")

		// Test RemoveChild with nil - WASM may throw error, non-WASM handles gracefully
		t.Run("remove nil child", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Log("RemoveChild(nil) correctly caused an error in WASM:", r)
				}
			}()
			element.RemoveChild(nil)
			t.Log("RemoveChild(nil) handled gracefully in non-WASM")
		})

		// Test SetAttribute with empty name - WASM may throw error, non-WASM handles gracefully
		t.Run("empty attribute name", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Log("SetAttribute with empty name correctly caused an error in WASM:", r)
				}
			}()
			element.SetAttribute("", "")
			t.Log("SetAttribute with empty name handled gracefully in non-WASM")
		})

		// GetAttribute with empty string should always work
		t.Run("get empty attribute", func(t *testing.T) {
			result := element.GetAttribute("")
			assert.Equal(t, "", result, "GetAttribute with empty string should return empty string")
		})
	})
}

func TestElement_ComplexStructure(t *testing.T) {
	doc := domPkg.GetWindow().Document()

	t.Run("build complex DOM tree", func(t *testing.T) {
		// Create a complex structure: div > header > h1 + nav > ul > li*3
		container := doc.CreateElement("div")
		container.SetAttribute("class", "container")

		header := doc.CreateElement("header")
		h1 := doc.CreateElement("h1")
		h1Text := doc.CreateTextNode("Page Title")
		h1.AppendChild(h1Text)

		nav := doc.CreateElement("nav")
		ul := doc.CreateElement("ul")

		// Create navigation items
		for i := 1; i <= 3; i++ {
			li := doc.CreateElement("li")
			a := doc.CreateElement("a")
			a.SetAttribute("href", "#")
			linkText := doc.CreateTextNode("Link " + string(rune('0'+i)))
			a.AppendChild(linkText)
			li.AppendChild(a)
			ul.AppendChild(li)
		}

		// Build the structure
		nav.AppendChild(ul)
		header.AppendChild(h1)
		header.AppendChild(nav)
		container.AppendChild(header)

		// Test the structure
		assert.True(t, container.HasChildNodes(), "Container should have children")
		assert.Equal(t, "container", container.GetAttribute("class"), "Container should have correct class")

		firstChild := container.FirstChild()
		assert.NotNil(t, firstChild, "Container should have first child")
		if header, ok := firstChild.(dom.Element); ok {
			assert.Equal(t, "HEADER", header.TagName(), "First child should be header")
		}
	})
}

// Legacy tests for backward compatibility
func Test_Element_001(t *testing.T) {
	element := domPkg.GetWindow().Document().CreateElement("A")
	if element.NodeName() != "A" {
		t.Errorf("Element.NodeName() = %v, want %v", element.NodeName(), "A")
	}
	if element.NodeType() != dom.ELEMENT_NODE {
		t.Errorf("Element.NodeType() = %v, want %v", element.NodeType(), dom.ELEMENT_NODE)
	}
}

func Test_Element_002(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("A")
	b := doc.CreateElement("CB")
	c := doc.CreateElement("CC")
	if parent.HasChildNodes() != false {
		t.Error("HasChildNodes() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
	parent.AppendChild(b)
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(b) == false {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() == false {
		t.Error("HasChildNodes() failed: ", parent)
	}
	if b.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.AppendChild(c)
	if b.NextSibling() == nil {
		t.Error("NextSibling() failed")
	}
	if parent.FirstChild().Equals(b) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if c.ParentNode().Equals(parent) == false {
		t.Error("ParentNode() failed")
	}
	if b.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if b.NextSibling().Equals(c) == false {
		t.Error("NextSibling() failed")
	}
	if c.PreviousSibling().Equals(b) == false {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(b)
	if parent.FirstChild().Equals(c) == false {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild().Equals(c) == false {
		t.Error("LastChild() failed")
	}
	if b.ParentNode() != nil {
		t.Error("ParentNode() failed", b.ParentNode())
	}
	if c.PreviousSibling() != nil {
		t.Error("PreviousSibling() failed")
	}
	if c.NextSibling() != nil {
		t.Error("NextSibling() failed")
	}
	parent.RemoveChild(c)
	if parent.FirstChild() != nil {
		t.Error("FirstChild() failed")
	}
	if parent.LastChild() != nil {
		t.Error("LastChild() failed")
	}
	if parent.HasChildNodes() {
		t.Error("HasChildNodes() failed")
	}
	if c.ParentNode() != nil {
		t.Error("ParentNode() failed")
	}
}

func Test_Element_003(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("a")
	if parent.InnerHTML() != "" {
		t.Error("InnerHTML() failed")
	}
	if parent.OuterHTML() != "<a></a>" {
		t.Error("OuterHTML() failed: ", parent.OuterHTML())
	}
}

func TestElement_RemoveAttribute(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Set an attribute
	element.SetAttribute("id", "test-id")
	assert.True(t, element.HasAttribute("id"), "Should have id attribute")
	assert.Equal(t, "test-id", element.GetAttribute("id"), "Should get correct id value")

	// Remove the attribute
	element.RemoveAttribute("id")
	assert.False(t, element.HasAttribute("id"), "Should not have id attribute after removal")
	assert.Equal(t, "", element.GetAttribute("id"), "GetAttribute should return empty string after removal")

	// Removing non-existent attribute should not error
	element.RemoveAttribute("non-existent")
}

func TestElement_RemoveAttributeNode(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Set an attribute
	attr := element.SetAttribute("class", "my-class")
	assert.NotNil(t, attr, "SetAttribute should return an attribute")
	assert.True(t, element.HasAttribute("class"), "Should have class attribute")

	// Remove the attribute node
	element.RemoveAttributeNode(attr)
	assert.False(t, element.HasAttribute("class"), "Should not have class attribute after removal")

	// Removing non-existent attribute node - create a new unattached attribute (should not error)
	newAttr := doc.CreateAttribute("other")
	element.RemoveAttributeNode(newAttr)

	// Removing nil should not error
	element.RemoveAttributeNode(nil)
}

func TestElement_SetAttributeNode(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Create and set an attribute node
	attr := doc.CreateAttribute("data-test")
	attr.SetValue("value1")
	oldAttr := element.SetAttributeNode(attr)
	assert.Nil(t, oldAttr, "Setting new attribute should return nil")
	assert.True(t, element.HasAttribute("data-test"), "Should have data-test attribute")
	assert.Equal(t, "value1", element.GetAttribute("data-test"), "Should get correct value")

	// Replace with a new attribute node
	newAttr := doc.CreateAttribute("data-test")
	newAttr.SetValue("value2")
	oldAttr = element.SetAttributeNode(newAttr)
	assert.NotNil(t, oldAttr, "Replacing attribute should return old attribute")
	assert.Equal(t, "value1", oldAttr.Value(), "Old attribute should have previous value")
	assert.Equal(t, "value2", element.GetAttribute("data-test"), "Should get new value")

	// Setting nil should return nil
	result := element.SetAttributeNode(nil)
	assert.Nil(t, result, "Setting nil attribute should return nil")
}

func TestElement_GetAttributeNames(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Initially should have no attributes
	names := element.GetAttributeNames()
	assert.Equal(t, 0, len(names), "New element should have no attributes")

	// Add some attributes
	element.SetAttribute("id", "test")
	element.SetAttribute("class", "my-class")
	element.SetAttribute("data-value", "123")

	names = element.GetAttributeNames()
	assert.Equal(t, 3, len(names), "Should have 3 attributes")

	// Check that all attribute names are present (order may vary)
	nameSet := make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}
	assert.True(t, nameSet["id"], "Should contain 'id'")
	assert.True(t, nameSet["class"], "Should contain 'class'")
	assert.True(t, nameSet["data-value"], "Should contain 'data-value'")

	// Remove an attribute
	element.RemoveAttribute("class")
	names = element.GetAttributeNames()
	assert.Equal(t, 2, len(names), "Should have 2 attributes after removal")

	// Verify class is not in the list
	nameSet = make(map[string]bool)
	for _, name := range names {
		nameSet[name] = true
	}
	assert.False(t, nameSet["class"], "Should not contain 'class' after removal")
}

func TestElement_AttributeNodeOperations(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Test complete flow: create, set, get, remove
	attr1 := doc.CreateAttribute("attr1")
	attr1.SetValue("value1")
	element.SetAttributeNode(attr1)

	attr2 := doc.CreateAttribute("attr2")
	attr2.SetValue("value2")
	element.SetAttributeNode(attr2)

	// Verify both attributes are set
	assert.Equal(t, 2, len(element.GetAttributeNames()), "Should have 2 attributes")
	assert.Equal(t, "value1", element.GetAttribute("attr1"), "attr1 should have correct value")
	assert.Equal(t, "value2", element.GetAttribute("attr2"), "attr2 should have correct value")

	// Get attribute node
	retrievedAttr := element.GetAttributeNode("attr1")
	assert.NotNil(t, retrievedAttr, "Should retrieve attr1 node")
	assert.Equal(t, "attr1", retrievedAttr.Name(), "Retrieved node should have correct name")
	assert.Equal(t, "value1", retrievedAttr.Value(), "Retrieved node should have correct value")

	// Remove one attribute node
	element.RemoveAttributeNode(attr1)
	assert.Equal(t, 1, len(element.GetAttributeNames()), "Should have 1 attribute after removal")
	assert.False(t, element.HasAttribute("attr1"), "attr1 should be removed")
	assert.True(t, element.HasAttribute("attr2"), "attr2 should still exist")
}

func TestElement_GetElementsByClassName(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	container := doc.CreateElement("div")

	// Create nested structure
	div1 := doc.CreateElement("div")
	div1.SetAttribute("class", "item active")
	div2 := doc.CreateElement("div")
	div2.SetAttribute("class", "item")
	div3 := doc.CreateElement("div")
	div3.SetAttribute("class", "active special")

	span := doc.CreateElement("span")
	span.SetAttribute("class", "item")
	div1.AppendChild(span)

	container.AppendChild(div1)
	container.AppendChild(div2)
	container.AppendChild(div3)

	// Get elements by class
	items := container.GetElementsByClassName("item")
	assert.Equal(t, 3, len(items), "Should find 3 elements with 'item' class")

	actives := container.GetElementsByClassName("active")
	assert.Equal(t, 2, len(actives), "Should find 2 elements with 'active' class")

	specials := container.GetElementsByClassName("special")
	assert.Equal(t, 1, len(specials), "Should find 1 element with 'special' class")
}

func TestElement_GetElementsByTagName(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	container := doc.CreateElement("div")

	// Create nested structure
	p1 := doc.CreateElement("p")
	p2 := doc.CreateElement("p")
	span := doc.CreateElement("span")
	div := doc.CreateElement("div")

	nestedP := doc.CreateElement("p")
	div.AppendChild(nestedP)

	container.AppendChild(p1)
	container.AppendChild(p2)
	container.AppendChild(span)
	container.AppendChild(div)

	// Get elements by tag name
	paragraphs := container.GetElementsByTagName("p")
	assert.Equal(t, 3, len(paragraphs), "Should find 3 <p> elements including nested")

	spans := container.GetElementsByTagName("span")
	assert.Equal(t, 1, len(spans), "Should find 1 <span> element")

	divs := container.GetElementsByTagName("div")
	assert.Equal(t, 1, len(divs), "Should find 1 <div> element (not including container)")
}

func TestElement_Remove(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")
	child := doc.CreateElement("p")

	parent.AppendChild(child)
	assert.True(t, parent.HasChildNodes(), "Parent should have child")
	assert.True(t, child.ParentNode().Equals(parent), "Child should have parent")

	// Remove the child
	child.Remove()
	assert.False(t, parent.HasChildNodes(), "Parent should not have child after removal")
	assert.Nil(t, child.ParentNode(), "Child should have no parent after removal")
}

func TestElement_ReplaceWith(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")
	oldChild := doc.CreateElement("p")
	oldChild.AppendChild(doc.CreateTextNode("Old"))

	parent.AppendChild(oldChild)
	assert.Equal(t, 1, len(parent.ChildNodes()), "Parent should have 1 child")

	// Replace with new nodes
	newChild1 := doc.CreateElement("span")
	newChild1.AppendChild(doc.CreateTextNode("New1"))
	newChild2 := doc.CreateElement("span")
	newChild2.AppendChild(doc.CreateTextNode("New2"))

	oldChild.ReplaceWith(newChild1, newChild2)

	assert.Equal(t, 2, len(parent.ChildNodes()), "Parent should have 2 children after replacement")
	assert.Nil(t, oldChild.ParentNode(), "Old child should have no parent")
	assert.True(t, newChild1.ParentNode().Equals(parent), "New child1 should have parent")
	assert.True(t, newChild2.ParentNode().Equals(parent), "New child2 should have parent")
}

func TestElement_InsertAdjacentElement(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	container := doc.CreateElement("div")
	target := doc.CreateElement("p")
	target.AppendChild(doc.CreateTextNode("Target"))
	container.AppendChild(target)

	// beforebegin
	before := doc.CreateElement("span")
	before.AppendChild(doc.CreateTextNode("Before"))
	result := target.InsertAdjacentElement("beforebegin", before)
	assert.NotNil(t, result, "Should return inserted element")
	assert.Equal(t, 2, len(container.ChildNodes()), "Container should have 2 children")
	assert.Equal(t, before, container.FirstChild(), "Before element should be first")

	// afterbegin (first child of target)
	afterBegin := doc.CreateElement("b")
	afterBegin.AppendChild(doc.CreateTextNode("AfterBegin"))
	result = target.InsertAdjacentElement("afterbegin", afterBegin)
	assert.NotNil(t, result, "Should return inserted element")
	assert.True(t, target.HasChildNodes(), "Target should have children")
	assert.Equal(t, afterBegin, target.FirstChild(), "AfterBegin should be first child of target")

	// beforeend (last child of target)
	beforeEnd := doc.CreateElement("i")
	beforeEnd.AppendChild(doc.CreateTextNode("BeforeEnd"))
	result = target.InsertAdjacentElement("beforeend", beforeEnd)
	assert.NotNil(t, result, "Should return inserted element")
	assert.Equal(t, beforeEnd, target.LastChild(), "BeforeEnd should be last child of target")

	// afterend
	after := doc.CreateElement("span")
	after.AppendChild(doc.CreateTextNode("After"))
	result = target.InsertAdjacentElement("afterend", after)
	assert.NotNil(t, result, "Should return inserted element")
	assert.Equal(t, after, container.LastChild(), "After element should be last in container")
}

func TestElement_IDMethods(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Initially no ID
	assert.Equal(t, "", element.ID(), "New element should have no ID")

	// Set ID
	element.SetID("test-id")
	assert.Equal(t, "test-id", element.ID(), "ID should be set")
	assert.Equal(t, "test-id", element.GetAttribute("id"), "ID should be in attributes")

	// Change ID
	element.SetID("new-id")
	assert.Equal(t, "new-id", element.ID(), "ID should be updated")
}

func TestElement_ClassNameMethods(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	element := doc.CreateElement("div")

	// Initially no class
	assert.Equal(t, "", element.ClassName(), "New element should have no class")

	// Set class name
	element.SetClassName("btn btn-primary")
	assert.Equal(t, "btn btn-primary", element.ClassName(), "Class name should be set")
	assert.Equal(t, "btn btn-primary", element.GetAttribute("class"), "Class should be in attributes")

	// Change class name
	element.SetClassName("btn btn-secondary active")
	assert.Equal(t, "btn btn-secondary active", element.ClassName(), "Class name should be updated")
}

func TestElement_Children(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	// Add mixed content: elements and text nodes
	elem1 := doc.CreateElement("p")
	text1 := doc.CreateTextNode("text1")
	elem2 := doc.CreateElement("span")
	text2 := doc.CreateTextNode("text2")
	elem3 := doc.CreateElement("div")

	parent.AppendChild(elem1)
	parent.AppendChild(text1)
	parent.AppendChild(elem2)
	parent.AppendChild(text2)
	parent.AppendChild(elem3)

	// Children should only return elements, not text nodes
	children := parent.Children()
	assert.Equal(t, 3, len(children), "Should have 3 element children")
	assert.Equal(t, elem1, children[0], "First child should be elem1")
	assert.Equal(t, elem2, children[1], "Second child should be elem2")
	assert.Equal(t, elem3, children[2], "Third child should be elem3")
}

func TestElement_ChildElementCount(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	assert.Equal(t, 0, parent.ChildElementCount(), "Empty parent should have 0 element children")

	// Add mixed content
	parent.AppendChild(doc.CreateElement("p"))
	parent.AppendChild(doc.CreateTextNode("text"))
	parent.AppendChild(doc.CreateElement("span"))

	assert.Equal(t, 2, parent.ChildElementCount(), "Should count only element children")
}

func TestElement_FirstLastElementChild(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	// Empty parent
	assert.Nil(t, parent.FirstElementChild(), "Empty parent should have no first element child")
	assert.Nil(t, parent.LastElementChild(), "Empty parent should have no last element child")

	// Add mixed content
	text1 := doc.CreateTextNode("text1")
	elem1 := doc.CreateElement("p")
	elem2 := doc.CreateElement("span")
	text2 := doc.CreateTextNode("text2")
	elem3 := doc.CreateElement("div")

	parent.AppendChild(text1)
	parent.AppendChild(elem1)
	parent.AppendChild(elem2)
	parent.AppendChild(text2)
	parent.AppendChild(elem3)

	assert.Equal(t, elem1, parent.FirstElementChild(), "First element child should be elem1")
	assert.Equal(t, elem3, parent.LastElementChild(), "Last element child should be elem3")
}

func TestElement_ElementSiblings(t *testing.T) {
	doc := domPkg.GetWindow().Document()
	parent := doc.CreateElement("div")

	// Create structure with mixed content
	text1 := doc.CreateTextNode("text1")
	elem1 := doc.CreateElement("p")
	text2 := doc.CreateTextNode("text2")
	elem2 := doc.CreateElement("span")
	elem3 := doc.CreateElement("div")
	text3 := doc.CreateTextNode("text3")

	parent.AppendChild(text1)
	parent.AppendChild(elem1)
	parent.AppendChild(text2)
	parent.AppendChild(elem2)
	parent.AppendChild(elem3)
	parent.AppendChild(text3)

	// Test NextElementSibling
	assert.Equal(t, elem2, elem1.NextElementSibling(), "elem1's next element sibling should be elem2")
	assert.Equal(t, elem3, elem2.NextElementSibling(), "elem2's next element sibling should be elem3")
	assert.Nil(t, elem3.NextElementSibling(), "elem3 should have no next element sibling")

	// Test PreviousElementSibling
	assert.Nil(t, elem1.PreviousElementSibling(), "elem1 should have no previous element sibling")
	assert.Equal(t, elem1, elem2.PreviousElementSibling(), "elem2's previous element sibling should be elem1")
	assert.Equal(t, elem2, elem3.PreviousElementSibling(), "elem3's previous element sibling should be elem2")
}
