package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// MorrisInorder returns inorder without recursion or a stack: it threads a
// temporary link from each left subtree's rightmost node back to its root,
// walks it, and removes the thread when it returns.
func MorrisInorder(root *TreeNode) []int {
	out := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			out = append(out, cur.Val)
			cur = cur.Right
			continue
		}
		// Find the inorder predecessor: rightmost node of the left subtree.
		pred := cur.Left
		for pred.Right != nil && pred.Right != cur {
			pred = pred.Right
		}
		if pred.Right == nil {
			pred.Right = cur // thread: remember the way back
			cur = cur.Left
		} else {
			pred.Right = nil // second arrival: left side done, unthread
			out = append(out, cur.Val)
			cur = cur.Right
		}
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	fmt.Println(MorrisInorder(root))
}
