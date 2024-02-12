package piscine

func TrimAtoi(s string) int {
	/* - = 45 , + = 43 */
	nmb := 0
	var list []int
	minus := false
	/* parcours string */
	for _, x := range s {
		/* check si nmb */
		if int(x) >= 48 && int(x) <= 57 {
			/* si 0 après premier nombre rencontré ou si nombre > 0 */
			if (int(x) == 48 && len(list) > 0) || int(x) > 48 {
				list = append(list, int(x))
			}
		} else if int(x) == 45 && len(list) == 0 {
			minus = true
		}
	}
	/* list en int */
	/* puissance */
	p := 1
	for i := 1; i < len(list); i++ {
		p = p * 10
	}
	/* ajout dans nmb */
	for i := 0; i < len(list); i++ {
		if p == 0 {
			nmb = list[i] - 48
		} else {
			nmb = nmb + p*(list[i]-48)
		}
		p = p / 10
	}
	/* si negatif */
	if minus == true {
		nmb = nmb * -1
	}
	return nmb
}
