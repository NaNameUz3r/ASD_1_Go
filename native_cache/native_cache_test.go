package main

import (
	"fmt"
	"testing"
)

func TestNativeCache(t *testing.T) {
	fmt.Println("TEST BEGIN")

	testNC := Init[string](10, 3)

	if testNC.size != 10 {
		t.Error("Wrong initial size")
	}

	testNC.Put("hello", "friend")
	testNC.Put("hello friend?", "That's lame")
	testNC.Put("Maybe I should", "give you a name")
	testNC.Put("But that's", "a slippery slope")
	testNC.Put("You are only", "in my head")

	get1, err1 := testNC.Get("hello")
	if get1 != "friend" {
		t.Errorf("Error! Not found existing key")
	}
	if err1 != nil {
		t.Errorf("value found, but error also raised")
	}

	if testNC.hits[testNC.getCacheIdx("hello")] != 1 {
		t.Errorf("Get func does not increase hit counter")
	}

	get2, err2 := testNC.Get("We have to remember that. Shit.")
	if get2 != "" {
		t.Errorf("Get with non existing key does not return false")
	}

	if err2 == nil {
		t.Errorf("Getting non existing key not raising an error")
	}

	checkGetVal, checkGetErr := testNC.Get("hello friend?")

	if checkGetErr != nil {
		t.Errorf("Trying get existing cached value return error")
	}

	if checkGetVal != "That's lame" {
		t.Errorf("Wrong existing cached value returned. That's lame")
	}

	get3, err3 := testNC.Get("hello")
	if get3 != "friend" {
		t.Errorf("Error! Not found existing key")
	}
	if err3 != nil {
		t.Errorf("value found, but error also raised")
	}

	if testNC.hits[testNC.getCacheIdx("hello")] != 2 {
		t.Errorf("Get func does not increase hit counter")
	}

	// LET FLOOD BEGIN!

	testNC.Put("foo", "bar")
	testNC.Put("chilly", "willy")
	testNC.Put("GNU", "LINUX")
	testNC.Put("FREE", "BSD")
	testNC.Put("DOOM", "GUY")

	if len(testNC.values) != 10 {
		t.Errorf("Cache not flooded enough for following tests")
	}

	for a := 0; a < 3; a++ { testNC.Get("foo") }
	for b := 0; b < 4; b++ { testNC.Get("chilly") }
	for c := 0; c < 5; c++ { testNC.Get("GNU") }
	for d := 0; d < 6; d++ { testNC.Get("FREE") }
	for e := 0; e < 7; e++ { testNC.Get("DOOM") }
	testNC.Get("Maybe I should")
	testNC.Get("But that's")



	if testNC.hits[testNC.getCacheIdx("foo")] != 3 {
		t.Errorf("Wrong hit counter for foo key")
	}

	if testNC.hits[testNC.getCacheIdx("DOOM")] != 7 {
		t.Errorf("Wrong hit counter for foo key")
	}

	testNC.Put("Everything is in the head", "of Master Mind")

	if testNC.hits[testNC.getCacheIdx("Everything is in the head")] != 0 {
		t.Errorf("Put func when purged other cached value by LRU dont purge old hits counter")
	}

	get4, err4 := testNC.Get("Everything is in the head")
	if get4 != "of Master Mind" {
		t.Errorf("Error! Not found existing key")
	}
	if err4 != nil {
		t.Errorf("value found, but error also raised")
	}

	// for i, _ := range testNC.slots {
	// 	fmt.Println(testNC.slots[i], "\t" ,testNC.values[i], "\t", testNC.hits[i])
	// }

	if testNC.hits[testNC.getCacheIdx("Everything is in the head")] != 1 {
		t.Errorf("Get func does not increase hit counter")
	}

	get5, err5 := testNC.Get("You are only")

	if get5 != "" {
		t.Errorf("Get of non existing (purged by LRU) key ne zeroValue")
	}

	if err5 == nil {
		t.Errorf("Get of non existing (purged by LRU) key not raising an error")
	}

}