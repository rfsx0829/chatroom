package main

import (
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/controller"
	"github.com/rfsx0829/chatroom/server/mysql"
)

func main() {
	source := "root:123456@tcp(0.0.0.0:8806)/wow"

	err := mysql.InitDefault(source)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/api/main", controller.Handler)

	log.Println("Listen on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
