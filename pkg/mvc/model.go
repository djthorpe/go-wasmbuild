package mvc

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Model is a generic observable slice. It holds a cached array of T and
// notifies named listeners whenever the data changes via Set, Append or Delete.
//
// All methods are safe to call from the WASM main goroutine (single-threaded).
//
// Basic usage:
//
//	type Row struct { Name, Role string }
//
//	var m mvc.Model[Row]
//	m.AddEventListener("my-table", func(rows []Row) {
//	    // re-render table from rows
//	})
//	m.Set([]Row{{"Alice", "Engineer"}, {"Bob", "Designer"}})
type Model[T any] struct {
	items     []T
	listeners map[string]func([]T)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Items returns a read-only snapshot of the current slice. The caller must not
// mutate the returned slice.
func (m *Model[T]) Items() []T {
	return m.items
}

// Len returns the number of items in the model.
func (m *Model[T]) Len() int {
	return len(m.items)
}

// Set replaces the entire slice and notifies all listeners.
func (m *Model[T]) Set(items []T) {
	m.items = items
	m.notify()
}

// Append adds one or more items to the end of the slice and notifies all
// listeners.
func (m *Model[T]) Append(items ...T) {
	m.items = append(m.items, items...)
	m.notify()
}

// Delete removes the item at index i and notifies all listeners. It panics if
// i is out of range.
func (m *Model[T]) Delete(i int) {
	m.items = append(m.items[:i], m.items[i+1:]...)
	m.notify()
}

// AddEventListener registers fn under the given name. If a listener with that
// name already exists it is replaced. fn is called immediately with the current
// items so the subscriber can perform an initial render, then again on every
// subsequent mutation.
func (m *Model[T]) AddEventListener(name string, fn func([]T)) {
	if m.listeners == nil {
		m.listeners = make(map[string]func([]T))
	}
	m.listeners[name] = fn
	// Initial call so the subscriber can render the current state.
	fn(m.items)
}

// RemoveEventListener removes the listener registered under name. It is a
// no-op if no listener by that name exists.
func (m *Model[T]) RemoveEventListener(name string) {
	delete(m.listeners, name)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (m *Model[T]) notify() {
	for _, fn := range m.listeners {
		fn(m.items)
	}
}
