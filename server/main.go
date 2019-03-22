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

	http.HandleFunc("/api/user/signup", controller.SignUp)
	http.HandleFunc("/api/user/signin", controller.SignIn)
	http.HandleFunc("/api/email/add", controller.AddEmail)
	http.HandleFunc("/api/email/signin", controller.EmailSignIn)
	http.HandleFunc("/api/room/create", controller.CreateRoom)
	http.HandleFunc("/api/room/enter", controller.EnterRoom)
	http.HandleFunc("/api/room/leave", controller.LeaveRoom)
	http.HandleFunc("/api/send/box", controller.SendToBox)
	http.HandleFunc("/api/send/mes", controller.SendMes)

	log.Println("Listen on http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", nil))
}

/*

"/api/user/signup"      ->  signup
{
	"name": "username",
	"pass": "password"
}

"/api/user/signin"      ->  signin
{
	"name": "username",
	"pass": "password"
}

"/api/email/signin"   ->  emailSignin
{
	"email": "email@email.com",
	"pass": "password"
}

"/api/email/add"    ->  addEmail
{
	"uid": 1234,
	"email": "email@email.com"
}

"/api/room/create"  ->  createRoom
{
	"uid": 1234,

	"rname": "badroom",
	"rpass": "password",
	"pri": "true"
}

"/api/room/enter"   ->  enterRoom
{
	"uid": 1234,
	"rid": 1001,
	"rpass": "password"
}

"/api/room/leave"   ->  leaveRoom
{
	"uid": 1234,
}

"/api/send/mes"     ->  sendMes
{
	"uid": 1234,
	"mes": "Hello World"
}

"/api/send/box"   ->  sendToBox
{
	"uid": 1234,
	"sendto": 1235,
	"mes": "Hello World"
}

*/
