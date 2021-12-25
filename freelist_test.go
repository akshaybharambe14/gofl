package gofl

import (
	"reflect"
	"testing"
)

type testType []int

func (t testType) Reset() {
	for i := range t {
		t[i] = 0
	}
}

func TestFreeList_Get(t *testing.T) {
	newFn := func() testType {
		return make(testType, 4)
	}

	type args struct {
		fl FreeList[testType]
	}

	tests := []struct {
		name                 string
		args                 args
		expectedT            testType
		expectedFromFreeList bool
	}{
		{
			name:                 "with valid new function",
			args:                 args{fl: NewFreeList(1, newFn)},
			expectedT:            testType{0, 0, 0, 0},
			expectedFromFreeList: false,
		},
		{
			name:                 "without new function",
			args:                 args{fl: NewFreeList[testType](1, nil)},
			expectedT:            nil,
			expectedFromFreeList: false,
		},
		{
			name: "with existing T on free list",
			args: args{
				fl: func() FreeList[testType] {
					fl := NewFreeList(1, newFn)
					fl.put(testType{1, 2, 3, 4})
					return fl
				}(),
			},
			expectedT:            testType{0, 0, 0, 0},
			expectedFromFreeList: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualT, actualFromFreeList := tt.args.fl.get()

			if actualFromFreeList != tt.expectedFromFreeList {
				t.Errorf("get() actualFromFreeList = %v, expectedFromFreeList %v", actualFromFreeList, tt.expectedFromFreeList)
			}

			if !reflect.DeepEqual(actualT, tt.expectedT) {
				t.Errorf("get() actualT = %v, expectedT %v", actualT, tt.expectedT)
			}
		})
	}
}

func TestFreeList_Put(t *testing.T) {
	newFn := func() testType {
		return make(testType, 4)
	}

	type args struct {
		fl FreeList[testType]
		t  testType
	}

	tests := []struct {
		name          string
		args          args
		expectedAdded bool
	}{
		{
			name: "with no space available free list",
			args: args{
				fl: func() FreeList[testType] {
					fl := NewFreeList(1, newFn)
					fl.put(newFn())
					return fl
				}(),
				t: testType{0, 0, 0, 0},
			},
			expectedAdded: false,
		},
		{
			name: "with space available on free list",
			args: args{
				fl: NewFreeList(1, newFn),
				t:  testType{0, 0, 0, 0},
			},
			expectedAdded: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualAdded := tt.args.fl.put(tt.args.t)

			if !reflect.DeepEqual(actualAdded, tt.expectedAdded) {
				t.Errorf("Get() actualAdded = %v, expectedAdded %v", actualAdded, tt.expectedAdded)
			}
		})
	}
}
