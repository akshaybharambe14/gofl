# gofl

GOFL provides a Generic Free List implementation for Go.

## Installation

This package requires go1.18 or later.

```sh
go get github.com/akshaybharambe14/gofl
```

## Why?

Free list pattern provides a way to reuse already initialized memory. GOFL helps you to grab an item from the free list and put it back to the free list once it is no longer needed. This helps in allocating limited amount of memory, resulting in better performance.

You can have a look at [leaky buffer example from effective go documentation](https://go.dev/doc/effective_go). This package is inspired by the same implementation.

A general usage would be:

```go
newFn := func() T {...}
max := uint(10)
fl := gofl.NewFreeList(max, newFn)
...
t := fl.Get() // get t from free list
process(..., t) // reuse memory
fl.Put(t) // put t back to free list; will be reused eventually
```

Following benchmarks depict the use of free list in finding duplicates in a slice.

```sh
$ github.com/akshaybharambe14/gofl/example> go1.18beta1 test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/akshaybharambe14/gofl/example
cpu: Intel(R) Core(TM) i5-4200M CPU @ 2.50GHz
BenchmarkFindDuplicate-4                 1693132               681.7 ns/op           162 B/op          1 allocs/op
BenchmarkFindDuplicateWithFreeList-4     2151346               554.3 ns/op             2 B/op          0 allocs/op
PASS
ok      github.com/akshaybharambe14/gofl/example        3.855s
```

As you can see, it avoids creation and allocations of temporary map, resulting in better performance.

For more details see [example](example/main_test.go).

## Contact

[Akshay Bharambe](https://twitter.com/akshaybharambe1)

---

If this is not something you are looking for, you can check other similar packages on [pkg.go.dev](https://pkg.go.dev/), [go.libhunt.com](https://go.libhunt.com).

Do let me know if you have any feedback. Leave a ⭐ this helps you!
