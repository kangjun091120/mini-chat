package user

import {
	uuid "github.com/satori/go.uuid"
}

type User struct {
	Name   string
	Id int
	Message []byte
}
