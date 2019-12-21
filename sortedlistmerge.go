package student

// SortedListMerge mege 2 and kepp sorted
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	copy := n1

	for iterator := n2; iterator.Next != nil; iterator = iterator.Next {
		SortListInsert(copy, iterator.Data)
	}

	return copy
}
