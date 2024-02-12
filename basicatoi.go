package piscine

func BasicAtoi(s string) int {
	nmb := 0
	for i, x := range s {
		if int(x)-48 != 0 || nmb > 0 {
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
	return nmb
}
