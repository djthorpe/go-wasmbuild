package bs

import (
	"fmt"
	"testing"

	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TestViewHeadingConstant(t *testing.T) {
	expected := "mvc-bs-heading"
	if ViewHeading != expected {
		t.Errorf("ViewHeading = %v, want %v", ViewHeading, expected)
	}
}

func TestHeadingRegistered(t *testing.T) {
	// Check if the heading view is registered during init
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Heading view should be registered, but got panic: %v", r)
		}
	}()

	// This will panic if ViewHeading is not registered
	_ = Heading(1)
}

func TestHeadingType(t *testing.T) {
	// Verify the heading type exists and has correct structure
	var _ *heading = &heading{}

	// Verify heading embeds View
	h := &heading{}
	if _, ok := any(h).(interface{ View }); !ok {
		t.Error("heading should embed View")
	}
}

func TestHeadingCreation(t *testing.T) {
	tests := []struct {
		name  string
		level int
		opts  []Opt
	}{
		{
			name:  "create H1 without options",
			level: 1,
			opts:  nil,
		},
		{
			name:  "create H2 with ID",
			level: 2,
			opts:  []Opt{WithID("main-title")},
		},
		{
			name:  "create H3 with classes",
			level: 3,
			opts:  []Opt{WithClass("text-center", "my-3")},
		},
		{
			name:  "create H4 with multiple options",
			level: 4,
			opts: []Opt{
				WithID("section-heading"),
				WithClass("display-4"),
				WithAttr("data-level", "4"),
			},
		},
		{
			name:  "create H5",
			level: 5,
			opts:  nil,
		},
		{
			name:  "create H6",
			level: 6,
			opts:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Heading() with valid level and options should not panic: %v", r)
				}
			}()

			h := Heading(tt.level, tt.opts...)
			if h == nil {
				t.Fatal("Heading() returned nil")
			}

			if h.View == nil {
				t.Error("Heading().View should not be nil")
			}

			if h.Name() != ViewHeading {
				t.Errorf("Heading().Name() = %v, want %v", h.Name(), ViewHeading)
			}
		})
	}
}

func TestHeadingLevels(t *testing.T) {
	tests := []struct {
		level       int
		expectPanic bool
		description string
	}{
		{1, false, "H1 should be valid"},
		{2, false, "H2 should be valid"},
		{3, false, "H3 should be valid"},
		{4, false, "H4 should be valid"},
		{5, false, "H5 should be valid"},
		{6, false, "H6 should be valid"},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.expectPanic {
					t.Errorf("Heading(%d) panic = %v, wantPanic %v", tt.level, r != nil, tt.expectPanic)
				}
			}()

			h := Heading(tt.level)
			if !tt.expectPanic && h == nil {
				t.Errorf("Heading(%d) returned nil for valid level", tt.level)
			}
		})
	}
}

func TestHeadingReturnsPointer(t *testing.T) {
	// Verify Heading returns a pointer to heading
	h := Heading(1)
	if h == nil {
		t.Fatal("Heading() should not return nil")
	}

	// Test that we can call it multiple times and get different instances
	h1 := Heading(1)
	h2 := Heading(1)

	if h1 == h2 {
		t.Error("Heading() should return different instances on each call")
	}
}

func TestHeadingWithOptions(t *testing.T) {
	t.Run("with ID", func(t *testing.T) {
		h := Heading(2, WithID("my-heading"))
		if h.ID() != "my-heading" {
			t.Errorf("Heading(2, WithID('my-heading')).ID() = %v, want %v", h.ID(), "my-heading")
		}
	})

	t.Run("with multiple classes", func(t *testing.T) {
		h := Heading(3, WithClass("display-1", "text-primary"))
		// We can't easily test the classes without DOM, but verify it doesn't panic
		if h == nil {
			t.Error("Heading() with classes should not return nil")
		}
	})

	t.Run("with attributes", func(t *testing.T) {
		h := Heading(4, WithAttr("role", "heading"))
		// Verify it doesn't panic with attributes
		if h == nil {
			t.Error("Heading() with attributes should not return nil")
		}
	})

	t.Run("with combined options", func(t *testing.T) {
		h := Heading(1,
			WithID("page-title"),
			WithClass("display-2", "fw-bold"),
			WithAttr("aria-level", "1"),
		)

		if h == nil {
			t.Fatal("Heading() with combined options should not return nil")
		}

		if h.ID() != "page-title" {
			t.Errorf("Heading().ID() = %v, want %v", h.ID(), "page-title")
		}
	})
}

func TestHeadingViewInterface(t *testing.T) {
	// Verify heading implements or embeds View properly
	h := Heading(2)

	// Test that View methods are accessible
	t.Run("Name method", func(t *testing.T) {
		if name := h.Name(); name != ViewHeading {
			t.Errorf("Name() = %v, want %v", name, ViewHeading)
		}
	})

	t.Run("Root method", func(t *testing.T) {
		root := h.Root()
		if root == nil {
			t.Error("Root() should not return nil")
		}
	})

	t.Run("ID method", func(t *testing.T) {
		// Should not panic even with no ID set
		_ = h.ID()
	})
}

func TestHeadingChaining(t *testing.T) {
	// Test that View methods can be chained through the embedded View
	t.Run("chaining pattern compiles", func(t *testing.T) {
		_ = func(h *heading) View { return h.Opts(WithClass("test")) }

		// These would work in a real DOM environment:
		// h := Heading(1)
		// h.Append("Title").Apply(WithClass("highlight"))
	})
}

func TestNewHeadingFromElement(t *testing.T) {
	// Test the constructor function used for element recovery
	t.Run("function exists", func(t *testing.T) {
		// Verify the function signature
		var _ ViewConstructorFunc = newHeadingFromElement
	})

	// Note: Testing with actual elements requires a DOM environment
	// Those tests would go in integration tests or WASM-specific tests
}

func TestHeadingEmbedding(t *testing.T) {
	// Verify that heading properly embeds View
	h := Heading(3)

	// heading should be able to be used as a View
	var v View = h
	if v.Name() != ViewHeading {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), ViewHeading)
	}
}

func TestHeadingMultipleInstances(t *testing.T) {
	// Create multiple headings and verify they're independent
	h1 := Heading(1, WithID("heading-1"))
	h2 := Heading(2, WithID("heading-2"))
	h3 := Heading(3)

	if h1 == h2 || h1 == h3 || h2 == h3 {
		t.Error("Each Heading() call should create a distinct instance")
	}

	if h1.ID() == h2.ID() {
		t.Error("Headings with different IDs should maintain separate state")
	}
}

func TestHeadingDifferentLevels(t *testing.T) {
	// Create headings at different levels and verify they work independently
	headings := make([]*heading, 6)
	for i := 0; i < 6; i++ {
		headings[i] = Heading(i+1, WithID(fmt.Sprintf("h%d", i+1)))
	}

	for i, h := range headings {
		if h == nil {
			t.Errorf("Heading(%d) returned nil", i+1)
		}
		expectedID := fmt.Sprintf("h%d", i+1)
		if h.ID() != expectedID {
			t.Errorf("Heading(%d).ID() = %v, want %v", i+1, h.ID(), expectedID)
		}
	}
}

func TestHeadingSameLevel(t *testing.T) {
	// Create multiple headings at the same level with different content
	h1 := Heading(2, WithID("section-1"))
	h2 := Heading(2, WithID("section-2"))
	h3 := Heading(2, WithID("section-3"))

	ids := []string{h1.ID(), h2.ID(), h3.ID()}
	expected := []string{"section-1", "section-2", "section-3"}

	for i, id := range ids {
		if id != expected[i] {
			t.Errorf("Heading %d: ID = %v, want %v", i+1, id, expected[i])
		}
	}
}
