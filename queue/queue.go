package queue

import "fmt"

const SIZE = 3
var front, rear = -1, -1
var queue[SIZE]int

func Enqueue(x int) {
	if rear == SIZE-1  {
		fmt.Println("Queue is full")
	} else {
		if front == -1 {
			front = 0
		}
		rear++;
		queue[rear] = x
		fmt.Println(x, " added in queue")
	}
}

func Dequeue() {
	if front == -1 {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println(queue[front], " is dequeue from queue")
		front++;
		if front > rear {
			front = -1
			rear = -1
		}
	}
}

func DisplayQueue() {
	if rear == -1 {
		fmt.Println("Queue is Empty")
	} else {
		fmt.Print("Queue: ")
		for i := front; i <= rear; i++ {
			fmt.Print(queue[i], " ")
		}
	}
}