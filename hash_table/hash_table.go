package main

import (
        // "os"
        // "strconv"
)

type HashTable struct {
        size  int
        step  int
        slots []string
}

func Init(sz int, stp int) HashTable {
        ht := HashTable{size: sz, step: stp, slots: nil}
        ht.slots = make([]string, sz)
        return ht
}

func (ht *HashTable) HashFun(value string) int {
		var hashIdx int
		var hashSum int
		for _, runeval := range value {
			hashSum = hashSum + int(runeval)
		}
		hashIdx = hashSum % ht.size
        return hashIdx
}

func (ht *HashTable) SeekSlot(value string) int {
		hashIdx := ht.HashFun(value)

		if ht.slots[hashIdx] == "" {
			return hashIdx
		}

		seekIdx := 0

		for {
			if seekIdx >= ht.size {
				break
			}

			hashIdx = hashIdx + ht.step

			if hashIdx >= ht.size {
				hashIdx = hashIdx - ht.size
			}

			if ht.slots[hashIdx] == "" {
				return hashIdx
			} else {
				seekIdx = seekIdx + 1
			}
		}
		hashIdx = -1
		return hashIdx
}

func (ht *HashTable) Put(value string) int {
		var slotIdx = ht.SeekSlot(value)

		if slotIdx != -1 {
			ht.slots[slotIdx] = value
		}

		return slotIdx
}

func (ht *HashTable) Find(value string) int {
		var checkHashIdx = ht.HashFun(value)

		if ht.slots[checkHashIdx] == value {
			return checkHashIdx
		}

		findIdx := 0

		for {
			if findIdx >= ht.size {
				break
			}

			checkHashIdx += ht.step
			if checkHashIdx >= ht.size {
				checkHashIdx = checkHashIdx - ht.size
			}

			if ht.slots[checkHashIdx] == value {
				return checkHashIdx
			} else {
				findIdx = findIdx + 1
			}
		}

		checkHashIdx = -1
		return checkHashIdx

}