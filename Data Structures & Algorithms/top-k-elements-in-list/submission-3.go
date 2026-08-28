
type Item struct {
	val int
	freq int
}
type MaxHeap []Item

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) { *h = append(*h, x.(Item)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n - 1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) == k { return nums }

	count := map[int]int{}
	for _, val := range nums {
		count[val]++
	}

	h := &MaxHeap{}
	heap.Init(h)
	for num, freq := range count {
		heap.Push(h, Item{val: num, freq: freq})
	}

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(h).(Item).val
	}

	return result
}
