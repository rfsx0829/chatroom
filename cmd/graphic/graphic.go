package graphic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Graphic struct {
	Size struct {
		Length int `json:"length"`
		Width  int `json:"width"`
		Llen   int `json:"llen"`
		Rlen   int `json:"rlen"`
		H1     int `json:"h1"`
		H2     int `json:"h2"`
		H3     int `json:"h3"`
	} `json:"size"`
}

func NewGraphic() (*Graphic, error) {
	var g Graphic
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *Graphic) BasicFrame() {
	div1 := StrRepeat("-", g.Size.Llen)
	div2 := StrRepeat("-", g.Size.Length-g.Size.Llen-g.Size.Rlen)
	div3 := StrRepeat("-", g.Size.Rlen)
	space1 := StrRepeat(" ", g.Size.Llen)
	space2 := StrRepeat(" ", g.Size.Length-g.Size.Llen-g.Size.Rlen)
	space3 := StrRepeat(" ", g.Size.Rlen)

	divLine := StrsLine("+", div1, div2, div3)
	emptyLine := StrsLine("|", space1, space2, space3)

	fmt.Println(divLine)
	for i := 0; i < g.Size.H1; i++ {
		fmt.Println(emptyLine)
	}
	fmt.Println(divLine)
	for i := 0; i < g.Size.H2; i++ {
		fmt.Println(emptyLine)
	}
	fmt.Println(StrsLine("+", div1, space2, div3))
	for i := 0; i < g.Size.Width-g.Size.H1-g.Size.H2-g.Size.H3; i++ {
		fmt.Println(emptyLine)
	}
	fmt.Println(StrsLine("+", space1, div2, space3))
	for i := 0; i < g.Size.H3; i++ {
		fmt.Println(emptyLine)
	}
	fmt.Println(divLine)
}
