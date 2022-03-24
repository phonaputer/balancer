package balancer

import "testing"

func TestNewWeightedRoundRobin_EmptyList_ReturnsEmptyBalancer(t *testing.T) {
	rr := NewWeightedRoundRobin[int](nil)

	assertEqual(t, 0, rr.Next())
	assertEqual(t, 0, len(rr.Elements()))
}

func TestWeightedRoundRobin_Elements_ListOfThree_ReturnsListInTheOrderItWasInput(t *testing.T) {
	type idStruct struct {
		ID int
	}

	input := []WeightedValue[idStruct]{
		{Weight: 5, Value: idStruct{ID: 1}},
		{Weight: 5, Value: idStruct{ID: 2}},
		{Weight: 5, Value: idStruct{ID: 3}},
	}

	wrr := NewWeightedRoundRobin(input)
	elems := wrr.Elements()

	assertEqual(t, 3, len(elems))
	assertEqual(t, 1, elems[0].ID)
	assertEqual(t, 2, elems[1].ID)
	assertEqual(t, 3, elems[2].ID)
}

func TestWeightedRoundRobin_Next_ThreeElements_IteratesInOrderRespectingWeights(t *testing.T) {
	type idStruct struct {
		ID int
	}

	input := []WeightedValue[idStruct]{
		{Weight: 2, Value: idStruct{ID: 1}},
		{Weight: 1, Value: idStruct{ID: 2}},
		{Weight: 7, Value: idStruct{ID: 3}},
	}

	wrr := NewWeightedRoundRobin(input)

	for i := 0; i < 100; i++ {
		weightedIdx := i % 10
		if weightedIdx < 2 {
			assertEqual(t, 1, wrr.Next().ID)
		} else if weightedIdx < 3 {
			assertEqual(t, 2, wrr.Next().ID)
		} else {
			assertEqual(t, 3, wrr.Next().ID)
		}
	}
}

func TestWeightedRoundRobin_Next_Uint64Overflow_CanWrapBackToBeginningAfterOverflow(t *testing.T) {
	type idStruct struct{ ID int }

	input := []WeightedValue[idStruct]{
		{Weight: 2, Value: idStruct{ID: 1}},
		{Weight: 1, Value: idStruct{ID: 2}},
		{Weight: 7, Value: idStruct{ID: 3}},
	}

	wrr := newWeightedRoundRobin(input)
	ctr := uint64(18446744073709551612) // max value for uint32 is 18446744073709551615
	wrr.ctr = &ctr

	// get to the overflow
	assertEqual(t, 3, wrr.Next().ID)
	assertEqual(t, 3, wrr.Next().ID)
	assertEqual(t, 3, wrr.Next().ID)

	for i := 0; i < 100; i++ {
		weightedIdx := i % 10
		if weightedIdx < 2 {
			assertEqual(t, 1, wrr.Next().ID)
		} else if weightedIdx < 3 {
			assertEqual(t, 2, wrr.Next().ID)
		} else {
			assertEqual(t, 3, wrr.Next().ID)
		}
	}
}
