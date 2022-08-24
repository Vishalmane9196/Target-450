package main

import "fmt"

func enqueue(q queue, data int) {
	q.arr = append(q.arr, data)
	if len(q.arr) < 0 {
		q.rear = 
	}
}

func dequeue_(q queue) { 

	q.arr = q.arr[1:]
	if len(q.arr) < 0 {
		q.front = -1
	}else{
		q.front = q.arr[0]
	}
}

func front(q queue) int {
	return q.arr[q.front]
}

func print(q queue) {

	for i := 0; i < len(q.arr); i++ {

	}

}

func main() {

	//dequeue implmentation 
	var dq = &dequeue{[]int{}, -1, -1}
	addFirst(dq, 1)
	addFirst(dq, 2)
	addFirst(dq, 3)
	addLast(dq, 4)
	addLast(dq, 5)
	addLast(dq, 6)
	printDequeue(dq)
	firstElement := getFirst(dq)
	fmt.Println("First element of dequeue : ", firstElement)
	lastElement := getLast(dq)
	fmt.Println("Last element of dequeue : ", lastElement)
  felement := removeFirst(dq)
	fmt.Println("removed the front elment from dequeue : ", felement)
	lelement := removeLast(dq)
	fmt.Println("removed the last elment from dequeue : ", lelement)
	printDequeue(dq)
}