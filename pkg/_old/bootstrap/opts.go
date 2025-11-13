package bootstrap

import (
	"strings"

	// Packages
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type opts struct {
	name       name
	id         string
	classList  TokenList
	attributes map[string]string
}

type Opt func(*opts) error

// Breakpoint defines the responsive breakpoint for containers
type Breakpoint string

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	// BreakpointDefault is the default responsive container (max-width at each breakpoint)
	BreakpointDefault Breakpoint = ""
	// BreakpointSmall creates a container that is 100% wide until ≥576px
	BreakpointSmall Breakpoint = "sm"
	// BreakpointMedium creates a container that is 100% wide until ≥768px
	BreakpointMedium Breakpoint = "md"
	// BreakpointLarge creates a container that is 100% wide until ≥992px
	BreakpointLarge Breakpoint = "lg"
	// BreakpointXLarge creates a container that is 100% wide until ≥1200px
	BreakpointXLarge Breakpoint = "xl"
	// BreakpointXXLarge creates a container that is 100% wide until ≥1400px
	BreakpointXXLarge Breakpoint = "xxl"
	// BreakpointFluid creates a full-width container (100% at all breakpoints)
	BreakpointFluid Breakpoint = "fluid"
)

// Size defines button sizes
type Size string

const (
	// SizeDefault is the default button size
	SizeDefault Size = ""
	// SizeSmall creates a small button (btn-sm)
	SizeSmall Size = "sm"
	// SizeLarge creates a large button (btn-lg)
	SizeLarge Size = "lg"
)

// Cursor defines the cursor type
type Cursor string

const (
	// CursorAuto - Browser determines cursor (default)
	CursorAuto Cursor = "auto"
	// CursorDefault - Platform-dependent default cursor (usually an arrow)
	CursorDefault Cursor = "default"
	// CursorPointer - Hand with pointing finger (typically for links)
	CursorPointer Cursor = "pointer"
	// CursorWait - Indicates the program is busy
	CursorWait Cursor = "wait"
	// CursorText - Indicates text can be selected
	CursorText Cursor = "text"
	// CursorMove - Indicates something can be moved
	CursorMove Cursor = "move"
	// CursorNotAllowed - Indicates action is not allowed
	CursorNotAllowed Cursor = "not-allowed"
	// CursorHelp - Help information is available
	CursorHelp Cursor = "help"
	// CursorCrosshair - Simple crosshair
	CursorCrosshair Cursor = "crosshair"
	// CursorGrab - Indicates something can be grabbed
	CursorGrab Cursor = "grab"
	// CursorGrabbing - Indicates something is being grabbed
	CursorGrabbing Cursor = "grabbing"
	// CursorZoomIn - Indicates zoom in
	CursorZoomIn Cursor = "zoom-in"
	// CursorZoomOut - Indicates zoom out
	CursorZoomOut Cursor = "zoom-out"
	// CursorNone - No cursor is rendered
	CursorNone Cursor = "none"
	// CursorProgress - Indicates program is busy but user can still interact
	CursorProgress Cursor = "progress"
	// CursorCopy - Indicates something will be copied
	CursorCopy Cursor = "copy"
	// CursorAlias - Indicates an alias or shortcut will be created
	CursorAlias Cursor = "alias"
	// CursorContextMenu - Context menu is available
	CursorContextMenu Cursor = "context-menu"
	// CursorCell - Indicates table cell selection
	CursorCell Cursor = "cell"
	// CursorVerticalText - Indicates vertical text can be selected
	CursorVerticalText Cursor = "vertical-text"
	// CursorNResize - North resize
	CursorNResize Cursor = "n-resize"
	// CursorEResize - East resize
	CursorEResize Cursor = "e-resize"
	// CursorSResize - South resize
	CursorSResize Cursor = "s-resize"
	// CursorWResize - West resize
	CursorWResize Cursor = "w-resize"
	// CursorNEResize - North-East resize
	CursorNEResize Cursor = "ne-resize"
	// CursorNWResize - North-West resize
	CursorNWResize Cursor = "nw-resize"
	// CursorSEResize - South-East resize
	CursorSEResize Cursor = "se-resize"
	// CursorSWResize - South-West resize
	CursorSWResize Cursor = "sw-resize"
	// CursorEWResize - East-West resize
	CursorEWResize Cursor = "ew-resize"
	// CursorNSResize - North-South resize
	CursorNSResize Cursor = "ns-resize"
	// CursorNESWResize - North-East/South-West resize
	CursorNESWResize Cursor = "nesw-resize"
	// CursorNWSEResize - North-West/South-East resize
	CursorNWSEResize Cursor = "nwse-resize"
	// CursorColResize - Column resize
	CursorColResize Cursor = "col-resize"
	// CursorRowResize - Row resize
	CursorRowResize Cursor = "row-resize"
	// CursorAllScroll - Indicates scrolling in any direction
	CursorAllScroll Cursor = "all-scroll"
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewOpts(name name, opt ...Opt) (*opts, error) {
	o := &opts{
		name:       name,
		classList:  dom.NewTokenList(),
		attributes: make(map[string]string),
	}
	if err := o.apply(opt...); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *opts) apply(opts ...Opt) error {
	for _, opt := range opts {
		if opt != nil {
			if err := opt(o); err != nil {
				return err
			}
		}
	}
	return nil
}

// Apply options to any dom.Element
func (component *component) applyTo(root Element, opts ...Opt) error {
	opt, err := NewOpts(component.name, opts...)
	if err != nil {
		return err
	}

	// Set class list first (Bootstrap convention: class comes before other attributes)
	classes := opt.classList.Values()
	if len(classes) > 0 {
		root.SetAttribute("class", strings.Join(classes, " "))
	}

	// Set data-component attribute second (for consistency in output)
	root.SetAttribute("data-component", string(component.name))

	// Set ID if provided
	if opt.id != "" {
		root.SetAttribute("id", opt.id)
	}

	// Set other attributes
	for key, value := range opt.attributes {
		root.SetAttribute(key, value)
	}

	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC FUNCTIONS

func WithClass(class ...string) Opt {
	return func(o *opts) error {
		o.classList.Add(class...)
		return nil
	}
}

func WithoutClass(class ...string) Opt {
	return func(o *opts) error {
		o.classList.Remove(class...)
		return nil
	}
}

// WithBreakpoint sets the container breakpoint type.
// Use BreakpointFluid for full-width, BreakpointSmall/Medium/Large/XLarge/XXLarge for responsive containers,
// or BreakpointDefault for the default container behavior.
func WithBreakpoint(breakpoint Breakpoint) Opt {
	return func(o *opts) error {
		// Error if this isn't a container
		if o.name != ContainerComponent {
			return ErrBadParameter.Withf("Cannot use WithBreakpoint with component of type %q", o.name)
		}

		// Remove any existing container classes
		o.classList.Remove("container", "container-sm", "container-md", "container-lg", "container-xl", "container-xxl", "container-fluid")

		// Add the appropriate container class
		if breakpoint == BreakpointDefault || breakpoint == "" {
			o.classList.Add("container")
		} else {
			o.classList.Add("container-" + string(breakpoint))
		}
		return nil
	}
}

func WithBorder(position Position, color ...Color) Opt {
	return func(o *opts) error {
		// WithBorder works with heading, container, badge, alert, and button-group components
		if o.name != HeadingComponent && o.name != ContainerComponent && o.name != BadgeComponent && o.name != AlertComponent && o.name != ButtonGroupComponent {
			return ErrBadParameter.Withf("Cannot use WithBorder with component of type %q", o.name)
		}

		// Remove all existing border classes
		o.classList.Remove(BorderAll.borderClassNames()...)

		// Add new border classes
		if pposition := position & BorderAll; pposition != 0 {
			o.classList.Add(pposition.borderClassNames()...)
		}

		// TODO: Remove any existing border color classes

		// Add border color
		if len(color) == 1 {
			o.classList.Add(color[0].className("border"))
		} else if len(color) > 1 {
			return ErrBadParameter.Withf("Only one color can be specified for borders")
		}

		// Return success
		return nil
	}
}

func WithBackground(color Color) Opt {
	return func(o *opts) error {
		// TODO: Remove any existing background color classes

		// Add background color
		o.classList.Add(color.className("bg"))

		// Return success
		return nil
	}
}

// WithColor sets the color for badge, alert, button, icon, link, card, table, and navbar components.
// For badges: <span class="badge text-bg-primary">Primary</span>
// For alerts: <div class="alert alert-primary" role="alert">Primary</div>
// For buttons: <button class="btn btn-primary">Primary</button>
// For icons: <i class="bi bi-icon text-primary"></i>
// For links: <a class="link-primary">Link</a>
// For cards: <div class="card text-bg-primary">Card</div>
// For tables: <table class="table table-primary">Table</table>
// For navbars: <nav class="navbar bg-primary">Navbar</nav>
func WithColor(color Color) Opt {
	return func(o *opts) error {
		// Define all possible color suffixes
		colorSuffixes := []string{
			"primary", "primary-subtle",
			"secondary", "secondary-subtle",
			"success", "success-subtle",
			"danger", "danger-subtle",
			"warning", "warning-subtle",
			"info", "info-subtle",
			"light", "light-subtle",
			"dark", "dark-subtle",
			"white", "black",
		}

		// Remove existing color classes based on component type
		switch o.name {
		case BadgeComponent, CardComponent:
			// Remove text-bg-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("text-bg-" + suffix)
			}
		case AlertComponent:
			// Remove alert-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("alert-" + suffix)
			}
		case ButtonComponent:
			// Remove btn-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("btn-" + suffix)
			}
		case IconComponent:
			// Remove text-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("text-" + suffix)
			}
		case LinkComponent:
			// Remove link-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("link-" + suffix)
			}
		case TableComponent:
			// Remove table-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("table-" + suffix)
			}
		case NavBarComponent:
			// Remove bg-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("bg-" + suffix)
			}
		case ToastComponent:
			// Remove text-bg-* classes
			for _, suffix := range colorSuffixes {
				o.classList.Remove("text-bg-" + suffix)
			}
		default:
			return ErrBadParameter.Withf("Cannot use WithColor with component of type %q", o.name)
		}

		// Add appropriate color class based on component type (skip if TRANSPARENT)
		if color != "" {
			switch o.name {
			case BadgeComponent:
				o.classList.Add(color.className("text-bg"))
			case AlertComponent:
				o.classList.Add(color.className("alert"))
			case ButtonComponent:
				o.classList.Add(color.className("btn"))
			case IconComponent:
				o.classList.Add(color.className("text"))
			case LinkComponent:
				o.classList.Add(color.className("link"))
			case CardComponent:
				o.classList.Add(color.className("text-bg"))
			case TableComponent:
				o.classList.Add(color.className("table"))
			case NavBarComponent:
				o.classList.Add(color.className("bg"))
			case ToastComponent:
				o.classList.Add(color.className("text-bg"))
			}
		}

		// Return success
		return nil
	}
}

// WithResponsive sets the responsive breakpoint for tables or navbars.
// For tables: makes the table horizontally scrollable on smaller screens.
// For navbars: sets when the navbar should expand from collapsed to full horizontal layout.
// Use specific breakpoints (sm, md, lg, xl, xxl) to make responsive up to that breakpoint.
// For tables: <table class="table table-responsive">Table</table>
// For specific breakpoint: <table class="table table-responsive-md">Table</table>
// For navbars: <nav class="navbar navbar-expand-lg">Navbar</nav>
func WithResponsive(breakpoint Breakpoint) Opt {
	return func(o *opts) error {
		// WithResponsive works with table and navbar components
		if o.name != TableComponent && o.name != NavBarComponent {
			return ErrBadParameter.Withf("Cannot use WithResponsive with component of type %q", o.name)
		}

		// Add responsive class based on component type
		if o.name == TableComponent {
			if breakpoint == BreakpointDefault || breakpoint == "" {
				o.classList.Add("table-responsive")
			} else {
				o.classList.Add("table-responsive-" + string(breakpoint))
			}
		} else if o.name == NavBarComponent {
			// For navbars, add navbar-expand-{breakpoint}
			if breakpoint == BreakpointDefault || breakpoint == "" {
				// No expand class means always collapsed
				// (usually you want a breakpoint, but we allow this)
			} else {
				o.classList.Add("navbar-expand-" + string(breakpoint))
			}
		}

		// Return success
		return nil
	}
}

// WithMargin sets the margin for the specified position, where position can be one or more of
// TOP, BOTTOM, START, END. Margin is an integer from -5 to 5, where negative values indicate
// negative margins, 0 indicates no margin, and positive values indicate increasing margin sizes.
func WithMargin(position Position, margin int) Opt {
	return func(o *opts) error {
		// TODO: Remove any existing margin classes

		// Add margin classes
		o.classList.Add(position.marginClassNames(margin)...)

		// Return success
		return nil
	}
}

// WithPadding sets the padding for the specified position, where position can be one or more of
// TOP, BOTTOM, START, END. Padding is an integer from 1 to 5, where negative or zero values are not allowed.
func WithPadding(position Position, padding int) Opt {
	return func(o *opts) error {
		// Validate padding size (must be positive, no zero or negative)
		if padding <= 0 {
			return ErrBadParameter.Withf("Padding must be a positive value, got %d", padding)
		}

		// TODO: Remove any existing padding classes

		// Add padding classes
		o.classList.Add(position.paddingClassNames(padding)...)

		// Return success
		return nil
	}
}

// WithSize sets the size for button, button group, or table components.
// For buttons: btn-sm, btn-lg
// For button groups: btn-group-sm, btn-group-lg
// For tables: table-sm (only SizeSmall and SizeDefault supported)
func WithSize(size Size) Opt {
	return func(o *opts) error {
		// WithSize only works with button, button-group, or table components
		if o.name != ButtonComponent && o.name != ButtonGroupComponent && o.name != TableComponent {
			return ErrBadParameter.Withf("Cannot use WithSize with component of type %q", o.name)
		}

		// For tables, only allow SizeDefault and SizeSmall
		if o.name == TableComponent {
			if size != SizeDefault && size != SizeSmall {
				return ErrBadParameter.Withf("Table component only supports SizeDefault and SizeSmall, got %q", size)
			}

			// Remove existing size classes
			o.classList.Remove("table-sm")

			// Add size class if not default
			if size == SizeSmall {
				o.classList.Add("table-sm")
			}

			return nil
		}

		// Skip if default size
		if size == SizeDefault {
			return nil
		}

		// TODO: Remove any existing size classes for buttons

		// Add size class based on component type
		if o.name == ButtonGroupComponent {
			o.classList.Add("btn-group-" + string(size))
		} else {
			o.classList.Add("btn-" + string(size))
		}

		// Return success
		return nil
	}
}

// WithAriaLabel sets the aria-label attribute for accessibility
func WithAriaLabel(label string) Opt {
	return func(o *opts) error {
		if label != "" {
			o.attributes["aria-label"] = label
		}
		return nil
	}
}

// WithRole sets the role attribute for accessibility
func WithRole(role string) Opt {
	return func(o *opts) error {
		if role != "" {
			o.attributes["role"] = role
		}
		return nil
	}
}

// WithAttribute sets a custom HTML attribute
func WithAttribute(key, value string) Opt {
	return func(o *opts) error {
		if key != "" {
			o.attributes[key] = value
		}
		return nil
	}
}

// WithFlex sets flexbox alignment classes based on position
// - CENTER or MIDDLE: adds "d-flex align-items-center"
// - START or TOP: adds "d-flex align-items-start"
// - END or BOTTOM: adds "d-flex align-items-end"
// - START|END: adds "d-flex flex-row" (horizontal flow)
// - TOP|BOTTOM: adds "d-flex flex-column" (vertical flow)
func WithFlex(position Position) Opt {
	return func(o *opts) error {
		// Always add d-flex
		o.classList.Add("d-flex")

		// Handle alignment
		switch {
		case position&CENTER != 0 || position&MIDDLE != 0:
			o.classList.Add("align-items-center")
		case position&START != 0 && position&END == 0:
			o.classList.Add("align-items-start")
		case position&TOP != 0 && position&BOTTOM == 0:
			o.classList.Add("align-items-start")
		case position&END != 0 && position&START == 0:
			o.classList.Add("align-items-end")
		case position&BOTTOM != 0 && position&TOP == 0:
			o.classList.Add("align-items-end")
		}

		// Handle flow direction
		if (position&START != 0) && (position&END != 0) {
			// Both START and END means horizontal flow
			o.classList.Add("flex-row")
		} else if (position&TOP != 0) && (position&BOTTOM != 0) {
			// Both TOP and BOTTOM means vertical flow
			o.classList.Add("flex-column")
		}

		return nil
	}
}

func WithTextAlign(position Position) Opt {
	return func(o *opts) error {
		switch {
		case position&CENTER != 0:
			o.classList.Add("text-center")
		case position&START != 0:
			o.classList.Add("text-start")
		case position&END != 0:
			o.classList.Add("text-end")
		}
		return nil
	}
}

// WithCursor sets the CSS cursor style for the component
// Accepts any Cursor constant (e.g., CursorPointer, CursorNotAllowed, CursorGrab, etc.)
func WithCursor(cursor Cursor) Opt {
	return func(o *opts) error {
		if cursor != "" && cursor != CursorAuto {
			o.attributes["style"] = "cursor: " + string(cursor) + ";"
		}
		return nil
	}
}

// WithID sets the ID attribute for any component.
// This is a general-purpose option that can be used by all components.
func WithID(id string) Opt {
	return func(o *opts) error {
		o.id = id
		return nil
	}
}

// WithPosition sets the position for components that support positioning.
// For offcanvas: TOP, BOTTOM, START, END
// The specific class name will be determined by the component.
func WithPosition(position Position) Opt {
	return func(o *opts) error {
		// Store position in a special attribute that components can read
		// Components will translate this to their specific class names
		if className := position.className(""); className != "" {
			o.attributes["data-position"] = className
		}
		return nil
	}
}

// WithTheme sets the Bootstrap theme.
// Only LIGHT or DARK colors are supported for themes.
// Applies data-bs-theme attribute and adds text-bg-dark class for DARK theme.
func WithTheme(theme Color) Opt {
	return func(o *opts) error {
		if theme != LIGHT && theme != DARK {
			return ErrBadParameter.Withf("WithTheme only accepts LIGHT or DARK colors, got %q", theme)
		}
		o.attributes["data-bs-theme"] = string(theme)
		if theme == DARK {
			o.classList.Add("text-bg-dark")
		}
		return nil
	}
}
