package str_test

import (
	"bytes"
	"testing"
)

//go:noinline
func String(who string) string {
	return "Hello, " + who
}

func prepBytes(b []byte, capacity int) []byte {
	if cap(b) < capacity {
		b = make([]byte, capacity)
	} else {
		b = b[:]
	}
	return b
}

//go:noinline
func Bytes(b []byte, who []byte) []byte {
	b = prepBytes(b, len("Hello, ")+len(who))
	b = append(b[:0], "Hello, "...)
	b = append(b, who...)
	return b
}

func BenchmarkString(b *testing.B) {
	for b.Loop() {
		x := String("World!")
		_ = x
	}
}

func BenchmarkBytes(b *testing.B) {
	for b.Loop() {
		x := Bytes(make([]byte, 20), []byte("World!"))
		_ = x
	}
}

func TestString(t *testing.T) {
	x := String("World!")
	if x != "Hello, World!" {
		t.Errorf("expected \"Hello, World!\", got %q", x)
	}
}

func TestBytes(t *testing.T) {
	x := Bytes(make([]byte, 20), []byte("World!"))
	if !bytes.Equal(x, []byte("Hello, World!")) {
		t.Errorf("expected \"Hello, World!\", got %q", x)
	}
}
