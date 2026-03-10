# Benchmark Results: New vs Init

The following benchmarks were conducted on an Apple M2 Pro (darwin/arm64) to compare the performance of creating a new instance using a constructor-like function (`New`) versus initializing an existing instance (`Init`).

## Results

| Benchmark       | Iterations  | Time per Op | Memory per Op | Allocs per Op |
| --------------- | ----------- | ----------- | ------------- | ------------- |
| `BenchmarkNew`  | 86,299,579  | 13.84 ns/op | 16 B/op       | 1 allocs/op   |
| `BenchmarkInit` | 597,296,486 | 2.011 ns/op | 0 B/op        | 0 allocs/op   |

## Analysis

- **Performance**: `BenchmarkInit` is approximately **6.8x faster** than `BenchmarkNew`.
- **Memory**: `BenchmarkInit` performs **zero allocations**, while `BenchmarkNew` allocates 16 bytes per operation.
- **Conclusion**: Reusing an existing structure and initializing it via an `Init` method significantly reduces heap allocation overhead and improves throughput in high-frequency scenarios.
