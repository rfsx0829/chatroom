package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/mysql"

	"github.com/rfsx0829/chatroom/server/service"
)

func EmailSignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.EmailSignIn...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Email string `json:"email"`
		Pass  string `json:"pass"`
	}
	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println("Controller.SignIn: ", x)

	u, err := service.Client.SignInWithEmail(w, r, x.Email, x.Pass)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	data, _ := json.Marshal(struct {
		Uid int `json:"uid"`
	}{
		Uid: u,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(data)
	log.Printf("uid = %d , SignIn OK !", u)
}

func AddEmail(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.AddEmail...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Uid   int    `json:"uid"`
		Email string `json:"email"`
	}

	log.Println("In AddEmail:", x)

	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	err = mysql.Default.AddEmail(x.Uid, x.Email)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Add Email OK !"))
	log.Println("AddEmail OK !")
}
