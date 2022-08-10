package main

import (
        "testing"
)

func TestBloomFilterHashes(t *testing.T) {
        testtestBF := BloomFilter{}
        testtestBF.filter_len = 32

        if testtestBF.filter_len != 32 {
                t.Error("Wrong initial filter length")
        }

        checkHash1 := testtestBF.Hash1("0123456789")

        if checkHash1 != 13 {
                t.Error("Wrong Hash1 function result, ne 13")
        }

        checkHash2 := testtestBF.Hash2("0123456789")

        if checkHash2 != 5 {
                t.Error("Wrong Hash2 function result, ne 5")
        }
}

func TestBloomFilterAddValueCheckValue(t *testing.T) {
        var testBF = BloomFilter{}
        testBF.filter_len = 32

        testBF.Add("0123456789")
        testBF.Add("1234567890")
        testBF.Add("2345678901")
        testBF.Add("3456789012")
        testBF.Add("4567890123")
        testBF.Add("5678901234")

        if testBF.filter_len != 32 {
                t.Error("Wrong filter length")
        }

        var check = testBF.IsValue("2222222222")
        if check != false {
                t.Error("2222222222 isValue eq True, but value was not added")
        }

        check = testBF.IsValue("5678901234")
        if check != true {
                t.Error("5678901234 is value ne True, but value was added")
        }

        check = testBF.IsValue("4567890123")
        if check != true {
                t.Error("4567890123 is value ne True, but value was added")
        }

        check = testBF.IsValue("3456789012")
        if check != true {
                t.Error("3456789012 is value ne True, but value was added")
        }

        check = testBF.IsValue("2345678901")
        if check != true {
                t.Error("2345678901 is value ne True, but value was added")
        }

        check = testBF.IsValue("1234567890")
        if check != true {
                t.Error("1234567890 is value ne True, but value was added")
        }

        check = testBF.IsValue("0123456789")
        if check != true {
                t.Error("0123456789 is value ne True, but value was added")
        }
}