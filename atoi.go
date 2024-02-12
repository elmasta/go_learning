package piscine

func Atoi(s string) int {
	/* - = 45 , + = 43 */
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
