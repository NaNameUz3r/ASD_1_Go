package main

import (
	stack "ADS/stack_on_slice"
)

type QueueStacks[T any] struct {
	firstStack stack.Stack[T]
	secondStack stack.Stack[T]
}

func (q *QueueStacks[T]) Enqueue(item T) {
	q.firstStack.Push(item)
}

func (q *QueueStacks[T]) Dequeue() T {
	var resElement T
	if q.secondStack.Size() == 0 {
		for {
			if q.firstStack.Size() > 0 {
				element, _ := q.firstStack.Pop()
				q.secondStack.Push(element)
			} else {
				break
			}
		}
	}

	if q.secondStack.Size() > 0 {
		resElement, _ = q.secondStack.Pop()
	}

	return resElement
}

func (q *QueueStacks[T]) Size() int {
	var size int
	if q.firstStack.Size() == 0 && q.secondStack.Size() > 0 {
		size = q.secondStack.Size()
	} else if q.firstStack.Size() > 0 {
		size = q.firstStack.Size()
	}
	return size
}