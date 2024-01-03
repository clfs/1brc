# obrc

An unoptimized, ~155 sec. [One Billion Row Challenge](https://www.morling.dev/blog/one-billion-row-challenge/) solution in Go.

(I skipped the output sorting because I'm lazy.)

## Usage

Generate a CSV:

```text
go run ./cmd/gen -n 1000 > test.csv
```

Run the solver:

```text
go run ./cmd/stats -f test.csv
```

## Benchmark

MacBook Pro 14-inch, 2021, M1 Max, 64 GB:

```text
$ go test -bench=. .
goos: darwin
goarch: arm64
pkg: github.com/clfs/obrc
BenchmarkGenerateCSV1e3-10          4960            233921 ns/op
BenchmarkGenerateCSV1e4-10           512           2340646 ns/op
BenchmarkGenerateCSV1e5-10            49          23557021 ns/op
BenchmarkGenerateCSV1e6-10             5         238529000 ns/op
BenchmarkComputeStats1e3-10         6564            175773 ns/op
BenchmarkComputeStats1e4-10          747           1591551 ns/op
BenchmarkComputeStats1e5-10           76          15301464 ns/op
BenchmarkComputeStats1e6-10            7         150111482 ns/op
PASS
ok      github.com/clfs/obrc    12.140s
```

```text
$ go run ./cmd/gen -n 1_000_000_000 > test.csv
$ time go run ./cmd/stats -f test.csv
...
... 150.54s user 3.88s system 100% cpu 2:33.27 total
```