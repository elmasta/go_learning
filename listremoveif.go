package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	i := l.Head
	var previous *NodeL
	for i != nil {
		if i.Data == data_ref {
			if i == l.Head {
				l.Head = i.Next
			} else {
				previous.Next = i.Next
			}
		} else {
			previous = i
		}
		i = i.Next
	}
}
