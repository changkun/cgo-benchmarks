#!/bin/bash

gcc -O2 c/emptycall_bench.c -o empty_c
go test -bench=. -count=1000

for i in `seq 1 1000`; do
    ./empty_c
done
rm empty_c