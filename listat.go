package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	iterator := l
	for i := 0; i < pos; i++ {
		if iterator.Next != nil {
			// fmt.Print("i= ", i, "  ")
			// fmt.Print("value ", iterator.Data, "\n")
			iterator = iterator.Next
			// fmt.Print("i=", i, "  ")
		} else {
			return nil
		}
	}
	// fmt.Print("pos=", pos, " ")
	return iterator
}
