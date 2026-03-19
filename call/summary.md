# Benchmark Summary: "Do" Pattern vs Standard Function

## Benchmark Results

Running the benchmarks in `call_test.go` on an Apple M2 Pro architecture yields the following results:

```text
goos: darwin
goarch: arm64
pkg: github.com/DanielvNiek/goperfm/call
cpu: Apple M2 Pro
BenchmarkFunction-10    39466600        30.10 ns/op       40 B/op        2 allocs/op
BenchmarkDo-10          39562232        29.83 ns/op       40 B/op        2 allocs/op
```

## Analysis

As the benchmark clearly demonstrates, there is **zero meaningful performance difference** between the two approaches.

### 1. Speed
The standard function (`BenchmarkFunction-10`) runs at roughly **30.10 ns/op**, while the Method Object / "Do" pattern (`BenchmarkDo-10`) runs at **29.83 ns/op**. These differences are entirely within the expected margin of error for Go microbenchmarks.

### 2. Memory Allocations
Both patterns trigger exactly **2 allocations** for a total of **40 bytes** per operation. 
* The first allocation is likely the struct instance itself escaping to the heap (since it is added to the global `users` slice indirectly, causing the compiler's escape analysis to push it to the heap). 
* The second allocation is the internal `*User` struct being created when appending.

### 3. Under the Hood
In Go, a method on a struct is essentially just syntactic sugar for a regular function where the struct is passed implicitly as the first argument.

```go
// Syntactically different
func insertUser(req *insertUserRequest) int64
func (i *InsertUser) Do() int64

// But conceptually identical to the compiler
func insertUser(req *insertUserRequest) int64
func InsertUser_Do(i *InsertUser) int64
```

Because of this, the Go compiler optimizes both patterns into practically identical machine code.

## Conclusion

From a performance perspective, the "Do" pattern (Method Object pattern) is completely free compared to the standard request-object pattern. You can confidently adopt the "Do" pattern to improve the readability, modularity, and cleanliness of complex business logic functions without worrying about adding any overhead to your Go applications.