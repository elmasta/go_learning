package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{}
	n.Data = data
	if l.Head == nil {
		n.Next = nil
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
