package service

type Response struct {
	Status int         `json:"status"`
	Text   string      `json:"text"`
	Extra  interface{} `json:"extra"`
}
