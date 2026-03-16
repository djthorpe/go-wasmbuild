package mvc

// Selectable is implemented by views (such as a nav sidebar) that can mark
// one or more child views as "active" and deactivate the rest.
// Call Select with zero arguments to deselect everything.
type Selectable interface {
	Select(views ...View)
}
