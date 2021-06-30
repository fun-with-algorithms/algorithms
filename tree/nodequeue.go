package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Queue []*Node

func (q *Queue) Insert(element *Node) {
	*q = append(*q, element)
}

func (q *Queue) Remove() *Node {
	element := (*q)[0]
	*q = (*q)[1:]

	return element
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Print() {
	sb := strings.Builder{}

	for _, elem := range *q {
		_, _ = sb.WriteString(strconv.FormatInt(int64(elem.Data), 10))
		_, _ = sb.WriteString(", ")
	}

	str := sb.String()

	if sb.Len() != 0 {
		str = sb.String()[:(sb.Len() - 2)]
	}

	fmt.Println(str)
}
