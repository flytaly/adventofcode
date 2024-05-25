package utils

type PItem[K comparable] struct {
	Value    K
	Priority int
	index    int
}

type PQ[K comparable] []*PItem[K]

func (pq PQ[_]) Len() int {
	return len(pq)
}

func (pq PQ[_]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PQ[_]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ[K]) Push(x any) {
	n := len(*pq)
	item := x.(*PItem[K])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PQ[_]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
