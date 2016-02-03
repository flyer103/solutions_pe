package main

import (
	"testing"
)

var (
	expectedAnswer = 4613732
)

func TestMethod_0(t *testing.T) {
	data := method_0()
	if data != expectedAnswer {
		t.Errorf("expected %d, but got %d", expectedAnswer, data)
	}
}

func BenchmarkMethod_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		method_0()
	}
}

func TestMethod_1(t *testing.T) {
	data := method_1()
	if data != expectedAnswer {
		t.Errorf("expected %d, but got %d", expectedAnswer, data)
	}
}

func BenchmarkMethod_1(b *testing.B) {
	for i := 1; i < b.N; i++ {
		method_1()
	}
}
