package main

import "testing"

func BenchmarkDenominationChange(b *testing.B) {
	for n := 0; n < b.N; n++ {
		denominationChange(1972.78)
	}
}
