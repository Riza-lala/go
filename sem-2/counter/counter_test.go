package counter

import "testing"

func TestCounter(t *testing.T) {
	next := New(10, 3)

	want := []int{13, 16, 19, 22}
	for i, expected := range want {
		if got := next(); got != expected {
			t.Fatalf("call %d returned %d, want %d", i+1, got, expected)
		}
	}
}

func TestCountersAreIndependent(t *testing.T) {
	first := New(0, 1)
	second := New(100, -10)

	if got := first(); got != 1 {
		t.Fatalf("first() = %d, want 1", got)
	}

	if got := first(); got != 2 {
		t.Fatalf("first() = %d, want 2", got)
	}

	if got := second(); got != 90 {
		t.Fatalf("second() = %d, want 90", got)
	}

	if got := first(); got != 3 {
		t.Fatalf("first() after second() = %d, want 3", got)
	}

	if got := second(); got != 80 {
		t.Fatalf("second() = %d, want 80", got)
	}
}

func TestZeroStep(t *testing.T) {
	next := New(42, 0)
	for range 3 {
		if got := next(); got != 42 {
			t.Fatalf("New(42, 0) returned %d, want 42", got)
		}
	}
}
