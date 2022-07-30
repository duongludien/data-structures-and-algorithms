package sorting

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
