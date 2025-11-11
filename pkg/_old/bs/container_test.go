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
// BASIC CONTAINER TESTS

func TestContainerCreation(t *testing.T) {
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
			opts: []Opt{WithID("main-container")},
		},
		{
			name: "create with classes",
			opts: []Opt{WithClass("shadow", "rounded")},
		},
		{
			name: "create with combined options",
			opts: []Opt{
				WithID("content-wrapper"),
				WithClass("mt-5"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			if c == nil {
				t.Fatal("bs.Container() returned nil")
			}
			if c.Name() != bs.ViewContainer {
				t.Errorf("bs.Container().Name() = %v, want %v", c.Name(), bs.ViewContainer)
			}
		})
	}
}

func TestContainerViewInterface(t *testing.T) {
	c := bs.Container()

	t.Run("Name returns correct view name", func(t *testing.T) {
		if c.Name() != bs.ViewContainer {
			t.Errorf("Container.Name() = %v, want %v", c.Name(), bs.ViewContainer)
		}
	})

	t.Run("Root returns non-nil element", func(t *testing.T) {
		root := c.Root()
		if root == nil {
			t.Error("Container.Root() returned nil")
		}
	})

	t.Run("ID method should not panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Container.ID() panicked: %v", r)
			}
		}()
		_ = c.ID()
	})

	t.Run("Root has correct tag name", func(t *testing.T) {
		root := c.Root()
		if root.TagName() != "DIV" {
			t.Errorf("Container.Root().TagName() = %v, want DIV", root.TagName())
		}
	})
}

func TestContainerDefaultClass(t *testing.T) {
	c := bs.Container()
	classList := c.Root().ClassList()

	t.Run("has container class by default", func(t *testing.T) {
		if !classList.Contains("container") {
			t.Error("Container should have 'container' class by default")
		}
	})
}

func TestContainerEmbedding(t *testing.T) {
	// Verify that container properly embeds View
	c := bs.Container()

	// container should be able to be used as a View
	var v View = c
	if v.Name() != bs.ViewContainer {
		t.Errorf("Embedded View.Name() = %v, want %v", v.Name(), bs.ViewContainer)
	}
}

func TestContainerMultipleInstances(t *testing.T) {
	// Create multiple containers and verify they're independent
	c1 := bs.Container(WithID("container-1"))
	c2 := bs.Container(WithID("container-2"))

	if c1 == c2 {
		t.Error("Each bs.Container() call should create a distinct instance")
	}

	if c1.ID() == c2.ID() {
		t.Error("Containers with different IDs should maintain separate state")
	}
}

func TestContainerWithID(t *testing.T) {
	c := bs.Container(WithID("app-container"))
	if c.ID() != "app-container" {
		t.Errorf("bs.Container(WithID('app-container')).ID() = %v, want %v", c.ID(), "app-container")
	}
}

func TestContainerWithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		opts            []Opt
		expectedClass   string
		mustHaveClasses []string
	}{
		{
			name:            "container with single additional class",
			opts:            []Opt{WithClass("shadow")},
			expectedClass:   "shadow",
			mustHaveClasses: []string{"container", "shadow"},
		},
		{
			name:            "container with multiple additional classes",
			opts:            []Opt{WithClass("shadow", "rounded", "p-3")},
			expectedClass:   "shadow",
			mustHaveClasses: []string{"container", "shadow", "rounded", "p-3"},
		},
		{
			name:            "container with utility classes",
			opts:            []Opt{WithClass("mt-5", "mb-3", "mx-auto")},
			expectedClass:   "mt-5",
			mustHaveClasses: []string{"container", "mt-5", "mb-3", "mx-auto"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.mustHaveClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have class %q", class)
				}
			}
		})
	}
}

func TestContainerEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		opts []Opt
	}{
		{
			name: "empty options",
			opts: []Opt{},
		},
		{
			name: "duplicate classes",
			opts: []Opt{WithClass("shadow", "shadow")},
		},
		{
			name: "empty ID",
			opts: []Opt{WithID("")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Container with %q should not panic: %v", tt.name, r)
				}
			}()

			c := bs.Container(tt.opts...)
			if c == nil {
				t.Errorf("Container with %q should not return nil", tt.name)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// RESPONSIVE BREAKPOINT TESTS (WithSize)

func TestContainerResponsiveBreakpoints(t *testing.T) {
	tests := []struct {
		name          string
		size          bs.Size
		useSize       bool
		expectedClass string
		mustNotHave   []string
	}{
		{
			name:          "default container",
			size:          bs.SizeDefault,
			useSize:       false, // Don't pass WithSize for default
			expectedClass: "container",
			mustNotHave:   []string{"container-sm", "container-md", "container-lg", "container-xl", "container-xxl", "container-fluid"},
		},
		{
			name:          "small breakpoint",
			size:          bs.SizeSmall,
			useSize:       true,
			expectedClass: "container-sm",
			mustNotHave:   []string{"container-md", "container-lg", "container-xl", "container-xxl", "container-fluid"},
		},
		{
			name:          "medium breakpoint",
			size:          bs.SizeMedium,
			useSize:       true,
			expectedClass: "container-md",
			mustNotHave:   []string{"container-sm", "container-lg", "container-xl", "container-xxl", "container-fluid"},
		},
		{
			name:          "large breakpoint",
			size:          bs.SizeLarge,
			useSize:       true,
			expectedClass: "container-lg",
			mustNotHave:   []string{"container-sm", "container-md", "container-xl", "container-xxl", "container-fluid"},
		},
		{
			name:          "extra large breakpoint",
			size:          bs.SizeXLarge,
			useSize:       true,
			expectedClass: "container-xl",
			mustNotHave:   []string{"container-sm", "container-md", "container-lg", "container-xxl", "container-fluid"},
		},
		{
			name:          "extra extra large breakpoint",
			size:          bs.SizeXXLarge,
			useSize:       true,
			expectedClass: "container-xxl",
			mustNotHave:   []string{"container-sm", "container-md", "container-lg", "container-xl", "container-fluid"},
		},
		{
			name:          "fluid container",
			size:          bs.SizeFluid,
			useSize:       true,
			expectedClass: "container-fluid",
			mustNotHave:   []string{"container-sm", "container-md", "container-lg", "container-xl", "container-xxl"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var classList interface {
				Contains(string) bool
			}
			if tt.useSize {
				c := bs.Container(bs.WithSize(tt.size))
				classList = c.Root().ClassList()
			} else {
				c := bs.Container()
				classList = c.Root().ClassList()
			}

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Container should have class %q", tt.expectedClass)
			}

			for _, class := range tt.mustNotHave {
				if classList.Contains(class) {
					t.Errorf("Container should not have class %q", class)
				}
			}
		})
	}
}

func TestContainerBreakpointWithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		size            bs.Size
		opts            []Opt
		expectedSize    string
		additionalClass string
	}{
		{
			name:            "small with margin",
			size:            bs.SizeSmall,
			opts:            []Opt{bs.WithSize(bs.SizeSmall), WithClass("mt-3")},
			expectedSize:    "container-sm",
			additionalClass: "mt-3",
		},
		{
			name:            "medium with padding",
			size:            bs.SizeMedium,
			opts:            []Opt{bs.WithSize(bs.SizeMedium), WithClass("p-4")},
			expectedSize:    "container-md",
			additionalClass: "p-4",
		},
		{
			name:            "large with multiple utilities",
			size:            bs.SizeLarge,
			opts:            []Opt{bs.WithSize(bs.SizeLarge), WithClass("shadow", "rounded")},
			expectedSize:    "container-lg",
			additionalClass: "shadow",
		},
		{
			name:            "extra large with custom class",
			size:            bs.SizeXLarge,
			opts:            []Opt{bs.WithSize(bs.SizeXLarge), WithClass("my-custom-class")},
			expectedSize:    "container-xl",
			additionalClass: "my-custom-class",
		},
		{
			name:            "extra extra large with text-center",
			size:            bs.SizeXXLarge,
			opts:            []Opt{bs.WithSize(bs.SizeXXLarge), WithClass("text-center")},
			expectedSize:    "container-xxl",
			additionalClass: "text-center",
		},
		{
			name:            "fluid with margin",
			size:            bs.SizeFluid,
			opts:            []Opt{bs.WithSize(bs.SizeFluid), WithClass("mx-auto")},
			expectedSize:    "container-fluid",
			additionalClass: "mx-auto",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			classList := c.Root().ClassList()

			if !classList.Contains(tt.expectedSize) {
				t.Errorf("Container should have class %q", tt.expectedSize)
			}

			if !classList.Contains(tt.additionalClass) {
				t.Errorf("Container should have class %q", tt.additionalClass)
			}
		})
	}
}

func TestContainerBreakpointReplacement(t *testing.T) {
	tests := []struct {
		name        string
		firstSize   bs.Size
		secondSize  bs.Size
		finalClass  string
		mustNotHave string
	}{
		{
			name:        "large overrides medium",
			firstSize:   bs.SizeMedium,
			secondSize:  bs.SizeLarge,
			finalClass:  "container-lg",
			mustNotHave: "container-md",
		},
		{
			name:        "fluid overrides large",
			firstSize:   bs.SizeLarge,
			secondSize:  bs.SizeFluid,
			finalClass:  "container-fluid",
			mustNotHave: "container-lg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// First create with firstSize, then apply secondSize
			c := bs.Container(bs.WithSize(tt.firstSize))
			c.Opts(bs.WithSize(tt.secondSize))
			classList := c.Root().ClassList()

			if !classList.Contains(tt.finalClass) {
				t.Errorf("Container should have class %q", tt.finalClass)
			}

			if classList.Contains(tt.mustNotHave) {
				t.Errorf("Container should not have class %q after replacement", tt.mustNotHave)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// MARGIN TESTS

func TestContainerWithMargin(t *testing.T) {
	tests := []struct {
		name          string
		position      bs.Position
		size          int
		expectedClass string
	}{
		{
			name:          "margin all sides size 3",
			position:      bs.All,
			size:          3,
			expectedClass: "m-3",
		},
		{
			name:          "margin top size 5",
			position:      bs.Top,
			size:          5,
			expectedClass: "mt-5",
		},
		{
			name:          "margin bottom size 2",
			position:      bs.Bottom,
			size:          2,
			expectedClass: "mb-2",
		},
		{
			name:          "margin start size 4",
			position:      bs.Start,
			size:          4,
			expectedClass: "ms-4",
		},
		{
			name:          "margin end size 1",
			position:      bs.End,
			size:          1,
			expectedClass: "me-1",
		},
		{
			name:          "margin horizontal (x) size 3",
			position:      bs.X,
			size:          3,
			expectedClass: "mx-3",
		},
		{
			name:          "margin vertical (y) size 4",
			position:      bs.Y,
			size:          4,
			expectedClass: "my-4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(bs.WithMargin(tt.position, tt.size))
			classList := c.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Container should have margin class %q", tt.expectedClass)
			}
		})
	}
}

func TestContainerMarginReplacement(t *testing.T) {
	// Test that margin utilities replace previous values
	c := bs.Container(bs.WithMargin(bs.All, 2))
	classList := c.Root().ClassList()

	if !classList.Contains("m-2") {
		t.Error("Container should initially have m-2")
	}

	// Apply new margin - should replace previous
	c.Opts(bs.WithMargin(bs.All, 4))
	classList = c.Root().ClassList()

	if !classList.Contains("m-4") {
		t.Error("Container should have m-4 after replacement")
	}

	if classList.Contains("m-2") {
		t.Error("Container should not have m-2 after replacement")
	}
}

func TestContainerMarginCombinations(t *testing.T) {
	tests := []struct {
		name    string
		margins []struct {
			pos  bs.Position
			size int
		}
		expectedClasses []string
	}{
		{
			name: "top and bottom margins",
			margins: []struct {
				pos  bs.Position
				size int
			}{
				{bs.Top, 3},
				{bs.Bottom, 2},
			},
			expectedClasses: []string{"mt-3", "mb-2"},
		},
		{
			name: "start and end margins",
			margins: []struct {
				pos  bs.Position
				size int
			}{
				{bs.Start, 4},
				{bs.End, 1},
			},
			expectedClasses: []string{"ms-4", "me-1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := make([]Opt, 0, len(tt.margins))
			for _, m := range tt.margins {
				opts = append(opts, bs.WithMargin(m.pos, m.size))
			}

			c := bs.Container(opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have margin class %q", class)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// PADDING TESTS

func TestContainerWithPadding(t *testing.T) {
	tests := []struct {
		name          string
		position      bs.Position
		size          int
		expectedClass string
	}{
		{
			name:          "padding all sides size 3",
			position:      bs.All,
			size:          3,
			expectedClass: "p-3",
		},
		{
			name:          "padding top size 5",
			position:      bs.Top,
			size:          5,
			expectedClass: "pt-5",
		},
		{
			name:          "padding bottom size 2",
			position:      bs.Bottom,
			size:          2,
			expectedClass: "pb-2",
		},
		{
			name:          "padding start size 4",
			position:      bs.Start,
			size:          4,
			expectedClass: "ps-4",
		},
		{
			name:          "padding end size 1",
			position:      bs.End,
			size:          1,
			expectedClass: "pe-1",
		},
		{
			name:          "padding horizontal (x) size 3",
			position:      bs.X,
			size:          3,
			expectedClass: "px-3",
		},
		{
			name:          "padding vertical (y) size 4",
			position:      bs.Y,
			size:          4,
			expectedClass: "py-4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(bs.WithPadding(tt.position, tt.size))
			classList := c.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Container should have padding class %q", tt.expectedClass)
			}
		})
	}
}

func TestContainerPaddingReplacement(t *testing.T) {
	// Test that padding utilities replace previous values
	c := bs.Container(bs.WithPadding(bs.All, 2))
	classList := c.Root().ClassList()

	if !classList.Contains("p-2") {
		t.Error("Container should initially have p-2")
	}

	// Apply new padding - should replace previous
	c.Opts(bs.WithPadding(bs.All, 5))
	classList = c.Root().ClassList()

	if !classList.Contains("p-5") {
		t.Error("Container should have p-5 after replacement")
	}

	if classList.Contains("p-2") {
		t.Error("Container should not have p-2 after replacement")
	}
}

func TestContainerPaddingCombinations(t *testing.T) {
	tests := []struct {
		name     string
		paddings []struct {
			pos  bs.Position
			size int
		}
		expectedClasses []string
	}{
		{
			name: "top and bottom padding",
			paddings: []struct {
				pos  bs.Position
				size int
			}{
				{bs.Top, 3},
				{bs.Bottom, 2},
			},
			expectedClasses: []string{"pt-3", "pb-2"},
		},
		{
			name: "start and end padding",
			paddings: []struct {
				pos  bs.Position
				size int
			}{
				{bs.Start, 4},
				{bs.End, 1},
			},
			expectedClasses: []string{"ps-4", "pe-1"},
		},
		{
			name: "vertical and horizontal",
			paddings: []struct {
				pos  bs.Position
				size int
			}{
				{bs.Y, 3},
				{bs.X, 4},
			},
			expectedClasses: []string{"py-3", "px-4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := make([]Opt, 0, len(tt.paddings))
			for _, p := range tt.paddings {
				opts = append(opts, bs.WithPadding(p.pos, p.size))
			}

			c := bs.Container(opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have padding class %q", class)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BORDER TESTS

func TestContainerWithBorder(t *testing.T) {
	tests := []struct {
		name          string
		position      bs.Position
		color         bs.Color
		expectedClass string
	}{
		{
			name:          "border on all sides",
			position:      bs.All,
			color:         bs.TRANSPARENT,
			expectedClass: "border",
		},
		{
			name:          "border top only",
			position:      bs.Top,
			color:         bs.TRANSPARENT,
			expectedClass: "border-top",
		},
		{
			name:          "border bottom only",
			position:      bs.Bottom,
			color:         bs.TRANSPARENT,
			expectedClass: "border-bottom",
		},
		{
			name:          "border start only",
			position:      bs.Start,
			color:         bs.TRANSPARENT,
			expectedClass: "border-start",
		},
		{
			name:          "border end only",
			position:      bs.End,
			color:         bs.TRANSPARENT,
			expectedClass: "border-end",
		},
		{
			name:          "border all with primary color",
			position:      bs.All,
			color:         bs.PRIMARY,
			expectedClass: "border-primary",
		},
		{
			name:          "border top with danger color",
			position:      bs.Top,
			color:         bs.DANGER,
			expectedClass: "border-danger",
		},
		{
			name:          "border all with success color",
			position:      bs.All,
			color:         bs.SUCCESS,
			expectedClass: "border-success",
		},
		{
			name:          "border bottom with warning color",
			position:      bs.Bottom,
			color:         bs.WARNING,
			expectedClass: "border-warning",
		},
		{
			name:          "border with subtle color",
			position:      bs.All,
			color:         bs.PRIMARY_SUBTLE,
			expectedClass: "border-primary-subtle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(bs.WithBorder(tt.position, tt.color))
			classList := c.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Container should have border class %q", tt.expectedClass)
			}
		})
	}
}

func TestContainerBorderCombinations(t *testing.T) {
	tests := []struct {
		name    string
		borders []struct {
			pos   bs.Position
			color bs.Color
		}
		expectedClasses []string
	}{
		{
			name: "top and bottom borders",
			borders: []struct {
				pos   bs.Position
				color bs.Color
			}{
				{bs.Top, bs.TRANSPARENT},
				{bs.Bottom, bs.TRANSPARENT},
			},
			expectedClasses: []string{"border-top", "border-bottom"},
		},
		{
			name: "start and end borders",
			borders: []struct {
				pos   bs.Position
				color bs.Color
			}{
				{bs.Start, bs.TRANSPARENT},
				{bs.End, bs.TRANSPARENT},
			},
			expectedClasses: []string{"border-start", "border-end"},
		},
		{
			name: "all sides with primary color",
			borders: []struct {
				pos   bs.Position
				color bs.Color
			}{
				{bs.All, bs.TRANSPARENT},
				{bs.All, bs.PRIMARY},
			},
			expectedClasses: []string{"border", "border-primary"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := make([]Opt, 0, len(tt.borders))
			for _, b := range tt.borders {
				opts = append(opts, bs.WithBorder(b.pos, b.color))
			}

			c := bs.Container(opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have border class %q", class)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// BACKGROUND COLOR TESTS

func TestContainerWithBackground(t *testing.T) {
	tests := []struct {
		name          string
		color         bs.Color
		expectedClass string
	}{
		{
			name:          "primary background",
			color:         bs.PRIMARY,
			expectedClass: "bg-primary",
		},
		{
			name:          "secondary background",
			color:         bs.SECONDARY,
			expectedClass: "bg-secondary",
		},
		{
			name:          "success background",
			color:         bs.SUCCESS,
			expectedClass: "bg-success",
		},
		{
			name:          "danger background",
			color:         bs.DANGER,
			expectedClass: "bg-danger",
		},
		{
			name:          "warning background",
			color:         bs.WARNING,
			expectedClass: "bg-warning",
		},
		{
			name:          "info background",
			color:         bs.INFO,
			expectedClass: "bg-info",
		},
		{
			name:          "light background",
			color:         bs.LIGHT,
			expectedClass: "bg-light",
		},
		{
			name:          "dark background",
			color:         bs.DARK,
			expectedClass: "bg-dark",
		},
		{
			name:          "white background",
			color:         bs.WHITE,
			expectedClass: "bg-white",
		},
		{
			name:          "black background",
			color:         bs.BLACK,
			expectedClass: "bg-black",
		},
		{
			name:          "primary subtle background",
			color:         bs.PRIMARY_SUBTLE,
			expectedClass: "bg-primary-subtle",
		},
		{
			name:          "danger subtle background",
			color:         bs.DANGER_SUBTLE,
			expectedClass: "bg-danger-subtle",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(bs.WithBackground(tt.color))
			classList := c.Root().ClassList()

			if !classList.Contains(tt.expectedClass) {
				t.Errorf("Container should have background class %q", tt.expectedClass)
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// COMPLEX COMBINATIONS

func TestContainerComplexCombinations(t *testing.T) {
	tests := []struct {
		name            string
		opts            []Opt
		expectedClasses []string
	}{
		{
			name: "container with size, margin, padding",
			opts: []Opt{
				bs.WithSize(bs.SizeLarge),
				bs.WithMargin(bs.Y, 5),
				bs.WithPadding(bs.All, 4),
			},
			expectedClasses: []string{"container-lg", "my-5", "p-4"},
		},
		{
			name: "fluid container with border and background",
			opts: []Opt{
				bs.WithSize(bs.SizeFluid),
				bs.WithBorder(bs.All, bs.PRIMARY),
				bs.WithBackground(bs.LIGHT),
			},
			expectedClasses: []string{"container-fluid", "border", "border-primary", "bg-light"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have class %q", class)
				}
			}
		})
	}
}

func TestContainerWithBorderAndBackground(t *testing.T) {
	tests := []struct {
		name            string
		opts            []Opt
		expectedClasses []string
	}{
		{
			name: "border with background",
			opts: []Opt{
				bs.WithBorder(bs.All, bs.PRIMARY),
				bs.WithBackground(bs.PRIMARY_SUBTLE),
			},
			expectedClasses: []string{"border", "border-primary", "bg-primary-subtle"},
		},
		{
			name: "top border with danger background",
			opts: []Opt{
				bs.WithBorder(bs.Top, bs.DANGER),
				bs.WithBackground(bs.DANGER_SUBTLE),
			},
			expectedClasses: []string{"border-top", "border-danger", "bg-danger-subtle"},
		},
		{
			name: "all borders with success colors",
			opts: []Opt{
				bs.WithBorder(bs.All, bs.SUCCESS),
				bs.WithBackground(bs.SUCCESS_SUBTLE),
			},
			expectedClasses: []string{"border", "border-success", "bg-success-subtle"},
		},
		{
			name: "multiple borders with background",
			opts: []Opt{
				bs.WithBorder(bs.Top, bs.PRIMARY),
				bs.WithBorder(bs.Bottom, bs.PRIMARY),
				bs.WithBackground(bs.LIGHT),
			},
			expectedClasses: []string{"border-top", "border-bottom", "border-primary", "bg-light"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have class %q", class)
				}
			}
		})
	}
}

func TestContainerWithBorderAndBackgroundWithBreakpoint(t *testing.T) {
	tests := []struct {
		name            string
		opts            []Opt
		expectedClasses []string
	}{
		{
			name: "fluid container with border and background",
			opts: []Opt{
				bs.WithSize(bs.SizeFluid),
				bs.WithBorder(bs.All, bs.TRANSPARENT),
				bs.WithBackground(bs.PRIMARY),
			},
			expectedClasses: []string{"container-fluid", "border", "bg-primary"},
		},
		{
			name: "small container with top border and background",
			opts: []Opt{
				bs.WithSize(bs.SizeSmall),
				bs.WithBorder(bs.Top, bs.DANGER),
				bs.WithBackground(bs.LIGHT),
			},
			expectedClasses: []string{"container-sm", "border-top", "border-danger", "bg-light"},
		},
		{
			name: "large container with all sides border and dark background",
			opts: []Opt{
				bs.WithSize(bs.SizeLarge),
				bs.WithBorder(bs.All, bs.LIGHT),
				bs.WithBackground(bs.DARK),
			},
			expectedClasses: []string{"container-lg", "border", "border-light", "bg-dark"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			classList := c.Root().ClassList()

			for _, class := range tt.expectedClasses {
				if !classList.Contains(class) {
					t.Errorf("Container should have class %q", class)
				}
			}
		})
	}
}

///////////////////////////////////////////////////////////////////////////////
// ADDITIONAL UTILITY TESTS

func TestContainerWithAttributes(t *testing.T) {
	c := bs.Container(
		WithAttr("role", "main"),
		WithAttr("data-section", "content"),
	)

	root := c.Root()
	if root.GetAttribute("role") != "main" {
		t.Errorf("Container should have role='main', got %q", root.GetAttribute("role"))
	}

	if root.GetAttribute("data-section") != "content" {
		t.Errorf("Container should have data-section='content', got %q", root.GetAttribute("data-section"))
	}
}

func TestContainerOuterHTML(t *testing.T) {
	tests := []struct {
		name     string
		opts     []Opt
		contains []string
	}{
		{
			name:     "default container",
			opts:     nil,
			contains: []string{"<div", "class=\"container\"", "</div>"},
		},
		{
			name:     "container with ID",
			opts:     []Opt{WithID("main")},
			contains: []string{"id=\"main\"", "class=\"container\""},
		},
		{
			name:     "fluid container",
			opts:     []Opt{bs.WithSize(bs.SizeFluid)},
			contains: []string{"class=\"container-fluid\""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := bs.Container(tt.opts...)
			html := c.Root().OuterHTML()

			for _, substr := range tt.contains {
				if !strings.Contains(html, substr) {
					t.Errorf("OuterHTML should contain %q, got: %s", substr, html)
				}
			}
		})
	}
}
