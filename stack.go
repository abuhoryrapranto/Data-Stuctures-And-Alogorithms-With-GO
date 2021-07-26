package main

import "fmt"

const SIZE = 3

var top int = -1
var stack [SIZE]int

func push(x int) {
	if top < SIZE-1 {
		top = top + 1
		stack[top] = x
		fmt.Println(x, " is added in Stack")
	} else {
		fmt.Println("Stack is full")
	}
}

func pop() {
	if top >= 0 {
		top = top - 1
		fmt.Println("Poped from stack")
	} else {
		fmt.Println("Satck is empty")
	}
}

func peek() {
	if top >= 0 {
		fmt.Println("Top is ", stack[top])
	} else {
		fmt.Println("Satck is empty")
	}
}

func main() {
	push(5)
	push(4)
	push(3)
	pop()
	peek()
}
