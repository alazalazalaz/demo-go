package main

import (
	"fmt"
)

func main(){
	nodeDemo := ListNodeFB()
	nodeDemo2 := ListNodeFB()
	_outputList(reverseList(nodeDemo))
	_outputList(reverseListV2(nodeDemo2))
}

//使用数组
func reverseList(head *ListNodeB) *ListNodeB {
	calHead := head
	var stackList []*ListNodeB
	for ; calHead != nil; {
		stackList = append(stackList, calHead)
		calHead = calHead.Next
	}
	headNode := &ListNodeB{
		Val:0,
		Next:nil,
	}

	for i:= 0; i<len(stackList) ; i++{
		if i==0{
			stackList[i].Next = nil
			headNode = stackList[i]
		}else{
			stackList[i].Next = headNode
			headNode = stackList[i]
		}
	}
	return headNode
}

//空间O(1) 时间O(n)
func reverseListV2( pHead *ListNodeB ) *ListNodeB {
	temp := &ListNodeB{
		Val :0,
		Next: nil,
	}
	curr := &ListNodeB{
		Val :0,
		Next: nil,
	}
	nextNode := &ListNodeB{
		Val :0,
		Next: nil,
	}
	temp, curr, nextNode = nil, nil, nil
	for pHead != nil{
		nextNode = pHead.Next
		temp = curr
		curr = pHead
		curr.Next = temp
		pHead = nextNode
	}
	return curr
}

type ListNodeB struct{
	Val int
	Next *ListNodeB
}

func ListNodeFB() *ListNodeB {
	headNode := &ListNodeB{
		Val: 1,
		Next: nil,
	}

	firstNode := new(ListNodeB)
	firstNode.Val = 2

	secondNode := new(ListNodeB)
	secondNode.Val = 3

	thirdNode := new(ListNodeB)
	thirdNode.Val = 4

	headNode.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	return headNode
}

func _outputList(node *ListNodeB){
	fmt.Printf("输出节点：")
	for node != nil{
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Printf("节点输出完毕\r\n")
}