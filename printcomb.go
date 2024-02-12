package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	/* 48 à 57 */
	l := [3]int{48, 48, 48}
	for i := 0; i <= 1000; i++ {
		if l[0] < l[1] && l[1] < l[2] && l[0] < l[2] {
			for _, s := range l {
				z01.PrintRune(rune(s))
			}
			if l[0] < 55 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
		if l[2]+1 == 58 {
			l[2] = 48
			l[1] = l[1] + 1
			if l[1]+1 == 58 {
				l[1] = 48
				l[0] = l[0] + 1
			}
		} else {
			l[2] = l[2] + 1
		}
	}
}
