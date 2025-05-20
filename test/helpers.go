package test

import (
	"testing"
)

// ShouldPanic fails the test if the test did not panic.
// It should be called at the start of your test with:
//
//	defer test.ShouldPanic(t)
func ShouldPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("Expected panic, but didn't get one")
		t.FailNow()
	}
}

// CompareValues fails the test if the two provided values are not equal.
func CompareValues[T comparable](a T, b T, t *testing.T) {
	if a != b {
		t.Errorf("Expected %v, got %v", a, b)
		t.FailNow()
	}
}

// CompareSlices fails the test if the two slices are unequal in length, or if the
// elements are not equal for every index.
func CompareSlices[T comparable](a []T, b []T, t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("Slices are not the same length")
		t.FailNow()
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("Expected %v at [%d], got %v", a[i], i, b[i])
			t.FailNow()
		}
	}
}

// CompareSlicesFunc fails the test if the two slices are unequal in length, or if the
// elements are not equal for every index.
//
// NOTE: This differs from CompareSlices in that it allows you to provide a custom comparison function.
func CompareSlicesFunc[T any](a []T, b []T, compare func(T, T, *testing.T), t *testing.T) {
	if len(a) != len(b) {
		t.Errorf("Slices are not the same length")
		t.FailNow()
	}
	for i := 0; i < len(a); i++ {
		compare(a[i], b[i], t)
	}
}
