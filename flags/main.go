package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}

func SwapInTable(table []rune, index int) []rune {
	a, b := table[index-1], table[index]
	table[index-1], table[index] = b, a
	return table
}

func main() {
	if len(os.Args) == 1 {
		printHelp()
	} else {
		var temp string
		var tempTwo string
		toOrder := false
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			printHelp()
		} else {
			for i, v := range os.Args {
				if i > 0 && len(v) > 0 {
					if v == "-o" || v == "--order" {
						toOrder = true
					} else if len(v) == 1 {
						temp = temp + v
					} else if len(v) < 9 {
						if v[0:2] == "-i" {
							tempTwo = tempTwo + v[3:]
						} else {
							temp = temp + v
						}
					} else {
						if v[0:8] == "--insert" {
							tempTwo = tempTwo + v[9:]
						} else if v[0:2] == "-i" {
							tempTwo = tempTwo + v[3:]
						} else {
							temp = temp + v
						}
					}
				}
			}
			temp = temp + tempTwo
			if toOrder && len(temp) > 1 {
				tR := []rune(temp)
				for i := 1; i < len(tR); i++ {
					if tR[i] < tR[i-1] {
						tR = SwapInTable(tR, i)
						i = 0
					}
				}
				temp = string(tR)
			}
			fmt.Println(temp)
		}
	}
}
