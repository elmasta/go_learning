package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	}
	if n1 != nil {
		i := n1
		for ; i.Next != nil; i = i.Next {
		}
		i.Next = n2
		return ListSort(n1)
	}
	return ListSort(n2)
}
