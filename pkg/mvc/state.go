package mvc

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// ActiveState is implemented by a view that can be marked active or inactive.
type ActiveState interface {
	Active() bool
}

// EnabledState is implemented by a view that can be marked enabled or disabled.
type EnabledState interface {
	Enabled() bool
}

// VisibleState is implemented by a view that can be shown or hidden.
type VisibleState interface {
	Visible() bool
	SetVisible(bool) View
}

// ActiveGroup is implemented by a container that manages which of its member
// views are active. Calling SetActive with no arguments deactivates all members.
type ActiveGroup interface {
	SetActive(views ...View) View
}

// EnabledGroup is implemented by a container that manages the enabled/disabled
// state of its member views. The passed views are enabled; all others are
// disabled. Calling SetEnabled with no arguments disables all members.
type EnabledGroup interface {
	SetEnabled(views ...View)
}

// VisibleGroup is implemented by a container that manages the visible/hidden
// state of its member views. The passed views are made visible; all others are
// hidden. Calling SetVisible with no arguments hides all members.
type VisibleGroup interface {
	SetVisible(views ...View)
}
