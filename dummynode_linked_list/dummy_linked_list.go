package dummynode_linked_list

import (
	// "os"
	"reflect"
	"errors"
	// "fmt"
)
type Node struct {
	prev  *Node
	next  *Node
	value int
	dummy bool
}

type LinkedListDummy struct {
	head *Node
	tail *Node
}

func NewList() *LinkedListDummy {
	l := new(LinkedListDummy)
	l.head = &Node{prev: nil, next: nil, value: -1, dummy: true}
	l.tail = &Node{prev: nil, next: nil, value: -1, dummy: true}
	l.head.next = l.tail
	l.tail.prev = l.head

	return l
}

func (l *LinkedListDummy) AddInTail(item Node) {
	if l.head.next.dummy == true {
		l.head.next = &item
		l.tail.prev = &item
		item.next = l.tail
		item.prev = l.head
	} else {
		l.tail.prev.next = &item
		item.prev = l.tail.prev
		item.next = l.tail
		l.tail.prev = &item
	}
}

func (l *LinkedListDummy) GetListVals() []int {
	var node_values []int
	current_node := l.head.next
	for {
		if current_node.dummy == true {
			break
		}
		node_values = append(node_values, current_node.value)
		current_node = current_node.next
	}
	return node_values

}

func (l *LinkedListDummy) Count() int {
	count := 0
	current_node := l.head.next
	for {
		if current_node.dummy == true {
			break
		}
		count++
		current_node = current_node.next
	}
	return count
}

// error не nil, если узел не найден
func (l *LinkedListDummy) Find(n int) (Node, error) {
	if l.head.next.dummy == true {
		return Node{prev: nil, next: nil, value: -1}, nil
	}
	current_node := l.head.next
	for {
		if current_node.dummy == true {
			return Node{prev: nil, next: nil, value: -1}, errors.New("Node not found")
		}
		if current_node.value == n {
			return *current_node, nil
		}
		current_node = current_node.next
	}
}

func (l *LinkedListDummy) FindAll(n int) []Node {
	var nodes []Node
	current_node := l.head.next
	for {
		if current_node.dummy == true {
			break
		}
		if current_node.value == n {
			nodes = append(nodes, *current_node)
		}
		current_node = current_node.next
	}
	return nodes
}

func (l *LinkedListDummy) Delete(n int, all bool) {
	current_node := l.head.next
	if all == true {
		// Delete all nodes with argument value
		purgedList := *NewList()
		for {
			if current_node.dummy == true {
				break
			}
			if current_node.value != n {
				purgedList.AddInTail(Node{prev: nil, next: nil, value: current_node.value, dummy: false})
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
						l.head.prev = nil
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

func (l *LinkedListDummy) Insert(after *Node, add Node) {
	if reflect.DeepEqual(after, l.tail.prev) == true {
		l.AddInTail(add)
	} else {
		add.next = after.next
		add.prev = after
		after.next.prev = &add
		after.next = &add
	}
}

func (l *LinkedListDummy) InsertFirst(first Node) {
	if l.head.dummy == true && l.head.next.dummy == true {
		l.AddInTail(first)
	} else {
		first.next = l.head.next
		l.head.next.prev = &first
		l.head.next = &first
	}
}

func (l *LinkedListDummy) Clean() {
	l.head.next = l.tail
	newList := *NewList()
	*l = newList

}
