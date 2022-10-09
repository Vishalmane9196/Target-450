package main

import (
	"fmt"
	"math"
)

// Link: https://takeuforward.org/data-structure/implement-queue-using-array/
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

// Link: https://leetcode.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	s1    []int
	s2    []int
	front int
}

func Constructor() MyQueue {
	return MyQueue{[]int{}, []int{}, math.MinInt}
}

func (this *MyQueue) Push(x int) {
	//check s1 is empty or not
	if len(this.s1) == 0 {
		this.front = x
	}
	this.s1 = append(this.s1, x)
}

func (this *MyQueue) Pop() int {
	//check stack s2 is empty or not,
	//if empty then the push all element of s1 to s2 one by one
	if len(this.s2) == 0 {
		for len(this.s1) != 0 {
			var num int = this.s1[len(this.s1)-1]
			this.s1 = this.s1[:len(this.s1)-1]

			this.s2 = append(this.s2, num)
		}
	}

	var popElement = this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]

	for len(this.s2) > 0 {
		last := this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		this.Push(last)
	}
	return popElement
}

func (this *MyQueue) Peek() int {
	if len(this.s2) != 0 {
		return this.s2[len(this.s2)-1]
	}
	return this.front
}

func (this *MyQueue) Empty() bool {
	return len(this.s1) == 0 && len(this.s2) == 0
}

func queueUsingStack() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Println(" top : ", q.Peek())
	fmt.Println("st : ", q)
	fmt.Println(" top : ", q.Pop())
	fmt.Println(" isEmpty : ", q.Empty())
	fmt.Println("st : ", q)
}
