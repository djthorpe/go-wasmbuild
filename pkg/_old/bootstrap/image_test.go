package bootstrap

import (
	"strings"
	"testing"
)

func TestImage_Basic(t *testing.T) {
	img := Image("test.jpg")

	// Check element type
	if tag := img.Element().TagName(); tag != "IMG" {
		t.Errorf("Expected tag name 'IMG', got '%s'", tag)
	}

	// Check src attribute
	if src := img.Element().GetAttribute("src"); src != "test.jpg" {
		t.Errorf("Expected src 'test.jpg', got '%s'", src)
	}
}

func TestImage_WithAlt(t *testing.T) {
	img := Image("photo.jpg", WithAriaLabel("Profile photo"))

	// Check aria-label attribute
	if label := img.Element().GetAttribute("aria-label"); label != "Profile photo" {
		t.Errorf("Expected aria-label 'Profile photo', got '%s'", label)
	}
}

func TestImage_Fluid(t *testing.T) {
	img := Image("responsive.jpg", WithClass("img-fluid"))

	// Check img-fluid class
	if class := img.Element().GetAttribute("class"); !strings.Contains(class, "img-fluid") {
		t.Errorf("Expected class to contain 'img-fluid', got '%s'", class)
	}
}

func TestImage_Thumbnail(t *testing.T) {
	img := Image("thumb.jpg", WithClass("img-thumbnail"))

	// Check img-thumbnail class
	if class := img.Element().GetAttribute("class"); !strings.Contains(class, "img-thumbnail") {
		t.Errorf("Expected class to contain 'img-thumbnail', got '%s'", class)
	}
}

func TestImage_Rounded(t *testing.T) {
	img := Image("avatar.jpg", WithClass("rounded-circle"))

	// Check rounded-circle class
	if class := img.Element().GetAttribute("class"); !strings.Contains(class, "rounded-circle") {
		t.Errorf("Expected class to contain 'rounded-circle', got '%s'", class)
	}
}

func TestImage_WithMultipleClasses(t *testing.T) {
	img := Image("banner.jpg", WithClass("img-fluid", "rounded", "shadow"))

	class := img.Element().GetAttribute("class")
	expectedClasses := []string{"img-fluid", "rounded", "shadow"}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected class to contain '%s', got '%s'", expected, class)
		}
	}
}

func TestImage_Component(t *testing.T) {
	img := Image("test.jpg")

	// Check component name
	if name := img.Name(); name != "image" {
		t.Errorf("Expected component name 'image', got '%s'", name)
	}

	// Check that Element() returns the same as root
	if elem := img.Element(); elem != img.root {
		t.Error("Element() should return the root element")
	}
}

func TestImage_EmptySrc(t *testing.T) {
	img := Image("")

	// Check that src is empty
	if src := img.Element().GetAttribute("src"); src != "" {
		t.Errorf("Expected empty src, got '%s'", src)
	}
}

func TestImage_WithWidth(t *testing.T) {
	img := Image("sized.jpg", WithClass("w-50"))

	// Check width class
	if class := img.Element().GetAttribute("class"); !strings.Contains(class, "w-50") {
		t.Errorf("Expected class to contain 'w-50', got '%s'", class)
	}
}

func TestImage_Float(t *testing.T) {
	img := Image("float.jpg", WithClass("float-start"))

	// Check float class
	if class := img.Element().GetAttribute("class"); !strings.Contains(class, "float-start") {
		t.Errorf("Expected class to contain 'float-start', got '%s'", class)
	}
}
