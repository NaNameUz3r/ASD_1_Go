package stackLL

import (
		// "os"
		"fmt"
		ll "ADS/dummynode_linked_list"
)

type Stack[T any] struct {
        stack func()
}

func InitStack() *Stack {
	s := Stack{
		stack: func() {ll.NewList()},
	}

}
func (st *Stack[T]) Size() int {
        return st.stack.Count()
}

func (st *Stack[T]) Peek() (T, error) {
        var result T
        var err error

        if st.Size() == 0 {
                err = fmt.Errorf("Peeking from empty stack")
        } else {
                result = st.stack.head
        }
        return result, err
}

func (st *Stack) Pop() (T, error) {
        var result T
        var err error

        if st.Size() == 0 {
                err = fmt.Errorf("Poping from empty stack")
        } else {
                result = st.stack[len(st.stack) - 1]
                st.stack = st.stack[:len(st.stack) - 1]
        }
        return result, err
}

func (st *Stack) Push(itm T) {
		st.stack = append(st.stack, itm)
}