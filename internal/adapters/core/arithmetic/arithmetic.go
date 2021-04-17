// Package arithmetic is just a placeholder to show where
// youre core/domain should sit. Packages within this
// "core" directory should contain only business logic.
package arithmetic

import ()

// Adapter is compatible with the CorePort
type Adapter struct {
}

// NewAdapter creates an arithmetic adapter that is
// compatible with CorePort
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Addition adds two numbers together
func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

// Subtraction subtracts one number from the other
func (arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

// Multiplication multiplies one number by the other
func (arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

// Division divides one number by the other
func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
