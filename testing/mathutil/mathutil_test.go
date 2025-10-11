package mathutil

import (
	"testing"
)

func TestAdd(t *testing.T){
	got := Add(2,3)
	want := 5

	if got != want {
		t.Errorf("Add(2,3) = %d; want %d", got, want)
	}
}

func TestSubtract(t *testing.T){
	got := Subtract(5,3)
	want := 2

	if got != want {
		t.Errorf("Subtract(5,3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{4, 0, 0},
	}

	for _, tt := range tests {
		t.Run("Multiply Test", func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Multiply(%d,%d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected error when dividing by zero, got nil")
	}

	got, err := Divide(10, 2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	want := 5
	if got != want {
		t.Errorf("Divide(10,2) = %d; want %d", got, want)
	}
}