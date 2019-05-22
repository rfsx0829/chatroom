package main

import (
	"log"

	"github.com/rfsx0829/chatroom/server/refac"
	"go.uber.org/dig"
)

func main() {
	con := dig.New()

	con.Provide(refac.NewPlat)
	con.Provide(refac.NewController)

	if err := con.Invoke(refac.EntryPoint); err != nil {
		log.Println(err)
	}
}
