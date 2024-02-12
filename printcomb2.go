package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	/* 48 à 57 */
	l := [4]int{48, 48, 48, 48}
	for i := 0; i <= 9999; i++ {
		if l[0]*10+l[1] < l[2]*10+l[3] {
			for i, s := range l {
				z01.PrintRune(rune(s))
				if i == 1 {
					z01.PrintRune(' ')
				}
			}
			if l[0]+l[1] == 113 && l[0] == 57 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		if l[2] == 57 && l[3] == 57 {
			if l[1]+1 == 58 {
				l[1] = 48
				l[0] = l[0] + 1
			} else {
				l[1] = l[1] + 1
			}
			l[2] = 48
			l[3] = 48
		} else {
			if l[3]+1 == 58 {
				l[3] = 48
				l[2] = l[2] + 1
			} else {
				l[3] = l[3] + 1
			}
		}
	}
}
