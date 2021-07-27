package main
import (
	"go/stack"
	"go/queue"
)

func main() {

	/* Start Stack */

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	stack.Pop()

	stack.Peek()
	
	/* End Stack */

	/* Start Queue */

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	queue.Dequeue()

	queue.DisplayQueue()

	/* End Queue */
}
