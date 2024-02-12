package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{}
	n.Data = data
	n.Next = nil
	if l.Head == nil {
		l.Head = n
	} else {
		i := l.Head
		for ; i.Next != nil; i = i.Next {
		}
		i.Next = n
	}
}
