package piscine

func IndexS(s string, toFind string) []int {
	var toReturn []int
	if len(toFind) > 0 && len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				found := false
				tempI := i
				for ind := 0; ind < len(toFind); ind++ {
					if i < len(s) {
						if s[i] != toFind[ind] {
							found = false
							break
						} else {
							found = true
						}
						i++
					}
				}
				if found {
					toReturn = append(toReturn, i-len(toFind))
				}
				i = tempI
			}
		}
		return toReturn
	} else {
		return toReturn
	}
}

func Split(s, sep string) []string {
	var toReturn []string
	if len(sep) == 0 {
		toReturn = append(toReturn, s)
	} else {
		temp := IndexS(s, sep)
		if len(temp) > 0 {
			for i, v := range temp {
				if i != 0 && i+1 != len(temp) {
					if v+len(sep) != temp[i+1] {
						toReturn = append(toReturn, s[v+len(sep):temp[i+1]])
					}
				} else if i == 0 {
					if v > 0 {
						toReturn = append(toReturn, s[:v])
					}
					toReturn = append(toReturn, s[v+len(sep):temp[i+1]])
				} else if v+len(sep) != len(s) {
					toReturn = append(toReturn, s[v+len(sep):])
				}
			}
		}
	}
	var final []string
	for _, v := range toReturn {
		if len(v) > 0 {
			final = append(final, v)
		}
	}
	return final
}
