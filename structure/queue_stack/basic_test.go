package queuestack_test

import (
	"container/list"
	"testing"
)

type Stack struct {
	items []interface{}
}

type Queue struct {
	items *list.List
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func (q *Queue) Enqueue(item interface{}) {
	q.items.PushBack(item)
}

func (q *Queue) Dequeue() interface{} {
	if q.items.Len() == 0 {
		return nil
	}
	item := q.items.Front()
	q.items.Remove(item)
	return item.Value
}

func TestQueueStack(t *testing.T) {
	stack := Stack{}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stackItem := stack.Pop()
	t.Log("Stack Pop Item: ", stackItem)
	t.Log("Stack Items: ", stack.items)

	queue := Queue{items: list.New()}
	t.Log("Basic Queue Items: ", queue.items)
	queue.Enqueue("a")
	queue.Enqueue("b")
	queue.Enqueue("c")
	queue.Enqueue("d")
	queueItem := queue.Dequeue()
	t.Log("Queue Pop Item: ", queueItem)

	for i := queue.items.Front(); i != nil; i = i.Next() {
		t.Log("queueItem:", i.Value)
	}

}
