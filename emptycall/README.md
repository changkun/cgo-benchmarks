# Benchmark: CGO vs GO vs C in Empty Calls

## Run

```bash
sh run.sh
```

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/changkun/cgo-benchmarks/emptycall
BenchmarkEmptyCgoCalls-4        20000000                 55.9 ns/op
BenchmarkEmptyGoCalls-4         2000000000               0.29 ns/op
BenchmarkEmptyCCalls            2000000000                 00 ns/op
PASS
ok      github.com/changkun/cgo-benchmarks/emptycall    1.807s
```

## Conclusions

- Pure Go call is `(55.9 - 0.29) / 0.29 = 191.76%` faster than Cgo call.
- Pure C call is `(0.29 - 0.00) / 0.00 = infinity` faster than Go call.
- Pure C call is `(55.9 - 0.00) / 0.00 = higher infinity` faster than Cgo call.

## Related researches

- https://github.com/draffensperger/go-interlang

## License

MIT &copy; [Changkun Ou](https://changkun.de)