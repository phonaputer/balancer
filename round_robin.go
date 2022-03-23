package balancer

import (
	"sync/atomic"
)

// roundRobin is a Balancer implementation using Round Robin for selecting amongst its contained elements.
type roundRobin[T any] struct {
	elements    []T
	lenElements uint32
	ctr         *uint32 // This relies on uint32 wrapping to 0 on overflow
}

// NewRoundRobin returns a Balancer implemented using the Round Robin method for selecting elements.
//
// Note: behavior is undefined for a list of elements larger than the maximum value for the uint32 type.
func NewRoundRobin[T any](elements []T) Balancer[T] {
	return newRoundRobin(elements)
}

func newRoundRobin[T any](elements []T) *roundRobin[T] {

	// since Next increments before fetching, this will cause the balancer to start from element 0
	ctr := uint32(len(elements) - 1)

	return &roundRobin[T]{
		elements:    elements,
		lenElements: uint32(len(elements)),
		ctr:         &ctr,
	}
}

// Next selects and returns an element from the input list using round-robin
func (b *roundRobin[T]) Next() T {
	idx := atomic.AddUint32(b.ctr, 1)

	return b.elements[idx%b.lenElements]
}

// Elements gets back a slice of the elements which are included in this balancer.
// The types may be mutable, but mutating their order in the slice will not affect
// the roundRobin's internal slice order.
func (b *roundRobin[T]) Elements() []T {
	res := make([]T, b.lenElements)

	copy(res, b.elements)

	return res
}
