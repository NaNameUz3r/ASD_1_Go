package stackLL

import (
		// "os"
		"fmt"
		ll "ADS/dummy_linked_list"
)

type Stack = ll.LinkedListDummy

func InitStack() Stack {
        return *ll.NewList()
}

func Size(st Stack) int {
        return st.Count()
}

func Peek(st Stack) (interface{}, error) {
        var result interface{}
        var err error

        if Size(st) == 0 {
                err = fmt.Errorf("Peeking from empty stack")
        } else {
                result = st.GetHeadValue()
        }
        return result, err
}

func Pop(st Stack) (interface{}, error) {
        var result interface{}
        var err error

        if Size(st) == 0 {
                err = fmt.Errorf("Poping from empty stack")
        } else {
                result = st.Head.Next.Value
                st.DeleteHead()
        }
        return result, err
}

func Push(st Stack, value interface{}) {
        n := ll.Node{Prev: nil, Next: nil, Value: value, Dummy: false}
        st.InsertFirst(n)
}