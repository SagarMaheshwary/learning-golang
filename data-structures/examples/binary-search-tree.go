package examples

import (
	"fmt"
	"slices"
)

func BinarySearchTree() {
	tree := BSTTree{Root: BSTNode{Val: 4}}

	tree.insert(2, &tree.Root)
	tree.insert(7, &tree.Root)
	tree.insert(1, &tree.Root)
	tree.insert(3, &tree.Root)
	tree.insert(6, &tree.Root)
	tree.insert(9, &tree.Root)

	fmt.Println(countNodes(tree.Root))
	dfsPrint(tree.Root)
	bfsPrint(tree.Root)
	invertTree(&tree.Root)
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

type BSTTree struct {
	Root BSTNode
}

func (t *BSTTree) insert(val int, root *BSTNode) {
	if root.Val < val {
		t.insertRight(val, root)
	} else {
		t.insertLeft(val, root)
	}
}

func (t *BSTTree) insertRight(val int, node *BSTNode) {
	if node.Right != nil {
		t.insert(val, node.Right)
	} else {
		node.Right = &BSTNode{Val: val}
	}
}

func (t *BSTTree) insertLeft(val int, node *BSTNode) {
	if node.Left != nil {
		t.insert(val, node.Left)
	} else {
		node.Left = &BSTNode{Val: val}
	}
}

func invertTree(root *BSTNode) *BSTNode {
	if root == nil {
		return root
	}

	nodes := []*BSTNode{root}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = slices.Delete(nodes, 0, 1)

		temp := node.Left
		node.Left = node.Right
		node.Right = temp

		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}

		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
	}

	return root
}

func bfsPrint(root BSTNode) {
	nodes := []BSTNode{root}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = slices.Delete(nodes, 0, 1)

		fmt.Println("CURRENT NODE", node.Val)

		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}

		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
	}

}

func dfsPrint(root BSTNode) {
	nodes := []BSTNode{root}

	for len(nodes) > 0 {
		nodesCount := len(nodes)
		node := nodes[nodesCount-1]

		fmt.Println("CURRENT NODE", node.Val)

		nodes = slices.Delete(nodes, nodesCount-1, nodesCount)

		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}

		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
	}
}

func countNodes(root BSTNode) int {
	nodes := []BSTNode{root}

	visited := map[int]bool{}
	count := 0

	for len(nodes) > 0 {
		nodesCount := len(nodes)
		node := nodes[len(nodes)-1]
		nodes = slices.Delete(nodes, nodesCount-1, nodesCount)

		if !visited[node.Val] {
			visited[node.Val] = true
			count++
		}

		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}

		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
	}

	return count
}
