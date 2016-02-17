"""
"""

from functools import partial

from ..solution_05 import is_prime, get_all_primes, \
    get_max_multiple, get_smallest_number


def test_is_prime():
    res_1 = is_prime(1)
    assert not res_1

    res_2 = is_prime(2)
    assert res_2

    res_3 = is_prime(3)
    assert res_3

    res_1000 = is_prime(1000)
    assert not res_1000


def test_get_all_primes(benchmark):
    expected = [2, 3, 5, 7, 11, 13, 17, 19]
    res = benchmark(partial(get_all_primes, 1, 20))

    assert sorted(res) == expected


def test_get_max_multile(benchmark):
    expected = [11, 12, 13, 14, 15, 16, 17, 18, 19, 20]
    res = benchmark(partial(get_max_multiple, 1, 20))

    assert sorted(res) == expected


def test_get_smallest_number(benchmark):
    expected = 232792560
    res = benchmark(partial(get_smallest_number, 1, 20))

    assert res == expected
