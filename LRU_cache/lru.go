package main

import "container/list"

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	if node, exist := c.m[key]; exist {
		val := node.Value.(*list.Element).Value.(Pair).value
		//moved node to front
		c.l.MoveToFront(node)
		return val
	}
	return -1
}

func (c *LRUCache) Put(key, val int) {

	//if node exist
	if node, exist := c.m[key]; exist {
		//move the node to front
		c.l.MoveToFront(node)

		node.Value.(*list.Element).Value = Pair{key, val}
	} else {
		//if node does not exist
		//delete the node of list is full
		if c.l.Len() == c.cap {
			//get the key you want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			//delete node from map
			delete(c.m, idx)
			//remove the last node
			c.l.Remove(c.l.Back())
		}
		//initialize node
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: val,
			},
		}
		//push the new node to list
		ptr := c.l.PushFront(node)
		//add new node to map
		c.m[key] = ptr
	}

}
