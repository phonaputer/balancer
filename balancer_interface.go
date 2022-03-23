package balancer

// Balancer is a load balancer which balances over a list of elements of type T contained within it.
type Balancer[T any] interface {

	// Next gets a element from the balancer in order according to its internal balancing rules.
	// This is thread safe.
	Next() T

	// Elements returns the slice of elements being balanced over by this balancer.
	Elements() []T
}
