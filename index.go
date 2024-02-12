package piscine

func Index(s string, toFind string) int {
	toReturn := -1
	if len(toFind) > 0 && len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				if toReturn < 0 {
					toReturn = i
				}
				for ind := 0; ind < len(toFind); ind++ {
					if s[i] != toFind[ind] {
						toReturn = -1
						break
					}
					if len(s) == i+1 && len(s) > 1 {
						toReturn = -1
						break
					}
					i++
				}
			}
			if toReturn > -1 || i == len(s) {
				break
			}
		}
		return toReturn
	} else if len(toFind) == 0 && len(s) > 0 {
		return 0
	} else {
		return -1
	}
}
