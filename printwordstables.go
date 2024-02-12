package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, v := range a {
		for i, s := range v {
			z01.PrintRune(s)
			if i+1 == len(v) {
				z01.PrintRune('\n')
			}
		}
	}
}
