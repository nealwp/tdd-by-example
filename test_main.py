class Dollar:
    def __init__(self, amount) -> None:
        self._amount = amount
    
    def times(self, multiplier):
        return Dollar(self._amount * multiplier)
    
    def equals(self, dollar):
        return self._amount == dollar._amount
def test_multiplication():
    five = Dollar(5)
    assert Dollar(10).equals(five.times(2))
    assert Dollar(15).equals(five.times(3))

def test_franc_multiplication():
    five = Franc(5)
    assert Franc(10).equals(five.times(2))
    assert Franc(15).equals(five.times(3))

def test_equality():
    assert Dollar(5).equals(Dollar(5))
    assert not Dollar(5).equals(Dollar(6))