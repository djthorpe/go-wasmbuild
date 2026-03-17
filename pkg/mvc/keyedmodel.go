package mvc

///////////////////////////////////////////////////////////////////////////////
// TYPES

// Keyed is implemented by types that have a unique, comparable primary key.
// KeyedModel uses this to emit fine-grained "added" and "deleted" events.
type Keyed[K comparable] interface {
	PrimaryKey() K
}

// AddedEvent is emitted when a new item is inserted into a KeyedModel.
type AddedEvent[T any] struct {
	Item  T   // the item that was added
	Index int // its position in the list
	Items []T // full slice after insertion (do not modify)
}

// DeletedEvent is emitted when an item is removed from a KeyedModel.
type DeletedEvent[T any] struct {
	Item  T   // the item that was removed
	Index int // the position it occupied before removal
	Items []T // full slice after removal (do not modify)
}

// ChangedEvent is emitted when an existing item is updated in place.
type ChangedEvent[T any] struct {
	Item  T   // the updated item
	Index int // its position in the list
	Items []T // full slice after update (do not modify)
}

// KeyedModel stores an ordered slice of items of type T (which must implement
// Keyed[K]).
//
// Set and Clear only fire OnSet listeners — added/deleted/changed are not fired.
// Append fires OnAdded for each new key, then OnSet.
// Remove fires OnDeleted, then OnSet.
// Update fires OnChanged (existing key) or OnAdded (new key), then OnSet.
type KeyedModel[K comparable, T Keyed[K]] struct {
	items            []T
	index            map[K]int // key → position in items
	setListeners     []func([]T)
	addedListeners   []func(AddedEvent[T])
	deletedListeners []func(DeletedEvent[T])
	changedListeners []func(ChangedEvent[T])
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Len returns the number of items in the model.
func (m *KeyedModel[K, T]) Len() int {
	return len(m.items)
}

// Items returns a shallow copy of the stored items.
func (m *KeyedModel[K, T]) Items() []T {
	out := make([]T, len(m.items))
	copy(out, m.items)
	return out
}

// Get returns the item with the given key and true, or the zero value and
// false if no such item exists.
func (m *KeyedModel[K, T]) Get(key K) (T, bool) {
	if m.index != nil {
		if i, ok := m.index[key]; ok {
			return m.items[i], true
		}
	}
	var zero T
	return zero, false
}

// Set replaces all items wholesale. Only OnSet listeners fire; added,
// deleted, and changed listeners are not called.
func (m *KeyedModel[K, T]) Set(items []T) {
	m.items = make([]T, len(items))
	copy(m.items, items)
	m.rebuildIndex()
	m.emitSet()
}

// Append adds items that do not already exist (by key). For each new item,
// OnAdded fires before OnSet. Items with an existing key are silently skipped.
func (m *KeyedModel[K, T]) Append(items ...T) {
	changed := false
	for _, item := range items {
		key := item.PrimaryKey()
		if _, exists := m.lookupKey(key); exists {
			continue
		}
		m.initIndex()
		idx := len(m.items)
		m.index[key] = idx
		m.items = append(m.items, item)
		m.emitAdded(AddedEvent[T]{Item: item, Index: idx, Items: m.items})
		changed = true
	}
	if changed {
		m.emitSet()
	}
}

// Remove deletes the item with the given key. OnDeleted fires (with the former
// index) then OnSet fires. No-op if the key does not exist.
func (m *KeyedModel[K, T]) Remove(key K) {
	i, exists := m.lookupKey(key)
	if !exists {
		return
	}
	removed := m.items[i]
	m.items = append(m.items[:i], m.items[i+1:]...)
	m.rebuildIndex()
	m.emitDeleted(DeletedEvent[T]{Item: removed, Index: i, Items: m.items})
	m.emitSet()
}

// Update replaces the item that shares the same primary key, firing OnChanged
// then OnSet. If no item with that key exists yet, it is appended and OnAdded
// fires instead.
func (m *KeyedModel[K, T]) Update(item T) {
	key := item.PrimaryKey()
	if i, exists := m.lookupKey(key); exists {
		m.items[i] = item
		m.emitChanged(ChangedEvent[T]{Item: item, Index: i, Items: m.items})
		m.emitSet()
	} else {
		m.initIndex()
		idx := len(m.items)
		m.index[key] = idx
		m.items = append(m.items, item)
		m.emitAdded(AddedEvent[T]{Item: item, Index: idx, Items: m.items})
		m.emitSet()
	}
}

// Clear removes all items. Only OnSet listeners fire.
func (m *KeyedModel[K, T]) Clear() {
	m.items = m.items[:0]
	m.index = nil
	m.emitSet()
}

// OnSet registers fn to be called after any mutation with the full item
// slice. fn must not modify the slice.
func (m *KeyedModel[K, T]) OnSet(fn func([]T)) {
	m.setListeners = append(m.setListeners, fn)
}

// OnAdded registers fn to be called when a new item is inserted.
func (m *KeyedModel[K, T]) OnAdded(fn func(AddedEvent[T])) {
	m.addedListeners = append(m.addedListeners, fn)
}

// OnDeleted registers fn to be called when an item is removed.
func (m *KeyedModel[K, T]) OnDeleted(fn func(DeletedEvent[T])) {
	m.deletedListeners = append(m.deletedListeners, fn)
}

// OnChanged registers fn to be called when an existing item is updated.
func (m *KeyedModel[K, T]) OnChanged(fn func(ChangedEvent[T])) {
	m.changedListeners = append(m.changedListeners, fn)
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (m *KeyedModel[K, T]) initIndex() {
	if m.index == nil {
		m.index = make(map[K]int, len(m.items))
	}
}

func (m *KeyedModel[K, T]) rebuildIndex() {
	m.index = make(map[K]int, len(m.items))
	for i, item := range m.items {
		m.index[item.PrimaryKey()] = i
	}
}

func (m *KeyedModel[K, T]) lookupKey(key K) (int, bool) {
	if m.index != nil {
		i, ok := m.index[key]
		return i, ok
	}
	return 0, false
}

func (m *KeyedModel[K, T]) emitSet() {
	for _, fn := range m.setListeners {
		fn(m.items)
	}
}

func (m *KeyedModel[K, T]) emitAdded(e AddedEvent[T]) {
	for _, fn := range m.addedListeners {
		fn(e)
	}
}

func (m *KeyedModel[K, T]) emitDeleted(e DeletedEvent[T]) {
	for _, fn := range m.deletedListeners {
		fn(e)
	}
}

func (m *KeyedModel[K, T]) emitChanged(e ChangedEvent[T]) {
	for _, fn := range m.changedListeners {
		fn(e)
	}
}
