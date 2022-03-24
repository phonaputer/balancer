package balancer

// Balancer is a load balancer which balances over a list of elements of type T contained within it.
type Balancer[T any] interface {

	// Next gets a element from the balancer in order according to its internal balancing rules.
	// This function is thread safe.
	//
	// If the balancer contains a list of 0 elements, this function will always return the zero value of type T.
	Next() T

	// Elements returns the slice of elements being balanced over by this balancer.
	//
	// If the slice of elements is empty, this function returns nil.
	Elements() []T
}
