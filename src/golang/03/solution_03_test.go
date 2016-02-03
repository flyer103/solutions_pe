package main

import (
	"testing"
)

var (
	testData       = int64(600851475143)
	expectedAnswer = int64(6857)
)

func TestMethod_0(t *testing.T) {
	data := method0(testData)
	if data != expectedAnswer {
		t.Errorf("expected %d, but got %d", expectedAnswer, data)
	}
}

func BenchmarkMethod_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method0(testData)
	}
}

func TestMethod_1(t *testing.T) {
	data := method1(testData)
	if data != expectedAnswer {
		t.Errorf("expected %d, but got %d", expectedAnswer, data)
	}
}

func BenchmarkMethod_1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		method1(testData)
	}
}
