package main

import (
	"log"
	"net/http"
	"time"

	"github.com/rfsx0829/chatroom/server/controller"
	"github.com/rfsx0829/chatroom/server/mysql"
)

func main() {
	source := "root:123456@tcp(192.168.99.104:31255)/wow"

	err := mysql.InitDefault(source)
	count := 0
	for err != nil && count < 5 {
		time.Sleep(time.Second * 5)
		err = mysql.InitDefault(source)
		count++
	}

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/main", controller.Handler)

	log.Println("Listen on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
