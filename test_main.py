class Dollar:
    def __init__(self, amount) -> None:
        self._amount = amount
    
    def times(self, multiplier):
        return Dollar(self._amount * multiplier)

    def __eq__(self, __o: object) -> bool:
        return self._amount == __o._amount

class Franc:
    def __init__(self, amount) -> None:
        self._amount = amount

    def times(self, multiplier):
        return Franc(self._amount * multiplier)

    def __eq__(self, __o: object) -> bool:
        return self._amount == __o._amount

def test_multiplication():
    five = Dollar(5)
    assert Dollar(10) == five.times(2)
    assert Dollar(15) == five.times(3)

def test_franc_multiplication():
    five = Franc(5)
    assert Franc(10) == five.times(2)
    assert Franc(15) == five.times(3)

def test_equality():
    assert Dollar(5) == Dollar(5)
    assert not Dollar(5) == Dollar(6)