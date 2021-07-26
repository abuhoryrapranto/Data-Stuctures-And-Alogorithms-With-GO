package stack

import "fmt"

const SIZE = 3

var top int = -1
var stack [SIZE]int

func Push(x int) {
	if top < SIZE-1 {
		top = top + 1
		stack[top] = x
		fmt.Println(x, " is added in Stack")
	} else {
		fmt.Println("Stack is full")
	}
}

func Pop() {
	if top >= 0 {
		top = top - 1
		fmt.Println("Poped from stack")
	} else {
		fmt.Println("Satck is empty")
	}
}

func Peek() {
	if top >= 0 {
		fmt.Println("Top is ", stack[top])
	} else {
		fmt.Println("Satck is empty")
	}
}
