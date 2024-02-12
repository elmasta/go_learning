package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	i := l.Head
	for i != nil {
		if comp(ref, i.Data) {
			return &i.Data
		}
		i = i.Next
	}
	return nil
}
