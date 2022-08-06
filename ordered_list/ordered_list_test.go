package main

import (
	"testing"
)


func TestOrdereListAscMethods(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = true


	var node1 Node[int]
	node1.value = 1

	testOL.Add(&node1)

	if testOL.head != &node1 || testOL.tail != &node1 {
		t.Errorf("Wrong first node addition, something wrong with head or tail pointers")
	}

	if testOL.Count() != 1 {
		t.Errorf("Wrong list count, ne 1")
	}
	var node2 Node[int]
	node2.value = 9
	testOL.Add(&node2)

	if testOL.head.value != 1 || testOL.tail.value != 9 {
		t.Errorf("List elements ordered incorrectly, ne [1 9]")
	}

	if testOL.Count() != 2 {
		t.Errorf("Wrong list count, ne 2")
	}

	var node3 Node[int]
	node3.value = 5
	testOL.Add(&node3)

	if testOL.head.value != 1 || testOL.tail.value != 9 || testOL.head.next.value != 5 {
		t.Errorf("List asc elements ordered incorrectly, ne [1 5 9]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	var node4 Node[int]
	node4.value = 15
	testOL.Add(&node4)

	if testOL.head.value != 1 || testOL.tail.prev.value != 9 || testOL.head.next.value != 5 || testOL.tail.value != 15 {
		t.Errorf("List asc elements ordered incorrectly, ne [1 5 9 15]")
	}

	if testOL.Count() != 4 {
		t.Errorf("Wrong list count, ne 4")
	}

	find1, err1 := testOL.Find(1)

	if node1 != find1 {
		t.Errorf("Wrong find method working, finded node should ve eq to head")
	}
	if err1 != nil {
		t.Errorf("Node found, but error was raised.")
	}

	_, err2 := testOL.Find(999)
	if err2 == nil {
		t.Errorf("Node not found, but error was not raised.")
	}


	testOL.Delete(5)
	if testOL.head.value != 1 || testOL.tail.value != 15 || testOL.head.next.value != 9 {
		t.Errorf("List asc elements ordered incorrectly after node.value 5 deletion, ne [1 9 15]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	testOL.Delete(1)
	testOL.Delete(9)
	testOL.Delete(15)

	_, err3 := testOL.Find(15)

	if err3 == nil {
		t.Errorf("Attempt to find something in empty list not rasing an error")
	}
	if testOL.head != nil || testOL.tail != nil {
		t.Errorf("Wrong all elements delition, head or tail is not nil")
	}

	if testOL._ascending != true {
		t.Errorf("Wrong _ascending Odereder List field value after all elements one by one deletion. By default this value should stay the same initialized.")
	}
	if testOL.Count() != 0 {
		t.Errorf("Wrong list count, ne 0")
	}

}

func TestOrdereListDescMethods(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = false


	var node1 Node[int]
	node1.value = 1

	testOL.Add(&node1)

	if testOL.head != &node1 || testOL.tail != &node1 {
		t.Errorf("Wrong first node addition, something wrong with head or tail pointers")
	}

	if testOL.Count() != 1 {
		t.Errorf("Wrong list count, ne 1")
	}
	var node2 Node[int]
	node2.value = 9
	testOL.Add(&node2)

	if testOL.head.value != 9 || testOL.tail.value != 1 {
		t.Errorf("List elements ordered incorrectly, ne [9 1]")
	}

	if testOL.Count() != 2 {
		t.Errorf("Wrong list count, ne 2")
	}

	var node3 Node[int]
	node3.value = 5
	testOL.Add(&node3)

	if testOL.head.value != 9 || testOL.tail.value != 1 || testOL.head.next.value != 5 {
		t.Errorf("List asc elements ordered incorrectly, ne [9 5 1]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	var node4 Node[int]
	node4.value = 15
	testOL.Add(&node4)

	if testOL.head.value != 15 || testOL.tail.prev.value != 5 || testOL.head.next.value != 9 || testOL.tail.value != 1 {
		t.Errorf("List asc elements ordered incorrectly, ne [15 9 5 1]")
	}

	if testOL.Count() != 4 {
		t.Errorf("Wrong list count, ne 4")
	}

	find1, err1 := testOL.Find(1)

	if node1 != find1 {
		t.Errorf("Wrong find method working, finded node should ve eq to head")
	}
	if err1 != nil {
		t.Errorf("Node found, but error was raised.")
	}

	_, err2 := testOL.Find(999)
	if err2 == nil {
		t.Errorf("Node not found, but error was not raised.")
	}


	testOL.Delete(5)
	if testOL.head.value != 15 || testOL.tail.value != 1 || testOL.head.next.value != 9 {
		t.Errorf("List asc elements ordered incorrectly after node.value 5 deletion, ne [15 9 1]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	testOL.Delete(1)
	testOL.Delete(9)
	testOL.Delete(15)

	_, err3 := testOL.Find(15)

	if err3 == nil {
		t.Errorf("Attempt to find something in empty list not rasing an error")
	}
	if testOL.head != nil || testOL.tail != nil {
		t.Errorf("Wrong all elements delition, head or tail is not nil")
	}

	if testOL._ascending != false {
		t.Errorf("Wrong _ascending Odereder List field value after all elements one by one deletion. By default this value should stay the same initialized.")
	}
	if testOL.Count() != 0 {
		t.Errorf("Wrong list count, ne 0")
	}

}
