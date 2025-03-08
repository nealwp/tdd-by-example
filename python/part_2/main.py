class TestCase:
    def __init__(self, name: str):
        self.name = name

    def setup(self):
        pass

    def teardown(self):
        pass

    def run(self, result):
        result.test_started()
        self.setup()

        try:
            method = getattr(self, self.name)
            method()
        except:
            result.test_failed()

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

    def test_broken_method(self):
        raise Exception

    def teardown(self):
        self.log = self.log + "teardown "


class TestResult:
    def __init__(self) -> None:
        self.run_count = 0
        self.error_count = 0

    def test_started(self):
        self.run_count += 1

    def test_failed(self):
        self.error_count += 1

    def summary(self):
        return f"{self.run_count} run, {self.error_count} failed"


class TestSuite:
    def __init__(self):
        self.tests = []

    def add(self, test):
        self.tests.append(test)

    def run(self, result):
        for test in self.tests:
            test.run(result)


class TestCaseTest(TestCase):
    def setup(self):
        self.result = TestResult()

    def test_template_method(self):
        test = WasRun("test_method")
        test.run(self.result)
        assert "setup test_method teardown " == test.log

    def test_result(self):
        test = WasRun("test_method")
        test.run(self.result)
        assert "1 run, 0 failed" == self.result.summary()

    def test_failed_result(self):
        test = WasRun("test_broken_method")
        test.run(self.result)
        assert "1 run, 1 failed" == self.result.summary()

    def test_failed_result_formatting(self):
        self.result.test_started()
        self.result.test_failed()
        assert "1 run, 1 failed" == self.result.summary()

    def test_suite(self):
        suite = TestSuite()
        suite.add(WasRun("test_method"))
        suite.add(WasRun("test_broken_method"))
        suite.run(self.result)
        assert "2 run, 1 failed" == self.result.summary()


suite = TestSuite()
suite.add(TestCaseTest("test_template_method"))
suite.add(TestCaseTest("test_result"))
suite.add(TestCaseTest("test_failed_result"))
suite.add(TestCaseTest("test_failed_result_formatting"))
suite.add(TestCaseTest("test_suite"))
result = TestResult()
suite.run(result)
print(result.summary())
