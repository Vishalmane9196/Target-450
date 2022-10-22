package main

import "fmt"

func main() {

	//heap implementation
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 45, 87, 78, 12, 55}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

	//Problems
	// words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	// k := 2
	// res := topKFrequent(words, k)
	// fmt.Println("top k frequent words  : ", res)
}
