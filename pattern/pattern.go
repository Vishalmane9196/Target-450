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

/*
---------------
*
**
***
****
*****
****
***
**
*
---------------
*/
func p9(n int) {
	for i := 0; i < 2*n; i++ {

		if i < n {
			for j := 0; j <=i; j++ {
				fmt.Printf("*")
			}
		}else{
			for j := 0; j <2*n-1-i; j++ {
				fmt.Printf("*")
			}
		}
		fmt.Println()
	}
}

/*
---------------
    *
   * *
  * * *
 * * * *
* * * * *
 * * * *
  * * *
   * *
    *
---------------
*/
func p10(n int) {

	for row := 0; row < 2*n; row++ {

		var totalColInRow int = 0
		if row > n {
			totalColInRow = 2*n-row
		}else{
			totalColInRow = row
		}

		var numberOfSpaces = n-totalColInRow
		// fmt.Println("s : ",s)
		// fmt.Println("k : ",k)
		for j:= 0; j<numberOfSpaces; j++ {
			fmt.Printf(" ")
		}
		for col:=0; col<totalColInRow; col++ {
			fmt.Printf("* ")
		}

		fmt.Println()
	}
}

/*
1      1
12    12
123  123
12341234
*/
func p11(n int){

	for row:=1; row<=n; row++ {
		for col:=1;col<=row; col++ {
			fmt.Printf("%v",col)
		}
		var noSpaces = 2*n-2*row
		for col:=1;col<=noSpaces; col++ {
			fmt.Printf(" ")
		}
		for col:=1;col<=row; col++ {
			fmt.Printf("%v",col)
		}
		fmt.Println()
	}
}

/*
-------------
****
*  *
*  *
*  *
****
---------------
*/
func p12(row, col int) {

	for i:=1; i<=row; i++ {
		for j:=1; j<=col; j++ {

			if i==1 || i==row || j==1 || j==col{
        fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
			// fmt.Printf("%v ==> %v\n",i,j)
		}
		fmt.Println()
	}
}

/*
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/
func p13(row int) {
 var count = 1
	for i:=1; i<=row; i++ {
		for j:=1; j<=i; j++ {
      fmt.Printf("%v ",count)
			count++
		}
		fmt.Println()
	}
}

/*
-----------
*        *
**      **
***    ***
****  ****
**********
**********
****  ****
***    ***
**      **
*        *
------------
*/
func p14(row int) {

	for i:=1; i<=row; i++ {
		for j:=1; j<=i; j++ {
      fmt.Printf("*")
		}
		var noSpaces = 2*row-2*i
		for j:=1; j<=noSpaces; j++ {
      fmt.Printf(" ")
		}
		for j:=1; j<=i; j++ {
      fmt.Printf("*")
		}
		fmt.Println()
	}

	for i:=row; i>=1; i-- {
		for j:=1; j<=i; j++ {
      fmt.Printf("*")
		}
		var noSpaces = 2*row-2*i
		for j:=1; j<=noSpaces; j++ {
      fmt.Printf(" ")
		}
		for j:=1; j<=i; j++ {
      fmt.Printf("*")
		}
		fmt.Println()
	}
}

/*
1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
*/
func p15(row int) {
	for i:=1; i<=row; i++ {
		for j:=1; j<=i; j++ {
			if (i+j)%2!=0{
				fmt.Printf("0")
			}else{
				fmt.Printf("1")
			}
		}
		fmt.Println()
	}
}

/*
A
A B
A B C
A B C D
A B C D E
*/
func p16() {
	for i:='A'; i<='E'; i++ {
		for j:='A'; j<=i; j++ {
				fmt.Printf("%v ",string(j))
		}
		fmt.Println()
	}
}

/*
A B C D E
A B C D
A B C
A B
A
*/
func p17() {
	for i:='E'; i>='A'; i-- {
		for j:='A'; j<=i; j++ {
				fmt.Printf("%v ",string(j))
		}
		fmt.Println()
	}
}

/*
A
B B
C C C
D D D D
E E E E E
*/
func p18() {
	for i:='A'; i<='E'; i++ {
		for j:='A'; j<=i; j++ {
				fmt.Printf("%v ",string(i))
		}
		fmt.Println()
	}
}

/*
E
D E
C D E
B C D E
A B C D E
*/
func p19() {
	for i:='E'; i>='A'; i-- {
		for j:=i; j<='E'; j++ {
				fmt.Printf("%v ",string(j))
		}
		fmt.Println()
	}
}

/*
E
D E
C D E
B C D E
A B C D E
*/	
func p20() {
	for i:='A'; i<='D'; i++ {
		var noSpaces = 'D'-i
		for j:='A'; j<=noSpaces; j++ {
				fmt.Printf(" ")
		}
		for j:='A'; j<=i; j++ {
				fmt.Printf("%v ",string(j))
		}
		for j:=i+1; j<='A'; j-- {
				fmt.Printf("%v ",string(j))
		}
		fmt.Println()
	}
}