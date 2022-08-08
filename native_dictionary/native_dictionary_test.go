package main

import (
	"testing"
)


func TestNativeDict(t *testing.T) {

	var testND = Init[string](15)

	if testND.size != 15 {
		t.Errorf("Wrong initial size")
	}
	testND.Put("hello", "friend")
	testND.Put("hello friend?", "That's lame")
	testND.Put("Maybe I should", "give you a name")
	testND.Put("But that's", "a slippery slope")
	testND.Put("You are only", "in my head")

	// test IsKey

	if testND.IsKey("hello") != true {
		t.Errorf("Error! Not found existing key")
	}

	if testND.IsKey("We have to remember that. Shit.") != false {
		t.Errorf("IsKey with non existing key does not return false")
	}

	// test Get existing

	returnValue1, err2 := testND.Get("hello")

	if returnValue1 != "friend" {
		t.Errorf("Returned wrong value")
	}

	if err2 != nil {
		t.Errorf("Raised error trying get existing key")
	}

	// test Get non existing

	returnValue2, err2 := testND.Get("We have to remember that. Shit.")

	if returnValue2 != "" {
		t.Errorf("not empty string returned trying get non existing key")
	}

	if err2 == nil {
		t.Errorf("Not raising an error trying get non existing key")
	}
}