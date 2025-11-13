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
// BASIC LINK TESTS

func TestLinkCreation(t *testing.T) {
	tests := []struct {
		name string
		href string
		opts []Opt
	}{
		{
			name: "create without options",
			href: "#",
		},
		{
			name: "create with absolute URL",
			href: "https://example.com",
		},
		{
			name: "create with relative URL",
			href: "/page",
		},
		{
			name: "create with ID",
			href: "#section",
			opts: []Opt{WithID("main-link")},
		},
		{
			name: "create with classes",
			href: "#",
			opts: []Opt{WithClass("fw-bold", "text-decoration-none")},
		},
		{
			name: "create with combined options",
			href: "https://example.com",
			opts: []Opt{
				WithID("external-link"),
				WithClass("external"),
				WithAttr("target", "_blank"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link(tt.href, tt.opts...)
			if link == nil {
				t.Fatal("bs.Link() returned nil")
			}
			if link.Name() != bs.ViewLink {
				t.Errorf("bs.Link().Name() = %v, want %v", link.Name(), bs.ViewLink)
			}
		})
	}
}

func TestLinkViewInterface(t *testing.T) {
	link := bs.Link("#")

	t.Run("Name returns correct view name", func(t *testing.T) {
		if link.Name() != bs.ViewLink {
			t.Errorf("Link.Name() = %v, want %v", link.Name(), bs.ViewLink)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := link.Root()
		if root == nil {
			t.Error("Link.Root() returned nil")
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Link.ID() panicked: %v", r)
			}
		}()
		_ = link.ID()
	})

	t.Run("Root has correct tag name", func(t *testing.T) {
		root := link.Root()
		if root.TagName() != "A" {
			t.Errorf("Link.Root().TagName() = %v, want A", root.TagName())
		}
	})
}

func TestLinkHrefAttribute(t *testing.T) {
	tests := []struct {
		name string
		href string
	}{
		{
			name: "anchor link",
			href: "#",
		},
		{
			name: "section anchor",
			href: "#section",
		},
		{
			name: "relative path",
			href: "/about",
		},
		{
			name: "absolute URL",
			href: "https://example.com",
		},
		{
			name: "mailto link",
			href: "mailto:test@example.com",
		},
		{
			name: "tel link",
			href: "tel:+1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link(tt.href)
			href := link.Root().GetAttribute("href")
			if href != tt.href {
				t.Errorf("Link href = %q, want %q", href, tt.href)
			}
		})
	}
}

func TestLinkEmbedding(t *testing.T) {
	// Verify that link properly embeds View
	link := bs.Link("#")

	t.Run("can append text", func(t *testing.T) {
		link.Append("Click here")
		// Should not panic
	})

	t.Run("can use Opts method", func(t *testing.T) {
		link.Opts(WithClass("test"))
		if !link.Root().ClassList().Contains("test") {
			t.Error("Link should have 'test' class after Opts")
		}
	})
}

func TestLinkMultipleInstances(t *testing.T) {
	link1 := bs.Link("#link1", WithID("link1"))
	link2 := bs.Link("#link2", WithID("link2"))

	t.Run("links are independent", func(t *testing.T) {
		if link1.ID() == link2.ID() {
			t.Error("Different link instances should have different IDs")
		}
	})

	t.Run("each link has correct href", func(t *testing.T) {
		if href := link1.Root().GetAttribute("href"); href != "#link1" {
			t.Errorf("link1 href = %q, want #link1", href)
		}
		if href := link2.Root().GetAttribute("href"); href != "#link2" {
			t.Errorf("link2 href = %q, want #link2", href)
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// LINK OPACITY TESTS

func TestLinkOpacity(t *testing.T) {
	tests := []struct {
		name      string
		className string
		opacity   string
	}{
		{
			name:      "opacity 10",
			className: "link-opacity-10",
			opacity:   "10",
		},
		{
			name:      "opacity 25",
			className: "link-opacity-25",
			opacity:   "25",
		},
		{
			name:      "opacity 50",
			className: "link-opacity-50",
			opacity:   "50",
		},
		{
			name:      "opacity 75",
			className: "link-opacity-75",
			opacity:   "75",
		},
		{
			name:      "opacity 100",
			className: "link-opacity-100",
			opacity:   "100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkOpacityHover(t *testing.T) {
	tests := []struct {
		name      string
		className string
	}{
		{
			name:      "hover opacity 10",
			className: "link-opacity-10-hover",
		},
		{
			name:      "hover opacity 25",
			className: "link-opacity-25-hover",
		},
		{
			name:      "hover opacity 50",
			className: "link-opacity-50-hover",
		},
		{
			name:      "hover opacity 75",
			className: "link-opacity-75-hover",
		},
		{
			name:      "hover opacity 100",
			className: "link-opacity-100-hover",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkOpacityCombinations(t *testing.T) {
	link := bs.Link("#", WithClass("link-opacity-25", "link-opacity-100-hover"))
	classList := link.Root().ClassList()

	t.Run("has both opacity classes", func(t *testing.T) {
		if !classList.Contains("link-opacity-25") {
			t.Error("Link should have 'link-opacity-25' class")
		}
		if !classList.Contains("link-opacity-100-hover") {
			t.Error("Link should have 'link-opacity-100-hover' class")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// LINK UNDERLINE COLOR TESTS

func TestLinkUnderlineColor(t *testing.T) {
	tests := []struct {
		name      string
		className string
		color     string
	}{
		{
			name:      "primary underline",
			className: "link-underline-primary",
			color:     "primary",
		},
		{
			name:      "secondary underline",
			className: "link-underline-secondary",
			color:     "secondary",
		},
		{
			name:      "success underline",
			className: "link-underline-success",
			color:     "success",
		},
		{
			name:      "danger underline",
			className: "link-underline-danger",
			color:     "danger",
		},
		{
			name:      "warning underline",
			className: "link-underline-warning",
			color:     "warning",
		},
		{
			name:      "info underline",
			className: "link-underline-info",
			color:     "info",
		},
		{
			name:      "light underline",
			className: "link-underline-light",
			color:     "light",
		},
		{
			name:      "dark underline",
			className: "link-underline-dark",
			color:     "dark",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// LINK UNDERLINE OFFSET TESTS

func TestLinkUnderlineOffset(t *testing.T) {
	tests := []struct {
		name      string
		className string
		offset    string
	}{
		{
			name:      "offset 1",
			className: "link-offset-1",
			offset:    "1",
		},
		{
			name:      "offset 2",
			className: "link-offset-2",
			offset:    "2",
		},
		{
			name:      "offset 3",
			className: "link-offset-3",
			offset:    "3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkUnderlineOffsetHover(t *testing.T) {
	tests := []struct {
		name      string
		className string
	}{
		{
			name:      "hover offset 1",
			className: "link-offset-1-hover",
		},
		{
			name:      "hover offset 2",
			className: "link-offset-2-hover",
		},
		{
			name:      "hover offset 3",
			className: "link-offset-3-hover",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// LINK UNDERLINE OPACITY TESTS

func TestLinkUnderlineOpacity(t *testing.T) {
	tests := []struct {
		name      string
		className string
	}{
		{
			name:      "underline opacity 0",
			className: "link-underline-opacity-0",
		},
		{
			name:      "underline opacity 10",
			className: "link-underline-opacity-10",
		},
		{
			name:      "underline opacity 25",
			className: "link-underline-opacity-25",
		},
		{
			name:      "underline opacity 50",
			className: "link-underline-opacity-50",
		},
		{
			name:      "underline opacity 75",
			className: "link-underline-opacity-75",
		},
		{
			name:      "underline opacity 100",
			className: "link-underline-opacity-100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// link-underline class is required for underline opacity to work
			link := bs.Link("#", WithClass("link-underline", tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains("link-underline") {
				t.Error("Link should have 'link-underline' class")
			}
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkUnderlineOpacityHover(t *testing.T) {
	tests := []struct {
		name      string
		className string
	}{
		{
			name:      "hover underline opacity 0",
			className: "link-underline-opacity-0-hover",
		},
		{
			name:      "hover underline opacity 10",
			className: "link-underline-opacity-10-hover",
		},
		{
			name:      "hover underline opacity 25",
			className: "link-underline-opacity-25-hover",
		},
		{
			name:      "hover underline opacity 50",
			className: "link-underline-opacity-50-hover",
		},
		{
			name:      "hover underline opacity 75",
			className: "link-underline-opacity-75-hover",
		},
		{
			name:      "hover underline opacity 100",
			className: "link-underline-opacity-100-hover",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass("link-underline", tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// COLORED LINKS TESTS

func TestLinkColors(t *testing.T) {
	tests := []struct {
		name      string
		className string
		color     string
	}{
		{
			name:      "primary link",
			className: "link-primary",
			color:     "primary",
		},
		{
			name:      "secondary link",
			className: "link-secondary",
			color:     "secondary",
		},
		{
			name:      "success link",
			className: "link-success",
			color:     "success",
		},
		{
			name:      "danger link",
			className: "link-danger",
			color:     "danger",
		},
		{
			name:      "warning link",
			className: "link-warning",
			color:     "warning",
		},
		{
			name:      "info link",
			className: "link-info",
			color:     "info",
		},
		{
			name:      "light link",
			className: "link-light",
			color:     "light",
		},
		{
			name:      "dark link",
			className: "link-dark",
			color:     "dark",
		},
		{
			name:      "body emphasis link",
			className: "link-body-emphasis",
			color:     "body-emphasis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := bs.Link("#", WithClass(tt.className))
			classList := link.Root().ClassList()
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// COMPLEX LINK COMBINATIONS TESTS

func TestLinkComplexCombination1(t *testing.T) {
	// Primary link with offset, underline opacity, and hover effect
	// <a href="#" class="link-primary link-offset-2 link-underline-opacity-25 link-underline-opacity-100-hover">Primary link</a>
	link := bs.Link("#", WithClass(
		"link-primary",
		"link-offset-2",
		"link-underline-opacity-25",
		"link-underline-opacity-100-hover",
	))
	classList := link.Root().ClassList()

	tests := []struct {
		name      string
		className string
	}{
		{"has link-primary", "link-primary"},
		{"has link-offset-2", "link-offset-2"},
		{"has link-underline-opacity-25", "link-underline-opacity-25"},
		{"has link-underline-opacity-100-hover", "link-underline-opacity-100-hover"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkComplexCombination2(t *testing.T) {
	// Link with offset, underline, and opacity variations on hover
	// <a class="link-offset-2 link-offset-3-hover link-underline link-underline-opacity-0 link-underline-opacity-75-hover" href="#">
	link := bs.Link("#", WithClass(
		"link-offset-2",
		"link-offset-3-hover",
		"link-underline",
		"link-underline-opacity-0",
		"link-underline-opacity-75-hover",
	))
	classList := link.Root().ClassList()

	tests := []struct {
		name      string
		className string
	}{
		{"has link-offset-2", "link-offset-2"},
		{"has link-offset-3-hover", "link-offset-3-hover"},
		{"has link-underline", "link-underline"},
		{"has link-underline-opacity-0", "link-underline-opacity-0"},
		{"has link-underline-opacity-75-hover", "link-underline-opacity-75-hover"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

func TestLinkComplexCombination3(t *testing.T) {
	// Multiple link utilities combined
	link := bs.Link("#", WithClass(
		"link-danger",
		"link-offset-1",
		"link-opacity-50",
		"link-opacity-75-hover",
		"link-underline-danger",
	))
	classList := link.Root().ClassList()

	tests := []struct {
		name      string
		className string
	}{
		{"has link-danger", "link-danger"},
		{"has link-offset-1", "link-offset-1"},
		{"has link-opacity-50", "link-opacity-50"},
		{"has link-opacity-75-hover", "link-opacity-75-hover"},
		{"has link-underline-danger", "link-underline-danger"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !classList.Contains(tt.className) {
				t.Errorf("Link should have %q class", tt.className)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// LINK CONTENT TESTS

func TestLinkWithTextContent(t *testing.T) {
	link := bs.Link("#")
	link.Append("Click here")

	t.Run("can append text", func(t *testing.T) {
		html := link.Root().InnerHTML()
		if !strings.Contains(html, "Click here") {
			t.Error("Link should contain 'Click here' text")
		}
	})
}

func TestLinkWithMultipleChildren(t *testing.T) {
	link := bs.Link("#")
	link.Append("Read more ", bs.Icon("arrow-right"))

	t.Run("can append multiple children", func(t *testing.T) {
		html := link.Root().InnerHTML()
		if !strings.Contains(html, "Read more") {
			t.Error("Link should contain 'Read more' text")
		}
		// Icon will be present as <i> element
		if !strings.Contains(html, "<i") {
			t.Error("Link should contain icon element")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// LINK EDGE CASES

func TestLinkEmptyHref(t *testing.T) {
	link := bs.Link("")

	t.Run("accepts empty href", func(t *testing.T) {
		if link == nil {
			t.Fatal("bs.Link() with empty href should not return nil")
		}
	})

	t.Run("has empty href attribute", func(t *testing.T) {
		href := link.Root().GetAttribute("href")
		if href != "" {
			t.Errorf("Link with empty href should have empty href attribute, got %q", href)
		}
	})
}

func TestLinkWithID(t *testing.T) {
	link := bs.Link("#", WithID("my-link"))

	t.Run("has correct ID", func(t *testing.T) {
		if link.ID() != "my-link" {
			t.Errorf("Link.ID() = %q, want my-link", link.ID())
		}
	})

	t.Run("ID is in HTML", func(t *testing.T) {
		html := link.Root().OuterHTML()
		if !strings.Contains(html, `id="my-link"`) {
			t.Error("Link HTML should contain id attribute")
		}
	})
}

func TestLinkWithAttributes(t *testing.T) {
	link := bs.Link("https://example.com",
		WithAttr("target", "_blank"),
		WithAttr("rel", "noopener noreferrer"),
		WithAttr("title", "External link"),
	)

	tests := []struct {
		name     string
		attr     string
		expected string
	}{
		{"has target attribute", "target", "_blank"},
		{"has rel attribute", "rel", "noopener noreferrer"},
		{"has title attribute", "title", "External link"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := link.Root().GetAttribute(tt.attr)
			if value != tt.expected {
				t.Errorf("Link %s = %q, want %q", tt.attr, value, tt.expected)
			}
		})
	}
}
