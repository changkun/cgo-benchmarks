#!/bin/bash

gcc -O2 c/emptycall_bench.c -o empty_c
go test -bench=.
./empty_c
rm empty_c