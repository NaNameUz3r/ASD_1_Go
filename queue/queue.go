package main

import (
	"os"
	"fmt"
)

type Queue[T any] struct {
	queue []T
}

func (q *Queue[T]) Size() (int) {
	var size int
	size = len(q.queue)
	return size
}

func (q *Queue[T]) Dequeue() (T, error) {
	var result T
	var err error
	if q.Size() > 0 {
		result, q.queue = q.queue[0], q.queue[1:]
	} else {
		err = fmt.Errorf("Dequeue empty queue")
	}
	return result, err
}

func (q *Queue[T]) Enqueue(itm T) {
	q.queue = append(q.queue, itm)
}

func (q *Queue[T]) SpinQueue(spinrange int) {
	for i := 0; i < spinrange; i++ {
		element, _ := q.Dequeue()
		q.Enqueue(element)
	}
}