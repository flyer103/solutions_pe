// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
// Answer: 906609

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var (
	start = 100
	end   = 999
)

type int64Arr []int64

func (a int64Arr) Len() int {
	return len(a)
}

func (a int64Arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a int64Arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

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
	var numbers int64Arr
	var total int
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			numbers = append(numbers, int64(i*j))
			total += 1
		}
	}
	sort.Sort(numbers)

	for total--; total >= 0; total-- {
		if isPalindrome(numbers[total]) {
			return numbers[total]
		}
	}

	return int64(-1)
}

func main() {
	fmt.Println(method0())
}
