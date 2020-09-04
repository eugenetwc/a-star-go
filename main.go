package main

import (
	"fmt"
)

func main() {
	coordMap := [][]int{
		{0, 0, 1, 0, 3, 0},
		{0, 0, 0, 1, 0, 0},
		{1, 1, 0, 1, 1, 0},
		{2, 0, 0, 0, 0, 0},
	}
	test := AStar(coordMap)
	fmt.Println(test)
}
