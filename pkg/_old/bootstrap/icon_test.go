package bootstrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func TestIcon_Basic(t *testing.T) {
	icon := Icon("heart-fill")
	assert.NotNil(t, icon)
	assert.Equal(t, "I", icon.Element().TagName())

	// Check classes
	classList := icon.Element().ClassList()
	assert.True(t, classList.Contains("bi"))
	assert.True(t, classList.Contains("bi-heart-fill"))
}

func TestIcon_WithColor(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		want  string
	}{
		{"primary", PRIMARY, "text-primary"},
		{"secondary", SECONDARY, "text-secondary"},
		{"success", SUCCESS, "text-success"},
		{"danger", DANGER, "text-danger"},
		{"warning", WARNING, "text-warning"},
		{"info", INFO, "text-info"},
		{"light", LIGHT, "text-light"},
		{"dark", DARK, "text-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := Icon("alarm", WithColor(tt.color))
			classList := icon.Element().ClassList()
			assert.True(t, classList.Contains("bi"))
			assert.True(t, classList.Contains("bi-alarm"))
			assert.True(t, classList.Contains(tt.want))
		})
	}
}

func TestIcon_WithSize(t *testing.T) {
	// Icons use font-size classes directly, not button sizes
	tests := []struct {
		name  string
		class string
	}{
		{"fs-1", "fs-1"},
		{"fs-2", "fs-2"},
		{"fs-3", "fs-3"},
		{"fs-4", "fs-4"},
		{"fs-5", "fs-5"},
		{"fs-6", "fs-6"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := Icon("star-fill", WithClass(tt.class))
			classList := icon.Element().ClassList()
			assert.True(t, classList.Contains("bi"))
			assert.True(t, classList.Contains("bi-star-fill"))
			assert.True(t, classList.Contains(tt.class))
		})
	}
}

func TestIcon_WithClass(t *testing.T) {
	icon := Icon("bootstrap", WithClass("fs-1"), WithClass("me-2"))
	classList := icon.Element().ClassList()
	assert.True(t, classList.Contains("bi"))
	assert.True(t, classList.Contains("bi-bootstrap"))
	assert.True(t, classList.Contains("fs-1"))
	assert.True(t, classList.Contains("me-2"))
}

func TestIcon_WithAriaLabel(t *testing.T) {
	icon := Icon("heart", WithAriaLabel("Favorite"))
	assert.Equal(t, "Favorite", icon.Element().GetAttribute("aria-label"))
}

func TestIcon_MultipleOptions(t *testing.T) {
	icon := Icon("check-circle-fill",
		WithColor(SUCCESS),
		WithClass("fs-3"),
		WithMargin(END, 2),
		WithAriaLabel("Success"),
	)

	classList := icon.Element().ClassList()
	assert.True(t, classList.Contains("bi"))
	assert.True(t, classList.Contains("bi-check-circle-fill"))
	assert.True(t, classList.Contains("text-success"))
	assert.True(t, classList.Contains("fs-3"))
	assert.True(t, classList.Contains("me-2"))
	assert.Equal(t, "Success", icon.Element().GetAttribute("aria-label"))
}

func TestIcon_DifferentIconNames(t *testing.T) {
	tests := []struct {
		name      string
		iconName  string
		wantClass string
	}{
		{"heart", "heart", "bi-heart"},
		{"heart-fill", "heart-fill", "bi-heart-fill"},
		{"alarm", "alarm", "bi-alarm"},
		{"alarm-fill", "alarm-fill", "bi-alarm-fill"},
		{"bootstrap", "bootstrap", "bi-bootstrap"},
		{"github", "github", "bi-github"},
		{"cloud-download", "cloud-download", "bi-cloud-download"},
		{"exclamation-triangle", "exclamation-triangle", "bi-exclamation-triangle"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icon := Icon(tt.iconName)
			classList := icon.Element().ClassList()
			assert.True(t, classList.Contains("bi"))
			assert.True(t, classList.Contains(tt.wantClass))
		})
	}
}

func TestIcon_Component(t *testing.T) {
	icon := Icon("star")

	// Test that it implements Component interface
	assert.NotNil(t, icon.Element())
	assert.Equal(t, "I", icon.Element().TagName())

	// Test that Name() returns the component name
	assert.NotEmpty(t, icon.Name())
}

func TestIcon_Append(t *testing.T) {
	// Icons typically don't have children, but test Append for completeness
	icon := Icon("info-circle")

	// Should not panic when calling Append (even if it doesn't make semantic sense)
	assert.NotPanics(t, func() {
		icon.Append("test")
	})
}
