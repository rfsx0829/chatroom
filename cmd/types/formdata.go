package types

type FormData struct {
	Oper Op       `json:"oper"`
	User UserInfo `json:"user"`
	Room Room     `json:"room"`
	Mes  Message  `json:"mes"`
}
