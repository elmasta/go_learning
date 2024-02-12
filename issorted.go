package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	toReturn := true
	if len(a) > 1 {
		if a[0] > a[1] {
			for i := 0; i < len(a); i++ {
				if i+1 < len(a) {
					result := f(a[i+1], a[i])
					if result > 0 {
						toReturn = false
					}
				}
			}
		} else {
			for i := 0; i < len(a); i++ {
				if i+1 < len(a) {
					result := f(a[i], a[i+1])
					if result > 0 {
						toReturn = false
					}
				}
			}
		}
	}
	return toReturn
}
