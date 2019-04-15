package platform

import "time"

type messageType int

// Types
const (
	TextType messageType = iota
	PictureType
)

type message struct {
	From int         `json:"from"`
	To   int         `json:"to"`
	Type messageType `json:"type"`
	Time time.Time   `json:"time"`
}
