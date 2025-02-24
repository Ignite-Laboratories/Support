// Package constraints provides a singular point for referencing 'Numeric' types of both Integer and Float.
package constraints

import (
	"golang.org/x/exp/constraints"
)

// Numeric represents any integer or floating-point type.
type Numeric interface {
	Integer | Float
}

// Integer represents any integer type.
type Integer interface {
	constraints.Integer
}

// Float represents any floating-point type.
type Float interface {
	constraints.Float
}
