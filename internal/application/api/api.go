package api

import (
	"hex/internal/ports"
)

// Application implements the APIPort interface
type Application struct {
	db    ports.DbPort
	arith Arithmetic
}

// NewApplication creates a new Application
func NewApplication(db ports.DbPort, arith Arithmetic) *Application {
	return &Application{db: db, arith: arith}
}

// GetAddition gets the result of adding parameters a and b
func (apia Application) GetAddition(a, b int32) (int32, error) {
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

// GetSubtraction gets the result of subtracting parameters a and b
func (apia Application) GetSubtraction(a, b int32) (int32, error) {
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

// GetMultiplication gets the result of multiplying parameters a and b
func (apia Application) GetMultiplication(a, b int32) (int32, error) {
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

// GetDivision gets the result of dividing parameters a and b
func (apia Application) GetDivision(a, b int32) (int32, error) {
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
