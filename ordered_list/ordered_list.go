package main

import (
        "constraints"
        "fmt"
        "os"
)

type Node[T constraints.Ordered] struct {
        prev  *Node[T]
        next  *Node[T]
        value T
}

type OrderedList[T constraints.Ordered] struct {
        head       *Node[T]
        tail       *Node[T]
        _ascending bool
}

func (l *OrderedList[T]) Count() int {
        nodeCounter := 0
        currentNode := l.head

        for {
                if currentNode == nil {
                        break
                }
                nodeCounter++
                currentNode = currentNode.next
        }
        return nodeCounter
}

func (l *OrderedList[T]) GetListVals() []T {
	var values []T
	currentNode := l.head

	for {
		if currentNode == nil {
			break
		}
		values = append(values, currentNode.value)
		currentNode = currentNode.next
	}

	return values
}

func (l *OrderedList[T]) Add(item *Node[T]) {
        currentNode := l.head
        if currentNode == nil {
                l.head = item
                l.tail = item
        } else {

                for {
                        if currentNode != nil {
                                if (l._ascending == true && l.Compare(item.value, currentNode.value) <= 0) || (
									l._ascending == false && l.Compare(item.value, currentNode.value) >= 0) {
                                        if currentNode.prev == nil {
                                                item.next = currentNode
                                                currentNode.prev = item
                                                l.head = item
                                        } else if currentNode.prev != nil {
                                                currentNode.prev.next = item
                                                item.prev = currentNode.prev
                                                item.next = currentNode
                                                currentNode.prev = item
                                        }
                                        break
                                } else if currentNode.next == nil {
                                        currentNode.next = item
                                        item.prev = currentNode
                                        l.tail = item
                                        break
                                }
                        }

                        currentNode = currentNode.next
                }
        }
}

func (l *OrderedList[T]) Find(n T) (Node[T], error) {
        var err error

		if l.head == nil {
			err = fmt.Errorf("Empty list nothing to give you")
	    	return Node[T]{value: n, next: nil, prev: nil}, err
		}

		currentNode := l.head
		for {
			if currentNode == nil {
				err = fmt.Errorf("Node with %T value is not found", n)
				return Node[T]{value: n, next: nil, prev: nil}, err

			}
			if (l._ascending == true && currentNode.value > n) || (
				l._ascending == false && currentNode.value < n) {
					err = fmt.Errorf("Node with %T value is not found", n)
					return Node[T]{value: n, next: nil, prev: nil}, err
				}
			if currentNode.value == n {
				return *currentNode, err
			}


			currentNode = currentNode.next
		}
}

func (l *OrderedList[T]) Delete(n T) {
        currentNode := l.head

        for {
                if currentNode == nil {
                        break
                }

                if currentNode.value == n && currentNode == l.head {
                        if l.head == l.tail {
                                l.Clear(l._ascending)
                                break
                        }

                        l.head = l.head.next
                        l.head.prev = nil
                        break
                }

                if currentNode.value == n && currentNode == l.tail {
                        l.tail = l.tail.prev
                        l.tail.next = nil
                        break
                }

                if currentNode.value == n {
                        currentNode.prev.next = currentNode.next
                        currentNode.next.prev = currentNode.prev
                }

                currentNode = currentNode.next
        }
}

func (l *OrderedList[T]) Clear(asc bool) {
        l.head, l.tail = nil, nil
        l._ascending = asc
}

func (l *OrderedList[T]) Compare(v1 T, v2 T) int {
        if v1 < v2 {
                return -1
        }
        if v1 > v2 {
                return +1
        }
        return 0
}