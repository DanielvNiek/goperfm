// Package stringbuilding benchmarks different ways of building strings
package stringbuilding

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	id    = "asfd"
	email = "test@example.com"
)

func BenchmarkFmt(b *testing.B) {
	for b.Loop() {
		s := fmt.Sprintf("The user (%s) email is %s", id, email)
		if s != "The user (asfd) email is test@example.com" {
			b.Fail()
		}
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for b.Loop() {
		sb := &strings.Builder{}
		sb.WriteString("The user (")
		sb.WriteString(id)
		sb.WriteString(") email is ")
		sb.WriteString(email)
		if sb.String() != "The user (asfd) email is test@example.com" {
			b.Fail()
		}
	}
}

func BenchmarkBytesBuilder(b *testing.B) {
	for b.Loop() {
		buf := bytes.NewBuffer(nil)
		buf.WriteString("The user (")
		buf.WriteString(id)
		buf.WriteString(") email is ")
		buf.WriteString(email)
		if buf.String() != "The user (asfd) email is test@example.com" {
			b.Fail()
		}
	}
}

func BenchmarkAddOperator(b *testing.B) {
	for b.Loop() {
		s := "The user (" + id + ") email is " + email
		if s != "The user (asfd) email is test@example.com" {
			b.Fail()
		}
	}
}
