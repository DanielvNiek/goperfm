package call

import (
	"crypto/rand"
	"testing"
	"time"
)

func insertUserFunc(name string, age int32, now time.Time) int64 {
	id := now.Unix()
	users = append(users, &User{
		id:   id,
		name: name,
		age:  age,
	})
	users = nil
	return id
}

type User struct {
	id   int64
	name string
	age  int32
}

var users = []*User{}

type insertUserRequest struct {
	name string
	age  int32
	now  time.Time
}

func insertUserReq(req *insertUserRequest) int64 {
	id := req.now.Unix()
	users = append(users, &User{
		id:   id,
		name: req.name,
		age:  req.age,
	})
	users = nil
	return id
}

type InsertUser struct {
	name string
	age  int32
	now  time.Time
}

func (i *InsertUser) Do() int64 {
	id := i.now.Unix()
	users = append(users, &User{
		id:   id,
		name: i.name,
		age:  i.age,
	})
	users = nil
	return id
}

func BenchmarkFunc(b *testing.B) {
	now := time.Now()
	for b.Loop() {
		b := make([]byte, 2048*4)
		rand.Read(b)
		name := string(b)
		i := insertUserFunc(name, 28, now)
		_ = i
	}
}

func BenchmarkReq(b *testing.B) {
	now := time.Now()
	for b.Loop() {
		b := make([]byte, 2048*4)
		rand.Read(b)
		name := string(b)
		i := insertUserReq(&insertUserRequest{
			name: name,
			age:  28,
			now:  now,
		})
		_ = i
	}
}

func BenchmarkDo(b *testing.B) {
	now := time.Now()
	for b.Loop() {
		b := make([]byte, 2048*4)
		rand.Read(b)
		name := string(b)
		i := (&InsertUser{
			name: name,
			age:  28,
			now:  now,
		}).Do()
		_ = i
	}
}
