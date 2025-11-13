package bootstrap

///////////////////////////////////////////////////////////////////////////////
// OFFCANVAS OPTIONS

// WithScroll enables body scrolling when offcanvas is open.
// By default, body scrolling is disabled when offcanvas is shown.
func WithScroll() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-scroll"] = "true"
		return nil
	}
}

// WithBackdrop configures the backdrop behavior for offcanvas.
// Accepts "false" (no backdrop) or "static" (backdrop doesn't close on click).
func WithBackdrop(value string) Opt {
	return func(o *opts) error {
		o.attributes["data-bs-backdrop"] = value
		return nil
	}
}

// WithoutBackdrop configures the backdrop behavior for offcanvas.
func WithoutBackdrop() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-backdrop"] = "false"
		return nil
	}
}

// WithoutKeyboard disables closing the offcanvas with the escape key.
// By default, pressing escape closes the offcanvas.
func WithoutKeyboard() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-keyboard"] = "false"
		return nil
	}
}
