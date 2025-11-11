package bootstrap

import (
	"fmt"
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// INIT

func init() {
	// Register the component factory with the dom package
	dom.RegisterComponentFactory(ComponentFromElement)
}

///////////////////////////////////////////////////////////////////////////////
// TYPES

type name string

type component struct {
	name name
	root Element
	body Element // Where content is appended; usually same as root
}

// Ensure that component implements Component interface
var _ Component = (*component)(nil)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// newComponent creates a new component with the given name and root element
// The data-component attribute will be set when applyTo is called
func newComponent(name_ name, root Element) *component {
	c := &component{
		name: name_,
		root: root,
		body: nil,
	}
	return c
}

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

const (
	ContainerComponent      name = "container"
	HeadingComponent        name = "heading"
	BadgeComponent          name = "badge"
	AlertComponent          name = "alert"
	SpanComponent           name = "span"
	ParaComponent           name = "para"
	RuleComponent           name = "rule"
	ButtonComponent         name = "button"
	ButtonGroupComponent    name = "button-group"
	IconComponent           name = "icon"
	LinkComponent           name = "link"
	CardComponent           name = "card"
	ImageComponent          name = "image"
	NavComponent            name = "nav"
	NavBarComponent         name = "nav-bar"
	NavItemComponent        name = "nav-item"
	NavDropdownComponent    name = "nav-dropdown"
	NavSpacerComponent      name = "nav-spacer"
	NavDividerComponent     name = "nav-divider"
	TableComponent          name = "table"
	TableRowComponent       name = "table-row"
	PaginationComponent     name = "pagination"
	PaginationItemComponent name = "pagination-item"
	OffcanvasComponent      name = "offcanvas"
	ToastComponent          name = "toast"
	FormComponent           name = "form"
	InputComponent          name = "input"
	LabelComponent          name = "label"
)

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (component *component) ID() string {
	return component.root.GetAttribute("id")
}

func (component *component) Name() string {
	return string(component.name)
}

func (component *component) Element() Element {
	return component.root
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (component *component) String() string {
	return fmt.Sprint(component.root)
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (component *component) Empty() Component {
	elem := component.body
	if elem == nil {
		elem = component.root
	}
	elem.SetInnerHTML("")
	return component
}

// Insert Component, Element or string children into the body element (or root if no body)
func (component *component) Insert(children ...any) Component {
	if len(children) == 0 {
		return component
	}
	target := component.body
	if target == nil {
		target = component.root
	}

	// Insert first child at the start of the target
	var next Node
	for i, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// First child inserted at beginning
		if i == 0 {
			next = target.InsertBefore(child.(Node), target.FirstChild())
			continue
		}

		// Append after previous child (insert before next sibling)
		next = target.InsertBefore(child.(Node), next.NextSibling())
	}

	// Return the component for chaining
	return component
}

// Append Component, Element or string children to the body element (or root if no body)
func (component *component) Append(children ...any) Component {
	if len(children) == 0 {
		return component
	}
	target := component.body
	if target == nil {
		target = component.root
	}

	for _, child := range children {
		// Convert to Element if necessary
		if component, ok := child.(Component); ok {
			child = component.Element()
		} else if str, ok := child.(string); ok {
			child = dom.GetWindow().Document().CreateTextNode(str)
		}

		// Append to target
		target.AppendChild(child.(Node))
	}

	// Return the component for chaining
	return component
}

// AddEventListener adds an event listener to the component's root element
func (component *component) AddEventListener(event string, handler func(Node)) Component {
	component.root.AddEventListener(event, handler)
	return component
}

// Apply applies options to the component's root element
// It preserves existing id, classes, and attributes by reading them first
func (component *component) Apply(opts ...any) Component {
	// Read existing attributes from the element
	existingOpts := make([]Opt, 0)

	// Preserve existing ID
	if id := component.root.GetAttribute("id"); id != "" {
		existingOpts = append(existingOpts, WithID(id))
	}

	// Preserve existing classes (split space-separated classes into individual tokens)
	if class := component.root.GetAttribute("class"); class != "" {
		// Split the class string by spaces to get individual class names
		classes := strings.Fields(class)
		if len(classes) > 0 {
			existingOpts = append(existingOpts, WithClass(classes...))
		}
	}

	// Preserve existing attributes (excluding id, class, and data-component)
	attrs := component.root.Attributes()
	for _, attr := range attrs {
		name := attr.Name()
		value := attr.Value()
		// Skip id, class, and data-component as they're handled separately
		if name != "id" && name != "class" && name != "data-component" {
			existingOpts = append(existingOpts, WithAttribute(name, value))
		}
	}

	// Convert any opts to Opt type and append them (new opts override existing)
	optList := make([]Opt, 0, len(existingOpts)+len(opts))
	optList = append(optList, existingOpts...)
	for _, opt := range opts {
		if o, ok := opt.(Opt); ok {
			optList = append(optList, o)
		}
	}

	// Apply options to root element
	if err := component.applyTo(component.root, optList...); err != nil {
		panic(err)
	}

	return component
}

///////////////////////////////////////////////////////////////////////////////
// FACTORY FUNCTION

// ComponentFromElement reconstructs a component from an element with data-component attribute
// Returns nil if the element doesn't have data-component or the component type is unknown
//
// This function is registered with the dom package and is called automatically by node.Component()
//
// Example usage:
//
//	buttonGroup.AddEventListener("click", func(node Node) {
//	    // Get the parent ButtonGroup component from any clicked element (including buttons)
//	    if comp := node.Component(); comp != nil {
//	        if bg, ok := comp.(*buttonGroup); ok {
//	            activeIndices := bg.Active()
//	            // ... work with the button group
//	        }
//	    }
//	})
//
// All bootstrap components automatically set data-component attribute via newComponent()
func ComponentFromElement(elem Element) Component {
	if !elem.HasAttribute("data-component") {
		return nil
	}

	componentType := name(elem.GetAttribute("data-component"))

	// Create a basic component wrapper
	c := &component{
		name: componentType,
		root: elem,
		body: nil,
	}

	// For certain component types, we can return specialized wrappers
	// This allows the returned component to have type-specific methods
	switch componentType {
	case ButtonComponent:
		return &button{component: *c}
	case ButtonGroupComponent:
		return &buttonGroup{component: *c}
	case TableComponent:
		return &table{component: *c}
	case TableRowComponent:
		return &tableRow{component: *c}
	case CardComponent:
		return &card{component: *c}
	case ContainerComponent:
		return &container{component: *c}
	case HeadingComponent:
		return &heading{component: *c}
	case BadgeComponent:
		return &badge{component: *c}
	case AlertComponent:
		return &alert{component: *c}
	case IconComponent:
		return &icon{component: *c}
	case NavComponent:
		return &nav{component: *c}
	case NavBarComponent:
		return &navbar{component: *c}
	case NavItemComponent:
		return &navItem{component: *c}
	case NavDropdownComponent:
		return &navDropdown{component: *c}
	case ImageComponent:
		return &image{component: *c}
	case SpanComponent:
		return &span{component: *c}
	case OffcanvasComponent:
		return &offcanvas{component: *c}
	case ToastComponent:
		return &toast{component: *c}
	case FormComponent:
		return &form{component: *c}
	case InputComponent:
		return &input{component: *c}
	case LabelComponent:
		return &label{component: *c}
	// Add more specialized types as needed
	default:
		return c
	}
}
