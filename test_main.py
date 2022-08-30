class Dollar:
    def __init__(self, amount) -> None:
        self.amount = amount
    
    def times(self, multiplier):
        self.amount *= multiplier
def test_multipication():
    five = Dollar(5)
    five.times(2)
    assert 10 == five.amount