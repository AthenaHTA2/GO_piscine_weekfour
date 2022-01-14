package piscine

import (
	"fmt"
)

func ListAt(l *NodeL, pos int) *NodeL{
	counter := 1
	iterator := l.Head
	for ; iterator.Next != nil; iterator = iterator.Next {
		counter = counter + 1
		if counter == pos {
			fmt.Println(iterator.Data)
		}
}