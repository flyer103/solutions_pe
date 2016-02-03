"""
"""

from ..solution_03 import Solution0, Solution1


def _test_solution(cls):
    return cls().get_max_prime()


def test_solution_0(benchmark):
    data = benchmark(_test_solution, Solution0)

    assert data == 6857


def test_solution_1(benchmark):
    data = benchmark(_test_solution, Solution1)

    assert data == 6857
