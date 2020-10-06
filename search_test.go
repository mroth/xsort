package xsort_test

import (
	"fmt"
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

	b.Run("pkgxsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchInts(xs, n-1)
		}
	})

	if !testing.Short() {
		b.Run("sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchInts(xs, n-1)
			}
		})
	}
}

func BenchmarkSearchFloat64s(b *testing.B) {
	xs := make([]float64, n)
	for i := range xs {
		xs[i] = float64(i)
	}

	b.Run("pkgxsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchFloat64s(xs, n-1)
		}
	})

	if !testing.Short() {
		b.Run("sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchFloat64s(xs, n-1)
			}
		})
	}
}

func BenchmarkSearchStrings(b *testing.B) {
	xs := make([]string, n)
	for i := range xs {
		xs[i] = fmt.Sprint(i)
	}

	b.Run("pkgxsort", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			xsort.SearchStrings(xs, fmt.Sprint(n-1))
		}
	})

	if !testing.Short() {
		b.Run("sort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sort.SearchStrings(xs, fmt.Sprint(n-1))
			}
		})
	}
}
