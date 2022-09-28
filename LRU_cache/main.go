package main

import "fmt"

func main() {

	obj := Constructor(2)
	obj.Put(1, 10)
	obj.Put(2, 20)
	fmt.Println(obj.Get(1))
	obj.Put(3, 30)
	fmt.Println(obj.Get(2))
	obj.Put(4, 40)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))

}
