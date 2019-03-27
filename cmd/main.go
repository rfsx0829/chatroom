package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rfsx0829/chatroom/cmd/prt/curse"
	"github.com/rfsx0829/chatroom/cmd/types"
)

type App struct {
	user    *types.Person
	conn    *types.Conn
	graphic *Graphic
}

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

func main() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var g Graphic

	err = json.Unmarshal(data, &g)
	if err != nil {
		panic(err)
	}

	g.BasicFrame()
}

func main_bak() {
	curse.MoveUp(12)
	var (
		regist = flag.Bool("add", false, "use this flag to regist an acount")
		fdd    = types.FormData{}
	)
	flag.Parse()

	u := &url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/api/main",
	}

	con, err := types.NewConn(u)
	if err != nil {
		panic(err)
	}
	defer con.Close(&fdd)

	if *regist {
		err = register(con)
	} else {
		err = MainLoop()
	}

	if err != nil {
		panic(err)
	}
}

func MainLoop() error {
	return nil
}

func register(con types.Conn) error {
	fd := &types.FormData{
		Oper: types.SignUp,
	}

	fmt.Print("Please Input Your Username:")
	fmt.Scanf("%s", &fd.User.Name)
	fmt.Print("Please Input Your Password:")
	fmt.Scanf("%s", &fd.User.Pass)

	res, err := con.WriteAndRead(fd)
	if err != nil {
		return err
	}

	if res.Status != http.StatusOK {
		return errors.New(res.Text)
	}

	fmt.Println(res.Text)
	return nil
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

func CenterStr(str string, length int) string {
	l := len(str)
	if l&1 != length&1 {
		str += " "
		l++
	}
	return fmt.Sprintf("%*s%s%*s", (length-l)/2, "", str, (length-l)/2, "")
}

func StrsLine(div string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}

	str := div
	for i := 0; i < len(strs); i++ {
		str += strs[i] + div
	}
	return str
}

func StrRepeat(single string, times int) string {
	str := ""
	for times > 0 {
		times--
		str += single
	}
	return str
}
