package student

// SortedListMerge mege 2 and kepp sorted
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return n2
	}

	if n2 == nil {
		return n1
	}

	result := ListPushBack(nil, n2.Data)

	for iterator := n2.Next; iterator != nil; iterator = iterator.Next {
		result = SortListInsert(result, iterator.Data)
	}

	for iterator := n1; iterator != nil; iterator = iterator.Next {
		result = SortListInsert(result, iterator.Data)
	}

	return result
}
