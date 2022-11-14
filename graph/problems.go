package main

//Link: https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
func removeStones(stones [][]int) int {

	rows := make(map[int][]int)
	cols := make(map[int][]int)

	for i := 0; i < len(stones); i++ {
		rows[stones[i][0]] = append(rows[stones[i][0]], i)
		cols[stones[i][1]] = append(cols[stones[i][1]], i)
	}

	visited := make([]bool, len(stones))
	var groups int
	for i := 0; i < len(stones); i++ {
		if !visited[i] {
			visited[i] = true
			groups++
			removeStonesHelper(rows, cols, visited, stones, i)
		}
	}
	return len(stones) - groups
}

func removeStonesHelper(rows, cols map[int][]int, visited []bool, stones [][]int, curr int) {

	for _, rowNeighbor := range rows[stones[curr][0]] {
		if !visited[rowNeighbor] {
			visited[rowNeighbor] = true
			removeStonesHelper(rows, cols, visited, stones, rowNeighbor)
		}
	}

	for _, columnNeighbor := range cols[stones[curr][1]] {
		if !visited[columnNeighbor] {
			visited[columnNeighbor] = true
			removeStonesHelper(rows, cols, visited, stones, columnNeighbor)
		}
	}
}
