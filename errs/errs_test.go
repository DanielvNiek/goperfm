package errs_test

import (
	"errors"
	"fmt"
	"testing"
)

var (
	id    = "asfd"
	email = "test@example.com"
)

func BenchmarkFmt(b *testing.B) {
	for b.Loop() {
		err := doFmt()
		if err.Error() != "Error reading from postgres database. Expected user asfd to have email test@example.com" {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doFmt() error {
	return fmt.Errorf("Error reading from postgres database. Expected user %s to have email %s", id, email)
}

func BenchmarkErrorsNew(b *testing.B) {
	for b.Loop() {
		err := doErrorsNew()
		if err.Error() != "Error reading from postgres database. Expected user asfd to have email test@example.com" {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doErrorsNew() error {
	return errors.New("Error reading from postgres database. Expected user " + id + " to have email " + email)
}

type StringError string

func (e StringError) Error() string {
	return string(e)
}

func BenchmarkStringError(b *testing.B) {
	for b.Loop() {
		err := doStringError()
		if err.Error() != "Error reading from postgres database. Expected user asfd to have email test@example.com" {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doStringError() error {
	return StringError("Error reading from postgres database. Expected user " + id + " to have email " + email)
}

func BenchmarkStringErrorNoInterface(b *testing.B) {
	for b.Loop() {
		err := doStringErrorNoInterface()
		if err != "Error reading from postgres database. Expected user asfd to have email test@example.com" {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doStringErrorNoInterface() StringError {
	return StringError("Error reading from postgres database. Expected user " + id + " to have email " + email)
}

type Error uint8

const (
	_               Error = iota
	UnexpectedEmail Error = iota
	UnexpectedID    Error = iota
)

func (e Error) Error() string {
	return fmt.Sprintf("%T", e)
}

func BenchmarkIntError(b *testing.B) {
	for b.Loop() {
		err := doIntError()
		if err != UnexpectedEmail {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doIntError() error {
	return UnexpectedEmail
}

func BenchmarkIntErrorNoInterface(b *testing.B) {
	for b.Loop() {
		err := doIntErrorNoInterface()
		if err != UnexpectedEmail {
			b.Fatal("unexpected error")
		}
	}
}

//go:noinline
func doIntErrorNoInterface() Error {
	return UnexpectedEmail
}
