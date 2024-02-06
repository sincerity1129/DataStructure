package queuestack_test

import (
	"container/list"
	"fmt"
	"testing"
)

type Stack []int

type Queue struct {
	list *list.List
}

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		panic("Stack is empty")
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Push(value int) {
	q.list.PushBack(value)
}

func (q *Queue) Pop() int {
	if q.list.Len() == 0 {
		panic("Queue is empty")
	}
	elem := q.list.Front()
	q.list.Remove(elem)
	return elem.Value.(int)
}

func TestQueueStack(t *testing.T) {
	stack := make(Stack, 0)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("-------stack--------")
	for len(stack) > 0 {
		value := stack.Pop()
		fmt.Println(value) // 출력: 3, 2, 1
	}

	queue := NewQueue()

	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	fmt.Println("-------queue--------")
	for queue.list.Len() > 0 {
		value := queue.Pop()
		fmt.Println(value) // 출력: 1, 2, 3
	}
}
