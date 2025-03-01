package main

import "testing"

func assertEquals(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func assertTrue(t *testing.T, r bool) {
	if !r {
		t.Error("expected true, got false")
	}
}

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	product := five.Times(2)
	assertEquals(t, 10, product.amount)
	product = five.Times(3)
	assertEquals(t, 15, product.amount)
}

func TestEquality(t *testing.T) {
	dollar := Dollar{5}
	assertTrue(t, dollar.Equals(Dollar{5}))
}
