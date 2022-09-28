package main

import "fmt"

//Link: https://takeuforward.org/data-structure/implement-queue-using-array/
type Queue struct {
	arr   []int
	front int
	rear  int
}

func Init() *Queue {
	return &Queue{
		front: -1,
		rear:  -1,
	}
}

func (q *Queue) enqueue(data int) {

	if len(q.arr) == 0 {
		q.arr = append(q.arr, data)
		q.front++
		q.rear++
	} else {
		q.arr = append(q.arr, data)
		q.rear++
	}
}

func (q *Queue) dequeue() {

	if len(q.arr) == 0 {
		return
	}
	q.arr = q.arr[1:]
}

func (q *Queue) getSize() int {
	return len(q.arr)
}

func (q *Queue) getFront() int {
	return q.front
}

func (q *Queue) getRear() int {
	return q.rear
}

func (q *Queue) printQueue() {
	fmt.Printf("[Front]<-> ")
	for _, v := range q.arr {
		fmt.Printf(" [%v]<->", v)
	}
	fmt.Printf(" [Rear]\n")
}

//Link: https://takeuforward.org/data-structure/implement-stack-using-single-queue/
