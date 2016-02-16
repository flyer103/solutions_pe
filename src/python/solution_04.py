"""
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
Answer: 906609
"""


def is_palindrome(num):
    num = str(num)
    if num == num[::-1]:
        return True

    return False


def method0():
    maxNum = 0

    for i in range(999, 99, -1):
        for j in range(999, 99, -1):
            product = i * j
            if product > maxNum and is_palindrome(product):
                maxNum = product

    return maxNum


def method1():
    """
    The palindrome can be written as:
        abccba
    Which then simpifies to:
        100000a + 10000b + 1000c + 100c + 10b + a
    And then:
        100001a + 10010b + 1100c
    Factoring out 11, you get:
        11(9091a + 910b + 100c)
    So the palindrome must be divisible by 11.
    Seeing as 11 is prime, at least one of the numbers must be divisible by 11.
    """
    maxNum = 0

    for i in range(999, 99, -1):
        for j in range(990, 99, -11):
            product = i * j
            if product > maxNum and is_palindrome(product):
                maxNum = product

    return maxNum


if __name__ == '__main__':
    print(method0())
    print(method1())
