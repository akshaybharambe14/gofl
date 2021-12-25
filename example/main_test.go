package main

import (
	"testing"

	"github.com/akshaybharambe14/gofl"
)

func BenchmarkFindDuplicate(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	for i := 0; i < b.N; i++ {
		findDuplicate(arr)
	}
}

func BenchmarkFindDuplicateWithFreeList(b *testing.B) {
	f := func() temp {
		return make(temp, 15)
	}
	fl := gofl.NewFreeList(uint(b.N)/4, f)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	for i := 0; i < b.N; i++ {
		t := fl.Get()
		findDuplicateWithFreeList(arr, t)
		fl.Put(t)
	}
}

func Test_findDuplicateWithFreeList(t *testing.T) {
	f := func() temp {
		return make(temp, 15)
	}
	fl := gofl.NewFreeList(1, f)

	type args struct {
		t   temp
		arr []int
	}
	tests := []struct {
		cleanup func(t temp)
		name    string
		args    args
		want    bool
	}{
		{
			name: "duplicate",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1},
				t:   fl.Get(),
			},
			want: true,
			cleanup: func(t temp) {
				fl.Put(t)
			},
		},
		{
			name: "reused T should return correct result",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				t:   fl.Get(),
			},
			want: false,
			cleanup: func(t temp) {
				fl.Put(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateWithFreeList(tt.args.arr, tt.args.t); got != tt.want {
				t.Errorf("findDuplicateWithFreeList() = %v, want %v", got, tt.want)
			}

			if tt.cleanup != nil {
				tt.cleanup(tt.args.t)
			}
		})
	}
}
