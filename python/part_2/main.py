class TestCase:
    def __init__(self, name: str):
        self.name = name

    def setup(self):
        pass

    def teardown(self):
        pass

    def run(self):
        result = TestResult()
        result.test_started()
        self.setup()
        method = getattr(self, self.name)
        method()
        self.teardown()
        return result


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


class TestCaseTest(TestCase):
    def test_template_method(self):
        self.test = WasRun("test_method")
        self.test.run()
        assert "setup test_method teardown " == self.test.log

    def test_result(self):
        test = WasRun("test_method")
        result = test.run()
        assert "1 run, 0 failed" == result.summary()

    def test_failed_result(self):
        test = WasRun("test_broken_method")
        result = test.run()
        assert "1 run, 1 failed" == result.summary()

    def test_failed_result_formatting(self):
        result = TestResult()
        result.test_started()
        result.test_failed()
        assert "1 run, 1 failed" == result.summary()


TestCaseTest("test_template_method").run()
TestCaseTest("test_result").run()
# TestCaseTest("test_failed_result").run()
TestCaseTest("test_failed_result_formatting").run()
