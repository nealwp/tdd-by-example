class WasRun:
    def __init__(self, name: str):
        self.was_run = None
        self.name = name

    def test_method(self):
        self.was_run = 1

    def run(self):
        method = getattr(self, self.name)
        method()


test = WasRun("test_method")
print(test.was_run)
test.run()
print(test.was_run)
