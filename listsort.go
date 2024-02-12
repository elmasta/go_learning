package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func Sorter(table []int) []int {
	i := 0
	for true {
		if len(table) > 1 {
			if i == len(table)-2 {
				if table[i] > table[i+1] {
					table = SwapInTable(table, i)
					i = 0
				} else {
					break
				}
			} else if table[i] > table[i+1] {
				table = SwapInTable(table, i)
				i = 0
			} else {
				i++
			}
		} else {
			break
		}
	}
	return table
}

func listPushBackNodeI(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var lInt []int
	i := l
	for ; i.Next != nil; i = i.Next {
		lInt = append(lInt, i.Data)
	}
	lInt = append(lInt, i.Data)
	lInt = Sorter(lInt)
	var link *NodeI
	for _, v := range lInt {
		link = listPushBackNodeI(link, v)
	}
	return link
}
