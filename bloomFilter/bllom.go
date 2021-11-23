package main

import (
	"fmt"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	m int
	k int
	a []byte
}

func NewBloomFilter(m, k int) *BloomFilter {
	return &BloomFilter{m, k, make([]byte, m)}
}

func (bf *BloomFilter) Add(s string) {
	h := hash32(s) % bf.m
	bf.a[h] = 1
	h2 := hash64(s) % bf.m
	bf.a[h2] = 1

}

func hash32(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func hash64(s string) int {
	h := murmur3.New32()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func (bf *BloomFilter) Test(s string) bool {
	h := hash32(s) % bf.m
	if bf.a[h] == 0 {
		return false
	}
	h2 := hash64(s) % bf.m
	if bf.a[h2] == 0 {
		return false
	}
	return true
}

func main() {
	bf := NewBloomFilter(25, 5)
	bf.Add("hello")
	bf.Add("world")
	bf.Add("oy")
	fmt.Println(bf)
	fmt.Println(bf.Test("bro"))
	fmt.Println(bf.Test("hello"))
	fmt.Println(bf.Test("world"))
	fmt.Println(bf.Test("yo"))
}
