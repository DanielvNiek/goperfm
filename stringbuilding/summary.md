# String Building Benchmarks

Benchmarks comparing different ways to build strings in Go.

## Results

```text
goos: darwin
goarch: arm64
pkg: goperfm/stringbuilding
cpu: Apple M2 Pro
BenchmarkFmt-10              	11887582	        84.85 ns/op	      80 B/op	       3 allocs/op
BenchmarkStringBuilder-10    	21320209	        55.47 ns/op	     112 B/op	       3 allocs/op
BenchmarkBytesBuilder-10     	26793464	        44.03 ns/op	     112 B/op	       2 allocs/op
BenchmarkAddOperator-10      	42216543	        28.54 ns/op	      48 B/op	       1 allocs/op
```

## Summary

- **BenchmarkAddOperator** is the fastest and most memory-efficient method for this simple case.
- **BenchmarkBytesBuilder** and **BenchmarkStringBuilder** follow, with `bytes.Buffer` showing slightly better performance than `strings.Builder`.
- **BenchmarkFmt** is the slowest and has more overhead due to the use of `fmt.Sprintf`.
