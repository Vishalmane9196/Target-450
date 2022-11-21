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

//Link: https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
type State struct {
	row, col, length int
}

func nearestExit(maze [][]byte, entrance []int) int {

	//to store the whole maze element state
	dataMap := make(map[State]struct{})

	rowLength, colLength := len(maze), len(maze[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	//var declaration
	var poppedState, pState State
	var currentRow, currentCol int

	//queued entrace elemenet to queue
	queue := []State{{entrance[0], entrance[1], 0}}
	for len(queue) > 0 {
		//dequeued the first elment of queue
		poppedState, queue = queue[0], queue[1:]

		pState = State{poppedState.row, poppedState.col, 0}
		if _, ok := dataMap[pState]; ok {
			continue
		}
		dataMap[pState] = struct{}{}

		//check the current element is bundary element or not
		if (poppedState.row == 0 || poppedState.row == rowLength-1 || poppedState.col == 0 || poppedState.col == colLength-1) && (poppedState.row != entrance[0] || poppedState.col != entrance[1]) {
			return poppedState.length
		}

		//traverse over all the directions of popped element
		for _, direction := range directions {
			currentRow = poppedState.row + direction[0]
			currentCol = poppedState.col + direction[1]
			if currentRow >= 0 && currentRow < rowLength && currentCol >= 0 && currentCol < colLength && maze[currentRow][currentCol] == '.' {
				queue = append(queue, State{currentRow, currentCol, poppedState.length + 1})
			}
		}

	}
	return -1
}
