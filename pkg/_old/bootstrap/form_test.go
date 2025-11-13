package bootstrap_test

import (
	"testing"

	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestForm(t *testing.T) {
	form := bs.Form()
	assert.NotNil(t, form)

	elem := form.Element()
	assert.NotNil(t, elem)
	assert.Equal(t, "FORM", elem.TagName())
}

func TestInput(t *testing.T) {
	input := bs.Input("testfield")
	assert.NotNil(t, input)

	elem := input.Element()
	assert.NotNil(t, elem)
	assert.Equal(t, "INPUT", elem.TagName())

	// Check default attributes
	assert.Equal(t, "text", elem.GetAttribute("type"))
	assert.Equal(t, "testfield", elem.GetAttribute("name"))
	assert.True(t, elem.ClassList().Contains("form-control"))
}

func TestInputValue(t *testing.T) {
	input := bs.Input("testfield")

	// Test setting value
	returnedValue := input.Value("test value")
	assert.Equal(t, "test value", returnedValue)

	// Verify value was set
	value := input.Value()
	assert.Equal(t, "test value", value)
	assert.Equal(t, "test value", input.Element().GetAttribute("value"))
}

func TestNumberInput(t *testing.T) {
	input := bs.NumberInput("quantity")
	assert.NotNil(t, input)

	elem := input.Element()
	assert.NotNil(t, elem)
	assert.Equal(t, "INPUT", elem.TagName())

	// Check that type is "number" instead of "text"
	assert.Equal(t, "number", elem.GetAttribute("type"))

	// Check other defaults are still present
	assert.Equal(t, "quantity", elem.GetAttribute("name"))
	assert.True(t, elem.ClassList().Contains("form-control"))
}

func TestLabel(t *testing.T) {
	label := bs.Label("Test Label")
	assert.NotNil(t, label)

	elem := label.Element()
	assert.NotNil(t, elem)
	assert.Equal(t, "LABEL", elem.TagName())

	// Check default class
	assert.True(t, elem.ClassList().Contains("form-label"))

	// Check text content
	assert.Equal(t, "Test Label", elem.TextContent())
}

func TestLabelEmpty(t *testing.T) {
	label := bs.Label("")
	assert.NotNil(t, label)

	elem := label.Element()
	assert.NotNil(t, elem)
	assert.Equal(t, "LABEL", elem.TagName())
	assert.True(t, elem.ClassList().Contains("form-label"))
	assert.Equal(t, "", elem.TextContent())
}

func TestWithReadOnly(t *testing.T) {
	input := bs.Input("testfield", bs.WithReadOnly())

	elem := input.Element()
	assert.NotNil(t, elem)

	// Check readonly attribute is present
	assert.Equal(t, "", elem.GetAttribute("readonly"))
}

func TestWithPlaceholder(t *testing.T) {
	input := bs.Input("testfield", bs.WithPlaceholder("Enter text"))

	elem := input.Element()
	assert.NotNil(t, elem)

	// Check placeholder attribute
	assert.Equal(t, "Enter text", elem.GetAttribute("placeholder"))
}

func TestWithValue(t *testing.T) {
	input := bs.Input("testfield", bs.WithValue("initial"))

	elem := input.Element()
	assert.NotNil(t, elem)

	// Check value attribute
	assert.Equal(t, "initial", elem.GetAttribute("value"))
}

func TestCombinedOptions(t *testing.T) {
	input := bs.Input("testfield",
		bs.WithReadOnly(),
		bs.WithPlaceholder("Enter text"),
		bs.WithValue("initial"))

	elem := input.Element()
	assert.NotNil(t, elem)

	// Check all attributes are present
	assert.Equal(t, "", elem.GetAttribute("readonly"))
	assert.Equal(t, "Enter text", elem.GetAttribute("placeholder"))
	assert.Equal(t, "initial", elem.GetAttribute("value"))
}

func TestNumberInputWithOptions(t *testing.T) {
	input := bs.NumberInput("quantity",
		bs.WithPlaceholder("Enter a number"),
		bs.WithValue("10"))

	elem := input.Element()
	assert.NotNil(t, elem)

	// Check type is still "number"
	assert.Equal(t, "number", elem.GetAttribute("type"))

	// Check additional options were applied
	assert.Equal(t, "Enter a number", elem.GetAttribute("placeholder"))
	assert.Equal(t, "10", elem.GetAttribute("value"))
}
