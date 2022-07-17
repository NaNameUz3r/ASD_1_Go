	package main

	import (
		"testing"
		"fmt"
	)

// basic LinkedList creation

func TestLinkedListCreate(t *testing.T){
	testList := LinkedList{}
	fmt.Println(testList)

	testList.AddInTail(Node{nil, 1})
	testList.AddInTail(Node{nil, 2})
	testList.AddInTail(Node{nil, 3})

	want 
}