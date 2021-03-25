package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
	//good thing about linkedlist is that you have to contain only the head of the node and length
}

func (l *linkedlist) prepend(n *node) {
	//(l *linkedlist1) is a reciever and it's a pointer
	//if we dont use pointer we would be working on copy of it
	//(n *node) to be added in front
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedlist) printlistdata() {
	//not using pointer cuz we arent going to change anything
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d", toPrint.data)
		toPrint = toPrint.next
		l.length--
		fmt.Printf("\n")
	}

}

func (l *linkedlist) deletevalue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			//if value dont exist
			return
		}
		previousToDelete = previousToDelete.next
		//goes untill find  A match
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedlist{}
	node1 := &node{data: 28}
	node2 := &node{data: 20}
	node3 := &node{data: 19}
	node4 := &node{data: 10}
	node5 := &node{data: 11}
	node6 := &node{data: 87}
	node7 := &node{data: 90}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.prepend(node7)
	// fmt.Println(mylist)
	mylist.printlistdata()
	mylist.deletevalue(87)
	fmt.Printf("\n")
	mylist.printlistdata()
}
