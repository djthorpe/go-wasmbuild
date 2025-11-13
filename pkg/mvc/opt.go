package mvc

import (
	"slices"
	"strings"

	// Namespace imports
	dom "github.com/djthorpe/go-wasmbuild"
)

/////////////////////////////////////////////////////////////////////////////////
// TYPES

// Opt is a function which can apply options to a view
type Opt func(OptSet) error

// opt is a private struct which holds options
type opt struct {
	doc   dom.Document
	name  string
	id    string
	class []string
	attr  map[string]string
	child dom.Node
}

// OptSet interface for applying options
type OptSet interface {
	// Return the name of the view
	Name() string

	// Return the classes of the view
	Classes() []string
}

// Ensure that opt implements OptSet interface
var _ OptSet = (*opt)(nil)

/////////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func applyOpts(element dom.Element, opts ...Opt) error {
	var o opt

	// Existing name from element
	if element == nil {
		return dom.ErrBadParameter.Withf("Missing Element")
	} else if name := element.GetAttribute(DataComponentAttrKey); name == "" {
		o.name = element.TagName()
	} else {
		o.name = name
	}

	// Set existing Document, ID and class
	o.doc = element.OwnerDocument()
	o.id = element.ID()
	o.class = element.ClassList().Values()

	// Set existing attributes
	o.attr = make(map[string]string)
	for _, attr := range element.Attributes() {
		attrName := attr.Name()

		// Skip id, class, and data-component as they're handled separately
		if attrName != "id" && attrName != "class" && attrName != DataComponentAttrKey {
			o.attr[attrName] = attr.Value()
		}
	}

	// Apply each option to the opt struct
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return err
		}
	}

	// Apply ID if set
	if o.id != "" {
		element.SetID(o.id)
	}

	// Apply classes if set
	if len(o.class) > 0 {
		element.SetAttribute("class", strings.Join(o.Classes(), " "))
	}

	// Apply attributes if set
	for key, value := range o.attr {
		element.SetAttribute(key, value)
	}

	// Apply child node if set
	if o.child != nil {
		element.SetInnerHTML("")
		element.AppendChild(o.child)
	}

	return nil
}

func gatherOpts(args ...any) ([]Opt, []any) {
	var opts []Opt
	var content []any
	for _, arg := range args {
		switch v := arg.(type) {
		case []any:
			o, c := gatherOpts(v...)
			opts = append(opts, o...)
			content = append(content, c...)
		case []Opt:
			opts = append(opts, v...)
		case Opt:
			opts = append(opts, v)
		default:
			content = append(content, v)
		}
	}
	return opts, content
}

/////////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func (o *opt) Name() string {
	return o.name
}

// Return array of non-empty class names
func (o *opt) Classes() []string {
	return slices.Compact(o.class)
}

/////////////////////////////////////////////////////////////////////////////////
// PUBLIC OPTIONS

func WithClass(classes ...string) Opt {
	return func(o OptSet) error {
		o.(*opt).class = append(o.(*opt).class, classes...)
		return nil
	}
}

func WithoutClass(classes ...string) Opt {
	return func(o OptSet) error {
		for _, class := range classes {
			o.(*opt).class = slices.DeleteFunc(o.(*opt).class, func(c string) bool {
				return c == class
			})
		}
		return nil
	}
}

func WithAttr(key, value string) Opt {
	return func(o OptSet) error {
		if o.(*opt).attr == nil {
			o.(*opt).attr = make(map[string]string)
		}
		o.(*opt).attr[key] = value
		return nil
	}
}

func WithoutAttr(keys ...string) Opt {
	return func(o OptSet) error {
		if o.(*opt).attr == nil {
			return nil
		}
		for _, key := range keys {
			delete(o.(*opt).attr, key)
		}
		return nil
	}
}

func WithID(id string) Opt {
	return func(o OptSet) error {
		o.(*opt).id = id
		return nil
	}
}

func WithInnerText(text string) Opt {
	return func(o OptSet) error {
		o.(*opt).child = o.(*opt).doc.CreateTextNode(text)
		return nil
	}
}

func WithStyle(style string) Opt {
	return func(o OptSet) error {
		return WithAttr("style", style)(o)
	}
}
