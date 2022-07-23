package main

import (
	"fmt"
	// "reflect"
	"testing"
	"reflect"
)

func TestLinkedListCreate(t *testing.T) {
	testList := LinkedList2{}

	testList.AddInTail(Node{nil, nil, 1})
	testList.AddInTail(Node{nil, nil, 2})
	testList.AddInTail(Node{nil, nil, 3})

	if testList.head.value != 1 {
		t.Errorf("head value is incorrect, it is not 1 but %d", testList.head.value)
	}

	if testList.head.next.value != 2 {
		t.Errorf("head next node value is not correct. It is equal to %d, should be 2", testList.head.next.value)
	}

	if testList.head.next.prev.value != 1 {
		t.Errorf("2nd node prev value is not correct. It is equal to %d, should be 1", testList.head.next.value)
	}

	if testList.tail.value != 3 {
		t.Errorf("wrong tail value - %d, should be 3", testList.tail.value)
	}

	if testList.tail.prev.value != 2 {
		t.Errorf("wrong tail prev node value - %d, should be 2", testList.tail.value)
	}
	// check if tail next is nil

	if testList.tail.next != nil {
		t.Errorf("something went wrong, tail next 'value' is not nil")
	}

}

func TestLinkedListCount(t *testing.T) {
	testList := LinkedList2{}

	testList.AddInTail(Node{nil, nil, 1})
	testList.AddInTail(Node{nil, nil, 2})
	testList.AddInTail(Node{nil, nil, 3})

	if testList.Count() != 3 {
		t.Errorf("Count working incorrect")
	}
}

func TestFind(t *testing.T) {
	testList := LinkedList2{}
	node1 := Node{nil, nil, 1}
	node2 := Node{nil, nil, 2}
	node3 := Node{nil, nil, 3}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	somenode, err := testList.Find(1)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode, testList.head) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode2, err := testList.Find(2)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode2, testList.head.next) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode3, err := testList.Find(3)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode3, testList.tail) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestFindAll(t *testing.T) {
	testList := LinkedList2{}
	node1 := Node{nil, nil, 1}
	node2 := Node{nil, nil, 1}
	node3 := Node{nil, nil, 1}
	node4 := Node{nil, nil, 4}
	node5 := Node{nil, nil, 5}
	node6 := Node{nil, nil, 6}

	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)
	testList.AddInTail(node4)
	testList.AddInTail(node5)
	testList.AddInTail(node6)

	someNodes := testList.FindAll(1)

	checkSlice := []Node{*testList.head, *testList.head.next, *testList.head.next.next}
	if reflect.DeepEqual(someNodes, checkSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	testList2 := LinkedList2{}
	nodeX := Node{nil, nil, 1}
	nodeY := Node{nil, nil, 2}
	nodeZ := Node{nil, nil, 2}

	testList2.AddInTail(nodeX)
	testList2.AddInTail(nodeY)
	testList2.AddInTail(nodeZ)

	yetAnotherNodes := testList2.FindAll(2)
	yetAnotherSlice := []Node{*testList2.head.next, *testList2.tail}
	if reflect.DeepEqual(yetAnotherNodes, yetAnotherSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestDeleteAllSameHeadTail(t *testing.T) {
	// Test Delete of list of same values
	testList := LinkedList2{}
	node1 := Node{nil, nil, 6}
	node2 := Node{nil, nil, 6}
	node3 := Node{nil, nil, 2}
	node4 := Node{nil, nil, 2}
	node5 := Node{nil, nil, 1}
	node6 := Node{nil, nil, 6}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)
	testList.AddInTail(node4)
	testList.AddInTail(node5)
	testList.AddInTail(node6)

	// Test delete same consecutive in head and tail
	testList.Delete(6, true)

	if testList.head.value != 2 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.head.prev != nil {
		fmt.Printf("head prev != nil")
		t.Error()
	}

	if testList.head.next.prev.value != 2 {
		fmt.Printf("Incorrect head next node prev value")
		t.Error()
	}

	if testList.tail.value != 1 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.tail.prev.value != 2 {
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
	testListSameVals := LinkedList2{}
	node1 := Node{nil, nil, 1}
	node2 := Node{nil, nil, 1}
	node3 := Node{nil, nil, 1}
	testListSameVals.AddInTail(node1)
	testListSameVals.AddInTail(node2)
	testListSameVals.AddInTail(node3)

	testListSameVals.Delete(1, true)

	if testListSameVals.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}

	if testListSameVals.head != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}
}

func TestDeleteAllSameMiddle(t *testing.T) {
	// Test Delete of list of same values
	testList := LinkedList2{}

	node1 := Node{nil, nil, 1}
	node2 := Node{nil, nil, 2}
	node3 := Node{nil, nil, 2}
	node4 := Node{nil, nil, 2}
	node5 := Node{nil, nil, 2}
	node6 := Node{nil, nil, 6}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)
	testList.AddInTail(node4)
	testList.AddInTail(node5)
	testList.AddInTail(node6)
	// Test delete same consecutive in list middle
	testList.Delete(2, true)

	if testList.tail.prev.value != 1 {
		fmt.Printf("Incorrect tail prev value")
		t.Error()
	}

	if reflect.DeepEqual(testList.tail.prev, testList.head) != true {
		fmt.Printf("Tail prev is not head")
		t.Error()
	}

	if reflect.DeepEqual(testList.head.next, testList.tail) != true {
		fmt.Printf("Head next is not tail")
		t.Error()
	}
}

func TestDeleteOneNode(t *testing.T) {
	// Test Delete of list of same values
	testList := LinkedList2{}

	node1 := Node{nil, nil, 1}
	node2 := Node{nil, nil, 2}
	node3 := Node{nil, nil, 3}
	node4 := Node{nil, nil, 4}
	node5 := Node{nil, nil, 5}
	node6 := Node{nil, nil, 6}

	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)
	testList.AddInTail(node4)
	testList.AddInTail(node5)
	testList.AddInTail(node6)
	// Delete head

	testList.Delete(1, false)

	if testList.head.value != 2 {
		fmt.Printf("Incorrect head value")
		t.Error()
	}

	if testList.head.next.value != 3 {
		fmt.Printf("Incorrect head next node value")
		t.Error()
	}

	if testList.head.next.prev.value != 2 {
		fmt.Printf("Incorrect head value")
		t.Error()
	}

	if reflect.DeepEqual(testList.head.next.prev, testList.head) != true {
		fmt.Printf("Head next prev is not head")
		t.Error()
	}

	// Delete tail

	testList.Delete(6, false)

	if testList.tail.value != 5 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.tail.next != nil {
		fmt.Printf("Tail next is not nil")
		t.Error()
	}

	if testList.tail.prev.next.value != 5 {
		fmt.Printf("Incorrect tail prev node next value")
		t.Error()
	}

	if reflect.DeepEqual(testList.tail.prev.next, testList.tail) != true {
		fmt.Printf("Tail prev next is not tail")
		t.Error()
	}

	// Delete second node
	testList.Delete(3, false)

	if testList.head.next.value != 4 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.head.next.prev.value != 2 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	if testList.tail.prev.value != 4 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}

	// Delete rest of nodes
	testList.Delete(2, false)
	testList.Delete(4, false)
	// Last node delition is also implicit Clean method test
	testList.Delete(5, false)

	if testList.head != nil {
		fmt.Printf("List head is not <nil")
		t.Error()
	}

	if testList.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}
}

func TestInsert(t *testing.T) {
	testList := LinkedList2{}

	node1 := Node{nil, nil,  1}
	node2 := Node{nil, nil,  2}
	node3 := Node{nil, nil,  3}
	nodeadd1 := Node{nil, nil,  777}
	nodeadd2 := Node{nil, nil,  888}
	nodeadd3 := Node{nil, nil,  999}
	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	// insert after head
	testList.Insert(testList.head, nodeadd1)

	if testList.head.next.value != 777 {
		fmt.Printf("head next value wrong")
		t.Error()
	}

	if testList.head.next.prev.value != 1 {
		fmt.Printf("head next node prev value wrong")
		t.Error()
	}

	if reflect.DeepEqual(testList.head.next.prev, testList.head) != true {
		fmt.Printf("head hext node prev not equal head")
		t.Error()
	}

	// insert after tail
	testList.Insert(testList.tail, nodeadd2)

	if testList.tail.value != 888 {
		fmt.Printf("tail value wrong")
		t.Error()
	}

	if testList.tail.next != nil {
		fmt.Printf("List tail next node is not <nil")
		t.Error()
	}

	if testList.tail.prev.next.value != 888 {
		fmt.Printf("tail prev node next value wrong")
		t.Error()
	}
	// insert after third node
	testList.Insert(testList.head.next.next, nodeadd3)

	if testList.head.next.next.next.value != 999 {
		fmt.Printf("wrong fourth node value")
		t.Error()
	}

	if testList.head.next.next.next.next.value != 3 {
		fmt.Printf("wrong fourth node next value")
		t.Error()
	}

	if testList.head.next.next.next.prev.value != 2 {
		fmt.Printf("wrong fourth node prev value")
		t.Error()
	}

	if testList.head.next.next.next.prev.next.value != 999 {
		fmt.Printf("wrong fourth node prev-next value")
		t.Error()
	}

	if testList.head.next.next.next.next.prev.value != 999 {
		fmt.Printf("wrong fourth node next-prev value")
		t.Error()
	}
}


func TestInserFirst(t *testing.T) {
	testList := LinkedList2{}

	node1 := Node{nil, nil,  1}
	node2 := Node{nil, nil,  2}
	node3 := Node{nil, nil,  3}
	nodeadd := Node{nil, nil,  777}

	testList.AddInTail(node1)
	testList.AddInTail(node2)
	testList.AddInTail(node3)

	testList.InsertFirst(nodeadd)

	fmt.Println(testList.GetListVals())

	if testList.head.value != 777 {
		fmt.Printf("head value wrong")
		t.Error()
	}
	if testList.head.next.value != 1 {
		fmt.Printf("head value wrong")
		t.Error()
	}

	if testList.head.next.prev.value != 777 {
		fmt.Printf("head next node prev value wrong")
		t.Error()
	}
}