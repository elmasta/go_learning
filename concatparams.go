package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return "\n"
	} else {
		toReturn := ""
		for i, v := range args {
			toReturn += v
			if i < len(args)-1 {
				toReturn += "\n"
			}
		}
		return toReturn
	}
}
