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
	five := Dollar{BaseMoney{amount: 5, currency: "USD"}}
	assertEquals(t, Dollar{BaseMoney{amount: 10, currency: "USD"}}, five.Times(2))
	assertEquals(t, Dollar{BaseMoney{amount: 15, currency: "USD"}}, five.Times(3))
}

func TestEquality(t *testing.T) {
	df := Dollar{BaseMoney{amount: 5, currency: "USD"}}
	assertTrue(t, df.Equals(Dollar{BaseMoney{amount: 5, currency: "USD"}}))
	assertFalse(t, df.Equals(Dollar{BaseMoney{amount: 6, currency: "USD"}}))
	ff := Franc{BaseMoney{amount: 5, currency: "CHF"}}
	assertTrue(t, ff.Equals(Franc{BaseMoney{amount: 5, currency: "CHF"}}))
	assertFalse(t, ff.Equals(Franc{BaseMoney{amount: 6, currency: "CHF"}}))

	assertFalse(t, ff.Equals(df))
}

func TestFrancMultiplication(t *testing.T) {
	five := Franc{BaseMoney{amount: 5, currency: "CHF"}}
	assertEquals(t, Franc{BaseMoney{amount: 10, currency: "CHF"}}, five.Times(2))
	assertEquals(t, Franc{BaseMoney{amount: 15, currency: "CHF"}}, five.Times(3))
}

func TestCurrency(t *testing.T) {
	assertStringEquals(t, "USD", Dollar{BaseMoney{amount: 5, currency: "USD"}}.GetCurrency())
	assertStringEquals(t, "CHF", Franc{BaseMoney{amount: 6, currency: "CHF"}}.GetCurrency())
}
