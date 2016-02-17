"""
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
Answer: 232792560
"""
import math


def is_prime(num):
    if num <= 1:
        return False

    sqrt_num = int(math.sqrt(num))
    for i in range(2, sqrt_num+1):
        if num % i == 0:
            return False

    return True


def get_all_primes(start=1, end=20):
    primes = []
    for i in range(start, end+1):
        if is_prime(i):
            primes.append(i)

    return primes


def get_max_multiple(start=1, end=20):
    nums = set()
    for i in range(start, end+1):
        max_multiple = i
        for j in range(2, end+1):
            multiple = i * j
            if max_multiple < multiple <= end:
                max_multiple = multiple
        nums.add(max_multiple)

    return list(nums)


def get_smallest_number(start=1, end=20):
    all_primes = get_all_primes(start, end)
    max_multiples = get_max_multiple(start, end)

    min_product_primes = all_primes[0]
    for i in all_primes[1:]:
        min_product_primes *= i

    multiplier = 1
    while True:
        product = min_product_primes * multiplier
        for i in max_multiples:
            if product % i != 0:
                break
        else:
            return product

        multiplier += 1


if __name__ == '__main__':
    print(get_smallest_number())
