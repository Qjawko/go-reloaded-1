package student

// NodeL q
type NodeL struct {
	Data interface{}
	Next *NodeL
}

// List q
type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListRemoveIf q
func ListRemoveIf(l *List, dataRef interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.Data == dataRef {
		l.Head = l.Head.Next
		ListRemoveIf(l, dataRef)
		return
	}

	iterator := l.Head
	for iterator.Next != nil {

		if iterator.Next.Data == dataRef {
			iterator.Next = iterator.Next.Next
			ListRemoveIf(l, dataRef)
			return
		}

		iterator = iterator.Next
	}
}
