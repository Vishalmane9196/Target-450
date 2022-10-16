package main

import "fmt"

//Lecture 1
//basic problem

//fibonacci porblem

func fibo(n int) int {

	//base case
	if n <= 1 {
		return n
	}

	//recurssive relation and processing
	ans := fibo(n-1) + fibo(n-2)

	return ans
}

func printDecreasingOrder(n int) {
	//base case
	if n < 1 {
		return
	}

	//processing
	fmt.Printf("%v ", n)
	//recurssive relation
	printDecreasingOrder(n - 1)
}

//string
func reverseString(str string) string {

	// strbyte := []byte(str)
	// reverseStringHelper(strbyte, 0, len(str)-1)
	// return string(strbyte)

	strbyte := []byte(str)
	var ans []byte
	reverseStringHelper1(strbyte, &ans, 0, len(str))
	return string(ans)
}

func reverseStringHelper(str []byte, i, j int) {
	if i >= j {
		return
	}

	if str[i] != str[j] {
		str[i], str[j] = str[j], str[i]
		fmt.Println("str : ", string(str))
	}
	reverseStringHelper(str, i+1, j-1)
}

func reverseStringHelper1(str []byte, ans *[]byte, currLength, totalLength int) {
	if currLength == totalLength {
		return
	}
	reverseStringHelper1(str, ans, currLength+1, totalLength)
	*ans = append(*ans, str[currLength])
	fmt.Println("ans  : ", ans)
}

//Linked list

/*
1. print the linked list

solution==>
   i)
     print the number until head become nil
	 func solve (node *Node) {
		 if head == nil { return }
		 fmt.Println(head.data)
        solve(node.Next)
	 }


2. print the kth element from last/end
      solution ==>
	  trverse till end of linked list and thn while returning do the operations
	  handle if list is empty and k is grater than length of linked list

	    func solve(node *Node, k int) {

			if head == nil { return }
			solve(node.Next, k)
             k = k-1
			if k == 0 {
				fmt.Println(node.data)
			}
		}
*/
