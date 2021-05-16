package arithmetic

// Arith impliments the Arithmetic interface
type Arith struct {
}

// NewArith creates a new Arith
func New() *Arith {
	return &Arith{}
}

// Addition gets the result of adding parameters a and b
func (arith Arith) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

// Subtraction gets the result of subtracting parameters a and b
func (arith Arith) Subtraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

// Multiplication gets the result of multiplying parameters a and b
func (arith Arith) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

// Division gets the result of dividing parameters a and b
func (arith Arith) Division(a int32, b int32) (int32, error) {
	return a / b, nil
}
