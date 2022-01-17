package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if pos < 0 {
		return nil
	}
	iterator := l
	for i := 0; i < pos; i++ {
		if iterator.Next != nil {
			iterator = iterator.Next
		} else {
			return nil
		}
	}
	return iterator
}
