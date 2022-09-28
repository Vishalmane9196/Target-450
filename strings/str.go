package main

import (
	"fmt"
	"strings"
)

// Link: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
func parenthesisDepth(str string) int {
	fmt.Println("String : ", str)
	max, count := 0, 0

	for _, v := range str {
		if v == '(' {
			count += 1
			//set max to count for depth
			if max < count {
				max = count
			}
		}

		if v == ')' {
			count -= 1
		}
	}
	return max
}

// Link: https://leetcode.com/problems/remove-outermost-parentheses/
func removeOutermostParenthesis(str string) string {
	fmt.Println("String : ", str)
	var sb strings.Builder
	sb.Grow(len(str))
	depth := 0

	for _, r := range str {

		if r == '(' && depth > 0 || r == ')' && depth > 1 {
			sb.WriteRune(r)
		}

		if r == '(' {
			depth++
		} else {
			depth--
		}
	}
	return sb.String()
}

// Link: https://takeuforward.org/data-structure/reverse-words-in-a-string/
func reverseWords(words string) string {

	//Brute force
	//create stack , append each word to it and while popping elment from stack just add space between them
	// stack := make([]string, 0)
	// resultString := ""
	// str := strings.Split(words, " ")
	// for _, v := range str {
	// 	stack = append(stack, v)
	// }

	// for i := len(stack) - 1; i >= 0; i-- {
	// 	resultString = resultString + " " + stack[i]
	// }
	// return resultString

	//another approach
	//here we have added each word from sentence in slice and by using 2 pointer we are just swapping from left to right
	str := strings.TrimSpace(words)
	word := strings.Fields(str)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return strings.Join(word, " ")
}

// Link: https://leetcode.com/problems/largest-odd-number-in-string/
func largestOddNumber(num string) string {

	// storing the length index of number
	var numberLength = len(num) - 1

	//checking whether the number with till last index is divisible by 2 and index of number
	for numberLength >= 0 && num[numberLength]%2 == 0 {
		numberLength--
	}
	//returning the number till index where number is not divisible by 2
	return num[:numberLength+1]
}

// Link: https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	var ans strings.Builder

	for i := 0; i < len(strs[0]); i++ {
		//storing current character in curr
		cur := strs[0][i]
		for j := 1; j < len(strs); j++ {
			//checking curr is mactches with other words from string slices
			if i == len(strs[j]) || strs[j][i] != cur {
				return ans.String()
			}
		}
		//writing the curr to ans stringbuilder
		ans.WriteString(string(cur))
	}

	return ans.String()
}

// Link: https://leetcode.com/problems/isomorphic-strings/
func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	smap := make(map[byte]byte) //*using byte type to avoid unneccesary byte to string conversion*
	tmap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := smap[s[i]]; !ok {
			smap[s[i]] = t[i] //*If we dont find the element in map, we add them mapping byte code of character from 's' to that of 't' as key , value respectively*
		}
		if _, ok := tmap[t[i]]; !ok {
			tmap[t[i]] = s[i] //* Similarily, mapping byte code of character from 't' to that of 's' as key , value respectively*
		}
		if smap[s[i]] != t[i] || tmap[t[i]] != s[i] {
			return false // If at any point we come across these characters are not mapped one-to-one i.e one character from s is mapped to two different characters of t or viceversa, we return False
		}
	}
	return true
	// If we dont come across the above condition , that means our input has one-to-one mapped character //set from 's' and 't' strings, hence we return True*
}

// Link: https://leetcode.com/problems/rotate-string/
func rotateString(s string, goal string) bool {

	if len(s) != len(goal) {
		return false
	}

	return strings.Contains(goal+goal, s)
}

// Link: https://takeuforward.org/data-structure/check-if-two-strings-are-anagrams-of-each-other/
func isAnagram(s string, t string) bool {

	var lenS = len(s)
	var lenT = len(t)

	if lenS != lenT {
		return false
	}

	//map to store each character and count number of it
	var sm = make(map[string]int)

	for _, v := range s {
		sm[string(v)]++
	}

	for _, v := range t {
		sm[string(v)]--
	}

	for _, v := range s {
		if sm[string(v)] != 0 {
			return false
		}
	}

	return true
}
