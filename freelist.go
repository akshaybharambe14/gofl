package gofl

type Resetter interface {
	// Reset tries to keep the memory allocated and safe to reuse.
	Reset()
}

// NewFn should create a new instance of T.
type NewFn[T Resetter] func() T

// FreeList represents a list of T that can be reused. It is safe for concurrent use but it does not protect underlying T from concurrent access.
//
// Zero value is not usable, use gofl.NewFreeList() to create a new instance.
type FreeList[T Resetter] struct {
	buffer chan T
	newFn  NewFn[T]
	zero   T
}

// NewFreeList creates a new free list with given capacity and newFn.
// If newFn is nil then zero value of T will be returned on Get operation.
func NewFreeList[T Resetter](max uint, f NewFn[T]) FreeList[T] {
	return FreeList[T]{
		buffer: make(chan T, max),
		newFn:  f,
	}
}

// Get returns a T from the free list. It creates a new instance of T with provided newFn, if not available.
func (fl FreeList[T]) Get() T {
	t, _ := fl.get()
	return t
}

func (fl FreeList[T]) get() (T, bool) {
	select {
	case v := <-fl.buffer:
		return v, true
	default:
		if fl.newFn == nil {
			return fl.zero, false
		}
		return fl.newFn(), false
	}
}

// Put returns a T to the free list after calling Reset(). T will be dropped if list is full.
func (fl FreeList[T]) Put(v T) {
	_ = fl.put(v)
}

func (fl FreeList[T]) put(v T) bool {
	v.Reset() // TODO: wasted efforts if buffer is full

	select {
	case fl.buffer <- v:
		return true
	default:
		// drop item; buffer is full
		return false
	}
}
