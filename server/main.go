package main

import (
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/controller"
	"github.com/rfsx0829/chatroom/server/plat"
	"github.com/rfsx0829/chatroom/server/redis"
	"go.uber.org/dig"
)

func main() {
	con := dig.New()

	con.Provide(func() *redis.Option {
		return &redis.Option{
			Host: "172.17.0.5",
			Port: 6379,
			Pass: "zdan1214",
			DB:   1,
		}
	})

	con.Provide(redis.InitClient)
	con.Provide(plat.New)
	con.Provide(controller.New)

	if err := con.Invoke(entryPoint); err != nil {
		log.Println(err)
	}
}

func entryPoint(c *controller.Controller) error {
	http.HandleFunc("/au", c.AddUser)
	http.HandleFunc("/ac", c.AddConn)
	http.HandleFunc("/du", c.DelUser)
	http.HandleFunc("/cr", c.CreateRoom)
	http.HandleFunc("/dr", c.DeleteRoom)
	http.HandleFunc("/er", c.EnterRoom)
	http.HandleFunc("/lr", c.LeaveRoom)
	http.HandleFunc("/gr", c.RoomList)

	log.Println("Listening At localhost:8089")
	return http.ListenAndServe(":8089", nil)
}
