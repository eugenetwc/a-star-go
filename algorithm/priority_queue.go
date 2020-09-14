//https://golang.org/pkg/container/heap
package algorithm

type priorityQueue []*node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].f < pq[j].f
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node 
}