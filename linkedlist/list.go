package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Data int64
	Next *Node
}

func (n *Node) GetNext() *Node {
	if n == nil {
		return nil
	}

	return n.Next
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}

	return "{ " + strconv.FormatInt(n.Data, 10) + ", " + n.Next.String() + " }"
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Print() {
	if ll.Head == nil {
		return
	}

	curr := ll.Head
	var sb strings.Builder

	for curr != nil {
		fmt.Fprintf(&sb, "%d->", curr.Data)
		curr = curr.Next
	}

	fmt.Printf("List: %s\n", sb.String()[:len(sb.String())-2])
}
