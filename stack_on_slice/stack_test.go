package main

import (
	// "fmt"
	"testing"
)

func TestStackMethods(t *testing.T) {
		var testStack Stack[int]

		if testStack.Size() != 0 {
			t.Errorf("Wrong initial stack size, ne 0")
		}
		_ , err := testStack.Peek()

		if err == nil {
			t.Errorf("Error is nil wheen Peek empty stack")
		}


		testStack.Push(1)

		item2, err2 := testStack.Peek()

		if item2 != 1 {
			t.Errorf("Wrong item Peeked")
		}
		if testStack.Size() != 1 {
			t.Errorf("Wrong stack size, ne 1")
		}

		if err2 != nil {
			t.Errorf("Error is not nil wheen Peek stack with elements")
		}

		testStack.Push(2)
		testStack.Push(3)
		testStack.Push(4)

		if testStack.Size() != 4 {
			t.Errorf("Wrong stack Size, not 4, but 4 elements was pushed")
		}

		itemPoped, err3 := testStack.Pop()

		if itemPoped != 4 {
			t.Errorf("Wrong item Poped, not 4 int, but this is last pushed")
		}

		if err3 != nil {
			t.Errorf("There is some error when Poping from not empty stack")
		}

		testStack.Pop()
		testStack.Pop()

		if testStack.Size() != 1 {
			t.Errorf("Wrong stack Size, not 1, but only 1 elements should remain after Pops")
		}

		testStack.Pop()
		_, err4 := testStack.Pop()

		if err4 == nil {
			t.Errorf("Poping from empty stack not causing an error")
		}


}
