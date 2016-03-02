package main

import (
	"testing"
)

var (
	nums []int
)

func init() {
	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}
}

func TestSumOfSquare(t *testing.T) {
	expected := 385
	res := sumOfSquare(nums)

	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

func BenchmarkSumOfSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumOfSquare(nums)
	}
}

func TestSquareOfSum(t *testing.T) {
	expected := 3025
	res := squareOfSum(nums)

	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

func BenchmarkSquareOfSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		squareOfSum(nums)
	}
}

func TestMethod0(t *testing.T) {
	expected := 2640
	res := method0(nums)

	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

func BenchmarkMethod0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method0(nums)
	}
}
