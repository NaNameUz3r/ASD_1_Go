package main

import (
        "testing"
)

func TestAppendAndReallocate(t *testing.T) {
        var testDynArr DynArray[int]

        testDynArr.Init()

        if testDynArr.capacity != 16 {
                t.Errorf("incorrect initial DynArr initialization, capacity is %d, but should be 16", testDynArr.capacity)
        }

        if testDynArr.count != 0 {
                t.Errorf("incorrect initial DynArr elements counter (not 0)")
        }

        for i := 1; i <= 16; i++ {
                testDynArr.Append(i)
                if testDynArr.array[i-1] != i {
                        t.Errorf("Wrong values of %d indexed element, shoul be %d", i-1, i)
                }
        }

        // test reallocate

        testDynArr.Append(999)

        if testDynArr.count != 17 {
                t.Errorf("Wrong counter value, should be 17, but %d", testDynArr.count)
        }

        if testDynArr.capacity != 32 {
                t.Errorf("Wrong counter value, should be 17, but %d", testDynArr.capacity)
        }
}

func TestRestOfMethodsAndCapacityReallocation(t *testing.T) {
        var testDynArr DynArray[int]

        testDynArr.Init()

        if testDynArr.capacity != 16 {
                t.Errorf("Wrong capacity value, should be 16, but %d", testDynArr.capacity)
        }

        for i := 1; i <= 16; i++ {
                testDynArr.Append(i)
                if testDynArr.array[i-1] != i {
                        t.Errorf("Wrong values of %d indexed element, shoul be %d", i-1, i)
                }
        }

        testDynArr.Append(999)

        if testDynArr.capacity != 32 {
                t.Errorf("Wrong capacity value, should be 32, but %d", testDynArr.capacity)
        }
        if testDynArr.count != 17 {
                t.Errorf("Wrong counter value, should be 17, but %d", testDynArr.count)
        }

        testDynArr.Remove(0)

        if testDynArr.array[0] != 2 {
                t.Errorf("Wrong index 0 item value, should be 2, but %d", testDynArr.array[0])
        }
        if testDynArr.array[testDynArr.count-1] != 999 {
                t.Errorf("Wrong index 0 item value, should be 999, but %d", testDynArr.array[0])
        }

        testDynArr.Remove(0)

        if testDynArr.capacity != 21 {
                t.Errorf("Wrong capacity value, should be 21, but %d", testDynArr.capacity)
        }
        if testDynArr.count != 15 {
                t.Errorf("Wrong counter value, should be 15, but %d", testDynArr.count)
        }

        testDynArr.Remove(0)
        testDynArr.Remove(0)
        testDynArr.Remove(0)
        testDynArr.Remove(0)
        testDynArr.Remove(0)

        if testDynArr.capacity != 16 {
                t.Errorf("Wrong capacity value, should be 21, but %d", testDynArr.capacity)
        }
        if testDynArr.count != 10 {
                t.Errorf("Wrong counter value, should be 15, but %d", testDynArr.count)
        }

        if testDynArr.array[0] != 8 {
                t.Errorf("Wrong index 0 item value, should be 2, but %d", testDynArr.array[0])
        }
        if testDynArr.array[testDynArr.count-1] != 999 {
                t.Errorf("Wrong index 0 item value, should be 999, but %d", testDynArr.array[testDynArr.count-1])
        }

        getValue, err := testDynArr.GetItem(9)
        if err != nil {
                t.Errorf("Catch error while getting value")

        }
        if testDynArr.array[testDynArr.count-1] != getValue {
                t.Errorf("Wrong imp of GetItem Method")
        }

        getValue2, err := testDynArr.GetItem(3)
        if err != nil {
                t.Errorf("Catch error while getting value")

        }
        if testDynArr.array[3] != getValue2 {
                t.Errorf("Wrong imp of GetItem Method")
        }

}