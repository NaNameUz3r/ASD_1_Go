package main

import (
	// "fmt"
	"testing"
)

func TestQueueStacks(t *testing.T) {
		var testQueue QueueStacks[int]

		if testQueue.Size() != 0 {
			t.Errorf("Wrong initial queue size, ne 0")
		}

		testQueue.Enqueue(1)
		testQueue.Enqueue(2)
		testQueue.Enqueue(3)

		if testQueue.Size() != 3 {
			t.Errorf("Wrong queue size, ne 3")
		}

		element1 := testQueue.Dequeue()

		if testQueue.Size() != 2 {
			t.Errorf("Wrong queue size, ne 2")
		}

		if element1 != 1 {
			t.Errorf("Wrong element dequeued, not 1")
		}

		element2 := testQueue.Dequeue()

		if testQueue.Size() != 1 {
			t.Errorf("Wrong queue size, ne 1")
		}

		if element2 != 2 {
			t.Errorf("Wrong element dequeued, not 2")
		}
}