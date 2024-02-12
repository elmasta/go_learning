package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	rA := []string{}
	for i := len(os.Args); i > 1; i-- {
		if i > 0 {
			rA = append(rA, os.Args[i-1])
		}
	}
	for _, v := range rA {
		v := []rune(v)
		for _, value := range v {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
