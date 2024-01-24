package examples

import "fmt"

func LinkedList() {
	ll := List{}

	ll.insert("A")
	ll.insert("B")

	ll.display()
}

type Node struct {
	value string
	next  *Node
}

type List struct {
	node *Node
}

func (ll *List) insert(val string) {
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

func (ll *List) display() {
	node := ll.node

	for node != nil {
		fmt.Println("NODE:", node.value)

		node = node.next
	}
}
