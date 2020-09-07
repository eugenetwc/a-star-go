package astar

import (
	"C"
	"container/heap"
	"math"
)

// Location is position
type Location struct {
	x int
	y int
}

// Node is potential position within grid
type node struct {
	parent *node
	position Location
	f, g, h int
	index int
}

// AStar function
func AStar(coordMap [][]int) []Location {
	var startPos, goalPos Location
	// find start and end points
	for y, row := range coordMap {
		for x, val := range row {
			if  val == 2 { // start coord where val == 2
				startPos = Location{x, y}
			} else if val == 3 { // end coord where val == 3
				goalPos = Location{x, y}
			}
			if (startPos != Location{} && goalPos != Location{}) {
				break
			}
		}
	}

	startNode := node{parent: nil, position: startPos}
	// Init priorityQueue open list with zero capacity
	var open priorityQueue
	heap.Init(&open)
	var closed []*node
	// Push startNode into open list
	heap.Push(&open, &startNode)	

	for open.Len() > 0 {
		currNode := heap.Pop(&open).(*node)
		closed = append(closed, currNode)

		if currNode.position == goalPos {
			var path []Location
			current := currNode
			for current != nil {
				path = append([]Location{current.position}, path...)
				current = current.parent
			}
			return path
		}

		successors := getNeighbours(currNode, coordMap)

		for _, childPos := range successors {
			tentativeGScore := currNode.g + 1

			openVisited := false
			for _, openNode := range open {
				if openNode.position == childPos {
					if openNode.g <= tentativeGScore {
						openVisited = true
						break
					} else {
						heap.Remove(&open, openNode.index)
						break
					}
				}
			}
			if openVisited == true {continue}

			closedVisited := false
			closedIndex := 0
			for i, visitedNode := range closed {
				if visitedNode.position == childPos {
					if visitedNode.g <= tentativeGScore {
						closedVisited = true
						break
					} else {
						closedIndex = i	
						break
					}
				}
			}
			if closedVisited == true {continue}
			if closedIndex != 0 {
				closed = append(closed[closedIndex:], closed[:closedIndex+1]...)
			}

			newNode := node{parent: currNode, position: childPos}
			newNode.g = tentativeGScore
			newNode.h = manhattan(childPos, goalPos)
			newNode.f = newNode.g + newNode.h

			heap.Push(&open, &newNode)
		}
	}
	return nil
}

func getNeighbours(currNode *node, coordMap [][]int) []Location {
	mapSize := [2]int{len(coordMap[0]), len(coordMap)}
	x, y := currNode.position.x, currNode.position.y
	successorPos := []Location{
		Location{x: x, y: y+1},
		Location{x: x, y: y-1},
		Location{x: x+1, y: y},
		Location{x: x-1, y: y},
	}

	validPos := successorPos[:0]
	for _, pos := range successorPos {
		if (pos.x < mapSize[0] && pos.x >= 0 && pos.y < mapSize[1] && pos.y >= 0) {
			if coordMap[pos.y][pos.x] != 1 {
				validPos = append(validPos, pos)
			}
		}
	}
	return validPos
}

func manhattan (a, b Location) int {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return int(math.Abs(dx) + math.Abs(dy))
}