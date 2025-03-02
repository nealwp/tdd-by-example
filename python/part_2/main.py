class TestCase:
    def __init__(self, name: str):
        self.name = name

    def setup(self):
        pass

    def run(self):
        self.setup()
        method = getattr(self, self.name)
        method()


class WasRun(TestCase):
    def __init__(self, name: str):
        self.setup()
        TestCase.__init__(self, name=name)

    def setup(self):
        self.was_run = False
        self.log = "setup "

    def test_method(self):
        self.was_run = True
        self.log = self.log + "test_method "


class TestCaseTest(TestCase):
    def setup(self):
        self.test = WasRun("test_method")

    def test_setup(self):
        self.test.run()
        assert "setup test_method " == self.test.log

    def test_running(self):
        self.test.run()
        assert self.test.was_run


TestCaseTest("test_running").run()
TestCaseTest("test_setup").run()
