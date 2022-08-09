package main

import (
        "testing"
        // "fmt"
        // "reflect"
)

func TestPowerSet(t *testing.T) {

        var testPS PowerSet[string]

        // test HashFun

        val1 := testPS.HashFun("friend")

        if val1 != 632 {
                t.Errorf("Hash func doesnot work")
        }

        val2 := testPS.HashFun("hello")

        if val2 != 532 || val1 == val2 {
                t.Errorf("Hash func doesnot work")
        }

        // Test Put and Size

        testPS.Put("123")

        if testPS.slots[150] != "123" {
                t.Errorf("Wrong put")
        }

        if testPS.Size() != 1 {
                t.Errorf("Wrong size")
        }

        // Test Get

        if testPS.Get("123") != true {
                t.Errorf("Wrong get existing")
        }

        if testPS.Get("OLOLO") != false {
                t.Errorf("Wrong get non existing")
        }

        // Test Remove

        isRemoved := testPS.Remove("123")

        if isRemoved != true {
                t.Errorf("Wrong remove return value when deleting existing")
        }

        if testPS.Size() != 0 {
                t.Errorf("Wrong size after removing element")
        }
        if testPS.Get("123") != false {
                t.Errorf("Wrong get deleted value (non existing)")
        }

        // Test Intersect

        testPS.Put("hello")
        testPS.Put("friend")
        testPS.Put("hello friend?")
        testPS.Put("That's lame")
        testPS.Put("Maybe I should")
        testPS.Put("give you a name")
        testPS.Put("But that's")
        testPS.Put("a slippery slope")
        testPS.Put("You are only")
        testPS.Put("in my head")

        // what about size?

        if testPS.Size() != 10 {
                t.Errorf("Wrong size!")
        }
        var testPS2 PowerSet[string]

        testPS2.Put("hello")
        testPS2.Put("a slippery slope")
        testPS2.Put("in my head")

        intersect1 := testPS.Intersection(testPS2)

        if intersect1.Get("hello") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }
        if intersect1.Get("a slippery slope") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }
        if intersect1.Get("in my head") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }

        intersect2 := testPS2.Intersection(testPS)

        if intersect2.Get("hello") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }
        if intersect2.Get("a slippery slope") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }
        if intersect2.Get("in my head") == true {
                t.Errorf("Wrong intersection bigger with smaller")
        }

        // Test Union

        var testPS3 PowerSet[string]
        testPS3.Put("hello")
        testPS3.Put("friend")
        testPS3.Put("hello friend?")
        testPS3.Put("That's lame")
        testPS3.Put("Maybe I should")

        var testPS4 PowerSet[string]
        testPS4.Put("give you a name")
        testPS4.Put("But that's")
        testPS4.Put("a slippery slope")
        testPS4.Put("You are only")
        testPS4.Put("in my head")

        union1 := testPS3.Union(testPS4)

        // for _, item := range union1.slots {
        //         if union1.Get(item) {
        //                 fmt.Println(item)
        //         }
        // }

        if union1.Get("hello") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("friend") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("hello friend?") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("That's lame") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("Maybe I should") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("give you a name") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("But that's") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("a slippery slope") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("You are only") != true {
                t.Errorf("Wont union working")
        }
        if union1.Get("in my head") != true {
                t.Errorf("Wont union working")
        }

        // Test Remove Empty

        var testPS5 PowerSet[string]
        isRemoved2 := testPS5.Remove("qwerty")

        if testPS5.Size() != 0 {
                t.Errorf("Wrong initial size")
        }

        if isRemoved2 == true {
                t.Errorf("Remove from empty power set suppose to return FALSE, but it returned TRUE")
        }
}

func TestGetFromEmptyPS(t *testing.T) {
        var testPS6 PowerSet[string]

        if testPS6.Get("qwerty") != false {
                t.Errorf("Wrong get non existing value from empty PowerSet, ne FALSE")
        }
}

func TestSizePS(t *testing.T) {
        var testPS7 PowerSet[string]

        if testPS7.Size() != 0 {
                t.Errorf("Wrong initial size ne 0")
        }


        testPS7.Put("qwerty")

        if testPS7.Size() != 1 {
                t.Errorf("Wrong size after put first element, ne 1")
        }

        testPS7.Put("ytrewq")


        if testPS7.Size() != 2 {
                t.Errorf("Wrong size after put seconf element, ne 2")
        }

        // And Remove...

        testPS7.Remove("qwerty")
        // for _, item := range testPS7.slots {
        //         itemIdx := testPS7.Find(item)
        //         if itemIdx > -1 && itemIdx != 0 {
        //                 fmt.Println(testPS7.slots[itemIdx], itemIdx)
        //         }
        // }

        if testPS7.Size() != 1 {
                t.Errorf("Wrong size after deletion one element, ne 1")
        }
}