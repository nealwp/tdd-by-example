package main

import "testing"

func assertEquals(t *testing.T, got int, expected int) {
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	five.times(2)
	assertEquals(t, 10, five.amount)
}
