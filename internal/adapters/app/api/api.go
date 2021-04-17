package api

import (
	"hex-arch-go-grpc/internal/ports"
)

// Adapter is compatible with the
type Adapter struct {
	db    ports.DbPort
	arith ports.ArithemeticPort
}

// NewAdapter creates an rpc adapter that is
// compatible with
func NewAdapter(db ports.DbPort, arith ports.ArithemeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

// GetAddition returns the result of adding two numbers
func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	var err error

	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

// GetSubtraction returns the result of subtracting one number form the other
func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	var err error

	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

// GetMultiplication returns the result of multiplying one number by the other
func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	var err error

	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

// GetDivision returns the result of dividing one number by the other
func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	var err error

	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
