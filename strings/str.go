package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
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

// palindrome problem
func isPalindrome(str string) bool {
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-i-1] {
			return false
		}
	}
	return true
}

// getShortestPath
func getShortestPath(str string) float64 {

	x, y := 0, 0
	for _, v := range str {
		if v == 'E' {
			x++
		} else if v == 'W' {
			x--
		} else if v == 'S' {
			y--
		} else {
			y++
		}
	}

	x1 := float64(x * x)
	y1 := float64(y * y)
	return math.Sqrt(x1 + y1)
}

// toUpperCase
func toUpperCase(str string) string {

	var sb strings.Builder
	ch := strings.ToUpper(string(str[0]))
	sb.WriteString(ch)

	for i := 1; i < len(str); i++ {
		if str[i] == ' ' && i < len(str)-1 {
			sb.WriteByte(str[i])
			i++
			sb.WriteString(strings.ToUpper(string(str[i])))
		} else {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}

// compressString
func compressString(str string) string {
	var sb strings.Builder

	for i := 0; i < len(str); i++ {
		count := 1
		for i < len(str)-1 && str[i] == str[i+1] {
			count++
			i++
		}

		sb.WriteByte(str[i])
		if count > 1 {
			stringCount := strconv.Itoa(count)
			sb.WriteString(stringCount)
		}
	}
	return sb.String()
}

// Link: https://takeuforward.org/data-structure/reverse-words-in-a-string/
// Link: https://leetcode.com/problems/reverse-words-in-a-string/
func reverseWords(words string) string {

	//Brute force
	//create stack , append each word to it and while popping elment from stack just add space between them
	stack := make([]string, 0)
	resultString := ""
	words1 := strings.TrimSpace(words)
	str := strings.Split(words1, " ")
	for i, v := range str {
		fmt.Printf(" %v ===>%v\n", i, v)
		if v == "" {
			continue
		} else {
			stack = append(stack, v)
		}
	}

	fmt.Println("stack : ", stack)
	for i := len(stack) - 1; i >= 0; i-- {
		if len(stack) == i+1 {
			resultString = stack[i]
		} else {
			resultString = resultString + " " + stack[i]
		}
	}
	return resultString

	//another approach
	//here we have added each word from sentence in slice and by using 2 pointer we are just swapping from left to right
	// str := strings.TrimSpace(words)
	// word := strings.Fields(str)

	// for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
	// 	word[i], word[j] = word[j], word[i]
	// }
	// return strings.Join(word, " ")
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

// Link: https://leetcode.com/problems/break-a-palindrome/
func breakPalindrome(palindrome string) string {

	n := len(palindrome)
	//case 1: if string consist of only character
	if n == 1 {
		return ""
	}

	for i := 0; i < n/2; i++ {
		if palindrome[i] != 'a' {
			str := palindrome[:i] + "a" + palindrome[i+1:]
			return str
		}
	}

	str := palindrome[:n-1] + "b"
	return str
}

// Link: https://leetcode.com/problems/string-compression/
func compress(chars []byte) int {

	count := 1    // count of consecutive character
	resIndex := 0 // index of result chars

	//append the first character at first index
	chars[resIndex] = chars[0]
	for i := 1; i < len(chars); i++ {
		//start from 1st index
		if chars[resIndex] != chars[i] {
			// new character
			resIndex++
			resIndex = resIndex + assignCount(chars[resIndex:], count)
			// fmt.Printf("resIndex : %v  || chars[i]  :  %v \n", resIndex, chars[i])
			chars[resIndex] = chars[i]
			// fmt.Println("chars  : ", chars)
			count = 1
		} else {
			// same character
			count++
		}
	}
	resIndex++
	resIndex = resIndex + assignCount(chars[resIndex:], count)
	// fmt.Println("resIndex   :", resIndex)
	// fmt.Println("chars : ", chars)
	chars = chars[:resIndex+1]
	// fmt.Println("chars : ", chars)
	return resIndex
}

// convert count to byte slice and assign it to target byte slice
func assignCount(target []byte, count int) int {
	if count == 1 {
		// no need to assign
		return 0
	}
	countStr := strconv.Itoa(count)
	// fmt.Println("countstr : ", countStr)
	for i, v := range countStr {
		target[i] = byte(v)
		// fmt.Println("tartget[i] : ", string(v))
	}
	// fmt.Println("countstr : ", countStr)
	// fmt.Println("countstr len: ", len(countStr))
	return len(countStr)
}

// Link: https://leetcode.com/problems/maximum-nesting-depth-of-the-parentheses/
func maxDepth(s string) int {
	max, count := 0, 0

	for _, v := range s {
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

// Link: https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {

	romanMap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	result := romanMap[s[len(s)-1]]

	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}
	return result
}

// Link: https://leetcode.com/problems/integer-to-roman/
func IntToRoman(num int) string {

	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for num > 0 {
		for i := range value {
			if num >= value[i] {
				result += symbol[i]
				num -= value[i]
				break
			}
		}
	}
	return result
}

// Link: https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/
func allStringPermutations(str string) []string {
	index := 0
	ans := make([]string, 0)
	strByte := []byte(str)
	allStringPermutationsHelper(strByte, &ans, index)
	return ans
}

func allStringPermutationsHelper(str []byte, ans *[]string, index int) {
	//base condition
	if index == len(str) {
		*ans = append(*ans, string(str))
		// fmt.Println("str : ", string(str))
		return
	}

	for i := index; i < len(str); i++ {
		str[index], str[i] = str[i], str[index]
		allStringPermutationsHelper(str, ans, index+1)
		str[index], str[i] = str[i], str[index]
	}
}

// Link: https://www.geeksforgeeks.org/number-substrings-string/
func subString(str string) int {
	ans := 0
	index := 0
	strByte := []byte(str)
	subStringHelper(strByte, index, "", &ans)
	return ans

}

func subStringHelper(str []byte, index int, ans string, count *int) {

	//base condition
	if index == len(str) {
		fmt.Println("ans : ", ans)
		*count = *count + 1
		return
	}

	//include
	subStringHelper(str, index+1, ans+string(str[index]), count)

	//exclude
	subStringHelper(str, index+1, ans, count)
}

// Link: https://leetcode.com/problems/sum-of-beauty-of-all-substrings/
func beautySum(s string) int {
	min, max := 1, 1
	charMap := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		charMap[s[i]] = charMap[s[i]] + 1
		if count, exist := charMap[s[i]]; exist {
			fmt.Printf("  %v ===> %v  \n", s[i], charMap[s[i]])
			if count > max {
				max = count
			}
			if count < min {
				min = count
			}
		}
	}
	fmt.Printf("min : %v  max: %v \n", min, max)
	return max - min
}

// Link: https://leetcode.com/problems/count-and-say/
func countAndSay(n int) string {

	var ans string

	for i := 0; i < n; i++ {
		ans = countAndSayHelper(ans)
	}
	return ans
}

func countAndSayHelper(prevAns string) string {
	if prevAns == "" {
		return "1"
	}
	var ans strings.Builder
	count := 1
	for i := 1; i < len(prevAns); i++ {
		if prevAns[i] == prevAns[i-1] {
			count++
		} else {
			ans.WriteString(strconv.Itoa(count))
			ans.WriteByte(prevAns[i-1])
			count = 1
		}
	}
	ans.WriteString(strconv.Itoa(count))
	ans.WriteByte(prevAns[len(prevAns)-1])
	return ans.String()
}

// Link: https://leetcode.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {

	var (
		rem        = 0
		counter    = make(map[byte]int)
		ret        = string(make([]byte, len(s)))
		start, end = 0, 0
	)

	for i := range t {
		rem++
		counter[t[i]]++
	}
	if rem > len(s) {
		return ""
	}

	for end < len(s) {
		if v, ok := counter[s[end]]; ok {
			if v > 0 {
				rem--
			}
			counter[s[end]]--
		}
		for rem <= 0 {
			if len(ret) >= len(s[start:end+1]) {
				ret = s[start : end+1]
			}
			if _, ok := counter[s[start]]; ok {
				counter[s[start]]++
				if counter[s[start]] > 0 {
					rem++
				}
			}
			start++
		}
		end++
	}

	if ret == string(make([]byte, len(s))) {
		return ""
	}
	return ret
}

// Link: https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
func arrayStringsAreEqual(word1 []string, word2 []string) bool {

	completeWord1 := ""
	completeWord2 := ""
	if len(word1) > 1 {
		completeWord1 = concat(word1)
	} else {
		completeWord1 = word1[0]
	}

	if len(word2) > 1 {
		completeWord2 = concat(word2)
	} else {
		completeWord2 = word2[0]
	}
	return completeWord1 == completeWord2
}

func concat(arr []string) string {
	res := ""
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

// Link: https://leetcode.com/problems/group-anagrams/
type CharCount [26]int

func groupAnagrams(strs []string) [][]string {
	hash := make(map[CharCount][]string)

	for _, word := range strs {
		charCount := CharCount{}
		for i := 0; i < len(word); i++ {
			charCount[word[i]-'a']++
		}

		hash[charCount] = append(hash[charCount], word)
	}

	var result [][]string
	for _, value := range hash {
		result = append(result, value)
	}
	return result
}

// Link: https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
func longestPalindrome(words []string) int {
	n := len(words)
	count := make(map[string]int, n)
	length := 0

	// for every word that has its corresponding reversed word, it will definitely be included in the final palindrome length
	for _, word := range words {
		reversed := reverse(word)
		if v, ok := count[reversed]; ok {
			length += 4
			if v == 1 {
				delete(count, reversed)
			} else {
				count[reversed]--
			}
		} else {
			count[word]++
		}
	}

	// among unused words, if there is a word with the same letters, we can add it to the final palindrome length
	for word, _ := range count {
		if isPair(word) {
			length += 2
			break
		}
	}
	return length
}

func reverse(s string) string {
	ss := make([]byte, 2)
	ss[0], ss[1] = s[1], s[0]
	return string(ss)
}

func isPair(s string) bool {
	return s[0] == s[1]
}

// Link: https://leetcode.com/problems/reverse-vowels-of-a-string/
func reverseVowels(s string) string {
	r := []byte(s)
	var index []int

	for i, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' ||
			v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
			index = append(index, i)
		}
	}

	if len(index) >= 2 {
		i := 0
		j := len(index) - 1
		for i < j {
			r[index[i]], r[index[j]] = r[index[j]], r[index[i]]
			i++
			j--
		}
	}
	return string(r)
}

// Link: https://leetcode.com/problems/orderly-queue/
func orderlyQueue(s string, k int) string {
	ans := ""
	sB := []byte(s)
	if k == 1 {
		sBmin := make([]byte, len(sB))
		copy(sBmin, sB)
		// when k is equal to 1, we run through every possibility, looking for the order that has the smallest lexigraphical value.
		for range s {
			if bytes.Compare(sB, sBmin) < 0 {
				copy(sBmin, sB)
			}
			sB = append(sB[1:], sB[0])
		}
		ans = string(sBmin)
	} else {
		sort.Slice(sB, func(a, b int) bool {
			return sB[a] < sB[b]
		})
		ans = string(sB)
	}
	return ans
}

// Link: https://leetcode.com/problems/make-the-string-great/
func makeGood(s string) string {

	if len(s) <= 1 {
		return s
	}

	ans := make([]byte, 0)
	//push first elment in stack
	ans = append(ans, s[0])

	for i := 1; i < len(s); i++ {
		var topElement byte
		if len(ans) == 0 {
			topElement = 0
		} else {
			topElement = ans[len(ans)-1]
		}

		//pop the top element if we found same elment in stack as of s string otherwise append current element
		if topElement == s[i]-32 || s[i] == topElement-32 {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

// Link: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {

	if len(s) <= 1 {
		return s
	}

	ans := make([]byte, 0)

	//push first elment in stack
	ans = append(ans, s[0])
	for i := 1; i < len(s); i++ {
		var topElement byte
		if len(ans) == 0 {
			topElement = 0
		} else {
			topElement = ans[len(ans)-1]
		}
		//pop the top element if we found same elment in stack as of s string otherwise append current element
		if topElement == s[i] {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

// Link: https://leetcode.com/problems/determine-if-string-halves-are-alike/
func halvesAreAlike(s string) bool {
	//s = "book"
	n := len(s)
	vowelMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	var firstCount, secondCount int

	for i, v := range s {
		if i < n/2 {
			_, found := vowelMap[v]
			if found {
				firstCount++
			}
		} else {
			_, found := vowelMap[v]
			if found {
				secondCount++
			}
		}
	}
	return firstCount == secondCount
}

/*
Question 1657: https://leetcode.com/problems/determine-if-two-strings-are-close/

Idea:
- Count occurence of each char in word1 and word2
- 2 words are close if for any character 'a1' in word1 with occurence x there will be a character
a2 in word2 with the same occurence.

Optimization:
- Checks the length of word1, word2.
- While creating the bucket for word2 we can also do check to see if there is any character that
only exist in word2.
**/

func closeStrings(word1 string, word2 string) bool {

	if len(word1) != len(word2) {
		return false
	}

	bucket1 := make([]int, 26)
	bucket2 := make([]int, 26)
	for _, char := range word1 {
		bucket1[int(char)-int('a')]++
	}

	for _, char := range word2 {
		if bucket1[int(char)-int('a')] == 0 {
			return false
		}
		bucket2[int(char)-int('a')]++
	}

	sort.Ints(bucket1)
	sort.Ints(bucket2)
	for i := range bucket1 {
		if bucket1[i] != bucket2[i] {
			return false
		}
	}

	return true
}

// Link: https://leetcode.com/problems/sort-characters-by-frequency/
type runeFreq struct {
	r   rune
	cnt int
}

func frequencySort(s string) string {

	freq := make([]runeFreq, 128)
	for _, r := range s {
		freq[r-'0'].cnt++
		freq[r-'0'].r = r
	}

	sort.Slice(freq, func(i, j int) bool {
		if freq[i].cnt == freq[j].cnt {
			return freq[i].r < freq[j].r
		}
		return freq[i].cnt > freq[j].cnt
	})
	sb := strings.Builder{}
	for _, fr := range freq {
		sb.WriteString(strings.Repeat(string(fr.r), fr.cnt))
	}

	return sb.String()
}
