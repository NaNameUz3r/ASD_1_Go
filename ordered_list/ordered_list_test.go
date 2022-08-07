package main

import (
	"testing"
	// "fmt"
)


func TestOrdereListAscMethods(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = true



	testOL.Add(1)

	if testOL.head.value != 1 || testOL.tail.value != 1 {
		t.Errorf("Wrong first node addition, something wrong with head or tail pointers")
	}

	if testOL.Count() != 1 {
		t.Errorf("Wrong list count, ne 1")
	}
	testOL.Add(9)

	if testOL.head.value != 1 || testOL.tail.value != 9 {
		t.Errorf("List elements ordered incorrectly, ne [1 9]")
	}

	if testOL.Count() != 2 {
		t.Errorf("Wrong list count, ne 2")
	}

	testOL.Add(5)

	if testOL.head.value != 1 || testOL.tail.value != 9 || testOL.head.next.value != 5 {
		t.Errorf("List asc elements ordered incorrectly, ne [1 5 9]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	testOL.Add(15)

	if testOL.head.value != 1 || testOL.tail.prev.value != 9 || testOL.head.next.value != 5 || testOL.tail.value != 15 {
		t.Errorf("List asc elements ordered incorrectly, ne [1 5 9 15]")
	}

	if testOL.Count() != 4 {
		t.Errorf("Wrong list count, ne 4")
	}

	find1, err1 := testOL.Find(1)

	if find1.value != 1 || find1 != *testOL.head {
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


	testOL.Add(1)

	if testOL.head.value != 1 || testOL.tail.value != 1 {
		t.Errorf("Wrong first node addition, something wrong with head or tail pointers")
	}

	if testOL.Count() != 1 {
		t.Errorf("Wrong list count, ne 1")
	}
	testOL.Add(9)

	if testOL.head.value != 9 || testOL.tail.value != 1 {
		t.Errorf("List elements ordered incorrectly, ne [9 1]")
	}

	if testOL.Count() != 2 {
		t.Errorf("Wrong list count, ne 2")
	}

	testOL.Add(5)

	if testOL.head.value != 9 || testOL.tail.value != 1 || testOL.head.next.value != 5 {
		t.Errorf("List asc elements ordered incorrectly, ne [9 5 1]")
	}

	if testOL.Count() != 3 {
		t.Errorf("Wrong list count, ne 3")
	}


	testOL.Add(15)

	if testOL.head.value != 15 || testOL.tail.prev.value != 5 || testOL.head.next.value != 9 || testOL.tail.value != 1 {
		t.Errorf("List asc elements ordered incorrectly, ne [15 9 5 1]")
	}

	if testOL.Count() != 4 {
		t.Errorf("Wrong list count, ne 4")
	}

	find1, err1 := testOL.Find(1)

	if &find1 != testOL.tail && find1.value != 1 {
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


func TestDescOLDeletion(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = false


	testOL.Add(128)
	testOL.Add(256)
	testOL.Add(512)

	if testOL.head.value != 512 || testOL.head.next.value != 256 || testOL.tail.value != 128 {
		t.Errorf("Something went wrong when adding in Asc ordered list, ne [512 256 128]")
	}
	// Delete non existin element

	testOL.Delete(777)

	if testOL.head.value != 512 || testOL.head.next.value != 256 || testOL.tail.value != 128 {
		t.Errorf("Something went wrong deleting non existing element from list, ne [512 256 128] after deletion")
	}

	// Delete existing element, lowest value in tail
	testOL.Delete(128)

	if testOL.head.value != 512 || testOL.tail.value != 256 {
		t.Errorf("Something went wrong deleting existing lowest element from list, ne [256 512] after deletion")
	}

	// Delete existing element, max value from head

	testOL.Delete(512)

	if testOL.head.value != 256 || testOL.tail.value != 256 {
		t.Errorf("Something went wrong deleting existing max element from list, ne [256] after deletion")
	}


	// Delete last element

	testOL.Delete(256)

	if testOL.head != nil || testOL.tail != nil {
		t.Errorf("Something went wrong deleting last element from list, ne [] after deletion")
	}

}

func TestAscOLDeletion(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = true


	testOL.Add(128)
	testOL.Add(256)
	testOL.Add(512)

	if testOL.head.value != 128 || testOL.head.next.value != 256 || testOL.tail.value != 512 {
		t.Errorf("Something went wrong when adding in Asc ordered list, ne [128 256 512]")
	}
	// Delete non existin element

	testOL.Delete(777)

	if testOL.head.value != 128 || testOL.head.next.value != 256 || testOL.tail.value != 512 {
		t.Errorf("Something went wrong deleting non existing element from list, ne [128 256 512] after deletion")
	}

	// Delete existing element, lowest value in head
	testOL.Delete(128)

	if testOL.head.value != 256 || testOL.tail.value != 512 {
		t.Errorf("Something went wrong deleting existing lowest element from list, ne [256 512] after deletion")
	}

	// Delete existing element, max value from tail

	testOL.Delete(512)

	if testOL.head.value != 256 || testOL.tail.value != 256 {
		t.Errorf("Something went wrong deleting existing max element from list, ne [256] after deletion")
	}


	// Delete last element

	testOL.Delete(256)

	if testOL.head != nil || testOL.tail != nil {
		t.Errorf("Something went wrong deleting last element from list, ne [] after deletion")
	}

}
