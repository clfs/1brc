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
BenchmarkGenerateCSV1e3-10       	    5054	    237691 ns/op
BenchmarkGenerateCSV1e4-10       	     506	   2380981 ns/op
BenchmarkGenerateCSV1e5-10       	      50	  23524732 ns/op
BenchmarkGenerateCSV1e6-10       	       5	 232802217 ns/op
BenchmarkTakeRecordings1e3-10    	    6760	    175922 ns/op
BenchmarkTakeRecordings1e4-10    	     758	   1584696 ns/op
BenchmarkTakeRecordings1e5-10    	      79	  15280674 ns/op
BenchmarkTakeRecordings1e6-10    	       7	 150463435 ns/op
PASS
ok  	github.com/clfs/obrc	12.240s
```

```text
$ go run ./cmd/gen -n 1_000_000_000 > test.csv
$ time go run ./cmd/stats -f test.csv
...
... 149.98s user 4.32s system 100% cpu 2:33.41 total
```