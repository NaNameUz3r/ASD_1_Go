package main

import (
	// "os"
	"reflect"
	"errors"
	// "fmt"
)

type Node struct {
	next  *Node
	value int
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddInTail(item Node) {
	if l.head == nil {
		l.head = &item
	} else {
		l.tail.next = &item
	}

	l.tail = &item
}

func (l *LinkedList) Count() int {
	count := 0
	current_node := l.head
	for {
		if current_node == nil { break }
		count ++
		current_node = current_node.next
	}
	return count
}

// error не nil, если узел не найден
func (l *LinkedList) Find(n int) (Node, error) {
	if l.head == nil { return Node{value: -1, next: nil}, nil }
	current_node := l.head
	for {
		if current_node == nil {
			return Node{value: -1, next: nil}, errors.New("Node not found")
		}
		if current_node.value == n { return *current_node, nil }
		current_node = current_node.next
	}
}

func (l *LinkedList) FindAll(n int) []Node {
	var nodes []Node
	current_node := l.head
	for {
		if current_node == nil { break }
		if current_node.value == n {
			nodes = append(nodes, *current_node)
		}
		current_node = current_node.next
		}
	return nodes
	}


func (l *LinkedList) Delete(n int, all bool) {
	current_node := l.head

	if all == true {
		// Delete all nodes with argument value
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
					}
				}
			} else if current_node.next.value == n {
				if reflect.DeepEqual(current_node.next, l.tail) == true {
					l.tail = current_node
					l.tail.next = nil
					break
				}
			}
			current_node = current_node.next
		}
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
					}
					l.head = l.head.next
					break
				}
			} else if current_node.next.value == n {
				if reflect.DeepEqual(current_node.next, l.tail) == true {
					l.tail = current_node
					l.tail.next = nil
					break
				} else {
					current_node.next = current_node.next.next
					break
				}
			}
			current_node = current_node.next
		}
}
}

func (l *LinkedList) Insert(after *Node, add Node) {

}

func (l *LinkedList) InsertFirst(first Node) {

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
	l = &LinkedList{}

}

func main() {
}