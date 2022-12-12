package main

import (
	"math/rand"
	"strconv"
)

// Link: https://leetcode.com/problems/maximum-69-number/
func maximum69Number(num int) int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	used := false
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 6 && !used {
			nums[i] = 9
			used = true
		}
		ans = ans*10 + nums[i]
	}

	return ans
}

// Link: https://leetcode.com/problems/rectangle-area/
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	first := area(ax1, ay1, ax2, ay2)
	second := area(bx1, by1, bx2, by2)

	leftOverlap := max(ax1, bx1)
	downOverlap := max(ay1, by1)

	rightOverlap := min(ax2, bx2)
	upOverlap := min(ay2, by2)

	overlap := area(leftOverlap, downOverlap, rightOverlap, upOverlap)

	return first + second - overlap

}

// Returns 0 if right < left or up < down
func area(left, down, right, up int) int {
	if left > right || down > up {
		return 0
	}
	return (right - left) * (up - down)
}

// math.Max accepts and returns floats and makes a mess
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Link: https://leetcode.com/problems/ugly-number/
func isUgly(n int) bool {
	arr := [...]int{2, 3, 5}
	for _, v := range arr {
		for n >= v && n%v == 0 {
			n /= v
		}
	}
	return n == 1
}

// Link: https://leetcode.com/problems/valid-sudoku/
func isValidSudoku(board [][]byte) bool {
	rMap, cMap, boxMap := make(map[string]byte), make(map[string]byte), make(map[string]byte)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if "." == string(board[i][j]) {
				continue
			}

			_, row := rMap[strconv.Itoa(i)+string(board[i][j])]
			_, col := cMap[strconv.Itoa(j)+string(board[i][j])]
			_, box := boxMap[strconv.Itoa(i/3)+strconv.Itoa(j/3)+string(board[i][j])]

			if row || col || box {
				return false
			}

			rMap[strconv.Itoa(i)+string(board[i][j])] = 1
			cMap[strconv.Itoa(j)+string(board[i][j])] = 1
			boxMap[strconv.Itoa(i/3)+strconv.Itoa(j/3)+string(board[i][j])] = 1
		}
	}
	return true
}

// Link: https://leetcode.com/problems/insert-delete-getrandom-o1/
type RandomizedSet struct {
	elementMap  map[int]int //store the lement and its corrosponding index of lementlist
	elementlist []int       //store all the elments from list
}

func Constructor() RandomizedSet {
	elementMap := make(map[int]int)
	elementlist := make([]int, 0)
	return RandomizedSet{elementMap, elementlist}
}

func (this *RandomizedSet) Insert(val int) bool {
	//checked if val exist in map or not
	if _, ok := this.elementMap[val]; ok {
		return false
	}
	//if val is not present then add the value in map and list
	this.elementMap[val] = len(this.elementlist)
	this.elementlist = append(this.elementlist, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	//check if value present in map or not
	if _, ok := this.elementMap[val]; !ok {
		return false
	}
	//fetch the current length of list and index of val from map
	listLength, valIndex := len(this.elementlist), this.elementMap[val]
	//swap the both fetched values
	this.elementlist[valIndex], this.elementlist[listLength-1] = this.elementlist[listLength-1], this.elementlist[valIndex]
	//update the correct index for val
	this.elementMap[this.elementlist[valIndex]] = valIndex
	//remove the val from list
	this.elementlist = this.elementlist[:listLength-1]
	//remove the val from map
	delete(this.elementMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.elementlist))
	return this.elementlist[i]
}

// Link: https://leetcode.com/problems/minimum-average-difference/
func minimumAverageDifference(nums []int) int {
	n, total := len(nums), 0
	for _, num := range nums {
		total += num
	}
	min, prefix, idx := total, 0, 0
	for i := 0; i < n-1; i++ {
		prefix += nums[i]
		cur := prefix/(i+1) - (total-prefix)/(n-i-1)
		if cur < 0 {
			cur = -cur
		}
		if cur < min {
			min, idx = cur, i
		}
	}
	if total/n < min {
		return n - 1
	}
	return idx
}

// Link: https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {

	dp := make([]int, n+1)
	if n <= 2 {
		return n
	}
	dp[0], dp[1], dp[2] = 0, 1, 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
