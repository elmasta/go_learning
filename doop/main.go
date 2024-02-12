package main

import (
	"os"
)

func Atoi(s string) int {
	nmb := 0
	plusCount := 0
	minusCount := 0
	for i, x := range s {
		if int(x) >= 48 && int(x) <= 57 {
			if int(x) != 48 || nmb > 0 {
				p := 10
				for d := len(s) - i - 2; d > 0; d-- {
					p = p * 10
				}
				if i+1 == len(s) {
					nmb = nmb + int(x) - 48
				} else {
					nmb = nmb + (int(x)-48)*p
				}
			}
		}
		if int(x) == 43 {
			if nmb > 0 {
				nmb = 0
				break
			} else {
				plusCount = plusCount + 1
			}
		}
		if int(x) == 45 {
			if nmb > 0 {
				nmb = 0
				break
			} else {
				minusCount = minusCount + 1
			}
		}
		if int(x) < 48 || int(x) > 57 {
			if int(x) != 43 {
				if int(x) != 45 {
					nmb = 0
					break
				}
			}
		}
		if plusCount > 1 || minusCount > 1 {
			nmb = 0
			break
		} else if plusCount == 1 && minusCount == 1 {
			nmb = 0
			break
		}
	}
	if nmb > 0 && minusCount == 1 {
		nmb = nmb * -1
	}
	return nmb
}

func check(str string) bool {
	isOk := false
	for i, v := range str {
		if (i == 0 && rune(v) == '-') || (rune(v) >= '0' && rune(v) <= '9') {
			isOk = true
		} else {
			isOk = false
		}
	}
	return isOk
}

func checkSign(str string) bool {
	isOk := false
	signs := []rune{'+', '-', '*', '/', '%'}
	for _, v := range signs {
		if len(str) == 1 && rune(str[0]) == v {
			isOk = true
			break
		} else {
			isOk = false
		}
	}
	return isOk
}

func PrintNumbr(n int) {
	var l []rune
	var toPrint string
	countN := n
	for true {
		if n > 0 {
			if countN/10 > 0 {
				l = append(l, rune(48+countN%10))
				countN = countN / 10
			} else {
				l = append(l, rune(48+countN%10))
				break
			}
		} else {
			if countN/10 < 0 {
				l = append(l, rune(48+countN%10*-1))
				countN = countN / 10
			} else {
				l = append(l, rune(48+countN%10*-1))
				break
			}
		}
	}
	if n < 0 {
		l = append(l, '-')
	}
	for i := len(l) - 1; i >= 0; i-- {
		toPrint += string(l[i])
	}
	toPrint += string("\n")
	os.Stdout.Write([]byte(toPrint))
}

func Operation(a int, b int, c rune) {
	var totalNum int
	if c == '+' {
		totalNum = a + b
		PrintNumbr(totalNum)
	} else if c == '-' {
		totalNum = a - b
		PrintNumbr(totalNum)
	} else if c == '*' {
		totalNum = a * b
		PrintNumbr(totalNum)
	} else if c == '/' {
		if b != 0 {
			totalNum = a / b
			PrintNumbr(totalNum)
		} else {
			os.Stdout.Write([]byte("No division by 0\n"))
		}
	} else {
		if b != 0 {
			totalNum = a % b
			PrintNumbr(totalNum)
		} else {
			os.Stdout.Write([]byte("No modulo by 0\n"))
		}
	}
}

func main() {
	if len(os.Args) == 4 {
		numbOne := 0
		checkOne := false
		numbTwo := 0
		checkTwo := false
		var operand rune
		checkThree := false
		for i, v := range os.Args {
			if i > 0 {
				if i == 1 && check(v) && len(v) < 19 {
					numbOne = Atoi(v)
					checkOne = true
				} else if i == 3 && check(v) && len(v) < 19 {
					numbTwo = Atoi(v)
					checkThree = true
				} else if i == 2 && checkSign(v) && len(v) == 1 {
					operand = rune(v[0])
					checkTwo = true
				}
			}
		}
		if checkOne && checkTwo && checkThree {
			Operation(numbOne, numbTwo, operand)
		}
	}
}
