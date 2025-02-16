package support

import "sync"
import "golang.org/x/exp/constraints"

// GetAverage calculates the average of a slice of numeric values.
func GetAverage[T constraints.Integer](data ...T) T {
	if len(data) == 0 {
		return 0
	}
	total := uint64(0)
	for _, v := range data {
		total += uint64(v)
	}
	return T(total / uint64(len(data)))
}

// Max takes two integers and returns which is greater than the other.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min takes two integers and returns which is smaller than the other.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Subdivide takes a generic slice and subdivides it into sub-slices up to the provided width.
func Subdivide[T any](width int, data []T) [][]T {
	output := make([][]T, 0, (len(data)+width-1)/width)
	if width <= 0 {
		return output
	}

	for i := 0; i < len(data); i += width {
		if i+width > len(data) {
			output = append(output, data[i:])
		} else {
			output = append(output, data[i:i+width])
		}
	}
	return output
}

// RunInParallel takes a slice of slices and, for each outer slice, calls a transformation function
// and returns a slice of the output from each transformation in their respective order.
func RunInParallel[TIn any, TOut any](sliceOfSlices [][]TIn, transformation func([]TIn) TOut) []TOut {
	output := make([]TOut, len(sliceOfSlices))
	var wg sync.WaitGroup

	for i, innerSlice := range sliceOfSlices {
		wg.Add(1)

		go func(index int, data []TIn) {
			defer wg.Done()
			output[index] = transformation(data)
		}(i, innerSlice)
	}

	wg.Wait()
	return output
}
