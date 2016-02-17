// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
// Answer: 232792560
package main

import (
	"fmt"
	"math"
	"sort"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func getAllPrimes(start, end int) (data []int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			data = append(data, i)
		}
	}

	return data
}

func getMaxMultiple(start, end int) (data []int) {
	numSet := make(map[int]struct{})

	for i := start; i <= end; i++ {
		maxNum := i
		for j := 2; j <= end; j++ {
			product := i * j
			if product > maxNum && product <= end {
				maxNum = product
			}
		}
		numSet[maxNum] = struct{}{}
	}

	for num, _ := range numSet {
		data = append(data, num)
	}
	sort.Ints(data)

	return data
}

func getSmallestNumber(start, end int) int {
	allPrimes := getAllPrimes(start, end)
	multiplers := getMaxMultiple(start, end)

	minProduct := 1
	for _, i := range allPrimes {
		minProduct *= i
	}

	multipler := 1
	for {
	LOOP:
		product := minProduct * multipler
		for _, i := range multiplers {
			if product%i != 0 {
				multipler++
				goto LOOP
			}
		}

		return product
	}
}

func main() {
	fmt.Println(getSmallestNumber(1, 20))
}
