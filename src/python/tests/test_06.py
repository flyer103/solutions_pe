"""
"""

from functools import partial

from ..solution_06 import sum_of_square, square_of_sum, method_0


def test_sum_of_square(benchmark):
    nums = list(range(1, 11))

    expected = 385
    res = benchmark(partial(sum_of_square, nums))

    assert res == expected


def test_square_of_sum(benchmark):
    nums = list(range(1, 11))

    expected = 3025
    res = benchmark(partial(square_of_sum, nums))

    assert res == expected


def test_method_0(benchmark):
    nums = list(range(1, 11))

    expected = 2640
    res = benchmark(partial(method_0, nums))

    assert res == expected
