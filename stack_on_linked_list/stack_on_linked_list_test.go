package stackLL

import (
        // "fmt"
        "testing"
)

func TestStackMethods(t *testing.T) {
        testStack := InitStack()

        if Size(testStack) != 0 {
                t.Errorf("Wrong initial stack size, ne 0")
        }
        _, err := Peek(testStack)

        if err == nil {
                t.Errorf("Error is nil wheen Peek empty stack")
        }

        Push(testStack, 1)

        item2, err2 := Peek(testStack)

        if item2 != 1 {
                t.Errorf("Wrong item Peeked")
        }
        if Size(testStack) != 1 {
                t.Errorf("Wrong stack size, ne 1")
        }

        if err2 != nil {
                t.Errorf("Error is not nil wheen Peek stack with elements")
        }

        Push(testStack, 2)
        Push(testStack, 3)
        Push(testStack, 4)

        if Size(testStack) != 4 {
                t.Errorf("Wrong stack Size, not 4, but 4 elements was pushed")
        }

        itemPoped, err3 := Pop(testStack)

        if itemPoped != 4 {
                t.Errorf("Wrong item Poped, not 4 int, but this is last pushed")
        }

        if err3 != nil {
                t.Errorf("There is some error when Poping from not empty stack")
        }

        Pop(testStack)
        Pop(testStack)

        if Size(testStack) != 1 {
                t.Errorf("Wrong stack Size, not 1, but only 1 elements should remain after Pops")
        }

        Pop(testStack)
        _, err4 := Pop(testStack)

        if err4 == nil {
                t.Errorf("Poping from empty stack not causing an error")
        }

}