package emptycall_test

import (
	"testing"

	"github.com/changkun/cgo-benchmarks/emptycall"
)

func BenchmarkEmptyCgoCalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptycall.Cempty()
	}
}

func BenchmarkEmptyGoCalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		emptycall.Empty()
	}
}
