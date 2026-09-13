package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	// test
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
