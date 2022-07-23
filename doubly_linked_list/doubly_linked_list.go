package main

import (
	"os"
	"reflect"
	"errors"
)
type Node struct {
	prev  *Node
	next  *Node
	value int
}

type LinkedList2 struct {
	head *Node
	tail *Node
}

func (l *LinkedList2) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
		l.head.next = nil
		l.head.prev = nil
	} else {
		l.tail.next = &item
		item.prev = l.tail
	}

	l.tail = &item
	l.tail.next = nil
}

func (l *LinkedList2) GetListVals() []int {
	var node_values []int
	current_node := l.head
	for {
		if current_node == nil {
			break
		}
		node_values = append(node_values, current_node.value)
		current_node = current_node.next
	}
	return node_values

}

func (l *LinkedList2) Count() int {
	count := 0
	current_node := l.head
	for {
		if current_node == nil {
			break
		}
		count++
		current_node = current_node.next
	}
	return count
}

// error не nil, если узел не найден
func (l *LinkedList2) Find(n int) (Node, error) {
	if l.head == nil {
		return Node{prev: nil, next: nil, value: -1}, nil
	}
	current_node := l.head
	for {
		if current_node == nil {
			return Node{prev: nil, next: nil, value: -1}, errors.New("Node not found")
		}
		if current_node.value == n {
			return *current_node, nil
		}
		current_node = current_node.next
	}
}

func (l *LinkedList2) FindAll(n int) []Node {
	var nodes []Node
	current_node := l.head
	for {
		if current_node == nil {
			break
		}
		if current_node.value == n {
			nodes = append(nodes, *current_node)
		}
		current_node = current_node.next
	}
	return nodes
}

func (l *LinkedList2) Delete(n int, all bool) {
	current_node := l.head
	if all == true {
		// Delete all nodes with argument value
		purgedList := LinkedList2{}
		for {
			if current_node == nil {
				break
			}
			if current_node.value != n {
				purgedList.AddInTail(Node{prev: nil, next: nil, value: current_node.value})
			}
			current_node = current_node.next
		}
		*l = purgedList

	} else {
		// Delete first met node with argument value
		for {
			if current_node == nil {
				break
			}
			if current_node.value == n {
				if reflect.DeepEqual(current_node, l.head) == true {
					if reflect.DeepEqual(l.head, l.tail) == true {
						l.Clean()
						break
					} else {
						l.head = l.head.next
						break
					}
				} else if reflect.DeepEqual(current_node, l.tail) == true {
					l.tail = l.tail.prev
					l.tail.next = nil
					break
				} else {
					current_node.prev.next = current_node.next
					current_node.next.prev = current_node.prev
					break
				}
			} else {
				current_node = current_node.next
			}
		}
	}
}

func (l *LinkedList2) Insert(after *Node, add Node) {
	if reflect.DeepEqual(after, l.tail) == true {
		l.AddInTail(add)
	} else {
		add.next = after.next
		add.prev = after
		after.next.prev = &add
		after.next = &add
	}
}

func (l *LinkedList2) InsertFirst(first Node) {
	if l.head == nil {
		l.AddInTail(first)
	} else {
		first.next = l.head
		l.head.prev = &first
		l.head = &first
	}
}

func (l *LinkedList2) Clean() {
	l.head = nil
	l.tail = nil
	l = &LinkedList2{}

}
