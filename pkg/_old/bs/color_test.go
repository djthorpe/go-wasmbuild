package bs

import (
	"testing"
)

func TestColorConstants(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"TRANSPARENT", TRANSPARENT, ""},
		{"PRIMARY", PRIMARY, "primary"},
		{"PRIMARY_SUBTLE", PRIMARY_SUBTLE, "primary-subtle"},
		{"SECONDARY", SECONDARY, "secondary"},
		{"SECONDARY_SUBTLE", SECONDARY_SUBTLE, "secondary-subtle"},
		{"SUCCESS", SUCCESS, "success"},
		{"SUCCESS_SUBTLE", SUCCESS_SUBTLE, "success-subtle"},
		{"DANGER", DANGER, "danger"},
		{"DANGER_SUBTLE", DANGER_SUBTLE, "danger-subtle"},
		{"WARNING", WARNING, "warning"},
		{"WARNING_SUBTLE", WARNING_SUBTLE, "warning-subtle"},
		{"INFO", INFO, "info"},
		{"INFO_SUBTLE", INFO_SUBTLE, "info-subtle"},
		{"LIGHT", LIGHT, "light"},
		{"LIGHT_SUBTLE", LIGHT_SUBTLE, "light-subtle"},
		{"DARK", DARK, "dark"},
		{"DARK_SUBTLE", DARK_SUBTLE, "dark-subtle"},
		{"WHITE", WHITE, "white"},
		{"BLACK", BLACK, "black"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.color) != tt.expected {
				t.Errorf("%s = %v, want %v", tt.name, string(tt.color), tt.expected)
			}
		})
	}
}

func TestColorType(t *testing.T) {
	// Verify Color is a string type
	var c Color = "test"
	if _, ok := any(c).(string); ok {
		t.Error("Color should not be directly castable to string without explicit conversion")
	}

	// But string conversion should work
	if string(c) != "test" {
		t.Errorf("string(Color) = %v, want %v", string(c), "test")
	}
}

func TestColorClassName(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		prefix   string
		expected string
	}{
		{
			name:     "transparent with text prefix",
			color:    TRANSPARENT,
			prefix:   "text",
			expected: "text",
		},
		{
			name:     "primary with text prefix",
			color:    PRIMARY,
			prefix:   "text",
			expected: "text-primary",
		},
		{
			name:     "danger with bg prefix",
			color:    DANGER,
			prefix:   "bg",
			expected: "bg-danger",
		},
		{
			name:     "success-subtle with text prefix",
			color:    SUCCESS_SUBTLE,
			prefix:   "text",
			expected: "text-success-subtle",
		},
		{
			name:     "white with border prefix",
			color:    WHITE,
			prefix:   "border",
			expected: "border-white",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.color.className(tt.prefix)
			if result != tt.expected {
				t.Errorf("className(%q) = %v, want %v", tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestColorAllClassNames(t *testing.T) {
	t.Run("returns all color classes with prefix", func(t *testing.T) {
		result := PRIMARY.allClassNames("text")

		// Should return 19 classes (all colors)
		if len(result) != len(allColors) {
			t.Errorf("allClassNames() returned %d classes, want %d", len(result), len(allColors))
		}

		// Check some expected values
		expectedClasses := []string{
			"text",
			"text-primary",
			"text-secondary",
			"text-danger",
		}

		for _, expected := range expectedClasses {
			found := false
			for _, class := range result {
				if class == expected {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("allClassNames() missing expected class %q", expected)
			}
		}
	})

	t.Run("different colors return same list", func(t *testing.T) {
		result1 := PRIMARY.allClassNames("bg")
		result2 := DANGER.allClassNames("bg")

		if len(result1) != len(result2) {
			t.Errorf("Different colors should return same length list")
		}
	})
}

func TestColorPrefixForView(t *testing.T) {
	tests := []struct {
		viewName string
		expected string
	}{
		{ViewHeading, "text"},
		{ViewText, "text"},
		{ViewContainer, "text"},
		{"unknown-view", ""},
	}

	for _, tt := range tests {
		t.Run(tt.viewName, func(t *testing.T) {
			result := colorPrefixForView(tt.viewName)
			if result != tt.expected {
				t.Errorf("colorPrefixForView(%q) = %v, want %v", tt.viewName, result, tt.expected)
			}
		})
	}
}

func TestWithColor(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		viewName string
	}{
		{
			name:     "primary color function exists",
			color:    PRIMARY,
			viewName: ViewHeading,
		},
		{
			name:     "danger color function exists",
			color:    DANGER,
			viewName: ViewText,
		},
		{
			name:     "transparent color function exists",
			color:    TRANSPARENT,
			viewName: ViewHeading,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := WithColor(tt.color)
			if opt == nil {
				t.Fatal("WithColor() returned nil")
			}

			// WithColor returns a valid Opt function
			// Full testing requires integration with actual views in DOM environment
		})
	}
}

func TestAllColorsArray(t *testing.T) {
	// Verify allColors contains all defined colors
	if len(allColors) != 19 {
		t.Errorf("allColors length = %d, want 19", len(allColors))
	}

	// Verify specific colors are present
	expectedColors := []Color{
		TRANSPARENT,
		PRIMARY,
		SECONDARY,
		SUCCESS,
		DANGER,
		WARNING,
		INFO,
		LIGHT,
		DARK,
		WHITE,
		BLACK,
	}

	for _, expected := range expectedColors {
		found := false
		for _, c := range allColors {
			if c == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("allColors missing expected color %q", expected)
		}
	}
}

func TestColorSubtleVariants(t *testing.T) {
	// Test that subtle variants exist and are distinct
	subtleTests := []struct {
		base   Color
		subtle Color
	}{
		{PRIMARY, PRIMARY_SUBTLE},
		{SECONDARY, SECONDARY_SUBTLE},
		{SUCCESS, SUCCESS_SUBTLE},
		{DANGER, DANGER_SUBTLE},
		{WARNING, WARNING_SUBTLE},
		{INFO, INFO_SUBTLE},
		{LIGHT, LIGHT_SUBTLE},
		{DARK, DARK_SUBTLE},
	}

	for _, tt := range subtleTests {
		t.Run(string(tt.base), func(t *testing.T) {
			if tt.base == tt.subtle {
				t.Errorf("Base color %q and subtle %q should be different", tt.base, tt.subtle)
			}

			// Subtle should contain "-subtle"
			subtleStr := string(tt.subtle)
			if len(subtleStr) < 7 || subtleStr[len(subtleStr)-7:] != "-subtle" {
				t.Errorf("Subtle color %q should end with '-subtle'", tt.subtle)
			}
		})
	}
}
