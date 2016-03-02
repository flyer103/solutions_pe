// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.
// Find the difference between the sum of the squares of the first one
// hundred natural numbers and the square of the sum.

// Answer: 25164150

package main

import (
	"fmt"
)

func sumOfSquare(nums []int) (ret int) {
	for _, num := range nums {
		ret += num * num
	}

	return ret
}

func squareOfSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum * sum
}

func method0(nums []int) int {
	return squareOfSum(nums) - sumOfSquare(nums)
}

func main() {
	nums := []int{}
	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}

	fmt.Println(method0(nums))
}
