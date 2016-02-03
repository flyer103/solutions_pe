"""
"""

from ..solution_02 import Solution0, Solution1


def _solution(cls):
    return cls.sum_fib()


def test_solution_0(benchmark):
    res = benchmark(_solution, Solution0)

    assert res == 4613732


def test_solution_1(benchmark):
    res = benchmark(_solution, Solution1)

    assert res == 4613732
