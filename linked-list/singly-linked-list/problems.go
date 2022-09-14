package main

import "fmt"

func (list *LinkedList) middleOfLinkedList() {
    if list.headNode == nil || list.headNode.nextNode == nil {
		fmt.Println("middle : ", list.headNode)
        return 
    }
    
    if list.headNode.nextNode.nextNode == nil {
		fmt.Println("middle : ", list.headNode.nextNode)
        return 
    }
    
    var slow *Node = list.headNode
    var fast *Node = list.headNode.nextNode
    
    for fast != nil {
        fast = fast.nextNode
        if fast != nil {
            fast = fast.nextNode
        }
        slow = slow.nextNode
    }
	fmt.Println("middle : ", slow)
}

func (list *LinkedList) detectCycle() {

	if list.headNode == nil || list.headNode.nextNode == nil {
		fmt.Println("there is no cycle in linked list")
        return
    }
    
    // floyd algo for cycle detection
    var slow *Node = list.headNode
    var fast *Node = list.headNode
    
    for slow != nil && fast != nil {
        fast = fast.nextNode
        
        if fast != nil {
            fast = fast.nextNode
        }
        slow = slow.nextNode
        
        if slow == fast {
			fmt.Println("there is cycle in linked list")
            return 
        }
    }
	fmt.Println("there is no cycle in linked list")
    return 
}