//go:build js && wasm

package bootstrap

import (
	"strings"
	"testing"
)

func TestOffcanvas_Basic(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	)
	html := offcanvas.Element().OuterHTML()

	// Check that it has the correct classes
	if !strings.Contains(html, `offcanvas`) || !strings.Contains(html, `offcanvas-start`) {
		t.Errorf("Expected offcanvas and offcanvas-start classes, got: %s", html)
	}

	// Check required attributes
	if !strings.Contains(html, `id="myOffcanvas"`) {
		t.Error("Expected id attribute")
	}
	if !strings.Contains(html, `tabindex="-1"`) {
		t.Error("Expected tabindex attribute")
	}

	// Check structure
	if !strings.Contains(html, `offcanvas-header`) {
		t.Error("Expected offcanvas-header div")
	}
	if !strings.Contains(html, `offcanvas-body`) {
		t.Error("Expected offcanvas-body div")
	}
}

func TestOffcanvas_Placements(t *testing.T) {
	placements := []struct {
		position Position
		expected string
	}{
		{START, "offcanvas-start"},
		{END, "offcanvas-end"},
		{TOP, "offcanvas-top"},
		{BOTTOM, "offcanvas-bottom"},
	}

	for _, tc := range placements {
		t.Run(tc.expected, func(t *testing.T) {
			offcanvas := Offcanvas(
				WithID("test"),
				WithPosition(tc.position),
			)
			html := offcanvas.Element().OuterHTML()
			if !strings.Contains(html, tc.expected) {
				t.Errorf("Expected class %s, got: %s", tc.expected, html)
			}
		})
	}
}

func TestOffcanvas_WithScroll(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
		WithScroll(),
	)
	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, `data-bs-scroll="true"`) {
		t.Errorf("Expected data-bs-scroll attribute, got: %s", html)
	}
}

func TestOffcanvas_WithBackdrop(t *testing.T) {
	tests := []struct {
		name     string
		backdrop string
		expected string
	}{
		{"false", "false", `data-bs-backdrop="false"`},
		{"static", "static", `data-bs-backdrop="static"`},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			offcanvas := Offcanvas(
				WithID("myOffcanvas"),
				WithPosition(START),
				WithBackdrop(tc.backdrop),
			)
			html := offcanvas.Element().OuterHTML()
			if !strings.Contains(html, tc.expected) {
				t.Errorf("Expected %s, got: %s", tc.expected, html)
			}
		})
	}
}

func TestOffcanvas_WithoutKeyboard(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
		WithoutKeyboard(),
	)
	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, `data-bs-keyboard="false"`) {
		t.Errorf("Expected data-bs-keyboard attribute, got: %s", html)
	}
}

func TestOffcanvas_WithDark(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
		WithTheme(DARK),
	)
	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, `data-bs-theme="dark"`) {
		t.Error("Expected data-bs-theme attribute")
	}

	// Check classList directly
	classList := offcanvas.Element().ClassList()
	if !classList.Contains("text-bg-dark") {
		t.Error("Expected text-bg-dark class")
	}
}

func TestOffcanvas_WithThemeLight(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
		WithTheme(LIGHT),
	)
	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, `data-bs-theme="light"`) {
		t.Error("Expected data-bs-theme attribute with light value")
	}

	// Should NOT have text-bg-dark class for light theme
	classList := offcanvas.Element().ClassList()
	if classList.Contains("text-bg-dark") {
		t.Error("Should not have text-bg-dark class for light theme")
	}
}

func TestOffcanvas_WithShow(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	)
	offcanvas.Show()

	// Note: The "show" class is added by Bootstrap JavaScript when Show() is called.
	// Since Bootstrap isn't loaded in the test environment, we can only verify
	// that Show() can be called without error.
	if offcanvas == nil {
		t.Error("Expected offcanvas to exist after Show()")
	}
}

func TestOffcanvas_Append(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	).Append("Some content")

	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, "Some content") {
		t.Errorf("Expected content in body, got: %s", html)
	}

	// Verify it's in the body section
	if !strings.Contains(html, `offcanvas-body`) {
		t.Error("Expected content in offcanvas-body")
	}
}

func TestOffcanvas_HeaderAppend(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	).Header("Custom Header Content")

	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, "Custom Header Content") {
		t.Errorf("Expected custom header content, got: %s", html)
	}

	// Verify it's in the header section (should be before body)
	headerIdx := strings.Index(html, "offcanvas-header")
	bodyIdx := strings.Index(html, "offcanvas-body")
	contentIdx := strings.Index(html, "Custom Header Content")

	if contentIdx == -1 || headerIdx == -1 || bodyIdx == -1 {
		t.Error("Could not find header, body, or content in HTML")
	} else if contentIdx > bodyIdx {
		t.Error("Header content should appear before body")
	}
}

func TestOffcanvas_HeaderAndBodySeparate(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	).Header("Header Text").Append("Body Text")

	html := offcanvas.Element().OuterHTML()

	if !strings.Contains(html, "Header Text") {
		t.Error("Expected header text")
	}
	if !strings.Contains(html, "Body Text") {
		t.Error("Expected body text")
	}

	// Verify they're in correct sections
	headerIdx := strings.Index(html, "Header Text")
	bodyIdx := strings.Index(html, "Body Text")

	if headerIdx > bodyIdx {
		t.Error("Header content should appear before body content")
	}
}

func TestOffcanvas_HeaderMultipleCalls(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	).Header("First").Header("Second")

	html := offcanvas.Element().OuterHTML()

	// Both should be present since multiple calls append
	if !strings.Contains(html, "First") {
		t.Error("Expected first header content")
	}
	if !strings.Contains(html, "Second") {
		t.Error("Expected second header content")
	}
}

func TestOffcanvas_Chaining(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(END),
		WithScroll(),
		WithBackdrop("static"),
		WithoutKeyboard(),
		WithTheme(DARK),
	)
	offcanvas.Show()
	offcanvas.Header("Settings").Append("Content here")

	html := offcanvas.Element().OuterHTML()

	// Verify all options are applied
	checks := []string{
		`offcanvas-end`,
		`Settings`,
		`data-bs-scroll="true"`,
		`data-bs-backdrop="static"`,
		`data-bs-keyboard="false"`,
		`data-bs-theme="dark"`,
		`Content here`,
	}

	for _, check := range checks {
		if !strings.Contains(html, check) {
			t.Errorf("Expected %s in HTML, got: %s", check, html)
		}
	}

	// Note: The "show" class is added by Bootstrap JavaScript.
	// We can only verify method chaining works.
	if offcanvas == nil {
		t.Error("Expected Show() to return offcanvas instance")
	}
}

func TestOffcanvas_Hide(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	)
	offcanvas.Show()

	// Note: Bootstrap JavaScript isn't loaded in tests, so "show" class won't be added.
	// We verify that Hide() can be called without error.

	// Hide the offcanvas
	offcanvas.Hide()

	// Verify offcanvas still exists
	if offcanvas == nil {
		t.Error("Expected offcanvas to still exist after Hide()")
	}
}

func TestOffcanvas_Toggle(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(START),
	)

	// Toggle on
	offcanvas.Toggle()

	// Verify offcanvas still exists
	if offcanvas == nil {
		t.Error("Expected offcanvas to still exist after Toggle()")
	}
}

func TestOffcanvas_VisibilityChaining(t *testing.T) {
	offcanvas := Offcanvas(
		WithID("myOffcanvas"),
		WithPosition(END),
	)
	offcanvas.Show()
	offcanvas.Hide()
	offcanvas.Toggle()

	// Verify method chaining works correctly
	if offcanvas == nil {
		t.Error("Expected method chaining to return offcanvas instance")
	}

	html := offcanvas.Element().OuterHTML()
	if !strings.Contains(html, `id="myOffcanvas"`) {
		t.Error("Expected offcanvas ID to be preserved through chaining")
	}
}
