package main

import (
	"testing"
)
// BenchmarkAStar is for benchmarking
func BenchmarkAStar(b *testing.B) {
	coordMap := [][]int{
		{0, 0, 1, 0, 3, 0},
		{0, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 1, 0},
		{2, 0, 0, 0, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		c := AStar(coordMap)
		for range c {
		}
	}
}