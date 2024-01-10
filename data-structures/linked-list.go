package main

import "fmt"

func main() {
	fmt.Println("Hello")

	ll := new(LinkedList)

	ll.Insert("A")
	ll.Insert("B")

	ll.Display()
}

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	node *Node
	size int
}

func (ll *LinkedList) Insert(val string) {
	if ll.node == nil {
		ll.node = &Node{
			value: val,
		}
	} else {
		node := ll.node

		for node != nil {
			if node.next == nil {
				node.next = &Node{
					value: val,
				}

				break
			}

			node = node.next
		}
	}
}

func (ll *LinkedList) Display() {
	node := ll.node

	for node != nil {
		fmt.Println("NODE:", node.value)

		node = node.next
	}
}
