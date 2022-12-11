package main

import (
	"fmt"
	"reflect"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status Status
	Message string
}

func (s StatusErr) Error() string {
	return s.Message
}

func (me StatusErr) Is(target error) bool {
	if me2, ok := target.(StatusErr); ok {
		return reflect.DeepEqual(me, me2)
	}
	return false
}

func login(uid, pwd string) error {
	return nil
}

func get(file string) ([]byte, error) {
	return []byte{}, nil
}

func LoginAndGet(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status: InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user: %s", uid),
		}
	}

	data, err := get(file)
	if err != nil {
		return nil, StatusErr{
			Status: NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}

	return data, nil
}

func main() {
	
}