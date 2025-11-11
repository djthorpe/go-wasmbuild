package bs_test

import (
	"strings"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

///////////////////////////////////////////////////////////////////////////////
// BASIC BUTTON TESTS

func TestButtonCreation(t *testing.T) {
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
			opts: []Opt{WithID("submit-btn")},
		},
		{
			name: "create with color class",
			opts: []Opt{WithClass("btn-success")},
		},
		{
			name: "create with combined options",
			opts: []Opt{
				WithID("action-btn"),
				WithClass("btn-lg"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(tt.opts...)
			if b == nil {
				t.Fatal("bs.Button() returned nil")
			}
			if b.Name() != bs.ViewButton {
				t.Errorf("bs.Button().Name() = %v, want %v", b.Name(), bs.ViewButton)
			}
		})
	}
}

func TestButtonViewInterface(t *testing.T) {
	b := bs.Button()

	t.Run("Name returns correct view name", func(t *testing.T) {
		if b.Name() != bs.ViewButton {
			t.Errorf("Button.Name() = %v, want %v", b.Name(), bs.ViewButton)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := b.Root()
		if root == nil {
			t.Error("Button.Root() returned nil")
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Button.ID() panicked: %v", r)
			}
		}()
		_ = b.ID()
	})

	t.Run("Root has correct tag name", func(t *testing.T) {
		root := b.Root()
		if root.TagName() != "BUTTON" {
			t.Errorf("Button.Root().TagName() = %v, want BUTTON", root.TagName())
		}
	})
}

func TestButtonDefaultClasses(t *testing.T) {
	b := bs.Button()
	classList := b.Root().ClassList()

	t.Run("has btn class", func(t *testing.T) {
		if !classList.Contains("btn") {
			t.Error("Button should have 'btn' class")
		}
	})

	t.Run("has btn-primary class by default", func(t *testing.T) {
		if !classList.Contains("btn-primary") {
			t.Error("Button should have 'btn-primary' class by default")
		}
	})
}

func TestButtonDefaultAttributes(t *testing.T) {
	b := bs.Button()
	root := b.Root()

	t.Run("has type='button' attribute", func(t *testing.T) {
		if !root.HasAttribute("type") {
			t.Error("Button should have 'type' attribute")
		}
		if root.GetAttribute("type") != "button" {
			t.Errorf("Button type = %v, want 'button'", root.GetAttribute("type"))
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON COLOR VARIANTS (Bootstrap 5.3 Documentation)

func TestButtonColorVariants(t *testing.T) {
	tests := []struct {
		name          string
		opts          []Opt
		expectedClass string
	}{
		{"primary button", []Opt{}, "btn-primary"},
		{"secondary button", []Opt{WithClass("btn-secondary")}, "btn-secondary"},
		{"success button", []Opt{WithClass("btn-success")}, "btn-success"},
		{"danger button", []Opt{WithClass("btn-danger")}, "btn-danger"},
		{"warning button", []Opt{WithClass("btn-warning")}, "btn-warning"},
		{"info button", []Opt{WithClass("btn-info")}, "btn-info"},
		{"light button", []Opt{WithClass("btn-light")}, "btn-light"},
		{"dark button", []Opt{WithClass("btn-dark")}, "btn-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(tt.opts...)
			classList := b.Root().ClassList()

			if !classList.Contains("btn") {
				t.Error("Button should contain 'btn' class")
			}
			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Button should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// OUTLINE BUTTON VARIANTS (Bootstrap 5.3 Documentation)

func TestOutlineButtonCreation(t *testing.T) {
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
			opts: []Opt{WithID("outline-btn")},
		},
		{
			name: "create with size",
			opts: []Opt{WithClass("btn-lg")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.OutlineButton(tt.opts...)
			if b == nil {
				t.Fatal("bs.OutlineButton() returned nil")
			}
			if b.Name() != bs.ViewButton {
				t.Errorf("bs.OutlineButton().Name() = %v, want %v", b.Name(), bs.ViewButton)
			}
		})
	}
}

func TestOutlineButtonDefaultClasses(t *testing.T) {
	b := bs.OutlineButton()
	classList := b.Root().ClassList()

	t.Run("has btn class", func(t *testing.T) {
		if !classList.Contains("btn") {
			t.Error("OutlineButton should have 'btn' class")
		}
	})

	t.Run("has btn-outline-primary class by default", func(t *testing.T) {
		if !classList.Contains("btn-outline-primary") {
			t.Error("OutlineButton should have 'btn-outline-primary' class by default")
		}
	})

	t.Run("has btn-outline prefix class", func(t *testing.T) {
		if !classList.Contains("btn-outline") {
			t.Error("OutlineButton should have 'btn-outline' prefix class")
		}
	})
}

func TestOutlineButtonColorVariants(t *testing.T) {
	tests := []struct {
		name          string
		opts          []Opt
		expectedClass string
	}{
		{"outline primary button", []Opt{}, "btn-outline-primary"},
		{"outline secondary button", []Opt{WithClass("btn-outline-secondary")}, "btn-outline-secondary"},
		{"outline success button", []Opt{WithClass("btn-outline-success")}, "btn-outline-success"},
		{"outline danger button", []Opt{WithClass("btn-outline-danger")}, "btn-outline-danger"},
		{"outline warning button", []Opt{WithClass("btn-outline-warning")}, "btn-outline-warning"},
		{"outline info button", []Opt{WithClass("btn-outline-info")}, "btn-outline-info"},
		{"outline light button", []Opt{WithClass("btn-outline-light")}, "btn-outline-light"},
		{"outline dark button", []Opt{WithClass("btn-outline-dark")}, "btn-outline-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.OutlineButton(tt.opts...)
			classList := b.Root().ClassList()

			if !classList.Contains("btn") {
				t.Error("OutlineButton should contain 'btn' class")
			}
			if !classList.Contains(tt.expectedClass) {
				t.Errorf("OutlineButton should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON SIZES (Bootstrap 5.3 Documentation)

func TestButtonSizes(t *testing.T) {
	tests := []struct {
		name          string
		size          bs.Size
		expectedClass string
	}{
		{"small button", bs.SizeSmall, "btn-sm"},
		{"large button", bs.SizeLarge, "btn-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(bs.WithSize(tt.size))
			classList := b.Root().ClassList()

			if !classList.Contains("btn") {
				t.Error("Button should contain 'btn' class")
			}
			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Button should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

func TestOutlineButtonSizes(t *testing.T) {
	tests := []struct {
		name          string
		size          bs.Size
		expectedClass string
	}{
		{"small outline button", bs.SizeSmall, "btn-sm"},
		{"large outline button", bs.SizeLarge, "btn-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.OutlineButton(bs.WithSize(tt.size))
			classList := b.Root().ClassList()

			if !classList.Contains("btn") {
				t.Error("OutlineButton should contain 'btn' class")
			}
			if !classList.Contains("btn-outline-primary") {
				t.Error("OutlineButton should contain 'btn-outline-primary' class")
			}
			if !classList.Contains(tt.expectedClass) {
				t.Errorf("OutlineButton should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON DISABLED STATE (Bootstrap 5.3 Documentation)

func TestButtonDisabled(t *testing.T) {
	t.Run("button without disabled attribute", func(t *testing.T) {
		b := bs.Button()
		if b.Disabled() {
			t.Error("Button should not be disabled by default")
		}
	})

	t.Run("button with disabled attribute", func(t *testing.T) {
		b := bs.Button(WithAttr("disabled", ""))
		if !b.Disabled() {
			t.Error("Button should be disabled when disabled attribute is set")
		}
	})

	t.Run("outline button with disabled attribute", func(t *testing.T) {
		b := bs.OutlineButton(WithAttr("disabled", ""))
		if !b.Disabled() {
			t.Error("OutlineButton should be disabled when disabled attribute is set")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON ACTIVE STATE (Bootstrap 5.3 Documentation)

func TestButtonActive(t *testing.T) {
	t.Run("button without active class", func(t *testing.T) {
		b := bs.Button()
		if b.Active() {
			t.Error("Button should not be active by default")
		}
	})

	t.Run("button with active class", func(t *testing.T) {
		b := bs.Button(WithClass("active"))
		if !b.Active() {
			t.Error("Button should be active when active class is set")
		}
	})

	t.Run("outline button with active class", func(t *testing.T) {
		b := bs.OutlineButton(WithClass("active"))
		if !b.Active() {
			t.Error("OutlineButton should be active when active class is set")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON WITH ADDITIONAL CLASSES

func TestButtonWithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		options         []Opt
		expectedClasses []string
	}{
		{
			name:            "button with single additional class",
			options:         []Opt{WithClass("custom-btn")},
			expectedClasses: []string{"btn", "btn-primary", "custom-btn"},
		},
		{
			name:            "button with multiple additional classes",
			options:         []Opt{WithClass("class1", "class2", "class3")},
			expectedClasses: []string{"btn", "btn-primary", "class1", "class2", "class3"},
		},
		{
			name:            "button with text-nowrap",
			options:         []Opt{WithClass("text-nowrap")},
			expectedClasses: []string{"btn", "btn-primary", "text-nowrap"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(tt.options...)
			classList := b.Root().ClassList()

			for _, expectedClass := range tt.expectedClasses {
				if !classList.Contains(expectedClass) {
					t.Errorf("Button should contain class '%s'", expectedClass)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON WITH UTILITY OPTIONS

func TestButtonWithMargin(t *testing.T) {
	tests := []struct {
		name          string
		position      bs.Position
		margin        int
		expectedClass string
	}{
		{"margin top 2", bs.Top, 2, "mt-2"},
		{"margin bottom 3", bs.Bottom, 3, "mb-3"},
		{"margin start 1", bs.Start, 1, "ms-1"},
		{"margin end 4", bs.End, 4, "me-4"},
		{"margin x 2", bs.X, 2, "mx-2"},
		{"margin y 3", bs.Y, 3, "my-3"},
		{"margin all 2", bs.All, 2, "m-2"},
		{"negative margin top 1", bs.Top, -1, "mt-n1"},
		{"negative margin bottom 2", bs.Bottom, -2, "mb-n2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(bs.WithMargin(tt.position, tt.margin))
			classList := b.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Button should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

func TestButtonWithPadding(t *testing.T) {
	tests := []struct {
		name          string
		position      bs.Position
		padding       int
		expectedClass string
	}{
		{"padding top 2", bs.Top, 2, "pt-2"},
		{"padding bottom 3", bs.Bottom, 3, "pb-3"},
		{"padding start 1", bs.Start, 1, "ps-1"},
		{"padding end 4", bs.End, 4, "pe-4"},
		{"padding x 2", bs.X, 2, "px-2"},
		{"padding y 3", bs.Y, 3, "py-3"},
		{"padding all 2", bs.All, 2, "p-2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(bs.WithPadding(tt.position, tt.padding))
			classList := b.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Button should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON SPACING COMBINATIONS

func TestButtonWithCombinedMarginAndPadding(t *testing.T) {
	b := bs.Button(
		bs.WithMargin(bs.All, 2),
		bs.WithPadding(bs.X, 3),
	)
	classList := b.Root().ClassList()

	if !classList.Contains("m-2") {
		t.Error("Button should contain 'm-2' class")
	}
	if !classList.Contains("px-3") {
		t.Error("Button should contain 'px-3' class")
	}
}

func TestButtonWithMultipleSpacingUtilities(t *testing.T) {
	// Test that applying margin twice replaces the first one
	t.Run("applying margin twice replaces first", func(t *testing.T) {
		b := bs.Button(
			bs.WithMargin(bs.Top, 2),
			bs.WithMargin(bs.Top, 3), // This should replace mt-2 with mt-3
		)
		classList := b.Root().ClassList()

		// Should only have mt-3, not mt-2
		if classList.Contains("mt-2") {
			t.Error("Button should not contain 'mt-2' class after being replaced")
		}
		if !classList.Contains("mt-3") {
			t.Error("Button should contain 'mt-3' class")
		}
	})

	// Test that applying padding twice replaces the first one
	t.Run("applying padding twice replaces first", func(t *testing.T) {
		b := bs.Button(
			bs.WithPadding(bs.Start, 2),
			bs.WithPadding(bs.Start, 4), // This should replace ps-2 with ps-4
		)
		classList := b.Root().ClassList()

		// Should only have ps-4, not ps-2
		if classList.Contains("ps-2") {
			t.Error("Button should not contain 'ps-2' class after being replaced")
		}
		if !classList.Contains("ps-4") {
			t.Error("Button should contain 'ps-4' class")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON ACCESSIBILITY

func TestButtonAccessibilityAttributes(t *testing.T) {
	tests := []struct {
		name      string
		options   []Opt
		validator func(t *testing.T, root Element)
	}{
		{
			name:    "with aria-label",
			options: []Opt{WithAttr("aria-label", "Close dialog")},
			validator: func(t *testing.T, root Element) {
				if root.GetAttribute("aria-label") != "Close dialog" {
					t.Error("Button should have aria-label='Close dialog'")
				}
			},
		},
		{
			name:    "with role",
			options: []Opt{WithAttr("role", "button")},
			validator: func(t *testing.T, root Element) {
				if root.GetAttribute("role") != "button" {
					t.Error("Button should have role='button'")
				}
			},
		},
		{
			name: "with aria-pressed for toggle",
			options: []Opt{
				WithAttr("data-bs-toggle", "button"),
				WithAttr("aria-pressed", "false"),
			},
			validator: func(t *testing.T, root Element) {
				if root.GetAttribute("data-bs-toggle") != "button" {
					t.Error("Button should have data-bs-toggle='button'")
				}
				if root.GetAttribute("aria-pressed") != "false" {
					t.Error("Button should have aria-pressed='false'")
				}
			},
		},
		{
			name:    "with aria-disabled",
			options: []Opt{WithAttr("aria-disabled", "true")},
			validator: func(t *testing.T, root Element) {
				if root.GetAttribute("aria-disabled") != "true" {
					t.Error("Button should have aria-disabled='true'")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(tt.options...)
			tt.validator(t, b.Root())
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON TOGGLE STATE (Bootstrap 5.3 Documentation)

func TestButtonToggleState(t *testing.T) {
	t.Run("toggle button with data-bs-toggle", func(t *testing.T) {
		b := bs.Button(
			WithAttr("data-bs-toggle", "button"),
			WithClass("active"),
			WithAttr("aria-pressed", "true"),
		)
		root := b.Root()

		if root.GetAttribute("data-bs-toggle") != "button" {
			t.Error("Button should have data-bs-toggle='button'")
		}
		if !root.ClassList().Contains("active") {
			t.Error("Button should contain 'active' class")
		}
		if root.GetAttribute("aria-pressed") != "true" {
			t.Error("Button should have aria-pressed='true'")
		}
	})

	t.Run("disabled toggle button", func(t *testing.T) {
		b := bs.Button(
			WithAttr("data-bs-toggle", "button"),
			WithAttr("disabled", ""),
		)

		if !b.Disabled() {
			t.Error("Button should be disabled")
		}
		if b.Root().GetAttribute("data-bs-toggle") != "button" {
			t.Error("Button should have data-bs-toggle='button'")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON MULTIPLE INSTANCES

func TestButtonMultipleInstances(t *testing.T) {
	b1 := bs.Button(WithID("btn-1"), WithClass("btn-success"))
	b2 := bs.Button(WithID("btn-2"), WithClass("btn-danger"))
	b3 := bs.OutlineButton(WithID("btn-3"), WithClass("btn-outline-info"))

	if b1 == b2 || b1 == b3 || b2 == b3 {
		t.Error("Each button call should create a distinct instance")
	}

	if b1.ID() == b2.ID() || b1.ID() == b3.ID() || b2.ID() == b3.ID() {
		t.Error("Buttons with different IDs should maintain separate state")
	}

	if !b1.Root().ClassList().Contains("btn-success") {
		t.Error("Button 1 should have btn-success class")
	}
	if !b2.Root().ClassList().Contains("btn-danger") {
		t.Error("Button 2 should have btn-danger class")
	}
	if !b3.Root().ClassList().Contains("btn-outline-info") {
		t.Error("Button 3 should have btn-outline-info class")
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON EDGE CASES

func TestButtonEdgeCases(t *testing.T) {
	t.Run("button with duplicate classes", func(t *testing.T) {
		b := bs.Button(WithClass("duplicate", "duplicate"))
		classList := b.Root().ClassList()
		if !classList.Contains("duplicate") {
			t.Error("Button should contain 'duplicate' class")
		}
	})

	t.Run("button with empty ID", func(t *testing.T) {
		b := bs.Button(WithID(""))
		// Should not panic and should create button
		if b == nil {
			t.Error("Button should be created even with empty ID")
		}
	})

	t.Run("outline button with conflicting color classes", func(t *testing.T) {
		b := bs.OutlineButton(WithClass("btn-outline-success"))
		classList := b.Root().ClassList()
		// Both btn-outline-primary and btn-outline-success will be present
		if !classList.Contains("btn-outline-primary") {
			t.Error("OutlineButton should still have default btn-outline-primary class")
		}
		if !classList.Contains("btn-outline-success") {
			t.Error("OutlineButton should have btn-outline-success class")
		}
	})
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON GROUP TESTS (Bootstrap 5.3 Documentation)

func TestButtonGroupCreation(t *testing.T) {
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
			opts: []Opt{WithID("btn-group-1")},
		},
		{
			name: "create with aria-label",
			opts: []Opt{WithAttr("aria-label", "Basic example")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg := bs.ButtonGroup(tt.opts...)
			if bg == nil {
				t.Fatal("bs.ButtonGroup() returned nil")
			}
			if bg.Name() != bs.ViewButtonGroup {
				t.Errorf("bs.ButtonGroup().Name() = %v, want %v", bg.Name(), bs.ViewButtonGroup)
			}
		})
	}
}

func TestButtonGroupViewInterface(t *testing.T) {
	bg := bs.ButtonGroup()

	t.Run("Name returns correct view name", func(t *testing.T) {
		if bg.Name() != bs.ViewButtonGroup {
			t.Errorf("ButtonGroup.Name() = %v, want %v", bg.Name(), bs.ViewButtonGroup)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := bg.Root()
		if root == nil {
			t.Error("ButtonGroup.Root() returned nil")
		}
	})

	t.Run("Root has correct tag name", func(t *testing.T) {
		root := bg.Root()
		if root.TagName() != "DIV" {
			t.Errorf("ButtonGroup.Root().TagName() = %v, want DIV", root.TagName())
		}
	})
}

func TestButtonGroupDefaultClasses(t *testing.T) {
	bg := bs.ButtonGroup()
	classList := bg.Root().ClassList()

	t.Run("has btn-group class", func(t *testing.T) {
		if !classList.Contains("btn-group") {
			t.Error("ButtonGroup should have 'btn-group' class")
		}
	})
}

func TestButtonGroupDefaultAttributes(t *testing.T) {
	bg := bs.ButtonGroup()
	root := bg.Root()

	t.Run("has role='group' attribute", func(t *testing.T) {
		if !root.HasAttribute("role") {
			t.Error("ButtonGroup should have 'role' attribute")
		}
		if root.GetAttribute("role") != "group" {
			t.Errorf("ButtonGroup role = %v, want 'group'", root.GetAttribute("role"))
		}
	})
}

func TestButtonGroupWithSize(t *testing.T) {
	tests := []struct {
		name          string
		size          bs.Size
		expectedClass string
	}{
		{"small button group", bs.SizeSmall, "btn-group-sm"},
		{"large button group", bs.SizeLarge, "btn-group-lg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bg := bs.ButtonGroup(bs.WithSize(tt.size))
			classList := bg.Root().ClassList()

			if !classList.Contains("btn-group") {
				t.Error("ButtonGroup should contain 'btn-group' class")
			}
			if !classList.Contains(tt.expectedClass) {
				t.Errorf("ButtonGroup should contain '%s' class", tt.expectedClass)
			}
		})
	}
}

func TestButtonGroupWithAriaLabel(t *testing.T) {
	bg := bs.ButtonGroup(WithAttr("aria-label", "Button group example"))
	root := bg.Root()

	if root.GetAttribute("aria-label") != "Button group example" {
		t.Error("ButtonGroup should have aria-label='Button group example'")
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON COMBINED FEATURES

func TestButtonCombinedFeatures(t *testing.T) {
	tests := []struct {
		name      string
		options   []Opt
		validator func(t *testing.T, root Element)
	}{
		{
			name: "large button with margin and custom class",
			options: []Opt{
				bs.WithSize(bs.SizeLarge),
				bs.WithMargin(bs.All, 3),
				WithClass("custom-btn"),
			},
			validator: func(t *testing.T, root Element) {
				classList := root.ClassList()
				if !classList.Contains("btn-lg") {
					t.Error("Button should have btn-lg class")
				}
				if !classList.Contains("m-3") {
					t.Error("Button should have m-3 class")
				}
				if !classList.Contains("custom-btn") {
					t.Error("Button should have custom-btn class")
				}
			},
		},
		{
			name: "outline button with padding and margin",
			options: []Opt{
				bs.WithPadding(bs.X, 4),
				bs.WithMargin(bs.Y, 2),
			},
			validator: func(t *testing.T, root Element) {
				classList := root.ClassList()
				if !classList.Contains("px-4") {
					t.Error("Button should have px-4 class")
				}
				if !classList.Contains("my-2") {
					t.Error("Button should have my-2 class")
				}
			},
		},
		{
			name: "button with all accessibility features",
			options: []Opt{
				WithID("submit-btn"),
				WithAttr("aria-label", "Submit the form"),
				WithAttr("role", "button"),
				WithAttr("tabindex", "0"),
			},
			validator: func(t *testing.T, root Element) {
				if root.GetAttribute("id") != "submit-btn" {
					t.Error("Button should have id='submit-btn'")
				}
				if root.GetAttribute("aria-label") != "Submit the form" {
					t.Error("Button should have aria-label='Submit the form'")
				}
				if root.GetAttribute("role") != "button" {
					t.Error("Button should have role='button'")
				}
				if root.GetAttribute("tabindex") != "0" {
					t.Error("Button should have tabindex='0'")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "outline") {
				b := bs.OutlineButton(tt.options...)
				tt.validator(t, b.Root())
			} else {
				b := bs.Button(tt.options...)
				tt.validator(t, b.Root())
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON WITH ID

func TestButtonWithID(t *testing.T) {
	tests := []struct {
		name string
		id   string
	}{
		{"button with simple ID", "submit-btn"},
		{"button with complex ID", "action-button-123"},
		{"outline button with ID", "outline-btn-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if strings.Contains(tt.name, "outline") {
				b := bs.OutlineButton(WithID(tt.id))
				if b.ID() != tt.id {
					t.Errorf("Button.ID() = %v, want %v", b.ID(), tt.id)
				}
				if b.Root().GetAttribute("id") != tt.id {
					t.Errorf("Button root id = %v, want %v", b.Root().GetAttribute("id"), tt.id)
				}
			} else {
				b := bs.Button(WithID(tt.id))
				if b.ID() != tt.id {
					t.Errorf("Button.ID() = %v, want %v", b.ID(), tt.id)
				}
				if b.Root().GetAttribute("id") != tt.id {
					t.Errorf("Button root id = %v, want %v", b.Root().GetAttribute("id"), tt.id)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BUTTON POSITION UTILITIES

func TestButtonWithPositionUtilities(t *testing.T) {
	tests := []struct {
		name          string
		options       []Opt
		expectedClass string
	}{
		{
			name:          "button with text-center",
			options:       []Opt{WithClass("text-center")},
			expectedClass: "text-center",
		},
		{
			name:          "button with d-flex",
			options:       []Opt{WithClass("d-flex")},
			expectedClass: "d-flex",
		},
		{
			name:          "button with justify-content-center",
			options:       []Opt{WithClass("justify-content-center")},
			expectedClass: "justify-content-center",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Button(tt.options...)
			classList := b.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Button should contain '%s' class", tt.expectedClass)
			}
		})
	}
}
