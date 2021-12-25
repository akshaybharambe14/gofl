package main

import (
	"fmt"

	"github.com/akshaybharambe14/gofl"
)

type temp map[int]struct{}

// interface guard
var _ gofl.Resetter = temp{}

// Reset T so that it can be reused
func (t temp) Reset() {
	for k := range t {
		delete(t, k)
	}
}

func main() {
	f := func() temp {
		return make(temp, 15)
	}
	fl := gofl.NewFreeList(1, f) // maintain 1 item in the free list
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	found := findDuplicateWithFreeList(arr, fl.Get())
	fmt.Println("found:", found)

	// now reuse temporary memory with free list
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // no duplicates
	found = findDuplicateWithFreeList(arr, fl.Get() /* gets existing map */)
	fmt.Println("found:", found) // should be false this time
}

func findDuplicate(arr []int) bool {
	t := make(temp, len(arr)/2)
	for _, v := range arr {
		if _, ok := t[v]; ok {
			return true
		}
		t[v] = struct{}{}
	}
	return false
}

func findDuplicateWithFreeList(arr []int, t temp) bool {
	for _, v := range arr {
		if _, ok := t[v]; ok {
			return true
		}
		t[v] = struct{}{}
	}
	return false
}
