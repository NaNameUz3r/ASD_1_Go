package main

import (
		"constraints"
		"fmt"
		"os"
		"strconv"
)

var CONSTANT_SIZE = 20000

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

func (ps *PowerSet[T]) Size() (int) {
		return ps.itemsCounter
}

func (ps *PowerSet[T]) Put(value T) {
		var hashIdx int
		var zeroType T
		hashIdx = ps.HashFun(value)
		if ps.slots[hashIdx] == zeroType {
			ps.itemsCounter++
			ps.slots[hashIdx] = value
		}

}

func (ps *PowerSet[T]) Get(value T) (bool) {
		var hashIdx = ps.HashFun(value)
		var zeroType T
		return ps.slots[hashIdx] == zeroType
}

func (ps *PowerSet[T]) Remove(value T) (bool) {
		var hashIdx = ps.HashFun(value)
		var zeroType T

		if ps.slots[hashIdx] != zeroType {
			ps.itemsCounter--
			ps.slots[hashIdx] = zeroType
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