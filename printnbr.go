package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var l []int
	if n < 0 {
		z01.PrintRune('-')
	}
	countN := n
	for true {
		if n > 0 {
			if countN/10 > 0 {
				l = append(l, 48+countN%10)
				countN = countN / 10
			} else {
				l = append(l, 48+countN%10)
				break
			}
		} else {
			if countN/10 < 0 {
				l = append(l, 48+countN%10*-1)
				countN = countN / 10
			} else {
				l = append(l, 48+countN%10*-1)
				break
			}
		}
	}
	for i := len(l) - 1; i >= 0; i-- {
		z01.PrintRune(rune(l[i]))
	}
}
