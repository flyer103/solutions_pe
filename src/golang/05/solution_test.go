package main

import (
	"reflect"
	"testing"
)

func TestIsPrime(t *testing.T) {
	if isPrime(1) {
		t.Error("1 is not prime")
	}

	if !isPrime(2) {
		t.Error("2 is prime")
	}

	if !isPrime(3) {
		t.Error("3 is prime")
	}

	if isPrime(1000) {
		t.Error("1000 is prime")
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(1000)
	}
}

func TestGetAllPrimes(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19}
	data := getAllPrimes(1, 20)

	if !reflect.DeepEqual(data, expected) {
		t.Error("not equal")
	}
}

func BenchmarkGetAllPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getAllPrimes(1, 20)
	}
}

func TestGetMaxMultiple(t *testing.T) {
	expected := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	data := getMaxMultiple(1, 20)

	if !reflect.DeepEqual(data, expected) {
		t.Error("not equal")
	}
}

func BenchmarkGetMaxMultiple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMaxMultiple(1, 20)
	}
}

func TestGetSmallestNumber(t *testing.T) {
	const expected = 232792560
	data := getSmallestNumber(1, 20)
	if data != expected {
		t.Errorf("expected %d, but got %d\n", expected, data)
	}
}

func BenchmarkGetSmallestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getSmallestNumber(1, 20)
	}
}

func TestGcd(t *testing.T) {
	const expected = 4
	data := gcd(12, 8)
	if data != expected {
		t.Errorf("expected %d, but got %d\n", expected, data)
	}
}

func BenchmarkGcd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gcd(12, 8)
	}
}

func TestLcm(t *testing.T) {
	const expected = 12
	data := lcm(4, 6)
	if data != expected {
		t.Errorf("expected %d, but got %d\n", expected, data)
	}
}

func BenchmarkLcm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcm(4, 6)
	}
}

func TestLcmm(t *testing.T) {
	const expected = 232792560
	data := lcmm(1, 20)
	if data != expected {
		t.Errorf("expected %d, but got %d\n", expected, data)
	}
}

func BenchmarkLcmm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lcmm(1, 20)
	}
}
