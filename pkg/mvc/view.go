package mvc

import (
	"fmt"
	"os"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// View represents a UI component in the interface
type View interface {
	// Return the view name
	Name() string

	// Return the self container for the view
	Self() View

	// Return the view ID, if set
	ID() string

	// Return the view's root element
	Root() dom.Element

	// Append a view or element to the view's body, and set the body
	Body(any) View

	// Set the body's content to the given text, Element or View children
	// If no arguments are given, the content is cleared
	Content(children ...any) View

	// Append text, Element or View children at the bottom of the view body
	Append(children ...any) View

	// Add an event listener to the view's root element
	AddEventListener(event string, handler func(dom.Event)) View

	// Set options on the view
	Opts(opts ...Opt) View
}

// ViewWithState represents a UI component with active and disabled states
type ViewWithState interface {
	View

	// Indicates whether the view is active
	Active() bool

	// Indicates whether the view is disabled
	Disabled() bool
}

// ViewWithGroupState represents a UI component with a group of active and disabled states
type ViewWithGroupState interface {
	View

	// Returns any elements which are active
	Active() []dom.Element

	// Returns any elements which are disabled
	Disabled() []dom.Element
}

// ViewWithCaption represents a UI component with a header and footer
type ViewWithCaption interface {
	View

	// Sets the caption of the view and returns the view
	Caption(...any) ViewWithCaption
}

// ViewWithHeaderFooter represents a UI component with a header and footer
type ViewWithHeaderFooter interface {
	View

	// Sets the header and returns the view
	Header(...any) ViewWithHeaderFooter

	// Returns the footer element
	Footer(...any) ViewWithHeaderFooter
}

// ViewWithHeaderFooter represents a UI component with a header and footer
type ViewWithVisibility interface {
	View

	// Returns true if the view is visible
	Visible() bool

	// Makes the view visible and returns the view
	Show() ViewWithVisibility

	// Hides the view and returns the view
	Hide() ViewWithVisibility
}

// View represents a UI component in the interface
type ViewWithSelf interface {
	View
	SetView(view View)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE TYPES

// Implementation of View interface
type view struct {
	self View
	name string
	root dom.Element
	body dom.Element
}

// Ensure that view implements View interface
var _ View = (*view)(nil)

// Constructor function for views
type ViewConstructorFunc func(dom.Element) View

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// The attribute key which identifies a wasmbuild view
	DataComponentAttrKey = "data-wasmbuild"
)

var (
	// All the registered views
	views = make(map[string]ViewConstructorFunc, 50)
)

// RegisterView registers a view constructor function for a given name,
// so that the view can be created on-demand
func RegisterView(name string, constructor ViewConstructorFunc) {
	if _, exists := views[name]; exists {
		panic("View already registered: " + name)
	}
	views[name] = constructor
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new empty view, applying any options to it
func NewView(self View, name string, tagName string, opts ...Opt) View {
	if _, exists := views[name]; !exists {
		panic(fmt.Sprintf("NewView: View not registered %q", name))
	}

	// Create the view
	v := &view{
		self: self,
		name: name,
		root: elementFactory(tagName),
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", name))
	} else {
		self_.SetView(v)
	}

	// Set the data-wasmbuild-component attribute
	v.root.SetAttribute(DataComponentAttrKey, name)

	// Apply options to the view
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}

	// Return the view
	return v.self
}

// Create view from an existing element, applying any options to it
func NewViewWithElement(self View, element dom.Element, opts ...Opt) View {
	if element == nil {
		panic("NewViewWithElement: missing element")
	} else if self == nil {
		panic("NewViewWithElement: missing self")
	}

	// Create the view
	v := &view{
		self: self,
		name: element.GetAttribute(DataComponentAttrKey),
		root: element,
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", v.name))
	} else {
		self_.SetView(v)
	}

	// Apply options to the view
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}

	// Return self
	return v.self
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v *view) String() string {
	return v.Root().OuterHTML()
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (v *view) Self() View {
	return v.self
}

func (v *view) Name() string {
	return v.name
}

func (v *view) ID() string {
	return v.root.ID()
}

func (v *view) Root() dom.Element {
	return v.root
}

func (v *view) Body(content any) View {
	node := NodeFromAny(content)
	if element, ok := node.(dom.Element); ok {
		v.body = element
	} else if view, ok := node.(View); ok {
		v.body = view.Root()
	} else {
		panic(fmt.Sprint("view.Body: invalid content type", node.NodeType()))
	}

	// If the body is not already a child of the root, clear the root and set it
	if v.body.ParentNode() == nil {
		v.root.SetInnerHTML("")
		v.root.AppendChild(v.body)
	}

	return v.self
}

func (v *view) Content(children ...any) View {
	// Determine target element
	target := v.body
	if target == nil {
		target = v.root
	}

	// Clear the content
	if len(children) == 0 {
		target.SetInnerHTML("")
		return v
	}

	// Append each child
	return v.Append(children...)
}

func (v *view) Append(children ...any) View {
	target := v.body
	if target == nil {
		target = v.root
	}
	for _, child := range children {
		target.AppendChild(NodeFromAny(child))
	}

	return v.self
}

func (v *view) AddEventListener(event string, handler func(dom.Event)) View {
	v.root.AddEventListener(event, handler)
	return v.self
}

func (v *view) Opts(opts ...Opt) View {
	if err := applyOpts(v.root, opts...); err != nil {
		panic(err)
	}
	return v.self
}

///////////////////////////////////////////////////////////////////////////////
// UTILITY METHODS

// NodeFromAny returns a Node from a string, Element, Tag or View
// or returns nil if the type is unsupported
func NodeFromAny(child any) dom.Node {
	switch c := child.(type) {
	case string:
		return textFactory(c)
	case dom.Element:
		return c
	case dom.Node:
		if c.NodeType() == dom.TEXT_NODE || c.NodeType() == dom.COMMENT_NODE {
			return c
		}
	case View:
		return c.Root()
	}
	panic(dom.ErrInternalAppError.Withf("NodeFromAny: unsupported: %T", child))
}

// ViewFromNode returns a View from a Node, or nil if the type is unsupported
func ViewFromNode(node dom.Node) View {
	if element, ok := node.(dom.Element); ok {
		// Work up the chain until a view is found
		for {
			if view, err := viewFromElement(element); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return nil
			} else if view != nil {
				return view
			}
			element = element.ParentElement()
			if element == nil {
				break
			}
		}

	}
	return nil
}

func viewFromElement(element dom.Element) (View, error) {
	if !element.HasAttribute(DataComponentAttrKey) {
		return nil, nil
	}
	name := element.GetAttribute(DataComponentAttrKey)
	if constructor, exists := views[name]; !exists {
		return nil, dom.ErrInternalAppError.Withf("viewFromElement: no constructor for view %q", name)
	} else if constructor == nil {
		return nil, dom.ErrInternalAppError.Withf("viewFromElement: constructor for view %q is nil", name)
	} else if view := constructor(element); view == nil {
		return nil, dom.ErrInternalAppError.Withf("viewFromElement: constructor for view %q returned nil", name)
	} else {
		return view, nil
	}
}
