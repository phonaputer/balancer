package balancer

import "sync/atomic"

type weightedRoundRobin[T any] struct {
	weightTable *weightTable[T]
	ctr         *uint64 // This relies on uint64 wrapping to 0 on overflow
}

// NewWeightedRoundRobin returns a Balancer implemented using the Weighted Round-robin method for selecting elements.
//
// When calling Next the list of WeightedValues will be iterated in order
// with each element being returned a number of times equal to its weight before moving on to the next element.
//
// Note: behavior is undefined for a total sum of weights larger than the maximum value for the uint64 type.
func NewWeightedRoundRobin[T any](elements []WeightedValue[T]) Balancer[T] {
	if len(elements) < 1 {
		return &emptyBalancer[T]{}
	}

	return newWeightedRoundRobin(elements)
}

func newWeightedRoundRobin[T any](elements []WeightedValue[T]) *weightedRoundRobin[T] {
	table := newWeightTable(elements)

	// since Next increments before fetching, this will cause the balancer to start from element 0
	ctr := table.totalWeight - 1

	return &weightedRoundRobin[T]{
		weightTable: table,
		ctr:         &ctr,
	}
}

// Next selects and returns an element from the input list using weighted round-robin
func (w *weightedRoundRobin[T]) Next() T {
	idx := atomic.AddUint64(w.ctr, 1)

	return w.weightTable.get(idx)
}

// Elements gets back a slice of the elements which are included in this balancer.
// The types may be mutable, but mutating their order in the slice will not affect
// the weightedRoundRobin's internal slice order.
func (w *weightedRoundRobin[T]) Elements() []T {
	res := make([]T, len(w.weightTable.storedValues))

	copy(res, w.weightTable.storedValues)

	return res
}
