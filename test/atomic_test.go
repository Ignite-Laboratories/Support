package test

import (
	"github.com/ignite-laboratories/core/test"
	"github.com/ignite-laboratories/support/atomic"
	"testing"
)

func Test_Atomic_Slice_NewSlice(t *testing.T) {
	s := atomic.NewSlice[int]()
	s.Add(5)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 9
	})
	r := s.All()
	test.CompareSlices(r, []int{5, 7}, t)
}

func Test_Atomic_Slice_Add(t *testing.T) {
	s := atomic.NewSlice[int]()
	s.Add(5)
	s.Add(7)
	s.Add(9)
	r := s.All()
	test.CompareSlices(r, []int{5, 7, 9}, t)
}

func Test_Atomic_Slice_RemoveIf(t *testing.T) {
	s := atomic.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 9
	})
	r := s.All()
	test.CompareSlices(r, []int{5, 7}, t)
}

func Test_Atomic_Slice_RemoveIf_NoMatches(t *testing.T) {
	s := atomic.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 1
	})
	r := s.All()
	test.CompareSlices(r, []int{5, 9, 7, 9}, t)
}

func Test_Atomic_Slice_Length(t *testing.T) {
	s := atomic.NewSlice[int]()
	test.CompareValues(s.Length(), 0, t)
	s.Add(5)
	test.CompareValues(s.Length(), 1, t)
	s.Add(9)
	test.CompareValues(s.Length(), 2, t)
	s.Add(7)
	test.CompareValues(s.Length(), 3, t)
	s.Add(9)
	test.CompareValues(s.Length(), 4, t)
}

func Test_Atomic_Slice_Get(t *testing.T) {
	s := atomic.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)

	test.CompareValues(5, s.Get(0), t)
	test.CompareValues(9, s.Get(1), t)
	test.CompareValues(7, s.Get(2), t)
	test.CompareValues(9, s.Get(3), t)
}

func Test_Atomic_Slice_All(t *testing.T) {
	s := atomic.NewSlice[int]()
	r := s.All()
	test.CompareSlices(r, []int{}, t)

	s.Add(5)
	s.Add(7)
	s.Add(9)
	r = s.All()
	test.CompareSlices(r, []int{5, 7, 9}, t)
}
