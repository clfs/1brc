# obrc

An unoptimized, 2m8s [One Billion Row Challenge](https://www.morling.dev/blog/one-billion-row-challenge/) solution in Go.

(I skipped the output sorting because I'm lazy.)

## Usage

Generate a CSV:

```text
go run ./cmd/gen -n 1000 > test.csv
```

Run the solver:

```text
go run ./cmd/stats -input test.csv
```

## Benchmark

MacBook Pro 14-inch, 2021, M1 Max, 64 GB:

```text
$ go test -bench=. .
goos: darwin
goarch: arm64
pkg: github.com/clfs/obrc
BenchmarkGenerateCSV1e3-10       	    5113	    233530 ns/op
BenchmarkGenerateCSV1e4-10       	     512	   2328699 ns/op
BenchmarkGenerateCSV1e5-10       	      50	  23345009 ns/op
BenchmarkGenerateCSV1e6-10       	       5	 232703967 ns/op
BenchmarkTakeRecordings1e3-10    	    7803	    152262 ns/op
BenchmarkTakeRecordings1e4-10    	     889	   1345091 ns/op
BenchmarkTakeRecordings1e5-10    	      93	  12778112 ns/op
BenchmarkTakeRecordings1e6-10    	       8	 126402234 ns/op
PASS
ok  	github.com/clfs/obrc	12.131s
```

```text
$ go run ./cmd/gen -n 1_000_000_000 > test.csv
$ time go run ./cmd/stats -input test.csv
...
... 125.79s user 3.53s system 100% cpu 2:08.90 total
```