package main

import "github.com/01-edu/z01"

func main() {
	Array := "abcdefghijklmnopqrstuvwxyz\n"
	for i := 0; i <= 26; i++ {
		z01.PrintRune(rune(Array[i]))
	}
}
