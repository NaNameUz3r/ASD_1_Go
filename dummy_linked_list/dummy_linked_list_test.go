package dummynode_linked_list

import (
	"fmt"
	"reflect"
	"testing"
	// "reflect"
)

func TestLinkedListCreateAndCount(t *testing.T) {
	testList := NewList()


	if testList.Head.Dummy != true {
		t.Errorf("Head is not Dummy")
	}
	if testList.Tail.Dummy != true {
		t.Errorf("Tail is not Dummy")
	}


	testList.AddInTail(Node{nil, nil, 1, false})
	testList.AddInTail(Node{nil, nil, 2, false})
	testList.AddInTail(Node{nil, nil, 3, false})

	if testList.Count() != 3 {
		t.Errorf("Count works wrong, there is 3 nodes, but Counted %d", testList.Count())
	}

	if testList.Head.Next.Value != 1 {
		t.Errorf("Head Next Value is wrong")
	}

	if testList.Tail.Prev.Value != 3 {
		t.Errorf("Tail Prev Value is wrong")
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

	if reflect.DeepEqual(&somenode, testList.Head.Next) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode2, err := testList.Find(2)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode2, testList.Head.Next.Next) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}

	somenode3, err := testList.Find(3)
	if err != nil {
		fmt.Printf("Error")
		t.Error()
	}

	if reflect.DeepEqual(&somenode3, testList.Tail.Prev) != true {
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

	checkSlice := []Node{*testList.Head.Next, *testList.Head.Next.Next, *testList.Head.Next.Next.Next}
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
	yetAnotherSlice := []Node{*testList2.Tail.Prev.Prev, *testList2.Tail.Prev}
	if reflect.DeepEqual(yetAnotherNodes, yetAnotherSlice) != true {
		fmt.Printf("Nodes are not equal")
		t.Error()
	}
}

func TestDeleteAllSameHeadTail(t *testing.T) {
	// Test Delete of list of same Values
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

	// Test delete same consecutive in Head and Tail
	testList.Delete(6, true)

	if testList.Head.Next.Value != 2 {
		fmt.Printf("Incorrect Tail Value")
		t.Error()
	}

	if testList.Head.Dummy != true {
		fmt.Printf("Head is not Dummy")
		t.Error()
	}

	if testList.Head.Next.Value != 2 {
		fmt.Printf("Incorrect Head Next.Next node Value")
		t.Error()
	}

	if testList.Head.Next.Next.Value != 2 {
		fmt.Printf("Incorrect Head Next.Next node Value")
		t.Error()
	}

	if testList.Head.Next.Next.Prev.Value != 2 {
		fmt.Printf("Incorrect Head Next.Next node Prev Value")
		t.Error()
	}


	if testList.Tail.Dummy != true {
		fmt.Printf("Tail is not Dummy")
		t.Error()
	}

	if testList.Tail.Prev.Value != 1 {
		fmt.Printf("Incorrect Tail Prev node Value")
		t.Error()
	}

	if testList.Tail.Next != nil {
		fmt.Printf("List Tail Next node is not <nil")
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

	if testListSameVals.Tail.Dummy != true {
		fmt.Printf("List Tail is not Dummy")
		t.Error()
	}

	if testListSameVals.Head.Dummy != true {
		fmt.Printf("List Tail is not Dummy")
		t.Error()
	}

	if testListSameVals.Head.Next.Dummy != true {
		fmt.Printf("Head Next is not Dummy")
		t.Error()
	}
}

func TestDeleteAllSameMiddle(t *testing.T) {
	// Test Delete of list of same Values
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


	if testList.Tail.Prev.Value != 6 {
		fmt.Printf("Incorrect Tail Prev Value")
		t.Error()
	}

	if testList.Head.Next.Value != 1 {
		fmt.Printf("Incorrect Head Next Value")
		t.Error()
	}

	if reflect.DeepEqual(testList.Tail.Prev, testList.Head.Next.Next) != true {
		fmt.Printf("Tail Prev is not Head Next")
		t.Error()
	}

	if reflect.DeepEqual(testList.Head.Next, testList.Tail.Prev.Prev) != true {
		fmt.Printf("Head Next is not Tail Prev Prev")
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

	// insert after Head Next
	testList.Insert(testList.Head.Next, nodeadd1)

	if testList.Head.Next.Next.Value != 777 {
		fmt.Printf("Head Next Value wrong")
		t.Error()
	}

	if testList.Head.Next.Value != 1 {
		fmt.Printf("Head Next node Value wrong")
		t.Error()
	}

	if reflect.DeepEqual(testList.Head.Next.Prev, testList.Head) != true {
		fmt.Printf("Head hext node Prev not equal DummyHead")
		t.Error()
	}

	// insert after actual Tail
	testList.Insert(testList.Tail.Prev, nodeadd2)

	if testList.Tail.Prev.Value != 888 {
		fmt.Printf("Tail Value wrong")
		t.Error()
	}

	if testList.Tail.Next != nil {
		fmt.Printf("List Tail Next node is not <nil>")
		t.Error()
	}

	if testList.Tail.Dummy != true {
		fmt.Printf(" Tail is not Dummy")
		t.Error()
	}



	if testList.Tail.Prev.Value != 888 {
		fmt.Printf("Tail Prev node Value wrong")
		t.Error()
	}
	// insert after third node
	testList.Insert(testList.Head.Next.Next.Next, nodeadd3)

	if testList.Head.Next.Next.Next.Next.Value != 999 {
		fmt.Printf("wrong fourth node Value")
		t.Error()
	}

	if testList.Head.Next.Next.Next.Next.Next.Value != 2 {
		fmt.Printf("wrong fourth node Next Value")
		t.Error()
	}

	if testList.Head.Next.Next.Next.Next.Next.Prev.Value != 999 {
		fmt.Printf("wrong fourth node Prev Value")
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

	if testList.Head.Next.Value != 777 {
		fmt.Printf("Head Value wrong")
		t.Error()
	}
	if testList.Head.Next.Next.Value != 1 {
		fmt.Printf("Head Value wrong")
		t.Error()
	}

	if testList.Head.Next.Next.Prev.Value != 777 {
		fmt.Printf("Head Next node Prev Value wrong")
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
	if testList.Head.Next.Dummy != true {
		fmt.Printf("Head Next node is not Dummy")
		t.Error()
	}
}