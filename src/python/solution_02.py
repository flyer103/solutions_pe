"""
Each new term in the Fibonacci sequence is generated by adding the previous
two terms. By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed
four million, find the sum of the even-valued terms.

Answer: 4613732
"""


class Solution0:

    MAX = 4000000

    def fib_tail(self, n=100):
        if n == 1:
            return 1
        elif n == 2:
            return 2

        i = 3
        before = 1
        res = 2
        while i <= n:
            before, res = res, res+before
            i += 1

        return res

    def sum_fib(self):
        res = 0
        n = 1
        while True:
            data = self.fib_tail(n)
            if data > self.MAX:
                break

            if data % 2 == 0:
                res += data

            n += 1

        return res


class Solution1:
    """
    The Fibonacci series is:
    1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610...
    Now, replacing an odd number with O and an even with E, we get:
    O, O, E, O, O, E, O, O, E, O, O, E, O, O, E...
    And so each third number is even.
    We don't need to calculate the odd numbers.
    Starting from an two odd terms x, y, the series is:
    x, y, x + y, x + 2y, 2x + 3y, 3x + 5y
    """

    MAX = 4000000

    def sum_fib(self):
        res = 0
        x = 1
        y = 1
        while True:
            data = x + y
            if data > self.MAX:
                break

            res += data
            x, y = x+2*y, 2*x+3*y

        return res


if __name__ == '__main__':
    solution_0 = Solution0()
    print(solution_0.sum_fib())

    solution_1 = Solution1()
    print(solution_1.sum_fib())
