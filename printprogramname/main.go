package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rA := []rune(os.Args[0])
	for i, v := range rA {
		if i > 1 {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
