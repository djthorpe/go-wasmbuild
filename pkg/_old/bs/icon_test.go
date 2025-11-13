package bs_test

import (
	"fmt"
	"strings"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// BASIC ICON TESTS

func TestIconCreation(t *testing.T) {
	tests := []struct {
		name     string
		iconName string
	}{
		{
			name:     "heart icon",
			iconName: "heart",
		},
		{
			name:     "star icon",
			iconName: "star",
		},
		{
			name:     "house icon",
			iconName: "house",
		},
		{
			name:     "search icon",
			iconName: "search",
		},
		{
			name:     "gear icon",
			iconName: "gear",
		},
		{
			name:     "bell icon",
			iconName: "bell",
		},
		{
			name:     "hyphenated icon name",
			iconName: "arrow-up-circle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := bs.Icon(tt.iconName)
			if icon == nil {
				t.Fatal("bs.Icon() returned nil")
			}
			if icon.Name() != bs.ViewIcon {
				t.Errorf("bs.Icon().Name() = %v, want %v", icon.Name(), bs.ViewIcon)
			}
		})
	}
}

func TestIconViewInterface(t *testing.T) {
	icon := bs.Icon("heart")

	t.Run("Name returns correct view name", func(t *testing.T) {
		if icon.Name() != bs.ViewIcon {
			t.Errorf("Icon.Name() = %v, want %v", icon.Name(), bs.ViewIcon)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := icon.Root()
		if root == nil {
			t.Error("Icon.Root() returned nil")
		}
	})

	t.Run("Root returns I element", func(t *testing.T) {
		tagName := icon.Root().TagName()
		if tagName != "I" {
			t.Errorf("Icon root element tag = %v, want I", tagName)
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Icon.ID() panicked: %v", r)
			}
		}()
		_ = icon.ID()
	})
}

func TestIconEmbedding(t *testing.T) {
	// Verify that icon properly embeds View
	icon := bs.Icon("star")

	// icon should be able to be used as a View
	var v View = icon
	if v.Name() != bs.ViewIcon {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), bs.ViewIcon)
	}
}

func TestIconMultipleInstances(t *testing.T) {
	// Create multiple icon elements and verify they're independent
	icon1 := bs.Icon("heart")
	icon2 := bs.Icon("star")

	if icon1 == icon2 {
		t.Error("Each bs.Icon() call should create a distinct instance")
	}

	if icon1.Root() == icon2.Root() {
		t.Error("Icons should have separate DOM elements")
	}
}

func TestIconWithID(t *testing.T) {
	tests := []struct {
		name     string
		iconName string
		id       string
	}{
		{
			name:     "heart icon with ID",
			iconName: "heart",
			id:       "favorite-icon",
		},
		{
			name:     "gear icon with ID",
			iconName: "gear",
			id:       "settings-icon",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Icons are created without options in the constructor,
			// so we need to apply ID after creation
			icon := bs.Icon(tt.iconName)
			icon.Opts(WithID(tt.id))

			if icon.ID() != tt.id {
				t.Errorf("Icon.ID() = %v, want %v", icon.ID(), tt.id)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// ICON CLASS TESTS

func TestIconBootstrapClass(t *testing.T) {
	tests := []struct {
		name          string
		iconName      string
		expectedClass string
	}{
		{
			name:          "heart icon",
			iconName:      "heart",
			expectedClass: "bi-heart",
		},
		{
			name:          "star icon",
			iconName:      "star",
			expectedClass: "bi-star",
		},
		{
			name:          "search icon",
			iconName:      "search",
			expectedClass: "bi-search",
		},
		{
			name:          "hyphenated icon name",
			iconName:      "arrow-up-circle",
			expectedClass: "bi-arrow-up-circle",
		},
		{
			name:          "icon with fill variant",
			iconName:      "heart-fill",
			expectedClass: "bi-heart-fill",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := bs.Icon(tt.iconName)
			classList := icon.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Icon should have class %q", tt.expectedClass)
			}
		})
	}
}

func TestIconWithAdditionalClasses(t *testing.T) {
	icon := bs.Icon("heart")
	icon.Opts(WithClass("text-danger", "fs-3"))

	classList := icon.Root().ClassList()

	if !classList.Contains("bi-heart") {
		t.Error("Icon should have default 'bi-heart' class")
	}

	if !classList.Contains("text-danger") {
		t.Error("Icon should have 'text-danger' class")
	}

	if !classList.Contains("fs-3") {
		t.Error("Icon should have 'fs-3' class")
	}
}

///////////////////////////////////////////////////////////////////////////////
// ICON APPEND/INSERT RESTRICTION TESTS

func TestIconAppendPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Icon.Append() should panic")
		} else {
			// Verify the panic message mentions append not being supported
			msg := fmt.Sprint(r)
			if !strings.Contains(strings.ToLower(msg), "append") {
				t.Errorf("Panic message should mention 'append', got: %v", r)
			}
		}
	}()

	icon := bs.Icon("heart")
	icon.Append("text") // Should panic
}

func TestIconInsertPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Icon.Insert() should panic")
		} else {
			// Verify the panic message mentions insert not being supported
			msg := fmt.Sprint(r)
			if !strings.Contains(strings.ToLower(msg), "insert") {
				t.Errorf("Panic message should mention 'insert', got: %v", r)
			}
		}
	}()

	icon := bs.Icon("star")
	icon.Content("text") // Should panic
}

///////////////////////////////////////////////////////////////////////////////
// ICON WITH OPTIONS TESTS

func TestIconWithOptions(t *testing.T) {
	tests := []struct {
		name        string
		iconName    string
		opts        []Opt
		checkFunc   func(*testing.T, View)
		description string
	}{
		{
			name:        "icon with ID",
			iconName:    "heart",
			opts:        []Opt{WithID("my-icon")},
			description: "should have ID set",
			checkFunc: func(t *testing.T, v View) {
				if v.ID() != "my-icon" {
					t.Errorf("ID = %v, want 'my-icon'", v.ID())
				}
			},
		},
		{
			name:        "icon with data attribute",
			iconName:    "star",
			opts:        []Opt{WithAttr("data-toggle", "tooltip")},
			description: "should have data attribute",
			checkFunc: func(t *testing.T, v View) {
				if v.Root().GetAttribute("data-toggle") != "tooltip" {
					t.Error("Should have data-toggle attribute")
				}
			},
		},
		{
			name:        "icon with aria-label",
			iconName:    "search",
			opts:        []Opt{WithAttr("aria-label", "Search")},
			description: "should have aria-label",
			checkFunc: func(t *testing.T, v View) {
				if v.Root().GetAttribute("aria-label") != "Search" {
					t.Error("Should have aria-label attribute")
				}
			},
		},
		{
			name:        "icon with multiple classes",
			iconName:    "bell",
			opts:        []Opt{WithClass("text-primary", "me-2")},
			description: "should have additional classes",
			checkFunc: func(t *testing.T, v View) {
				classList := v.Root().ClassList()
				if !classList.Contains("text-primary") {
					t.Error("Should have text-primary class")
				}
				if !classList.Contains("me-2") {
					t.Error("Should have me-2 class")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := bs.Icon(tt.iconName)
			icon.Opts(tt.opts...)
			tt.checkFunc(t, icon)
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// ICON VARIANTS TESTS

func TestIconVariants(t *testing.T) {
	tests := []struct {
		name          string
		iconName      string
		expectedClass string
		description   string
	}{
		{
			name:          "filled icon",
			iconName:      "heart-fill",
			expectedClass: "bi-heart-fill",
			description:   "should have fill variant class",
		},
		{
			name:          "outlined icon",
			iconName:      "heart",
			expectedClass: "bi-heart",
			description:   "should have outline variant class",
		},
		{
			name:          "circle icon",
			iconName:      "check-circle",
			expectedClass: "bi-check-circle",
			description:   "should have circle variant class",
		},
		{
			name:          "square icon",
			iconName:      "check-square",
			expectedClass: "bi-check-square",
			description:   "should have square variant class",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := bs.Icon(tt.iconName)
			classList := icon.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Icon should have class %q for %s", tt.expectedClass, tt.description)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// OUTER HTML TESTS

func TestIconOuterHTML(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() View
		contains []string
	}{
		{
			name: "basic heart icon",
			setup: func() View {
				return bs.Icon("heart")
			},
			contains: []string{
				"<i",
				"class=\"bi-heart\"",
				"</i>",
			},
		},
		{
			name: "icon with ID",
			setup: func() View {
				icon := bs.Icon("star")
				icon.Opts(WithID("my-star"))
				return icon
			},
			contains: []string{
				"<i",
				"id=\"my-star\"",
				"class=\"bi-star\"",
			},
		},
		{
			name: "icon with additional classes",
			setup: func() View {
				icon := bs.Icon("search")
				icon.Opts(WithClass("text-muted", "me-2"))
				return icon
			},
			contains: []string{
				"<i",
				"bi-search",
				"text-muted",
				"me-2",
			},
		},
		{
			name: "icon with aria-label",
			setup: func() View {
				icon := bs.Icon("bell")
				icon.Opts(WithAttr("aria-label", "Notifications"))
				return icon
			},
			contains: []string{
				"<i",
				"bi-bell",
				"aria-label=\"Notifications\"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := tt.setup()
			html := icon.Root().OuterHTML()

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

func TestIconEdgeCases(t *testing.T) {
	t.Run("empty icon name", func(t *testing.T) {
		icon := bs.Icon("")
		classList := icon.Root().ClassList()

		// Should have "bi-" prefix even with empty name
		if !classList.Contains("bi-") {
			t.Error("Icon should have 'bi-' class even with empty name")
		}
	})

	t.Run("icon with numbers in name", func(t *testing.T) {
		icon := bs.Icon("1-circle")
		classList := icon.Root().ClassList()

		if !classList.Contains("bi-1-circle") {
			t.Error("Icon should support numbers in name")
		}
	})

	t.Run("icon with multiple hyphens", func(t *testing.T) {
		icon := bs.Icon("arrow-up-right-circle")
		classList := icon.Root().ClassList()

		if !classList.Contains("bi-arrow-up-right-circle") {
			t.Error("Icon should support multiple hyphens in name")
		}
	})

	t.Run("icon name with uppercase converts to class", func(t *testing.T) {
		// Note: Bootstrap Icons are lowercase, but the function should accept any string
		icon := bs.Icon("HEART")
		classList := icon.Root().ClassList()

		if !classList.Contains("bi-HEART") {
			t.Error("Icon should accept uppercase names (though not standard)")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// ATTRIBUTES TESTS

func TestIconWithAttributes(t *testing.T) {
	icon := bs.Icon("gear")
	icon.Opts(
		WithID("settings-icon"),
		WithAttr("data-action", "open-settings"),
		WithAttr("aria-label", "Settings"),
		WithAttr("role", "button"),
	)

	root := icon.Root()

	if icon.ID() != "settings-icon" {
		t.Errorf("ID = %q, want 'settings-icon'", icon.ID())
	}

	if root.GetAttribute("data-action") != "open-settings" {
		t.Error("Should have data-action attribute")
	}

	if root.GetAttribute("aria-label") != "Settings" {
		t.Error("Should have aria-label attribute")
	}

	if root.GetAttribute("role") != "button" {
		t.Error("Should have role attribute")
	}
}

///////////////////////////////////////////////////////////////////////////////
// COMMON ICON NAMES TESTS

func TestCommonBootstrapIcons(t *testing.T) {
	// Test a variety of common Bootstrap Icons
	commonIcons := []string{
		"alarm",
		"archive",
		"arrow-right",
		"bag",
		"bell",
		"bookmark",
		"calendar",
		"camera",
		"cart",
		"check",
		"chevron-down",
		"clock",
		"cloud",
		"download",
		"envelope",
		"eye",
		"file",
		"folder",
		"gear",
		"heart",
		"home",
		"image",
		"info-circle",
		"link",
		"list",
		"lock",
		"menu",
		"person",
		"phone",
		"play",
		"plus",
		"question-circle",
		"search",
		"share",
		"star",
		"trash",
		"upload",
		"x",
		"zoom-in",
	}

	for _, iconName := range commonIcons {
		t.Run(iconName, func(t *testing.T) {
			icon := bs.Icon(iconName)
			if icon == nil {
				t.Errorf("Failed to create icon %q", iconName)
			}

			classList := icon.Root().ClassList()
			expectedClass := "bi-" + iconName
			if !classList.Contains(expectedClass) {
				t.Errorf("Icon should have class %q", expectedClass)
			}
		})
	}
}
