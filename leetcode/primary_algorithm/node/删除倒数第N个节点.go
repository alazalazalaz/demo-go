package main

import "fmt"

//func main(){
//	head := listNode()
//	result := removeNthFromEnd(head, 1)
//	displayNode(result)
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//计算长度
	fmt.Println(head)
	nodeLen := lenNode(head)
	dummyNode := &ListNode{
		Val: 0,
		Next: head,
	}
	tempHead := dummyNode
	for i:= 0; i<nodeLen - n; i++{
		tempHead = tempHead.Next
	}
	tempHead.Next = tempHead.Next.Next

	return dummyNode.Next
}

func lenNode(head *ListNode) int {
	i:= 0
	for ; head != nil; {
		i++
		head = head.Next
	}
	return i
}

type ListNode struct{
	Val int
	Next *ListNode
}

func listNode() *ListNode {
	headNode := &ListNode{
		Val: 0,
		Next: nil,
	}

	firstNode := new(ListNode)
	firstNode.Val = 111

	secondNode := new(ListNode)
	secondNode.Val = 222

	thirdNode := new(ListNode)
	thirdNode.Val = 333

	headNode.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	return headNode
}

func displayNode(no *ListNode){
	for ; no != nil; {
		fmt.Println(no)
		no = no.Next
	}
}
