package mvc

import (
	"fmt"
	"os"
	"slices"
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

	// Return self
	Self() View

	// Return a slot element by name, or nil if not found
	Slot(string) dom.Element

	// Replace a named slot with a node amnd apply options to the slot
	ReplaceSlot(string, any, ...Opt) View

	// Replace the children of a named slot with views, elememts or text and apply options to the slot
	ReplaceSlotChildren(name string, args ...any) View

	// Apply class and attribute options to the root element
	Apply(...Opt) View

	// Replace the content of the view
	Content(...any) View

	// Add an event listener to the view's root element
	AddEventListener(string, func(dom.Event)) View

	// Remove an event listener from the view's root element
	RemoveEventListener(string) View
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
	ContentSlot = "body"
)

var (
	// All the registered views
	views = make(map[string]ViewConstructorFunc, 50)

	// All the registered events
	events = make(map[string][]string, 50)
)

// RegisterView registers a view constructor function for a given name,
// so that the view can be created on-demand, and zero or more event types that
// a controller which attaches to this view should listen for.
func RegisterView(name string, constructor ViewConstructorFunc, eventtypes ...string) {
	if _, exists := views[name]; exists {
		panic("View already registered: " + name)
	}
	views[name] = constructor
	events[name] = eventtypes
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Create a new empty view, applying any options to it
func NewView(self View, name string, template string, fn func(View, View), args ...any) View {
	if _, exists := views[name]; !exists {
		panic(fmt.Sprintf("NewView[%s]: view not registered", name))
	}

	var root dom.Element
	var slot map[string]dom.Element
	if reTagName.MatchString(template) {
		root = HTML(template)
		slot = map[string]dom.Element{
			ContentSlot: root,
		}
	} else {
		root, slot = elementFromTemplate(template)
	}

	// Create the view
	v := &view{self: self, name: name, root: root, slot: slot}

	// Set component identifier
	root.SetAttribute(DataComponentAttrKey, name)

	// Apply options and content to the view
	opts, content := gatherOpts(args...)
	if len(opts) > 0 {
		v.Apply(opts...)
	}

	// Call the initialization function to establish the relationship between
	// the view and its child view
	if fn != nil {
		fn(v.Self(), v)
	}

	// Insert content into the view
	if len(content) > 0 {
		v.Self().Content(content...)
	}

	// Return the view
	return v.Self()
}

// Create view from an existing element, applying any options to it
func NewViewWithElement(self View, element dom.Element, fn func(View, View), opts ...Opt) View {
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

	// Apply options to the view
	if len(opts) > 0 {
		if err := applyOpts(v.root, opts...); err != nil {
			panic(err)
		}
	}

	// TODO: Set the view slots

	// Call the initialization function to establish the relationship between
	// the view and its child view
	if fn != nil {
		fn(v.Self(), v)
	}

	// Return self
	return v.Self()
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

// Return parent view
func (v *view) Parent() View {
	e := v.root
	for {
		// Work up the chain until a view is found
		e = e.ParentElement()
		if e == nil {
			break
		}
		if view, err := viewFromElement(e); err != nil {
			fmt.Fprintf(os.Stderr, "Parent[%s]: %v\n", v.Name(), err)
			break
		} else if view != nil {
			return view
		}
	}
	return nil
}

// Return self
func (v *view) Self() View {
	if v.self == nil {
		return v
	}
	return v.self
}

// Replace the content of the view
func (v *view) Content(args ...any) View {
	return v.ReplaceSlotChildren(ContentSlot, args...)
}

// Return a slot element by name. Returns nil if the slot does not exist
func (v *view) Slot(name string) dom.Element {
	if name = strings.TrimSpace(name); name == "" {
		name = ContentSlot
	}
	return v.slot[name]
}

// Replace a slot with a view, text or element
func (v *view) ReplaceSlot(name string, root any, opts ...Opt) View {
	slot := v.Slot(name)
	if slot == nil {
		panic(fmt.Sprintf("ReplaceSlot[%s]: slot %q does not exist", v.Name(), name))
	}

	// Replace the slot content
	if root != nil {
		node := NodeFromAny(root)
		if node.NodeType() != dom.ELEMENT_NODE {
			panic(fmt.Sprintf("ReplaceSlot[%s]: unsupported node type %v", v.Name(), node.NodeType()))
		}

		// Set the data-slot attribute on the new element
		element, ok := node.(dom.Element)
		if !ok {
			panic(dom.ErrInternalAppError.Withf("ReplaceSlot[%s]: node is not an Element on slot %q", v.Name(), name).Error())
		} else {
			// Actually replace the content in the slot and set it to the new content
			element.SetAttribute(DataSlotAttrKey, name)
			slot.ReplaceWith(element)
			v.slot[name] = element
		}
	}

	// Apply options to the slot element
	if err := applyOpts(slot, opts...); err != nil {
		panic(err)
	}

	// Return self for chaining
	return v.Self()
}

// Replace slot children with  view, text or elements and apply options to the slot
func (v *view) ReplaceSlotChildren(name string, args ...any) View {
	slot := v.Slot(name)
	if slot == nil {
		panic(fmt.Sprintf("ReplaceSlotChildren[%s]: slot %q does not exist", v.Name(), name))
	}

	// Apply options and content to the view
	opts, children := gatherOpts(args...)
	if len(opts) > 0 {
		v.Apply(opts...)
	}

	slot.SetInnerHTML("")
	for _, child := range children {
		slot.AppendChild(NodeFromAny(child))
	}
	return v.Self()
}

// Apply class and attribute options to the view root element
func (v *view) Apply(opts ...Opt) View {
	if err := applyOpts(v.Root(), opts...); err != nil {
		panic(err)
	}
	return v.Self()
}

func (v *view) AddEventListener(event string, handler func(dom.Event)) View {
	v.root.AddEventListener(event, handler)
	return v.Self()
}

func (v *view) RemoveEventListener(event string) View {
	v.root.RemoveEventListener(event)
	return v.Self()
}

///////////////////////////////////////////////////////////////////////////////
// UTILITY METHODS

// ViewFromEvent returns a View from an Event, or nil if the type is unsupported
// or not found. If one or more view names are provided, only views with those names
// are returned.
func ViewFromEvent(e dom.Event, views ...string) View {
	if e == nil {
		return nil
	}
	// Work up the chain until a view is found
	switch element := e.Target().(type) {
	case dom.Element:
		for {
			if view, err := viewFromElement(element); err != nil {
				fmt.Fprintf(os.Stderr, "ViewFromEvent: %v\n", err)
				return nil
			} else if view != nil {
				if len(views) == 0 || slices.Contains(views, view.Name()) {
					return view
				}
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

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

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
		slotmap[ContentSlot] = root
		return root, slotmap
	}

	// Otherwise enumerate the slots, using the 'data-slot' attribute or 'name' attribute
	for _, slot := range slots {
		name := slotNameFromElement(slot)
		if name == "" {
			name = ContentSlot
		}
		if _, exists := slotmap[name]; exists {
			panic("elementFromTemplate: duplicate slot name " + name)
		}
		// Set the slot
		slotmap[name] = slot
	}

	// Ensure a default slot exists
	if _, exists := slotmap[ContentSlot]; !exists {
		slotmap[ContentSlot] = root
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
