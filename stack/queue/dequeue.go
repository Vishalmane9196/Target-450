package main

import "fmt"

type dequeue struct {
	arr   []int
	front int
	rear  int
}

//addfirst/last
//getfirst/last
//removefirst/last

func addFirst(q *dequeue, val int) {
	if len(q.arr) == 0 {
		q.rear = val
	}
	// old_queue_copy := q.arr
	var new_queue []int
	new_queue = append(new_queue, val)
	for len(q.arr) > 0 {
		var num int = q.arr[0]
		q.arr = q.arr[1:]
		new_queue = append(new_queue, num)
	}
	q.arr = new_queue
	q.front = q.arr[0]
}

func addLast(q *dequeue, val int) {
	if len(q.arr) == 0 {
		q.front = val
	}

	q.arr = append(q.arr, val)
	q.rear = q.arr[len(q.arr)-1]

}

func getFirst (q *dequeue) int {

  if len(q.arr) == 0 {
		return -1
	}
	return q.front
}

func getLast (q *dequeue) int {

  if len(q.arr) == 0 {
		return -1
	} 
	return q.rear
}

func removeFirst(q *dequeue) int{
	if len(q.arr) == 0 {
		return -1
	}
  var num = q.arr[0]
	q.arr = q.arr[1:]
	q.front = q.arr[0]
	
	return num
}

func removeLast(q *dequeue) int{
	if len(q.arr) == 0 {
		return -1
	}
  var num = q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	q.rear = q.arr[len(q.arr)-1]
	
	return num
}

func printDequeue(q *dequeue) {

	fmt.Printf("[BEG] <-> ")
	for _,v:=range q.arr {
		fmt.Printf("[%d] <-> ", v)
	}
	fmt.Printf("[END]\n")
	fmt.Printf("Front : %v\n", q.front)
	fmt.Printf("Rear : %v\n", q.rear)
}