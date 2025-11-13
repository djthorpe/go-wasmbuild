package bs_test

import (
	"strings"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// BASIC FIGURE TESTS

func TestFigureCreation(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{
			name: "create without options",
			opts: nil,
		},
		{
			name: "create with ID",
			opts: []Opt{WithID("figure-1")},
		},
		{
			name: "create with additional classes",
			opts: []Opt{WithClass("shadow", "rounded")},
		},
		{
			name: "create with combined options",
			opts: []Opt{
				WithID("figure-main"),
				WithClass("my-figure"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := bs.Figure(tt.opts...)
			if f == nil {
				t.Fatal("bs.Figure() returned nil")
			}
			if f.Name() != bs.ViewFigure {
				t.Errorf("bs.Figure().Name() = %v, want %v", f.Name(), bs.ViewFigure)
			}
		})
	}
}

func TestFigureViewInterface(t *testing.T) {
	f := bs.Figure()

	t.Run("Name returns correct view name", func(t *testing.T) {
		if f.Name() != bs.ViewFigure {
			t.Errorf("Figure.Name() = %v, want %v", f.Name(), bs.ViewFigure)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := f.Root()
		if root == nil {
			t.Error("Figure.Root() returned nil")
		}
	})

	t.Run("Root returns FIGURE element", func(t *testing.T) {
		tagName := f.Root().TagName()
		if tagName != "FIGURE" {
			t.Errorf("Figure root element tag = %v, want FIGURE", tagName)
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Figure.ID() panicked: %v", r)
			}
		}()
		_ = f.ID()
	})
}

func TestFigureDefaultClass(t *testing.T) {
	f := bs.Figure()
	classList := f.Root().ClassList()

	if !classList.Contains("figure") {
		t.Error("Figure should have 'figure' class by default")
	}
}

func TestFigureEmbedding(t *testing.T) {
	// Verify that figure properly embeds View
	f := bs.Figure()

	// figure should be able to be used as a View
	var v View = f
	if v.Name() != bs.ViewFigure {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), bs.ViewFigure)
	}
}

func TestFigureMultipleInstances(t *testing.T) {
	// Create multiple figure elements and verify they're independent
	f1 := bs.Figure(WithID("figure-1"))
	f2 := bs.Figure(WithID("figure-2"))

	if f1 == f2 {
		t.Error("Each bs.Figure() call should create a distinct instance")
	}

	if f1.ID() == f2.ID() {
		t.Error("Figures with different IDs should maintain separate state")
	}

	if f1.Root() == f2.Root() {
		t.Error("Figures should have separate DOM elements")
	}
}

func TestFigureWithID(t *testing.T) {
	f := bs.Figure(WithID("test-figure"))

	if f.ID() != "test-figure" {
		t.Errorf("Figure.ID() = %v, want 'test-figure'", f.ID())
	}
}

func TestFigureWithAdditionalClasses(t *testing.T) {
	f := bs.Figure(WithClass("custom-class", "another-class"))
	classList := f.Root().ClassList()

	if !classList.Contains("figure") {
		t.Error("Figure should have default 'figure' class")
	}

	if !classList.Contains("custom-class") {
		t.Error("Figure should have 'custom-class'")
	}

	if !classList.Contains("another-class") {
		t.Error("Figure should have 'another-class'")
	}
}

///////////////////////////////////////////////////////////////////////////////
// FIGURE WITH IMAGE TESTS

func TestFigureWithSingleImage(t *testing.T) {
	img := bs.Image("placeholder.jpg")
	f := bs.Figure()
	f.Append(img)

	classList := img.Root().ClassList()

	if !classList.Contains("figure-img") {
		t.Error("Image in figure should have 'figure-img' class")
	}

	if !classList.Contains("image-fluid") {
		t.Error("Image should retain its 'image-fluid' class")
	}
}

func TestFigureWithRoundedImage(t *testing.T) {
	img := bs.RoundedImage("placeholder.jpg")
	f := bs.Figure()
	f.Append(img)

	classList := img.Root().ClassList()

	if !classList.Contains("figure-img") {
		t.Error("Rounded image in figure should have 'figure-img' class")
	}

	if !classList.Contains("rounded") {
		t.Error("Rounded image should retain 'rounded' class")
	}
}

func TestFigureWithImageAndFluidClass(t *testing.T) {
	img := bs.Image("test.jpg", WithClass("img-fluid"))
	f := bs.Figure()
	f.Append(img)

	classList := img.Root().ClassList()

	if !classList.Contains("figure-img") {
		t.Error("Image should have 'figure-img' class")
	}

	if !classList.Contains("img-fluid") {
		t.Error("Image should have 'img-fluid' class")
	}
}

///////////////////////////////////////////////////////////////////////////////
// FIGURE CAPTION TESTS

func TestFigureWithCaption(t *testing.T) {
	f := bs.Figure()
	img := bs.Image("placeholder.jpg")
	f.Append(img)
	f.Caption("A caption for the image")

	// Check that the figure has a figcaption element
	lastChild := f.Root().LastElementChild()
	if lastChild == nil {
		t.Fatal("Figure should have a last child element")
	}

	if lastChild.TagName() != "FIGCAPTION" {
		t.Errorf("Last child should be FIGCAPTION, got %v", lastChild.TagName())
	}

	classList := lastChild.ClassList()
	if !classList.Contains("figure-caption") {
		t.Error("Caption should have 'figure-caption' class")
	}
}

func TestFigureCaptionContent(t *testing.T) {
	f := bs.Figure()
	img := bs.Image("test.jpg")
	f.Append(img)
	f.Caption("Test caption text")

	caption := f.Root().LastElementChild()
	if caption == nil {
		t.Fatal("Figure should have caption element")
	}

	textContent := caption.TextContent()
	if textContent != "Test caption text" {
		t.Errorf("Caption text = %q, want %q", textContent, "Test caption text")
	}
}

func TestFigureCaptionReplacement(t *testing.T) {
	f := bs.Figure()
	img := bs.Image("test.jpg")
	f.Append(img)

	// Add first caption
	f.Caption("First caption")
	firstCaptionText := f.Root().LastElementChild().TextContent()

	if firstCaptionText != "First caption" {
		t.Errorf("First caption text = %q, want %q", firstCaptionText, "First caption")
	}

	// Replace with second caption
	f.Caption("Second caption")
	secondCaptionText := f.Root().LastElementChild().TextContent()

	if secondCaptionText != "Second caption" {
		t.Errorf("Second caption text = %q, want %q", secondCaptionText, "Second caption")
	}

	// Verify only one caption exists
	children := f.Root().Children()
	captionCount := 0
	for _, child := range children {
		if child != nil && child.TagName() == "FIGCAPTION" {
			captionCount++
		}
	}

	if captionCount != 1 {
		t.Errorf("Figure should have exactly 1 caption, found %d", captionCount)
	}
}

func TestFigureCaptionWithTextAlignment(t *testing.T) {
	tests := []struct {
		name          string
		alignmentOpt  Opt
		expectedClass string
	}{
		{
			name:          "caption with text-start",
			alignmentOpt:  WithClass("text-start"),
			expectedClass: "text-start",
		},
		{
			name:          "caption with text-center",
			alignmentOpt:  WithClass("text-center"),
			expectedClass: "text-center",
		},
		{
			name:          "caption with text-end",
			alignmentOpt:  WithClass("text-end"),
			expectedClass: "text-end",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := bs.Figure(tt.alignmentOpt)
			img := bs.Image("test.jpg")
			f.Append(img)
			f.Caption("Aligned caption")

			caption := f.Root().LastElementChild()
			if caption == nil {
				t.Fatal("Caption should exist")
			}

			// Check if parent figure has the alignment class
			figureClassList := f.Root().ClassList()
			if !figureClassList.Contains(tt.expectedClass) {
				t.Errorf("Figure should have alignment class %q", tt.expectedClass)
			}
		})
	}
}

func TestFigureCaptionAtEnd(t *testing.T) {
	f := bs.Figure()
	img := bs.Image("test.jpg")
	f.Append(img)
	f.Caption("Caption text")

	// Caption should be the last child
	lastChild := f.Root().LastElementChild()
	if lastChild == nil || lastChild.TagName() != "FIGCAPTION" {
		t.Error("Caption should be the last child of figure")
	}

	// Image should come before caption
	firstChild := f.Root().FirstElementChild()
	if firstChild == nil || firstChild.TagName() != "IMG" {
		t.Error("Image should be the first child of figure")
	}
}

///////////////////////////////////////////////////////////////////////////////
// FIGURE INSERT TESTS

func TestFigureInsertImage(t *testing.T) {
	f := bs.Figure()
	img := bs.Image("inserted.jpg")
	f.Content(img)

	classList := img.Root().ClassList()
	if !classList.Contains("figure-img") {
		t.Error("Inserted image should have 'figure-img' class")
	}
}

func TestFigureInsertWithCaption(t *testing.T) {
	f := bs.Figure()
	f.Caption("Caption first")

	img := bs.Image("test.jpg")
	f.Content(img)

	// Caption should still be last
	lastChild := f.Root().LastElementChild()
	if lastChild == nil || lastChild.TagName() != "FIGCAPTION" {
		t.Error("Caption should remain the last child after Content()")
	}

	// Caption text should be preserved
	captionText := lastChild.TextContent()
	if captionText != "Caption first" {
		t.Errorf("Caption text should be preserved, got %q", captionText)
	}

	// Image should be first
	firstChild := f.Root().FirstElementChild()
	if firstChild == nil || firstChild.TagName() != "IMG" {
		t.Error("Image should be first child")
	}
}

func TestFigureInsertMultipleImagesWithCaption(t *testing.T) {
	f := bs.Figure()
	img1 := bs.Image("first.jpg")
	f.Content(img1)
	f.Caption("Test caption")

	// Insert another image
	img2 := bs.Image("second.jpg")
	f.Content(img2)

	// Caption should still be last
	lastChild := f.Root().LastElementChild()
	if lastChild == nil || lastChild.TagName() != "FIGCAPTION" {
		t.Error("Caption should remain last after inserting second image")
	}

	// Caption text should be preserved
	captionText := lastChild.TextContent()
	if captionText != "Test caption" {
		t.Errorf("Caption text = %q, want %q", captionText, "Test caption")
	}

	// Should have 3 children: 2 images + caption
	children := f.Root().Children()
	if len(children) != 3 {
		t.Errorf("Figure should have 3 children, got %d", len(children))
	}
}

///////////////////////////////////////////////////////////////////////////////
// MULTIPLE IMAGES TESTS

func TestFigureWithMultipleImages(t *testing.T) {
	f := bs.Figure()
	img1 := bs.Image("first.jpg")
	img2 := bs.Image("second.jpg")

	f.Append(img1)
	f.Append(img2)

	// Both images should have figure-img class
	if !img1.Root().ClassList().Contains("figure-img") {
		t.Error("First image should have 'figure-img' class")
	}

	if !img2.Root().ClassList().Contains("figure-img") {
		t.Error("Second image should have 'figure-img' class")
	}
}

func TestFigureMultipleImagesWithCaption(t *testing.T) {
	f := bs.Figure()
	img1 := bs.Image("first.jpg")
	img2 := bs.Image("second.jpg")

	f.Append(img1)
	f.Append(img2)
	f.Caption("Caption for both images")

	// Caption should be last
	lastChild := f.Root().LastElementChild()
	if lastChild == nil || lastChild.TagName() != "FIGCAPTION" {
		t.Error("Caption should be last child")
	}

	// Should have 3 children total (2 images + caption)
	children := f.Root().Children()
	if len(children) != 3 {
		t.Errorf("Figure should have 3 children, got %d", len(children))
	}
}

///////////////////////////////////////////////////////////////////////////////
// FIGURE CAPTION PERSISTENCE TESTS

func TestFigureCaptionPersistsAcrossAppends(t *testing.T) {
	f := bs.Figure()
	img1 := bs.Image("first.jpg")
	f.Append(img1)
	f.Caption("Initial caption")

	// Add another image
	img2 := bs.Image("second.jpg")
	f.Append(img2)

	// Caption should still exist and be last
	lastChild := f.Root().LastElementChild()
	if lastChild == nil || lastChild.TagName() != "FIGCAPTION" {
		t.Error("Caption should persist and remain last after appending more images")
	}

	captionText := lastChild.TextContent()
	if captionText != "Initial caption" {
		t.Errorf("Caption text should be preserved, got %q", captionText)
	}
}

///////////////////////////////////////////////////////////////////////////////
// OUTER HTML TESTS

func TestFigureOuterHTML(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() View
		contains []string
	}{
		{
			name: "basic figure",
			setup: func() View {
				return bs.Figure()
			},
			contains: []string{
				"<figure",
				"class=\"figure\"",
				"</figure>",
			},
		},
		{
			name: "figure with ID",
			setup: func() View {
				return bs.Figure(WithID("my-figure"))
			},
			contains: []string{
				"<figure",
				"id=\"my-figure\"",
				"class=\"figure\"",
			},
		},
		{
			name: "figure with image",
			setup: func() View {
				f := bs.Figure()
				img := bs.Image("test.jpg")
				f.Append(img)
				return f
			},
			contains: []string{
				"<figure",
				"<img",
				"class=\"image-fluid figure-img\"",
				"src=\"test.jpg\"",
			},
		},
		{
			name: "figure with image and caption",
			setup: func() View {
				f := bs.Figure()
				img := bs.Image("test.jpg")
				f.Append(img)
				f.Caption("Test caption")
				return f
			},
			contains: []string{
				"<figure",
				"<img",
				"<figcaption",
				"class=\"figure-caption\"",
				"Test caption",
				"</figcaption>",
			},
		},
		{
			name: "figure with rounded image",
			setup: func() View {
				f := bs.Figure()
				img := bs.RoundedImage("test.jpg")
				f.Append(img)
				f.Caption("Rounded image caption")
				return f
			},
			contains: []string{
				"rounded",
				"image-fluid",
				"figure-img",
				"<figcaption",
				"Rounded image caption",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.setup()
			html := f.Root().OuterHTML()

			for _, expected := range tt.contains {
				if !strings.Contains(html, expected) {
					t.Errorf("OuterHTML should contain %q, got: %s", expected, html)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// EDGE CASES

func TestFigureEdgeCases(t *testing.T) {
	t.Run("caption without image", func(t *testing.T) {
		f := bs.Figure()
		f.Caption("Caption without image")

		caption := f.Root().LastElementChild()
		if caption == nil || caption.TagName() != "FIGCAPTION" {
			t.Error("Should be able to add caption without image")
		}
	})

	t.Run("empty caption", func(t *testing.T) {
		f := bs.Figure()
		img := bs.Image("test.jpg")
		f.Append(img)
		f.Caption("")

		caption := f.Root().LastElementChild()
		if caption == nil || caption.TagName() != "FIGCAPTION" {
			t.Error("Should be able to add empty caption")
		}

		if caption.TextContent() != "" {
			t.Errorf("Empty caption should have no text content, got %q", caption.TextContent())
		}
	})

	t.Run("caption with multiple arguments", func(t *testing.T) {
		f := bs.Figure()
		img := bs.Image("test.jpg")
		f.Append(img)
		f.Caption("Part 1", " Part 2", " Part 3")

		caption := f.Root().LastElementChild()
		if caption == nil {
			t.Fatal("Caption should exist")
		}

		text := caption.TextContent()
		expected := "Part 1 Part 2 Part 3"
		if text != expected {
			t.Errorf("Caption text = %q, want %q", text, expected)
		}
	})

	t.Run("append multiple children should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Appending multiple children should panic")
			}
		}()

		f := bs.Figure()
		img1 := bs.Image("first.jpg")
		img2 := bs.Image("second.jpg")
		f.Append(img1, img2) // Should panic
	})

	t.Run("insert multiple children should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Inserting multiple children should panic")
			}
		}()

		f := bs.Figure()
		img1 := bs.Image("first.jpg")
		img2 := bs.Image("second.jpg")
		f.Content(img1, img2) // Should panic
	})

	t.Run("append non-image should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Appending non-image should panic")
			}
		}()

		f := bs.Figure()
		badge := bs.Badge()
		f.Append(badge) // Should panic
	})

	t.Run("insert non-image should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Inserting non-image should panic")
			}
		}()

		f := bs.Figure()
		badge := bs.Badge()
		f.Content(badge) // Should panic
	})
}

///////////////////////////////////////////////////////////////////////////////
// ATTRIBUTES TESTS

func TestFigureWithAttributes(t *testing.T) {
	f := bs.Figure(
		WithID("main-figure"),
		WithAttr("data-category", "nature"),
		WithAttr("aria-label", "Nature photograph"),
	)

	root := f.Root()

	if f.ID() != "main-figure" {
		t.Errorf("ID = %q, want 'main-figure'", f.ID())
	}

	if root.GetAttribute("data-category") != "nature" {
		t.Error("Should have data-category attribute")
	}

	if root.GetAttribute("aria-label") != "Nature photograph" {
		t.Error("Should have aria-label attribute")
	}
}
