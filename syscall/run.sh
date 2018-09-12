#!/bin/bash
gcc -O2 c/syscall_bench.c -o syscall_pure_c
go test -bench=. -count=1000 -v
for i in `set 1 1000`; do
    ./syscall_pure_c
done
rm syscall_pure_c