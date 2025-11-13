package bootstrap

import (
	"strings"
	"testing"

	. "github.com/djthorpe/go-wasmbuild"
)

func TestCard_Basic(t *testing.T) {
	card := Card()

	// Check element type
	if tag := card.Element().TagName(); tag != "DIV" {
		t.Errorf("Expected tag name 'DIV', got '%s'", tag)
	}

	// Check card class
	if class := card.Element().GetAttribute("class"); !strings.Contains(class, "card") {
		t.Errorf("Expected class to contain 'card', got '%s'", class)
	}

	// Check that card-body is automatically created
	children := card.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child (card-body), got %d", len(children))
	}

	if len(children) > 0 {
		bodyClass := children[0].(Element).GetAttribute("class")
		if bodyClass != "card-body" {
			t.Errorf("Expected card-body class, got '%s'", bodyClass)
		}
	}
}

func TestCard_WithContent(t *testing.T) {
	card := Card().Append("Card content")

	// Check that content is in card root
	children := card.Element().ChildNodes()
	if len(children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(children))
	}

	if len(children) > 0 {
		textContent := children[0].TextContent()
		if textContent != "Card content" {
			t.Errorf("Expected 'Card content', got '%s'", textContent)
		}
	}
}

func TestCard_WithHeading(t *testing.T) {
	card := Card().Header("Featured")

	// Check that we have header and body
	children := card.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (header + body), got %d", len(children))
	}

	if len(children) >= 2 {
		// Check header
		headerClass := children[0].(Element).GetAttribute("class")
		if headerClass != "card-header" {
			t.Errorf("Expected card-header class, got '%s'", headerClass)
		}

		headerText := children[0].TextContent()
		if headerText != "Featured" {
			t.Errorf("Expected 'Featured', got '%s'", headerText)
		}

		// Check body
		bodyClass := children[1].(Element).GetAttribute("class")
		if bodyClass != "card-body" {
			t.Errorf("Expected card-body class, got '%s'", bodyClass)
		}
	}
}

func TestCard_WithFooter(t *testing.T) {
	card := Card().Footer("Last updated 3 mins ago")

	// Check that we have body and footer
	children := card.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (body + footer), got %d", len(children))
	}

	if len(children) >= 2 {
		// Check body
		bodyClass := children[0].(Element).GetAttribute("class")
		if bodyClass != "card-body" {
			t.Errorf("Expected card-body class, got '%s'", bodyClass)
		}

		// Check footer
		footerClass := children[1].(Element).GetAttribute("class")
		if footerClass != "card-footer" {
			t.Errorf("Expected card-footer class, got '%s'", footerClass)
		}

		footerText := children[1].TextContent()
		if footerText != "Last updated 3 mins ago" {
			t.Errorf("Expected 'Last updated 3 mins ago', got '%s'", footerText)
		}
	}
}

func TestCard_WithHeadingAndFooter(t *testing.T) {
	c := Card().Header("Featured")
	c.Append("Card content")
	c.Footer("2 days ago")

	// Check that we have header, content, and footer
	children := c.Element().ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children (header + content + footer), got %d", len(children))
	}

	if len(children) >= 3 {
		// Check header
		headerClass := children[0].(Element).GetAttribute("class")
		if headerClass != "card-header" {
			t.Errorf("Expected card-header class, got '%s'", headerClass)
		}

		// Check content (in the middle, should be text node)
		contentText := children[1].TextContent()
		if contentText != "Card content" {
			t.Errorf("Expected 'Card content', got '%s'", contentText)
		}

		// Check footer (should be last)
		footerClass := children[2].(Element).GetAttribute("class")
		if footerClass != "card-footer" {
			t.Errorf("Expected card-footer class, got '%s'", footerClass)
		}
	}
}

func TestCard_WithColor(t *testing.T) {
	card := Card(WithColor(PRIMARY))

	class := card.Element().GetAttribute("class")
	if !strings.Contains(class, "card") {
		t.Errorf("Expected class to contain 'card', got '%s'", class)
	}
	if !strings.Contains(class, "text-bg-primary") {
		t.Errorf("Expected class to contain 'text-bg-primary', got '%s'", class)
	}
}

func TestCard_MultipleHeadingCalls(t *testing.T) {
	card := Card().
		Header("First heading").
		Header("Second heading")

	// Should have 2 children (header + body)
	children := card.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (header + body), got %d", len(children))
	}

	if len(children) >= 1 {
		headerText := children[0].TextContent()
		if headerText != "Second heading" {
			t.Errorf("Expected 'Second heading', got '%s'", headerText)
		}
	}
}

func TestCard_MultipleFooterCalls(t *testing.T) {
	card := Card().
		Footer("Last updated").
		Footer("3 mins ago")

	// Should have 2 children (body + footer)
	children := card.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children (body + footer), got %d", len(children))
	}

	if len(children) >= 2 {
		footerText := children[1].TextContent()
		if footerText != "3 mins ago" {
			t.Errorf("Expected '3 mins ago', got '%s'", footerText)
		}
	}
}

func TestCard_WithComponent(t *testing.T) {
	icon := Icon("star-fill")
	card := Card().
		Header("Featured ", icon).
		Append("Card with icon in header")

	// Check that icon is in header
	children := card.Element().ChildNodes()
	if len(children) < 1 {
		t.Fatal("Expected at least header child")
	}

	headerChildren := children[0].ChildNodes()
	if len(headerChildren) != 2 {
		t.Errorf("Expected 2 children in header (text + icon), got %d", len(headerChildren))
	}
}

func TestCard_WithClass(t *testing.T) {
	card := Card(WithClass("text-center", "mb-3"))

	class := card.Element().GetAttribute("class")
	expectedClasses := []string{"card", "text-center", "mb-3"}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected class to contain '%s', got '%s'", expected, class)
		}
	}
}

func TestCard_Component(t *testing.T) {
	card := Card()

	// Check component name
	if name := card.Name(); name != "card" {
		t.Errorf("Expected component name 'card', got '%s'", name)
	}

	// Check that Element() returns the same as root
	if elem := card.Element(); elem != card.root {
		t.Error("Element() should return the root element")
	}
}

func TestCard_MethodChaining(t *testing.T) {
	c := Card(WithClass("text-center")).
		Header("Card Title")

	c.Append("Card body content")
	c.Footer("Card footer")

	// Verify structure: header + content + footer
	children := c.Element().ChildNodes()
	if len(children) != 3 {
		t.Errorf("Expected 3 children, got %d", len(children))
	}

	// Verify classes
	class := c.Element().GetAttribute("class")
	if !strings.Contains(class, "card") || !strings.Contains(class, "text-center") {
		t.Errorf("Expected class to contain 'card' and 'text-center', got '%s'", class)
	}
}
