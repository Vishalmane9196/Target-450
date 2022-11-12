package main

func main() {

	/********Max Heap implementation*******/
	// m := &MaxHeap{}
	// buildHeap := []int{10, 20, 40, 80, 100, 70}
	// for _, v := range buildHeap {
	// 	m.Insert(v)
	// }
	// m.PrintHeapTree()
	// m.ExtractMax()
	// m.PrintHeapTree()
	// m.DecreaseKey(150, 3)
	// m.PrintHeapTree()
	// m.delete(40)
	// m.PrintHeapTree()
	// fmt.Println("checkMaxHeap : ", m.checkMaxHeap())

	/********Min Heap implementation*******/
	//m := &MinHeap{}
	//buildHeap := []int{10, 20, 30, 5, 45, 87, 78, 12, 55}
	//for _, v := range buildHeap {
	//	m.Insert(v)
	//}
	//m.PrintHeapTree()
	//m.ExtractMin()
	//m.PrintHeapTree()
	//m.DecreaseKey(150, 3)
	//m.PrintHeapTree()
	//m.delete(3)
	//m.PrintHeapTree()
	//fmt.Println("checkMinHeap : ", m.checkMinHeap())

	/********Min Heap implementation*******/
	// m := &MinHeap{}
	// buildHeap := []int{10, 20, 30, 5, 45, 87, 78, 12, 55}
	// for _, v := range buildHeap {
	// 	m.Insert(v)
	// }
	// m.PrintHeapTree()
	// m.Extract()
	// m.PrintHeapTree()
	// for i := 0; i < 5; i++ {
	// 	// fmt.Println(m)
	// }

	/*********Problems*********/
	// words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	// k := 2
	// res := topKFrequentWord(words, k)
	// fmt.Println("top k frequent words  : ", res)

	// minHeap := []int{3, 5, 9, 6, 8, 20, 10, 12, 18, 9}
	// res := convertMaxHeap(minHeap)
	// fmt.Println("convertMaxHeap  : ", res)

	// arr := []int{99, 5, 9, 6, 8, 20, 10, 12, 18, 9}
	// // arr := []int{6, 5, 3, 7, 2, 8, -1}
	// res := heapSort(arr, len(arr))
	// fmt.Println("heap sort : ", res)

	// nums := []int{1, 1, 1, 2, 2, 3}
	// k := 2
	// res := topKFrequent(nums, k)
	// fmt.Println("top k frequent nums  : ", res)

	// arr := []int{3, 2, 1, 5, 6, 4}
	// k := 2
	// arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	// k := 4
	// res := findKthLargest(arr, k)
	// fmt.Println("findKthLargest : ", res)

	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	obj.FindMedian()
	obj.AddNum(3)
	obj.FindMedian()

}
