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

func assertStringEquals[T string](t *testing.T, expected, got T) {
	t.Helper()
	if expected != got {
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
	five := Money{}.Dollar(5)
	assertEquals(t, Money{}.Dollar(10), five.Times(2))
	assertEquals(t, Money{}.Dollar(15), five.Times(3))
}

func TestEquality(t *testing.T) {
	df := Money{}.Dollar(5)
	assertTrue(t, df.Equals(Money{}.Dollar(5)))
	assertFalse(t, df.Equals(Money{}.Dollar(6)))
	ff := Money{amount: 5, currency: "CHF"}
	assertTrue(t, ff.Equals(Money{amount: 5, currency: "CHF"}))
	assertFalse(t, ff.Equals(Money{amount: 6, currency: "CHF"}))

	assertFalse(t, ff.Equals(df))
}

func TestFrancMultiplication(t *testing.T) {
	five := Money{amount: 5, currency: "CHF"}
	assertEquals(t, Money{amount: 10, currency: "CHF"}, five.Times(2))
	assertEquals(t, Money{amount: 15, currency: "CHF"}, five.Times(3))
}

func TestCurrency(t *testing.T) {
	assertStringEquals(t, "USD", Money{}.Dollar(5).GetCurrency())
	assertStringEquals(t, "CHF", Money{amount: 6, currency: "CHF"}.GetCurrency())
}
