package main

import (
	"fmt"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

//Binary search tree implementaion

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) insert(data int) {
	if tree.root == nil {
		tree.root = &Node{Val: data, Left: nil, Right: nil}
	} else {
		tree.root.insert(data)
	}
}

func (node *Node) insert(data int) {
	if node == nil {
		return
	} else if data < node.Val {
		if node.Left == nil {
			node.Left = &Node{Val: data, Left: nil, Right: nil}
		} else {
			node.Left.insert(data)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Val: data, Left: nil, Right: nil}
		} else {
			node.Right.insert(data)
		}
	}
}

func (tree BinarySearchTree) PrintBinarySearchTree() {

	str := tree.toTreeString("", true, "", tree.root)
	fmt.Println(str)
}

//------------------------------------------------------------------------------

func (tree BinarySearchTree) toTreeString(prefix string, top bool, str string, node *Node) string {

	Left := new(Node)
	Right := new(Node)

	if node == nil {
		return ""
	}

	Left = node.Left
	Right = node.Right

	if Right != nil {
		temp := tree.path(top, ""+prefix, "│   ", "    ")
		str = tree.toTreeString(temp, false, str, Right)
	}

	str = tree.path(top, str+prefix, "└──", "┌──")
	str = str + " " + strconv.Itoa(node.Val) + "\n"

	if Left != nil {
		temp := tree.path(top, ""+prefix, "    ", "│   ")
		str = tree.toTreeString(temp, true, str, Left)
	}
	return str
}

func (tree BinarySearchTree) path(condition bool, str string, choice1 string, choice2 string) string {

	if condition {
		str += choice1
	} else {
		str += choice2
	}
	return str
}

//--------------------------------------------------DSA Problem---------------------------

// Link: https://leetcode.com/problems/search-in-a-binary-search-tree/
func (tree BinarySearchTree) searchBST(val int) *Node {

	for tree.root != nil && tree.root.Val != val {
		if val < tree.root.Val {
			tree.root = tree.root.Left
		} else {
			tree.root = tree.root.Right
		}
	}
	return tree.root
}

// Link: https://practice.geeksforgeeks.org/problems/minimum-element-in-bst/1
// Returns minimum value in a given Binary Tree
func FindMin(root *Node) int {
	//1. Check if the given node is NULL
	if root == nil {
		return MaxInt
	}

	//2. Return maximum of 3 values: 1) Root's data 2) Max in Left Subtree 3) Max in Right subtree
	res := root.Val
	lres := FindMin(root.Left)
	rres := FindMin(root.Right)

	if lres < res {
		res = lres
	}
	if rres < res {
		res = rres
	}

	return res
}

// Returns maximum value in a given Binary Tree
func FindMax(root *Node) int {
	//1. Check if the given node is NULL
	if root == nil {
		return MinInt
	}

	//2. Return maximum of 3 values: 1) Root's data 2) Max in Left Subtree 3) Max in Right subtree
	res := root.Val
	lres := FindMax(root.Left)
	rres := FindMax(root.Right)

	if lres > res {
		res = lres
	}
	if rres > res {
		res = rres
	}

	return res
}

// Link: https://leetcode.com/problems/insert-into-a-binary-search-tree/
func insertIntoBST(root *Node, val int) *Node {

	if root == nil {
		return &Node{Val: val}
	}

	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// Link: https://leetcode.com/problems/delete-node-in-a-bst/
func deleteNode(root *Node, key int) *Node {

	if root == nil {
		return root
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}
	return root
}

func findMin(root *Node) *Node {
	for root.Left != nil {
		root = root.Left
	}

	return root
}

// Link: https://takeuforward.org/data-structure/kth-largest-smallest-element-in-binary-search-tree/
func kthSmallest(root *Node, key int) int {

	res := kthSmallesthelper(root, &key)
	if res != nil {
		return res.Val
	}
	return -1
}

func kthSmallesthelper(root *Node, key *int) *Node {
	if root == nil {
		return root
	}

	var Left *Node = kthSmallesthelper(root.Left, key)
	if Left != nil {
		return Left
	}
	*(key)--

	if *(key) == 0 {
		return root
	}
	return kthSmallesthelper(root.Right, key)
}

// kth largest element
func kthLargest(root *Node, key int) int {

	res := kthLargesthelper(root, &key)
	if res != nil {
		return res.Val
	}
	return -1
}

func kthLargesthelper(root *Node, key *int) *Node {
	if root == nil {
		return root
	}

	var Right *Node = kthLargesthelper(root.Right, key)
	if Right != nil {
		return Right
	}
	*(key)--

	if *(key) == 0 {
		return root
	}
	return kthLargesthelper(root.Left, key)
}

// Link: https://leetcode.com/problems/validate-binary-search-tree/
// Link: https://afteracademy.com/blog/check-if-a-binary-tree-is-BST-or-not
func isValidBST(root *Node) bool {

	var arr []int
	storeInorder(root, &arr)
	for i := 1; i < len(arr); i++ {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func storeInorder(root *Node, arr *[]int) {

	if root == nil {
		return
	}
	storeInorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	storeInorder(root.Right, arr)
}

// Link: https://www.geeksforgeeks.org/floor-and-ceil-from-a-bst/
func ceil(root *Node, key int) int {

	var ceil = -1
	for root != nil {
		if key == root.Val {
			ceil = root.Val
			return ceil
		}

		if key > root.Val {
			root = root.Right
		} else {
			ceil = root.Val
			root = root.Left
		}
	}
	return ceil
}

func floor(root *Node, key int) int {

	var floor = -1
	for root != nil {
		if key == root.Val {
			floor = root.Val
			return floor
		}

		if key < root.Val {
			root = root.Left
		} else {
			floor = root.Val
			root = root.Right
		}
	}
	return floor
}

// Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
func lowestCommonAncestor(root, p, q *Node) *Node {

	if root == nil {
		return nil
	}
	var cur = root.Val

	if p.Val > cur && q.Val > cur {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if p.Val < cur && q.Val < cur {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

// Link: https://www.geeksforgeeks.org/inorder-predecessor-successor-given-key-bst/
func inOrderSuccessor(root, p *Node) *Node {
	var sucessor *Node = nil

	for root != nil {
		if p.Val >= root.Val {
			root = root.Right
		} else {
			sucessor = root
			root = root.Left
		}
	}
	return sucessor
}

func inOrderpredecessor(root, p *Node) *Node {
	var predecessor *Node = nil

	for root != nil {
		if p.Val > root.Val {
			predecessor = root
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return predecessor
}

// Link: https://leetcode.com/problems/binary-search-tree-iterator/
type BSTIterator struct {
	item []*Node
}

func Constructor(root *Node) BSTIterator {
	itr := BSTIterator{}
	for root != nil {
		itr.item = append(itr.item, root)
		root = root.Left
	}
	return itr
}

func (this *BSTIterator) Next() int {

	n := len(this.item)
	cur := this.item[n-1]
	this.item = this.item[:n-1]
	res := cur.Val
	cur = cur.Right

	for cur != nil {
		this.item = append(this.item, cur)
		cur = cur.Left
	}
	return res
}

func (this *BSTIterator) HasNext() bool {
	return len(this.item) > 0
}

// Link: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

// res tsore the the map of elment we require and target sum
type Res struct {
	nodeMap   map[int]int
	targetSum int
}

func findTarget(root *Node, k int) bool {
	r := Res{nodeMap: make(map[int]int), targetSum: k}
	return r.dfs(root)
}

func (r *Res) dfs(node *Node) bool {

	if node == nil {
		return false
	}
	//search if current node is present in map
	//if it is present then return true
	if _, ok := r.nodeMap[node.Val]; ok {
		return true
	}
	//if elment is not present then calculate the required node from target and current node
	//add the required node in map
	r.nodeMap[r.targetSum-node.Val]++

	return r.dfs(node.Left) || r.dfs(node.Right)
}
