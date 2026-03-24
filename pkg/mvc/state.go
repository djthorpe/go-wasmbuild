package mvc

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// PaginationState is implemented by a view that tracks paginated collection
// state in terms of offset, limit, and total count.
//
// Page numbers and navigation state are derived from these values rather than
// stored independently.
type PaginationState interface {
	Offset() uint
	SetOffset(uint) View
	Limit() uint
	SetLimit(uint) View
	Count() uint
	SetCount(uint) View
}

// LabelState is implemented by a view that can set a label on a control
type LabelState interface {
	Label() string
	SetLabel(string) View
}

// ValueState is implemented by a view that exposes a string value.
type ValueState interface {
	Value() string
	SetValue(string) View
}

// ActiveState is implemented by a view that can be marked active or inactive.
type ActiveState interface {
	Active() bool
	SetActive(bool) View
}

// EnabledState is implemented by a view that can be marked enabled or disabled.
type EnabledState interface {
	Enabled() bool
	SetEnabled(bool) View
}

// VisibleState is implemented by a view that can be shown or hidden.
type VisibleState interface {
	Visible() bool
	SetVisible(bool) View
}

// ActiveGroup is implemented by a container that manages which of its member
// views are active. Calling SetActive with no arguments deactivates all members.
type ActiveGroup interface {
	Active() []View
	SetActive(views ...View) View
}

// EnabledGroup is implemented by a container that manages the enabled/disabled
// state of its member views. The passed views are enabled; all others are
// disabled. Calling SetEnabled with no arguments disables all members.
type EnabledGroup interface {
	Enabled() []View
	SetEnabled(views ...View) View
}

// VisibleGroup is implemented by a container that manages the visible/hidden
// state of its member views. The passed views are made visible; all others are
// hidden. Calling SetVisible with no arguments hides all members.
type VisibleGroup interface {
	Visible() []View
	SetVisible(views ...View) View
}

///////////////////////////////////////////////////////////////////////////////
// HELPERS

// Page returns the current 1-based page number derived from offset and limit.
// Returns 0 when the page cannot be determined because limit is zero.
func Page(state PaginationState) uint {
	if state == nil || state.Limit() == 0 {
		return 0
	}
	return (state.Offset() / state.Limit()) + 1
}

// PageCount returns the total number of pages derived from count and limit.
// Returns 0 when the page count cannot be determined because limit is zero.
func PageCount(state PaginationState) uint {
	if state == nil || state.Limit() == 0 {
		return 0
	}
	count := state.Count()
	limit := state.Limit()
	if count == 0 {
		return 1
	}
	return ((count - 1) / limit) + 1
}

// PageStart returns the zero-based inclusive index of the first item on the current page.
func PageStart(state PaginationState) uint {
	if state == nil {
		return 0
	}
	return state.Offset()
}

// PageEnd returns the zero-based exclusive index of the last item on the current page,
// clamped to the known total count.
func PageEnd(state PaginationState) uint {
	if state == nil {
		return 0
	}
	end := state.Offset() + state.Limit()
	count := state.Count()
	if count > 0 && end > count {
		return count
	}
	return end
}

// HasPreviousPage reports whether there is a page before the current one.
func HasPreviousPage(state PaginationState) bool {
	return state != nil && state.Offset() > 0
}

// HasNextPage reports whether there is a page after the current one.
func HasNextPage(state PaginationState) bool {
	if state == nil || state.Limit() == 0 {
		return false
	}
	count := state.Count()
	if count == 0 {
		return false
	}
	return state.Offset()+state.Limit() < count
}
