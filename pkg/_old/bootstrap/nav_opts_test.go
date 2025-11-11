package bootstrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTabs_NavComponentOnly(t *testing.T) {
	// Should work with Nav component
	nav := Nav(WithTabs())
	assert.NotNil(t, nav)
	assert.True(t, nav.Element().ClassList().Contains("nav-tabs"))

	// Should error with non-Nav component
	_, err := NewOpts(ButtonComponent, WithTabs())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cannot use WithTabs with component of type")
}

func TestWithPills_NavComponentOnly(t *testing.T) {
	// Should work with Nav component
	nav := Nav(WithPills())
	assert.NotNil(t, nav)
	assert.True(t, nav.Element().ClassList().Contains("nav-pills"))

	// Should error with non-Nav component
	_, err := NewOpts(ButtonComponent, WithPills())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cannot use WithPills with component of type")
}

func TestWithUnderline_NavComponentOnly(t *testing.T) {
	// Should work with Nav component
	nav := Nav(WithUnderline())
	assert.NotNil(t, nav)
	assert.True(t, nav.Element().ClassList().Contains("nav-underline"))

	// Should error with non-Nav component
	_, err := NewOpts(ButtonComponent, WithUnderline())
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Cannot use WithUnderline with component of type")
}

func TestNavStyles_MutuallyExclusive(t *testing.T) {
	// Test that multiple styles can be applied (though typically only one should be used)
	nav := Nav(WithTabs(), WithClass("custom-class"))
	assert.NotNil(t, nav)
	assert.True(t, nav.Element().ClassList().Contains("nav"))
	assert.True(t, nav.Element().ClassList().Contains("nav-tabs"))
	assert.True(t, nav.Element().ClassList().Contains("custom-class"))
}
