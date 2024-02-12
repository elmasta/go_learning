package piscine

func Compare(a, b string) int {
	sA := []byte(a)
	sB := []byte(b)
	if len(sA) > len(sB) {
		for i := 0; i < len(sA)-1; i++ {
			if len(sB)-1 != i {
				if sA[i] > sB[i] {
					return 1
				} else if sA[i] < sB[i] {
					return -1
				}
			} else {
				return 1
			}
		}
	} else {
		if a == b {
			return 0
		} else {
			for i := 0; i <= len(sB); i++ {
				if len(sA) != i {
					if sA[i] > sB[i] {
						return 1
					} else if sA[i] < sB[i] {
						return -1
					}
				} else {
					return -1
				}
			}
		}
	}
	return 0
}
