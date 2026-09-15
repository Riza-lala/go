package divmod

import (
	"errors"
	"testing"
)

func TestDivMod(t *testing.T) {
	testCases := []struct {
		name      string
		dividend  int
		divisor   int
		quotient  int
		remainder int
	}{
		{name: "exact", dividend: 12, divisor: 3, quotient: 4},
		{name: "remainder", dividend: 17, divisor: 5, quotient: 3, remainder: 2},
		{name: "negative dividend", dividend: -17, divisor: 5, quotient: -3, remainder: -2},
		{name: "negative divisor", dividend: 17, divisor: -5, quotient: -3, remainder: 2},
		{name: "zero dividend", dividend: 0, divisor: 7, quotient: 0, remainder: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			quotient, remainder, err := DivMod(tc.dividend, tc.divisor)
			if err != nil {
				t.Fatalf("DivMod(%d, %d) returned error: %v", tc.dividend, tc.divisor, err)
			}

			if quotient != tc.quotient || remainder != tc.remainder {
				t.Fatalf(
					"DivMod(%d, %d) = (%d, %d), want (%d, %d)",
					tc.dividend,
					tc.divisor,
					quotient,
					remainder,
					tc.quotient,
					tc.remainder,
				)
			}
		})
	}
}

func TestDivModByZero(t *testing.T) {
	quotient, remainder, err := DivMod(42, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Fatalf("DivMod(42, 0) error = %v, want ErrDivisionByZero", err)
	}

	if quotient != 0 || remainder != 0 {
		t.Fatalf("DivMod(42, 0) = (%d, %d), want (0, 0)", quotient, remainder)
	}
}
