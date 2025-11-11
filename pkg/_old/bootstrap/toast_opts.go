package bootstrap

import (
	"strconv"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TOAST OPTIONS

// WithToastPosition sets the position of the toast on the screen.
// Accepts combinations of TOP/BOTTOM and START/END using bitwise OR.
// Example: WithToastPosition(TOP | END) places the toast in the top-right corner.
// Automatically adds position-fixed class.
func WithToastPosition(position Position) Opt {
	return func(o *opts) error {
		// Add position-fixed to make it stick to viewport
		o.classList.Add("position-fixed")

		// Add vertical position (TOP or BOTTOM)
		if position&TOP != 0 {
			o.classList.Add("top-0")
		} else if position&BOTTOM != 0 {
			o.classList.Add("bottom-0")
		}

		// Add horizontal position (START or END)
		if position&START != 0 {
			o.classList.Add("start-0")
		} else if position&END != 0 {
			o.classList.Add("end-0")
		}

		return nil
	}
}

// WithoutAnimation disables the fade animation when showing/hiding the toast.
// By default, toasts use CSS fade transitions.
func WithoutAnimation() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-animation"] = "false"
		return nil
	}
}

// WithTimeout sets the delay (in milliseconds) before automatically hiding the toast.
// This also sets autohide to true. Pass 0 to disable autohide.
// Bootstrap default is 5000ms (5 seconds).
func WithTimeout(duration time.Duration) Opt {
	return func(o *opts) error {
		ms := int(duration.Milliseconds())
		if ms == 0 {
			o.attributes["data-bs-autohide"] = "false"
		} else {
			o.attributes["data-bs-autohide"] = "true"
			o.attributes["data-bs-delay"] = strconv.Itoa(ms)
		}
		return nil
	}
}
