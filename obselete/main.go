package main

import (
	"fmt"
	"time"
)

func main() {
	coordMap := [][]int{
		{0, 0, 1, 0, 3, 0},
		{0, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 1, 0},
		{2, 0, 0, 0, 0, 0},
	}

	var path []Location
	start := time.Now()
	for i := 0; i < 1000; i++ {
		path = AStar(coordMap)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("solution:", path)
	fmt.Println("time elapsed:", elapsed)
}
