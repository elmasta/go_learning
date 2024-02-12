package piscine

func ListSize(l *List) int {
	counter := 0
	i := l.Head
	if l.Head != nil {
		for ; i.Next != nil; i = i.Next {
			counter++
		}
		counter++
	}
	return counter
}
