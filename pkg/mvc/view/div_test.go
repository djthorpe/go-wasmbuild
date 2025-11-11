package views_test

import (
	"testing"

	// Packages
	mvc "github.com/djthorpe/go-wasmbuild/pkg/mvc"
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
		div2 := mvc.NewViewWithElement(div.Root())
		assert.NotNil(div2)
		assert.Equal(ViewDiv, div2.Name())
		assert.NotNil(div2.Root())
		assert.Equal("DIV", div2.Root().TagName())
	})
}
