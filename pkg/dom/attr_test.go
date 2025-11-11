package dom_test

import (
	"fmt"
	"testing"

	// Packages
	core "github.com/djthorpe/go-wasmbuild"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
)

func Test_Attr_001(t *testing.T) {
	doc := mustDocument(t)
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"name", "", `name=""`},
		{"name", "test", `name="test"`},
		{"name", "&", `name="&amp;"`},
		{"name", "<test>", `name="&lt;test&gt;"`},
		{"name", `"test"`, `name="&#34;test&#34;"`},
		{"name", `'test'`, `name="&#39;test&#39;"`},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("String:%s=%q", test.name, test.value), func(t *testing.T) {
			attr := doc.CreateAttribute(test.name)
			if attr == nil {
				t.Fatal("expected attribute, got nil")
			}
			if attr.Name() != test.name {
				t.Fatalf("expected name %q, got %q", test.name, attr.Name())
			}
			if attr.Value() != "" {
				t.Fatalf("expected empty value, got %q", attr.Value())
			}
			if test.value != "" {
				attr.SetValue(test.value)
				if attr.Value() != test.value {
					t.Fatalf("expected value %q, got %q", test.value, attr.Value())
				}
			}
			if str := fmt.Sprint(attr); str != test.expected {
				t.Fatalf("expected %q, got %q", test.expected, str)
			}
		})
	}
}

func TestAttrOwnerElementLifecycle(t *testing.T) {
	doc := mustDocument(t)
	element := mustElement(t, doc, "div")
	attr := doc.CreateAttribute("data-id")
	if attr == nil {
		t.Fatal("expected attribute, got nil")
	}
	if attr.OwnerElement() != nil {
		t.Fatalf("expected nil owner before attachment")
	}
	attr.SetValue("123")
	element.SetAttributeNode(attr)
	owner := attr.OwnerElement()
	if owner == nil {
		t.Fatalf("expected owner element after attachment")
	}
	if !owner.Equals(element) {
		t.Fatalf("owner element mismatch")
	}
	if got := element.GetAttribute("data-id"); got != "123" {
		t.Fatalf("expected attribute value '123', got %q", got)
	}
	if fmt.Sprint(attr) != `data-id="123"` {
		t.Fatalf("expected serialised attribute, got %q", fmt.Sprint(attr))
	}
	element.RemoveAttribute("data-id")
	if attr.OwnerElement() != nil {
		t.Fatalf("expected owner element to be nil after removal")
	}
	if element.HasAttribute("data-id") {
		t.Fatalf("expected element to have attribute removed")
	}
}

func TestAttrReplacement(t *testing.T) {
	doc := mustDocument(t)
	element := mustElement(t, doc, "span")
	first := doc.CreateAttribute("data-value")
	first.SetValue("first")
	element.SetAttributeNode(first)
	second := doc.CreateAttribute("data-value")
	second.SetValue("second")
	element.SetAttributeNode(second)
	if !second.OwnerElement().Equals(element) {
		t.Fatalf("expected second attribute owner to match element")
	}
	if first.OwnerElement() != nil {
		t.Fatalf("expected first attribute to be detached after replacement")
	}
	if got := element.GetAttribute("data-value"); got != "second" {
		t.Fatalf("expected element attribute to be 'second', got %q", got)
	}
}

func TestAttrSetValueReflectsOnElement(t *testing.T) {
	doc := mustDocument(t)
	element := mustElement(t, doc, "section")
	attr := element.SetAttribute("data-role", "button")
	if attr == nil {
		t.Fatal("expected attribute from SetAttribute")
	}
	if attr.OwnerElement() == nil {
		t.Fatalf("expected owner element after SetAttribute")
	}
	attr.SetValue("dialog")
	if got := element.GetAttribute("data-role"); got != "dialog" {
		t.Fatalf("expected attribute value to reflect set value, got %q", got)
	}
	if attr.Value() != "dialog" {
		t.Fatalf("expected attr value to be updated, got %q", attr.Value())
	}
	if fmt.Sprint(attr) != `data-role="dialog"` {
		t.Fatalf("expected serialised attribute, got %q", fmt.Sprint(attr))
	}
}

func mustDocument(tb testing.TB) core.Document {
	tb.Helper()
	win := dom.GetWindow()
	if win == nil {
		tb.Fatal("window is unavailable")
	}
	doc := win.Document()
	if doc == nil {
		tb.Fatal("document is unavailable")
	}
	return doc
}

func mustElement(tb testing.TB, doc core.Document, tag string) core.Element {
	tb.Helper()
	element := doc.CreateElement(tag)
	if element == nil {
		tb.Fatalf("expected element %q", tag)
	}
	return element
}
