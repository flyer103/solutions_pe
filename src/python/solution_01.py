"""
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.

Answer: 233168
"""


class Solution0:

    @staticmethod
    def compute(num=1000):
        res = 0
        for i in range(1, num):
            if i % 3 == 0 or i % 5 == 0:
                res += i

        return res


class Solution1:

    @staticmethod
    def compute(num=1000):
        num -= 1

        num_3 = int(num/3)
        num_5 = int(num/5)
        num_15 = int(num/15)

        res = 3*(num_3+1)*num_3/2 \
              + 5*(num_5+1)*num_5/2 \
              - 15*(num_15+1)*num_15/2 \

        return int(res)


if __name__ == '__main__':
    print(Solution0.compute())
    print(Solution1.compute())
