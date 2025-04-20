// Package performs simple mathematical functions
// such as addition and in the future subtraction.
package workspace_lib

import (
	"golang.org/x/exp/constraints"
)

// Number type interface that includes both integer and float
// from the constraint module.
type Number interface {
	constraints.Integer | constraints.Float
}

// AddNums adds two Number interface variables and returns the result as a Number interface.
// Inspired by [MathIsFun].
//
// It has two parameters: a and b of type Number and returns a Number type.
// The Number is used from the constraint module.
//
// [MathIsFun]: https://www.mathsisfun.com/numbers/addition.html
func AddNums[T Number](a, b T) T {
	return a + b
}
