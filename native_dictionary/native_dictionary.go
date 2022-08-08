package main

import (
		"fmt"
		"os"
		"strconv"
)

type NativeDictionary[T any] struct {
		size int
		slots []string
		values []T
}

// создание экземпляра словаря
func Init[T any](sz int) NativeDictionary[T] {
		nd := NativeDictionary[T] { size : sz, slots : nil, values : nil }
		nd.slots = make([]string, sz)
		nd.values = make([]T, sz)
		return nd
}


func (nd *NativeDictionary[T]) HashFun(value string) (int) {
		var hashIdx int
		var hashSum int
		for _, runeval := range value {
			hashSum = hashSum + int(runeval)
		}
		hashIdx = hashSum % nd.size
		return hashIdx
}

func (nd *NativeDictionary[T]) IsKey(key string) (bool) {
		return nd.slots[nd.HashFun(key)] == key
}

func (nd *NativeDictionary[T]) Get(key string) (T, error) {
		var result T
		var err error

		if nd.IsKey(key) == true {
			itemIdx := nd.HashFun(key)
			return nd.values[itemIdx], err
		}

		err = fmt.Errorf("No such key in dict")
		return result, err
}

func (nd *NativeDictionary[T]) Put(key string, value T) {
		var newValueIdx = nd.HashFun(key)
		nd.slots[newValueIdx], nd.values[newValueIdx] = key, value
}