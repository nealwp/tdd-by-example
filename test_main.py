class Money:
    def __init__(self, amount) -> None:
        self._amount = amount

    def __eq__(self, __o: object) -> bool:
        return self._amount == __o._amount and type(self) == type(__o)

    def times(self, multiplier):
        return Money(self._amount * multiplier)
    
    def dollar(amount):
        return Dollar(amount)
    
    def franc(amount):
        return Franc(amount)

class Dollar(Money):
    
    def times(self, multiplier):
        return Dollar(self._amount * multiplier)

    def currency(self):
        return "USD"

class Franc(Money):

    def __init__(self, amount):
        self._amount = amount
        self._currency = "CHF"

    def times(self, multiplier):
        return Franc(self._amount * multiplier)

    def currency(self):
        return "CHF"

def test_multiplication():
    five = Money.dollar(5)
    assert Money.dollar(10) == five.times(2)
    assert Money.dollar(15) == five.times(3)

def test_franc_multiplication():
    five = Franc(5)
    assert Franc(10) == five.times(2)
    assert Franc(15) == five.times(3)

def test_equality():
    assert Money.dollar(5) == Money.dollar(5)
    assert not Money.dollar(5) == Money.dollar(6)
    assert Money.franc(5) == Money.franc(5)
    assert not Money.franc(5) == Money.franc(6)
    assert not Money.franc(5) == Money.dollar(5)

def test_currency():
    assert "USD" == Money.dollar(1).currency()
    assert "CHF" == Money.franc(1).currency()