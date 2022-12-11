package main

import (
	"errors"
	"fmt"
)

const (
	ResourceErrDatabase = "Database"
	ResourceErrRedis    = "Redis"
)

const (
	ResourceErrDatabaseCode = iota + 1
	ResourceErrRedisCode
)

type ResourceErr struct {
	Resource string
	Code     int
	Message  string
}

func NewDatabaseErr(msg string) ResourceErr {
	return ResourceErr{
		Resource: ResourceErrDatabase,
		Code: ResourceErrDatabaseCode,
		Message: msg,
	}
}

func NewRedisErr(msg string) ResourceErr {
	return ResourceErr{
		Resource: ResourceErrRedis,
		Code: ResourceErrRedisCode,
		Message: msg,
	}
}

func (r ResourceErr) Error() string {
	return r.Message
}

func (r ResourceErr) Is(target error) bool {
	if other, ok := target.(ResourceErr); ok {
		ignoreResource := other.Resource == ""
		ignoreCode := other.Code == 0
		matchResource := other.Resource == r.Resource
		matchCode := other.Code == r.Code
		return matchResource && matchCode ||
			matchResource && ignoreCode ||
			ignoreResource && matchCode
	}
	return false
}

func main() {
	err := NewDatabaseErr("Couldn't connect to the database")
	if errors.Is(err, ResourceErr{Resource: ResourceErrDatabase}) {
		fmt.Println("The database is broken:", err)
	}

	err = NewRedisErr("Couldn't connect to the redis")
	if errors.Is(err, ResourceErr{Resource: ResourceErrRedis}) {
		fmt.Println("The redis is broken:", err)
	}

	var myErr ResourceErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Resource, myErr.Code)
	}
}
