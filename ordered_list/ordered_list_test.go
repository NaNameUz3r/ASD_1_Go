package main

import (
	"fmt"
	"testing"
)


func TestOrdereListAscMethods(t *testing.T) {
	var testOL OrderedList[int]
	testOL._ascending = true


	var node1 Node[int]
	node1.value = 1

	testOL.Add(&node1)
	fmt.Println(testOL.GetListVals())

	var node2 Node[int]
	node2.value = 9
	testOL.Add(&node2)
	fmt.Println(testOL.GetListVals())

	var node3 Node[int]
	node3.value = 5
	testOL.Add(&node3)
	fmt.Println(testOL.GetListVals())

	var node4 Node[int]
	node4.value = 15
	testOL.Add(&node4)
	fmt.Println(testOL.GetListVals())

	testOL.Delete(5)

	testOL.Clear(true)

	fmt.Println("ololo", testOL.GetListVals())

	somenode, err := testOL.Find(15)
	fmt.Println(somenode, node4, err)

}
