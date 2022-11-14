package main

import "fmt"

func main() {

	/******Graph Implementataion*******/
	// g := &Graph{}
	// var V = []int{1, 2, 3, 4, 5, 6}
	// var E = []Edge{{1, 2}, {1, 6}, {2, 6}, {2, 5}, {2, 3}, {3, 4}, {3, 6}, {3, 5}, {4, 5}, {5, 6}}

	// g = create_graph()

	// //add vertex
	// for i := 0; i < len(V); i++ {
	// 	add_vertex(g, V[i])
	// }

	// //add edges
	// for i := 0; i < len(E); i++ {
	// 	add_edge(g, E[i].v_start, E[i].v_end)
	// }

	// //Initial state
	// show_graph(g, "Initial state")

	// remove_edge(g, 2, 5)
	// remove_edge(g, 2, 3)
	// remove_edge(g, 6, 2)
	// remove_edge(g, 6, 5)

	// show_graph(g, "After removing edges, (2, 5), (2, 3), (6, 2), (6, 5):")

	// destroy_graph(g)
	// show_graph(g, "After destroying graph")

	/********Problems********/
	stones := [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}
	res := removeStones(stones)
	fmt.Println("res : ", res)

}
