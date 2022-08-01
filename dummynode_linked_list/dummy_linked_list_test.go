package dummynode_linked_list

import (
	"fmt"
	"reflect"
	"testing"
	// "reflect"
)

func TestLinkedListCreateAndCount(t *testing.T) {
	testList := NewList()


	if testList.head.dummy != true {
		t.Errorf("Head is not dummy")
	}
	if testList.tail.dummy != true {
		t.Errorf("Tail is not dummy")
	}


	testList.AddInTail(Node{nil, nil, 1, false})
	testList.AddInTail(Node{nil, nil, 2, false})
	testList.AddInTail(Node{nil, nil, 3, false})

	if testList.Count() != 3 {
		t.Errorf("Count works wrong, there is 3 nodes, but Counted %d", testList.Count())
	}

	if testList.head.next.value != 1 {
		t.Errorf("Head next value is wrong")
	}

	if testList.tail.prev.value != 3 {
		t.Errorf("Tail prev value is wrong")
	}

}

func TestFind(t *testing.T) {
	testList := NewList()
	node1 := &Node{nil, nil, 1, false}
	node2 := &Node{nil, nil, 2, false}
	node3 := &Node{nil, nil, 3, false}
	testList.AddInTail(*node1)
	testList.AddInTail(*node2)
	testList.AddInTail(*node3)

	somenode, err := testList.Find(1)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode, testList.head.next) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode2, err := testList.Find(2)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode2, testList.head.next.next) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode3, err := testList.Find(3)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode3, testList.tail.prev) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestFindAll(t *testing.T) {
	testList := NewList()
	node1 := &Node{nil, nil, 1, false}
	node2 := &Node{nil, nil, 1, false}
	node3 := &Node{nil, nil, 1, false}
	node4 := &Node{nil, nil, 4, false}
	node5 := &Node{nil, nil, 5, false}
	node6 := &Node{nil, nil, 6, false}

	testList.AddInTail(*node1)
	testList.AddInTail(*node2)
	testList.AddInTail(*node3)
	testList.AddInTail(*node4)
	testList.AddInTail(*node5)
	testList.AddInTail(*node6)

	someNodes := testList.FindAll(1)

	checkSlice := []Node{*testList.head.next, *testList.head.next.next, *testList.head.next.next.next}
	if reflect.DeepEqual(someNodes, checkSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	testList2 := NewList()
	nodeX := &Node{nil, nil, 1, false}
	nodeY := &Node{nil, nil, 2, false}
	nodeZ := &Node{nil, nil, 2, false}

	testList2.AddInTail(*nodeX)
	testList2.AddInTail(*nodeY)
	testList2.AddInTail(*nodeZ)

	yetAnotherNodes := testList2.FindAll(2)
	yetAnotherSlice := []Node{*testList2.tail.prev.prev, *testList2.tail.prev}
	if reflect.DeepEqual(yetAnotherNodes, yetAnotherSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestDeleteAllSameHeadTail(t *testing.T) {
	// Test Delete of list of same values
	testList := NewList()
	node1 := &Node{nil, nil, 6, false}
	node2 := &Node{nil, nil, 6, false}
	node3 := &Node{nil, nil, 2, false}
	node4 := &Node{nil, nil, 2, false}
	node5 := &Node{nil, nil, 1, false}
	node6 := &Node{nil, nil, 6, false}
	testList.AddInTail(*node1)
	testList.AddInTail(*node2)
	testList.AddInTail(*node3)
	testList.AddInTail(*node4)
	testList.AddInTail(*node5)
	testList.AddInTail(*node6)

	// Test delete same consecutive in head and tail
	testList.Delete(6, true)

	if testList.head.next.value != 2 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.head.dummy != true {
		fmt.Printf("head is not dummy")
		t.Error()
	}

	if testList.head.next.value != 2 {
		fmt.Printf("Incorrect head next.next node value")
		t.Error()
	}

	if testList.head.next.next.value != 2 {
		fmt.Printf("Incorrect head next.next node value")
		t.Error()
	}

	if testList.head.next.next.prev.value != 2 {
		fmt.Printf("Incorrect head next.next node prev value")
		t.Error()
	}


	if testList.tail.dummy != true {
		fmt.Printf("Tail is not dummy")
		t.Error()
	}

	if testList.tail.prev.value != 1 {
		fmt.Printf("Incorrect tail prev node value")
		t.Error()
	}

	if testList.tail.next != nil {
		fmt.Printf("List tail next node is not <nil")
		t.Error()
	}
}

func TestDeleteAllSame(t *testing.T) {
	// Test Delete List with same consecutive
	testListSameVals := NewList()
	node1 := &Node{nil, nil, 1, false}
	node2 := &Node{nil, nil, 1, false}
	node3 := &Node{nil, nil, 1, false}
	testListSameVals.AddInTail(*node1)
	testListSameVals.AddInTail(*node2)
	testListSameVals.AddInTail(*node3)

	testListSameVals.Delete(1, true)

	if testListSameVals.tail.dummy != true {
		fmt.Printf("List tail is not dummy")
		t.Error()
	}

	if testListSameVals.head.dummy != true {
		fmt.Printf("List tail is not dummy")
		t.Error()
	}

	if testListSameVals.head.next.dummy != true {
		fmt.Printf("head next is not dummy")
		t.Error()
	}
}

func TestDeleteAllSameMiddle(t *testing.T) {
	// Test Delete of list of same values
	testList := NewList()
	node1 := &Node{nil, nil, 1, false}
	node2 := &Node{nil, nil, 2, false}
	node3 := &Node{nil, nil, 2, false}
	node4 := &Node{nil, nil, 2, false}
	node5 := &Node{nil, nil, 2, false}
	node6 := &Node{nil, nil, 6, false}
	testList.AddInTail(*node1)
	testList.AddInTail(*node2)
	testList.AddInTail(*node3)
	testList.AddInTail(*node4)
	testList.AddInTail(*node5)
	testList.AddInTail(*node6)
	// Test delete same consecutive in list middle
	testList.Delete(2, true)


	if testList.tail.prev.value != 6 {
		fmt.Printf("Incorrect tail prev value")
		t.Error()
	}

	if testList.head.next.value != 1 {
		fmt.Printf("Incorrect head next value")
		t.Error()
	}

	if reflect.DeepEqual(testList.tail.prev, testList.head.next.next) != true {
		fmt.Printf("Tail prev is not head next")
		t.Error()
	}

	if reflect.DeepEqual(testList.head.next, testList.tail.prev.prev) != true {
		fmt.Printf("Head next is not tail prev prev")
		t.Error()
	}
}

func TestInsert(t *testing.T) {
	testList := NewList()

	node1 := Node{nil, nil, 1, false}
	node2 := Node{nil, nil, 2, false}
	node3 := Node{nil, nil, 2, false}
	nodeadd1 := Node{nil, nil,  777, false}
	nodeadd2 := Node{nil, nil,  888, false}
	nodeadd3 := Node{nil, nil,  999, false}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	// insert after head next
	testList.Insert(testList.head.next, nodeadd1)

	if testList.head.next.next.value != 777 {
		fmt.Printf("head next value wrong")
		t.Error()
	}

	if testList.head.next.value != 1 {
		fmt.Printf("head next node value wrong")
		t.Error()
	}

	if reflect.DeepEqual(testList.head.next.prev, testList.head) != true {
		fmt.Printf("head hext node prev not equal dummyhead")
		t.Error()
	}

	// insert after actual tail
	testList.Insert(testList.tail.prev, nodeadd2)

	if testList.tail.prev.value != 888 {
		fmt.Printf("tail value wrong")
		t.Error()
	}

	if testList.tail.next != nil {
		fmt.Printf("List tail next node is not <nil>")
		t.Error()
	}

	if testList.tail.dummy != true {
		fmt.Printf(" tail is not dummy")
		t.Error()
	}



	if testList.tail.prev.value != 888 {
		fmt.Printf("tail prev node value wrong")
		t.Error()
	}
	// insert after third node
	testList.Insert(testList.head.next.next.next, nodeadd3)

	if testList.head.next.next.next.next.value != 999 {
		fmt.Printf("wrong fourth node value")
		t.Error()
	}

	if testList.head.next.next.next.next.next.value != 2 {
		fmt.Printf("wrong fourth node next value")
		t.Error()
	}

	if testList.head.next.next.next.next.next.prev.value != 999 {
		fmt.Printf("wrong fourth node prev value")
		t.Error()
	}
}

func TestInserFirst(t *testing.T) {
	testList := NewList()

	node1 := Node{nil, nil, 1, false}
	node2 := Node{nil, nil, 2, false}
	node3 := Node{nil, nil, 2, false}
	nodeadd1 := Node{nil, nil,  777, false}

	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	testList.InsertFirst(nodeadd1)

	if testList.head.next.value != 777 {
		fmt.Printf("head value wrong")
		t.Error()
	}
	if testList.head.next.next.value != 1 {
		fmt.Printf("head value wrong")
		t.Error()
	}

	if testList.head.next.next.prev.value != 777 {
		fmt.Printf("head next node prev value wrong")
		t.Error()
	}
}


func TestClean(t *testing.T) {
	testList := NewList()

	node1 := Node{nil, nil, 1, false}
	node2 := Node{nil, nil, 2, false}
	node3 := Node{nil, nil, 2, false}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	testList.Clean()
	if testList.head.next.dummy != true {
		fmt.Printf("head next node is not dummy")
		t.Error()
	}
}