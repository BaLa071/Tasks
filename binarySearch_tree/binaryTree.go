package main

import "fmt"

type node struct {
	left  *node
	data  int
	right *node
}

func creatnode(data int) *node {
	newnode := new(node)
	newnode.data = data
	newnode.left = nil
	newnode.right = nil
	return newnode
}

func insertnode(data int, root *node) *node {
	if root == nil {
		return creatnode(data)
	}
	if data < root.data {
		root.left = insertnode(data, root.left)
	} else if data > root.data {
		root.right = insertnode(data, root.right)
	}
	return root
}

func display(head *node) {
	if head == nil {
		return
	}
	display(head.left)
	fmt.Printf("%d ", head.data)
	display(head.right)
}

func search(data int, head *node) {
	if head == nil {
		return
	}
	if data == head.data {
		fmt.Printf("Found %d\n", head.data)
	}
	if data < head.data {
		search(data, head.left)
	}
	if data > head.data {
		search(data, head.right)
	}
}

func main() {

	root := creatnode(1)
	head := root
	a := []int{
		5, 6, 2, 4, 9, 0,
	}

	for i := 0; i < len(a); i++ {
		head = insertnode(a[i], head)
	}

	display(head)
	search(9, head)

}
