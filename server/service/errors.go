package service

import "errors"

var (
	NameExists  = errors.New("Name exists !")
	InvalidPass = errors.New("Invalid password !")
	InvalidUid  = errors.New("Invalid user id !")
	InvalidRid  = errors.New("Room not exists !")
)
