package main

import (
	"testing"
)

func TestCompute_0(t *testing.T) {
	numTarget := 1000
	resExpected := 233168
	resComputed := compute_0(numTarget)

	if resComputed != resExpected {
		t.Errorf("expected %d for %d, but get %d", resExpected, numTarget, resComputed)
	}
}

func BenchmarkCompute_0(b *testing.B) {
	numTarget := 1000

	for i := 0; i < b.N; i++ {
		compute_0(numTarget)
	}
}

func TestCompute_1(t *testing.T) {
	numTarget := 1000
	resExpected := 233168
	resComputed := compute_1(numTarget)

	if resComputed != resExpected {
		t.Errorf("expected %d for %d, but get %d", resExpected, numTarget, resComputed)
	}
}

func BenchmarkCompute_1(b *testing.B) {
	numTarget := 1000

	for i := 0; i < b.N; i++ {
		compute_1(numTarget)
	}
}
