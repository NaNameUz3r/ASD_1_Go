package main

import (
		"constraints"
		"fmt"
		// "os"
		// "strconv"
)

var CONSTANT_SIZE = 20000
var CONSTANT_STEP = 3

type PowerSet[T constraints.Ordered] struct {
		slots []T
		itemsCounter int
}

func (ps *PowerSet[T]) HashFun(value T) int {
		var hashIdx int
		var hashSum int

		if len(ps.slots) == 0 {
			ps.slots = make([]T, CONSTANT_SIZE)
		}

		processValue := fmt.Sprintf("%v", value)
		for _, runeval := range processValue {
			hashSum = hashSum + int(runeval)
		}
		hashIdx = hashSum % CONSTANT_SIZE
        return hashIdx
}

func (ps *PowerSet[T]) SeekSlot(value T) int {
		var hashIdx int
		hashIdx = ps.HashFun(value)
		var zeroType T

		if ps.slots[hashIdx] == zeroType {
			return hashIdx
		}

		seekIdx := 0

		for {
			if seekIdx >= CONSTANT_SIZE {
				break
			}

			hashIdx = hashIdx + CONSTANT_STEP

			if hashIdx >= CONSTANT_SIZE {
				hashIdx = hashIdx - CONSTANT_SIZE
			}

			if ps.slots[hashIdx] == zeroType {
				return hashIdx
			} else {
				seekIdx = seekIdx + 1
			}
		}
		hashIdx = -1
		return hashIdx
}

func (ps *PowerSet[T]) Size() (int) {
		return ps.itemsCounter
}

func (ps *PowerSet[T]) Put(value T) {
		var slotIdx int
		slotIdx = ps.SeekSlot(value)

		if slotIdx > -1 {
			ps.itemsCounter++
			ps.slots[slotIdx] = value
		}
}

func (ps *PowerSet[T]) Find(value T) int {
		var checkHashIdx int
		checkHashIdx = ps.HashFun(value)

		if ps.slots[checkHashIdx] == value {
			return checkHashIdx
		}

		findIdx := 0

		for {
			if findIdx >= CONSTANT_SIZE {
				break
			}

			checkHashIdx += CONSTANT_STEP
			if checkHashIdx >= CONSTANT_SIZE {
				checkHashIdx = checkHashIdx - CONSTANT_SIZE
			}

			if ps.slots[checkHashIdx] == value {
				return checkHashIdx
			} else {
				findIdx = findIdx + 1
			}
		}

		checkHashIdx = -1
		return checkHashIdx
}


func (ps *PowerSet[T]) Get(value T) (bool) {
		return ps.Find(value) > -1
}

func (ps *PowerSet[T]) Remove(value T) (bool) {
		var findIdx = ps.Find(value)
		var zeroType T

		if findIdx > -1 {
			ps.itemsCounter--
			ps.slots[findIdx] = zeroType
			return true
		}
		return false
}

func (ps *PowerSet[T]) Intersection(set2 PowerSet[T]) (PowerSet[T]) {
		var resultPs PowerSet[T]

	 	if ps.Size() >= set2.Size() {
			for _, item := range set2.slots {
				if !ps.Get(item) {
					resultPs.Put(item)
				}
			}
		} else if set2.Size() >= ps.Size() {
			for _, item := range ps.slots {
				if !set2.Get(item) {
					resultPs.Put(item)
				}
			}
		}
		return resultPs
}

func (ps *PowerSet[T]) Union(set2 PowerSet[T]) (PowerSet[T]) {
		// объединение текущего множества и set2
		var resultPs PowerSet[T]

		for _, item := range ps.slots {
			resultPs.Put(item)
		}

		for _, item := range set2.slots {
			resultPs.Put(item)
		}
		return resultPs
}

func (ps *PowerSet[T]) Difference(set2 PowerSet[T]) (PowerSet[T]) {
		var resultPs PowerSet[T]

		for _, item := range ps.slots {
			if set2.Get(item) == false {
				resultPs.Put(item)
			}
		}

		return resultPs
}

func (ps *PowerSet[T]) IsSubset(set2 PowerSet[T]) (bool) {

		for _, item := range set2.slots {
			if ps.Get(item) == false {
				return false
			}
		}
		return true
}