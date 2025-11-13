package dom_test

import (
	"fmt"
	"strings"
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestDocument_NodeMetadata(t *testing.T) {
	doc := mustDocument(t)

	if name := doc.NodeName(); name != "#document" {
		t.Fatalf("expected node name #document, got %q", name)
	}
	if nodeType := doc.NodeType(); nodeType != DOCUMENT_NODE {
		t.Fatalf("expected node type DOCUMENT_NODE, got %v", nodeType)
	}
	if owner := doc.OwnerDocument(); owner != nil {
		t.Fatalf("expected owner document to be nil, got %v", owner)
	}
	if doc.ParentElement() != nil {
		t.Fatal("expected parent element to be nil")
	}
	if !doc.HasChildNodes() {
		t.Fatal("expected document to report child nodes")
	}
}

func TestDocument_HeadAndBodyAvailable(t *testing.T) {
	doc := mustDocument(t)

	head := doc.Head()
	if head == nil {
		t.Fatal("expected head element")
	}
	if head.TagName() != "HEAD" {
		t.Fatalf("expected HEAD tag, got %q", head.TagName())
	}
	if owner := head.OwnerDocument(); owner == nil || owner.NodeType() != DOCUMENT_NODE {
		t.Fatal("expected head owner to be the document")
	}

	body := doc.Body()
	if body == nil {
		t.Fatal("expected body element")
	}
	if body.TagName() != "BODY" {
		t.Fatalf("expected BODY tag, got %q", body.TagName())
	}
	if owner := body.OwnerDocument(); owner == nil || owner.NodeType() != DOCUMENT_NODE {
		t.Fatal("expected body owner to be the document")
	}
}

func TestDocument_CreateElementLifecycle(t *testing.T) {
	doc := mustDocument(t)
	body := doc.Body()
	if body == nil {
		t.Fatal("expected body element")
	}

	section := doc.CreateElement("section")
	if section == nil {
		t.Fatal("expected section element")
	}
	if owner := section.OwnerDocument(); owner == nil || owner.NodeType() != DOCUMENT_NODE {
		t.Fatal("expected section owner to be the document")
	}
	if section.ParentNode() != nil {
		t.Fatal("expected new element to have no parent")
	}

	content := fmt.Sprintf("document-section-content-%s", t.Name())
	text := doc.CreateTextNode(content)
	if text == nil {
		t.Fatal("expected text node")
	}
	if text.ParentNode() != nil {
		t.Fatal("expected new text node to have no parent")
	}
	section.AppendChild(text)
	if text.ParentNode() == nil {
		t.Fatal("expected text node to be attached after append")
	}

	body.AppendChild(section)
	appended := true
	defer func() {
		if appended {
			body.RemoveChild(section)
		}
	}()

	if !body.Contains(section) {
		t.Fatal("expected body to contain appended section")
	}
	if parent := section.ParentNode(); parent == nil || !parent.Equals(body) {
		t.Fatal("expected section parent to be body")
	}

	body.RemoveChild(section)
	appended = false

	if section.ParentNode() != nil {
		t.Fatal("expected section parent to be nil after removal")
	}
	if parent := text.ParentNode(); parent == nil || !parent.Equals(section) {
		t.Fatal("expected text to remain attached to section after removal")
	}
	if body.Contains(section) {
		t.Fatal("expected body to no longer contain section after removal")
	}
}

func TestDocument_TextContentTracksBodyChanges(t *testing.T) {
	doc := mustDocument(t)
	body := doc.Body()
	if body == nil {
		t.Fatal("expected body element")
	}

	unique := uniqueDocumentText(doc, t.Name())

	paragraph := doc.CreateElement("p")
	if paragraph == nil {
		t.Fatal("expected paragraph element")
	}
	text := doc.CreateTextNode(unique)
	paragraph.AppendChild(text)

	body.AppendChild(paragraph)
	appended := true
	defer func() {
		if appended {
			body.RemoveChild(paragraph)
		}
	}()

	if !strings.Contains(body.TextContent(), unique) {
		t.Fatalf("expected body text content to include %q", unique)
	}

	body.RemoveChild(paragraph)
	appended = false

	if strings.Contains(body.TextContent(), unique) {
		t.Fatalf("expected body text content to drop %q after removal", unique)
	}
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

func uniqueDocumentText(doc Document, seed string) string {
	base := fmt.Sprintf("go-wasmbuild-%s", seed)
	body := doc.Body()
	if body == nil {
		return base
	}
	if !strings.Contains(body.TextContent(), base) {
		return base
	}
	for i := 0; ; i++ {
		candidate := fmt.Sprintf("%s-%d", base, i)
		if !strings.Contains(body.TextContent(), candidate) {
			return candidate
		}
	}
}
