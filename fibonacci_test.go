package main

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	for _, tt := range []struct {
		in, out int
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3},
		{5, 5}, {6, 8}, {7, 13}, {8, 21}, {9, 34},
	} {
		t.Run(fmt.Sprintf("Fibonacci(%d)", tt.in), func(t *testing.T) {
			if got := Fibonacci(tt.in); got != tt.out {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.in, got, tt.out)
			}
		})
	}
}
