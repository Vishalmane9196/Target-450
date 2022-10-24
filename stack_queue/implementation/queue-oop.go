package main

import "fmt"

// Link: https://takeuforward.org/data-structure/implement-queue-using-array/
type Queue struct {
	arr   []int
	front int //store front index
	rear  int //store rear index
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
		q.front = -1
		q.rear = -1
		return
	} else if len(q.arr) == 1 {
		q.front = -1
		q.rear = -1
		q.arr = q.arr[1:]
	} else {
		q.arr = q.arr[1:]
		q.rear--
	}
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
	fmt.Printf("[FRONT]<-> ")
	for _, v := range q.arr {
		fmt.Printf(" [%v]<->", v)
	}
	fmt.Printf(" [REAR]\n")
}

func main() {

	/*******Queue implementation*****/
	st := Init()
	st.enqueue(1)
	st.enqueue(2)
	st.enqueue(3)
	st.enqueue(4)
	st.enqueue(5)
	st.printQueue()
	fmt.Println(" Front : ", st.getFront())
	fmt.Println(" Rear  : ", st.getRear())
	fmt.Println(" size of stack : ", st.getSize())
	st.dequeue()
	st.dequeue()
	st.dequeue()
	st.dequeue()
	st.dequeue()
	fmt.Println(" Front : ", st.getFront())
	fmt.Println(" Rear  : ", st.getRear())
	fmt.Println(" size of stack : ", st.getSize())
	st.printQueue()
}
