# Cgo benchmarks

This repository provides a series of benchmarks between Cgo, Go and C.

## Benchmark: CGO vs GO vs C in System Calls `read/write`

In `./syscall`

### Usage

```bash
cd syscall && sh run.sh
```

### Result

```
goos: darwin
goarch: amd64
pkg: github.com/sjtushi/guacamole-go/benchmarks/syscall
BenchmarkEmptyCgoCalls-4        10000000               136 ns/op
BenchmarkEmptyGoCalls-4        200000000              8.66 ns/op
BenchmarkReadWriteCgoCalls-4      300000              4282 ns/op
BenchmarkReadWriteGoCalls-4       500000              3715 ns/op
BenchmarkReadWritePureCCalls      300000              2682 ns/op
PASS
ok      github.com/sjtushi/guacamole-go/benchmarks/syscall      7.126s
```

### Conclusions

- Pure Go call is `(136 - 8.66) / 136 = 93.6%` faster than Cgo call. 
- Pure Go system call is `(4684 - 3518) / 4684 = 24.8%` faster than Cgo call.
- Pure C system call is `(3715 - 2682) / 3715 = 27.8%` faster than Go system call.
- Pure C system call is `(4684 - 2682) / 3715 = 53.9%` faster than Cgo system call.

### Related researches

- https://github.com/golang/go/issues/19563
- https://github.com/golang/go/issues/19574
- https://github.com/draffensperger/go-interlang

## Benchmark: TODO

## License

MIT &copy; [Changkun Ou](https://changkun.de)