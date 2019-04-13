package service

type Response struct {
	Status int         `json:"status"`
	Oper   int         `json:"oper"`
	Text   string      `json:"text"`
	Extra  interface{} `json:"extra"`
}
