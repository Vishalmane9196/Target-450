package main

import (
	"container/heap"
	"fmt"
)

// heap implementation
// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array lenth is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp process
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown process
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex { // when left child is the only child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// Compare array value of the current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap((index), childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int {
	return 2*i + 1
}

// Get the right child index
func right(i int) int {
	return 2*i + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Link:https://leetcode.com/problems/top-k-frequent-words/
type wordFreq struct {
	word      string
	frequency int
}

type pq []wordFreq

func (q pq) Len() int {
	return len(q)
}

func (q *pq) Push(a interface{}) {
	*q = append(*q, a.(wordFreq))
}

func (q *pq) Pop() interface{} {
	len := len(*q)
	elem := (*q)[len-1]
	*q = (*q)[:len-1]
	return elem
}

func (q pq) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q pq) Less(i int, j int) bool {

	if q[i].frequency == q[j].frequency {
		return q[i].word < q[j].word
	}
	return q[i].frequency > q[j].frequency
}

func topKFrequent(words []string, k int) []string {

	ans := make([]string, k)
	freqMap := make(map[string]int)
	pqObj := new(pq)

	for _, word := range words {
		if _, ok := freqMap[word]; ok {
			freqMap[word]++
		} else {
			freqMap[word] = 1
		}
	}

	for i, v := range freqMap {
		heap.Push(pqObj, wordFreq{word: i, frequency: v})
	}

	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(pqObj).(wordFreq).word
	}

	return ans
}
