package main

import "github.com/01-edu/z01"

func main() {
	Array := "0123456789\n"
	for i := 0; i <= 10; i++ {
		z01.PrintRune(rune(Array[i]))
	}
}
