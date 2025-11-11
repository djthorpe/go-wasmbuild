package mvc

import (
	"fmt"
	"strings"
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

type testComponent struct {
	View
	internal *view
}

func (t *testComponent) SetView(v View) {
	t.View = v
	if internal, ok := v.(*view); ok {
		t.internal = internal
	}
}

type testComponentWithStructure struct {
	testComponent
}

var _ ViewWithHeaderFooter = (*testComponentWithStructure)(nil)

func (t *testComponentWithStructure) Header(children ...any) ViewWithHeaderFooter {
	if t.internal == nil {
		panic("testComponentWithStructure: missing internal view")
	}
	return t.internal.Header(children...)
}

func (t *testComponentWithStructure) Footer(children ...any) ViewWithHeaderFooter {
	if t.internal == nil {
		panic("testComponentWithStructure: missing internal view")
	}
	return t.internal.Footer(children...)
}

type testComponentWithCaption struct {
	testComponent
}

var _ ViewWithCaption = (*testComponentWithCaption)(nil)

func (t *testComponentWithCaption) Caption(children ...any) ViewWithCaption {
	if t.internal == nil {
		panic("testComponentWithCaption: missing internal view")
	}
	return t.internal.Caption(children...)
}

func TestRegisterView(t *testing.T) {
	// Save original views map and restore it after test
	originalViews := views
	defer func() {
		views = originalViews
	}()

	// Reset views for testing
	views = make(map[string]ViewConstructorFunc, 50)

	tests := []struct {
		name        string
		viewName    string
		constructor ViewConstructorFunc
		shouldPanic bool
	}{
		{
			name:        "register new view",
			viewName:    "test-view",
			constructor: nil,
			shouldPanic: false,
		},
		{
			name:        "register duplicate view",
			viewName:    "test-view",
			constructor: nil,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("RegisterView() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
			}()
			RegisterView(tt.viewName, tt.constructor)
		})
	}
}

func TestDataComponentAttrKey(t *testing.T) {
	if DataComponentAttrKey != "data-wasmbuild" {
		t.Errorf("DataComponentAttrKey = %v, want %v", DataComponentAttrKey, "data-wasmbuild")
	}
}

func TestNewView(t *testing.T) {
	// Save and restore original views
	originalViews := views
	defer func() {
		views = originalViews
	}()

	tests := []struct {
		name        string
		viewName    string
		tagName     string
		setup       func()
		shouldPanic bool
		panicMsg    string
	}{
		{
			name:     "create registered view",
			viewName: "test-view",
			tagName:  "DIV",
			setup: func() {
				views = make(map[string]ViewConstructorFunc, 50)
				RegisterView("test-view", nil)
			},
			shouldPanic: false,
		},
		{
			name:     "create unregistered view",
			viewName: "unregistered-view",
			tagName:  "DIV",
			setup: func() {
				views = make(map[string]ViewConstructorFunc, 50)
			},
			shouldPanic: true,
			panicMsg:    "view not registered",
		},
		{
			name:     "use reserved view name",
			viewName: componentPartHeader,
			tagName:  "DIV",
			setup: func() {
				views = make(map[string]ViewConstructorFunc, 50)
				views[componentPartHeader] = nil
			},
			shouldPanic: true,
			panicMsg:    "reserved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("NewView() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
				if r != nil && tt.panicMsg != "" {
					// Check panic message contains expected text
					if msg, ok := r.(string); ok {
						if !strings.Contains(msg, tt.panicMsg) {
							t.Errorf("NewView() panic message = %v, want to contain %v", msg, tt.panicMsg)
						}
					}
				}
			}()

			v := NewView(new(testComponent), tt.viewName, tt.tagName)
			if !tt.shouldPanic {
				if v == nil {
					t.Fatal("NewView() returned nil")
				}
				if v.Name() != tt.viewName {
					t.Errorf("NewView().Name() = %v, want %v", v.Name(), tt.viewName)
				}
			}
		})
	}
}

func TestNewViewWithHeaderCreatesBody(t *testing.T) {
	originalViews := views
	defer func() {
		views = originalViews
	}()

	views = make(map[string]ViewConstructorFunc, 50)
	RegisterView("test-struct-view", nil)

	header := elementFactory("header")
	component := new(testComponentWithStructure)
	result := NewViewEx(component, "test-struct-view", "div", header, nil, nil, nil)
	if result == nil {
		t.Fatal("NewViewEx() returned nil")
	}

	internal, ok := component.View.(*view)
	if !ok {
		t.Fatalf("expected underlying view type, got %T", component.View)
	}
	if internal.body == nil {
		t.Fatal("body should not be nil")
	}
	if internal.body == internal.root {
		t.Fatal("body should not be the root when structural parts are provided")
	}
	if attr := internal.body.GetAttribute(DataComponentAttrKey); attr != componentPartBody {
		t.Fatalf("body attribute = %q, want %q", attr, componentPartBody)
	}

	component.View.Content("hello")
	children := component.View.Root().ChildNodes()
	if len(children) < 2 {
		t.Fatalf("expected header and body children, got %d", len(children))
	}
	firstChild, ok := children[0].(Element)
	if !ok {
		t.Fatalf("first child is not an element: %T", children[0])
	}
	if attr := firstChild.GetAttribute(DataComponentAttrKey); attr != componentPartHeader {
		t.Fatalf("first child attribute = %q, want %q", attr, componentPartHeader)
	}
}

func TestNewViewWithElementRequiresBodyForStructuralParts(t *testing.T) {
	originalViews := views
	defer func() {
		views = originalViews
	}()

	views = make(map[string]ViewConstructorFunc, 50)
	RegisterView("test-struct-view", nil)

	root := elementFactory("div")
	root.SetAttribute(DataComponentAttrKey, "test-struct-view")
	header := elementFactory("header")
	header.SetAttribute(DataComponentAttrKey, componentPartHeader)
	root.AppendChild(header)

	component := new(testComponent)
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic due to missing body component")
		}
	}()

	NewViewWithElement(component, root)
}

func TestViewHeaderMethod(t *testing.T) {
	originalViews := views
	defer func() {
		views = originalViews
	}()

	views = make(map[string]ViewConstructorFunc, 50)
	RegisterView("test-struct-view", nil)

	component := new(testComponentWithStructure)
	view := NewView(component, "test-struct-view", "div")
	if view == nil {
		t.Fatal("NewView returned nil")
	}

	component.Header("Title")

	if component.internal.header == nil {
		t.Fatal("header element was not created")
	}
	if attr := component.internal.header.GetAttribute(DataComponentAttrKey); attr != componentPartHeader {
		t.Fatalf("header attribute = %q, want %q", attr, componentPartHeader)
	}
	if component.internal.header.InnerHTML() != "Title" {
		t.Fatalf("header content = %q, want %q", component.internal.header.InnerHTML(), "Title")
	}

	if component.internal.body == nil || component.internal.body == component.internal.root {
		t.Fatal("body element should be created when using header")
	}

	children := component.internal.root.ChildNodes()
	if len(children) < 2 {
		t.Fatalf("expected at least header and body nodes, got %d", len(children))
	}
	if !children[0].Equals(component.internal.header) {
		t.Fatal("header should be the first child of the root")
	}

	component.Header("Updated")
	if component.internal.header.InnerHTML() != "Updated" {
		t.Fatalf("header content after update = %q, want %q", component.internal.header.InnerHTML(), "Updated")
	}
}

func TestViewFooterMethod(t *testing.T) {
	originalViews := views
	defer func() {
		views = originalViews
	}()

	views = make(map[string]ViewConstructorFunc, 50)
	RegisterView("test-struct-view", nil)

	component := new(testComponentWithStructure)
	NewView(component, "test-struct-view", "div")

	component.Footer("Footer")

	if component.internal.footer == nil {
		t.Fatal("footer element was not created")
	}
	if attr := component.internal.footer.GetAttribute(DataComponentAttrKey); attr != componentPartFooter {
		t.Fatalf("footer attribute = %q, want %q", attr, componentPartFooter)
	}
	if component.internal.footer.InnerHTML() != "Footer" {
		t.Fatalf("footer content = %q, want %q", component.internal.footer.InnerHTML(), "Footer")
	}

	if component.internal.body == nil || component.internal.body == component.internal.root {
		t.Fatal("body element should be created when using footer")
	}

	component.Footer("New Footer")
	if component.internal.footer.InnerHTML() != "New Footer" {
		t.Fatalf("footer content after update = %q, want %q", component.internal.footer.InnerHTML(), "New Footer")
	}

	children := component.internal.root.ChildNodes()
	if len(children) < 2 {
		t.Fatalf("expected at least body and footer nodes, got %d", len(children))
	}
	if !children[len(children)-1].Equals(component.internal.footer) {
		t.Fatal("footer should be the last structural child of the root")
	}
}

func TestViewCaptionMethod(t *testing.T) {
	originalViews := views
	defer func() {
		views = originalViews
	}()

	views = make(map[string]ViewConstructorFunc, 50)
	RegisterView("test-caption-view", nil)

	component := new(testComponentWithCaption)
	NewView(component, "test-caption-view", "div")

	component.Caption("Summary")

	if component.internal.caption == nil {
		t.Fatal("caption element was not created")
	}
	if attr := component.internal.caption.GetAttribute(DataComponentAttrKey); attr != componentPartCaption {
		t.Fatalf("caption attribute = %q, want %q", attr, componentPartCaption)
	}
	if component.internal.caption.InnerHTML() != "Summary" {
		t.Fatalf("caption content = %q, want %q", component.internal.caption.InnerHTML(), "Summary")
	}

	if component.internal.body == nil || component.internal.body == component.internal.root {
		t.Fatal("body element should be created when using caption")
	}

	component.Caption("Updated Summary")
	if component.internal.caption.InnerHTML() != "Updated Summary" {
		t.Fatalf("caption content after update = %q, want %q", component.internal.caption.InnerHTML(), "Updated Summary")
	}

	children := component.internal.root.ChildNodes()
	if !children[len(children)-1].Equals(component.internal.caption) {
		t.Fatal("caption should be the last child of the root")
	}
}

func TestAppContentAppendsElements(t *testing.T) {
	appView := New()
	if appView == nil {
		t.Fatal("New returned nil")
	}

	result := appView.Content(HTML("h1", WithClass("text-center")))
	if result == nil {
		t.Fatal("Content returned nil")
	}

	app, ok := result.(*app)
	if !ok {
		t.Fatalf("expected Content to return *app, got %T", result)
	}

	children := app.root.ChildNodes()
	if len(children) == 0 {
		t.Fatal("expected body to have child nodes")
	}

	last := children[len(children)-1]
	childElement, ok := last.(Element)
	if !ok {
		t.Fatalf("expected last child to be Element, got %T", last)
	}
	if childElement.TagName() != "H1" {
		t.Fatalf("expected last child to be H1, got %s", childElement.TagName())
	}

	if html := fmt.Sprint(result); !strings.Contains(strings.ToLower(html), "<h1") {
		t.Fatalf("expected string output to include <h1, got %s", html)
	}
}

func TestViewName(t *testing.T) {
	v := &view{name: "test-component"}
	if got := v.Name(); got != "test-component" {
		t.Errorf("Name() = %v, want %v", got, "test-component")
	}
}

func TestViewString(t *testing.T) {
	// This test verifies the String() method exists
	// We can't test actual HTML output without a real DOM environment
	t.Skip("String() requires DOM environment - tested in integration tests")
}

func TestNodeFromAny(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		shouldPanic bool
		wantType    string
	}{
		{
			name:        "unsupported type - int",
			input:       123,
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - bool",
			input:       true,
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - slice",
			input:       []string{"a", "b"},
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - map",
			input:       map[string]string{"key": "value"},
			shouldPanic: true,
			wantType:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("NodeFromAny() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
			}()

			_ = NodeFromAny(tt.input)
		})
	}
}

func TestViewChaining(t *testing.T) {
	// Test that methods return the view for chaining (compile-time check)
	t.Run("method chaining pattern", func(t *testing.T) {
		// This is a compile-time check that the methods have the right signatures
		// for method chaining. We can't actually test them without a DOM.

		// These should all compile - they return View for chaining
		_ = func(v View) View { return v.Content("test") }
		_ = func(v View) View { return v.Append("test") }
		_ = func(v View) View { return v.AddEventListener("click", nil) }
		_ = func(v View) View { return v.Opts() }

		// Body doesn't return View, so it breaks the chain intentionally
		_ = func(v View) { v.Body(nil) }
	})
}

func TestViewInterfaceImplementation(t *testing.T) {
	// Compile-time check that view implements View interface
	var _ View = (*view)(nil)

	// Runtime check
	v := &view{name: "test"}
	if _, ok := any(v).(View); !ok {
		t.Error("view does not implement View interface")
	}
}

func TestViewConstructorFunc(t *testing.T) {
	// Test that ViewConstructorFunc type works as expected
	var constructor ViewConstructorFunc = func(e Element) View {
		return &view{name: "constructed"}
	}

	if constructor == nil {
		t.Error("ViewConstructorFunc should not be nil")
	}
}

func TestViewStructFields(t *testing.T) {
	v := &view{
		name: "test-view",
		root: nil,
		body: nil,
	}

	if v.name != "test-view" {
		t.Errorf("view.name = %v, want %v", v.name, "test-view")
	}

	if v.root != nil {
		t.Errorf("view.root should be nil")
	}

	if v.body != nil {
		t.Errorf("view.body should be nil")
	}
}
