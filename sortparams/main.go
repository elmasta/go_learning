package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := len(os.Args)
	for i := 0; i < a; i++ {
		for j := 1; j < a-1; j++ {
			if os.Args[j] > os.Args[j+1] {
				temp := os.Args[j]
				os.Args[j] = os.Args[j+1]
				os.Args[j+1] = temp
			}
		}
	}
	for i := 1; i < len(os.Args); i++ {
		r := []rune(os.Args[i])
		for x := 0; x < len(r); x++ {
			z01.PrintRune(r[x])
		}
		z01.PrintRune('\n')
	}
}
