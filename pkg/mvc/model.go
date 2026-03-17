package mvc

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Model stores an ordered slice of items of type T and notifies registered
// listeners whenever the contents change.
type Model[T any] struct {
	items     []T
	listeners []func([]T)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Len returns the number of items in the model.
func (m *Model[T]) Len() int {
	return len(m.items)
}

// Items returns a shallow copy of the stored items.
func (m *Model[T]) Items() []T {
	out := make([]T, len(m.items))
	copy(out, m.items)
	return out
}

// Set replaces the stored items and notifies all listeners.
func (m *Model[T]) Set(items []T) {
	m.items = make([]T, len(items))
	copy(m.items, items)
	m.emit()
}

// Append adds items to the end of the slice and notifies all listeners.
func (m *Model[T]) Append(items ...T) {
	m.items = append(m.items, items...)
	m.emit()
}

// Clear removes all items and notifies all listeners.
func (m *Model[T]) Clear() {
	clear(m.items)
	m.items = m.items[:0]
	m.emit()
}

// AddEventListener registers fn to be called whenever the model changes.
// fn receives the current item slice (not a copy — do not modify it).
func (m *Model[T]) AddEventListener(fn func([]T)) {
	m.listeners = append(m.listeners, fn)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (m *Model[T]) emit() {
	for _, fn := range m.listeners {
		fn(m.items)
	}
}
