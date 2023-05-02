//go:build go1.18
// +build go1.18

// Package xsort contains manually inlined versions of the "search wrappers"
// in the standard sort library.
//
// In the standard library, these are convenience wrappers around the generic
// `sort.Search()` function, which takes a function parameter to determine
// truthfulness. However, since this function is utilized within a for loop, it
// cannot currently be inlined by the Go compiler, resulting in non-trivial
// performance overhead.
package xsort

import "golang.org/x/exp/constraints"

// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchInts(a []int, x int) int {
	return searchComparable(a, x)
}

// SearchFloat64s searches for x in a sorted slice of float64s and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchFloat64s(a []float64, x float64) int {
	return searchComparable(a, x)
}

// SearchStrings searches for x in a sorted slice of strings and returns the index
// as specified by Search. The return value is the index to insert x if x is not
// present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchStrings(a []string, x string) int {
	return searchComparable(a, x)
}

// searchComparable searches for x in a sorted slice of constraints.Ordered and
// returns the index as specified by Search. The return value is the index to
// insert x if x is not present (it could be len(a)).
//
// The slice must be sorted in ascending order.
func searchComparable[T constraints.Ordered](a []T, x T) int {
	i, j := 0, len(a)
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
