package api

// Arithmetic
type Arithmetic interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}
