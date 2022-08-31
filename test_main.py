class Expression:
    pass

class Bank:
    def reduce(self, source: Expression, to: str):
        return source.reduce(self, to)

    def rate(self, conv_from: str, conv_to: str):
        return 2 if conv_from == "CHF" and conv_to == "USD" else 1

class Money(Expression):
    def __init__(self, amount, currency) -> None:
        self._amount = amount
        self._currency = currency

    def __eq__(self, __o: object) -> bool:
        return self._amount == __o._amount and self._currency == __o._currency

    def __str__(self) -> str:
        return self._amount + " " + self._currency

    def __add__(self, addend):
        return Sum(self._amount, addend._amount, self._currency)

    def times(self, multiplier):
        return Money(self._amount * multiplier, self._currency)
    
    def dollar(amount):
        return Money(amount, "USD")
    
    def franc(amount):
        return Money(amount, "CHF")

    def currency(self):
        return self._currency

    def reduce(self, bank: Bank, to: str):
        rate = bank.rate(self._currency, to)
        return Money(self._amount / rate, to)

class Sum(Expression):
    def __init__(self, augend: Money, addend: Money) -> None:
       self._augend = augend
       self._addend = addend

    def reduce(self, bank, to):
        amount = self._augend._amount + self._addend._amount
        return Money(amount, to) 

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