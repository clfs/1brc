# obrc

An unoptimized, 2m9s [One Billion Row Challenge](https://www.morling.dev/blog/one-billion-row-challenge/) solution in Go.

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
BenchmarkGenerateCSV1e3-10       	    5095	    234506 ns/op
BenchmarkGenerateCSV1e4-10       	     512	   2333130 ns/op
BenchmarkGenerateCSV1e5-10       	      50	  23381521 ns/op
BenchmarkGenerateCSV1e6-10       	       5	 232976692 ns/op
BenchmarkTakeRecordings1e3-10    	    7417	    148938 ns/op
BenchmarkTakeRecordings1e4-10    	     903	   1321701 ns/op
BenchmarkTakeRecordings1e5-10    	      93	  12592483 ns/op
BenchmarkTakeRecordings1e6-10    	       9	 124397972 ns/op
PASS
ok  	github.com/clfs/obrc	13.300s
```

```text
$ go run ./cmd/gen -n 1_000_000_000 > test.csv
$ time go run ./cmd/stats -input test.csv
...
... 126.41s user 3.52s system 100% cpu 2:09.61 total
```