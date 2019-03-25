package prt

type Mode uint32
type FrontEndColor uint32
type BackEndColor uint32

const (
	Default Mode = iota
	HighLight
	Light
	Italic
	UnderLine
	Flashing
	Normal
	Invert
	Invisible
)

const (
	Black FrontEndColor = 30 + iota
	Red
	Green
	Yellow
	Blue
	Fuchsia
	Azure
	White
)

const (
	Black_BackEnd BackEndColor = 40 + iota
	Red_BackEnd
	Green_BackEnd
	Yellow_BackEnd
	Blue_BackEnd
	Fuchsia_BackEnd
	Azure_BackEnd
	White_BackEnd
)
