from main import *

def test_multiplication():
    five = Money.dollar(5)
    assert Money.dollar(10) == five.times(2)
    assert Money.dollar(15) == five.times(3)
    five = Money.franc(5)
    assert Money.franc(10) == five.times(2)
    assert Money.franc(15) == five.times(3)

def test_equality():
    assert Money.dollar(5) == Money.dollar(5)
    assert not Money.dollar(5) == Money.dollar(6)
    assert not Money.franc(5) == Money.dollar(5)

def test_currency():
    assert "USD" == Money.dollar(1).currency()
    assert "CHF" == Money.franc(1).currency()

def test_simple_addition():
    five = Money.dollar(5)
    sum = Sum(five, five)
    bank = Bank()
    reduced = bank.reduce(sum, "USD")
    assert reduced == Money.dollar(10)

def test_reduce_sum():
    sum = Sum(Money.dollar(3), Money.dollar(4))
    bank = Bank()
    result = bank.reduce(sum, "USD")
    assert Money.dollar(7) == result

def test_plus_returns_sum():
    five = Money.dollar(5)
    sum = Sum(five, five)
    assert five == sum._augend
    assert five == sum._addend

def test_reduce_money():
    bank = Bank()
    result = bank.reduce(Money.dollar(1), "USD")
    assert Money.dollar(1) == result

def test_reduce_money_different_currency():
    bank = Bank()
    bank.add_rate("CHF", "USD", 2)
    result = bank.reduce(Money.franc(2), "USD")
    assert Money.dollar(1) == result

def test_identity_rate():
    assert 1 == Bank().rate("USD", "USD")

def test_mixed_addition():
    fiveBucks = Money.dollar(5)
    tenFrancs = Money.franc(10)
    bank = Bank()
    bank.add_rate("CHF", "USD", 2)
    result = bank.reduce(Sum(fiveBucks, tenFrancs), "USD")
    assert Money.dollar(10) == result