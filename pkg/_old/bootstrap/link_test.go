package bootstrap

import (
	"strings"
	"testing"
)

func TestLink_Basic(t *testing.T) {
	link := Link("/test")

	// Check element type
	if tag := link.Element().TagName(); tag != "A" {
		t.Errorf("Expected tag name 'A', got '%s'", tag)
	}

	// Check href attribute
	if href := link.Element().GetAttribute("href"); href != "/test" {
		t.Errorf("Expected href '/test', got '%s'", href)
	}
}

func TestLink_WithContent(t *testing.T) {
	link := Link("#section").Append("Jump to Section")

	// Check text content
	if text := link.Element().TextContent(); text != "Jump to Section" {
		t.Errorf("Expected text 'Jump to Section', got '%s'", text)
	}
}

func TestLink_WithColor(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{PRIMARY, "link-primary"},
		{SECONDARY, "link-secondary"},
		{SUCCESS, "link-success"},
		{DANGER, "link-danger"},
		{WARNING, "link-warning"},
		{INFO, "link-info"},
		{LIGHT, "link-light"},
		{DARK, "link-dark"},
	}

	for _, tt := range tests {
		t.Run(string(tt.color), func(t *testing.T) {
			link := Link("#", WithColor(tt.color))
			class := link.Element().GetAttribute("class")
			if !strings.Contains(class, tt.expected) {
				t.Errorf("Expected class to contain '%s', got '%s'", tt.expected, class)
			}
		})
	}
}

func TestLink_WithClass(t *testing.T) {
	link := Link("#", WithClass("link-opacity-50", "link-offset-2"))
	class := link.Element().GetAttribute("class")

	if !strings.Contains(class, "link-opacity-50") {
		t.Errorf("Expected class to contain 'link-opacity-50', got '%s'", class)
	}
	if !strings.Contains(class, "link-offset-2") {
		t.Errorf("Expected class to contain 'link-offset-2', got '%s'", class)
	}
}

func TestLink_WithUnderlineClasses(t *testing.T) {
	link := Link("#", WithClass("link-underline", "link-underline-primary", "link-underline-opacity-50"))
	class := link.Element().GetAttribute("class")

	expectedClasses := []string{
		"link-underline",
		"link-underline-primary",
		"link-underline-opacity-50",
	}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected class to contain '%s', got '%s'", expected, class)
		}
	}
}

func TestLink_SetAttributeAfterCreation(t *testing.T) {
	link := Link("https://example.com")
	link.Element().SetAttribute("target", "_blank")
	link.Element().SetAttribute("rel", "noopener noreferrer")

	if target := link.Element().GetAttribute("target"); target != "_blank" {
		t.Errorf("Expected target '_blank', got '%s'", target)
	}

	if rel := link.Element().GetAttribute("rel"); rel != "noopener noreferrer" {
		t.Errorf("Expected rel 'noopener noreferrer', got '%s'", rel)
	}
}

func TestLink_MultipleOptions(t *testing.T) {
	link := Link("#", WithColor(PRIMARY), WithClass("link-offset-2", "link-underline", "link-underline-opacity-25"))
	class := link.Element().GetAttribute("class")

	expectedClasses := []string{
		"link-primary",
		"link-offset-2",
		"link-underline",
		"link-underline-opacity-25",
	}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected class to contain '%s', got '%s'", expected, class)
		}
	}
}

func TestLink_Component(t *testing.T) {
	link := Link("#")

	// Check component name
	if name := link.Name(); name != "link" {
		t.Errorf("Expected component name 'link', got '%s'", name)
	}

	// Check that Element() returns the same as root
	if elem := link.Element(); elem != link.root {
		t.Error("Element() should return the root element")
	}
}

func TestLink_WithIcon(t *testing.T) {
	// Test that links can contain icons
	icon := Icon("arrow-right")
	link := Link("#").Append("Next Page ", icon)

	// Check that link contains children
	children := link.Element().ChildNodes()
	if len(children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(children))
	}
}

func TestLink_EmptyHref(t *testing.T) {
	// Test link with empty href
	link := Link("")
	href := link.Element().GetAttribute("href")
	if href != "" {
		t.Errorf("Expected empty href, got '%s'", href)
	}
}

func TestLink_WithAriaLabel(t *testing.T) {
	link := Link("#", WithAriaLabel("Go to home"))
	ariaLabel := link.Element().GetAttribute("aria-label")
	if ariaLabel != "Go to home" {
		t.Errorf("Expected aria-label 'Go to home', got '%s'", ariaLabel)
	}
}
