package mvc

import (
	"fmt"
	"os"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// View represents a web component in the interface
type View interface {
	// Return the view name
	Name() string

	// Return the view ID, if set
	ID() string

	// Return the view's root element
	Root() dom.Element

	// Return the view's parent view
	Parent() View

	// Return a slot by name, or nil if not found
	Slot(string) dom.Element

	// Replace a named slot with a node
	ReplaceSlot(string, any) View

	// Add an event listener to the view's root element
	AddEventListener(string, func(dom.Event)) View

	// Return the value of the view as a string. The contents of the
	// string depends on the view type
	Value() string

	// Set the value of the view as a string. The interpretation of the
	// string depends on the view type
	Set(string) View

	// Apply class and attribute options to the view
	Apply(...Opt) View

	// TODO - Deprecate these methods in favour of ReplaceSlot

	// Set the view's content to the given text, Element or View children
	// If no arguments are given, the content is cleared
	Content(...any) View

	// Set the view's label element. Panics if the view does not have a slot
	// called "label"
	Label(...any) View

	// Set the view's header element. Panics if the view does not have a slot
	// called "header"
	Header(...any) View

	// Set the view's footer element. Panics if the view does not have a slot
	// called "footer"
	Footer(...any) View
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

///////////////////////////////////////////////////////////////////////////////
// PRIVATE TYPES

// Implementation of View interface
type view struct {
	self View
	name string
	root dom.Element
	slot map[string]dom.Element
}

// Ensure that view implements View interface
var _ View = (*view)(nil)

// Constructor function for views
type ViewConstructorFunc func(dom.Element) View

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	// The attribute key which identifies an mvc component
	DataComponentAttrKey = "data-mvc"

	// The attribute key which identifies a slot in a component
	DataSlotAttrKey = "data-slot"

	// The name of the default slot, when name atribute is missing
	defaultSlot = "body"
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

// Create a new view with template, applying any options to it
func NewViewExEx(self View, name string, template string, args ...any) View {
	if _, exists := views[name]; !exists {
		panic(fmt.Sprintf("NewViewExEx: view not registered %q", name))
	}
	// Create the element from the template, and return the slots on the view
	root, slots := elementFromTemplate(template)

	// Create the view
	v := &view{
		self: self,
		name: name,
		root: root,
		slot: slots,
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", name))
	} else {
		self_.SetView(v)
	}

	// Set the component identifier
	v.root.SetAttribute(DataComponentAttrKey, name)

	// Apply options to the view
	opts, content := gatherOpts(args...)
	v.Apply(opts...)

	// Set the content in the view
	if len(content) > 0 {
		v.self.Content(content...)
	}

	// Return the view
	return v.self
}

func elementFromTemplate(template string) (dom.Element, map[string]dom.Element) {
	// Create the root element
	root := elementFactory("div")
	root.SetInnerHTML(template)

	// There should be a single child element
	if root.ChildElementCount() != 1 {
		panic(fmt.Sprintf("elementFromTemplate: template must have a single root element, found %d", root.ChildElementCount()))
	} else {
		root = root.FirstElementChild()
	}

	// Find the slots in the template by <slot> elements, or data-slot attributes
	slots := root.GetElementsByTagName("slot")
	if len(slots) == 0 {
		slots = getElementsByAttribute(root, DataSlotAttrKey)
	}

	// Create the slot map
	slotmap := make(map[string]dom.Element, len(slots))

	// In the case there is no slot, use the root element as the default slot
	if len(slots) == 0 {
		slotmap[defaultSlot] = root
		return root, slotmap
	}

	// Otherwise enumerate the slots, using the 'data-slot' attribute or 'name' attribute
	for _, slot := range slots {
		name := slotNameFromElement(slot)
		if name == "" {
			name = defaultSlot
		}
		if _, exists := slotmap[name]; exists {
			panic("elementFromTemplate: duplicate slot name " + name)
		}
		// Set the slot
		slotmap[name] = slot
	}

	// Ensure a default slot exists
	if _, exists := slotmap[defaultSlot]; !exists {
		slotmap[defaultSlot] = root
	}

	// Return the root element and slot map
	return root, slotmap
}

func slotNameFromElement(element dom.Element) string {
	// Slot name is from 'name' attribute for <slot> elements, or 'data-slot' attribute otherwise
	if element.TagName() == "SLOT" {
		return element.GetAttribute("name")
	} else {
		return element.GetAttribute(DataSlotAttrKey)
	}
}

func getElementsByAttribute(root dom.Element, attr string) []dom.Element {
	var elements []dom.Element
	children := root.Children()
	for _, child := range children {
		// Recursively search child elements
		if child.HasAttribute(attr) {
			elements = append(elements, child)
		} else {
			elements = append(elements, getElementsByAttribute(child, attr)...)
		}
	}
	return elements
}

// Create a new empty view, applying any options to it
func NewView(self View, name string, tagName string, args ...any) View {
	if _, exists := views[name]; !exists {
		panic(fmt.Sprintf("NewView: view not registered %q", name))
	}

	// Create the view
	v := &view{
		self: self,
		name: name,
		root: elementFactory(tagName),
		slot: make(map[string]dom.Element),
	}

	// Set the view in self
	if self_, ok := self.(ViewWithSelf); !ok {
		panic(fmt.Sprintf("NewView: %v does not implement ViewWithSelf", name))
	} else {
		self_.SetView(v)
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
		panic("NewViewWithElement: element missing data-mvc attribute")
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

	// TODO: Set the view slots

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

func (v *view) Name() string {
	return v.name
}

func (v *view) ID() string {
	return v.root.ID()
}

func (v *view) Root() dom.Element {
	return v.root
}

func (v *view) Parent() View {
	e := v.root
	for {
		// Work up the chain until a view is found
		e = e.ParentElement()
		if e == nil {
			break
		}
		if view, err := viewFromElement(e); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			break
		} else if view != nil {
			return view
		}
	}
	return nil
}

// Return a slot element by name. Returns nil if the slot does not exist
func (v *view) Slot(name string) dom.Element {
	if name == "" {
		name = defaultSlot
	}
	return v.slot[name]
}

// Replace a slot with a view, text or element
func (v *view) ReplaceSlot(name string, root any) View {
	// Set name for default slot
	if name = strings.TrimSpace(name); name == "" {
		name = defaultSlot
	}

	// Ensure slot exists
	slot, exists := v.slot[name]
	if !exists {
		panic(fmt.Sprintf("ReplaceSlot: slot %q does not exist", name))
	}

	// Replace the slot content
	node := NodeFromAny(root)
	if node.NodeType() != dom.ELEMENT_NODE {
		panic(fmt.Sprintf("ReplaceSlot: unsupported node type %v", node.NodeType()))
	}

	// Set the data-slot attribute on the new element
	element, ok := node.(dom.Element)
	if !ok {
		panic(dom.ErrInternalAppError.Withf("ReplaceSlot: node is not an Element on slot %q of view %q", name, v.Name()).Error())
	} else {
		// Actually replace the content in the slot and set it to the new content
		element.SetAttribute(DataSlotAttrKey, name)
		slot.ReplaceWith(element)
		v.slot[name] = element
	}

	// Return self for chaining
	return v.self
}

// Apply class and attribute options to the view root element
func (v *view) Apply(opts ...Opt) View {
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}
	return v.self
}

// Set content of the default slot
func (v *view) Content(children ...any) View {
	target, exists := v.slot[defaultSlot]
	if !exists {
		target = v.root
	}
	return v.replaceChildContent(target, children...)
}

func (v *view) Header(children ...any) View {
	slot, exists := v.slot["header"]
	if !exists {
		panic("view.Header: view does not have a header slot")
	}
	return v.replaceChildContent(slot, children...)
}

func (v *view) Footer(children ...any) View {
	slot, exists := v.slot["footer"]
	if !exists {
		panic("view.Footer: view does not have a footer slot")
	}
	return v.replaceChildContent(slot, children...)
}

func (v *view) Label(children ...any) View {
	slot, exists := v.slot["label"]
	if !exists {
		panic("view.Label: view does not have a label slot")
	}
	return v.replaceChildContent(slot, children...)
}

func (v *view) AddEventListener(event string, handler func(dom.Event)) View {
	v.root.AddEventListener(event, handler)
	return v.self
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS - ViewWithValue

func (v *view) Value() string {
	return v.root.Value()
}

func (v *view) Set(value string) View {
	v.root.SetValue(value)
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

func (v *view) replaceChildContent(target dom.Element, children ...any) View {
	if target == nil {
		return v.self
	}
	target.SetInnerHTML("")
	for _, child := range children {
		target.AppendChild(NodeFromAny(child))
	}
	return v.self
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
