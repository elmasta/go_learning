package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return nil
	}
	var lInt []int
	i := l
	for ; i.Next != nil; i = i.Next {
		lInt = append(lInt, i.Data)
	}
	lInt = append(lInt, i.Data)
	lInt = append(lInt, data_ref)
	lInt = Sorter(lInt)
	var link *NodeI
	for _, v := range lInt {
		link = listPushBackNodeI(link, v)
	}
	return link
}
