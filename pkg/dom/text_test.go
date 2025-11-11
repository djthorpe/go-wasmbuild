package dom_test

import (
	"bytes"
	"fmt"
	"html"
	"runtime"
	"testing"

	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestText_BasicProperties(t *testing.T) {
	doc := mustDocument(t)
	data := "go-wasmbuild-text"
	text := mustText(t, doc, data)

	if text.NodeName() != "#text" {
		t.Fatalf("expected node name #text, got %q", text.NodeName())
	}
	if text.NodeType() != TEXT_NODE {
		t.Fatalf("expected text node type, got %v", text.NodeType())
	}
	if text.Data() != data {
		t.Fatalf("expected data %q, got %q", data, text.Data())
	}
	if text.Length() != len(data) {
		t.Fatalf("expected length %d, got %d", len(data), text.Length())
	}
	if text.TextContent() != data {
		t.Fatalf("expected text content %q, got %q", data, text.TextContent())
	}
	if owner := text.OwnerDocument(); owner == nil || owner.NodeType() != DOCUMENT_NODE {
		t.Fatal("expected owner document with document node type")
	}
	if text.ParentNode() != nil {
		t.Fatal("expected detached text node to have no parent")
	}
	if text.HasChildNodes() {
		t.Fatal("expected text node to have no children")
	}
	if len(text.ChildNodes()) != 0 {
		t.Fatalf("expected zero child nodes, got %d", len(text.ChildNodes()))
	}
	if text.FirstChild() != nil {
		t.Fatal("expected no first child for text node")
	}
	if text.LastChild() != nil {
		t.Fatal("expected no last child for text node")
	}
	other := mustText(t, doc, data+"-other")
	if text.Contains(other) {
		t.Fatal("expected text node not to contain unrelated nodes")
	}
}

func TestText_StringAndWriteEscaping(t *testing.T) {
	if runtime.GOARCH == "wasm" {
		t.Skip("Write is not implemented for wasm text nodes")
	}

	doc := mustDocument(t)
	raw := `text <with> risky & "characters"`
	text := mustText(t, doc, raw)

	var buf bytes.Buffer
	written, err := text.Write(&buf)
	if err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	escaped := html.EscapeString(raw)
	assert.Equal(t, len(escaped), written, "expected escape-aware byte count")
	assert.Equal(t, escaped, buf.String(), "expected escaped write output")
	assert.Equal(t, escaped, fmt.Sprint(text), "expected fmt.Sprint to mirror escaped write output")
}

func TestText_AttachDetachLifecycle(t *testing.T) {
	doc := mustDocument(t)
	parent := mustElement(t, doc, "section")
	text := mustText(t, doc, "attached")

	parent.AppendChild(text)

	if parent.FirstChild() == nil || !parent.FirstChild().Equals(text) {
		t.Fatal("expected text to be first child after append")
	}
	if !parent.Contains(text) {
		t.Fatal("expected parent to contain text after append")
	}
	if text.ParentNode() == nil || !text.ParentNode().Equals(parent) {
		t.Fatal("expected text parent to be element after append")
	}
	if parent.TextContent() != text.Data() {
		t.Fatalf("expected parent text content to match child data, got %q", parent.TextContent())
	}

	parent.RemoveChild(text)

	if parent.FirstChild() != nil {
		t.Fatal("expected parent to have no children after removal")
	}
	if parent.Contains(text) {
		t.Fatal("expected parent to no longer contain text after removal")
	}
	if text.ParentNode() != nil {
		t.Fatal("expected text parent to be nil after removal")
	}
}

func TestText_SiblingNavigation(t *testing.T) {
	doc := mustDocument(t)
	parent := mustElement(t, doc, "div")
	first := mustText(t, doc, "alpha")
	second := mustText(t, doc, "beta")
	third := mustText(t, doc, "gamma")

	parent.AppendChild(first)
	parent.AppendChild(second)
	parent.AppendChild(third)

	if next := first.NextSibling(); next == nil || !next.Equals(second) {
		t.Fatal("expected first.NextSibling to be second")
	}
	if prev := second.PreviousSibling(); prev == nil || !prev.Equals(first) {
		t.Fatal("expected second.PreviousSibling to be first")
	}
	if next := second.NextSibling(); next == nil || !next.Equals(third) {
		t.Fatal("expected second.NextSibling to be third")
	}
	if first.PreviousSibling() != nil {
		t.Fatal("expected first.PreviousSibling to be nil")
	}
	if third.NextSibling() != nil {
		t.Fatal("expected third.NextSibling to be nil")
	}
}

func TestText_TextContentAggregatesInParent(t *testing.T) {
	doc := mustDocument(t)
	parent := mustElement(t, doc, "p")
	first := mustText(t, doc, "alpha")
	second := mustText(t, doc, "beta")

	parent.AppendChild(first)
	parent.AppendChild(second)

	expected := first.Data() + second.Data()
	if parent.TextContent() != expected {
		t.Fatalf("expected combined text content %q, got %q", expected, parent.TextContent())
	}

	parent.RemoveChild(second)
	if parent.TextContent() != first.Data() {
		t.Fatalf("expected text content to revert to %q, got %q", first.Data(), parent.TextContent())
	}
}

func TestText_DisallowedChildOperations(t *testing.T) {
	doc := mustDocument(t)
	text := mustText(t, doc, "parent")
	other := mustText(t, doc, "other")

	tests := []struct {
		name string
		fn   func()
	}{
		{"AppendChild", func() { text.AppendChild(other) }},
		{"InsertBefore", func() { text.InsertBefore(other, nil) }},
		{"RemoveChild", func() { text.RemoveChild(other) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Fatalf("expected %s to panic on text node", tc.name)
				}
			}()
			tc.fn()
		})
	}
}

func TestText_EqualsVariants(t *testing.T) {
	doc := mustDocument(t)
	text := mustText(t, doc, t.Name())
	other := mustText(t, doc, t.Name())

	if !text.Equals(text) {
		t.Fatal("expected text to equal itself")
	}
	if text.Equals(other) {
		t.Fatal("expected distinct text nodes not to be equal")
	}
	if text.Equals(nil) {
		t.Fatal("expected text.Equals(nil) to be false")
	}
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

func mustText(tb testing.TB, doc Document, data string) Text {
	tb.Helper()
	text := doc.CreateTextNode(data)
	if text == nil {
		tb.Fatalf("expected text node for %q", data)
	}
	return text
}
