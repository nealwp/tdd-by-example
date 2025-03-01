package main

import "testing"

type Comparable[T any] interface {
	Equals(obj T) bool
	GetAmount() int
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
	five := Dollar{BaseMoney{5}}
	assertEquals(t, Dollar{BaseMoney{10}}, five.Times(2))
	assertEquals(t, Dollar{BaseMoney{15}}, five.Times(3))
}

func TestEquality(t *testing.T) {
	df := Dollar{BaseMoney{5}}
	assertTrue(t, df.Equals(Dollar{BaseMoney{5}}))
	assertFalse(t, df.Equals(Dollar{BaseMoney{6}}))
	ff := Franc{BaseMoney{5}}
	assertTrue(t, ff.Equals(Franc{BaseMoney{5}}))
	assertFalse(t, ff.Equals(Franc{BaseMoney{6}}))
}

func TestFrancMultiplication(t *testing.T) {
	five := Franc{BaseMoney{5}}
	assertEquals(t, Franc{BaseMoney{10}}, five.Times(2))
	assertEquals(t, Franc{BaseMoney{15}}, five.Times(3))
}
