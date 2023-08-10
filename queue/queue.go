package main

import "fmt"

func enqueue(queue []int, element int) []int {
	queue = append(queue, element)
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0]
	fmt.Println(element)
	return queue[1:]
}

func main() {
	var queue []int

	queue = enqueue(queue, 10)
	queue = enqueue(queue, 20)
	queue = enqueue(queue, 30)

	fmt.Println("Queue:", queue)

	queue = dequeue(queue)
	fmt.Println("Queue:", queue)

	queue = enqueue(queue, 40)
	fmt.Println("Queue:", queue)
}
