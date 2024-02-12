package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
	} else if l2.Head != nil {
		i := l1.Head
		for ; i.Next != nil; i = i.Next {
		}
		i.Next = l2.Head
	}
}
