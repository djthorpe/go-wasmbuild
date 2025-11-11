package bootstrap

///////////////////////////////////////////////////////////////////////////////
// FORM OPTIONS

// WithReadOnly sets the readonly attribute on an input element
func WithReadOnly() Opt {
	return func(o *opts) error {
		if o.attributes == nil {
			o.attributes = make(map[string]string)
		}
		o.attributes["readonly"] = ""
		return nil
	}
}

// WithPlaceholder sets the placeholder attribute on an input element
func WithPlaceholder(placeholder string) Opt {
	return func(o *opts) error {
		if o.attributes == nil {
			o.attributes = make(map[string]string)
		}
		o.attributes["placeholder"] = placeholder
		return nil
	}
}

// WithValue sets the value attribute on an input element
func WithValue(value string) Opt {
	return func(o *opts) error {
		if o.attributes == nil {
			o.attributes = make(map[string]string)
		}
		o.attributes["value"] = value
		return nil
	}
}
