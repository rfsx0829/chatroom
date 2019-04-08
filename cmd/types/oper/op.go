package oper

// Op is the shortcut of Operation
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
