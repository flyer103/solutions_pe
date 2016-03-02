"""
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten
natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one
hundred natural numbers and the square of the sum.

Answer: 25164150
"""


def sum_of_square(nums):
    ret = 0
    for i in nums:
        ret += i ** 2

    return ret


def square_of_sum(nums):
    return sum(nums) ** 2


def method_0(nums):
    return square_of_sum(nums) - sum_of_square(nums)


if __name__ == '__main__':
    nums = list(range(1, 101))
    print(method_0(nums))
