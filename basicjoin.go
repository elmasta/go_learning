package piscine

func BasicJoin(elems []string) string {
	toReturn := ""
	for _, v := range elems {
		toReturn += v
	}
	return toReturn
}
