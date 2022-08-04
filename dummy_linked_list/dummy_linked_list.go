package dummynode_linked_list

import (
	// "os"
	"reflect"
	"errors"
	// "fmt"
)
type Node struct {
	Prev  *Node
	Next  *Node
	Value interface{}
	Dummy bool
}

type LinkedListDummy struct {
	Head *Node
	Tail *Node
}

func NewList() *LinkedListDummy {
	l := new(LinkedListDummy)
	l.Head = &Node{Prev: nil, Next: nil, Value: nil, Dummy: true}
	l.Tail = &Node{Prev: nil, Next: nil, Value: nil, Dummy: true}
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head

	return l
}

func (l *LinkedListDummy) AddInTail(item Node) {
	if l.Head.Next.Dummy == true {
		l.Head.Next = &item
		l.Tail.Prev = &item
		item.Next = l.Tail
		item.Prev = l.Head
	} else {
		l.Tail.Prev.Next = &item
		item.Prev = l.Tail.Prev
		item.Next = l.Tail
		l.Tail.Prev = &item
	}
}

func (l *LinkedListDummy) GetHeadValue() interface{} {
	Value := l.Head.Next.Value
	return Value
}

func (l *LinkedListDummy) DeleteHead() error {

	var err error

	if l.Count() == 0 {
		err = errors.New("Delete head from empty list")
	} else if l.Count() == 1 {
		l.Clean()
	} else {
		l.Head.Next = l.Head.Next.Next
		l.Head.Next.Next.Prev = l.Head
	}

	return err
}

func (l *LinkedListDummy) GetListVals() []interface{} {
	var node_Values []interface{}
	current_node := l.Head.Next
	for {
		if current_node.Dummy == true {
			break
		}
		node_Values = append(node_Values, current_node.Value)
		current_node = current_node.Next
	}
	return node_Values

}

func (l *LinkedListDummy) Count() int {
	count := 0
	current_node := l.Head.Next
	for {
		if current_node.Dummy == true {
			break
		}
		count++
		current_node = current_node.Next
	}
	return count
}

// Find by Value
func (l *LinkedListDummy) Find(n interface{}) (Node, error) {
	if l.Head.Next.Dummy == true {
		return Node{Prev: nil, Next: nil, Value: -1}, nil
	}
	current_node := l.Head.Next
	for {
		if current_node.Dummy == true {
			return Node{Prev: nil, Next: nil, Value: -1}, errors.New("Node not found")
		}
		if current_node.Value == n {
			return *current_node, nil
		}
		current_node = current_node.Next
	}
}

func (l *LinkedListDummy) FindAll(n interface{}) []Node {
	var nodes []Node
	current_node := l.Head.Next
	for {
		if current_node.Dummy == true {
			break
		}
		if current_node.Value == n {
			nodes = append(nodes, *current_node)
		}
		current_node = current_node.Next
	}
	return nodes
}
// Delete by Value
func (l *LinkedListDummy) Delete(n interface{}, all bool) {
	current_node := l.Head.Next
	if all == true {
		// Delete all nodes with argument Value
		purgedList := *NewList()
		for {
			if current_node.Dummy == true {
				break
			}
			if current_node.Value != n {
				purgedList.AddInTail(Node{Prev: nil, Next: nil, Value: current_node.Value, Dummy: false})
			}
			current_node = current_node.Next
		}
		*l = purgedList

	} else {
		// Delete first met node with argument Value
		for {
			if current_node == nil {
				break
			}
			if current_node.Value == n {
				if reflect.DeepEqual(current_node, l.Head) == true {
					if reflect.DeepEqual(l.Head, l.Tail) == true {
						l.Clean()
						break
					} else {
						l.Head = l.Head.Next
						l.Head.Prev = nil
						break
					}
				} else if reflect.DeepEqual(current_node, l.Tail) == true {
					l.Tail = l.Tail.Prev
					l.Tail.Next = nil
					break
				} else {
					current_node.Prev.Next = current_node.Next
					current_node.Next.Prev = current_node.Prev
					break
				}
			} else {
				current_node = current_node.Next
			}
		}
	}
}

func (l *LinkedListDummy) Insert(after *Node, add Node) {
	if reflect.DeepEqual(after, l.Tail.Prev) == true {
		l.AddInTail(add)
	} else {
		add.Next = after.Next
		add.Prev = after
		after.Next.Prev = &add
		after.Next = &add
	}
}

func (l *LinkedListDummy) InsertFirst(first Node) {
	if l.Head.Dummy == true && l.Head.Next.Dummy == true {
		l.AddInTail(first)
	} else {
		first.Next = l.Head.Next
		l.Head.Next.Prev = &first
		l.Head.Next = &first
	}
}

func (l *LinkedListDummy) Clean() {
	l.Head.Next = l.Tail
	newList := *NewList()
	*l = newList

}
