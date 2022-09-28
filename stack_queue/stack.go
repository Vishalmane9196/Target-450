// package main

// import "fmt"

// type stack struct {
// 	item []int
// 	top  int
// }

// func push(arr *stack, data int) {
// 	arr.item[arr.top] = data
// }

// func show_stack(arr *stack) {
// 	fmt.Printf("[BEG] <-> ")
// 	for i := 0; i < len(arr.item); i++ {
// 		fmt.Printf("[%v] <-> ", arr.item[i])
// 	}
// 	fmt.Printf("END")
// }

// func main() {

// 	var newstack = &stack{}

// 	push(newstack, 1)
// 	push(newstack, 2)
// 	push(newstack, 3)
// 	push(newstack, 4)
// 	push(newstack, 5)
// 	show_stack(newstack)

// }