package tree

import "fmt"

func LevelOrder(root *Node) {
	q := &Queue{}

	q.Insert(root)

	for !q.IsEmpty() {
		elem := q.Remove()

		if elem.Left != nil {
			q.Insert(elem.Left)
		}

		if elem.Right != nil {
			q.Insert(elem.Right)
		}

		fmt.Println(elem.Data)
	}
}
