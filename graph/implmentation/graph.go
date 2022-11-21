package main

import "fmt"

/******Graph Implementataion*******/
type Hnode struct {
	v      int
	h_prev *Hnode
	h_next *Hnode
}

type Vnode struct {
	v            int
	ph_head_node *Hnode
	v_prev       *Vnode
	v_next       *Vnode
}

type Graph struct {
	pv_head_node *Vnode
	nr_vertices  int
	nr_edges     int
}

type Edge struct {
	v_start int
	v_end   int
}

func show_graph(g *Graph, msg string) {
	pv_run := &Vnode{}
	ph_run := &Hnode{}

	fmt.Println(msg)
	for pv_run = g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run.v_next {
		fmt.Printf("[%v]  <->  ", pv_run.v)
		for ph_run = pv_run.ph_head_node.h_next; ph_run != pv_run.ph_head_node; ph_run = ph_run.h_next {
			fmt.Printf("[%v]  <->  ", ph_run.v)
		}
		fmt.Printf("[END]\n")
	}
	fmt.Println("-------------------------------------------------------------")
}

func create_graph() *Graph {
	g := &Graph{}
	g.pv_head_node = v_create_list()
	g.nr_vertices = 0
	g.nr_edges = 0
	return g
}

func add_vertex(g *Graph, new_vertex int) string {
	p := &Vnode{}

	p = v_search_node(g.pv_head_node, new_vertex)
	if p != nil {
		return "vertex already exist"
	}

	v_insert_end(g.pv_head_node, new_vertex)
	return "Added the vertex successfully"
}

func remove_vertex(g *Graph, vertex int) string {
	pv_delete := &Vnode{}
	pv_h_in_vlist := &Vnode{}

	ph_run := &Hnode{}
	ph_run_next := &Hnode{}
	ph_delete_node_in_adj_list := &Hnode{}

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

	return "Removed the vertex successfully"
}

func add_edge(g *Graph, v_start int, v_end int) string {

	pv_start := &Vnode{}
	pv_end := &Vnode{}
	ph_start_in_end := &Hnode{}
	ph_end_in_start := &Hnode{}

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

	return "Added edge successfully"

}

func destroy_graph(g *Graph) string {
	pv_run := &Vnode{}
	pv_run_next := &Vnode{}

	for pv_run = g.pv_head_node.v_next; pv_run != g.pv_head_node; pv_run = pv_run_next {
		pv_run_next = pv_run.v_next
		remove_vertex(g, pv_run.v)
	}
	return "deleted the graph succesfully"
}

func remove_edge(g *Graph, v_start, v_end int) string {

	pv_start := &Vnode{}
	pv_end := &Vnode{}
	ph_end_in_start := &Hnode{}
	ph_start_in_end := &Hnode{}

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

//Graph helper routines -> Vertical List Management -> Vertical List Interface

func v_create_list() *Vnode {
	pv_head_node := &Vnode{}
	pv_head_node.ph_head_node = nil
	pv_head_node.v_prev = pv_head_node
	pv_head_node.v_next = pv_head_node

	return pv_head_node
}

func v_search_node(v_list *Vnode, search_vertex int) *Vnode {
	pv_run := &Vnode{}
	for pv_run = v_list.v_next; pv_run != v_list; pv_run = pv_run.v_next {
		if pv_run.v == search_vertex {
			return pv_run
		}
	}
	return nil
}

func v_insert_end(pv_list *Vnode, new_vertex int) {
	v_generic_insert(pv_list.v_prev, v_get_list_node(new_vertex), pv_list)
}

func v_generic_insert(pv_beg *Vnode, pv_mid *Vnode, pv_end *Vnode) {
	pv_mid.v_next = pv_end
	pv_mid.v_prev = pv_beg
	pv_beg.v_next = pv_mid
	pv_end.v_prev = pv_mid
}

func v_get_list_node(new_vertex int) *Vnode {

	pv_new_node := &Vnode{}
	pv_new_node.v = new_vertex
	pv_new_node.ph_head_node = h_create_list()
	pv_new_node.v_prev = nil
	pv_new_node.v_next = nil

	return pv_new_node
}

func v_generic_delete(pv_delete_node *Vnode) {
	pv_delete_node.v_next.v_prev = pv_delete_node.v_prev
	pv_delete_node.v_prev.v_next = pv_delete_node.v_next
}

//Graph helper routines -> Horizontal List Management -> Horizontal list interface

func h_create_list() *Hnode {

	ph_head_node := &Hnode{}

	ph_head_node = h_get_list_node(-1)
	ph_head_node.h_next = ph_head_node
	ph_head_node.h_prev = ph_head_node

	return ph_head_node
}

func h_get_list_node(new_vertex int) *Hnode {
	ph_new_node := &Hnode{}

	ph_new_node.v = new_vertex
	ph_new_node.h_next = nil
	ph_new_node.h_prev = nil

	return ph_new_node
}

func h_search_node(ph_list *Hnode, search_vertex int) *Hnode {
	ph_run := &Hnode{}

	for ph_run = ph_list.h_next; ph_run != ph_list; ph_run = ph_run.h_next {
		if ph_run.v == search_vertex {
			return ph_run
		}
	}
	return nil
}

func h_insert_end(ph_list *Hnode, new_vertex int) {
	h_generic_insert(ph_list.h_prev, h_get_list_node(new_vertex), ph_list)
}

func h_generic_insert(ph_beg, ph_mid, ph_end *Hnode) {
	ph_mid.h_next = ph_end
	ph_mid.h_prev = ph_beg
	ph_beg.h_next = ph_mid
	ph_end.h_prev = ph_mid
}

func h_generic_delete(ph_delete_node *Hnode) {
	ph_delete_node.h_next.h_prev = ph_delete_node.h_prev
	ph_delete_node.h_prev.h_next = ph_delete_node.h_next
}
