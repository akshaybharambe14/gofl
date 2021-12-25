/*
Package gofl provides a generic free list implementation. It aims to reuse existing initalized memory instead of allocating new memory(if not in provided limit).

Example:

	newFn := func() T {...}
	max := uint(10)
	fl := gofl.NewFreeList(max, newFn)
	...
	t := fl.Get() // get t from free list
	process(..., t) // reuse memory
	fl.Put(t) // put t back to free list; will be reused eventually

*/
package gofl // import "github.com/akshaybharambe14/gofl"
