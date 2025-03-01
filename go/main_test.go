package main

import "testing"

type Comparable interface {
	Equals(obj Dollar) bool
}

func assertEquals(t *testing.T, expected Comparable, got Dollar) {
	t.Helper()
	if !expected.Equals(got) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func assertTrue(t *testing.T, r bool) {
	t.Helper()
	if !r {
		t.Error("expected true, got false")
	}
}

func assertFalse(t *testing.T, r bool) {
	t.Helper()
	if r {
		t.Error("expected false, got true")
	}
}

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	assertEquals(t, Dollar{10}, five.Times(2))
	assertEquals(t, Dollar{15}, five.Times(3))
}

func TestEquality(t *testing.T) {
	dollar := Dollar{5}
	assertTrue(t, dollar.Equals(Dollar{5}))
	assertFalse(t, dollar.Equals(Dollar{6}))
}
