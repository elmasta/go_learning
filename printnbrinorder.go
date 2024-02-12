package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var l []int
	if n == 0 {
		z01.PrintRune('0')
	} else {
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
			}
		}
		/* sorter */
		i := 0
		for true {
			if len(l) > 1 {
				if i == len(l)-2 {
					if l[i] < l[i+1] {
						l = SwapInTable(l, i)
						i = 0
					} else {
						break
					}
				} else if l[i] < l[i+1] {
					l = SwapInTable(l, i)
					i = 0
				} else {
					i++
				}
			} else {
				break
			}
		}
		/* end sorter */
		for i := len(l) - 1; i >= 0; i-- {
			z01.PrintRune(rune(l[i]))
		}
	}
}
