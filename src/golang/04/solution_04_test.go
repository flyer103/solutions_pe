package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(int64(9009)) {
		t.Errorf("9009 is palindromic number, but got false")
	}

	if isPalindrome(int64(9008)) {
		t.Errorf("9008 is not palindromic number, but got true")
	}
}

func TestMethod_0(t *testing.T) {
	data := method0()
	expected := int64(906609)
	if data != expected {
		t.Errorf("Expected %d, but got %d", expected, data)
	}
}

func BenchmarkMethod_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method0()
	}
}
