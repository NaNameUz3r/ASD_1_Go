package main

import (
        // "os"
)

type BloomFilter struct {
        filter_len int
        bitmask    uint32
}

func (bf *BloomFilter) hash(s string, salt int) int {
        hashSum := 0
        for _, char := range s {
                code := int(char)
                hashSum = (hashSum*salt + code) % bf.filter_len
        }
        return hashSum
}

func (bf *BloomFilter) Hash1(s string) int {
        return bf.hash(s, 17)
}

func (bf *BloomFilter) Hash2(s string) int {
        return bf.hash(s, 223)
}

func (bf *BloomFilter) Add(s string) {
        hashSum1, hashSum2 := uint32(bf.Hash1(s)), uint32(bf.Hash2(s))
        bitIdx := hashSum1 % uint32(bf.filter_len)

        bf.bitmask |= 1 << uint(bitIdx&0x2C)
        bitIdx = hashSum2 % uint32(bf.filter_len)
        bf.bitmask |= 1 << uint(bitIdx&0x2C)
}

func (bf *BloomFilter) IsValue(s string) bool {
        hashSum1, hashSum2 := uint32(bf.Hash1(s)), uint32(bf.Hash2(s))
        bitIdx1, bitIdx2 := hashSum1%uint32(bf.filter_len), hashSum2%uint32(bf.filter_len)

        isValueIdx2 := (bf.bitmask >> uint(bitIdx2&0x2C)) & 1
        if isValueIdx2 == 0 {
                return false
        }

        isValueIdx1 := (bf.bitmask >> uint(bitIdx1&0x2C)) & 1
        if isValueIdx1 == 0 {
                return false
        }

        return true
}