package main

import "testing"

type Comparable interface {
	Equals(obj interface{}) bool
	GetAmount() int
}

func assertEquals(t *testing.T, expected interface{}, got interface{}) {
	t.Helper()
	switch e := expected.(type) {
	case int, string:
		if e != got {
			t.Errorf("expected: %v, got: %v", e, got)
		}
	case Comparable:
		if !e.Equals(got) {
			t.Errorf("expected: %v, got: %v", e, got)
		}
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

	ff := Money{}.Franc(5)
	assertTrue(t, ff.Equals(Money{}.Franc(5)))
	assertFalse(t, ff.Equals(Money{}.Franc(6)))

	assertFalse(t, ff.Equals(df))
}

func TestFrancMultiplication(t *testing.T) {
	five := Money{}.Franc(5)
	assertEquals(t, Money{}.Franc(10), five.Times(2))
	assertEquals(t, Money{}.Franc(15), five.Times(3))
}

func TestCurrency(t *testing.T) {
	assertEquals(t, "USD", Money{}.Dollar(5).GetCurrency())
	assertEquals(t, "CHF", Money{}.Franc(6).GetCurrency())
}

func TestSimpleAddition(t *testing.T) {
	five := Money{}.Dollar(5)
	sum := five.Plus(five)
	bank := Bank{}
	reduced := bank.Reduce(sum, "USD")
	assertEquals(t, Money{}.Dollar(10), reduced)
}

func TestPlusReturnsSum(t *testing.T) {
	five := Money{}.Dollar(5)
	result := five.Plus(five)
	sum := result.(Sum)
	assertEquals(t, five, sum.Augend)
	assertEquals(t, five, sum.Addend)
}

func TestReduceMoney(t *testing.T) {
	bank := Bank{}
	result := bank.Reduce(Money{}.Dollar(1), "USD")
	assertEquals(t, Money{}.Dollar(1), result)
}

func TestIdentityRate(t *testing.T) {
	bank := Bank{}
	result := bank.Rate("USD", "USD")
	assertEquals(t, 1, result)
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := Bank{}
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(Money{}.Franc(2), "USD")
	assertEquals(t, Money{}.Dollar(1), result)
}
