package balancer

// WeightedValue contains both a value and a weight for that value.
// This is used as input to NewWeightedRoundRobin.
type WeightedValue[T any] struct {

	// Weight is the number of times this value should be returned by the Balancer's Next function before switching
	// to the next element in the Balancer's list.
	Weight uint32

	// Value is the value of type T which is being included in the Balancer's list of elements.
	Value T
}
