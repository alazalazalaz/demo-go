package main

import "fmt"

//func main(){
//	nodeList := TreeNode1()
//	fmt.Println(isValidBST(nodeList))
//}


var valList []int

func isValidBST(root *TreeNode) bool {
	_centerOrder(root)
	fmt.Println(root, valList)
	if len(valList) <= 1{
		return true
	}
	for i:= 0; i<len(valList) - 1; i++{
		if valList[i+1] <= valList[i]{
			return false
		}
	}
	return true
}

func _centerOrder(root *TreeNode){
	if root == nil {
		return
	}
	_centerOrder(root.Left)
	valList = append(valList, root.Val)
	_centerOrder(root.Right)
	return
}

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func TreeNode1() *TreeNode {
	headNode := &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}

	secondNode := new(TreeNode)
	secondNode.Val = 2

	thirdNode := new(TreeNode)
	thirdNode.Val = 2
	//fourNode := new(TreeNode)
	//fourNode.Val = 3
	//fiveNode := new(TreeNode)
	//fiveNode.Val = 6

	headNode.Left = secondNode
	headNode.Right = thirdNode
	//thirdNode.Left = fourNode
	//thirdNode.Right = fiveNode

	return headNode
}