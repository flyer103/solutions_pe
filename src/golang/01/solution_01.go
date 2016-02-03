// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
// Answer: 233168

package main

import (
	"flag"
	"fmt"
)

var (
	number = flag.Int("number", 1000, "input number")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(compute_0(*number))
	fmt.Println(compute_1(*number))
}

func compute_0(number int) int {
	number -= 1
	num_3 := int(number / 3)
	num_5 := int(number / 5)
	num_15 := int(number / 15)

	res := 3*(num_3+1)*num_3/2 +
		5*(num_5+1)*num_5/2 -
		15*(num_15+1)*num_15/2

	return res
}

func compute_1(number int) (res int) {
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}

	return res
}
