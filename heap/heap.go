package main

import "container/heap"

//Link:https://leetcode.com/problems/top-k-frequent-words/
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
