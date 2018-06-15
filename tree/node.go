package tree

import (
"fmt"
)

type TreeNode struct {
	Value int
	Left,Right *TreeNode
}
//factory func
func CreateNode(Value int) *TreeNode{
	return &TreeNode{Value:Value}
}
func (node TreeNode)Print(){
	fmt.Println(node.Value)
}

func(node *TreeNode) SetValue(value int){
	node.Value=value
}

func (node *TreeNode) Traverse(){
	if node==nil{
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

