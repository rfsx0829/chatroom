package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rfsx0829/chatroom/server/service"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.SignUp...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println("Controller.SignUp: ", x)

	err = service.Client.SignUp(x.Name, x.Pass)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SignUp OK !"))
	log.Println("SignUp OK !")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("Controller.SignIn...")

	if checkMethod(w, r, "POST") {
		log.Println("NotPost")
		return
	}

	var x struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	err := unmarshalBody(r, &x)
	if checkError(w, r, err) {
		log.Println(err)
		return
	}

	log.Println("Controller.SignIn: ", x)

	u, err := service.Client.SignInWithName(w, r, x.Name, x.Pass)
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

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only Receive Post !"))
		return true
	}
	return false
}

func checkError(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return true
	}
	return false
}

func unmarshalBody(r *http.Request, x interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, x)
}
