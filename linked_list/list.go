package main

import "fmt"

type node struct {
	data int
	next *node
}

func createNode(val int) *node {
	return &node{
		data: val,
		next: nil,
	}
}

func insertAtFirst(val int, head **node) {
	var temp *node = createNode(val)
	temp.next = *head
	*head = temp
}

func insertAtLast(val int, root **node) {
	var temp *node = createNode(val)
	(*root).next = temp
	*root = (*root).next
}

func insertAtMiddle(val, place int, head *node) {
	var temp *node = createNode(val)
	var new_temp *node = head
	for new_temp != nil {
		if (new_temp).data == place {
			flag := new_temp.next
			new_temp.next = temp
			temp.next = flag
			break
		}
		new_temp = new_temp.next
	}
}

func main() {
	var head, root *node
	head = createNode(1)
	root = head
	insertAtFirst(0, &head)
	insertAtLast(2, &root)
	insertAtLast(3, &root)
	insertAtLast(5, &root)
	insertAtMiddle(4, 3, head)
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}

}
