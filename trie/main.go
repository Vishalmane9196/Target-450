package main

import "fmt"

func main() {

	/*********Trie Implementation*******/
	// triObj := Constructor()

	// words := []string{"aqua", "jack", "card", "care"}
	// for _, word := range words {
	// 	triObj.Insert(word)
	// }

	// words_Search := []string{"aqua", "jack", "card", "care", "cat", "dog", "can"}
	// for _, word := range words_Search {
	// 	found := triObj.Search(word)
	// 	fmt.Printf("Search [%s] : %v \n", word, found)
	// }

	// fmt.Println("StartsWith `aq` : ", triObj.StartsWith("aq"))

	/*******Problem*********/
	// var board1 [][]byte
	// board := [][]string{{"o", "a", "a", "n"}, {"e", "t", "a", "e"}, {"i", "h", "k", "r"}, {"i", "f", "l", "v"}}

	// for i, _ := range board {
	// 	temp := make([]byte, 0)
	// 	for j, _ := range board[i] {
	// 		temp = append(temp, []byte(board[i][j])...)
	// 	}
	// 	board1 = append(board1, temp)
	// }
	// words1 := []string{"oath", "pea", "eat", "rain"}
	// res := findWords(board1, words1)
	// fmt.Println("findWords : ", res)

	// strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	res := longestCommonPrefix(strs)
	fmt.Println("longestCommonPrefix : ", res)

}

/*
Links:
https://www.educative.io/answers/how-to-create-a-trie-in-go


*/
