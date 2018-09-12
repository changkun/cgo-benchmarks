#!/bin/bash

gcc -O2 c/syscall_bench.c -o syscall_pure_c
go test -bench=.
./syscall_pure_c
rm syscall_pure_c