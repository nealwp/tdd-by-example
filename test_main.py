class Dollar:
    def __init__(self, amount) -> None:
        self.amount = amount
    
    def times(self, multiplier):
        return Dollar(self.amount * multiplier)
def test_multipication():
    five = Dollar(5)
    product = five.times(2)
    assert 10 == product.amount
    product = five.times(3)
    assert 15 == product.amount