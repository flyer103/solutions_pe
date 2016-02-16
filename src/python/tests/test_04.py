"""
"""

from ..solution_04 import method0, method1

EXPECTED = 906609


def test_solution_0(benchmark):
    data = benchmark(method0)

    assert data == EXPECTED


def test_solution_1(benchmark):
    data = benchmark(method1)

    assert data == EXPECTED
