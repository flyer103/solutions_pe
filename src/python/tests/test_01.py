"""
"""


from ..solution_01 import Solution0, Solution1


def _test(cls):
    return cls.compute()


def test_solution_0(benchmark):
    data = benchmark(_test, Solution0)

    assert data == 233168


def test_solution_1(benchmark):
    data = benchmark(_test, Solution1)

    assert data == 233168
