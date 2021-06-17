package tree

import (
	"fmt"
)

func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.Data)
	PreOrder(root.Left)
	PreOrder(root.Right)
}

func InOrder(root *Node) {
	if root == nil {
		return
	}

	InOrder(root.Left)
	fmt.Println(root.Data)
	InOrder(root.Right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Println(root.Data)
}
