package main

import "fmt"

func main() {

	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	k := 2
	res := topKFrequent(words, k)
	fmt.Println("top k frequent words  : ", res)
}
