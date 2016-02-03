"""
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

Answer: 6857
"""

import math


class Solution0:

    @staticmethod
    def is_prime(num):
        if not isinstance(num, int) or num <= 1:
            return False

        end = int(math.sqrt(num))
        for i in range(2, end+1):
            if num % i == 0:
                return False

        return True

    @staticmethod
    def get_max_prime(num=600851475143):
        end = int(math.sqrt(num))
        for i in range(end, 1, -1):
            if num % i == 0 and Solution0.is_prime(i):
                return i

        return -1


class Solution1:

    @staticmethod
    def is_prime(num):
        if not isinstance(num, int) or num <= 1:
            return False

        end = int(math.sqrt(num))
        for i in range(2, end+1):
            if num % i == 0:
                return False

        return True

    @staticmethod
    def get_max_prime(num=600851475143):
        for i in range(2, int(num/2)):
            if num % i == 0:
                num = int(num/i)
                if Solution1.is_prime(num):
                    return num

                Solution1.get_max_prime(num)


if __name__ == '__main__':
    solution_0 = Solution0()
    print(solution_0.get_max_prime())

    solution_1 = Solution1()
    print(solution_1.get_max_prime())
