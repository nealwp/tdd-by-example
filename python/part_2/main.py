class TestCase:
    def __init__(self, name: str):
        self.name = name

    def run(self):
        method = getattr(self, self.name)
        method()


class WasRun(TestCase):
    def __init__(self, name: str):
        self.was_run = None
        TestCase.__init__(self, name=name)

    def test_method(self):
        self.was_run = 1


test = WasRun("test_method")
print(test.was_run)
test.run()
print(test.was_run)
