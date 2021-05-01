package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}
