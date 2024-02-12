package piscine

func Capitalize(s string) string {
	sR := []rune(s)
	for i, v := range s {
		if i == 0 && (v >= 'a' && v <= 'z') {
			sR[i] = v - 32
		} else if i > 0 {
			if v >= 'a' && v <= 'z' {
				if sR[i-1] < '0' || (sR[i-1] > '9' && sR[i-1] < 'A') || (sR[i-1] > 'Z' && sR[i-1] < 'a') || sR[i-1] > 'z' {
					sR[i] = v - 32
				}
			} else if v >= 'A' && v <= 'Z' {
				if (sR[i-1] >= '0' && sR[i-1] <= '9') || (sR[i-1] >= 'a' && sR[i-1] <= 'z') || (sR[i-1] >= 'A' && sR[i-1] <= 'Z') {
					sR[i] = v + 32
				}
			}
		}
	}
	return string(sR)
}
