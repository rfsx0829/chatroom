package refac

import (
	"net/http"

	"go.uber.org/dig"
)

type need struct{}

// Run Entry
func Run() {
	con := dig.New()
	con.Provide(NewPlat)
	con.Provide(NewController)
	con.Provide(router)

	if err := con.Invoke(run); err != nil {
		panic(err)
	}
}

func run(n need) error {
	return http.ListenAndServe(":8089", nil)
}

func router(c *Controller) need {
	http.HandleFunc("/au", c.AddUser)
	http.HandleFunc("/ac", c.AddConn)
	http.HandleFunc("/du", c.DelUser)
	http.HandleFunc("/cr", c.Create)
	http.HandleFunc("/dr", c.DelRoom)
	http.HandleFunc("/er", c.Enter)
	http.HandleFunc("/lr", c.Leave)

	return need{}
}
