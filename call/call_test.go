package call

import (
	"testing"
	"time"
)

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

func insertUser(req *insertUserRequest) int64 {
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

func BenchmarkFunction(b *testing.B) {
	now := time.Now()
	for b.Loop() {
		i := insertUser(&insertUserRequest{
			name: "daniel",
			age:  28,
			now:  now,
		})
		_ = i
	}
}

func BenchmarkDo(b *testing.B) {
	now := time.Now()
	for b.Loop() {
		i := (&InsertUser{
			name: "daniel",
			age:  28,
			now:  now,
		}).Do()
		_ = i
	}
}
