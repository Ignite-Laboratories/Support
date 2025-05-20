package test

import (
	"github.com/ignite-laboratories/support/threadSafe"
	"testing"
)

func Test_ThreadSafe_Slice_NewSlice(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	s.Add(5)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 9
	})
	r := s.All()
	CompareSlices(r, []int{5, 7}, t)
}

func Test_ThreadSafe_Slice_Add(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	s.Add(5)
	s.Add(7)
	s.Add(9)
	r := s.All()
	CompareSlices(r, []int{5, 7, 9}, t)
}

func Test_ThreadSafe_Slice_RemoveIf(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 9
	})
	r := s.All()
	CompareSlices(r, []int{5, 7}, t)
}

func Test_ThreadSafe_Slice_RemoveIf_NoMatches(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)
	s.RemoveIf(func(i int) bool {
		return i == 1
	})
	r := s.All()
	CompareSlices(r, []int{5, 9, 7, 9}, t)
}

func Test_ThreadSafe_Slice_Length(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	CompareValues(s.Length(), 0, t)
	s.Add(5)
	CompareValues(s.Length(), 1, t)
	s.Add(9)
	CompareValues(s.Length(), 2, t)
	s.Add(7)
	CompareValues(s.Length(), 3, t)
	s.Add(9)
	CompareValues(s.Length(), 4, t)
}

func Test_ThreadSafe_Slice_Get(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	s.Add(5)
	s.Add(9)
	s.Add(7)
	s.Add(9)

	CompareValues(5, s.Get(0), t)
	CompareValues(9, s.Get(1), t)
	CompareValues(7, s.Get(2), t)
	CompareValues(9, s.Get(3), t)
}

func Test_ThreadSafe_Slice_All(t *testing.T) {
	s := threadSafe.NewSlice[int]()
	r := s.All()
	CompareSlices(r, []int{}, t)

	s.Add(5)
	s.Add(7)
	s.Add(9)
	r = s.All()
	CompareSlices(r, []int{5, 7, 9}, t)
}
