package atomic

import (
	"sync"
	"sync/atomic"
)

// Slice is a thread-safe slice implementation.
type Slice[T any] struct {
	mu        sync.RWMutex
	atomicVal atomic.Value
	slice     []T
}

// NewSlice creates a new instance of a Slice.
func NewSlice[T any]() *Slice[T] {
	as := &Slice[T]{}
	as.atomicVal.Store(make([]T, 0))
	return as
}

// Add appends elements to the slice in a thread-safe manner.
func (as *Slice[T]) Add(elements ...T) {
	as.mu.Lock()
	defer as.mu.Unlock()
	current := as.atomicVal.Load().([]T)
	newSlice := append(current, elements...)
	as.atomicVal.Store(newSlice)
}

// RemoveIf removes elements that match the predicate in a thread-safe manner.
func (as *Slice[T]) RemoveIf(predicate func(T) bool) {
	as.mu.Lock()
	defer as.mu.Unlock()
	var result []T
	for _, v := range as.slice {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	as.slice = result
}

// Length returns the number of elements in the slice.
func (as *Slice[T]) Length() int {
	as.mu.RLock()
	defer as.mu.RUnlock()
	return len(as.slice)
}

// Get returns the element at the provided index in a thread-safe manner.
func (as *Slice[T]) Get(index int) T {
	as.mu.RLock()
	defer as.mu.RUnlock()
	return as.slice[index]
}

// All returns a copy of all elements in the slice in a thread-safe manner.
func (as *Slice[T]) All() []T {
	return as.atomicVal.Load().([]T)
}

// IfAny walks the slice and returns true as soon as an element satisfies the predicate.
func (as *Slice[T]) IfAny(predicate func(T) bool) bool {
	for _, v := range as.All() {
		if predicate(v) {
			return true
		}
	}
	return false
}
