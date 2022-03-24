package balancer

type emptyBalancer[T any] struct{}

func (e *emptyBalancer[T]) Next() T {
	var zeroValue T

	return zeroValue
}

func (e *emptyBalancer[T]) Elements() []T {
	return nil
}
