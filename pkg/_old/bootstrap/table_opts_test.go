package bootstrap

import (
	"strings"
	"testing"
)

func TestTableOpts_Striped(t *testing.T) {
	table := Table(WithStripedRows())
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-striped") {
		t.Errorf("Expected 'table-striped' in class, got '%s'", class)
	}
}

func TestTableOpts_StripedColumns(t *testing.T) {
	table := Table(WithStripedColumns())
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-striped-columns") {
		t.Errorf("Expected 'table-striped-columns' in class, got '%s'", class)
	}
}

func TestTableOpts_Hover(t *testing.T) {
	table := Table(WithHover())
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-hover") {
		t.Errorf("Expected 'table-hover' in class, got '%s'", class)
	}
}

func TestTableOpts_Bordered(t *testing.T) {
	table := Table(WithBordered())
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-bordered") {
		t.Errorf("Expected 'table-bordered' in class, got '%s'", class)
	}
}

func TestTableOpts_Borderless(t *testing.T) {
	table := Table(WithBorderless())
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-borderless") {
		t.Errorf("Expected 'table-borderless' in class, got '%s'", class)
	}
}

func TestTableOpts_Small(t *testing.T) {
	table := Table(WithSize(SizeSmall))
	class := table.Element().GetAttribute("class")
	if !strings.Contains(class, "table-sm") {
		t.Errorf("Expected 'table-sm' in class, got '%s'", class)
	}
}

func TestTableOpts_Variants(t *testing.T) {
	tests := []struct {
		name     string
		color    Color
		expected string
	}{
		{"Primary", PRIMARY, "table-primary"},
		{"Secondary", SECONDARY, "table-secondary"},
		{"Success", SUCCESS, "table-success"},
		{"Danger", DANGER, "table-danger"},
		{"Warning", WARNING, "table-warning"},
		{"Info", INFO, "table-info"},
		{"Light", LIGHT, "table-light"},
		{"Dark", DARK, "table-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := Table(WithColor(tt.color))
			class := table.Element().GetAttribute("class")
			if !strings.Contains(class, tt.expected) {
				t.Errorf("Expected '%s' in class, got '%s'", tt.expected, class)
			}
		})
	}
}

func TestTableOpts_MultipleOptions(t *testing.T) {
	table := Table(
		WithStripedRows(),
		WithHover(),
		WithBordered(),
		WithSize(SizeSmall),
	)

	class := table.Element().GetAttribute("class")
	expectedClasses := []string{"table", "table-striped", "table-hover", "table-bordered", "table-sm"}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected '%s' in class, got '%s'", expected, class)
		}
	}
}

func TestTableOpts_StripedWithVariant(t *testing.T) {
	table := Table(WithColor(DARK), WithStripedRows(), WithHover())

	class := table.Element().GetAttribute("class")
	expectedClasses := []string{"table", "table-dark", "table-striped", "table-hover"}

	for _, expected := range expectedClasses {
		if !strings.Contains(class, expected) {
			t.Errorf("Expected '%s' in class, got '%s'", expected, class)
		}
	}
}

func TestTableOpts_Responsive(t *testing.T) {
	tests := []struct {
		name       string
		breakpoint Breakpoint
		expected   string
	}{
		{"Responsive", BreakpointDefault, "table-responsive"},
		{"ResponsiveSm", BreakpointSmall, "table-responsive-sm"},
		{"ResponsiveMd", BreakpointMedium, "table-responsive-md"},
		{"ResponsiveLg", BreakpointLarge, "table-responsive-lg"},
		{"ResponsiveXl", BreakpointXLarge, "table-responsive-xl"},
		{"ResponsiveXxl", BreakpointXXLarge, "table-responsive-xxl"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := Table(WithResponsive(tt.breakpoint))
			class := table.Element().GetAttribute("class")
			if !strings.Contains(class, tt.expected) {
				t.Errorf("Expected '%s' in class, got '%s'", tt.expected, class)
			}
		})
	}
}
