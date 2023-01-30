package Router

import (
	"log"
	login "lorryservice/Login"
	"lorryservice/SignUp"
	"net/http"
)

func Route() {
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/signup", SignUp.SignUp)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
