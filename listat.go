package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	position := 0
	if l != nil {
		for ; l.Next != nil; l = l.Next {
			if position == pos {
				return l
			}
			position++
		}
		if position == pos {
			return l
		}
	}
	return nil
}
