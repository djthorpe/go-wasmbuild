package carbon

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

// Event type constants for Carbon views.
const (
	EventClick                  = "click"
	EventInput                  = "input"
	EventChange                 = "change"
	EventInvalid                = "invalid"
	EventHover                  = "mouseenter"
	EventNoHover                = "pointerleave" // mouseleave unreliable on web components; pointerleave respects pointer capture
	EventFocus                  = "focus"
	EventNoFocus                = "focusout"    // blur does not bubble; focusout does, crossing the shadow-DOM boundary
	EventHoverBubbled           = "pointerover" // bubbling hover signal for container-level listeners
	EventNoHoverBubbled         = "pointerout"  // bubbling hover-exit signal for container-level listeners
	EventFocusBubbled           = "focusin"     // bubbling focus signal for container-level listeners
	EventSectionToggle          = "cds-side-nav-menu-toggled"
	EventSectionToggling        = "cds-side-nav-menu-beingtoggled"
	EventSelected               = "cds-dropdown-selected"
	EventOverflowMenuItemClick  = "cds-overflow-menu-item-clicked"
	EventTagDismissibleClosed   = "cds-dismissible-tag-closed"
	EventTagOperationalSelected = "cds-operational-tag-selected"
)

const checkboxChangeEvent = "cds-checkbox-changed"
const numberInputChangeEvent = "cds-number-input"

// EventName maps a raw DOM event type string to its Go constant name.
// Returns the raw string if no mapping is found.
var EventName = map[string]string{
	EventClick:                  "EventClick",
	EventInput:                  "EventInput",
	EventChange:                 "EventChange",
	EventInvalid:                "EventInvalid",
	EventHover:                  "EventHover",
	EventNoHover:                "EventNoHover",
	EventFocus:                  "EventFocus",
	EventNoFocus:                "EventNoFocus",
	EventHoverBubbled:           "EventHover",
	EventNoHoverBubbled:         "EventNoHover",
	EventFocusBubbled:           "EventFocus",
	EventSectionToggle:          "EventSectionToggle",
	EventSectionToggling:        "EventSectionToggling",
	EventSelected:               "EventSelected",
	EventOverflowMenuItemClick:  "EventOverflowMenuItemClick",
	EventTagDismissibleClosed:   "EventTagDismissibleClosed",
	EventTagOperationalSelected: "EventTagOperationalSelected",
	numberInputChangeEvent:      "EventChange",
}

// GoName returns the Go constant name for a raw DOM event type, or the raw
// event type string itself if no mapping exists.
func GoName(eventType string) string {
	if name, ok := EventName[eventType]; ok {
		return name
	}
	return eventType
}
