package mvc

///////////////////////////////////////////////////////////////////////////////
// OPTIONS

// WithAriaLabel adds an aria-label attribute to a view
func WithAriaLabel(label string) Opt {
	return func(o OptSet) error {
		return WithAttr("aria-label", label)(o)
	}
}
