package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsIntStr(r []rune) bool {
	toReturn := false
	var temp rune
	for i, v := range r {
		if i == 0 {
			if v >= '0' && v <= '9' && len(r) == 1 {
				toReturn = true
			} else if v >= '1' && v <= '2' && i == 0 {
				temp = v
			} else {
				toReturn = false
				break
			}
		} else {
			if temp == '1' {
				if v >= '0' && v <= '9' {
					toReturn = true
				} else {
					toReturn = false
				}
			} else {
				if v >= '0' && v <= '6' {
					toReturn = true
				} else {
					toReturn = false
				}
			}
		}
	}
	return toReturn
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--upper" {
			for i := 2; i < len(os.Args); i++ {
				nmbTemp := 0
				aR := []rune(os.Args[i])
				if len(aR) < 3 && IsIntStr(aR) {
					if len(aR) == 2 {
						if aR[0] == '1' {
							nmbTemp += 10
						} else {
							nmbTemp += 20
						}
						nmbTemp = nmbTemp + int(aR[1]) - 48
					} else {
						nmbTemp = nmbTemp + int(aR[0]) - 48
					}
					z01.PrintRune(rune(nmbTemp + 64))
				} else {
					z01.PrintRune(' ')
				}
			}
		} else {
			for i := 1; i < len(os.Args); i++ {
				nmbTemp := 0
				aR := []rune(os.Args[i])
				if len(aR) < 3 && IsIntStr(aR) {
					if len(aR) == 2 {
						if aR[0] == '1' {
							nmbTemp += 10
						} else {
							nmbTemp += 20
						}
						nmbTemp = nmbTemp + int(aR[1]) - 48
					} else {
						nmbTemp = nmbTemp + int(aR[0]) - 48
					}
					z01.PrintRune(rune(nmbTemp + 96))
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
