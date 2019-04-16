package main

import (
	"log"
	"net/http"

	"go.uber.org/dig"

	"github.com/rfsx0829/chatroom/server/controller"
	"github.com/rfsx0829/chatroom/server/mysql"
)

func main() {
	container := dig.New()
	container.Provide(func() string {
		source := "root:123456@tcp(192.168.99.104:31255)/wow"
		return source
	})

	err := container.Provide(func(source string) error {
		return mysql.InitDefault(source)
	})

	if err != nil {
		panic(err)
	}

	container.Invoke(func() {
		http.HandleFunc("/api/main", controller.Handler)
		log.Println("Listen on http://localhost:8080")
		log.Println(http.ListenAndServe(":8080", nil))
	})
}
