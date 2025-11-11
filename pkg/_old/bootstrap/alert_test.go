package bootstrap_test

import (
	"strings"
	"testing"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	"github.com/stretchr/testify/assert"
)

func TestAlert_Basic(t *testing.T) {
	alert := bs.Alert()

	assert.NotNil(t, alert, "Alert should not be nil")
	assert.NotNil(t, alert.Element(), "Alert element should not be nil")
	assert.Equal(t, dom.ELEMENT_NODE, alert.Element().NodeType(), "Alert should be an element node")
	assert.Equal(t, "DIV", alert.Element().TagName(), "Alert should be a div element")
}

func TestAlert_DefaultClass(t *testing.T) {
	alert := bs.Alert()
	element := alert.Element()

	assert.True(t, element.HasAttribute("class"), "Alert should have class attribute")
	classList := element.ClassList()
	assert.NotNil(t, classList, "Alert should have class list")
	assert.True(t, classList.Contains("alert"), "Alert should contain 'alert' class")
}

func TestAlert_RoleAttribute(t *testing.T) {
	alert := bs.Alert()
	element := alert.Element()

	assert.True(t, element.HasAttribute("role"), "Alert should have role attribute")
	assert.Equal(t, "alert", element.GetAttribute("role"), "Alert should have role='alert'")
}

func TestAlert_WithAlertColor(t *testing.T) {
	tests := []struct {
		name          string
		color         bs.Color
		expectedClass string
	}{
		{"primary alert", bs.PRIMARY, "alert-primary"},
		{"secondary alert", bs.SECONDARY, "alert-secondary"},
		{"success alert", bs.SUCCESS, "alert-success"},
		{"danger alert", bs.DANGER, "alert-danger"},
		{"warning alert", bs.WARNING, "alert-warning"},
		{"info alert", bs.INFO, "alert-info"},
		{"light alert", bs.LIGHT, "alert-light"},
		{"dark alert", bs.DARK, "alert-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := bs.Alert(bs.WithColor(tt.color))
			element := alert.Element()

			classList := element.ClassList()
			assert.True(t, classList.Contains("alert"), "Alert should contain 'alert' class")
			assert.True(t, classList.Contains(tt.expectedClass), "Alert should contain '%s' class", tt.expectedClass)
		})
	}
}

func TestAlert_OuterHTML(t *testing.T) {
	tests := []struct {
		name            string
		constructor     func() dom.Component
		expectedClasses []string
		expectedRole    string
	}{
		{
			name:            "default alert",
			constructor:     func() dom.Component { return bs.Alert() },
			expectedClasses: []string{"alert"},
			expectedRole:    "alert",
		},
		{
			name:            "primary alert",
			constructor:     func() dom.Component { return bs.Alert(bs.WithColor(bs.PRIMARY)) },
			expectedClasses: []string{"alert", "alert-primary"},
			expectedRole:    "alert",
		},
		{
			name:            "danger alert",
			constructor:     func() dom.Component { return bs.Alert(bs.WithColor(bs.DANGER)) },
			expectedClasses: []string{"alert", "alert-danger"},
			expectedRole:    "alert",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := tt.constructor()
			element := alert.Element()

			// Verify tag name
			assert.Equal(t, "DIV", element.TagName(), "Alert should be a DIV element")

			// Verify role attribute
			assert.Equal(t, tt.expectedRole, element.GetAttribute("role"), "Alert should have correct role attribute")

			// Verify classes
			classList := element.ClassList()
			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Alert should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}

			// Verify OuterHTML contains expected parts (without checking attribute order)
			outerHTML := strings.ToLower(element.OuterHTML())
			assert.Contains(t, outerHTML, "<div", "OuterHTML should start with <div")
			assert.Contains(t, outerHTML, "</div>", "OuterHTML should end with </div>")
			assert.Contains(t, outerHTML, `role="alert"`, "OuterHTML should contain role attribute")
			for _, expectedClass := range tt.expectedClasses {
				assert.Contains(t, outerHTML, expectedClass, "OuterHTML should contain class '%s'", expectedClass)
			}
		})
	}
}

func TestAlert_WithAdditionalClasses(t *testing.T) {
	tests := []struct {
		name            string
		options         []bs.Opt
		expectedClasses []string
	}{
		{
			name:            "alert with single additional class",
			options:         []bs.Opt{bs.WithClass("my-custom-class")},
			expectedClasses: []string{"alert", "my-custom-class"},
		},
		{
			name:            "alert with multiple additional classes",
			options:         []bs.Opt{bs.WithClass("class1", "class2", "class3")},
			expectedClasses: []string{"alert", "class1", "class2", "class3"},
		},
		{
			name:            "primary alert with additional classes",
			options:         []bs.Opt{bs.WithColor(bs.PRIMARY), bs.WithClass("custom-alert")},
			expectedClasses: []string{"alert", "alert-primary", "custom-alert"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := bs.Alert(tt.options...)
			element := alert.Element()
			classList := element.ClassList()

			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"Alert should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}
		})
	}
}

func TestAlert_ComponentInterface(t *testing.T) {
	alert := bs.Alert()

	// Test that alert implements Component interface
	var component dom.Component = alert
	assert.NotNil(t, component, "Alert should implement Component interface")
	assert.NotNil(t, component.Element(), "Component Element() should return an element")
}

func TestAlert_AppendText(t *testing.T) {
	alert := bs.Alert()
	alert.Append("A simple primary alert—check it out!")

	element := alert.Element()
	assert.Equal(t, "A simple primary alert—check it out!", element.TextContent(), "Alert should contain text")
}

func TestAlert_AppendMultipleTextNodes(t *testing.T) {
	alert := bs.Alert()
	alert.Append("Hello", " ", "World")

	element := alert.Element()
	assert.Equal(t, "Hello World", element.TextContent(), "Alert should contain text 'Hello World'")
}

func TestAlert_ChainedAppends(t *testing.T) {
	alert := bs.Alert().Append("Alert").Append(" ").Append("Message")

	element := alert.Element()
	assert.Equal(t, "Alert Message", element.TextContent(), "Alert should contain text 'Alert Message'")
}

func TestAlert_ElementProperties(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "tag name is DIV",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.Equal(t, "DIV", alert.Element().TagName())
			},
		},
		{
			name: "node type is element",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.Equal(t, dom.ELEMENT_NODE, alert.Element().NodeType())
			},
		},
		{
			name: "node name is DIV",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.Equal(t, "DIV", alert.Element().NodeName())
			},
		},
		{
			name: "has role attribute",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.Equal(t, "alert", alert.Element().GetAttribute("role"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestAlert_ClassListOperations(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "initial state",
			test: func(t *testing.T) {
				alert := bs.Alert()
				classList := alert.Element().ClassList()
				assert.True(t, classList.Contains("alert"))
			},
		},
		{
			name: "add classes",
			test: func(t *testing.T) {
				alert := bs.Alert()
				classList := alert.Element().ClassList()
				classList.Add("new-class", "another-class")
				assert.True(t, classList.Contains("alert"))
				assert.True(t, classList.Contains("new-class"))
				assert.True(t, classList.Contains("another-class"))
			},
		},
		{
			name: "remove classes",
			test: func(t *testing.T) {
				alert := bs.Alert(bs.WithClass("removable"))
				classList := alert.Element().ClassList()
				assert.True(t, classList.Contains("removable"))
				classList.Remove("removable")
				assert.False(t, classList.Contains("removable"))
				assert.True(t, classList.Contains("alert"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestAlert_MultipleAlerts(t *testing.T) {
	alert1 := bs.Alert(bs.WithColor(bs.PRIMARY)).Append("Primary alert")
	alert2 := bs.Alert(bs.WithColor(bs.DANGER)).Append("Danger alert")
	alert3 := bs.Alert(bs.WithColor(bs.SUCCESS)).Append("Success alert")

	assert.Equal(t, "Primary alert", alert1.Element().TextContent())
	assert.Equal(t, "Danger alert", alert2.Element().TextContent())
	assert.Equal(t, "Success alert", alert3.Element().TextContent())

	assert.True(t, alert1.Element().ClassList().Contains("alert-primary"))
	assert.True(t, alert2.Element().ClassList().Contains("alert-danger"))
	assert.True(t, alert3.Element().ClassList().Contains("alert-success"))
}

func TestAlert_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "empty options",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.NotNil(t, alert)
				assert.True(t, alert.Element().ClassList().Contains("alert"))
			},
		},
		{
			name: "empty text content",
			test: func(t *testing.T) {
				alert := bs.Alert()
				assert.Equal(t, "", alert.Element().TextContent())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestAlert_DismissibleAlert(t *testing.T) {
	alert := bs.DismissibleAlert(bs.WithColor(bs.WARNING))

	classList := alert.Element().ClassList()
	assert.True(t, classList.Contains("alert"))
	assert.True(t, classList.Contains("alert-warning"))
	assert.True(t, classList.Contains("alert-dismissible"))
	assert.True(t, classList.Contains("fade"))
	assert.True(t, classList.Contains("show"))

	// Check that close button was automatically added
	element := alert.Element()
	assert.True(t, element.HasChildNodes(), "DismissibleAlert should have child nodes (close button)")

	// The close button should be the first child
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "DismissibleAlert should have a close button as first child")

	// Verify it's a button element with correct attributes
	if firstChild != nil {
		buttonElement, ok := firstChild.(dom.Element)
		assert.True(t, ok, "First child should be an Element")
		if ok {
			assert.Equal(t, "BUTTON", buttonElement.TagName())
			assert.Equal(t, "button", buttonElement.GetAttribute("type"))
			assert.Equal(t, "btn-close", buttonElement.GetAttribute("class"))
			assert.Equal(t, "alert", buttonElement.GetAttribute("data-bs-dismiss"))
			assert.Equal(t, "Close", buttonElement.GetAttribute("aria-label"))
		}
	}
}

func TestAlert_DismissibleAlertConstructor(t *testing.T) {
	tests := []struct {
		name            string
		color           bs.Color
		text            string
		expectedClasses []string
	}{
		{
			name:            "primary dismissible alert",
			color:           bs.PRIMARY,
			text:            "This is a dismissible alert",
			expectedClasses: []string{"alert", "alert-primary", "alert-dismissible", "fade", "show"},
		},
		{
			name:            "danger dismissible alert",
			color:           bs.DANGER,
			text:            "Error occurred!",
			expectedClasses: []string{"alert", "alert-danger", "alert-dismissible", "fade", "show"},
		},
		{
			name:            "success dismissible alert",
			color:           bs.SUCCESS,
			text:            "Success!",
			expectedClasses: []string{"alert", "alert-success", "alert-dismissible", "fade", "show"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := bs.DismissibleAlert(bs.WithColor(tt.color)).Append(tt.text)

			classList := alert.Element().ClassList()
			for _, expectedClass := range tt.expectedClasses {
				assert.True(t, classList.Contains(expectedClass),
					"DismissibleAlert should contain class '%s', actual classes: %v", expectedClass, classList.Values())
			}

			// Verify the text content (note: includes close button which has no text)
			assert.Contains(t, alert.Element().TextContent(), tt.text, "Alert should contain the expected text")

			// Verify close button is present
			element := alert.Element()
			assert.True(t, element.HasChildNodes(), "DismissibleAlert should have child nodes")

			// The close button should be the first child
			firstChild := element.FirstChild()
			assert.NotNil(t, firstChild, "DismissibleAlert should have a close button as first child")

			if firstChild != nil {
				buttonElement, ok := firstChild.(dom.Element)
				assert.True(t, ok, "First child should be an Element (button)")
				if ok {
					assert.Equal(t, "BUTTON", buttonElement.TagName(), "First child should be a BUTTON element")
					assert.Equal(t, "alert", buttonElement.GetAttribute("data-bs-dismiss"), "Button should have data-bs-dismiss='alert'")
				}
			}
		})
	}
}

func TestAlert_WithBorder(t *testing.T) {
	alert := bs.Alert(
		bs.WithColor(bs.PRIMARY),
		bs.WithBorder(bs.BorderAll, bs.PRIMARY),
	)

	classList := alert.Element().ClassList()
	assert.True(t, classList.Contains("alert"))
	assert.True(t, classList.Contains("alert-primary"))
	assert.True(t, classList.Contains("border"))
	assert.True(t, classList.Contains("border-primary"))
}

func TestAlert_WithPadding(t *testing.T) {
	alert := bs.Alert(
		bs.WithColor(bs.INFO),
		bs.WithPadding(bs.PaddingAll, 3),
	)

	classList := alert.Element().ClassList()
	assert.True(t, classList.Contains("alert"))
	assert.True(t, classList.Contains("alert-info"))
	assert.True(t, classList.Contains("p-3"))
}

func TestAlert_ComplexCombination(t *testing.T) {
	alert := bs.Alert(
		bs.WithColor(bs.SUCCESS),
		bs.WithBorder(bs.BorderAll, bs.SUCCESS),
		bs.WithPadding(bs.PaddingAll, 4),
		bs.WithClass("custom-alert", "shadow"),
	).Append("Complex alert with multiple options")

	classList := alert.Element().ClassList()
	expectedClasses := []string{"alert", "alert-success", "border", "border-success", "p-4", "custom-alert", "shadow"}

	for _, expectedClass := range expectedClasses {
		assert.True(t, classList.Contains(expectedClass),
			"Alert should contain class '%s', actual classes: %v", expectedClass, classList.Values())
	}

	assert.Equal(t, "Complex alert with multiple options", alert.Element().TextContent())
	assert.Equal(t, "alert", alert.Element().GetAttribute("role"))
}

func TestAlert_WithAlertColorExamples(t *testing.T) {
	// Test examples from Bootstrap documentation
	tests := []struct {
		name          string
		color         bs.Color
		text          string
		expectedClass string
	}{
		{"Primary", bs.PRIMARY, "A simple primary alert—check it out!", "alert-primary"},
		{"Secondary", bs.SECONDARY, "A simple secondary alert—check it out!", "alert-secondary"},
		{"Success", bs.SUCCESS, "A simple success alert—check it out!", "alert-success"},
		{"Danger", bs.DANGER, "A simple danger alert—check it out!", "alert-danger"},
		{"Warning", bs.WARNING, "A simple warning alert—check it out!", "alert-warning"},
		{"Info", bs.INFO, "A simple info alert—check it out!", "alert-info"},
		{"Light", bs.LIGHT, "A simple light alert—check it out!", "alert-light"},
		{"Dark", bs.DARK, "A simple dark alert—check it out!", "alert-dark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			alert := bs.Alert(bs.WithColor(tt.color)).Append(tt.text)

			classList := alert.Element().ClassList()
			assert.True(t, classList.Contains("alert"))
			assert.True(t, classList.Contains(tt.expectedClass))
			assert.Equal(t, tt.text, alert.Element().TextContent())
			assert.Equal(t, "alert", alert.Element().GetAttribute("role"))
		})
	}
}

func TestAlert_Heading(t *testing.T) {
	alert := bs.Alert(bs.WithColor(bs.SUCCESS))
	alert.Heading("Well done!")
	alert.Append("You successfully read this important alert message.")

	element := alert.Element()
	assert.True(t, element.HasChildNodes(), "Alert should have child nodes")

	// The heading should be the first child
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "Alert should have a heading as first child")

	if firstChild != nil {
		headingElement, ok := firstChild.(dom.Element)
		assert.True(t, ok, "First child should be an Element")
		if ok {
			assert.Equal(t, "H4", headingElement.TagName(), "Heading should be an H4 element")
			assert.Equal(t, "alert-heading", headingElement.GetAttribute("class"), "Heading should have 'alert-heading' class")
			assert.Equal(t, "Well done!", headingElement.TextContent(), "Heading should contain the text")
		}
	}

	// Check that the alert text comes after the heading
	assert.Contains(t, element.TextContent(), "Well done!")
	assert.Contains(t, element.TextContent(), "You successfully read this important alert message.")
}

func TestAlert_HeadingWithMultipleChildren(t *testing.T) {
	alert := bs.Alert(bs.WithColor(bs.INFO))
	alert.Heading("Info: ", "Multiple parts")
	alert.Append("Additional content")

	element := alert.Element()
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "Alert should have a heading")

	if firstChild != nil {
		headingElement, ok := firstChild.(dom.Element)
		assert.True(t, ok, "First child should be an Element")
		if ok {
			assert.Equal(t, "H4", headingElement.TagName())
			assert.Equal(t, "alert-heading", headingElement.GetAttribute("class"))
			assert.Equal(t, "Info: Multiple parts", headingElement.TextContent())
		}
	}
}

func TestAlert_HeadingWithComponent(t *testing.T) {
	alert := bs.Alert(bs.WithColor(bs.WARNING))
	icon := bs.Icon("exclamation-triangle-fill", bs.WithMargin(bs.END, 2))
	alert.Heading(icon, "Warning!")
	alert.Append("This is a warning message.")

	element := alert.Element()
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "Alert should have a heading")

	if firstChild != nil {
		headingElement, ok := firstChild.(dom.Element)
		assert.True(t, ok, "First child should be an Element")
		if ok {
			assert.Equal(t, "H4", headingElement.TagName())
			assert.Equal(t, "alert-heading", headingElement.GetAttribute("class"))
			assert.Contains(t, headingElement.TextContent(), "Warning!")
		}
	}
}

func TestAlert_DismissibleAlertWithHeading(t *testing.T) {
	alert := bs.DismissibleAlert(bs.WithColor(bs.DANGER))
	alert.Heading("Error!")
	alert.Append("Something went wrong.")

	element := alert.Element()
	assert.True(t, element.HasChildNodes(), "Dismissible alert should have child nodes")

	// The close button should be the first child
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "Dismissible alert should have a close button")

	if firstChild != nil {
		buttonElement, ok := firstChild.(dom.Element)
		assert.True(t, ok, "First child should be the close button")
		if ok {
			assert.Equal(t, "BUTTON", buttonElement.TagName())
		}
	}

	// The heading should be the second child (after close button)
	secondChild := firstChild.NextSibling()
	assert.NotNil(t, secondChild, "Dismissible alert should have a heading as second child")

	if secondChild != nil {
		headingElement, ok := secondChild.(dom.Element)
		assert.True(t, ok, "Second child should be an Element (heading)")
		if ok {
			assert.Equal(t, "H4", headingElement.TagName())
			assert.Equal(t, "alert-heading", headingElement.GetAttribute("class"))
			assert.Equal(t, "Error!", headingElement.TextContent())
		}
	}
}

func TestAlert_HeadingMethodChaining(t *testing.T) {
	alert := bs.Alert(bs.WithColor(bs.SUCCESS)).
		Heading("Success!").
		Append("Operation completed successfully.")

	element := alert.Element()
	assert.Contains(t, element.TextContent(), "Success!")
	assert.Contains(t, element.TextContent(), "Operation completed successfully.")

	// Verify the heading is present
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild)

	if firstChild != nil {
		headingElement, ok := firstChild.(dom.Element)
		assert.True(t, ok)
		if ok {
			assert.Equal(t, "H4", headingElement.TagName())
			assert.Equal(t, "alert-heading", headingElement.GetAttribute("class"))
		}
	}
}

func TestAlert_HeadingEmptyAlert(t *testing.T) {
	alert := bs.Alert(bs.WithColor(bs.PRIMARY))
	alert.Heading("Empty Alert")

	element := alert.Element()
	firstChild := element.FirstChild()
	assert.NotNil(t, firstChild, "Alert should have a heading")

	if firstChild != nil {
		headingElement, ok := firstChild.(dom.Element)
		assert.True(t, ok)
		if ok {
			assert.Equal(t, "H4", headingElement.TagName())
			assert.Equal(t, "Empty Alert", headingElement.TextContent())
		}
	}
}
