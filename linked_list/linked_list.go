package main

import (
	// "os"
	"reflect"
	"errors"
	"fmt"
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

func (l *LinkedList) PrintList() []int {
	var node_values []int
	current_node := l.head
	for {
		if current_node == nil { break }
		node_values = append(node_values, current_node.value)
		current_node = current_node.next
	}
	// fmt.Println(node_values)
	return node_values

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
		purgedList := LinkedList{}
		for {
			if current_node == nil { break }
			if current_node.value != n {
				purgedList.AddInTail(Node{nil, current_node.value})
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
	// seek_node, finderr := l.Find(after.value)
	// if reflect.DeepEqual(seek_node, l.head) == true {
	// 		add.next = l.head.next
	// 		*l.head.next = add
	// } else if reflect.DeepEqual(seek_node, l.tail) == true {
	// 	l.AddInTail(add)
	// } else if finderr == nil {
	// 	add.next = seek_node.next
	// 	*seek_node.next = add
	// }

	// current_node := l.head
	// for {
	// 	if current_node == nil {
	// 		break
	// 	}

	// 	if reflect.DeepEqual(current_node, l.head) == true {
	// 		if reflect.DeepEqual(l.head, after) == true {
	// 			add.next = l.head.next
	// 			*l.head.next = add
	// 			break
	// 		}
	// 	}

	// 	if reflect.DeepEqual(current_node, l.tail) == true {
	// 		if reflect.DeepEqual(l.tail, after) {
	// 			l.AddInTail(add)
	// 		}
	// 	}

	// 	if reflect.DeepEqual(current_node.next, after) {
	// 		add.next = current_node.next.next
	// 		*current_node.next = add
	// 		break
	// 	}

	// 	current_node = current_node.next
	// }
}

func (l *LinkedList) InsertFirst(first Node) {

}

func (l *LinkedList) Clean() {
	l.head = nil
	l.tail = nil
	l = &LinkedList{}

}

// func main() {
// 	testList := LinkedList{}
// 	node1 := Node{nil, 1}
// 	node2 := Node{nil, 2}
// 	node3 := Node{nil, 3}

// 	fmt.Println(node1, testList.head)
// 	testList.AddInTail(node1)
// 	fmt.Println(node1, testList.head)
// 	testList.AddInTail(node2)
// 	fmt.Println(&node1, testList.head)
// 	testList.AddInTail(node3)

// 	testList.Insert(&node1, Node{nil, 777})
// }