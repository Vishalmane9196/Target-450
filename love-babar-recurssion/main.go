package main

func main() {

	//basic questions

	// res := fibo(4)
	// fmt.Println("fibo ", res)

	// printDecreasingOrder(4)

	// res := reverseString("vishal")
	// fmt.Println(" res  : ", res)

}

/*

Recurssion -- function calls iteself
   solve one case and rest will be take care b y recurssion

problem's --> solution depend on solution of subproblem

There are baically 3 mains steps we need to figure out in problem

1. Base case               --- mandatory
2. Operation/processing    --- optional
3. Recurssive relation     --- mandatory

callstack is used to maintain the history of recurssion and keep the track of function calls

Types of recursion
1.Head recurssion
 In head recursion, the recursive call, when it happens, comes before other processing in the function (think of it happening at the top, or head, of the function).
 If the recursive call occurs at the beginning of a method, it is called a head recursion. The method saves the state before jumping into the next recursive call.

2.Tail recurssion
 In tail recursion, it’s the opposite—the processing occurs before the recursive call. Choosing between the two recursive styles may seem arbitrary, but the choice can make all the difference.
 If the recursive call occurs at the end of a method, it is called a tail recursion. The tail recursion is similar to a loop. The method executes all the statements before jumping into the next recursive call.


A function with a path with a single recursive call at the beginning of the path uses what is called head recursion. The factorial function of a previous exhibit uses head recursion. The first thing it does once it determines that recursion is needed is to call itself with the decremented parameter. A function with a single recursive call at the end of a path is using tail recursion. Refer this article
Example Recursion :

public void tail(int n)         |     public void head(int n)
{                               |     {
	//base case                 |        //base case
    if(n == 1)                  |         if(n == 0)
        return;                 |             return;
    else                        |         else
	//processing                |         //recurssive relation
        System.out.println(n);  |             head(n-1);
                                |
	//recurssive relation       |         //processing
    tail(n-1);                  |         System.out.println(n);
}                               |     }

*/
