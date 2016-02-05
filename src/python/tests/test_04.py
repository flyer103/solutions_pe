"""
"""

from ..solution_04 import method0

EXPECTED = 906609


def test_solution_0(benchmark):
    data = benchmark(method0)

    assert data == EXPECTED
