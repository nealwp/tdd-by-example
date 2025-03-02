class WasRun:
    def __init__(self, name: str):
        self.was_run = None

    def test_method(self):
        self.was_run = 1


test = WasRun("testMethod")
print(test.was_run)
test.test_method()
print(test.was_run)
