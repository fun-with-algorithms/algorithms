package linkedlist

func Reverse(list *LinkedList) {
	if list == nil || list.Head == nil {
		return
	}

	var p *Node = nil
	var n *Node = nil
	c := list.Head

	for c != nil {
		n = c.Next
		c.Next = p
		p = c
		c = n
	}

	list.Head = p
}
