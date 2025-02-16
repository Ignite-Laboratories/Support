package operate

import (
	"golang.org/x/exp/constraints"
	"log"
)

// Operator is a type of mathematical operate to be applied to integers.
type Operator int

const (
	// Add represents a+b
	Add Operator = iota
	// Subtract represents a-b
	Subtract
	// Multiply represents a*b
	Multiply
	// XOR represents a^b
	XOR
	// AND represents a&b
	AND
	// NAND represents ^(a&b)
	NAND
	// OR represents a|b
	OR
)

// OnEach applies the provided value using the passed operate for each member of the source slice of integers.
// NOTE: As this does not handle floating point numbers gracefully, division is not a provided operate
// to avoid blissfully unaware confusion.
func OnEach[T constraints.Integer](data []T, operation Operator, value T) []T {
	op := func(a, b T) T {
		switch operation {
		case Add:
			return a + b
		case Subtract:
			return a - b
		case Multiply:
			return a * b
		case XOR:
			return a ^ b
		case AND:
			return a & b
		case NAND:
			return ^(a & b)
		case OR:
			return a | b
		default:
			log.Panicf("Invalid operate %v", operation)
			return 0
		}
	}
	for i, v := range data {
		data[i] = op(v, value)
	}
	return data
}

// GetAverage calculates the average of a slice of integer values.
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
