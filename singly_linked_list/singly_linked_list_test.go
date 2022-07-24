package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIdiot(t *testing.T) {
	testList := LinkedList{}
	var node1 Node = Node{nil, 1}
	var node2 Node = Node{nil, 2}

	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	node1.value = 222

	fmt.Println(node1, testList.head)
}


// basic LinkedList creation

func TestLinkedListCreate(t *testing.T) {
	testList := LinkedList{}

	testList.AddInTail(&Node{nil, 1})
	testList.AddInTail(&Node{nil, 2})
	testList.AddInTail(&Node{nil, 3})

	if testList.head.value != 1 {
		t.Errorf("head value is incorrect, it is not 1 but %d", testList.head.value)
	}

	if testList.head.next.value != 2 {
		t.Errorf("head next node value is not correct. It is equal to %d, should be 2", testList.head.next.value)
	}

	if testList.tail.value != 3 {
		t.Errorf("wrong tail value - %d, should be 3", testList.tail.value)
	}

	// check if tail next is nil

	if testList.tail.next != nil {
		t.Errorf("something went wrong, tail next 'value' is not nil")
	}

}

func GotWantInt(got, want int) bool {
	if got != want {
		fmt.Printf("Got %d, but wanted %d", got, want)
		return false
	}
	return true
}

func TestLLCount(t *testing.T) {
	testList := LinkedList{}

	if GotWantInt(testList.Count(), 0) != true {
		t.Error()
	}

	testList.AddInTail(&Node{nil, 123})

	if GotWantInt(testList.Count(), 1) != true {
		t.Error()
	}
}

func TestFind(t *testing.T) {
	testList := LinkedList{}
	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 3}
	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)

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
	testList := LinkedList{}
	node1 := Node{nil, 1}
	node2 := Node{nil, 1}
	node3 := Node{nil, 1}
	node4 := Node{nil, 2}
	node5 := Node{nil, 3}
	node6 := Node{nil, 4}

	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)
	testList.AddInTail(&node4)
	testList.AddInTail(&node5)
	testList.AddInTail(&node6)

	someNodes := testList.FindAll(1)

	checkSlice := []Node{*testList.head, *testList.head.next, *testList.head.next.next}
	if reflect.DeepEqual(someNodes, checkSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	testList2 := LinkedList{}
	nodeX := Node{nil, 1}
	nodeY := Node{nil, 2}
	nodeZ := Node{nil, 2}

	testList2.AddInTail(&nodeX)
	testList2.AddInTail(&nodeY)
	testList2.AddInTail(&nodeZ)

	yetAnotherNodes := testList2.FindAll(2)
	yetAnotherSlice := []Node{*testList2.head.next, *testList2.tail}
	if reflect.DeepEqual(yetAnotherNodes, yetAnotherSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestDeleteOne(t *testing.T) {
	testList := LinkedList{}
	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 3}
	node4 := Node{nil, 4}
	node5 := Node{nil, 5}
	node6 := Node{nil, 6}

	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)
	testList.AddInTail(&node4)
	testList.AddInTail(&node5)
	testList.AddInTail(&node6)

	// Delete Head [1, 2, 3, 4, 5, 6]
	testList.Delete(1, false)
	if testList.head.value != 2 {
		fmt.Printf("Incorrect head value")
		t.Error()
	}

	if testList.head.next.value != 3 {
		fmt.Printf("Incorrect head next node value")
		t.Error()
	}
	// [2, 3, 4, 5, 6]
	testList.Delete(2, false)
	if testList.head.value != 3 {
		fmt.Printf("Incorrect head value")
		t.Error()
	}

	if testList.head.next.value != 4 {
		fmt.Printf("Incorrect head next node value")
		t.Error()
	}

	// Delete some in the middle [3, 4, 5, 6]
	testList.Delete(4, false)

	if testList.head.next.value != 5 {
		fmt.Printf("Incorrect head next node value")
		t.Error()
	}

	// Delete tail [3, 5, 6]
	testList.Delete(6, false)
	if testList.tail.value != 5 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}
	if testList.tail.next != nil {
		fmt.Printf("Tail next is not <nil>")
		t.Error()
	}

	testList.Delete(5, false)
	if testList.tail.value != 3 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}
	if testList.head.value != 3 {
		fmt.Printf("Incorrect head value")
		t.Error()
	}
	if testList.tail.next != nil {
		fmt.Printf("Tail next is not <nil>")
		t.Error()
	}
	if reflect.DeepEqual(testList.head, testList.tail) != true {
		fmt.Printf("Head is not a tail, but this is only one node in the LinkedList")
		t.Error()
	}

	// Delete last element [3]
	testList.Delete(3, false)
	if testList.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}

	if testList.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}

}

func TestDeleteAllSameHeadTail(t *testing.T) {
	// Test Delete of list of same values
	testList := LinkedList{}
	node1 := Node{nil, 6}
	node2 := Node{nil, 6}
	node3 := Node{nil, 2}
	node4 := Node{nil, 2}
	node5 := Node{nil, 1}
	node6 := Node{nil, 6}
	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)
	testList.AddInTail(&node4)
	testList.AddInTail(&node5)
	testList.AddInTail(&node6)

	// Test delete same consecutive in head and tail
	testList.Delete(6, true)

	if testList.head.value != 2 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}
	if testList.tail.value != 1 {
		fmt.Printf("Incorrect tail value")
		t.Error()
	}
	if testList.tail.next != nil {
		fmt.Printf("List tail next node is not <nil")
		t.Error()
	}
}

func TestDeleteAllSame(t *testing.T) {
	// Test Delete List with same consecutive
	testListSameVals := LinkedList{}
	node1 := Node{nil, 1}
	node2 := Node{nil, 1}
	node3 := Node{nil, 1}
	testListSameVals.AddInTail(&node1)
	testListSameVals.AddInTail(&node2)
	testListSameVals.AddInTail(&node3)

	testListSameVals.Delete(1, true)

	if testListSameVals.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}

	if testListSameVals.tail != nil {
		fmt.Printf("List tail is not <nil")
		t.Error()
	}
}

func TestDeleteAllSameMiddle(t *testing.T) {
	// Test Delete of list of same values
	testList := LinkedList{}

	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 2}
	node4 := Node{nil, 2}
	node5 := Node{nil, 2}
	node6 := Node{nil, 6}
	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)
	testList.AddInTail(&node4)
	testList.AddInTail(&node5)
	testList.AddInTail(&node6)
	// Test delete same consecutive in list middle
	testList.Delete(2, true)

}

func TestClean(t *testing.T) {
	testList := LinkedList{}
	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 3}
	node4 := Node{nil, 4}
	node5 := Node{nil, 5}
	node6 := Node{nil, 6}

	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)
	testList.AddInTail(&node4)
	testList.AddInTail(&node5)
	testList.AddInTail(&node6)

	testList.Clean()

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
	testList := LinkedList{}

	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 3}
	nodeadd := Node{nil, 777}
	nodeadd2 := Node{nil, 888}
	nodeadd3 := Node{nil, 999}
	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)

	testList.Insert(testList.head, nodeadd)

	if testList.head.next.value != 777 {
		fmt.Printf("head next value wrong")
		t.Error()
	}

	testList.Insert(testList.tail, nodeadd2)

	if testList.tail.value != 888 {
		fmt.Printf("tail value wrong")
		t.Error()
	}

	if testList.tail.next != nil {
		fmt.Printf("List tail next node is not <nil")
		t.Error()
	}

	testList.Insert(testList.head.next.next, nodeadd3)

	if testList.head.next.next.next.value != 999 {
		fmt.Printf("tail value wrong")
		t.Error()
	}

}

func TestInserFirst(t *testing.T) {
	testList := LinkedList{}

	node1 := Node{nil, 1}
	node2 := Node{nil, 2}
	node3 := Node{nil, 3}
	nodeadd := Node{nil, 777}

	testList.AddInTail(&node1)
	testList.AddInTail(&node2)
	testList.AddInTail(&node3)

	testList.InsertFirst(nodeadd)

	if testList.head.value != 777 {
		fmt.Printf("head value wrong")
		t.Error()
	}
	if testList.head.next.value != 1 {
		fmt.Printf("head value wrong")
		t.Error()
	}
}
