package balancer

import "testing"

func TestWeightTable_get_internalSliceIsEmpty_returnsZeroValue(t *testing.T) {
	wt := newWeightTable[string](nil)

	res := wt.get(1)

	assertEqual(t, "", res)
}
