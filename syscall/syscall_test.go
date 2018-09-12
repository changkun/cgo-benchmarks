package syscall

import (
	"syscall"
	"testing"
)

func writeAll(fd int, buf []byte) error {
	for len(buf) > 0 {
		n, err := syscall.Write(fd, buf)
		if n < 0 {
			return err
		}
		buf = buf[n:]
	}
	return nil
}

func BenchmarkReadWriteCgoCalls(b *testing.B) {
	fds, _ := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	message := "hello, world!"
	buffer := make([]byte, 13)
	for i := 0; i < b.N; i++ {
		CwriteAll(fds[0], []byte(message))
		Cread(fds[1], buffer)
	}
}

func BenchmarkReadWriteGoCalls(b *testing.B) {
	fds, _ := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_STREAM, 0)
	message := "hello, world!"
	buffer := make([]byte, 13)
	for i := 0; i < b.N; i++ {
		writeAll(fds[0], []byte(message))
		syscall.Read(fds[1], buffer)
	}
}
