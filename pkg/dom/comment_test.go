package dom

import (
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestCreateComment(t *testing.T) {
	doc := getDocument()
	if doc == nil {
		t.Fatal("Expected document, got nil")
	}

	comment := doc.CreateComment("This is a comment")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}
}

func TestComment_Data(t *testing.T) {
	doc := getDocument()
	expected := "Test comment data"

	comment := doc.CreateComment(expected)
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.Data(); got != expected {
		t.Errorf("Expected Data() = %q, got %q", expected, got)
	}
}

func TestComment_Length(t *testing.T) {
	doc := getDocument()
	data := "Hello comment"
	expected := len(data)

	comment := doc.CreateComment(data)
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.Length(); got != expected {
		t.Errorf("Expected Length() = %d, got %d", expected, got)
	}
}

func TestComment_NodeName(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.NodeName(); got != "#comment" {
		t.Errorf("Expected NodeName() = %q, got %q", "#comment", got)
	}
}

func TestComment_NodeType(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.NodeType(); got != COMMENT_NODE {
		t.Errorf("Expected NodeType() = %d, got %d", COMMENT_NODE, got)
	}
}

func TestComment_TextContent(t *testing.T) {
	doc := getDocument()
	expected := "Comment content"

	comment := doc.CreateComment(expected)
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.TextContent(); got != expected {
		t.Errorf("Expected TextContent() = %q, got %q", expected, got)
	}
}

func TestComment_HasChildNodes(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if comment.HasChildNodes() {
		t.Error("Comment node should not have child nodes")
	}
}

func TestComment_ChildNodes(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	children := comment.ChildNodes()
	if len(children) > 0 {
		t.Errorf("Expected no child nodes, got %d", len(children))
	}
}

func TestComment_FirstChild(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if child := comment.FirstChild(); child != nil {
		t.Error("Comment node should not have first child")
	}
}

func TestComment_LastChild(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if child := comment.LastChild(); child != nil {
		t.Error("Comment node should not have last child")
	}
}

func TestComment_Contains(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	otherComment := doc.CreateComment("other")
	if comment.Contains(otherComment) {
		t.Error("Comment node should not contain other nodes")
	}
}

func TestComment_CloneNode(t *testing.T) {
	doc := getDocument()
	original := doc.CreateComment("original comment")
	if original == nil {
		t.Fatal("Expected comment node, got nil")
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

func TestComment_OwnerDocument(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	owner := comment.OwnerDocument()
	if owner == nil {
		t.Error("Expected owner document, got nil")
	}
}

func TestComment_AppendChild_Panics(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	otherComment := doc.CreateComment("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling AppendChild on comment node")
		}
	}()

	comment.AppendChild(otherComment)
}

func TestComment_InsertBefore_Panics(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	otherComment := doc.CreateComment("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling InsertBefore on comment node")
		}
	}()

	comment.InsertBefore(otherComment, nil)
}

func TestComment_RemoveChild_Panics(t *testing.T) {
	doc := getDocument()
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	otherComment := doc.CreateComment("other")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when calling RemoveChild on comment node")
		}
	}()

	comment.RemoveChild(otherComment)
}

func TestComment_Equals(t *testing.T) {
	doc := getDocument()
	comment1 := doc.CreateComment("test")
	comment2 := doc.CreateComment("test")

	if comment1 == nil || comment2 == nil {
		t.Fatal("Expected comment nodes, got nil")
	}

	// Same content but different nodes
	if comment1.Equals(comment2) {
		t.Error("Different comment nodes should not be equal")
	}

	// Same node should equal itself
	if !comment1.Equals(comment1) {
		t.Error("Comment node should equal itself")
	}
}
