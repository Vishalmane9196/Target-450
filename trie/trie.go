package main

/*********Trie Implementation*******/
//Link: https://leetcode.com/problems/implement-trie-prefix-tree/
type trie_Node struct {
	childrens [26]*trie_Node
	wordEnds  bool
}

type Trie struct {
	root *trie_Node
}

func Constructor() Trie {
	t := Trie{nil}
	t.root = new(trie_Node)
	return t
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = new(trie_Node)
		}
		current = current.childrens[index]
	}
	current.wordEnds = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for _, wr := range word {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.wordEnds {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for _, wr := range prefix {
		index := wr - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	return true
}

/*******Problem*********/
//Link : https://leetcode.com/problems/word-search-ii/
type TrieNode struct {
	char     byte
	children map[byte]*TrieNode
	isEnd    bool
	word     string
}

func New(c byte) *TrieNode {
	return &TrieNode{c, make(map[byte]*TrieNode), false, ""}
}

func trieInsert(root *TrieNode, s string) {
	var curr = root
	for i := 0; i < len(s); i++ {
		var children = curr.children
		var node, ok = children[s[i]]

		if ok {

			// if the element is found and string is at the end
			// Make the isEnd true  and return
			if i == len(s)-1 {
				node.isEnd = true
				node.word = s
			}

			curr = node
		} else {
			var newNode = New(s[i])
			newNode.isEnd = i == len(s)-1
			children[s[i]] = newNode
			newNode.word = s
			curr = newNode
		}
	}
}

func dfs(board [][]byte, curr *TrieNode, row, col int, res *[]string) {
	var n = len(board)
	var m = len(board[0])
	var ch = board[row][col]

	var node, ok = curr.children[ch]
	if ch == '*' || !ok {
		return
	}

	curr = node
	if curr.isEnd {
		*res = append(*res, curr.word)
		curr.isEnd = false
	}

	// Mark the current cell as visited
	// We mark cell visited by changing its content to star
	// But we store the current content for backtracking
	var tmp = board[row][col]
	board[row][col] = '*'

	// Even if we have found one word we might want to check others
	// Do the DFS to check other adjacent cell
	if row+1 < n {
		dfs(board, curr, row+1, col, res)
	}
	if row-1 >= 0 {
		dfs(board, curr, row-1, col, res)
	}
	if col+1 < m {
		dfs(board, curr, row, col+1, res)
	}
	if col-1 >= 0 {
		dfs(board, curr, row, col-1, res)
	}

	// Backtrack
	board[row][col] = tmp
}

func findWords(board [][]byte, words []string) []string {

	var root = New('/')
	// Build the trie using words
	for i := 0; i < len(words); i++ {
		trieInsert(root, words[i])
	}
	// Traverse each word and calculate the result
	var result []string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, root, i, j, &result)
		}
	}
	return result
}

//remaining
// func longestCommonPrefix(strs []string) string {
// 	ans := ""

// 	triObj := Constructor()
// 	for _, word := range strs {
// 		triObj.Insert(word)
// 	}

// 	firstWord := strs[0]
// 	for i := 0; i < len(firstWord); i++ {
// 		input := string(firstWord[i])
// 		fmt.Println("input : ", input)
// 		found := triObj.StartsWith(input)
// 		if !found {
// 			break
// 		} else if found && len(input) < len(firstWord) {
// 			input = input + string(firstWord[i+1])
// 			ans = input
// 		}
// 	}
// 	fmt.Println("ans : ", ans)
// 	return ans
// }

func longestCommonPrefix(strs []string) string {
	var result = ""
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
}
