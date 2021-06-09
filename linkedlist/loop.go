package linkedlist

func DetectLoop(list *LinkedList) bool {
	if list == nil || list.Head == nil || list.Head.GetNext() == nil {
		return false
	}

	hare := list.Head
	tortoise := list.Head

	for hare != nil {
		hare = hare.GetNext().GetNext()
		tortoise = tortoise.GetNext()

		if hare == tortoise {
			return true
		}
	}

	return false
}
