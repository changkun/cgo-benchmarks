# Benchmark: CGO vs GO vs C in System Calls `read/write`

## Run

```bash
sh run.sh
```

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/changkun/cgo-benchmarks/syscall
BenchmarkReadWriteCgoCalls-4      500000              3215 ns/op
BenchmarkReadWriteGoCalls-4       500000              2781 ns/op
BenchmarkReadWritePureCCalls      500000              2348 ns/op
PASS
ok      github.com/changkun/cgo-benchmarks/syscall      3.085s
```

## Conclusions

- Pure Go system call is `(3215 - 2781) / 2781 = 15.61%` faster than Cgo call.
- Pure C system call is `(2781 - 2348) / 2348 = 18.44%` faster than Go system call.
- Pure C system call is `(3215 - 2348) / 2348 = 36.93%` faster than Cgo system call.

## Related researches

- https://github.com/golang/go/issues/19563
- https://github.com/golang/go/issues/19574

## License

MIT &copy; [Changkun Ou](https://changkun.de)