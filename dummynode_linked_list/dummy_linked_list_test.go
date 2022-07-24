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
