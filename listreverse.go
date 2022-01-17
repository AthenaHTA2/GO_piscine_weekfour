package piscine

func ListReverse(l *List) {
	currentNode := l.Head
	var next *NodeL
	var previousNode *NodeL
	l.Head = l.Tail

	for currentNode != nil {
		next, currentNode.Next = currentNode.Next, previousNode
		previousNode, currentNode = currentNode, next
	}
	l.Head = previousNode
}
