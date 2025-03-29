// Package operate provides general purpose mathematical operations for the JanOS ecosystem.
package operate

import (
	"fmt"
	"github.com/ignite-laboratories/core"
)

// Operators provides a slice of the permissible operators of this library.
var Operators = []Operator{Add, Subtract, Multiply, Divide, XOR, AND, NAND, OR}

// Operator is a type of mathematical operation to be applied to integers.
type Operator int

const (
	// Add represents a+b
	Add Operator = iota
	// Subtract represents a-b
	Subtract
	// Multiply represents a*b
	Multiply
	// Divide represents a/b
	Divide
	// XOR represents a^b
	XOR
	// AND represents a&b
	AND
	// NAND represents ^(a&b)
	NAND
	// OR represents a|b
	OR
)

// String provides a quick reference of what each operation performs "under the hood."
func (op Operator) String() string {
	switch op {
	case Add:
		return "a+b"
	case Subtract:
		return "a-b"
	case Multiply:
		return "a*b"
	case Divide:
		return "a/b"
	case XOR:
		return "a^b"
	case AND:
		return "a&b"
	case NAND:
		return "^(a&b)"
	case OR:
		return "a|b"
	default:
		panic(fmt.Sprintf("Invalid operation: %v", op))
	}
}

// OnEach applies the provided value using the passed operator for each member of the source integers.
// Every operation is performed within the bounds of the provided type.  This means overflow and underflow
// will occur if you cross beyond the boundaries of your provided type.  This is an intentional design for Spark.
// NOTE: Division loses precision as we specifically only operate on integers here!
// We constrain to integers only as the bitwise operators are not readily available on non-integer types.
func OnEach[T core.Integer](data []T, operation Operator, value T) []T {
	out := make([]T, len(data))
	for i, v := range data {
		out[i] = Operate(v, operation, value)
	}
	return out
}

// Operate applies the provided operation between a and b.
func Operate[T core.Integer](a T, operation Operator, b T) T {
	switch operation {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	case Divide:
		return a / b
	case XOR:
		return a ^ b
	case AND:
		return a & b
	case NAND:
		return ^(a & b)
	case OR:
		return a | b
	default:
		panic(fmt.Sprintf("Invalid operation: %v", operation))
	}
}

// GetAverage calculates the average of a slice of integer values and returns the result in the slice's type.
func GetAverage[T core.Integer](data ...T) T {
	if len(data) == 0 {
		return T(0)
	}
	total := uint64(0)
	for _, v := range data {
		total += uint64(v)
	}
	return T(total / uint64(len(data)))
}
