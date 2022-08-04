package main

import (
	// "fmt"
	"testing"
	"bytes"
)

func TestQueueMethods(t *testing.T) {
		var testQueue Queue[int]

		if testQueue.Size() != 0 {
			t.Errorf("Wrong initial queue size, ne 0")
		}
		_ , err := testQueue.Dequeue()

		if err == nil {
			t.Errorf("Error is nil wheen dequeue empty queue")
		}

		testQueue.Enqueue(1)
		testQueue.Enqueue(2)
		testQueue.Enqueue(3)

		if testQueue.Size() != 3 {
			t.Errorf("Wrong queue size, ne 3")
		}

		element1, err1 := testQueue.Dequeue()

		if testQueue.Size() != 2 {
			t.Errorf("Wrong queue size, ne 2")
		}

		if element1 != 1 {
			t.Errorf("Wrong element dequeued, not 1")
		}

		if err1 != nil {
			t.Errorf("Dequeue was ok, but error returned")
		}
}

func TestSpinQueue(t *testing.T) {
	var testQueue Queue[byte]
	testQueue.Enqueue(1)
	testQueue.Enqueue(2)
	testQueue.Enqueue(3)
	testQueue.Enqueue(4)
	testQueue.Enqueue(5)
	testQueue.Enqueue(6)

	testQueue.SpinQueue(2)

	reference := []byte{3, 4, 5, 6, 1, 2}
	res := bytes.Compare(reference, testQueue.queue)
	if res != 0 {
		t.Errorf("Wrong spin func work")
	}
}