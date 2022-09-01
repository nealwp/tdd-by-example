class Bank:
    def reduce(self, source, to: str):
        return source.reduce(self, to)

    def rate(self, conv_from: str, conv_to: str):
        return 2 if conv_from == "CHF" and conv_to == "USD" else 1

class Expression:
    pass

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