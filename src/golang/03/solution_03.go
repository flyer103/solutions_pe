// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143?
// Answer: 6857

package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	number = flag.Int64("number", 600851475143, "input number")
)

func init() {
	flag.Parse()
}

func isPrime(num int64) bool {
	if num <= 1 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}

	sqrtValue := int64(math.Sqrt(float64(num)))
	for i := int64(4); i <= sqrtValue; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func method0(number int64) (res int64) {
	sqrtValue := int64(math.Sqrt(float64(number)))
	for i := sqrtValue; i >= 2; i-- {
		if number%i == 0 && isPrime(i) {
			return i
		}
	}

	return -1
}

func method1(number int64) (res int64) {
	for i := int64(2); i <= number/2; i++ {
		if number%i == 0 {
			number = number / i
			if isPrime(number) {
				return number
			}
			method1(number)
		}
	}

	return -1
}

func main() {
	fmt.Println(method0(*number))
	fmt.Println(method1(*number))
}
