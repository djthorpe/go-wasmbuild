package dom

import (
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

func getDocument() Document {
	return NewWindow().Document()
}

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestCreateTextNode(t *testing.T) {
	doc := getDocument()
	if doc == nil {
		t.Fatal("Expected document, got nil")
	}

	text := doc.CreateTextNode("Hello, World!")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}
}

func TestText_Data(t *testing.T) {
	doc := getDocument()
	expected := "Test data"

	text := doc.CreateTextNode(expected)
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if got := text.Data(); got != expected {
		t.Errorf("Expected Data() = %q, got %q", expected, got)
	}
}

func TestText_Length(t *testing.T) {
	doc := getDocument()
	data := "Hello"
	expected := len(data)

	text := doc.CreateTextNode(data)
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if got := text.Length(); got != expected {
		t.Errorf("Expected Length() = %d, got %d", expected, got)
	}
}

func TestText_NodeName(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if got := text.NodeName(); got != "#text" {
		t.Errorf("Expected NodeName() = %q, got %q", "#text", got)
	}
}

func TestText_NodeType(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if got := text.NodeType(); got != TEXT_NODE {
		t.Errorf("Expected NodeType() = %d, got %d", TEXT_NODE, got)
	}
}

func TestText_TextContent(t *testing.T) {
	doc := getDocument()
	expected := "Text content"

	text := doc.CreateTextNode(expected)
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if got := text.TextContent(); got != expected {
		t.Errorf("Expected TextContent() = %q, got %q", expected, got)
	}
}

func TestText_HasChildNodes(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if text.HasChildNodes() {
		t.Error("Text node should not have child nodes")
	}
}

func TestText_ChildNodes(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	children := text.ChildNodes()
	if len(children) > 0 {
		t.Errorf("Expected no child nodes, got %d", len(children))
	}
}

func TestText_FirstChild(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if child := text.FirstChild(); child != nil {
		t.Error("Text node should not have first child")
	}
}

func TestText_LastChild(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	if child := text.LastChild(); child != nil {
		t.Error("Text node should not have last child")
	}
}

func TestText_Contains(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	otherText := doc.CreateTextNode("other")
	if text.Contains(otherText) {
		t.Error("Text node should not contain other nodes")
	}
}

func TestText_CloneNode(t *testing.T) {
	doc := getDocument()
	original := doc.CreateTextNode("original text")
	if original == nil {
		t.Fatal("Expected text node, got nil")
	}

	cloned := original.CloneNode(false)
	if cloned == nil {
		t.Fatal("Expected cloned node, got nil")
	}

	if cloned.TextContent() != original.TextContent() {
		t.Errorf("Expected cloned text = %q, got %q", original.TextContent(), cloned.TextContent())
	}

	// Ensure they are different nodes
	if cloned.Equals(original) {
		t.Error("Cloned node should not equal original node")
	}
}

func TestText_OwnerDocument(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	owner := text.OwnerDocument()
	if owner == nil {
		t.Error("Expected owner document, got nil")
	}
}

func TestText_AppendChild_Panics(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	otherText := doc.CreateTextNode("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling AppendChild on text node")
		}
	}()

	text.AppendChild(otherText)
}

func TestText_InsertBefore_Panics(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	otherText := doc.CreateTextNode("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling InsertBefore on text node")
		}
	}()

	text.InsertBefore(otherText, nil)
}

func TestText_RemoveChild_Panics(t *testing.T) {
	doc := getDocument()
	text := doc.CreateTextNode("test")
	if text == nil {
		t.Fatal("Expected text node, got nil")
	}

	otherText := doc.CreateTextNode("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling RemoveChild on text node")
		}
	}()

	text.RemoveChild(otherText)
}

func TestText_Equals(t *testing.T) {
	doc := getDocument()
	text1 := doc.CreateTextNode("test")
	text2 := doc.CreateTextNode("test")

	if text1 == nil || text2 == nil {
		t.Fatal("Expected text nodes, got nil")
	}

	// Same content but different nodes
	if text1.Equals(text2) {
		t.Error("Different text nodes should not be equal")
	}

	// Same node should equal itself
	if !text1.Equals(text1) {
		t.Error("Text node should equal itself")
	}
}
