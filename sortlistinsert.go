package student

// NodeI linked list node
type NodeI struct {
	Data int
	Next *NodeI
}

// SortListInsert inserts dataRef in the linked list l while keeping the list sorted in ascending order.
func SortListInsert(l *NodeI, dataRef int) *NodeI {

	// if no list
	if l == nil {
		return &NodeI{Data: dataRef}
	}

	// if dataRef is less than first item
	if l.Data > dataRef {
		return &NodeI{Next: l, Data: dataRef}
	}

	for iterator := l; iterator.Next != nil; iterator = iterator.Next {
		if iterator.Next.Data > dataRef {
			newItem := &NodeI{Next: iterator.Next, Data: dataRef}
			iterator.Next = newItem
			break
		}
	}

	return l
}
