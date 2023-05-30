package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func reverseList(list *LinkedList) {
	var previous *Node
	current := list.Head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	list.Head = previous
}

func printList(list *LinkedList) {
	current := list.Head

	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}

	fmt.Println()
}

func main() {
	list := LinkedList{
		Head: &Node{
			Value: 1,
			Next: &Node{
				Value: 2,
				Next: &Node{
					Value: 3,
					Next:  nil,
				},
			},
		},
	}

	printList(&list)
	reverseList(&list)
	printList(&list)
}
