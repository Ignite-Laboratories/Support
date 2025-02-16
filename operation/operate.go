package operation

import (
	"golang.org/x/exp/constraints"
	"log"
)

type Operation int

const (
	Addition Operation = iota
	Subtraction
	Multiplication
	XOR
	And
	Nand
	Or
)

// OnEach applies the provided value using the passed operation for each member of the source data.
// NOTE: As this does not handle floating point numbers gracefully, division is not a provided operation
// to avoid blissfully unaware confusion.
func OnEach[T constraints.Integer](data []T, value T, operation Operation) []T {
	op := func(a, b T) T {
		switch operation {
		case Addition:
			return a + b
		case Subtraction:
			return a - b
		case Multiplication:
			return a * b
		case XOR:
			return a ^ b
		case And:
			return a & b
		case Nand:
			return ^(a & b)
		case Or:
			return a | b
		default:
			log.Panicf("Invalid operation %v", operation)
			return 0
		}
	}
	for i, v := range data {
		data[i] = op(v, value)
	}
	return data
}

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
