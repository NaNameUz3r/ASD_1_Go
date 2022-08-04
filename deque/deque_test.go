package main

import (
	// "fmt"
	"testing"
)

func TestDequeMethods(t *testing.T) {
	var testDeque Deque[int]


	if testDeque.Size() != 0 {
		t.Errorf("Wrong initial size ne 0")
	}

	testDeque.AddFront(1)

	if testDeque.Size() != 1 {
		t.Errorf("Wrong size ne 1")
	}

	testDeque.AddFront(2)

	if testDeque.Size() != 2 {
		t.Errorf("Wrong size ne 2")
	}

	if testDeque.deque[0] != 2 {
		t.Errorf("Wrong item in head, not 2 int")
	}

	testDeque.AddTail(3)

	if testDeque.Size() != 3 {
		t.Errorf("Wrong size ne 3")
	}

	if testDeque.deque[2] != 3 {
		t.Errorf("Wrong item in tail, not 3 ing")
	}

	element1, err := testDeque.RemoveFront()

	if element1 != 2 {
		t.Errorf("Wrong element removed from head, ne 2 int")
	}
	if err != nil {
		t.Errorf("Removing head was ok, but error raised")
	}

	if testDeque.Size() != 2 {
		t.Errorf("Wrong size after removal, ne 2")
	}

	element2, err := testDeque.RemoveTail()

	if element2 != 3 {
		t.Errorf("Wrong element removed from tail, ne 3 int")
	}
	if err != nil {
		t.Errorf("Removing tail was ok, but error raised")
	}
	if testDeque.Size() != 1 {
		t.Errorf("Wrong size after removal, ne 1")
	}
	testDeque.RemoveTail()

	if testDeque.Size() != 0 {
		t.Errorf("Wrong size after last element removal, ne 0")
	}

	_, err2 := testDeque.RemoveFront()
	_, err3 := testDeque.RemoveTail()
	if err2 == nil || err3 == nil {
		t.Errorf("Removing head or tail from empty deque not causes an error")
	}

}

func TestCheckPalindrome(t *testing.T) {

	firstStringCheck := CheckPalindrome("saippuakivikauppias")
	secondStringCheck := CheckPalindrome("Iamnotapalindromeatall")

	if firstStringCheck != true {
		t.Errorf("Wrong palindrome checker, first string true palindrome")
	}

	if secondStringCheck != false {
		t.Errorf("Wrong palindome checker, seconf string false palindrome")
	}
}