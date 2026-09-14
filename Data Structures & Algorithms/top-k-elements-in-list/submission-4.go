type E struct {
	val, count int
}

type PriorityQueue []E

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(E))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	var res []int
	pq := &PriorityQueue{}

	h := make(map[int]int, len(nums))

	for _, val := range nums {
		h[val]++
	}

	for val, count := range h {
		heap.Push(pq, E{val, count} )
	}

	for pq.Len() > k {
		heap.Pop(pq)
	}

	for pq.Len() > 0 {
		pop := heap.Pop(pq).(E)

		res = append(res, pop.val)
	}

	return res
}
