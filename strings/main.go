package main

import "fmt"

func main() {

	// input := "(()())(())"
	// res := removeOutermostParenthesis(input)
	// fmt.Println("res : ", res)

	// input := "(1+(2*3)+((8)/4))+1"
	// res := parenthesisDepth(input)
	// fmt.Println("Depth of parenthesis  : ", res)

	// word := "racepar"
	// res := isPalindrome(word)
	// fmt.Println("isPalindrome : ", res)

	// directions := "WNEENESENNNE"
	// res := getShortestPath(directions)
	// fmt.Println("getShortestPath : ", res)

	// sentence := "hi i am vishal"
	// res := toUpperCase(sentence)
	// fmt.Println("toUpperCase : ", res)

	// sentence := "aaaaabbccccdd"
	// res := compressString(sentence)
	// fmt.Println("compressString : ", res)

	// input := "the sky is blue"
	// input := "  hello world  "
	// input := "example   good a"
	// res := reverseWords(input)
	// fmt.Println("res : ", res)

	// input := "34564"
	// res := largestOddNumber(input)
	// fmt.Println("res : ", res)

	// input :=  []string{"flower","flow","flight"}
	// res := longestCommonPrefix(input)
	// fmt.Println("res : ", res)

	// var s,t = "paper","title"
	// var s,t = "foo","bar"
	// res := isIsomorphic(s,t)
	// fmt.Println("res : ", res)

	// var s,t = "abcde","cdeab"
	// res := rotateString(s,t)
	// fmt.Println("res : ", res)

	// var s, t = "anagram", "nagaram"
	// res := isAnagram(s, t)
	// fmt.Println("res : ", res)

	// res := breakPalindrome("abccba")
	// fmt.Println("res : ", res)

	// str := []string{"a", "a", "b", "b", "c", "c", "c"}
	// str := "abbbbbbbbbbbb"
	// str := "aabbccc"
	// res := compress([]byte(str))
	// fmt.Println("compress1 : ", res)

	// str := "(1+(2*3)+((8)/4))+1"
	// res := maxDepth(str)
	// fmt.Println("MaxDepth : ", res)

	// str := "III"
	// str := "LVIII"
	// res := romanToInt(str)
	// fmt.Println("romanToInt : ", res)

	// number := 12
	// res := IntToRoman(number)
	// fmt.Println("IntToRoman : ", res)

	// str := "abc"
	// res := subString(str)
	// fmt.Println("subString : ", res)

	// str := "abc"
	// res := allStringPermutations(str)
	// fmt.Println("allStringPermutations : ", res)

	// arr := []int{1, 2, 3}
	// res := permute(arr)
	// fmt.Println("permute : ", res)

	// str := "aabcb"
	// res := beautySum(str)
	// fmt.Println("beautySum : ", res)

	// res := countAndSay(4)
	// fmt.Println("countAndSay : ", res)

	// var (
	// 	s = "ADOBECODEBANC"
	// 	t = "ABC"
	// )
	// res := minWindow(s, t)
	// fmt.Println("minWindow : ", res)

	// var (
	// 	// word1 = []string{"a", "cb"}
	// 	// word2 = []string{"ab", "c"}
	// 	word1 = []string{"abc", "d", "defg"}
	// 	word2 = []string{"abcddefg"}
	// )
	// res := arrayStringsAreEqual(word1, word2)
	// fmt.Println("arrayStringsAreEqual : ", res)

	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// res := groupAnagrams(strs)
	// fmt.Println("arrayStringsAreEqual : ", res)

	// strs := []string{"ab", "ty", "yt", "lc", "cl", "ab"}
	// res := longestPalindrome(strs)
	// fmt.Println("longestPalindrome : ", res)

	// strs := "leetcode"
	// res := reverseVowels(strs)
	// fmt.Println("reverseVowels : ", res)

	// strs := "baaca"
	// k := 3
	// res := orderlyQueue(strs, k)
	// fmt.Println("orderlyQueue : ", res)

	// strs := "leEeetcode"
	// res := makeGood(strs)
	// fmt.Println("makeGood : ", res)

	// strs := "abbaca"
	// res := removeDuplicates(strs)
	// fmt.Println("makeGood : ", res)

	// strs := "book"
	// strs := "Ieai"
	// strs := "textbook"
	// res := halvesAreAlike(strs)
	// fmt.Println("halvesAreAlike : ", res)

	// word1 := "abc"
	// word2 := "bca"
	// res := closeStrings(word1, word2)
	// fmt.Println("closeStrings : ", res)

	word := "tree"
	res := frequencySort(word)
	fmt.Println("frequencySort : ", res)

}
