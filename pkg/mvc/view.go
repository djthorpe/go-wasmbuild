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

// ViewWithVisibility represents a UI component with the ability to show or hide itself
type ViewWithVisibility interface {
	View

	// Returns true if the view is visible
	Visible() bool

	// Makes the view visible and returns the view
	Show() ViewWithVisibility

	// Hides the view and returns the view
	Hide() ViewWithVisibility
}

// ViewWithSelf represents a UI component that can set its own view
type ViewWithSelf interface {
	View
	SetView(view View)
}

// ViewWithValue represents a UI component that can set and get a value, typically
// for form elements
type ViewWithValue interface {
	View

	// Return the value of the view as a string
	Value() string

	// Set the value of the view as a string
	SetValue(string) ViewWithValue
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE TYPES

// Implementation of View interface
type view struct {
	self    View
	name    string
	root    dom.Element
	body    dom.Element
	header  dom.Element
	footer  dom.Element
	caption dom.Element
}

// Ensure that view implements View interface
var _ View = (*view)(nil)

// Constructor function for views
type ViewConstructorFunc func(dom.Element) View

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// The attribute key which identifies a wasmbuild component
	DataComponentAttrKey = "data-wasmbuild"

	componentPartHeader  = "header"
	componentPartBody    = "body"
	componentPartFooter  = "footer"
	componentPartCaption = "caption"
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
func NewView(self View, name string, tagName string, args ...any) View {
	return NewViewEx(self, name, tagName, nil, nil, nil, nil, args...)
}

// Create a new empty view with a header, footer and caption
func NewViewEx(self View, name string, tagName string, header, body, footer, caption dom.Element, args ...any) View {
	if _, exists := views[name]; !exists {
		panic(fmt.Sprintf("NewView: view not registered %q", name))
	}
	if isComponentPart(name) {
		panic(fmt.Sprintf("NewView: view name %q is reserved", name))
	}

	// Ensure a dedicated body exists whenever structural parts are provided so that
	// later content operations don't trample header, footer or caption nodes.
	if body == nil && (header != nil || footer != nil || caption != nil) {
		body = elementFactory("div")
	}

	// Create the view
	v := &view{
		self:    self,
		name:    name,
		root:    elementFactory(tagName),
		header:  header,
		body:    body,
		footer:  footer,
		caption: caption,
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", name))
	} else {
		self_.SetView(v)
	}

	// Check header, footer, caption
	if v.header != nil || v.footer != nil {
		if _, ok := self.(ViewWithHeaderFooter); !ok {
			panic(fmt.Sprintf("NewView: %v does not implement ViewWithHeaderFooter", name))
		}
	}
	if v.caption != nil {
		if _, ok := self.(ViewWithCaption); !ok {
			panic(fmt.Sprintf("NewView: %v does not implement ViewWithCaption", name))
		}
	}

	// Set the header, body, footer and caption
	if v.header != nil {
		if v.header.IsConnected() {
			panic("NewView: header element is already connected to the DOM")
		}
		markComponentPart(v.header, componentPartHeader)
		v.root.AppendChild(v.header)
	}
	if v.body != nil {
		if v.body.IsConnected() {
			panic("NewView: body element is already connected to the DOM")
		}
		if v.body != v.root {
			markComponentPart(v.body, componentPartBody)
		}
		v.root.AppendChild(v.body)
	} else {
		v.body = v.root
	}
	if v.footer != nil {
		if v.footer.IsConnected() {
			panic("NewView: footer element is already connected to the DOM")
		}
		markComponentPart(v.footer, componentPartFooter)
		v.root.AppendChild(v.footer)
	}
	if v.caption != nil {
		if v.caption.IsConnected() {
			panic("NewView: caption element is already connected to the DOM")
		}
		markComponentPart(v.caption, componentPartCaption)
		v.root.AppendChild(v.caption)
	}

	// Set the component identifier
	v.root.SetAttribute(DataComponentAttrKey, name)

	// Apply options to the view
	opts, content := gatherOpts(args...)
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}

	// Add content to the component
	if len(content) > 0 {
		v.self.Content(content...)
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
	if v.name == "" {
		panic("NewViewWithElement: element missing data-wasmbuild attribute")
	}
	if isComponentPart(v.name) {
		panic(fmt.Sprintf("NewViewWithElement: element uses reserved component value %q", v.name))
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", v.name))
	} else {
		self_.SetView(v)
	}

	// Discover structural elements from data attributes when present
	v.header = findComponentPart(element, componentPartHeader)
	if body := findComponentPart(element, componentPartBody); body != nil {
		v.body = body
	} else {
		v.body = v.root
	}
	v.footer = findComponentPart(element, componentPartFooter)
	v.caption = findComponentPart(element, componentPartCaption)

	if v.body == v.root && (v.header != nil || v.footer != nil || v.caption != nil) {
		panic("NewViewWithElement: element missing body component")
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
	prevBody := v.body
	node := NodeFromAny(content)
	var newBody dom.Element
	if element, ok := node.(dom.Element); ok {
		newBody = element
	} else if view, ok := node.(View); ok {
		newBody = view.Root()
	} else {
		panic(fmt.Sprint("view.Body: invalid content type ", node.NodeType()))
	}

	if prevBody != nil && prevBody.Equals(newBody) {
		return v.self
	}

	v.body = newBody

	// Remove previous body attribute and detach from root if needed
	if prevBody != nil && prevBody != v.root {
		if prevBody.HasAttribute(DataComponentAttrKey) && isComponentPart(prevBody.GetAttribute(DataComponentAttrKey)) {
			prevBody.RemoveAttribute(DataComponentAttrKey)
		}
		if parent := prevBody.ParentNode(); parent != nil && parent.Equals(v.root) {
			v.root.RemoveChild(prevBody)
		}
	}

	if v.body != v.root {
		markComponentPart(v.body, componentPartBody)
	}

	// Attach the body in the correct position if not already attached
	if v.body.ParentNode() == nil {
		if v.footer != nil && v.footer.ParentNode() != nil && v.footer.ParentNode().Equals(v.root) {
			v.root.InsertBefore(v.body, v.footer)
		} else if v.caption != nil && v.caption.ParentNode() != nil && v.caption.ParentNode().Equals(v.root) {
			v.root.InsertBefore(v.body, v.caption)
		} else {
			v.root.AppendChild(v.body)
		}
	}

	return v.self
}

func (v *view) Content(children ...any) View {
	// Determine target element
	target := v.body
	if target == nil {
		target = v.root
	}

	// Clear existing content before appending new children
	target.SetInnerHTML("")
	if len(children) == 0 {
		return v
	}

	// Append each child
	return v.self.Append(children...)
}

// Append appends text, Element or View children at the bottom of the view body
// and returns the view for chaining
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

func (v *view) Header(children ...any) ViewWithHeaderFooter {
	viewWithHeader, ok := v.self.(ViewWithHeaderFooter)
	if !ok {
		panic(fmt.Sprintf("view.Header: view %T does not implement ViewWithHeaderFooter", v.self))
	}

	header := v.ensureHeaderElement()
	v.replaceChildContent(header, children...)

	return viewWithHeader
}

func (v *view) Footer(children ...any) ViewWithHeaderFooter {
	viewWithFooter, ok := v.self.(ViewWithHeaderFooter)
	if !ok {
		panic(fmt.Sprintf("view.Footer: view %T does not implement ViewWithHeaderFooter", v.self))
	}

	footer := v.ensureFooterElement()
	v.replaceChildContent(footer, children...)

	return viewWithFooter
}

func (v *view) Caption(children ...any) ViewWithCaption {
	viewWithCaption, ok := v.self.(ViewWithCaption)
	if !ok {
		panic(fmt.Sprintf("view.Caption: view %T does not implement ViewWithCaption", v.self))
	}

	caption := v.ensureCaptionElement()
	v.replaceChildContent(caption, children...)

	return viewWithCaption
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

// ViewFromEvent returns a View from an Event, or nil if the type is unsupported
func ViewFromEvent(e dom.Event) View {
	if e == nil {
		return nil
	}
	switch element := e.Target().(type) {
	case dom.Element:
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

func (v *view) ensureBodyContainer() dom.Element {
	if v.root == nil {
		panic("view.ensureBodyContainer: missing root element")
	}
	if v.body == nil {
		v.body = v.root
	}
	if v.body != v.root {
		return v.body
	}

	body := elementFactory("div")
	markComponentPart(body, componentPartBody)

	for _, child := range v.root.ChildNodes() {
		if elem, ok := child.(dom.Element); ok {
			if elem.HasAttribute(DataComponentAttrKey) {
				if part := elem.GetAttribute(DataComponentAttrKey); isComponentPart(part) && part != componentPartBody {
					continue
				}
			}
		}
		body.AppendChild(child)
	}

	v.root.AppendChild(body)
	v.body = body

	return v.body
}

func (v *view) ensureHeaderElement() dom.Element {
	if v.header != nil {
		return v.header
	}
	body := v.ensureBodyContainer()
	header := elementFactory("header")
	markComponentPart(header, componentPartHeader)

	if parent := body.ParentNode(); parent != nil && parent.Equals(v.root) {
		v.root.InsertBefore(header, body)
	} else {
		v.root.AppendChild(header)
	}

	v.header = header
	return v.header
}

func (v *view) ensureFooterElement() dom.Element {
	if v.footer != nil {
		return v.footer
	}
	v.ensureBodyContainer()
	footer := elementFactory("footer")
	markComponentPart(footer, componentPartFooter)

	if v.caption != nil && v.caption.ParentNode() != nil && v.caption.ParentNode().Equals(v.root) {
		v.root.InsertBefore(footer, v.caption)
	} else {
		v.root.AppendChild(footer)
	}

	v.footer = footer
	return v.footer
}

func (v *view) ensureCaptionElement() dom.Element {
	if v.caption != nil {
		return v.caption
	}
	v.ensureBodyContainer()
	caption := elementFactory("caption")
	markComponentPart(caption, componentPartCaption)
	v.root.AppendChild(caption)
	v.caption = caption
	return v.caption
}

func (v *view) replaceChildContent(target dom.Element, children ...any) {
	if target == nil {
		return
	}
	target.SetInnerHTML("")
	for _, child := range children {
		target.AppendChild(NodeFromAny(child))
	}
}

func viewFromElement(element dom.Element) (View, error) {
	if !element.HasAttribute(DataComponentAttrKey) {
		return nil, nil
	}
	name := element.GetAttribute(DataComponentAttrKey)
	if isComponentPart(name) {
		return nil, nil
	}
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

func markComponentPart(element dom.Element, part string) {
	if element == nil {
		return
	}
	if element.HasAttribute(DataComponentAttrKey) {
		value := element.GetAttribute(DataComponentAttrKey)
		if value == part {
			return
		}
		if !isComponentPart(value) {
			panic(fmt.Sprintf("markComponentPart: element already bound to component %q", value))
		}
	}
	element.SetAttribute(DataComponentAttrKey, part)
}

func findComponentPart(root dom.Element, part string) dom.Element {
	if root == nil {
		return nil
	}
	if root.HasAttribute(DataComponentAttrKey) && root.GetAttribute(DataComponentAttrKey) == part {
		return root
	}
	for _, child := range root.ChildNodes() {
		if el, ok := child.(dom.Element); ok {
			if found := findComponentPart(el, part); found != nil {
				return found
			}
		}
	}
	return nil
}

func isComponentPart(value string) bool {
	switch value {
	case componentPartHeader, componentPartBody, componentPartFooter, componentPartCaption:
		return true
	default:
		return false
	}
}
