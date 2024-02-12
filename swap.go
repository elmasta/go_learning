package piscine

func Swap(a *int, b *int) {
	aVal := *a
	*a = *b
	*b = aVal
}
