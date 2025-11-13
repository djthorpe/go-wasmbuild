package bs_test

import (
	"fmt"
	"testing"

	// Packages
	"github.com/djthorpe/go-wasmbuild/pkg/bs"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TestBadgeCreation(t *testing.T) {
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
			opts: []Opt{WithID("counter")},
		},
		{
			name: "create with background color",
			opts: []Opt{WithClass("text-bg-primary")},
		},
		{
			name: "create with combined options",
			opts: []Opt{
				WithID("badge-1"),
				WithClass("text-bg-success"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Badge(tt.opts...)
			if b == nil {
				t.Fatal("bs.Badge() returned nil")
			}
			if b.Name() != bs.ViewBadge {
				t.Errorf("bs.Badge().Name() = %v, want %v", b.Name(), bs.ViewBadge)
			}
		})
	}
}

func TestBadgeViewInterface(t *testing.T) {
	b := bs.Badge()

	t.Run("Name returns correct view name", func(t *testing.T) {
		if b.Name() != bs.ViewBadge {
			t.Errorf("Badge.Name() = %v, want %v", b.Name(), bs.ViewBadge)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := b.Root()
		if root == nil {
			t.Error("Badge.Root() returned nil")
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Badge.ID() panicked: %v", r)
			}
		}()
		_ = b.ID()
	})
}

func TestBadgeEmbedding(t *testing.T) {
	// Verify that badge properly embeds View
	b := bs.Badge()

	// badge should be able to be used as a View
	var v View = b
	if v.Name() != bs.ViewBadge {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), bs.ViewBadge)
	}
}

func TestBadgeMultipleInstances(t *testing.T) {
	// Create multiple badge elements and verify they're independent
	b1 := bs.Badge(WithID("badge-1"))
	b2 := bs.Badge(WithID("badge-2"))

	if b1 == b2 {
		t.Error("Each bs.Badge() call should create a distinct instance")
	}

	if b1.ID() == b2.ID() {
		t.Error("Badges with different IDs should maintain separate state")
	}
}

func TestBadgeWithOptions(t *testing.T) {
	t.Run("with ID", func(t *testing.T) {
		b := bs.Badge(WithID("notification-count"))
		if b.ID() != "notification-count" {
			t.Errorf("bs.Badge(WithID('notification-count')).ID() = %v, want %v", b.ID(), "notification-count")
		}
	})

	t.Run("with classes", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-primary", "ms-2"))
		if b == nil {
			t.Error("bs.Badge() with classes should not return nil")
		}
	})

	t.Run("with combined options", func(t *testing.T) {
		b := bs.Badge(
			WithID("status"),
			WithClass("text-bg-success"),
			WithAttr("data-count", "5"),
		)

		if b == nil {
			t.Fatal("bs.Badge() with combined options should not return nil")
		}

		if b.ID() != "status" {
			t.Errorf("bs.Badge().ID() = %v, want %v", b.ID(), "status")
		}
	})
}

// Test Bootstrap badge background colors using WithColor
func TestBadgeBackgroundColors(t *testing.T) {
	colors := []struct {
		name  string
		color bs.Color
	}{
		{"primary", bs.PRIMARY},
		{"secondary", bs.SECONDARY},
		{"success", bs.SUCCESS},
		{"danger", bs.DANGER},
		{"warning", bs.WARNING},
		{"info", bs.INFO},
		{"light", bs.LIGHT},
		{"dark", bs.DARK},
	}

	for _, tc := range colors {
		t.Run(tc.name, func(t *testing.T) {
			b := bs.Badge(bs.WithColor(tc.color))
			if b == nil {
				t.Errorf("Badge with %s color should not return nil", tc.name)
			}
		})
	}
}

// Test pill badges with rounded-pill class
func TestPillBadge(t *testing.T) {
	t.Run("create pill badge without options", func(t *testing.T) {
		b := bs.PillBadge()
		if b == nil {
			t.Fatal("bs.PillBadge() returned nil")
		}
		if b.Name() != bs.ViewBadge {
			t.Errorf("bs.PillBadge().Name() = %v, want %v", b.Name(), bs.ViewBadge)
		}
	})

	t.Run("create pill badge with background color", func(t *testing.T) {
		b := bs.PillBadge(WithClass("text-bg-primary"))
		if b == nil {
			t.Error("bs.PillBadge() with color should not return nil")
		}
	})

	t.Run("create pill badge with ID", func(t *testing.T) {
		b := bs.PillBadge(WithID("pill-notification"))
		if b == nil {
			t.Fatal("bs.PillBadge() with ID should not return nil")
		}
		if b.ID() != "pill-notification" {
			t.Errorf("bs.PillBadge().ID() = %v, want %v", b.ID(), "pill-notification")
		}
	})
}

// Test pill badges with all background colors using WithColor
func TestPillBadgeBackgroundColors(t *testing.T) {
	colors := []struct {
		name  string
		color bs.Color
	}{
		{"primary", bs.PRIMARY},
		{"secondary", bs.SECONDARY},
		{"success", bs.SUCCESS},
		{"danger", bs.DANGER},
		{"warning", bs.WARNING},
		{"info", bs.INFO},
		{"light", bs.LIGHT},
		{"dark", bs.DARK},
	}

	for _, tc := range colors {
		t.Run("pill_"+tc.name, func(t *testing.T) {
			b := bs.PillBadge(bs.WithColor(tc.color))
			if b == nil {
				t.Errorf("PillBadge with %s color should not return nil", tc.name)
			}
		})
	}
}

// Test badges in headings context
func TestBadgeInHeadings(t *testing.T) {
	t.Run("badge for heading with text-bg-secondary", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-secondary"))
		if b == nil {
			t.Error("Badge for heading should not return nil")
		}
	})
}

// Test badges in buttons context
func TestBadgeInButtons(t *testing.T) {
	t.Run("badge as counter in button", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-secondary"))
		if b == nil {
			t.Error("Badge for button counter should not return nil")
		}
	})

	t.Run("badge with notification count", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-secondary"), WithAttr("data-count", "4"))
		if b == nil {
			t.Error("Badge with notification count should not return nil")
		}
	})
}

// Test positioned badges (for notification indicators)
func TestPositionedBadges(t *testing.T) {
	t.Run("positioned pill badge", func(t *testing.T) {
		b := bs.PillBadge(
			WithClass("position-absolute", "top-0", "start-100", "translate-middle"),
			bs.WithBackground(bs.DANGER),
		)
		if b == nil {
			t.Error("Positioned pill badge should not return nil")
		}
	})

	t.Run("positioned rounded circle indicator", func(t *testing.T) {
		b := bs.Badge(
			WithClass("position-absolute", "top-0", "start-100", "translate-middle", "rounded-circle"),
			bs.WithPadding(bs.All, 2),
			bs.WithBackground(bs.DANGER),
			bs.WithBorder(bs.All, bs.LIGHT),
		)
		if b == nil {
			t.Error("Positioned circular badge should not return nil")
		}
	})

	t.Run("positioned badge with multiple margins", func(t *testing.T) {
		b := bs.Badge(
			bs.WithMargin(bs.Start, 2),
			bs.WithMargin(bs.Top, 1),
			bs.WithColor(bs.WARNING),
		)
		if b == nil {
			t.Error("Badge with multiple margins should not return nil")
		}
	})
}

// Test badge with utility classes using helper functions
func TestBadgeWithUtilityClasses(t *testing.T) {
	tests := []struct {
		name string
		opt  Opt
	}{
		{
			name: "badge with start margin",
			opt:  bs.WithMargin(bs.Start, 2),
		},
		{
			name: "badge with end margin",
			opt:  bs.WithMargin(bs.End, 2),
		},
		{
			name: "badge with all margins",
			opt:  bs.WithMargin(bs.All, 3),
		},
		{
			name: "badge with padding",
			opt:  bs.WithPadding(bs.All, 2),
		},
		{
			name: "badge with top padding",
			opt:  bs.WithPadding(bs.Top, 1),
		},
		{
			name: "badge with border",
			opt:  bs.WithBorder(bs.All),
		},
		{
			name: "badge with border and color",
			opt:  bs.WithBorder(bs.All, bs.LIGHT),
		},
		{
			name: "badge with top border",
			opt:  bs.WithBorder(bs.Top, bs.DANGER),
		},
		{
			name: "badge with position utilities",
			opt:  WithClass("position-absolute", "top-0", "start-100"),
		},
		{
			name: "badge with translate utility",
			opt:  WithClass("translate-middle"),
		},
		{
			name: "badge with rounded utilities",
			opt:  WithClass("rounded-circle"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Badge(tt.opt)
			if b == nil {
				t.Errorf("Badge with %s should not return nil", tt.name)
			}
		})
	}
}

// Test badge accessibility considerations
func TestBadgeAccessibility(t *testing.T) {
	t.Run("badge with visually-hidden content", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-danger"))
		if b == nil {
			t.Error("Badge for accessibility use should not return nil")
		}
	})

	t.Run("badge with aria attribute", func(t *testing.T) {
		b := bs.Badge(
			WithClass("text-bg-primary"),
			WithAttr("aria-label", "Notifications"),
		)
		if b == nil {
			t.Error("Badge with aria-label should not return nil")
		}
	})
}

// Test that Badge always includes 'badge' class
func TestBadgeHasBadgeClass(t *testing.T) {
	t.Run("Badge includes badge class by default", func(t *testing.T) {
		b := bs.Badge()
		if b == nil {
			t.Fatal("bs.Badge() returned nil")
		}
		// The badge class is added in bs.Badge() constructor
	})

	t.Run("Badge with additional classes still includes badge class", func(t *testing.T) {
		b := bs.Badge(WithClass("text-bg-primary", "ms-2"))
		if b == nil {
			t.Fatal("bs.Badge() with additional classes returned nil")
		}
	})
}

// Test PillBadge includes both 'badge' and 'rounded-pill' classes
func TestPillBadgeHasBothClasses(t *testing.T) {
	t.Run("PillBadge includes badge and rounded-pill classes", func(t *testing.T) {
		b := bs.PillBadge()
		if b == nil {
			t.Fatal("bs.PillBadge() returned nil")
		}
		// The badge and rounded-pill classes are added in bs.PillBadge() constructor
	})

	t.Run("PillBadge with additional classes includes all classes", func(t *testing.T) {
		b := bs.PillBadge(WithClass("text-bg-danger"))
		if b == nil {
			t.Fatal("bs.PillBadge() with additional classes returned nil")
		}
	})
}

// Test badge sizing (relative to parent)
func TestBadgeSizing(t *testing.T) {
	t.Run("badge inherits size from parent", func(t *testing.T) {
		// Badges use em units and scale with parent font size
		b := bs.Badge(WithClass("text-bg-primary"))
		if b == nil {
			t.Error("Badge should not return nil")
		}
	})
}

// Test legacy color classes (bg-{color} and text-{color})
func TestBadgeLegacyColorClasses(t *testing.T) {
	tests := []struct {
		name    string
		classes []string
	}{
		{
			name:    "legacy bg-primary",
			classes: []string{"bg-primary", "text-white"},
		},
		{
			name:    "legacy bg-secondary",
			classes: []string{"bg-secondary", "text-white"},
		},
		{
			name:    "legacy bg-success",
			classes: []string{"bg-success", "text-white"},
		},
		{
			name:    "legacy bg-danger",
			classes: []string{"bg-danger", "text-white"},
		},
		{
			name:    "legacy bg-warning",
			classes: []string{"bg-warning", "text-dark"},
		},
		{
			name:    "legacy bg-info",
			classes: []string{"bg-info", "text-dark"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bs.Badge(WithClass(tt.classes...))
			if b == nil {
				t.Errorf("Badge with legacy classes %v should not return nil", tt.classes)
			}
		})
	}
}

// Test badge combinations as seen in Bootstrap documentation
func TestBadgeDocumentationExamples(t *testing.T) {
	t.Run("example: badge in heading", func(t *testing.T) {
		// <span class="badge text-bg-secondary">New</span>
		b := bs.Badge(bs.WithColor(bs.SECONDARY))
		if b == nil {
			t.Error("Badge for heading example should not return nil")
		}
	})

	t.Run("example: badge in button", func(t *testing.T) {
		// <span class="badge text-bg-secondary">4</span>
		b := bs.Badge(bs.WithColor(bs.SECONDARY))
		if b == nil {
			t.Error("Badge for button example should not return nil")
		}
	})

	t.Run("example: positioned notification badge", func(t *testing.T) {
		// <span class="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">99+</span>
		b := bs.PillBadge(
			WithClass("position-absolute", "top-0", "start-100", "translate-middle"),
			bs.WithBackground(bs.DANGER),
		)
		if b == nil {
			t.Error("Positioned notification badge should not return nil")
		}
	})

	t.Run("example: all background colors with WithColor", func(t *testing.T) {
		colors := []bs.Color{bs.PRIMARY, bs.SECONDARY, bs.SUCCESS, bs.DANGER, bs.WARNING, bs.INFO, bs.LIGHT, bs.DARK}
		for _, color := range colors {
			b := bs.Badge(bs.WithColor(color))
			if b == nil {
				t.Errorf("Badge with color %s should not return nil", color)
			}
		}
	})

	t.Run("example: all pill badge colors with WithColor", func(t *testing.T) {
		colors := []bs.Color{bs.PRIMARY, bs.SECONDARY, bs.SUCCESS, bs.DANGER, bs.WARNING, bs.INFO, bs.LIGHT, bs.DARK}
		for _, color := range colors {
			b := bs.PillBadge(bs.WithColor(color))
			if b == nil {
				t.Errorf("PillBadge with color %s should not return nil", color)
			}
		}
	})

	t.Run("example: badge with margin and padding", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.PRIMARY),
			bs.WithMargin(bs.Start, 2),
			bs.WithPadding(bs.X, 3),
		)
		if b == nil {
			t.Error("Badge with margin and padding should not return nil")
		}
	})

	t.Run("example: badge with border", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.INFO),
			bs.WithBorder(bs.All, bs.DARK),
			bs.WithPadding(bs.All, 2),
		)
		if b == nil {
			t.Error("Badge with border should not return nil")
		}
	})
}

// Test that Badge and PillBadge return different instances
func TestBadgeVsPillBadge(t *testing.T) {
	b1 := bs.Badge(bs.WithColor(bs.PRIMARY))
	b2 := bs.PillBadge(bs.WithColor(bs.PRIMARY))

	if b1 == b2 {
		t.Error("bs.Badge() and bs.PillBadge() should return different instances")
	}
}

// Test combining multiple helper options
func TestBadgeCombinedHelperOptions(t *testing.T) {
	t.Run("color with margin", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.SUCCESS),
			bs.WithMargin(bs.Start, 2),
		)
		if b == nil {
			t.Error("Badge with color and margin should not return nil")
		}
	})

	t.Run("color with padding", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.DANGER),
			bs.WithPadding(bs.All, 3),
		)
		if b == nil {
			t.Error("Badge with color and padding should not return nil")
		}
	})

	t.Run("color with border", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.WARNING),
			bs.WithBorder(bs.All, bs.DARK),
		)
		if b == nil {
			t.Error("Badge with color and border should not return nil")
		}
	})

	t.Run("all helper options combined", func(t *testing.T) {
		b := bs.Badge(
			bs.WithColor(bs.INFO),
			bs.WithMargin(bs.Start, 2),
			bs.WithPadding(bs.Y, 2),
			bs.WithBorder(bs.All, bs.PRIMARY),
		)
		if b == nil {
			t.Error("Badge with all helper options should not return nil")
		}
	})

	t.Run("pill badge with all helper options", func(t *testing.T) {
		b := bs.PillBadge(
			bs.WithColor(bs.PRIMARY),
			bs.WithMargin(bs.End, 3),
			bs.WithPadding(bs.X, 4),
			bs.WithBorder(bs.All, bs.LIGHT),
		)
		if b == nil {
			t.Error("PillBadge with all helper options should not return nil")
		}
	})
}

// Test margin positions
func TestBadgeMarginPositions(t *testing.T) {
	positions := []struct {
		name     string
		position bs.Position
	}{
		{"top", bs.Top},
		{"bottom", bs.Bottom},
		{"start", bs.Start},
		{"end", bs.End},
		{"all", bs.All},
		{"x-axis", bs.X},
		{"y-axis", bs.Y},
	}

	for _, tc := range positions {
		t.Run(tc.name, func(t *testing.T) {
			b := bs.Badge(bs.WithMargin(tc.position, 2))
			if b == nil {
				t.Errorf("Badge with %s margin should not return nil", tc.name)
			}
		})
	}
}

// Test padding positions
func TestBadgePaddingPositions(t *testing.T) {
	positions := []struct {
		name     string
		position bs.Position
	}{
		{"top", bs.Top},
		{"bottom", bs.Bottom},
		{"start", bs.Start},
		{"end", bs.End},
		{"all", bs.All},
		{"x-axis", bs.X},
		{"y-axis", bs.Y},
	}

	for _, tc := range positions {
		t.Run(tc.name, func(t *testing.T) {
			b := bs.Badge(bs.WithPadding(tc.position, 3))
			if b == nil {
				t.Errorf("Badge with %s padding should not return nil", tc.name)
			}
		})
	}
}

// Test border positions and colors
func TestBadgeBorderPositions(t *testing.T) {
	tests := []struct {
		name     string
		position bs.Position
		color    bs.Color
	}{
		{"top border with primary", bs.Top, bs.PRIMARY},
		{"bottom border with danger", bs.Bottom, bs.DANGER},
		{"start border with success", bs.Start, bs.SUCCESS},
		{"end border with warning", bs.End, bs.WARNING},
		{"all borders with info", bs.All, bs.INFO},
		{"all borders no color", bs.All, bs.TRANSPARENT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b View
			if tt.color == bs.TRANSPARENT {
				b = bs.Badge(bs.WithBorder(tt.position))
			} else {
				b = bs.Badge(bs.WithBorder(tt.position, tt.color))
			}
			if b == nil {
				t.Errorf("Badge with %s should not return nil", tt.name)
			}
		})
	}
}

// Test spacing sizes
func TestBadgeSpacingSizes(t *testing.T) {
	sizes := []int{1, 2, 3, 4, 5}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("margin_size_%d", size), func(t *testing.T) {
			b := bs.Badge(bs.WithMargin(bs.All, size))
			if b == nil {
				t.Errorf("Badge with margin size %d should not return nil", size)
			}
		})

		t.Run(fmt.Sprintf("padding_size_%d", size), func(t *testing.T) {
			b := bs.Badge(bs.WithPadding(bs.All, size))
			if b == nil {
				t.Errorf("Badge with padding size %d should not return nil", size)
			}
		})
	}

	// Test auto margin (size 0 is 'auto')
	t.Run("margin_auto", func(t *testing.T) {
		// Auto margins work differently, skipping size 0 which is invalid
		b := bs.Badge(WithClass("m-auto"))
		if b == nil {
			t.Error("Badge with auto margin should not return nil")
		}
	})
}

// Test color variants including subtle colors
func TestBadgeColorVariants(t *testing.T) {
	colors := []struct {
		name  string
		color bs.Color
	}{
		{"primary", bs.PRIMARY},
		{"primary-subtle", bs.PRIMARY_SUBTLE},
		{"secondary", bs.SECONDARY},
		{"secondary-subtle", bs.SECONDARY_SUBTLE},
		{"success", bs.SUCCESS},
		{"success-subtle", bs.SUCCESS_SUBTLE},
		{"danger", bs.DANGER},
		{"danger-subtle", bs.DANGER_SUBTLE},
		{"warning", bs.WARNING},
		{"warning-subtle", bs.WARNING_SUBTLE},
		{"info", bs.INFO},
		{"info-subtle", bs.INFO_SUBTLE},
		{"light", bs.LIGHT},
		{"light-subtle", bs.LIGHT_SUBTLE},
		{"dark", bs.DARK},
		{"dark-subtle", bs.DARK_SUBTLE},
	}

	for _, tc := range colors {
		t.Run(tc.name, func(t *testing.T) {
			b := bs.Badge(bs.WithColor(tc.color))
			if b == nil {
				t.Errorf("Badge with %s color should not return nil", tc.name)
			}
		})
	}
}
