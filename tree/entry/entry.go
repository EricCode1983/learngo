package main

import "learngo/tree"

func main() {

	var root tree.TreeNode

	root=tree.TreeNode{Value:3}
	root.Left=&tree.TreeNode{}
	root.Right=&tree.TreeNode{5,nil,nil}
	root.Right.Left=new(tree.TreeNode)
	//nodes:=[]treeNode{
	//	{value:3},{},{},
	//	{6,nil,&root},
	//}
	root.Left.Right=tree.CreateNode(2)

	println()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	root.Traverse()
}

