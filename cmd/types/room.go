package types

type Room struct {
	Rid  int    `json:"rid"`
	Name string `json:"rname"`
	Pass string `json:"rpass"`
}
