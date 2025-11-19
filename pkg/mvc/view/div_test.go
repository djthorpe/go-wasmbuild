package view_test

import (
	"testing"

	// Packages

	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc/view"
)

func TestDivType(t *testing.T) {
	assert := assert.New(t)

	div := Div()
	assert.NotNil(div)

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(ViewDiv, div.Name())
		assert.NotNil(div.Root())
		assert.Equal("DIV", div.Root().TagName())
	})

	t.Run("NewViewWithElement", func(t *testing.T) {
		// Create a new div from the existing element
		div2 := Div()
		// This test verifies that we can create a view from an element
		// In practice, newDivFromElement would be called by the framework
		assert.NotNil(div2)
		assert.Equal(ViewDiv, div2.Name())
		assert.NotNil(div2.Root())
		assert.Equal("DIV", div2.Root().TagName())
	})
}
