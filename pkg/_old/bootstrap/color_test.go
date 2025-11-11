package bootstrap

import (
	"testing"
)

func TestColorConstants(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
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
				t.Errorf("%s = %q, want %q", tt.name, tt.color, tt.expected)
			}
		})
	}
}

func TestColorClassName(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		prefix   string
		expected string
	}{
		{"PRIMARY with text", PRIMARY, "text", "text-primary"},
		{"PRIMARY with bg", PRIMARY, "bg", "bg-primary"},
		{"PRIMARY with btn", PRIMARY, "btn", "btn-primary"},
		{"PRIMARY with border", PRIMARY, "border", "border-primary"},
		{"PRIMARY with alert", PRIMARY, "alert", "alert-primary"},
		{"SECONDARY with text", SECONDARY, "text", "text-secondary"},
		{"SUCCESS with bg", SUCCESS, "bg", "bg-success"},
		{"DANGER with btn", DANGER, "btn", "btn-danger"},
		{"WARNING with alert", WARNING, "alert", "alert-warning"},
		{"INFO with text", INFO, "text", "text-info"},
		{"LIGHT with bg", LIGHT, "bg", "bg-light"},
		{"DARK with text", DARK, "text", "text-dark"},
		{"WHITE with bg", WHITE, "bg", "bg-white"},
		{"BLACK with text", BLACK, "text", "text-black"},
		{"PRIMARY_SUBTLE with bg", PRIMARY_SUBTLE, "bg", "bg-primary-subtle"},
		{"SECONDARY_SUBTLE with text", SECONDARY_SUBTLE, "text", "text-secondary-subtle"},
		{"SUCCESS_SUBTLE with alert", SUCCESS_SUBTLE, "alert", "alert-success-subtle"},
		{"DANGER_SUBTLE with border", DANGER_SUBTLE, "border", "border-danger-subtle"},
		{"WARNING_SUBTLE with btn", WARNING_SUBTLE, "btn", "btn-warning-subtle"},
		{"INFO_SUBTLE with bg", INFO_SUBTLE, "bg", "bg-info-subtle"},
		{"LIGHT_SUBTLE with text", LIGHT_SUBTLE, "text", "text-light-subtle"},
		{"DARK_SUBTLE with bg", DARK_SUBTLE, "bg", "bg-dark-subtle"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.color.className(tt.prefix)
			if result != tt.expected {
				t.Errorf("className(%q, %q) = %q, want %q", tt.color, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestColorClassNameWithEmptyPrefix(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{PRIMARY, "-primary"},
		{SECONDARY, "-secondary"},
		{SUCCESS, "-success"},
		{DANGER, "-danger"},
		{WARNING, "-warning"},
		{INFO, "-info"},
	}

	for _, tt := range tests {
		t.Run(string(tt.color), func(t *testing.T) {
			result := tt.color.className("")
			if result != tt.expected {
				t.Errorf("className(%q, \"\") = %q, want %q", tt.color, result, tt.expected)
			}
		})
	}
}

func TestColorClassNameConsistency(t *testing.T) {
	// Test that all color constants work with className
	colors := []Color{
		PRIMARY, PRIMARY_SUBTLE,
		SECONDARY, SECONDARY_SUBTLE,
		SUCCESS, SUCCESS_SUBTLE,
		DANGER, DANGER_SUBTLE,
		WARNING, WARNING_SUBTLE,
		INFO, INFO_SUBTLE,
		LIGHT, LIGHT_SUBTLE,
		DARK, DARK_SUBTLE,
		WHITE, BLACK,
	}

	prefix := "test"
	for _, color := range colors {
		t.Run(string(color), func(t *testing.T) {
			result := color.className(prefix)
			expected := prefix + "-" + string(color)
			if result != expected {
				t.Errorf("className(%q, %q) = %q, want %q", color, prefix, result, expected)
			}
		})
	}
}

func TestColorType(t *testing.T) {
	// Test that Color is a string type
	var color Color = "custom-color"
	if string(color) != "custom-color" {
		t.Errorf("Color type conversion failed: got %q, want %q", string(color), "custom-color")
	}
}

func TestColorBootstrapPrefixes(t *testing.T) {
	// Test common Bootstrap prefix patterns
	tests := []struct {
		prefix string
		color  Color
	}{
		{"text", PRIMARY},
		{"bg", SUCCESS},
		{"btn", DANGER},
		{"border", WARNING},
		{"alert", INFO},
		{"badge", LIGHT},
		{"table", DARK},
		{"btn-outline", SECONDARY},
		{"link", PRIMARY_SUBTLE},
		{"list-group-item", SUCCESS_SUBTLE},
	}

	for _, tt := range tests {
		t.Run(tt.prefix+"-"+string(tt.color), func(t *testing.T) {
			result := tt.color.className(tt.prefix)
			expected := tt.prefix + "-" + string(tt.color)
			if result != expected {
				t.Errorf("className() = %q, want %q", result, expected)
			}
		})
	}
}

func TestColorSubtleVariants(t *testing.T) {
	// Test that all subtle variants end with "-subtle"
	subtleColors := []Color{
		PRIMARY_SUBTLE,
		SECONDARY_SUBTLE,
		SUCCESS_SUBTLE,
		DANGER_SUBTLE,
		WARNING_SUBTLE,
		INFO_SUBTLE,
		LIGHT_SUBTLE,
		DARK_SUBTLE,
	}

	for _, color := range subtleColors {
		t.Run(string(color), func(t *testing.T) {
			colorStr := string(color)
			if len(colorStr) < 7 || colorStr[len(colorStr)-7:] != "-subtle" {
				t.Errorf("Subtle color %q does not end with '-subtle'", color)
			}
		})
	}
}

func TestColorNonSubtleVariants(t *testing.T) {
	// Test that non-subtle variants don't contain "-subtle"
	nonSubtleColors := []Color{
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

	for _, color := range nonSubtleColors {
		t.Run(string(color), func(t *testing.T) {
			colorStr := string(color)
			if len(colorStr) >= 7 && colorStr[len(colorStr)-7:] == "-subtle" {
				t.Errorf("Non-subtle color %q should not end with '-subtle'", color)
			}
		})
	}
}

func TestDangerTypo(t *testing.T) {
	// Note: There's a typo in the constant name DANGER_SUBTLE (should be DANGER_SUBTLE)
	// but the value is correct
	if string(DANGER_SUBTLE) != "danger-subtle" {
		t.Errorf("DANGER_SUBTLE (typo in constant name) has wrong value: got %q, want %q",
			DANGER_SUBTLE, "danger-subtle")
	}
}
