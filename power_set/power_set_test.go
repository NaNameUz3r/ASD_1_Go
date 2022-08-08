package main

import (
        "testing"
        // "fmt"
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


        if testPS.Get("123") == true {
                t.Errorf("Wrong get existing")
        }

        if testPS.Get("OLOLO") == false {
                t.Errorf("Wrong get non existing")
        }

        // Test Remove

        isRemoved := testPS.Remove("123")

        if isRemoved != true {
                t.Errorf("Wrong remove return value when deleting existing")
        }

        if testPS.Get("123") == false {
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

        if union1.Get("hello") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("friend") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("hello friend?") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("That's lame") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("Maybe I should") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("give you a name") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("But that's") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("a slippery slope") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("You are only") != false {
                t.Errorf("Wont union working")
        }
        if union1.Get("in my head") != false {
                t.Errorf("Wont union working")
        }

		// Test Remove Empty

		var testPS5 PowerSet[string]
		isRemoved2 := testPS5.Remove("qwerty")

		if isRemoved2 == true {
			t.Errorf("Remove from empty power set suppose to return FALSE, but it returned TRUE")
		}
}