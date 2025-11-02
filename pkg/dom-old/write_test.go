package dom_test

import (
	"bytes"
	"strings"
	"testing"

	domPkg "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func TestWrite_Document(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	// Add content to body
	body := doc.Body()
	if body != nil {
		p := doc.CreateElement("p")
		p.AppendChild(doc.CreateTextNode("Test content"))
		body.AppendChild(p)
	}

	buf := new(bytes.Buffer)
	n, err := win.Write(buf, doc)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output (%d bytes):\n%s", n, output)

	// Verify DOCTYPE is present
	if !strings.Contains(output, "<!DOCTYPE html>") {
		t.Error("Output missing DOCTYPE")
	}

	// Verify basic structure
	if !strings.Contains(output, "<html>") {
		t.Error("Output missing <html>")
	}

	// Verify our added content
	if body != nil && !strings.Contains(output, "<p>Test content</p>") {
		t.Error("Output missing test content")
	}
}

func TestWrite_Element(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	div := doc.CreateElement("div")
	div.SetAttribute("class", "container")
	div.SetAttribute("id", "main")
	p := doc.CreateElement("p")
	p.AppendChild(doc.CreateTextNode("Hello"))
	div.AppendChild(p)

	buf := new(bytes.Buffer)
	_, err := win.Write(buf, div)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	t.Logf("Output: %s", output)

	// Verify attributes
	if !strings.Contains(output, "class=\"container\"") {
		t.Error("Output missing class attribute")
	}
	if !strings.Contains(output, "id=\"main\"") {
		t.Error("Output missing id attribute")
	}
	if !strings.Contains(output, "<p>Hello</p>") {
		t.Error("Output missing paragraph content")
	}
}

func TestWrite_TextNode(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	text := doc.CreateTextNode("Hello, World!")
	buf := new(bytes.Buffer)
	_, err := win.Write(buf, text)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	if output != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", output)
	}
}

func TestWrite_Comment(t *testing.T) {
	win := domPkg.GetWindow()
	doc := win.Document()

	comment := doc.CreateComment("This is a comment")
	buf := new(bytes.Buffer)
	_, err := win.Write(buf, comment)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "<!--This is a comment-->") {
		t.Errorf("Expected comment, got '%s'", output)
	}
}

func TestWrite_ErrorCases(t *testing.T) {
	win := domPkg.GetWindow()

	buf := new(bytes.Buffer)
	_, err := win.Write(buf, nil)
	if err == nil {
		t.Error("Expected error when writing nil node")
	}
}
