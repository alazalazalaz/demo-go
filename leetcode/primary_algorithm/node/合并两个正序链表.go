package main

//func main(){
//	l1 := ListNode1Make()
//	l2 := ListNode1Make2()
//	re := mergeTwoLists(l1, l2)
//	fmt.Println(re)
//}

func mergeTwoLists(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	if l1 == nil || l2 == nil {
		return nil
	}
	dummy1 := new(ListNode1)
	dummyBase := dummy1
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val{
			dummy1.Next = l1
			dummy1 = l1
			l1 = l1.Next
		}else{
			dummy1.Next = l2
			dummy1 = l2
			l2 = l2.Next
		}
	}
	if l1 == nil {
		dummy1.Next = l2
	}
	if l2 == nil {
		dummy1.Next = l1
	}
	return dummyBase.Next
}

type ListNode1 struct{
	Val int
	Next *ListNode1
}

func ListNode1Make() *ListNode1 {
	headNode := &ListNode1{
		Val: 0,
		Next: nil,
	}

	firstNode := new(ListNode1)
	firstNode.Val = 1

	secondNode := new(ListNode1)
	secondNode.Val = 2

	thirdNode := new(ListNode1)
	thirdNode.Val = 4

	headNode.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	return headNode
}

func ListNode1Make2() *ListNode1 {
	headNode := &ListNode1{
		Val: 0,
		Next: nil,
	}

	firstNode := new(ListNode1)
	firstNode.Val = 1

	secondNode := new(ListNode1)
	secondNode.Val = 3

	thirdNode := new(ListNode1)
	thirdNode.Val = 5

	headNode.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	return headNode
}
