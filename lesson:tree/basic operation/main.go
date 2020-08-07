package main

import "fmt"

type TreeNode struct {
	Val interface{}
	Left *TreeNode
	Right *TreeNode
}

func NewTree(value interface{}) *TreeNode {
	return &TreeNode{
		Val: value,
		Left: nil,
		Right: nil,
	}
}

func (node *TreeNode)addLeftNode(value interface{}) {
	node.Left = &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

func (node *TreeNode)addRightNode(value interface{}) {
	node.Right = &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

func main() {
	root := NewTree("a")
	root.addLeftNode(123)
	root.addRightNode(true)

	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Printf("%v",root.Right.Val)


}
