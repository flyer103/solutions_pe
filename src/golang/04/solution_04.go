// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// Answer: 906609

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(num int64) bool {
	numStr := strconv.FormatInt(num, 10)
	arrStr := strings.Split(numStr, "")
	len := len(arrStr)

	var mid int
	if len%2 == 0 {
		mid = len/2 - 1
	} else {
		mid = len / 2
	}

	for i := 0; i <= mid; i++ {
		if arrStr[i] != arrStr[len-1-i] {
			return false
		}
	}

	return true
}

func method0() int64 {
	start := 100
	end := 999

	var maxNum int64
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			product := int64(i * j)
			if product > maxNum && isPalindrome(product) {
				maxNum = product
			}
		}
	}

	return maxNum
}

func method1() int64 {
	var maxNum int64

	for i := 999; i >= 100; i-- {
		for j := 990; j >= 100; j -= 11 {
			product := int64(i * j)
			if product > maxNum && isPalindrome(product) {
				maxNum = product
			}
		}
	}

	return maxNum
}

func main() {
	fmt.Println(method0())
	fmt.Println(method1())
}
