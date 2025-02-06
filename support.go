package support

import "sync"

var Ipsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque imperdiet libero eu neque facilisis, ac pretium nisi dignissim. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.\n"

// GetIpsum returns the provided number of paragraphs of 'Lorem ipsum' text in byte form.
func GetIpsum(paragraphs int) []byte {
	output := Ipsum
	for i := 0; i < paragraphs; i++ {
		output += Ipsum
	}
	return []byte(output)
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
