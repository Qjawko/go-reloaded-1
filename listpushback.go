package student

// ListPushBack push back
func ListPushBack(l *List, data interface{}) {
	new := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = new
		l.Tail = new
		return
	}

	l.Tail.Next = new
	l.Tail = new
}
