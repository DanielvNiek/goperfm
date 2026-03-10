package new_test

import (
	"testing"

	"goperfm/new"
)

func BenchmarkNew(b *testing.B) {
	for b.Loop() {
		p := new.New("test")
		if p.Name() != "test" {
			b.Errorf("Expected name to be test, got %s", p.Name())
		}
	}
}

func BenchmarkInit(b *testing.B) {
	for b.Loop() {
		p := (&new.Perfm{}).Init("test2")
		if p.Name() != "test2" {
			b.Errorf("Expected name to be test2, got %s", p.Name())
		}
	}
}
