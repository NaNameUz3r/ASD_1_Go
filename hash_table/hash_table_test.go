package main

import (
	"testing"
)

func TestHashFun(t *testing.T) {
	tsh := Init(17, 3)
	hashIdx1 := tsh.HashFun("HELLO")
	if hashIdx1 != 15 || hashIdx1 > tsh.size || hashIdx1 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx2 := tsh.HashFun("friend of mine")
	if hashIdx2 != 8 || hashIdx2 > tsh.size || hashIdx2 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx3 := tsh.HashFun("dont blame dogs please")
	if hashIdx3 != 1 || hashIdx3 > tsh.size || hashIdx3 < 0 {
		t.Error("Wrong hash index or out of range")
	}
}

func TestPut(t *testing.T) {
	tsh := Init(17, 3)
	hashIdx1 := tsh.Put("HELLO")
	if hashIdx1 != 15 || hashIdx1 > tsh.size || hashIdx1 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx2 := tsh.Put("friend of mine")
	if hashIdx2 != 8 || hashIdx2 > tsh.size || hashIdx2 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx3 := tsh.Put("dont blame dogs please")
	if hashIdx3 != 1 || hashIdx3 > tsh.size || hashIdx3 < 0 {
		t.Error("Wrong hash index or out of range")
	}

}


func TestSeek(t *testing.T) {
	tsh := Init(17, 3)
	hashIdx1 := tsh.SeekSlot("HELLO")
	if hashIdx1 != 15 || hashIdx1 > tsh.size || hashIdx1 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx2 := tsh.SeekSlot("friend of mine")
	if hashIdx2 != 8 || hashIdx2 > tsh.size || hashIdx2 < 0 {
		t.Error("Wrong hash index or out of range")
	}

	hashIdx3 := tsh.SeekSlot("dont blame dogs please")
	if hashIdx3 != 1 || hashIdx3 > tsh.size || hashIdx3 < 0 {
		t.Error("Wrong hash index or out of range")
	}

}


func TestPutAndFind(t *testing.T) {

	tsh := Init(17, 3)
	tsh.Put("HELLO")
	tsh.Put("qawerqweDFSDF;zxlkcjv;laksdjfrqweqwerrqwer")
	tsh.Put("qawerqwerqweqwe12341234rrqwer")
	tsh.Put("qawerqwerqweqwerrqwasdfasdfer")
	tsh.Put("qawasdgasgerqwerqweqwerrqwer")
	tsh.Put("qawerqwerqzxcvzxcvweqwerrqwer")
	tsh.Put("qawerqwerqweqwerrqfasdfgwer")
	tsh.Put("qawerqwerqweqwerrqwertttttttttttrrrrrrrrewwwwwwwwqqq")
	tsh.Put("qawerqwerqweqwerrqqwerwqerwer")
	tsh.Put("Hello, friend. Hello, friend? That's lame. Maybe I should give you a name, but that's a slippery slope. You're only in my head.")

	findResult := tsh.Find("Hello, friend. Hello, friend? That's lame. Maybe I should give you a name, but that's a slippery slope. You're only in my head.")

	if findResult == -1 {
		t.Errorf("Something went wrong, Find does not Find string that should be found")
	}
}