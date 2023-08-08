package main

import "fmt"

type node struct {
	data int
	prev *node
	next *node
}

func createNode(val int) *node {
	return &node{
		data: val,
		next: nil,
		prev: nil,
	}
}
func insertAtFirst(val int, head **node) {
	var temp *node = createNode(val)
	temp.next = *head
	(*head).prev = temp
	*head = temp
}

func insertAtLast(val int, tail **node) {
	var temp *node = createNode(val)
	(*tail).next = temp
	temp.prev = *tail
	(*tail) = (*tail).next
}

func insertAtMiddle(val int, position int, head *node) {
	var temp *node = createNode(val)
	var new_temp *node = head
	for new_temp != nil {
		if new_temp.data == position {
			var flag *node = new_temp.next
			temp.prev = new_temp
			new_temp.next = temp
			temp.next = flag
			flag.prev = temp
		}
		new_temp = new_temp.next
	}
}

func main() {
	var head, tail *node
	head = createNode(1)
	tail = head

	//inserting elements at various position
	insertAtFirst(0, &head)
	insertAtLast(2, &tail)
	insertAtLast(3, &tail)
	insertAtLast(5, &tail)

	insertAtMiddle(4, 3, head)

	//printing statement
	for head.next != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
