package main

import (
		// "os"
		"fmt"
)

type Stack[T any] struct {
        stack []T
}

func (st *Stack[T]) Size() int {
        var size int
        size = len(st.stack)
        return size
}

func (st *Stack[T]) Peek() (T, error) {
        var result T
        var err error

        if st.Size() == 0 {
                err = fmt.Errorf("Peeking from empty stack")
        } else {
                result = st.stack[len(st.stack) - 1]
        }
        return result, err
}

func (st *Stack[T]) Pop() (T, error) {
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

func (st *Stack[T]) Push(itm T) {
		st.stack = append(st.stack, itm)
}

func isBalanced(checkstring string) bool {
        var stack Stack[string]
        for _, char := range checkstring {
                peek, _ := stack.Peek()
                stringedChar := string(char)
                if stringedChar == ")" && peek == "(" {
                        stack.Pop()
                } else {
                        stack.Push(stringedChar)
                }
        }

        return stack.Size() == 0
}