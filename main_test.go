package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestFoo(t *testing.T) {
	expected := 3
	if actual := Sum(1, 2); actual != expected {
		t.Fatalf("expected = %d, actual = %d", actual, expected)
	}
}
