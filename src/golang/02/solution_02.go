// Each new term in the Fibonacci sequence is generated by adding the previous
// two terms. By starting with 1 and 2, the first 10 terms will be:
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed
// four million, find the sum of the even-valued terms.
// Answer: 4613732

package main

import (
	"fmt"
)

var (
	maxNum = 4000000
)

func fib_tail(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	}

	x := 0
	y := 1
	for i := 2; i <= num; i++ {
		x, y = y, x+y
	}

	return y
}

func method_0() (res int) {
	i := 0
	for {
		data := fib_tail(i)
		if data > maxNum {
			break
		}

		if data%2 == 0 {
			res += data
		}

		i += 1
	}

	return res
}

func method_1() (res int) {
	x, y := 1, 1

	for {
		data := x + y
		if data > maxNum {
			break
		}

		res += data
		x, y = x+2*y, 2*x+3*y
	}

	return res
}

func main() {
	fmt.Println(method_0())
	fmt.Println(method_1())
}