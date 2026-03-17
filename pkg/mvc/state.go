package mvc

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// ActiveState is implemented by a view that can be marked active or inactive.
type ActiveState interface {
	SetActive(bool)
}

// ActiveGroup is implemented by a container that manages which of its member
// views are active. Calling SetActive with no arguments deactivates all members.
type ActiveGroup interface {
	SetActive(views ...View)
}
