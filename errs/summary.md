# Benchmark Results Summary

Benchmarks were conducted on an Apple M2 Pro (darwin/arm64) to evaluate the performance of different error handling patterns in Go.

| Benchmark                         | Iterations  | Time (ns/op) | Memory (B/op) | Allocs (op) |
| :-------------------------------- | :---------- | :----------- | :------------ | :---------- |
| `BenchmarkFmt`                    | 7,687,845   | 135.8        | 144           | 4           |
| `BenchmarkErrorsNew`              | 23,473,010  | 50.33        | 112           | 2           |
| `BenchmarkStringError`            | 23,489,494  | 50.84        | 112           | 2           |
| `BenchmarkStringErrorNoInterface` | 32,902,854  | 36.31        | 96            | 1           |
| `BenchmarkIntError`               | 391,720,922 | 3.045        | 0             | 0           |
| `BenchmarkIntErrorNoInterface`    | 606,283,622 | 1.966        | 0             | 0           |

## Analysis

1.  **Integers are King**: Using `uint8` constants (sentinel errors) is by far the most efficient method, resulting in zero allocations and near-instant execution.
2.  **Interface Overhead**: Returning specific types instead of the `error` interface (e.g., `BenchmarkStringErrorNoInterface` vs `BenchmarkStringError`) consistently provides a performance boost by avoiding interface boxing.
3.  **String Concatenation vs. Formatting**: `errors.New` with string concatenation is significantly faster (~2.7x) and uses fewer allocations than `fmt.Errorf`.
4.  **Custom String Types**: Implementing the `error` interface on a custom string type (`StringError`) performs almost identically to `errors.New` when string concatenation is involved.
