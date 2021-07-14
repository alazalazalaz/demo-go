package main

import "fmt"

/**
链表
 */
type Node struct{
	Data string
	Next *Node
}

func main(){
	//链表for循环
	node := listNode()
	for ; node != nil; {
		fmt.Println(node)
		node = node.Next
		//output
		//&{head 0xc00000c0a0}
		//&{first 0xc00000c0c0}
		//&{second 0xc00000c0e0}
		//&{third <nil>}
	}
	fmt.Println(node)//nil 因为上面for循环最后是把thirdNode的next赋值给了node

	//节点复制后，指向的都是同一个链表哦，比如a赋值给b，清空链表b后，a也被清空了。
	clearNode()

	//删除倒数第N个节点
}

//在进行数组的插入、删除操作时，为了保持内存数据的连续性，需要做大量的数据搬移，所以速度较慢。
//而在链表中插入或者删除一个数据，我们并不需要为了保持内存的连续性而搬移结点，因为链表的存储空间本身就不是连续的。所以，在链表中插入和删除一个数据是非常快速的。
func clearNode(){
	fmt.Println("add node=========")
	node := listNode()
	testNode := node
	fmt.Println(node, testNode)
	testNode.Next = nil
	fmt.Println(node, testNode)

}

func listNode() *Node {
	headNode := &Node{
		Data: "head",
		Next: nil,
	}

	firstNode := new(Node)
	firstNode.Data = "first"

	secondNode := new(Node)
	secondNode.Data = "second"

	thirdNode := new(Node)
	thirdNode.Data = "third"

	headNode.Next = firstNode
	firstNode.Next = secondNode
	secondNode.Next = thirdNode

	return headNode
}



