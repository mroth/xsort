Package xsort contains manually inlined versions of the "search wrappers" in the
Go standard `sort` library (`SearchInts`, `SearchFloat64s`, `SearchStrings`).

In the standard library, these are convenience wrappers around the generic
`sort.Search()` function, which takes a function parameter to determine
truthfulness. However, since this function is utilized within a for loop, it
cannot currently be inlined by the Go compiler, resulting in non-trivial
performance overhead.

Some quick single threaded benchmarks on 10M element slices on my workstation:
```
$ sysctl -n machdep.cpu.brand_string
Intel(R) Xeon(R) W-2140B CPU @ 3.20GHz

$ go test -cpu=1 -bench=.           
goos: darwin
goarch: amd64
pkg: github.com/mroth/xsort
BenchmarkSearchInts/pkgxsort            60457520                19.1 ns/op
BenchmarkSearchInts/sort                19617564                60.7 ns/op
BenchmarkSearchFloat64s/pkgxsort        47584767                24.7 ns/op
BenchmarkSearchFloat64s/sort            19664412                60.7 ns/op
BenchmarkSearchStrings/pkgxsort          5624664               219 ns/op
BenchmarkSearchStrings/sort              4832034               250 ns/op
PASS
ok      github.com/mroth/xsort  11.255s
```
