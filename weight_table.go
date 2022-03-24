package balancer

type weightTable[T any] struct {
	storedValues []T
	weights      []weightTableEntry
	totalWeight  uint64
}

type weightTableEntry struct {
	start           uint64
	end             uint64
	storedValuesIdx int
}

func newWeightTable[T any](weightedValues []WeightedValue[T]) *weightTable[T] {
	if len(weightedValues) < 1 {
		return &weightTable[T]{totalWeight: 1}
	}

	storedValues := make([]T, len(weightedValues))
	weights := make([]weightTableEntry, len(weightedValues))
	currentWeight := uint64(0)

	for idx, weightedValue := range weightedValues {
		weights[idx] = weightTableEntry{
			start:           currentWeight,
			end:             currentWeight + uint64(weightedValue.Weight),
			storedValuesIdx: idx,
		}

		storedValues[idx] = weightedValue.Value

		currentWeight += uint64(weightedValue.Weight)
	}

	return &weightTable[T]{
		storedValues: storedValues,
		weights:      weights,
		totalWeight:  currentWeight,
	}
}

func (w *weightTable[T]) get(i uint64) T {
	weightIndex := i % w.totalWeight

	for _, entry := range w.weights {
		if entry.start <= weightIndex && weightIndex < entry.end {
			return w.storedValues[entry.storedValuesIdx]
		}
	}

	var zeroValue T

	return zeroValue
}
