package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, v := range os.Args {
		if i > 0 {
			rA := []rune(v)
			for index, value := range rA {
				if (i == 0 && index > 1) || i > 0 {
					z01.PrintRune(value)
				}
			}
			z01.PrintRune('\n')
		}
	}
}
