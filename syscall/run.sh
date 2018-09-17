#!/bin/bash
gcc -O2 c/syscall_bench.c -o syscall_pure_c
go test -bench=. -count=5 -timeout 20m -v
for i in `seq 1 5`; do
    ./syscall_pure_c
done
rm syscall_pure_c