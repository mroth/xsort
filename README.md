# xsort

[![Go Reference](https://pkg.go.dev/badge/github.com/mroth/xsort.svg)](https://pkg.go.dev/github.com/mroth/xsort)

Manually inlined versions of the "search wrappers" in the Go standard `sort`
library (`SearchInts`, `SearchFloat64s`, `SearchStrings`), which perform up to
80% faster. Usage is identical.

> [!IMPORTANT]
> This module is no longer necessary and should be considered deprecated for use in go1.24 and greater.

## Why?

In the standard library, these are convenience wrappers around the generic
`sort.Search()` function, which takes a function parameter to determine
truthfulness. However, since this function is utilized within a for loop,
[it cannot currently be inlined by the Go compiler][1], resulting in non-trivial
performance overhead.

## Performance (go1.20)

Some quick single threaded benchmarks on 10M element slices on my laptop:
```
$ go test -bench=. -count=10 -cpu=1 | benchstat -col="/lib@(std xsort)" -
goos: darwin
goarch: arm64
pkg: github.com/mroth/xsort
               │     std      │                xsort                │
               │    sec/op    │   sec/op     vs base                │
SearchInts        60.47n ± 0%   13.75n ± 1%  -77.25% (p=0.000 n=10)
SearchFloat64s    83.52n ± 0%   13.77n ± 0%  -83.51% (p=0.000 n=10)
SearchStrings    136.05n ± 0%   69.17n ± 0%  -49.16% (p=0.000 n=10)
geomean           88.24n        23.57n       -73.28%
```

![benchmark chart](docs/benchchart-go1.20-min.svg)

## :new: Updated performance for go 1.21+

> [!NOTE]
> * **<= go1.20**: `xsort` provides significant performance advantages over `sort`.
> * **go1.21**: `slices.BinarySearch` provides competitive performance to `xsort`, with the exception of `SearchStrings`, `sort` is still unoptimized.
> * **go1.22-1.23**: `slices.BinarySearch` provides roughly competitive performance to `xsort` across all types, but `sort` is still unoptimized.
> * **go1.24**: `sort` is now optimized, and provides competitive performance to `xsort`.  _`xsort` can now be deprecated._

### Go 1.18

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │     sort     │               xsort                │
               │    sec/op    │   sec/op     vs base               │
SearchInts       74.00n ±  7%   23.15n ± 1%  -68.72% (p=0.002 n=6)
SearchFloat64s   71.07n ± 13%   17.62n ± 8%  -75.20% (p=0.002 n=6)
SearchStrings    141.1n ±  6%   100.1n ± 0%  -29.03% (p=0.002 n=6)
geomean          90.52n         34.44n       -61.96%
```

### Go 1.19

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │    sort     │               xsort                │
               │   sec/op    │   sec/op     vs base               │
SearchInts       68.91n ± 9%   16.25n ± 5%  -76.42% (p=0.002 n=6)
SearchFloat64s   74.02n ± 6%   17.46n ± 9%  -76.40% (p=0.002 n=6)
SearchStrings    139.8n ± 1%   100.5n ± 0%  -28.05% (p=0.002 n=6)
geomean          89.33n        30.56n       -65.79%
```

### Go 1.20

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │     sort     │               xsort                │
               │    sec/op    │   sec/op     vs base               │
SearchInts       65.96n ±  7%   16.06n ± 0%  -75.65% (p=0.002 n=6)
SearchFloat64s   73.87n ± 16%   20.54n ± 2%  -72.20% (p=0.002 n=6)
SearchStrings    132.9n ±  2%   100.9n ± 0%  -24.08% (p=0.002 n=6)
geomean          86.52n         32.16n       -62.82%
```

### Go 1.21

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │     sort     │               slices               │                xsort                │
               │    sec/op    │   sec/op     vs base               │    sec/op     vs base               │
SearchInts       60.57n ± 20%   21.05n ± 1%  -65.24% (p=0.002 n=6)   16.30n ± 42%  -73.09% (p=0.002 n=6)
SearchFloat64s   68.92n ± 34%   26.89n ± 1%  -60.98% (p=0.002 n=6)   17.39n ±  1%  -74.76% (p=0.002 n=6)
SearchStrings    132.8n ±  0%   144.3n ± 1%   +8.66% (p=0.002 n=6)   100.5n ±  0%  -24.28% (p=0.002 n=6)
geomean          82.15n         43.39n       -47.18%                 30.55n        -62.81%
```

### Go 1.22

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │     sort     │               slices               │               xsort                │
               │    sec/op    │   sec/op     vs base               │   sec/op     vs base               │
SearchInts       75.00n ±  2%   20.79n ± 2%  -72.28% (p=0.002 n=6)   16.61n ± 0%  -77.85% (p=0.002 n=6)
SearchFloat64s   68.00n ±  1%   26.74n ± 2%  -60.68% (p=0.002 n=6)   18.27n ± 0%  -73.14% (p=0.002 n=6)
SearchStrings    140.2n ± 10%   112.8n ± 1%  -19.51% (p=0.002 n=6)   100.4n ± 0%  -28.42% (p=0.002 n=6)
geomean          89.42n         39.73n       -55.56%                 31.22n       -65.08%
```

### Go 1.23

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │    sort     │               slices               │               xsort                │
               │   sec/op    │   sec/op     vs base               │   sec/op     vs base               │
SearchInts       74.80n ± 0%   20.71n ± 2%  -72.31% (p=0.002 n=6)   16.60n ± 1%  -77.81% (p=0.002 n=6)
SearchFloat64s   67.90n ± 6%   26.55n ± 2%  -60.89% (p=0.002 n=6)   18.27n ± 1%  -73.10% (p=0.002 n=6)
SearchStrings    140.1n ± 1%   112.8n ± 0%  -19.51% (p=0.002 n=6)   100.4n ± 0%  -28.36% (p=0.002 n=6)
geomean          89.29n        39.59n       -55.66%                 31.22n       -65.03%
```

### Go 1.24

```
goos: linux
goarch: amd64
pkg: github.com/mroth/xsort
cpu: AMD EPYC 7763 64-Core Processor                
               │     sort     │               slices               │               xsort                │
               │    sec/op    │   sec/op     vs base               │    sec/op     vs base              │
SearchInts       16.46n ± 36%   20.75n ± 2%        ~ (p=0.063 n=6)   16.23n ± 42%       ~ (p=0.288 n=6)
SearchFloat64s   19.20n ±  3%   26.58n ± 1%  +38.45% (p=0.002 n=6)   18.28n ±  0%  -4.77% (p=0.002 n=6)
SearchStrings    100.3n ±  1%   112.5n ± 0%  +12.11% (p=0.002 n=6)   100.3n ±  1%       ~ (p=0.924 n=6)
geomean          31.65n         39.58n       +25.05%                 30.99n        -2.10%
```


[1]: https://github.com/golang/go/issues/15561
