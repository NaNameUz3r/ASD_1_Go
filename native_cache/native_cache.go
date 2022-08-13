package main

import (
        "fmt"
        "os"
)

type NativeCache[T any] struct {
        size   int
        step   int
        slots  []string
        values []T
        hits   []int
}

func Init[T any](sz int, stp int) NativeCache[T] {
        nc := NativeCache[T]{size: sz, step: stp, slots: nil}
        nc.slots = make([]string, sz)
        nc.values = make([]T, sz)
        nc.hits = make([]int, sz)
        return nc
}

func (nc *NativeCache[T]) HashFun(value string) int {
        var hashIdx int
        var hashSum int
        for _, runeval := range value {
                hashSum = hashSum + int(runeval)
        }
        hashIdx = hashSum % nc.size
        return hashIdx
}

func (nc *NativeCache[T]) SeekEmptySlot(value string) int {
        hashIdx := nc.HashFun(value)

        if nc.slots[hashIdx] == "" {
                return hashIdx
        }

        seekIdx := 0

        for {
                if seekIdx >= nc.size {
                        break
                }

                hashIdx = hashIdx + nc.step

                if hashIdx >= nc.size {
                        hashIdx = hashIdx - nc.size
                }

                if nc.slots[hashIdx] == "" {
                        return hashIdx
                } else {
                        seekIdx = seekIdx + 1
                }
        }
        hashIdx = -1
        return hashIdx
}

func (nc *NativeCache[T]) FindLruSlotIdx() int {
        var minHittedIdx = 0
        var minHitted = nc.hits[0]

        for hitIdx, hitsCounter := range nc.hits {
                if minHitted > hitsCounter {
                        minHitted = hitsCounter
                        minHittedIdx = hitIdx
                }
        }

        return minHittedIdx

}

func (nc *NativeCache[T]) Put(key string, value T) {
        var slotIdx = nc.SeekEmptySlot(key)

        if slotIdx == -1 {
                slotIdx = nc.FindLruSlotIdx()
                nc.hits[slotIdx] = 0
        }

        nc.slots[slotIdx] = key
        nc.values[slotIdx] = value
}

func (nc *NativeCache[T]) Get(key string) (T, error) {
        var checkHashIdx = nc.HashFun(key)
        var err error
        var zeroType T

        if nc.slots[checkHashIdx] == key {
                nc.hits[checkHashIdx]++
                return nc.values[checkHashIdx], err
        }

        findIdx := 0

        for {
                if findIdx >= nc.size {
                        break
                }

                checkHashIdx += nc.step
                if checkHashIdx >= nc.size {
                        checkHashIdx = checkHashIdx - nc.size
                }

                if nc.slots[checkHashIdx] == key {
                        nc.hits[checkHashIdx]++
                        return nc.values[checkHashIdx], err
                } else {
                        findIdx = findIdx + 1
                }
        }

        err = fmt.Errorf("no such key in cache")
        return zeroType, err

}

func (nc *NativeCache[T]) IsKey(key string) bool {
	return nc.getCacheIdx(key) > -1
}

func (nc *NativeCache[T]) getCacheIdx(key string) int {
        var checkHashIdx = nc.HashFun(key)

        if nc.slots[checkHashIdx] == key {
                return checkHashIdx
        }

        findIdx := 0

        for {
                if findIdx >= nc.size {
                        break
                }

                checkHashIdx += nc.step
                if checkHashIdx >= nc.size {
                        checkHashIdx = checkHashIdx - nc.size
                }

                if nc.slots[checkHashIdx] == key {
                        return checkHashIdx
                } else {
                        findIdx = findIdx + 1
                }
        }

        checkHashIdx = -1
        return checkHashIdx

}