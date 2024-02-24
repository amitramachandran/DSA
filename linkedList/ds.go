package linkedList

import "fmt"

type Data interface{}

type Node struct {
	data Data
	next *Node
	head bool
}

func NewNode(data Data) *Node {
	return &Node{data, nil, false}
}

// operations
func (n *Node) Head() bool {
	return n.head
}

func (n *Node) SetHead() {
	n.head = true
}

func (n *Node) GetNext() *Node {
	return n.next
}

func (n *Node) GetData() Data {
	return n.data
}

func (n *Node) Connect(to *Node) {
	if to.head {
		n.head = true
		to.head = false
	}
	// checks if we are adding middle of LL
	if n.next != nil && !n.head {
		to.next = n.next
	}
	n.next = to

}

func (n *Node) Delete(head *Node) {
	var temp *Node
	if head.Head() {
		temp = head
	} else {
		fmt.Println("Kindly pass first node/ head in the params to delete")
	}

	// deleting the head
	if n.data == temp.data {
		temp.next.head = true
		temp.head = false
		temp.next = nil
	}

	// to delete node in middle and last of LL
	for temp.next != nil {
		if n.data == temp.next.data {
			temp.next = temp.next.next
			break
		}
		temp = temp.next
	}

}
