package service

import "errors"

var (
	NameExists   = errors.New("Name exists !")
	InvalidPass  = errors.New("Invalid password !")
	InvalidUid   = errors.New("Invalid user id !")
	InvalidRid   = errors.New("Room not exists !")
	NameTooShort = errors.New("Name should longer than 3")
	PassTooShort = errors.New("Pass should longer than 4")
)
