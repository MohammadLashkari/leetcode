package main

import (
	"container/heap"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	var (
		freq   = make(map[byte]int)
		h      = &MaxHeap{}
		answer = 0
	)
	heap.Init(h)

	for _, task := range tasks {
		freq[task]++
	}
	for _, f := range freq {
		heap.Push(h, f)
	}

	for h.Len() != 0 {
		next_round := []int{}
		fleet := n + 1

		for fleet > 0 && h.Len() != 0 {
			task_freq := h.Pop().(int)
			fleet--
			answer++
			task_freq--
			if task_freq > 0 {
				next_round = append(next_round, task_freq)
			}
		}

		for _, task := range next_round {
			h.Push(task)
		}

		if h.Len() == 0 {
			break
		}
		answer += fleet
	}
	return answer
}

func main() {

	leastInterval([]byte{'a', 'a', 'b', 'a', 'a'}, 2)

}
