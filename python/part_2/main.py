class TestCase:
    def __init__(self, name: str):
        self.name = name

    def setup(self):
        pass

    def teardown(self):
        pass

    def run(self):
        self.setup()
        method = getattr(self, self.name)
        method()
        self.teardown()


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

    def teardown(self):
        self.log = self.log + "teardown "


class TestCaseTest(TestCase):
    def test_template_method(self):
        self.test = WasRun("test_method")
        self.test.run()
        assert "setup test_method teardown " == self.test.log


TestCaseTest("test_template_method").run()
