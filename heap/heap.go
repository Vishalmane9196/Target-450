package main

import (
	"container/heap"
	"fmt"
	"math"
)

/********Max Heap implementation*******/
// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap and calls maxHeapifyUp
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// ExtractMax returns the largest key, and removes it from heap
func (h *MaxHeap) ExtractMax() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp find the correct position for given element index
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
	largest := 0
	// loop while index has at least one child
	for l <= lastIndex { // when left child is the only child
		//finding the lowest value of index, it's left and right
		if l == lastIndex {
			largest = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			largest = l
		} else { // when right child is larger
			largest = r
		}
		//after finding the lowest value, compare and swap them
		// Compare array value of the current index to larger child and swap if smaller
		if h.array[index] < h.array[largest] {
			h.swap(index, largest)
			index = largest
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MaxHeap) DecreaseKey(val, index int) {

	h.array[index] = val
	for index != 0 && h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func (h *MaxHeap) delete(index int) {

	h.DecreaseKey(math.MaxInt, 3)
	h.ExtractMax()
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

// PrintHeapTree print binary tree
func (h MaxHeap) PrintHeapTree() {
	fmt.Printf("[START] <->")
	for i := range h.array {
		fmt.Printf("[%v] <-> ", h.array[i])
	}
	fmt.Printf("[END]\n")
}

/********Min Heap implementation*******/
// MinHeap struct has a slice that holds the array
type MinHeap struct {
	array []int
}

// Insert adds an element to the heap and calls maxHeapifyUp
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.minHeapifyUp(len(h.array) - 1)
}

// ExtractMin returns the smallest key, and removes it from heap
func (h *MinHeap) ExtractMin() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[1:]

	h.minHeapifyDown(0)
	return extracted
}

// minHeapifyUp find the correct position for given element index
func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// minHeapifyDown process
func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	smallest := 0
	// loop while index has at least one child
	for l <= lastIndex { // when left child is the only child
		//finding the lowest value of index, it's left and right
		if l == lastIndex {
			smallest = l
		} else if h.array[l] < h.array[r] { // when left child is larger
			smallest = l
		} else { // when right child is larger
			smallest = r
		}
		//after finding the lowest value, compare and swap them
		// Compare array value of the current index to larger child and swap if smaller
		if h.array[index] > h.array[smallest] {
			h.swap(index, smallest)
			index = smallest
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) DecreaseKey(val, index int) {
	h.array[index] = val
	for index != 0 && h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func (h *MinHeap) delete(index int) {
	h.DecreaseKey(math.MinInt, 3)
	h.ExtractMin()
}

// Swap keys in the array
func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// PrintHeapTree print binary tree
func (h MinHeap) PrintHeapTree() {
	fmt.Printf("[START] <->")
	for i := range h.array {
		fmt.Printf("[%v] <-> ", h.array[i])
	}
	fmt.Printf("[END]\n")
}

/*******Problems********/
// Link: https://helloacm.com/check-if-an-array-represents-a-max-heap-danny-heap-algorithm/
// Link: https://www.techiedelight.com/check-given-array-represents-min-heap-not/
func (h MaxHeap) checkMaxHeap() bool {
	nums := h.array
	n := len(nums)
	for i := 0; i < n; i++ {
		if (2*i+1 < n) && (nums[i] < nums[2*i+1]) {
			return false
		}
		if (2*i+2 < n) && (nums[i] < nums[2*i+2]) {
			return false
		}
	}
	return true
}

func (h MinHeap) checkMinHeap() bool {
	nums := h.array
	n := len(nums)
	for i := 0; i < n; i++ {
		if (2*i+1 < n) && (nums[i] > nums[2*i+1]) {
			return false
		}
		if (2*i+2 < n) && (nums[i] > nums[2*i+2]) {
			return false
		}
	}
	return true
}

// Link: https://afteracademy.com/blog/convert-a-min-heap-to-a-max-heap
func convertMaxHeap(arr []int) []int {
	maxheap := &MaxHeap{}
	for i := 0; i < len(arr); i++ {
		maxheap.Insert(arr[i])
	}
	return maxheap.array
}

// heap sort
func heapSort(arr []int, n int) []int {
	fmt.Println("original array : ", arr)
	//build heap
	m := &MaxHeap{}
	for _, v := range arr {
		m.Insert(v)
	}
	// fmt.Println("maxheap : ", m)
	//sorting
	for i := n - 1; i >= 0; i-- {
		arr[i] = m.array[0]
		m.array = m.array[1:]
		m.maxHeapifyDown(0)
	}
	// fmt.Println("arr : ", arr)
	return arr
}

// Link: https://leetcode.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	res := heapSort(nums, len(nums))
	fmt.Println("res : ", res)
	return res[k-1]
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
	length := len(*q)
	elem := (*q)[length-1]
	*q = (*q)[:length-1]
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

func topKFrequentWord(words []string, k int) []string {

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

// Link: https://leetcode.com/problems/top-k-frequent-elements/
type numsFreq struct {
	number    int
	frequency int
}

type npq []numsFreq

func (q npq) Len() int {
	return len(q)
}

func (q *npq) Push(a interface{}) {
	*q = append(*q, a.(numsFreq))
}

func (q *npq) Pop() interface{} {
	length := len(*q)
	elem := (*q)[length-1]
	*q = (*q)[:length-1]
	return elem
}

func (q npq) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q npq) Less(i int, j int) bool {
	if q[i].frequency == q[j].frequency {
		return q[i].number < q[j].number
	}
	return q[i].frequency > q[j].frequency
}

func topKFrequent(nums []int, k int) []int {

	ans := make([]int, k)
	freqMap := make(map[int]int)
	pqObj := new(npq)

	for _, num := range nums {
		if _, ok := freqMap[num]; ok {
			freqMap[num]++
		} else {
			freqMap[num] = 1
		}
	}

	for i, v := range freqMap {
		heap.Push(pqObj, numsFreq{number: i, frequency: v})
	}

	for i := 0; i < k; i++ {
		ans[i] = heap.Pop(pqObj).(numsFreq).number
	}
	return ans
}

// Link: https://leetcode.com/problems/find-median-from-data-stream/
// type MedianFinder struct {
// 	maxHeap *MaxHeap
// 	minHeap *MinHeap
// }

// // max heap + min heap
// func Constructor() MedianFinder {
// 	maxHeap := make(MaxHeap, 0)
// 	heap.Init(&maxHeap)

// 	minHeap := make(MinHeap, 0)
// 	heap.Init(&minHeap)

// 	return MedianFinder{
// 		maxHeap: &maxHeap,
// 		minHeap: &minHeap,
// 	}
// }

// func (this *MedianFinder) AddNum(num int) {
// 	if len(*this.maxHeap) == 0 || num > (*this.maxHeap)[0] {
// 		heap.Push(this.minHeap, num)
// 	} else {
// 		heap.Push(this.maxHeap, num)
// 	}

// 	if len(*this.maxHeap) > len(*this.minHeap) && len(*this.maxHeap)-len(*this.minHeap) >= 2 {
// 		v := heap.Pop(this.maxHeap)
// 		heap.Push(this.minHeap, v)
// 	}

// 	if len(*this.minHeap) > len(*this.maxHeap) && len(*this.minHeap)-len(*this.maxHeap) >= 2 {
// 		v := heap.Pop(this.minHeap)
// 		heap.Push(this.maxHeap, v)
// 	}
// }

// func (this *MedianFinder) FindMedian() float64 {
// 	if (len(*this.maxHeap)+len(*this.minHeap))%2 == 1 {
// 		if len(*this.maxHeap) > len(*this.minHeap) {
// 			return float64((*this.maxHeap)[0])
// 		} else {
// 			return float64((*this.minHeap)[0])
// 		}
// 	} else {
// 		small := (*this.maxHeap)[0]
// 		large := (*this.minHeap)[0]
// 		return (float64(small) + float64(large)) / 2
// 	}

// }

// type MinHeap []int

// func (h MinHeap) Len() int           { return len(h) }
// func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MinHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// type MaxHeap []int

// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
// func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *MaxHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MaxHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
