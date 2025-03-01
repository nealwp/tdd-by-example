package main

import "testing"

func assertEquals(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	product := five.Times(2)
	assertEquals(t, 10, product.amount)
	product = five.Times(3)
	assertEquals(t, 15, product.amount)
}
