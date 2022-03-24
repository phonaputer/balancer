package balancer

import (
	"testing"
)

func TestNewRoundRobin_EmptyList_ReturnsEmptyBalancer(t *testing.T) {
	rr := NewRoundRobin[int](nil)

	assertEqual(t, 0, rr.Next())
	assertEqual(t, 0, len(rr.Elements()))
}

func TestRoundRobin_Next_ThreeElems_IteratesInOrderIncludingWrapping(t *testing.T) {
	type idStruct struct{ ID int }

	testElemOne := &idStruct{ID: 1}
	testElemTwo := &idStruct{ID: 2}
	testElemThree := &idStruct{ID: 3}

	balancer := NewRoundRobin([]*idStruct{testElemOne, testElemTwo, testElemThree})

	for i := 0; i < 100; i++ {
		switch i % 3 {
		case 0:
			assertEqual(t, 1, balancer.Next().ID)
		case 1:
			assertEqual(t, 2, balancer.Next().ID)
		case 2:
			assertEqual(t, 3, balancer.Next().ID)
		}
	}
}

func TestRoundRobin_Next_Uint32Overflow_IteratesInOrderDespiteOverflow(t *testing.T) {
	type idStruct struct{ ID int }

	testElemOne := &idStruct{ID: 1}
	testElemTwo := &idStruct{ID: 2}
	testElemThree := &idStruct{ID: 3}

	balancer := newRoundRobin([]*idStruct{testElemOne, testElemTwo, testElemThree})
	ctr := uint32(4294967292) // max value for uint32 is 4294967295
	balancer.ctr = &ctr

	// get to the overflow
	assertEqual(t, 2, balancer.Next().ID)
	assertEqual(t, 3, balancer.Next().ID)
	assertEqual(t, 1, balancer.Next().ID)

	for i := 0; i < 100; i++ {
		switch i % 3 {
		case 0:
			assertEqual(t, 1, balancer.Next().ID)
		case 1:
			assertEqual(t, 2, balancer.Next().ID)
		case 2:
			assertEqual(t, 3, balancer.Next().ID)
		}
	}
}

func TestRoundRobin_Elements_GetsAllElementsBack(t *testing.T) {
	type idStruct struct{ ID int }

	testElemOne := &idStruct{ID: 1}
	testElemTwo := &idStruct{ID: 2}
	testElemThree := &idStruct{ID: 3}

	balancer := NewRoundRobin([]*idStruct{testElemOne, testElemTwo, testElemThree})
	elems := balancer.Elements()

	assertEqual(t, 3, len(elems))
	assertEqual(t, 1, elems[0].ID)
	assertEqual(t, 2, elems[1].ID)
	assertEqual(t, 3, elems[2].ID)
}

func assertEqual[T comparable](t *testing.T, expected T, actual T) {
	if expected != actual {
		t.Helper()
		t.Fatalf("Not equal! Expected %v but got %v", expected, actual)
	}
}
