package main

import "fmt"

// /*
// color code
// 	1 white
// 	2 grey
// 	3 black
// */

type hnode struct {
	v      int
	h_prev *hnode
	h_next *hnode
}

type vnode struct {
	v            int
	color        int
	ph_head_node *hnode
	v_prev       *vnode
	v_next       *vnode
}

type graph struct {
	pv_head_node *vnode
	nr_vertices  int
	nr_edges     int
}

type edge struct {
	v_start int
	v_end   int
}

type queue_node struct {
	pv_node *vnode
	q_prev  *queue_node
	q_next  *queue_node
}

func main() {

	g := &graph{}
	var V = []int{1, 2, 3, 4, 5, 6}
	var E = []edge{{1, 2}, {1, 6}, {2, 6}, {2, 5}, {2, 3}, {3, 4}, {3, 6}, {3, 5}, {4, 5}, {5, 6}}

	g = create_graph()

	//add vertex
	for i := 0; i < len(V); i++ {
		add_vertex(g, V[i])
	}

	//add edges
	for i := 0; i < len(E); i++ {
		add_edge(g, E[i].v_start, E[i].v_end)
	}

	show_graph(g, "Intial graph")

	fmt.Println("After DFS")
	dfs(g)

	//TODO:
	// fmt.Println("After BFS")
	// bfs(g, 3)

	destroy_graph(g)
}

func show_graph(g *graph, msg string) {
	pv_run := &vnode{}
	ph_run := &hnode{}

	fmt.Println(msg)

	for pv_run = g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run.v_next {
		fmt.Printf("[%v]\t<->\t", pv_run.v)
		for ph_run = pv_run.ph_head_node.h_next; ph_run != pv_run.ph_head_node; ph_run = ph_run.h_next {
			fmt.Printf("[%v]\t<->\t", ph_run.v)
		}
		fmt.Printf("[END]\n")
	}
	fmt.Println("-------------------------------------------------------------")

}

func create_graph() *graph {
	g := &graph{}
	g.pv_head_node = v_create_list()
	g.nr_vertices = 0
	g.nr_edges = 0
	return g
}

func add_vertex(g *graph, new_vertex int) string {
	p := &vnode{}

	p = v_search_node(g.pv_head_node, new_vertex)
	if p != nil {
		return "vertex already exist"
	}

	v_insert_end(g.pv_head_node, new_vertex)
	return "Added the vertex successfully"
}

func remove_vertex(g *graph, vertex int) string {
	pv_delete := &vnode{}
	pv_h_in_vlist := &vnode{}

	ph_run := &hnode{}
	ph_run_next := &hnode{}
	ph_delete_node_in_adj_list := &hnode{}

	pv_delete = v_search_node(g.pv_head_node, vertex)
	if pv_delete == nil {
		return "invalid vertex"
	}

	ph_run = pv_delete.ph_head_node.h_next

	for ph_run != pv_delete.ph_head_node {
		ph_run_next = ph_run.h_next

		pv_h_in_vlist = v_search_node(g.pv_head_node, ph_run.v)
		ph_delete_node_in_adj_list = h_search_node(pv_h_in_vlist.ph_head_node, vertex)
		h_generic_delete(ph_delete_node_in_adj_list)
		g.nr_edges -= 1

		ph_run = ph_run_next
	}

	v_generic_delete(pv_delete)
	g.nr_vertices -= 1

	return "removed the vertex succesfully"
}

func add_edge(g *graph, v_start int, v_end int) string {

	pv_start := &vnode{}
	pv_end := &vnode{}
	ph_start_in_end := &hnode{}
	ph_end_in_start := &hnode{}

	pv_start = v_search_node(g.pv_head_node, v_start)
	if pv_start == nil {
		return "invalid vertex"
	}
	pv_end = v_search_node(g.pv_head_node, v_end)
	if pv_end == nil {
		return "invalid vetex"
	}
	ph_start_in_end = h_search_node(pv_end.ph_head_node, v_start)
	if ph_start_in_end != nil {
		return "edge exist"
	}
	ph_end_in_start = h_search_node(pv_start.ph_head_node, v_end)
	if ph_end_in_start != nil {
		return "edge exist"
	}
	h_insert_end(pv_start.ph_head_node, v_end)
	h_insert_end(pv_end.ph_head_node, v_start)

	g.nr_edges += 1

	return "succefully added edge"

}

func destroy_graph(g *graph) string {
	pv_run := &vnode{}
	pv_run_next := &vnode{}

	for pv_run = g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run_next {
		pv_run_next = pv_run.v_next
		remove_vertex(g, pv_run.v)
	}
	return "deleted the graph succesfully"
}

func remove_edge(g *graph, v_start, v_end int) string {

	pv_start := &vnode{}
	pv_end := &vnode{}
	ph_end_in_start := &hnode{}
	ph_start_in_end := &hnode{}

	pv_start = v_search_node(g.pv_head_node, v_start)
	if pv_start == nil {
		return "invalid index"
	}

	pv_end = v_search_node(g.pv_head_node, v_end)
	if pv_end == nil {
		return "invalid index"
	}

	ph_end_in_start = h_search_node(pv_start.ph_head_node, v_end)
	if ph_end_in_start == nil {
		return "invalid edge"
	}
	ph_start_in_end = h_search_node(pv_end.ph_head_node, v_start)
	if ph_start_in_end == nil {
		return "invalid edge"
	}
	h_generic_delete(ph_end_in_start)
	h_generic_delete(ph_start_in_end)

	g.nr_edges -= 1

	return "successulyy removed the edge"
}

/*
graph helper routines -> Vertical List Management -> Vertical List Interface
*/

func v_create_list() *vnode {

	pv_head_node := &vnode{}
	pv_head_node.ph_head_node = nil
	pv_head_node.v_prev = pv_head_node
	pv_head_node.v_next = pv_head_node

	return pv_head_node
}

func v_search_node(v_list *vnode, search_vertex int) *vnode {
	pv_run := &vnode{}
	for pv_run = v_list.v_next; pv_run != v_list; pv_run = pv_run.v_next {
		if pv_run.v == search_vertex {
			return pv_run
		}
	}
	return nil
}

func v_insert_end(pv_list *vnode, new_vertex int) {
	v_generic_insert(pv_list.v_prev, v_get_list_node(new_vertex), pv_list)
}

func v_generic_insert(pv_beg *vnode, pv_mid *vnode, pv_end *vnode) {
	pv_mid.v_next = pv_end
	pv_mid.v_prev = pv_beg
	pv_beg.v_next = pv_mid
	pv_end.v_prev = pv_mid
}

func v_get_list_node(new_vertex int) *vnode {

	pv_new_node := &vnode{}
	pv_new_node.v = new_vertex
	pv_new_node.ph_head_node = h_create_list()
	pv_new_node.v_prev = nil
	pv_new_node.v_next = nil

	return pv_new_node
}

func v_generic_delete(pv_delete_node *vnode) {
	pv_delete_node.v_next.v_prev = pv_delete_node.v_prev
	pv_delete_node.v_prev.v_next = pv_delete_node.v_next
}

/* graph helper routines -> Horizontal List Management -> Horizontal list interface*/

func h_create_list() *hnode {

	ph_head_node := &hnode{}

	ph_head_node = h_get_list_node(-1)
	ph_head_node.h_next = ph_head_node
	ph_head_node.h_prev = ph_head_node

	return ph_head_node
}

func h_get_list_node(new_vertex int) *hnode {

	ph_new_node := &hnode{}

	ph_new_node.v = new_vertex
	ph_new_node.h_next = nil
	ph_new_node.h_prev = nil

	return ph_new_node
}

func h_search_node(ph_list *hnode, search_vertex int) *hnode {
	ph_run := &hnode{}

	for ph_run = ph_list.h_next; ph_run != ph_list; ph_run = ph_run.h_next {
		if ph_run.v == search_vertex {
			return ph_run
		}
	}
	return nil
}

func h_insert_end(ph_list *hnode, new_vertex int) {
	h_generic_insert(ph_list.h_prev, h_get_list_node(new_vertex), ph_list)
}

func h_generic_insert(ph_beg, ph_mid, ph_end *hnode) {
	ph_mid.h_next = ph_end
	ph_mid.h_prev = ph_beg
	ph_beg.h_next = ph_mid
	ph_end.h_prev = ph_mid
}

func h_generic_delete(ph_delete_node *hnode) {
	ph_delete_node.h_next.h_prev = ph_delete_node.h_prev
	ph_delete_node.h_prev.h_next = ph_delete_node.h_next

}

func reset(g *graph) {
	for pv_run := g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run.v_next {
		pv_run.color = 1
	}
}

func dfs(g *graph) {

	reset(g)
	fmt.Printf("[START]  <->  ")

	for pv_run := g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run.v_next {

		if pv_run.color == 1 {
			dfs_visit(g, pv_run)
		}
	}
	fmt.Printf("[END]")
}

func dfs_visit(g *graph, pv *vnode) {
	ph_run := &hnode{}
	pv_h_in_vlist := &vnode{}

	pv.color = 2

	fmt.Printf("[%v]  <-> ", pv.v)

	for ph_run = pv.ph_head_node.h_next; ph_run != pv.ph_head_node; ph_run = ph_run.h_next {

		pv_h_in_vlist = v_search_node(g.pv_head_node, ph_run.v)

		if pv_h_in_vlist.color == 1 {
			dfs_visit(g, pv_h_in_vlist)
		}
	}

	pv.color = 3
}

func create_prio_queue() *queue_node {

	var p_new_queue = &queue_node{}
	p_new_queue = q_get_node(nil)
	p_new_queue.q_prev = p_new_queue
	p_new_queue.q_next = p_new_queue
	return p_new_queue
}

func q_get_node(pv_node *vnode) *queue_node {
	var p_new_node = &queue_node{}
	p_new_node.pv_node = pv_node
	p_new_node.q_prev = nil
	p_new_node.q_next = nil
	return p_new_node
}

func prio_enqueue(prio_queue *queue_node, pv_node *vnode) {
	//insert at end
	// fmt.Println("arr data : ", pv_node.v)
	q_generic_insert(prio_queue.q_prev, q_get_node(pv_node), prio_queue)
}

func q_generic_insert(p_beg, p_mid, p_end *queue_node) {
	p_mid.q_prev = p_beg
	p_mid.q_next = p_end
	p_beg.q_next = p_mid
	p_end.q_next = p_mid
}

func q_generic_delete(p_node *queue_node) {
	p_node.q_prev.q_next = p_node.q_next
	p_node.q_next.q_prev = p_node.q_prev
	// p_node.v_prev.v_next =p_node.v_next
	// p_node.v_next.v_prev = p_node.v_prev
}
func generic_delete(p_node *vnode) {
	p_node.v_prev.v_next = p_node.v_next
	p_node.v_next.v_prev = p_node.v_prev
}

func is_prio_queue_empty(prio_queue *queue_node) bool {
	return (prio_queue.q_prev == prio_queue && prio_queue.q_next == prio_queue)
}

func prio_dequeue_min(prio_queue *queue_node, data **vnode) {

	if is_prio_queue_empty(prio_queue) == true {
		fmt.Println("prio queue is empty")
	}

	// var current_min_data int = prio_queue.v_next.v
	// var p_min_data_node *vnode = prio_queue.v_next
	var current_min_data int = prio_queue.q_next.pv_node.v
	var p_min_data_node *vnode = prio_queue.q_next.pv_node

	for p_run := prio_queue.q_next; p_run != prio_queue; p_run = p_run.q_next {

		if p_run.pv_node.v < current_min_data {
			current_min_data = p_run.pv_node.v
			p_min_data_node = p_run.pv_node
		}
	}
	*data = p_min_data_node
	generic_delete(p_min_data_node)
}

func destroy_prio_queue(prio_queue *queue_node) {
	//del at  beg
	for p_run := prio_queue.q_next; p_run != prio_queue; p_run = p_run.q_next {
		q_generic_delete(p_run)
	}
}

func show_queue(prio_queue *vnode) {

	fmt.Printf("[BEG]<->")
	for p_run := prio_queue.v_next; p_run != prio_queue; p_run = p_run.v_next {
		fmt.Printf("[%v]<->", p_run.v)
	}
	fmt.Printf("[END]\n")
}

func bfs(g *graph, v int) {

	pv_node := v_search_node(g.pv_head_node, v)
	if pv_node == nil {
		fmt.Println("Invalid vertex")
	}
	reset(g)
	pv_node.color = 2
	p_queue := create_prio_queue()
	prio_enqueue(p_queue, pv_node)
	fmt.Printf("[START] <-> ")

	for is_prio_queue_empty(p_queue) != true {
		pv := &vnode{}
		prio_dequeue_min(p_queue, &pv)
		fmt.Printf("[%v]<->", pv.v)

		for ph_run := pv.ph_head_node.h_next; ph_run != pv.ph_head_node; ph_run = ph_run.h_next {
			pv_h_in_vlist := v_search_node(g.pv_head_node, ph_run.v)

			if pv_h_in_vlist.color == 1 {
				pv_h_in_vlist.color = 2
				prio_enqueue(p_queue, pv_h_in_vlist)
			}
		}
		pv.color = 3
	}
	fmt.Printf("[END]\n")
	destroy_prio_queue(p_queue)
}
