class Money:
    def __init__(self, amount, currency) -> None:
        self._amount = amount
        self._currency = currency

    def __eq__(self, __o: object) -> bool:
        return self._amount == __o._amount and self._currency == __o._currency

    def __str__(self) -> str:
        return self._amount + " " + self._currency

    def times(self, multiplier):
        return Money(self._amount * multiplier, self._currency)
    
    def dollar(amount):
        return Money(amount, "USD")
    
    def franc(amount):
        return Money(amount, "CHF")

    def currency(self):
        return self._currency

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
    sum = Money.dollar(5) + Money.dollar(5)
    assert sum == Money.dollar(10)