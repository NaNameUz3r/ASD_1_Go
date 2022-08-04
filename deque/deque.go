package main

import (
	// "os"
	"fmt"
)

type Deque[T any] struct {
	deque []T
}

func (d *Deque[T]) Size() (int) {
	var size int
	size = len(d.deque)
	return size
}

func (d *Deque[T]) AddFront(itm T) {
	d.deque = append(d.deque, itm)
	copy(d.deque[1:], d.deque)
	d.deque[0] = itm
}

func (d *Deque[T]) AddTail(itm T) {
	d.deque = append(d.deque, itm)
}

func (d *Deque[T]) RemoveFront() (T, error) {
	var result T
	var err error

	if d.Size() > 0 {
		result, d.deque = d.deque[0], d.deque[1:]
	} else {
		err = fmt.Errorf("Removing head from empty deque")
	}

	return result, err
}

func (d *Deque[T]) RemoveTail() (T, error) {
	var result T
	var err error

	if d.Size() > 0 {
		result, d.deque = d.deque[len(d.deque) - 1], d.deque[:len(d.deque) - 1]
	} else {
		err = fmt.Errorf("Removing tail from empty deque")
	}
	return result, err
}