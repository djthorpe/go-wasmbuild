package bootstrap

import (
	"fmt"
	"testing"
)

func TestPositionConstants(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		expected uint
	}{
		{"TOP", TOP, 1},
		{"BOTTOM", BOTTOM, 2},
		{"START", START, 4},
		{"END", END, 8},
		{"CENTER", CENTER, 16},
		{"MIDDLE", MIDDLE, 32},
		{"NONE", NONE, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if uint(tt.position) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.position, tt.expected)
			}
		})
	}
}

func TestBorderAll(t *testing.T) {
	// BorderAll should be the combination of TOP, BOTTOM, START, and END
	expected := TOP | BOTTOM | START | END
	if BorderAll != expected {
		t.Errorf("BorderAll = %d, want %d", BorderAll, expected)
	}

	// Verify it equals 15 (1 + 2 + 4 + 8)
	if uint(BorderAll) != 15 {
		t.Errorf("BorderAll value = %d, want 15", BorderAll)
	}
}

func TestPositionBitwise(t *testing.T) {
	tests := []struct {
		name        string
		position    Position
		contains    []Position
		notContains []Position
	}{
		{
			name:        "TOP only",
			position:    TOP,
			contains:    []Position{TOP},
			notContains: []Position{BOTTOM, START, END},
		},
		{
			name:        "TOP | BOTTOM",
			position:    TOP | BOTTOM,
			contains:    []Position{TOP, BOTTOM},
			notContains: []Position{START, END},
		},
		{
			name:        "BorderAll",
			position:    BorderAll,
			contains:    []Position{TOP, BOTTOM, START, END},
			notContains: []Position{CENTER, MIDDLE},
		},
		{
			name:        "START | END",
			position:    START | END,
			contains:    []Position{START, END},
			notContains: []Position{TOP, BOTTOM},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check that expected positions are present
			for _, pos := range tt.contains {
				if tt.position&pos == 0 {
					t.Errorf("%s should contain %d but doesn't", tt.name, pos)
				}
			}

			// Check that unexpected positions are not present
			for _, pos := range tt.notContains {
				if tt.position&pos != 0 {
					t.Errorf("%s should not contain %d but does", tt.name, pos)
				}
			}
		})
	}
}

func TestClassName(t *testing.T) {
	tests := []struct {
		position Position
		prefix   string
		expected string
	}{
		{TOP, "border", "border-top"},
		{BOTTOM, "border", "border-bottom"},
		{START, "border", "border-start"},
		{END, "border", "border-end"},
		{CENTER, "border", ""},
		{MIDDLE, "border", ""},
		{NONE, "border", ""},
		{TOP | BOTTOM, "border", ""}, // Combined positions should return empty string
	}

	for _, tt := range tests {
		t.Run(tt.prefix+"_"+tt.expected, func(t *testing.T) {
			result := tt.position.borderClassName(tt.prefix)
			if result != tt.expected {
				t.Errorf("borderClassName(%d, %q) = %q, want %q", tt.position, tt.prefix, result, tt.expected)
			}
		})
	}
}

func TestBorderClassNames(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		expected []string
	}{
		{
			name:     "BorderAll",
			position: BorderAll,
			expected: []string{"border"},
		},
		{
			name:     "TOP",
			position: TOP,
			expected: []string{"border-top"},
		},
		{
			name:     "BOTTOM",
			position: BOTTOM,
			expected: []string{"border-bottom"},
		},
		{
			name:     "START",
			position: START,
			expected: []string{"border-start"},
		},
		{
			name:     "END",
			position: END,
			expected: []string{"border-end"},
		},
		{
			name:     "TOP | BOTTOM",
			position: TOP | BOTTOM,
			expected: []string{"border-top", "border-bottom"},
		},
		{
			name:     "START | END",
			position: START | END,
			expected: []string{"border-start", "border-end"},
		},
		{
			name:     "TOP | START",
			position: TOP | START,
			expected: []string{"border-top", "border-start"},
		},
		{
			name:     "NONE",
			position: NONE,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.borderClassNames()

			if len(result) != len(tt.expected) {
				t.Errorf("borderClassNames() returned %d classes, want %d: %v vs %v",
					len(result), len(tt.expected), result, tt.expected)
				return
			}

			// Check each expected class name is present
			for i, expected := range tt.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("borderClassNames()[%d] = %q, want %q (full result: %v)",
						i, result[i], expected, result)
				}
			}
		})
	}
}

func TestBorderClassNamesOrder(t *testing.T) {
	// Test that the order is consistent and follows TOP -> BOTTOM -> START -> END
	position := TOP | BOTTOM | START | END
	result := position.borderClassNames()

	// Should return just "border" for BorderAll
	if len(result) != 1 || result[0] != "border" {
		t.Errorf("BorderAll should return [\"border\"], got %v", result)
	}
}

func TestPositionIsPowerOfTwo(t *testing.T) {
	// Verify each non-NONE position is a power of 2 (has exactly one bit set)
	positions := []Position{TOP, BOTTOM, START, END, CENTER, MIDDLE}

	for _, pos := range positions {
		// Check if it's a power of 2 by checking if pos & (pos-1) == 0
		val := uint(pos)
		if val == 0 || (val&(val-1)) != 0 {
			t.Errorf("Position %d is not a power of 2", pos)
		}
	}
}

func TestNoneIsZero(t *testing.T) {
	if NONE != 0 {
		t.Errorf("NONE should be 0, got %d", NONE)
	}

	// NONE combined with anything should return that thing
	if NONE|TOP != TOP {
		t.Errorf("NONE | TOP should equal TOP")
	}

	// NONE & anything should be NONE
	if NONE&TOP != NONE {
		t.Errorf("NONE & TOP should equal NONE")
	}
}

func TestMarginClassNames(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		size     int
		expected []string
	}{
		{
			name:     "MarginAll size 0",
			position: MarginAll,
			size:     0,
			expected: []string{"m-0"},
		},
		{
			name:     "MarginAll size 1",
			position: MarginAll,
			size:     1,
			expected: []string{"m-1"},
		},
		{
			name:     "MarginAll size 5",
			position: MarginAll,
			size:     5,
			expected: []string{"m-5"},
		},
		{
			name:     "TOP only size 3",
			position: TOP,
			size:     3,
			expected: []string{"mt-3"},
		},
		{
			name:     "BOTTOM only size 4",
			position: BOTTOM,
			size:     4,
			expected: []string{"mb-4"},
		},
		{
			name:     "START only size 2",
			position: START,
			size:     2,
			expected: []string{"ms-2"},
		},
		{
			name:     "END only size 1",
			position: END,
			size:     1,
			expected: []string{"me-1"},
		},
		{
			name:     "TOP | BOTTOM size 3 (vertical)",
			position: TOP | BOTTOM,
			size:     3,
			expected: []string{"my-3"},
		},
		{
			name:     "START | END size 4 (horizontal)",
			position: START | END,
			size:     4,
			expected: []string{"mx-4"},
		},
		{
			name:     "TOP | START size 2",
			position: TOP | START,
			size:     2,
			expected: []string{"mt-2", "ms-2"},
		},
		{
			name:     "BOTTOM | END size 3",
			position: BOTTOM | END,
			size:     3,
			expected: []string{"mb-3", "me-3"},
		},
		{
			name:     "TOP | BOTTOM | START size 5",
			position: TOP | BOTTOM | START,
			size:     5,
			expected: []string{"my-5", "ms-5"}, // my-5 for vertical + ms-5 for start
		},
		{
			name:     "TOP | BOTTOM | END size 2",
			position: TOP | BOTTOM | END,
			size:     2,
			expected: []string{"my-2", "me-2"}, // my-2 for vertical + me-2 for end
		},
		{
			name:     "START | END | TOP size 4",
			position: START | END | TOP,
			size:     4,
			expected: []string{"mx-4", "mt-4"}, // mx-4 for horizontal + mt-4 for top
		},
		{
			name:     "START | END | BOTTOM size 3",
			position: START | END | BOTTOM,
			size:     3,
			expected: []string{"mx-3", "mb-3"}, // mx-3 for horizontal + mb-3 for bottom
		},
		{
			name:     "NONE size 3",
			position: NONE,
			size:     3,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.marginClassNames(tt.size)

			if len(result) != len(tt.expected) {
				t.Errorf("marginClassNames(%d) returned %d classes, want %d: got %v, want %v",
					tt.size, len(result), len(tt.expected), result, tt.expected)
				return
			}

			// Check each expected class name is present
			for i, expected := range tt.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("marginClassNames(%d)[%d] = %q, want %q (full result: %v)",
						tt.size, i, result[i], expected, result)
				}
			}
		})
	}
}

func TestMarginClassNamesNegativeSizes(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		size     int
		expected []string
	}{
		{
			name:     "MarginAll negative size",
			position: MarginAll,
			size:     -3,
			expected: []string{"m-n3"},
		},
		{
			name:     "TOP negative margin",
			position: TOP,
			size:     -2,
			expected: []string{"mt-n2"},
		},
		{
			name:     "START | END negative margin",
			position: START | END,
			size:     -1,
			expected: []string{"mx-n1"},
		},
		{
			name:     "BOTTOM negative margin",
			position: BOTTOM,
			size:     -4,
			expected: []string{"mb-n4"},
		},
		{
			name:     "TOP | BOTTOM negative margin",
			position: TOP | BOTTOM,
			size:     -5,
			expected: []string{"my-n5"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.marginClassNames(tt.size)

			if len(result) != len(tt.expected) {
				t.Errorf("marginClassNames(%d) returned %d classes, want %d: got %v, want %v",
					tt.size, len(result), len(tt.expected), result, tt.expected)
				return
			}

			for i, expected := range tt.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("marginClassNames(%d)[%d] = %q, want %q",
						tt.size, i, result[i], expected)
				}
			}
		})
	}
}

func TestMarginClassNamesSizeRange(t *testing.T) {
	// Test common Bootstrap margin sizes (0-5)
	sizes := []int{0, 1, 2, 3, 4, 5}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			// Test TOP position
			result := TOP.marginClassNames(size)
			expected := fmt.Sprintf("mt-%d", size)

			if len(result) != 1 {
				t.Errorf("TOP.marginClassNames(%d) should return 1 class, got %d", size, len(result))
			}

			if result[0] != expected {
				t.Errorf("TOP.marginClassNames(%d) = %q, want %q", size, result[0], expected)
			}
		})
	}
}

func TestMarginAllConstant(t *testing.T) {
	// MarginAll should be the combination of TOP, BOTTOM, START, and END
	expected := TOP | BOTTOM | START | END
	if MarginAll != expected {
		t.Errorf("MarginAll = %d, want %d", MarginAll, expected)
	}

	// Verify it equals 15 (1 + 2 + 4 + 8)
	if uint(MarginAll) != 15 {
		t.Errorf("MarginAll value = %d, want 15", MarginAll)
	}
}

func TestPaddingAllConstant(t *testing.T) {
	// PaddingAll should be the combination of TOP, BOTTOM, START, and END
	expected := TOP | BOTTOM | START | END
	if PaddingAll != expected {
		t.Errorf("PaddingAll = %d, want %d", PaddingAll, expected)
	}

	// Verify it equals 15 (1 + 2 + 4 + 8)
	if uint(PaddingAll) != 15 {
		t.Errorf("PaddingAll value = %d, want 15", PaddingAll)
	}
}

func TestMarginClassName(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		size     int
		expected string
	}{
		{
			name:     "TOP positive size",
			position: TOP,
			size:     3,
			expected: "mt-3",
		},
		{
			name:     "BOTTOM positive size",
			position: BOTTOM,
			size:     4,
			expected: "mb-4",
		},
		{
			name:     "START positive size",
			position: START,
			size:     2,
			expected: "ms-2",
		},
		{
			name:     "END positive size",
			position: END,
			size:     1,
			expected: "me-1",
		},
		{
			name:     "TOP negative size",
			position: TOP,
			size:     -2,
			expected: "mt-n2",
		},
		{
			name:     "BOTTOM negative size",
			position: BOTTOM,
			size:     -3,
			expected: "mb-n3",
		},
		{
			name:     "NONE",
			position: NONE,
			size:     3,
			expected: "",
		},
		{
			name:     "CENTER",
			position: CENTER,
			size:     3,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.marginClassName(tt.size)
			if result != tt.expected {
				t.Errorf("marginClassName(%d) = %q, want %q", tt.size, result, tt.expected)
			}
		})
	}
}

func TestPaddingClassNames(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		size     int
		expected []string
	}{
		{
			name:     "PaddingAll size 1",
			position: PaddingAll,
			size:     1,
			expected: []string{"p-1"},
		},
		{
			name:     "PaddingAll size 5",
			position: PaddingAll,
			size:     5,
			expected: []string{"p-5"},
		},
		{
			name:     "TOP only size 3",
			position: TOP,
			size:     3,
			expected: []string{"pt-3"},
		},
		{
			name:     "BOTTOM only size 4",
			position: BOTTOM,
			size:     4,
			expected: []string{"pb-4"},
		},
		{
			name:     "START only size 2",
			position: START,
			size:     2,
			expected: []string{"ps-2"},
		},
		{
			name:     "END only size 1",
			position: END,
			size:     1,
			expected: []string{"pe-1"},
		},
		{
			name:     "TOP | BOTTOM size 3 (vertical)",
			position: TOP | BOTTOM,
			size:     3,
			expected: []string{"py-3"},
		},
		{
			name:     "START | END size 4 (horizontal)",
			position: START | END,
			size:     4,
			expected: []string{"px-4"},
		},
		{
			name:     "TOP | START size 2",
			position: TOP | START,
			size:     2,
			expected: []string{"pt-2", "ps-2"},
		},
		{
			name:     "BOTTOM | END size 3",
			position: BOTTOM | END,
			size:     3,
			expected: []string{"pb-3", "pe-3"},
		},
		{
			name:     "TOP | BOTTOM | START size 5",
			position: TOP | BOTTOM | START,
			size:     5,
			expected: []string{"py-5", "ps-5"}, // py-5 for vertical + ps-5 for start
		},
		{
			name:     "TOP | BOTTOM | END size 2",
			position: TOP | BOTTOM | END,
			size:     2,
			expected: []string{"py-2", "pe-2"}, // py-2 for vertical + pe-2 for end
		},
		{
			name:     "START | END | TOP size 4",
			position: START | END | TOP,
			size:     4,
			expected: []string{"px-4", "pt-4"}, // px-4 for horizontal + pt-4 for top
		},
		{
			name:     "START | END | BOTTOM size 3",
			position: START | END | BOTTOM,
			size:     3,
			expected: []string{"px-3", "pb-3"}, // px-3 for horizontal + pb-3 for bottom
		},
		{
			name:     "NONE size 3",
			position: NONE,
			size:     3,
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.paddingClassNames(tt.size)

			if len(result) != len(tt.expected) {
				t.Errorf("paddingClassNames(%d) returned %d classes, want %d: got %v, want %v",
					tt.size, len(result), len(tt.expected), result, tt.expected)
				return
			}

			// Check each expected class name is present
			for i, expected := range tt.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("paddingClassNames(%d)[%d] = %q, want %q (full result: %v)",
						tt.size, i, result[i], expected, result)
				}
			}
		})
	}
}

func TestPaddingClassNamesSizeRange(t *testing.T) {
	// Test common Bootstrap padding sizes (1-5, no 0 or negative)
	sizes := []int{1, 2, 3, 4, 5}

	for _, size := range sizes {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			// Test TOP position
			result := TOP.paddingClassNames(size)
			expected := fmt.Sprintf("pt-%d", size)

			if len(result) != 1 {
				t.Errorf("TOP.paddingClassNames(%d) should return 1 class, got %d", size, len(result))
			}

			if result[0] != expected {
				t.Errorf("TOP.paddingClassNames(%d) = %q, want %q", size, result[0], expected)
			}
		})
	}
}

func TestPaddingClassName(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		size     int
		expected string
	}{
		{
			name:     "TOP positive size",
			position: TOP,
			size:     3,
			expected: "pt-3",
		},
		{
			name:     "BOTTOM positive size",
			position: BOTTOM,
			size:     4,
			expected: "pb-4",
		},
		{
			name:     "START positive size",
			position: START,
			size:     2,
			expected: "ps-2",
		},
		{
			name:     "END positive size",
			position: END,
			size:     1,
			expected: "pe-1",
		},
		{
			name:     "SIZE 5",
			position: TOP,
			size:     5,
			expected: "pt-5",
		},
		{
			name:     "NONE",
			position: NONE,
			size:     3,
			expected: "",
		},
		{
			name:     "CENTER",
			position: CENTER,
			size:     3,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.position.paddingClassName(tt.size)
			if result != tt.expected {
				t.Errorf("paddingClassName(%d) = %q, want %q", tt.size, result, tt.expected)
			}
		})
	}
}
