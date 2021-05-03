package arithmetic

// Adapter implements the ArithmeticPort interface
type Adapter struct {
}

// NewAdapter creates a new Adapter
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Addition gets the result of adding parameters a and b
func (arith Adapter) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

// Subtraction gets the result of subtracting parameters a and b
func (arith Adapter) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

// Multiplication gets the result of multiplying parameters a and b
func (arith Adapter) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

// Division gets the result of dividing parameters a and b
func (arith Adapter) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
