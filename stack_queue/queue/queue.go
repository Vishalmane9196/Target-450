package main

import (
	"fmt"
	"math"
)

// [Leetcode-232]
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
			lastIndex := len(this.s1) - 1
			num := this.s1[lastIndex]
			this.s1 = this.s1[:lastIndex]

			this.s2 = append(this.s2, num)
		}
	}

	//removing top element from s2 stack
	var popElement = this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]

	//appending all element from s2 to s1 one by one
	for len(this.s2) != 0 {
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

func (this *MyQueue) PrintQueue() {
	fmt.Printf("[FRONT] <-> ")
	for i := 0; i < len(this.s1); i++ {
		fmt.Printf("[%v] <-> ", this.s1[i])
	}
	fmt.Printf("[REAR]\n")
}

func queueUsingStack() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Println(" top : ", q.Peek())
	q.PrintQueue()
	fmt.Println(" top : ", q.Pop())
	fmt.Println(" isEmpty : ", q.Empty())
	q.PrintQueue()
}
