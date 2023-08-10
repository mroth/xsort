//go:build go1.21
// +build go1.21

package xsort_test

import (
	"fmt"
	"slices"
	"sort"
	"testing"

	"github.com/mroth/xsort"
)

const n = 10_000_000

func BenchmarkSearchInts(b *testing.B) {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = i
	}

	if !testing.Short() {
		b.Run("lib=sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchInts(xs, n-1)
			}
		})

		b.Run("lib=slices", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slices.BinarySearch(xs, n-1)
			}
		})

	}

	b.Run("lib=xsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchInts(xs, n-1)
		}
	})

}

func BenchmarkSearchFloat64s(b *testing.B) {
	xs := make([]float64, n)
	for i := range xs {
		xs[i] = float64(i)
	}

	if !testing.Short() {
		b.Run("lib=sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchFloat64s(xs, n-1)
			}
		})

		b.Run("lib=slices", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slices.BinarySearch(xs, n-1)
			}
		})
	}

	b.Run("lib=xsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchFloat64s(xs, n-1)
		}
	})
}

func BenchmarkSearchStrings(b *testing.B) {
	xs := make([]string, n)
	for i := range xs {
		xs[i] = fmt.Sprint(i)
	}
	x := fmt.Sprint(n - 1)

	if !testing.Short() {
		b.Run("lib=sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchStrings(xs, x)
			}
		})

		b.Run("lib=slices", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slices.BinarySearch(xs, x)
			}
		})
	}

	b.Run("lib=xsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchStrings(xs, x)
		}
	})
}
