package main

import "fmt"

/*
---------------
*****
*****
*****
*****
*****
---------------
*/
func p1(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

/*
---------------
*
**
***
****
*****
---------------
*/
func p2(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

/*
---------------
1
12
123
1234
12345
---------------
*/
func p3(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("%v",j+1)
		}
		fmt.Println()
	}
}

/*
---------------
1
12
123
1234
12345
---------------
*/
func p4(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("%v",i+1)
		}
		fmt.Println()
	}
}

/*
---------------
12345
1234
123
12
1
---------------
*/
func p5(n int) {
	for i := n-1; i >=0; i-- {
		for j := 0; j <=i; j++ {
			fmt.Printf("%v",j+1)
		}
		fmt.Println()
	}
}

/*
---------------
____*
___***
__*****
_*******
*********
---------------
*/
func p6(n int) {

	for i:=0; i <n; i++ {
		//decide how many spaces and how many star
	var k int = 0

		for j := 0; j < n-i-1; j++ {
			fmt.Printf("_")
		}
		for k:=0; k != (i+1)*2-1; k++ {
			fmt.Printf("*")
		}
		if k == (i+1)*2-1 {
			break
		}
		fmt.Println()
	}
}

/*
---------------
*********
_*******
__*****
___***
____*
---------------
*/
func p7(n int) {

	for i:=n-1; i >=0; i-- {
		//decide how many spaces and how many star
		var k int = 0

		for j := 0; j < n-i-1; j++ {
			fmt.Printf("_")
		}
		for k:=0; k != (i+1)*2-1; k++ {
			fmt.Printf("*")
		}
		if k == (i+1)*2-1 {
			break
		}
		fmt.Println()
	}
}

/*
---------------
1
12
123
1234
12345
---------------
*/
func p8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j <=i; j++ {
			fmt.Printf("%v",i+1)
		}
		fmt.Println()
	}
}