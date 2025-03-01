package main

import "testing"

type Comparable[T any] interface {
	Equals(obj T) bool
}

func assertEquals[T Comparable[T]](t *testing.T, expected Comparable[T], got T) {
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

func TestFrancMultiplication(t *testing.T) {
	five := Franc{5}
	assertEquals(t, Franc{10}, five.Times(2))
	assertEquals(t, Franc{15}, five.Times(3))
}
