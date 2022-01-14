package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "How are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
