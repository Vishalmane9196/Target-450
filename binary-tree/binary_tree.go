package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
)

/********Binary tree implementation***********/
type Node struct {
	key       int
	value     int
	leftNode  *Node
	rightNode *Node
}

type BinaryTree struct {
	rootNode *Node
	lock     sync.RWMutex
}

// InsertElement insert new node with given key and value
func (tree *BinaryTree) InsertElement(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var newNode = &Node{key: key, value: value, leftNode: nil, rightNode: nil}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		InsertElementHelper(tree.rootNode, newNode)
	}
}

// InsertElementHelper is helper function for InsertElement to find the correct place for a node in a tree
func InsertElementHelper(rootNode *Node, newNode *Node) {
	if newNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			InsertElementHelper(rootNode.leftNode, newNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			InsertElementHelper(rootNode.rightNode, newNode)
		}
	}
}

// SearchNode returns whether given node exist or not
func (tree *BinaryTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNodeHelper(tree.rootNode, key)
}

// searchNodeHelper is helper function for SearchNode
func searchNodeHelper(node *Node, searchKey int) bool {
	if node == nil {
		return false
	}
	if searchKey < node.key {
		return searchNodeHelper(node.leftNode, searchKey)
	}
	if searchKey > node.key {
		return searchNodeHelper(node.rightNode, searchKey)
	}
	return true
}

// RemoveNode removes node from binary tree with given key
func (tree *BinaryTree) RemoveNode(deleteKey int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNodeHelper(tree.rootNode, deleteKey)
}

// removeNodeHelper is helper function for RemoveNode
func removeNodeHelper(parentNode *Node, deleteKey int) *Node {
	if parentNode == nil {
		return nil
	}

	if deleteKey < parentNode.key {
		parentNode.leftNode = removeNodeHelper(parentNode.leftNode, deleteKey)
		return parentNode
	}
	if deleteKey > parentNode.key {
		parentNode.rightNode = removeNodeHelper(parentNode.rightNode, deleteKey)
		return parentNode
	}

	//case 1: both leaf node are nil
	if parentNode.leftNode == nil && parentNode.rightNode == nil {
		parentNode = nil
		return nil
	}

	//case 2: right node is not nil
	if parentNode.leftNode == nil {
		parentNode = parentNode.rightNode
		return parentNode
	}
	//case 3: left node is not nil
	if parentNode.rightNode == nil {
		parentNode = parentNode.leftNode
		return parentNode
	}

	/*case 4: both left and right node is present
	leftmostrightNode is the node which is greater than all the nodes of parent node and
	first element greater element than parent node
	this is the element we can keep in  place of deleted node inorder to satisfy the binary tree property
	*/
	leftmostRightNode := parentNode.rightNode
	for {
		//check if first right node of parent is nil or not
		//if that node is not nil then check it's left node nil or not
		//once left node become nil break out of loop
		if leftmostRightNode != nil && leftmostRightNode.leftNode != nil {
			leftmostRightNode = leftmostRightNode.leftNode
		} else {
			break
		}
	}

	//update the values of to be deleted node(parent node) with leftmostrightNode node
	parentNode.key, parentNode.value = leftmostRightNode.key, leftmostRightNode.value
	/*delete the rightnode of parent node because it is still pointing to old binary structure
	now the parentNode key and value is of lefmostRightNode, so we are deleteing the leftmostRightNode from right subtree of parentNode with parentNode key
	because the leftmostRightNode is still present in right tree of ParentNode
	*/
	parentNode.rightNode = removeNodeHelper(parentNode.rightNode, parentNode.key)
	return parentNode
}

// PrintBinaryTree print binary tree
func (tree BinaryTree) PrintBinaryTree() {

	str := tree.toTreeString("", true, "", tree.rootNode)
	fmt.Println("***************************************************************")
	fmt.Println(str)
	fmt.Println("***************************************************************")
}

// toTreeString is helper function for PrintBinaryTree
func (tree BinaryTree) toTreeString(prefix string, top bool, str string, node *Node) string {

	if node == nil {
		return ""
	}

	Left := node.leftNode
	Right := node.rightNode

	if Right != nil {
		temp := tree.path(top, ""+prefix, "│   ", "    ")
		str = tree.toTreeString(temp, false, str, Right)
	}

	str = tree.path(top, str+prefix, "└──", "┌──")
	str = str + " " + strconv.Itoa(node.value) + "\n"

	if Left != nil {
		temp := tree.path(top, ""+prefix, "    ", "│   ")
		str = tree.toTreeString(temp, true, str, Left)
	}
	return str
}

// path is helper function for toTreeString
func (tree BinaryTree) path(condition bool, str string, choice1 string, choice2 string) string {

	if condition {
		str += choice1
	} else {
		str += choice2
	}
	return str
}

// InOrderTraverseTree traversal ====> [Left <-> Print <-> Right]
func (tree *BinaryTree) InOrderTraverseTree() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	fmt.Println("In-order traversal :  ")
	InOrderTraverseTreeHelper(tree.rootNode)
}

// InOrderTraverseTreeHelper is a helper function for InOrderTraverseTree
func InOrderTraverseTreeHelper(Node *Node) {
	if Node != nil {
		InOrderTraverseTreeHelper(Node.leftNode)
		fmt.Printf("[%v]\t", Node.value)
		InOrderTraverseTreeHelper(Node.rightNode)
	}
}

// PreOrderTraversetraversal ====> [Print <-> Left <-> Right]
func (tree *BinaryTree) PreOrderTraverseTree() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println()
	fmt.Println("Pre-order traversal :  ")
	PreOrderTraverseTreeHelper(tree.rootNode)
}

// preOrderTraverseTree is a helper function for PreOrderTraverseTree
func PreOrderTraverseTreeHelper(Node *Node) {
	if Node != nil {
		fmt.Printf("[%v]\t", Node.value)
		PreOrderTraverseTreeHelper(Node.leftNode)
		PreOrderTraverseTreeHelper(Node.rightNode)
	}
}

// PostOrderTraverseTree ====> [Left <-> Right <-> Print]
func (tree *BinaryTree) PostOrderTraverseTree() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println()
	fmt.Println("Post-order traversal :  ")
	PostOrderTraverseTreeHelper(tree.rootNode)
}

// PostOrderTraverseTreeHelper is a helper function for PostOrderTraverseTree
func PostOrderTraverseTreeHelper(Node *Node) {
	if Node != nil {
		PostOrderTraverseTreeHelper(Node.leftNode)
		PostOrderTraverseTreeHelper(Node.rightNode)
		fmt.Printf("[%v]\t", Node.value)
	}
}

// MinNode method --> traverse to extreme left to find out min node
func (tree *BinaryTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	Node := tree.rootNode
	if Node == nil {
		return (*int)(nil)
	}
	//initiated an infinite loop --> break condition : node's left become nil
	for {
		if Node.leftNode == nil {
			return &Node.value
		}
		Node = Node.leftNode
	}
}

// MaxNode method ---> traverse to extreme right to find out maz node
func (tree *BinaryTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	Node := tree.rootNode
	if Node == nil {
		return (*int)(nil)
	}
	//initiated an infinite loop --> break condition : node's right become nil
	for {
		if Node.rightNode == nil {
			return &Node.value
		}
		Node = Node.rightNode
	}
}

/***********Problems***********/
/*
BFS traversal : level order traversal
DFS traversal : Preorder, PostOrder, InOrder traversal
*/

// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/
func levelOrder(root *Node) [][]int {
	//create 2d slice to store tree with levels
	var result [][]int

	if root == nil {
		return result
	}

	//created a queue to store node
	var queue []*Node
	//append root to queue
	queue = append(queue, root)

	//initiated a infinite loop
	//break condition: len(que) < 0
	for len(queue) > 0 {
		//store a queuesize because at runtime size of queue will change
		queueSize := len(queue)
		//slice to store the nodes from the level
		var levelNodes []int

		//iterate over all nodes till queuesize
		for i := 0; i < queueSize; i++ {
			//store the first node which will be popped
			dequeuedElement := (queue)[0]
			queue = queue[1:]

			//append the popped element in levelNodes slice
			levelNodes = append(levelNodes, dequeuedElement.value)

			//if popped element has left node then append that node to queue
			if dequeuedElement.leftNode != nil {
				queue = append(queue, dequeuedElement.leftNode)
			}
			//if popped element has right node then append that node to queue
			if dequeuedElement.rightNode != nil {
				queue = append(queue, dequeuedElement.rightNode)
			}
		}
		//append levelorder slice in result
		result = append(result, levelNodes)
	}
	return result
}

// Link: https://leetcode.com/problems/binary-tree-preorder-traversal/
func preorderTraversalIterative(root *Node) []int {

	//store the preorder traversal
	var res []int

	if root == nil {
		return res
	}

	//stack to store the node
	var stack []*Node
	//append the root to stack
	stack = append(stack, root)

	//Initiated infinite loop
	//break condition: len(stack) < 0
	for len(stack) > 0 {
		n := len(stack)
		//pop the topmost element
		poppedElement := stack[n-1]
		stack = stack[:len(stack)-1]
		//append opeed element to stack
		res = append(res, poppedElement.value)

		//append the right node of popped element in stack
		if poppedElement.rightNode != nil {
			stack = append(stack, poppedElement.rightNode)
		}

		//append the left node of popped element in stack
		if poppedElement.leftNode != nil {
			stack = append(stack, poppedElement.leftNode)
		}
	}
	return res
}

// Link: https://leetcode.com/problems/binary-tree-preorder-traversal/
func inorderTraversalIterative(root *Node) []int {

	//store the inorder traversal
	var res []int

	if root == nil {
		return res
	}

	//stack to store the node
	var stack []*Node
	var node *Node = root

	//Initiated infinite loop
	//break condition: len(stack) == 0 and node become nil
	for node != nil || len(stack) > 0 {

		//if node is not nil then append the element and go to left
		if node != nil {
			stack = append(stack, node)
			node = node.leftNode
		} else {

			//store topmost element and pop that element from stack
			lastEle := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			//store the popped element in res slice and move to right
			res = append(res, lastEle.value)
			node = lastEle.rightNode
		}
	}
	return res
}

// Link: https://leetcode.com/problems/binary-tree-preorder-traversal/
func postorderTraversalIterative(root *Node) []int {

	//store the preorder traversal
	var res []int

	if root == nil {
		return res
	}

	//stack to store the node
	var stack1, stack2 []*Node
	//append the root to stack
	stack1 = append(stack1, root)

	//Initiated infinite loop
	//break condition: len(stack) < 0
	for len(stack1) > 0 {
		n := len(stack1)
		//pop the topmost element
		poppedElement := stack1[n-1]
		stack1 = stack1[:n-1]
		stack2 = append(stack2, poppedElement)

		//append the left node of popped element in stack
		if poppedElement.leftNode != nil {
			stack1 = append(stack1, poppedElement.leftNode)
		}
		//append the right node of popped element in stack
		if poppedElement.rightNode != nil {
			stack1 = append(stack1, poppedElement.rightNode)
		}
	}

	for len(stack2) > 0 {
		n := len(stack2)
		res = append(res, stack2[n-1].value)
		stack2 = stack2[:n-1]
	}
	return res
}

// Link: https://takeuforward.org/data-structure/post-order-traversal-of-binary-tree/
// Link: https://leetcode.com/problems/binary-tree-postorder-traversal/
func postOrderTraversalUsing1Stack(root *Node) []int {
	//store the preorder traversal
	var res []int

	if root == nil {
		return res
	}

	//stack to store the node
	var stack []*Node
	//append the root to stack
	node := root

	for node != nil || len(stack) > 0 {

		if node != nil {
			stack = append(stack, node)
			node = node.leftNode
		} else {
			temp := stack[len(stack)-1].rightNode

			if temp == nil {
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				res = append(res, temp.value)

				for len(stack) > 0 && temp == stack[len(stack)-1].rightNode {
					temp = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					res = append(res, temp.value)
				}
			} else {
				node = temp
			}
		}
	}
	return res
}

// Link: https://takeuforward.org/data-structure/preorder-inorder-postorder-traversals-in-one-traversal/
type pair struct {
	node *Node
	num  int
}

func allTraversal(root *Node) ([]int, []int, []int) {
	var preOrder []int
	var inOrder []int
	var postOrder []int

	if root == nil {
		return preOrder, inOrder, postOrder
	}

	//stack to store the node
	var stack []pair
	//append the root to stack
	node := root

	stack = append(stack, pair{node, 1})

	for len(stack) > 0 {

		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		//case 1
		if it.num == 1 {
			preOrder = append(preOrder, it.node.value)
			it.num++
			stack = append(stack, it)

			if it.node.leftNode != nil {
				stack = append(stack, pair{it.node.leftNode, 1})
			}
		} else if it.num == 2 {
			//case 2
			inOrder = append(inOrder, it.node.value)
			it.num++
			stack = append(stack, it)

			if it.node.rightNode != nil {
				stack = append(stack, pair{it.node.rightNode, 1})
			}
		} else {
			// case 3
			postOrder = append(postOrder, it.node.value)
		}
	}

	return preOrder, inOrder, postOrder

}

// Link: https://takeuforward.org/data-structure/maximum-depth-of-a-binary-tree/
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	lh := maxDepth(node.leftNode)
	rh := maxDepth(node.rightNode)

	return 1 + maxi(lh, rh)
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Link: https://takeuforward.org/data-structure/check-if-the-binary-tree-is-balanced-binary-tree/
func isBalanced(root *Node) bool {
	return isBalancedHelper(root) != -1
}

func isBalancedHelper(root *Node) int {
	if root == nil {
		return 0
	}
	leftHeight := isBalancedHelper(root.leftNode)
	if leftHeight == -1 {
		return -1
	}
	rightHeight := isBalancedHelper(root.leftNode)
	if rightHeight == -1 {
		return -1
	}

	if Abs(leftHeight-rightHeight) > 1 {
		return -1
	}
	return maxi(leftHeight, rightHeight) + 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Link: https://takeuforward.org/data-structure/calculate-the-diameter-of-a-binary-tree/
func diameterOfBinaryTree(root *Node) int {
	var diameter int

	if root == nil {
		return 0
	}
	if root.leftNode == nil && root.rightNode == nil {
		return 0
	}
	diameterOfTree(root, &diameter)
	return diameter
}

func diameterOfTree(node *Node, diameter *int) int {
	if node == nil {
		return 0
	}
	lh := diameterOfTree(node.leftNode, diameter)
	rh := diameterOfTree(node.rightNode, diameter)

	*diameter = maxi(*diameter, lh+rh)
	return 1 + maxi(lh, rh)
}

// Link: https://leetcode.com/problems/binary-tree-maximum-path-sum/
// golabl varibale to store maximum sum
var maxSum int

func maxPathSum(root *Node) int {
	//initialize to minimum Value
	maxSum = math.MinInt32
	find(root)
	return maxSum
}

func find(root *Node) int {
	//Base Condition
	if root == nil {
		return 0
	}
	//find the sum of left subtree
	left := find(root.leftNode)
	//find the sum of right sub tree
	right := find(root.rightNode)
	//find th maxsumpath between root and and it's subtree
	temp := max(max(left, right)+root.value, root.value)
	// find th maxpath between its subtree a including node itself
	ans := max(temp, left+right+root.value)
	// check th max between them
	maxSum = max(maxSum, ans)
	return temp
}

// Link: https://leetcode.com/problems/same-tree/
func isSameTree(p *Node, q *Node) bool {

	if p == nil || q == nil {
		return p == q
	}
	return p.value == q.value && isSameTree(p.leftNode, q.leftNode) && isSameTree(p.rightNode, q.rightNode)
}

// Link: https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *Node) *Node {

	if root == nil {
		return nil
	}
	if root.value == p.value || root.value == q.value {
		return root
	}
	lcaLeft := lowestCommonAncestor(root.leftNode, p, q)
	lcaRight := lowestCommonAncestor(root.rightNode, p, q)

	if lcaLeft != nil && lcaRight != nil {
		return root
	}
	if lcaLeft != nil {
		return lcaLeft
	}
	return lcaRight
}

// Link: https://leetcode.com/problems/binary-tree-right-side-view/
// to get the left view just change the left to right seq of rec
func rightSideView(root *Node) []int {
	var ans []int
	var queue []*Node

	if root == nil {
		return nil
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		var helper []int
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			helper = append(helper, node.value)
			if node.leftNode != nil {
				queue = append(queue, node.leftNode)
			}
			if node.rightNode != nil {
				queue = append(queue, node.rightNode)
			}
		}
		// fmt.Println("helper : ", helper)
		ans = append(ans, helper[len(helper)-1])
	}
	return ans
}

// above solution in recursive form
func rightSideViewRec(root *Node) []int {
	var ans []int
	rec(root, 0, &ans)
	// fmt.Println("ans : ", ans)
	return ans
}

func rec(root *Node, level int, ans *[]int) {

	if root == nil {
		return
	}
	if len(*ans) == level {
		*ans = append(*ans, root.value)
	}
	rec(root.rightNode, level+1, ans)
	rec(root.leftNode, level+1, ans)
}

// Link: https://www.geeksforgeeks.org/bottom-view-binary-tree/
// Here we are checking if the line exist in map or not
// if it does not exist then we are adding that line with node value
// if it does exist we are update the map with corrosponding value
func bottomView(root *Node) []int {
	var ans []int
	mp := map[int]int{}
	var queue []pair
	if root == nil {
		return ans
	}

	queue = append(queue, pair{root, 0})
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]
		node := it.node
		line := it.num

		mp[line] = node.value

		if node.leftNode != nil {
			queue = append(queue, pair{node.leftNode, line - 1})
		}
		if node.rightNode != nil {
			queue = append(queue, pair{node.rightNode, line + 1})
		}
	}

	for _, v := range mp {
		ans = append(ans, v)
	}
	//sort map
	sort.Ints(ans)
	// fmt.Println("ans : ", ans)
	return ans
}

// Link: https://www.geeksforgeeks.org/bottom-view-binary-tree/
// For top view the only change other than bottom view is that we are checking if the line exist in map or not
// if it does not exist then we are adding that line with node value
// if it does exist we don't have update the map with corrosponding value
func topView(root *Node) []int {
	var ans []int
	mp := map[int]int{}
	var queue []pair
	if root == nil {
		return ans
	}

	queue = append(queue, pair{root, 0})
	for len(queue) > 0 {
		it := queue[0]
		queue = queue[1:]
		node := it.node
		line := it.num

		if _, exist := mp[line]; !exist {
			mp[line] = node.value
		}

		if node.leftNode != nil {
			queue = append(queue, pair{node.leftNode, line - 1})
		}
		if node.rightNode != nil {
			queue = append(queue, pair{node.rightNode, line + 1})
		}
	}

	for _, v := range mp {
		ans = append(ans, v)
	}
	//sort map
	sort.Ints(ans)
	// fmt.Println("ans : ", ans)
	return ans
}

//Link:

// Tuple to store each node's data such as column, row and value
type Tuple struct {
	col int
	row int
	val int
}

func verticalTraversal(root *Node) [][]int {
	//nodes will consist of all nodes data with tuple atribute
	nodes := dfs(root, 0, 0, []Tuple{})

	// fmt.Println("nodes : ", nodes)
	// Less function for lexicograplical order to sort the nodes based on columns ot each node
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		} else if nodes[i].row != nodes[j].row {
			return nodes[i].row < nodes[j].row
		} else if nodes[i].val != nodes[j].val {
			return nodes[i].val < nodes[j].val
		} else {
			return false
		}
	})
	// fmt.Println("After sort nodes : ", nodes)

	// Grouping sorted elements
	ans := make([][]int, 0)
	//This hack will be used to create new dynamic slice  for the first node
	prev := nodes[0].col + 1
	for _, node := range nodes {
		//if prev column is not the same of current column then we will create new slice
		if node.col != prev {
			ans = append(ans, []int{})
		}
		//and append the node's value
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
		//set prev node column as current node's columns
		prev = node.col
	}
	return ans
}

// dfs will give us list of tiples which consist of the information about each node such as row, col and value
func dfs(root *Node, row int, col int, nodes []Tuple) []Tuple {
	if root == nil {
		return nodes
	}

	nodes = append(nodes, Tuple{col, row, root.value})
	nodes = dfs(root.leftNode, row+1, col-1, nodes)
	nodes = dfs(root.rightNode, row+1, col+1, nodes)

	return nodes
}

// Link: https://leetcode.com/problems/symmetric-tree/
func isSymmetric(root *Node) bool {
	return root == nil || isSymmetricHelper(root.leftNode, root.rightNode)
}

func isSymmetricHelper(left, right *Node) bool {

	if left == nil || right == nil {
		return left == right
	}

	if left.value != right.value {
		return false
	}
	return isSymmetricHelper(left.leftNode, right.rightNode) && isSymmetricHelper(left.rightNode, right.leftNode)
}

// Link: https://takeuforward.org/data-structure/boundary-traversal-of-a-binary-tree/
func printBoundary(root *Node) []int {
	ans := make([]int, 0)
	if isLeaf(root) == false {
		ans = append(ans, root.value)
	}
	addLeftBoundary(root, &ans)
	addLeaves(root, &ans)
	addRightBoundary(root, &ans)
	return ans
}

func addLeftBoundary(node *Node, ans *[]int) {
	// fmt.Println("Inside addleft")
	cur := node.leftNode
	for cur != nil {
		if isLeaf(cur) == false {
			*ans = append(*ans, cur.value)
		}

		if cur.leftNode != nil {
			cur = cur.leftNode
		} else {
			cur = cur.rightNode
		}
	}
}

func addLeaves(node *Node, ans *[]int) {
	// fmt.Println("Inside add leaves")
	if isLeaf(node) {
		*ans = append(*ans, node.value)
		return
	}
	if node.leftNode != nil {
		addLeaves(node.leftNode, ans)
	}
	if node.rightNode != nil {
		addLeaves(node.rightNode, ans)
	}
}

func addRightBoundary(node *Node, ans *[]int) {
	// fmt.Println("Inside add right ")
	temp := make([]int, 0)
	cur := node.rightNode

	for cur != nil {
		if isLeaf(cur) == false {
			temp = append(temp, cur.value)
		}
		if cur.rightNode != nil {
			cur = cur.rightNode
		} else {
			cur = cur.leftNode
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		*ans = append(*ans, temp[i])
	}
}

func isLeaf(node *Node) bool {
	if node.leftNode == nil && node.rightNode == nil {
		return true
	}
	return false
}

// Link: https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/
func zigzagLevelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	var leftToRight = true

	for len(queue) > 0 {

		n := len(queue)
		var row = make([]int, n)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			var index int
			if leftToRight {
				index = i
			} else {
				index = n - 1 - i
			}
			// fmt.Println("index : ", index)
			row[index] = node.value

			if node.leftNode != nil {
				queue = append(queue, node.leftNode)
			}
			if node.rightNode != nil {
				queue = append(queue, node.rightNode)
			}
		}
		leftToRight = !leftToRight
		res = append(res, row)
	}
	return res
}

// Link: https://takeuforward.org/data-structure/print-root-to-node-path-in-a-binary-tree/
func getPath(root *Node, ans *[]int, searchData int) bool {

	if root == nil {
		return false
	}

	*ans = append(*ans, root.value)
	if root.value == searchData {
		return true
	}

	if getPath(root.leftNode, ans, searchData) || getPath(root.rightNode, ans, searchData) {
		return true
	}

	*ans = (*ans)[:len(*ans)-1]
	return false
}

// Link: https://leetcode.com/problems/maximum-width-of-binary-tree/
func widthOfBinaryTree(root *Node) int {

	var queue []pair
	var ans = 0
	if root == nil {
		return 0
	}

	queue = append(queue, pair{root, 0})
	for len(queue) > 0 {
		n := len(queue)
		curMin := queue[0].num
		var leftMost, rightMost int

		for i := 0; i < n; i++ {

			p := queue[0]
			cur_id := p.num - curMin
			temp := p.node
			queue = queue[1:]

			if i == 0 {
				leftMost = cur_id
			}
			if i == n-1 {
				rightMost = cur_id
			}

			if temp.leftNode != nil {
				queue = append(queue, pair{temp.leftNode, 2*cur_id + 1})
			}
			if temp.rightNode != nil {
				queue = append(queue, pair{temp.rightNode, 2*cur_id + 2})
			}
		}
		ans = maxi(ans, rightMost-leftMost+1)
	}
	return ans
}

// Link: https://leetcode.com/problems/add-one-row-to-tree/
func addOneRow(root *Node, val int, depth int) *Node {

	if depth == 1 {
		return &Node{val, val, root, nil}
	}

	queue := []*Node{}
	queue = append(queue, root)
	for len(queue) > 0 && depth != 1 {
		size := len(queue)
		depth -= 1

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.leftNode != nil {
				queue = append(queue, node.leftNode)
			}

			if node.rightNode != nil {
				queue = append(queue, node.rightNode)
			}

			if depth == 1 {
				node.leftNode = &Node{val, val, node.leftNode, nil}
				node.rightNode = &Node{val, val, nil, node.rightNode}
			}
		}
		if depth == 1 {
			break
		}
	}
	return root
}

// Link: https://leetcode.com/problems/count-complete-tree-nodes/
func countNodes(root *Node) int {

	var res int

	if root == nil {
		return res
	}

	var stack []*Node
	stack = append(stack, root)

	for len(stack) > 0 {
		n := len(stack)
		//pop the topmost element
		poppedElement := stack[n-1]
		stack = stack[:len(stack)-1]
		res += 1

		//append the right node of popped element in stack
		if poppedElement.rightNode != nil {
			stack = append(stack, poppedElement.rightNode)
		}

		//append the left node of popped element in stack
		if poppedElement.leftNode != nil {
			stack = append(stack, poppedElement.leftNode)
		}
	}
	return res
}

// Link: https://leetcode.com/problems/leaf-similar-trees/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := []int{}, []int{}
	dfs(root1, &s1)
	dfs(root2, &s2)
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*s = append(*s, root.Val)
	}
	dfs(root.Left, s)
	dfs(root.Right, s)
}

// Link: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	helper(root, &res, root.Val, root.Val)
	return res
}

func helper(cur *TreeNode, res *int, mn, mx int) {
	if cur == nil {
		return
	}
	currentMax := max(abs(cur.Val-mn), abs(cur.Val-mx))
	*res = max(*res, currentMax)
	currentMin := min(mn, cur.Val)
	cuurentMax := max(mx, cur.Val)
	helper(cur.Left, res, currentMin, cuurentMax)
	helper(cur.Right, res, currentMin, cuurentMax)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}
