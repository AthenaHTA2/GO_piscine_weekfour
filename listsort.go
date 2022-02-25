package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	count := 0
	var first *NodeI
	if l == nil || l.Next == nil {
		return l
	}
	for l != nil {
		next := l.Next
		if count == 0 {
			first = l
			count++
		}
		for next != nil {
			if l.Data > next.Data {
				l.Data, next.Data = next.Data, l.Data
			}
			next = next.Next
		}
		l = l.Next
	}
	return first
}
