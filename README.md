Package `xsort` contains manually inlined versions of the "search wrappers" in
the Go standard `sort` library (`SearchInts`, `SearchFloat64s`,
`SearchStrings`).

In the standard library, these are convenience wrappers around the generic
`sort.Search()` function, which takes a function parameter to determine
truthfulness. However, since this function is utilized within a for loop, it
[cannot currently be inlined by the Go compiler][1], resulting in non-trivial
performance overhead.

Some quick single threaded benchmarks on 10M element slices on my laptop:
```
$ go test -cpu=1 -bench=.
goos: darwin
goarch: arm64
pkg: github.com/mroth/xsort
BenchmarkSearchInts/xsort         	86329382	        13.85 ns/op
BenchmarkSearchInts/sort          	17089969	        70.16 ns/op
BenchmarkSearchFloat64s/xsort     	75447772	        15.81 ns/op
BenchmarkSearchFloat64s/sort      	12603822	        95.44 ns/op
BenchmarkSearchStrings/xsort      	 9170595	       131.3 ns/op
BenchmarkSearchStrings/sort       	 6044085	       197.3 ns/op
PASS
ok  	github.com/mroth/xsort	9.248s
```

[1]: https://github.com/golang/go/issues/15561
