"""
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
Answer: 906609
"""

START = 100
END = 999


def is_palindrome(num):
    num_str = str(num)
    total = len(num_str)
    if total % 2 == 0:
        mid = int(total/2) - 1
    else:
        mid = int(total/2)

    i = 0
    last = total - 1
    for i in range(0, mid+1):
        if num_str[i] != num_str[last-i]:
            return False

    return True


def method0():
    numbers = []
    total = 0
    for i in range(START, END+1):
        for j in range(START, END+1):
            numbers.append(i*j)
            total += 1

    numbers = sorted(numbers)
    for i in range(total-1, -1, -1):
        if is_palindrome(numbers[i]):
            return numbers[i]

    return -1


if __name__ == '__main__':
    print(method0())
