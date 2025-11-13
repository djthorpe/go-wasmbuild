package dom_test

import (
	"bytes"
	"fmt"
	"html"
	"strings"
	"testing"

	// Package imports
	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestCreateComment(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("This is a comment")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}
}

func TestComment_Data(t *testing.T) {
	doc := mustDocument(t)
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
	doc := mustDocument(t)
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
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.NodeName(); got != "#comment" {
		t.Errorf("Expected NodeName() = %q, got %q", "#comment", got)
	}
}

func TestComment_NodeType(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if got := comment.NodeType(); got != COMMENT_NODE {
		t.Errorf("Expected NodeType() = %d, got %d", COMMENT_NODE, got)
	}
}

func TestComment_TextContent(t *testing.T) {
	doc := mustDocument(t)
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
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if comment.HasChildNodes() {
		t.Error("Comment node should not have child nodes")
	}
}

func TestComment_ChildNodes(t *testing.T) {
	doc := mustDocument(t)
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
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if child := comment.FirstChild(); child != nil {
		t.Error("Comment node should not have first child")
	}
}

func TestComment_LastChild(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	if child := comment.LastChild(); child != nil {
		t.Error("Comment node should not have last child")
	}
}

func TestComment_Contains(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	otherComment := doc.CreateComment("other")
	if comment.Contains(otherComment) {
		t.Error("Comment node should not contain other nodes")
	}
}

func TestComment_OwnerDocument(t *testing.T) {
	doc := mustDocument(t)
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
	doc := mustDocument(t)
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
	doc := mustDocument(t)
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
	doc := mustDocument(t)
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
	doc := mustDocument(t)
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

func TestComment_Write(t *testing.T) {
	assert := assert.New(t)
	doc := mustDocument(t)

	var buf bytes.Buffer
	comment := doc.CreateComment("test")
	n, err := comment.Write(&buf)
	assert.NoError(err)
	assert.NotZero(n)
	assert.Equal("<!--test-->", buf.String())
}

func TestComment_WriteEscapesContent(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment(`<dangerous "comment"> & more`)

	var buf bytes.Buffer
	if _, err := comment.Write(&buf); err != nil {
		t.Fatalf("expected write to succeed: %v", err)
	}

	output := buf.String()
	if !strings.HasPrefix(output, "<!--") || !strings.HasSuffix(output, "-->") {
		t.Fatalf("expected comment delimiters, got %q", output)
	}

	escaped := html.EscapeString(`<dangerous "comment"> & more`)
	if got := strings.TrimSuffix(strings.TrimPrefix(output, "<!--"), "-->"); got != escaped {
		t.Fatalf("expected escaped payload %q, got %q", escaped, got)
	}
}

func TestComment_EqualsNil(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("test")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}
	if comment.Equals(nil) {
		t.Fatal("expected Equals(nil) to return false")
	}
}

func TestComment_StringMatchesWrite(t *testing.T) {
	doc := mustDocument(t)
	comment := doc.CreateComment("content & more")
	if comment == nil {
		t.Fatal("Expected comment node, got nil")
	}

	fromString := fmt.Sprint(comment)
	var buf bytes.Buffer
	if _, err := comment.Write(&buf); err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	if buf.String() != fromString {
		t.Fatalf("String() and Write() output mismatch: %q vs %q", fromString, buf.String())
	}
}

func TestComment_DetachKeepsParentNil(t *testing.T) {
	doc := mustDocument(t)
	elem := doc.CreateElement("div")
	elemComment := doc.CreateComment("child")

	// Append and remove via element API
	elem.AppendChild(elemComment)
	elem.RemoveChild(elemComment)

	if elemComment.ParentNode() != nil {
		t.Fatalf("expected comment parent to be nil after removal")
	}
}
