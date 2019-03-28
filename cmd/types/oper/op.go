package oper

type Op uint32

const (
	DefaultOper Op = iota
	SignUp
	SignIn
	AddEmail
	Create
	Enter
	Leave
	SendMes
	SendBox
	GetRoomList
	GetPersonsInRoom
	Close
)
