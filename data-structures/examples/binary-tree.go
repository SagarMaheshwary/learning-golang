package examples

import (
	"fmt"
	"slices"
)

func BinaryTree() {
	tree := Tree{Root: TreeNode{Val: 4}}

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

	// fmt.Println("-----------")
	// bfsPrint(tree.Root)

	// arr := []int{4, 2, 7, 1, 3, 6, 9, 11, 5, 7, 8, 6, 3}

	// for i := 1; i < len(arr); i += 2 {

	// 	fmt.Println("LEFT NODE", arr[i])

	// 	if !slices.Contains(arr, i+1) {
	// 		break
	// 	}

	// 	fmt.Println("RIGHT NODE", arr[i+1])
	// 	fmt.Println("")
	// }

	// fmt.Println(tree.Root.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root TreeNode
}

func (t *Tree) insert(val int, root *TreeNode) {
	if root.Val < val {
		t.insertRight(val, root)
	} else {
		t.insertLeft(val, root)
	}
}

func (t *Tree) insertRight(val int, node *TreeNode) {
	if node.Right != nil {
		t.insert(val, node.Right)
	} else {
		node.Right = &TreeNode{Val: val}
	}
}

func (t *Tree) insertLeft(val int, node *TreeNode) {
	if node.Left != nil {
		t.insert(val, node.Left)
	} else {
		node.Left = &TreeNode{Val: val}
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	nodes := []*TreeNode{root}

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

func bfsPrint(root TreeNode) {
	nodes := []TreeNode{root}

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

func dfsPrint(root TreeNode) {
	nodes := []TreeNode{root}

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

func countNodes(root TreeNode) int {
	nodes := []TreeNode{root}

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
