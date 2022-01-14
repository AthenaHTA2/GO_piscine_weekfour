package piscine

func ListSize(l *List) int {
	counter := 1
	if l.Head == nil {
		counter = 0
	} else {
		counter = 1
		iterator := l.Head
		for ; iterator.Next != nil; iterator = iterator.Next {
			counter = counter + 1
		}
	}
	return counter
}
